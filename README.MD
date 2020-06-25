# GoForest-tokenbucket

:8199

|方法|url|正确返回|错误返回|
|-|-|-|-|
|create|.../create?id=1&rate=1&capacity=5&max_wait=0|{"code":1,"msg":"success","detail":"token-bucket单元创建成功"}|{"code":-1,"msg":"ParamsErr","detail":"缺少id(令牌桶标识)"}|
|gettoken|.../?id=1|{"status":true,"detail":"","available":5,"rate":1,"capacity":5,"max_wait":0}|{"code":-1,"msg":"IDNotFound","detail":["1"]}|
|remove|.../remove?id=1|{"code":1,"msg":"removed","detail":{"Rate":1,"Capacity":5,"MaxWait":0,"Bucket":{}}}|-|


tokenbucket for goforest