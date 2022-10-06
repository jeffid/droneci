# Drone CI
Golang项目代码使用Drone

## 启动
### 代码仓库
目前使用Gitea, 已部署到公网

### Drone
drone-server和drone-runner使用docker-compose搭建测试环境
```shell
# 事先准备好.env文件
docker-compose pull
docker-compose up -d
```



## Gitea
公网服务器部署Gitea
### 注意事项
- 目录要先生成 /var/log/gitea/; 给予`custom``data`目录读写权限给管理账号
- 使用独立账号作为gitea的管理账号, 避免与真实用户账号发生ssh登录等冲突; `custom/conf/app.ini`中的账号名要一致
- git账号不要使用限制性`ssh`(`-s /usr/bin/git-shell`),默认就行 `useradd git`

### supervisor 配置
```shell
[program:gitea]
directory=/usr/local/gitea/
command=/usr/local/gitea/gitea web
autostart=true
autorestart=true
startsecs=10
stdout_logfile=/var/log/gitea/stdout.log
stdout_logfile_maxbytes=1MB
stdout_logfile_backups=10
stdout_capture_maxbytes=1MB
stderr_logfile=/var/log/gitea/stderr.log
stderr_logfile_maxbytes=1MB
stderr_logfile_backups=10
stderr_capture_maxbytes=1MB
user = git
environment = HOME="/home/git", USER="git"  
```
