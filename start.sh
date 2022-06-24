nohup  ./tools.gitlabhooks >> ./nohup`date +%Y-%m-%d`.out 2>&1 &

tail -f ./nohup`date +%Y-%m-%d`.out