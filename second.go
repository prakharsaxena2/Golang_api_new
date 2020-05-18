package main
 import (
	 "fmt" 
	
	"github.com/gin-gonic/gin"
	
)

func HomePage(c*gin.Context){

	c.JSON(200,gin.H{
		"message": "Hello World wqwqwqwqwewe",
	})
 }

 func PostHomePage(c*gin.Context){

	c.JSON(200,gin.H{
		"message": "Hello World Post homepage",
	})
 }


 func QueryString(c*gin.Context){
	 name := c.Query("name")
	 age := c.Query("age")


	c.JSON(200,gin.H{
		"name":name,
		"age":age,
	})
 }

 func PathParameters(c*gin.Context){
	name := c.Param("name")
	age := c.Param("age")


   c.JSON(200,gin.H{
	   "name":name,
	   "age":age,
   })
}
// func BodyParameter(c*gin.Context){
	
// 	body:= c.Request.Body
// 	value ,err :=ioutil.ReadAll(body)
// 	if err!=nil {
// 		fmt.Println(err.Error())
// 	} 


//    c.JSON(200,gin.H{
// 	   "message":string(body),
//    })
// }
 func main(){

	 fmt.Println("hello world")
	 r:=gin.Default()
	//  r.GET("/",func(c*gin.Context){

	// 	c.JSON(200,gin.H{
	// 		"message": "Hello World edfdfdddddddddddddddddddddddddddddddddf",
	// 	})

	//  })

	r.GET("/",HomePage)
	r.POST("/",PostHomePage)
	r.GET("/query",QueryString)//query?name=prakhar&gae=22
	r.GET("/path/:name/:age",PathParameters)//path/prakhar/22
	// r.GET("/bodyaccess",BodyParameter)

	r.Run()
	}