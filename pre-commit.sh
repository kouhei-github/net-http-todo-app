BRANCH_NAME=`git symbolic-ref HEAD | sed -e 's:^refs/heads/::'`
if test $BRANCH_NAME = main -o $BRANCH_NAME = staging -o $BRANCH_NAME = develop; then
    echo "${BRANCH_NAME} への直コミットは禁止されています。"
    exit 1
fi

files=`git status|grep -e "modified" -e "new file"|sed "s/new file:\(.*\)/\1/g"|sed "s/modified:\(.*\)/\1/g"|cut -f 2`

for file in $files
do
   result=`less $file | grep -e "fmt."`;
   # ファイル名がpre-commit.shだったらcontinue => console.logが入ってるから
   if [ $file = "pre-commit.sh" ]; then
     continue
   fi
   # デバッグ文(console.log)が入っていた処理を終わらせる
   if [ -n "$result" ]; then
       echo "${file} にデバッグ文が入っています。";
       exit 1
    fi
done
