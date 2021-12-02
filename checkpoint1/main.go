package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
// 	req, err := httptool.NewRequest(
// 		"",
// 		"",
// 		"",
// 		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(req)

	// write your code below
	// ...
	client := &http.Client{}
	req, _ := http.NewRequest("GET","http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code",nil)
	req.Header.Add("Code","120")			
	res, _ := client.Do(req)				//Do发送请求返回http回复
	by, _ := ioutil.ReadAll(res.Body)		//从res.body读取数据直到遇到error或EOF
	fmt.Println(string(by))	
}
