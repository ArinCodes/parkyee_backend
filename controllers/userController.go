package controllers

import(
"context"
"fmt"
"log"
"strconv"
"net/http"
"time"
"github.com/gin-gonic/gin"
"github.com/go-playground/validator/v10"
helper "parkyee_backend/helpers"
"parkyee_backend/models"
"parkyee_backend/helpers"
"golang.org/x/crypto/bcrypt"

"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection =  database.OpenCollection(database.Client,"user")
var validate = validator.New()

func HashPassword()

func VerifyPassword(userPassword string ,providePassword string)(bool ,string){
  err :=	bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
  check := true
 msg := ""

 if err!=nil{
	msg =fmt.Sprintf("email or password is incorrect")
	check = false
 }return check,msg

}

func Signup() gin.HandlerFunc{
  return func( c *gin.Context){
	var ctx,cancel =content.WithTimeout(content.Background(),100*time.Second)
	var user models.User

	if err =c.BindJSON(&user); err !=nil{
		c.JSON(htp.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	validationErr := validate.Struct(user)
	if validationErr !=nil{
		c.JSON(http.StatusBadRequest , gin.H{"error": validationErr.Error()})
		return
	}
	count , err :=userCollection.CountDocuments(ctx ,bson.M{"email":user.Email})
	defer cancel()
	if err!=nil{
		log.Panic(err)
		c.JSON(hrrp.StatusInternalServerError,gin.H{"error":"error occured while checking for the"})
	}
	count , err := userCollection.CountDocuments(ctx , bson.M{"phone": user.Phone})
	defer cancel()
	if err! =nil{
		log.Panic(err)
		c.JSON(hrrp.StatusInternalServerError,gin.H{"error":"error occured while checking for the"})
		
  }
  if count>0{
	c.JSON(http.StatusInternalServerError,gin.H{"error":"this email or phone number already exist"})
  }
  user.Created_at, _ = time.Parse(time.RFC3339 ,time.Now().Format(time.RFC3339))
  user.Updated_at, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
  user.ID= primitive.NewObjectID()
  user.User_id =user.ID.Hex()
  token,refreshToken,_=helper.GenerateAllTokens(*user.Email,*user.First_name, *user.Last_name,*user.User_type, *&user.User_id)
  user.Refresh_token =&refreshToken

  resultInsertionNumber , insertErr := userCollection.InsertOne(ctx ,user)
  if insertErr !=nil{
	msg :=fmt.Sprintf("User item was not created")
	c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
	return
  }
  defer cancel()
  c.JSON(http.StatusOk , resultInsertionNumber)



}
}

func Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		var foundUser models.User

        if err := c.BindJSON(&user); err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
			return
		}

     err:=  userCollection.FindOne(ctx,bson.M{"email":user.Email}).Decode(&foundUser)
	 defer cancel()
	 if err !=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"error":"email or password tis incorect"})
		return
	 }

	 passwordIsValid,msg:=VerifyPassword(*user.Password, *foundUser.Password)
	 defer cancel()
	}
}

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func( c *gin.Context){
userId := c.Param("user_id")

helper.MatchUserTypeToUid(c,userId); err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return
}
var ctx , cancel =context.WithTimeout(context.Background(),100*time.Second)

var user models.User
err := userCollection.FindOne(ctx , bson.M{"user_id" : userId}).Decode(&user)
	defer cancel()
	if err !=nil{
		c.JSON(http.StatusInternalServerError ,gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOk ,user)

}
}

