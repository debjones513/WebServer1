# Connect Issues
https://stackoverflow.com/questions/36732875/cant-connect-to-public-ip-for-ec2-instance  

If my Mac's IP changes, I have to update the firewall rule for the instance, or set it to 0.0.0.0/0

**To see what app is listening on a port**  
sudo netstat -tulpn | grep LISTEN  

**[ec2-user@ip-172-31-20-12 ~]$ netstat -tulpn | grep LISTEN**  
(Not all processes could be identified, non-owned process info
will not be shown, you would have to be root to see it all.)
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      -                   
tcp6       0      0 :::22                   :::*                    LISTEN      -                   
tcp6       0      0 :::8080                 :::*                    LISTEN      31553/./linux_binar **

**Deborahs-iMac:vtserver debjo$ nc -v 18.117.89.249 8080**  
Connection to 18.117.89.249 port 8080 [tcp/http-alt] succeeded!

### On my Mac in Chrome  
**http://44.234.131.118:8080/**
Make sure you use http - else you will get an SSL error.


**Output**  
*Web server is running!  
To test: http://localhost:8080/upload/vtserver_test*


**http://44.234.131.118:8080/upload/vtserver_test**  

**Output**  



### Resources
https://www.cyberciti.biz/faq/unix-linux-check-if-port-is-in-use-command/
https://stackoverflow.com/questions/36732875/cant-connect-to-public-ip-for-ec2-instance

# Build
**Build LInux on Mac**  
*env GOOS=linux GOARCH=386  go build*

# Upload to AWS
**Deborahs-iMac:~ debjo$ scp -i ".ssh/awsShortlnkED25519.pem" ~/GitHub/WebServer1/go/bin/ws  ec2-user@ec2-18-117-89-249.us-east-2.compute.amazonaws.com:shortlnk/ws**

**Deborahs-iMac:~ debjo$ scp -i ".ssh/awsShortlnkED25519.pem" ~/GitHub/WebServer1/go/src/shortlnk/ws/ws  ec2-user@ec2-18-117-89-249.us-east-2.compute.amazonaws.com:shortlnk/ws**



# SSH to EC2 in AWS
Deborahs-iMac:~ debjo$ ssh -i ".ssh/awsShortlnkED25519.pem" ec2-user@ec2-18-117-89-249.us-east-2.compute.amazonaws.com
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@         WARNING: UNPROTECTED PRIVATE KEY FILE!          @
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
**Permissions 0644 for '.ssh/awsShortlnkED25519.pem' are too open.**
It is required that your private key files are NOT accessible by others.
This private key will be ignored.
Load key ".ssh/awsShortlnkED25519.pem": bad permissions
ec2-user@ec2-3-144-196-146.us-east-2.compute.amazonaws.com: Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
**Deborahs-iMac:~ debjo$ chmod 400  .ssh/awsShortlnkED25519.pem**

**Deborahs-iMac:~ debjo$ ssh -i ".ssh/awsShortlnkED25519.pem" ec2-user@ec2-18-117-89-249.us-east-2.compute.amazonaws.com**
,     #_
~\_  ####_        Amazon Linux 2023
...

**[ec2-user@ip-172-31-16-45 ~]$ printenv**
SHELL=/bin/bash
HISTCONTROL=ignoredups
SYSTEMD_COLORS=false
HISTSIZE=1000
HOSTNAME=ip-172-31-16-45.us-east-2.compute.internal
PWD=/home/ec2-user
LOGNAME=ec2-user
XDG_SESSION_TYPE=tty
MOTD_SHOWN=pam
HOME=/home/ec2-user
LANG=C.UTF-8
LS_COLORS=rs=0:di=01;34:ln=01;36:mh=00:pi=40;33:so=01;35:do=01;35:bd=40;33;01:cd=40;33;01:or=40;31;01:mi=01;37;41:su=37;41:sg=30;43:ca=30;41:tw=30;42:ow=34;42:st=37;44:ex=01;32:*.tar=01;31:*.tgz=01;31:*.arc=01;31:*.arj=01;31:*.taz=01;31:*.lha=01;31:*.lz4=01;31:*.lzh=01;31:*.lzma=01;31:*.tlz=01;31:*.txz=01;31:*.tzo=01;31:*.t7z=01;31:*.zip=01;31:*.z=01;31:*.dz=01;31:*.gz=01;31:*.lrz=01;31:*.lz=01;31:*.lzo=01;31:*.xz=01;31:*.zst=01;31:*.tzst=01;31:*.bz2=01;31:*.bz=01;31:*.tbz=01;31:*.tbz2=01;31:*.tz=01;31:*.deb=01;31:*.rpm=01;31:*.jar=01;31:*.war=01;31:*.ear=01;31:*.sar=01;31:*.rar=01;31:*.alz=01;31:*.ace=01;31:*.zoo=01;31:*.cpio=01;31:*.7z=01;31:*.rz=01;31:*.cab=01;31:*.wim=01;31:*.swm=01;31:*.dwm=01;31:*.esd=01;31:*.jpg=01;35:*.jpeg=01;35:*.mjpg=01;35:*.mjpeg=01;35:*.gif=01;35:*.bmp=01;35:*.pbm=01;35:*.pgm=01;35:*.ppm=01;35:*.tga=01;35:*.xbm=01;35:*.xpm=01;35:*.tif=01;35:*.tiff=01;35:*.png=01;35:*.svg=01;35:*.svgz=01;35:*.mng=01;35:*.pcx=01;35:*.mov=01;35:*.mpg=01;35:*.mpeg=01;35:*.m2v=01;35:*.mkv=01;35:*.webm=01;35:*.webp=01;35:*.ogm=01;35:*.mp4=01;35:*.m4v=01;35:*.mp4v=01;35:*.vob=01;35:*.qt=01;35:*.nuv=01;35:*.wmv=01;35:*.asf=01;35:*.rm=01;35:*.rmvb=01;35:*.flc=01;35:*.avi=01;35:*.fli=01;35:*.flv=01;35:*.gl=01;35:*.dl=01;35:*.xcf=01;35:*.xwd=01;35:*.yuv=01;35:*.cgm=01;35:*.emf=01;35:*.ogv=01;35:*.ogx=01;35:*.aac=01;36:*.au=01;36:*.flac=01;36:*.m4a=01;36:*.mid=01;36:*.midi=01;36:*.mka=01;36:*.mp3=01;36:*.mpc=01;36:*.ogg=01;36:*.ra=01;36:*.wav=01;36:*.oga=01;36:*.opus=01;36:*.spx=01;36:*.xspf=01;36:
SSH_CONNECTION=50.47.115.215 62558 172.31.16.45 22
XDG_SESSION_CLASS=user
SELINUX_ROLE_REQUESTED=
TERM=xterm-256color
LESSOPEN=||/usr/bin/lesspipe.sh %s
USER=ec2-user
SELINUX_USE_CURRENT_RANGE=
SHLVL=1
XDG_SESSION_ID=3
XDG_RUNTIME_DIR=/run/user/1000
S_COLORS=auto
SSH_CLIENT=50.47.115.215 62558 22
which_declare=declare -f
PATH=/home/ec2-user/.local/bin:/home/ec2-user/bin:/usr/local/bin:/usr/bin:/usr/local/sbin:/usr/sbin
SELINUX_LEVEL_REQUESTED=
DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus
MAIL=/var/spool/mail/ec2-user
SSH_TTY=/dev/pts/0
BASH_FUNC_which%%=() {  ( alias;
 eval ${which_declare} ) | /usr/bin/which --tty-only --read-alias --read-functions --show-tilde --show-dot "$@"
}
_=/usr/bin/printenv
[ec2-user@ip-172-31-16-45 ~]$ 


**[ec2-user@ip-172-31-20-12 ~]$ ls -laf  /**   
.  ..  boot  dev  etc  local  proc  run  sys  tmp  var  usr  bin  sbin  lib  lib64  home  media  mnt  opt  root  srv



## Runcloud?
https://runcloud.io/blog/aws