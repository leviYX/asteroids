package assets

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed *
var assets embed.FS

var PlayerSprite = mustLoadImage("images/player.png")

func mustLoadImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

// go:embed *
//这是 编译期指令，告诉 go:embed 把当前源码文件所在目录下的所有文件/子目录统统打包进二进制。
//打包后只能运行时用 embed.FS 读，路径是相对于源码文件目录的相对路径。
//用 * 会递归包含子目录；如果你只想包某几个文件，可以写成

//_ "image/png"
//这是 空白导入（side-effect import）。
//下划线 _ 表示“我只想要这个包的 init() 被执行，但代码里不直接调用它的任何符号”。
//image/png 的 init() 会把 PNG 解码器注册到 image 包里，这样 image.Decode 才能识别 .png 文件。没有这一行，运行时会报 unknown format。
//一句话总结：
////go:embed * 负责把文件打进二进制；_ "image/png" 负责让 image.Decode 认识 PNG。
