package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "strings"
        "time"
        "strconv"
        "math/rand"
        "os"
        "crypto/md5"
        "encoding/hex"
        )

var file_path = "images/"
var dir_name = "images"

func Make_dir(dir_name string){
    err := os.Mkdir(dir_name, 0777)
    if err != nil {
    fmt.Println("mkdir root failed!")
    return
    // 建一个文件夹用来保存图片
   }
}

func Get_user_agent()(user_agent_list map[int]string){
    var user_agent map[int]string
    user_agent = make(map[int]string)
    user_agent[0] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[1] = "Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11"
    user_agent[2] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[3] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[4] = "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24"
    user_agent[5] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[6] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[7] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[8] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
    user_agent[9] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36(KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
    user_agent[10] = "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1"
    user_agent[11] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[12] = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[13] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[14] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24"
    user_agent[15] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1"
    user_agent[16] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6"
    user_agent[17] = "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5"
    user_agent[18] = "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[19] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1"
    user_agent[20] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
    user_agent[21] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5"
    user_agent[22] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[23] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3"
    user_agent_list = user_agent
    return
}

func Choose_agent() (user_agent string){
    var agent_list map[int]string = Get_user_agent()
    var random_index int = rand.Intn(24)
    user_agent = agent_list[random_index]
    return
}

func Get_page_content(url string)(content string){
    client := &http.Client{}
    request,err_1 := http.NewRequest("GET", url, nil)
    if err_1 != nil {
        fmt.Println(err_1)
        content = "0"
        return
    }
    // 设置一个模拟客户端，确认访问方式和网址
    var user_agent string = Choose_agent()
    // 从浏览器信息列表中随机选一个
    request.Header.Add("User-Agent",user_agent)
    request.Header.Add("Referer", url)
    // 添加请求头，尤其是浏览器信息
    response, err_2 := client.Do(request)
    // 进行请求
    if err_2 != nil {
        fmt.Println(err_2)
        time.Sleep(time.Duration(600)*time.Second)
        content = "0"
        return
    }
    defer response.Body.Close()
    data, err_4 := ioutil.ReadAll(response.Body)
    if err_4 != nil {
        fmt.Println("read error")
        content = "0"
        return
    }
    content = string(data)
    // 读取页面的html信息
    return
}   

func Get_page_range(url string)(max_page int){
    var html string
    var page_range string
    var max_page_content string
    var max_page_num string
    var max_page_url string
    html = Get_page_content(url)
    
    // fmt.Println(html)
    // 获取页数上限
    page_range = strings.Split(html,"<ul>")[2]
    page_list := strings.Split(page_range,"<li>")

    max_page_content = page_list[len(page_list)-1]
    max_page_url = strings.Split(max_page_content,"href='more_")[1]
    
    // fmt.Println(max_page_url)
    max_page_num = strings.Split(max_page_url,".html'")[0]
    // fmt.Println(max_page_num)
    max_page,err := strconv.Atoi(max_page_num)
    if err != nil{
        fmt.Println("类型转换失败")
        return
    }
    return

}

func Get_comment_list(url string)(result_list []string){
    var html string
    var html_content_with_picture string
    html = Get_page_content(url)

    if html != "0"{
        html_content_with_picture = strings.Split(strings.Split(html,"<ul class=\"wp-list clearfix\"")[1],"</ul>")[0]

        comment_content_list := strings.Split(html_content_with_picture,"<li class=\"wp-item\">")[1:]
        // fmt.Println(html_content_with_picture)

        var result_url string
        for _,comment_content := range comment_content_list{
            // fmt.Println(comment_content)
            comment_url_content := strings.Split(comment_content,"href=\"")[1]
            comment_url := strings.Split(comment_url_content,"\"><img src=\"")[0]
            if []byte(comment_url)[25] == '/'{
                result_url = comment_url[:25]+comment_url[26:]
            }else{
                result_url = comment_url
            }
            // fmt.Println(result_url)
            result_list = append(result_list,result_url)       
        }
        return    
    }else{
        result_list = append(result_list,"0")   
        return
    }
}

func Get_pic_list(url string){
    var html string
    var html_with_pic_list string
    var pic_url string
    var count int
    count = 0
    html = Get_page_content(url)
    if html != "0"{
        html_with_pic_list = strings.Split(strings.Split(html,"<div id=\"picture\">")[1],"</div>")[0]
        // fmt.Println(html_with_pic_list)
        picture_list := strings.Split(html_with_pic_list,"src=\"")[1:]
        for _,pic_content := range picture_list{
            pic_url = strings.Split(pic_content,"\"")[0]
            // fmt.Println(pic_url)
            count = count + Download_picture(pic_url)
        }
        fmt.Println("get",count,"pictures")
    }

}

func Download_picture(url string)(status_code int){
    var pic_content string
    var file_fore string
    var file_tail string
    var file_name string

    pic_content = Get_page_content(url)
    file_fore = Cal_md5(pic_content)
    file_tail = strings.Split(url,".")[3]
    file_name = file_fore+"."+file_tail
    fmt.Println(file_name)

    if pic_content != "0"{
        ioutil.WriteFile(file_path+file_name,[]byte(pic_content),0666)
        fmt.Println("Download ==> ",url,"successful")
        status_code = 1
    }else{
        fmt.Println("Download ==> ",url,"fail")
        status_code = 0
    }
    return

}

func Cal_md5(content string)(result string){
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(content))
    cipherStr := md5Ctx.Sum(nil)
    result = hex.EncodeToString(cipherStr)
    return
}

func main() {
    Make_dir(dir_name)
    var max_page int
    max_page = Get_page_range("http://www.meizitu.com/a/more_1.html")

    var url string
    var url_string string
    url_list := []string{}

    url = "http://www.meizitu.com/a/more_"
    for i := 1;i<= max_page;i++{
        url_string = url+strconv.Itoa(i)+".html"
        // fmt.Println(url_string)
        url_list = append(url_list,url_string)
    }
    for _,one_url := range url_list{
        comment_list := Get_comment_list(one_url)
        for _,comment_url := range comment_list{
            Get_pic_list(comment_url)
        }

    } 
    
}