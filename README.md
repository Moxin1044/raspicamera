# What is it?

这是一个使用Golang写的一个树莓派摄像头快速调用功能，因为发现树莓派的摄像头使用cv2很难受，所以调用了树莓派4B的一个小功能：`libcamera-jpeg`，快速指定命令并拍照。

This is a quick call function of raspberry pi camera written by Golang. Because the camera of raspberry pi is hard to use cv2, it calls a small function of raspberry pi 4B: 'libcamera-jpeg' to quickly specify commands and take pictures.

# How to use?

## ARM64

前往https://github.com/Moxin1044/raspicamera/releases/tag/raspicamera 下载 `raspicamera_ARM64.zip`，解压zip并给予可执行权限：`chmod +x raspicamera`，随后按照下面方式使用他：

```shell
./raspicamera hello.jpg
```

上面的代码将会保存`hello.jpg`

```shell
./raspicamera
```

上面的代码将会保存一个名为`save.jpg`的图片

go to https://github.com/Moxin1044/raspicamera/releases/tag/raspicamera Download `raspicamera_ ARM64. zip `, decompress the zip and grant the executable permission:`chmod+x raspicamera `, and then use it as follows:

```shell
./raspicamera hello.jpg
```

The above code will be saved ` hello.jpg`

```shell
./raspicamera
```

The above code will save an image named `save. jpg`

## ARMhf(arm、arm32)

由于手上没有空闲的树莓派和内存卡刷ARMhf系统，所以这里需要编译，不过已经给出源码了，需要使用的编译一下或者等我后面有闲心了再编译。

Since I have no free raspberry pi and memory card to brush the ARMhf system, I need to compile it here, but the source code has been given, and I need to compile it later or when I have leisure.