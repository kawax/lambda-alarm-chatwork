# lambda-alarm-chatwork

CloudWatchのアラームをChatWorkに通知。

昔はこれを使ってたけど途中で動かなくなったので自分で作り直した。
https://github.com/chatwork/lambda-sns-to-chatwork-function
LambdaがGolang対応したのでGoで再度作り直し。

## デプロイまでの構成
これを参考にCodeBuildやCodePipelineを使うように設定。
https://aws.amazon.com/jp/blogs/compute/announcing-go-support-for-aws-lambda/

難しい場合はローカルでビルドしてアップロードすればいい。このくらいの規模ならそれで十分。

## 通知までの設定
1. AWS SNSにトピックを作る。
2. トピックのサブスクリプションでLambdaを設定。
3. CloudWatch アラームの通知の送信先にSNSトピックを設定。

## 環境変数
Lambdaの環境変数で以下を設定。

- CHATWORK_API_KEY
- CHATWORK_ROOM_ID

## LICENSE
MIT