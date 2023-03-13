# goLearn
contains code created during learning golang
Steps to create a git in local machine and synchronize with github
1. In local machine, create a code folder hierarchy , for eg , for golang, hierachy is go->code->src->github.com->sachin19183
2. Important thing here is that in the src, we have created a folder hierarchy to mimic github. my username is sachin19183 in github which is reflected in the folder hierarchy.
3. Next step is i do a git init in the src folder
4. I create a new repo in github called goLearn.
5. I clone this repo under the folder github.com git clone <URL>
6. go in the goLearn folder created after cloning and execute git remote commmand
 we see origin written which represents the remote repo here
7. Next we add the code in the goLearn folder.after finishing the code, we add the files/folder to git repo (local) via the command git add <filename/folder name>
8. Next step is to commit the changes in local repo via command git commit <filename> -m "message".

9.once the code is committed, its time to push the changes in github remote.
The github repo is named main for me(instead of default master)
hence the command shall be git push origin main
