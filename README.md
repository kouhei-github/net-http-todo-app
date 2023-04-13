# 1. net-http-clean-architecture-sample 
Golangの**net/http**パッケージを使用したAPI開発   
下記のようなディレクトリ構造になっている   

```txt
repository -> database関連
route      -> ルーティング
controller -> コントローラーやハンドラー
config     -> configの設定（環境変数など）
service    -> サービス
log        -> ログ関連
```

Databaseには**GORM**   
configには**gopkg.in/ini.v1**   
logには**github.com/sirupsen/logrus**   

---

## 2. 起動方法

### 2.1 初期設定
```shell

# 環境変数の作成・編集
$ cp .env.sample .env

# config.iniファイルの作成・編集
$ cp config.ini.sample config.ini


```


### 2.2 コンテナ関連
```shell
# 構築方法
$ docker compose build

# 起動方法
$ docker compose up -d

# 停止方法
$ docker compose down
```
