#!/bin/bash
export UNRAID_GUID=GUID вашей флешки
export UNRAID_NAME=Имя вашего сервера на англ. яз.
export UNRAID_DATE=1654646400
export UNRAID_VERSION=Pro
/lib64/ld-linux-x86-64.so.2 /boot/config/unraider
#Start the Management Utility
/usr/local/sbin/emhttp &
