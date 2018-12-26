# MDhb
## Message-Digest Hibary
精简版的消息摘要算法，仅供娱乐，具体步骤如下：<br />
1、对数据进行填充，如果数据对32求余不为零则进行填充0<br />
2、初始化幻数A，B，C，D<br />
3、对每32bit数据进行操作，同MD5算法一样进行FF,GG等操作，只不过简化步骤<br />
4、将A，B，C，D转化为字符串形式<br />
使用时输入字符串，返回16进制的字符串形式，可以自己修改为byte数组
```go
println(MDhb("hello,world")) // ED9D12D
println(MDhb("hello world")) // CD9D1AD
```