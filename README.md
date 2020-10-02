# sound-ethernet-streaming

Стриминг аудио по UDP

Server - сервер для раздачи аудиосигнала

Player - клиент для приема и воспроизведения аудиосигнала

Recorder - клиент для получения аудиосигнала с микрофона

## ToDo

### Server

- [X] Стримминг аудиосигнала

  - [X] из wav файла
  - [ ] с recorder

- [X] Написание пакета для работы с микрофоном
- [X] RPC система управление

  - [X] player
  - [X] recorder

- [X] Запись сигнала в wav файл
- [ ] Наложение 2х дорожек

### Player

- [X] Выбор звуковой карты
- [X] Кэш
- [X] Рефакторинг
- [X] Управление с помощью RPC
- [ ] Регулировка громкости

### Recorder

- [X] Получение аудиосигнала с микрофона
- [X] Передача аудиосигнала на сервер
- [X] Управление с помощью RPC

## Запуск server

1. Скачать проект на машину, на которой будет развернут server

        git clone git@github.com:GeoIrb/sound-ethernet-streaming.git
2. Поместите аудиофайл, который необходимо будет стримить в папку `audio/`

3. Собрать образ сервера

        make build-server tag=IMAGE-NAME
4. Запуск сервера

        docker run -d --rm \
        -p PORT:PORT \ 
        -e ENVIROMENTS \ 
        IMAGE-NAME

**PORT** - порт, на который будет раздача (возможно это лишнее)

**ENVIROMENTS** - переменные окружения

- FILE=/audio/`FILE`.wav - файл для стримминга
- DST_ADDRESS="IP:PORT" - на какой IP и на какой PORT будет рассылка, по умолчанию 255.255.255.255:8080 - рассылка по всей сети на порт 8080

        make build-server server
        docker run -d --rm -p 8081:8081 -p 8082:8082 -e FILE=/audio/test.wav server

## Запуск player

1. Скачать проект на машину, на которой будет развернут player

        git clone git@github.com:GeoIrb/sound-ethernet-streaming.git
2. Собрать образ клиент

        make build-player tag=IMAGE-NAME
3. Запуск клиента

        docker run -d --rm \
        -p 0.0.0.0:8081:8081/tcp \ 
        -p 0.0.0.0:PORT:PORT -p 0.0.0.0:PORT:PORT/udp \
        --device /dev/snd \
        -e ENVIROMENTS \
        IMAGE-NAME

**PORT** - порт, на котором будет работать клиент

**ENVIROMENTS** - переменные окружения

- PORT - порт, на котором будет работать клиент
- PLAYBACK_DEVICE_NAME - устройство, на котором будет воспроизводиться принятый аудио сигнал

        make build-player tag player
        docker run -d --rm -p 0.0.0.0:8081:8081/tcp -p 0.0.0.0:8082:8082 -p 0.0.0.0:8082:8082/udp --device /dev/snd player
