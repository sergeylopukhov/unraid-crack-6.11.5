Уставнока UNRAID 6.11.5

1. Установить на флешку unraid с помощью unraid usb creator с сайта unraid.net

2. Во время установки перписать GUID флешки

3. Файл go на флешке в папке /config - отредактировать и привести к виду:

#!/bin/bash

export UNRAID_GUID=GUID вашей флешки

export UNRAID_NAME=Имя вашего сервера на англ. яз.

export UNRAID_DATE=1654646400

export UNRAID_VERSION=Pro

/lib64/ld-linux-x86-64.so.2 /boot/config/unraider

# Start the Management Utility

/usr/local/sbin/emhttp &

4. Скопировать файлы hook.so, trial.key unrader - в папку /config на флешке

5. Наслаждаться
