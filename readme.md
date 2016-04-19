### 概念
- Request:用户请求的信息，程序用来解析用户的请求，包括post，get，cookie，url等
- Response:服务器返回给客户端的信息
- Conn:用户的每次请求链接
- Handler:处理请求**Request**和生成返回信息**Response**的处理逻辑

### r *http.Request
r.Form里面是一个url.Vlues类型，存储类似于key:value的信息，可以使用r.Form.Get()的方式获取对应键值
r.Method返回当前用户请求所用的方法，GET,POST等
r.ParseForm()解析当前表单，Go不会自动解析表单
r.FormValue()也可以获取用户提交的表单参数，且不需要手动调用r.ParseForm()，但是该方法只返回同名阐述中的第一个若参数不存在则返回空字符串

r.ParseMultipartForm()处理文件上传，其中参数表示maxMemory，上传的文件保存在maxMemory这么大的内存空间中，如果文件大小大于该空间则将多出部分存储在系统临时文件中，可以使用r.FormFile获取该文件，io.Copy可以用来拷贝到服务器本地




### html/template
func HTMLEscape(w io.Writer, b[]byte)//把b进行转义后写入w
func HTMLEscapeString(s string) string//把s进行转义后返回转义结果字符串
func HTMLEscaper(args ...interface{}) string// 多个参数一起转义并返回转义结果字符串

template.HTML()以HTML代码方式输出参数内容

### w http.ResponseWriter

### http.HandleFunc()

注意和```http.HandlerFunc()```区分

### 上传文件步骤
- 表单中增加enctype="multipart/form-data"
- 服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
- 使用r.FormFile获取文件句柄，然后对文件进行存储等处理


