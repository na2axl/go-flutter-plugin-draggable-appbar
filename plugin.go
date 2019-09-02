package window

import (
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// Plugin is a go-flutter plugin that help interact with the GLFW window from
// any flutter widget of the app.
type Plugin struct {
	window           *glfw.Window
	windowTitle      string
	windowDragActive chan bool
}

type size struct {
	width  int
	height int
}

type position struct {
	x float64
	y float64
}

const channelName = "na2axl.github.io/go-flutter-plugin-window"

var _ flutter.Plugin = &Plugin{}     // compile-time type check
var _ flutter.PluginGLFW = &Plugin{} // compile-time type check

// InitPlugin creates a MethodChannel for "na2axl.github.io/go-flutter-plugin-titlebar"
func (p *Plugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.windowDragActive = make(chan bool)
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("onDragStart", p.onDragStart)
	channel.HandleFunc("onDragEnd", p.onDragEnd)
	return nil
}

// InitPluginGLFW is used to gain control over the glfw.Window
func (p *Plugin) InitPluginGLFW(window *glfw.Window) error {
	p.window = window
	return nil
}

// onDragStart initializes a new window drag session.
func (p *Plugin) onDragStart(arguments interface{}) (reply interface{}, err error) {
	argumentsMap := arguments.(map[interface{}]interface{})
	cursorPosX := int(argumentsMap["x"].(float64))
	cursorPosY := int(argumentsMap["y"].(float64))
	for {
		select {
		case <-p.windowDragActive:
			return nil, nil
		default:
			xpos, ypos := p.window.GetCursorPos()
			deltaX := int(xpos) - cursorPosX
			deltaY := int(ypos) - cursorPosY

			x, y := p.window.GetPos()
			p.window.SetPos(x+deltaX, y+deltaY)
		}
	}
}

// onDragEnd closes the current window drag session.
func (p *Plugin) onDragEnd(arguments interface{}) (reply interface{}, err error) {
	p.windowDragActive <- false
	return nil, nil
}

// maximize maximizes the window to the full screen.
func (p *Plugin) maximize(arguments interface{}) (reply interface{}, err error) {
	return nil, p.window.Maximize()
}

func (p *Plugin) restore(arguments interface{}) (reply interface{}, err error) {
	return nil, p.window.Restore()
}

func (p *Plugin) iconify(arguments interface{}) (reply interface{}, err error) {
	return nil, p.window.Iconify()
}

func (p *Plugin) focus(arguments interface{}) (reply interface{}, err error) {
	return nil, p.window.Focus()
}

func (p *Plugin) show(arguments interface{}) (reply interface{}, err error) {
	p.window.Show()
	return nil, nil
}

func (p *Plugin) hide(arguments interface{}) (reply interface{}, err error) {
	p.window.Hide()
	return nil, nil
}

func (p *Plugin) close(arguments interface{}) (reply interface{}, err error) {
	p.window.SetShouldClose(true)
	return nil, nil
}

func (p *Plugin) setTitle(arguments interface{}) (reply interface{}, err error) {
	p.windowTitle = arguments.(map[interface{}]interface{})["title"].(string)
	p.window.SetTitle(p.windowTitle)
	return nil, nil
}

func (p *Plugin) getTitle(arguments interface{}) (reply interface{}, err error) {
	return p.windowTitle, nil
}

func (p *Plugin) getWidth(arguments interface{}) (reply interface{}, err error) {
	w, _ := p.window.GetSize()
	return w, nil
}

func (p *Plugin) getHeight(arguments interface{}) (reply interface{}, err error) {
	_, h := p.window.GetSize()
	return h, nil
}

func (p *Plugin) getSize(arguments interface{}) (reply interface{}, err error) {
	w, h := p.window.GetSize()
	return size{width: w, height: h}, nil
}

func (p *Plugin) getPosition(arguments interface{}) (reply interface{}, err error) {
	x, y := p.window.GetPos()
	return position{x: float64(x), y: float64(y)}, nil
}

func (p *Plugin) getPositionX(arguments interface{}) (reply interface{}, err error) {
	x, _ := p.window.GetPos()
	return x, nil
}

func (p *Plugin) getPositionY(arguments interface{}) (reply interface{}, err error) {
	_, y := p.window.GetPos()
	return y, nil
}

func (p *Plugin) getCursorPosition(arguments interface{}) (reply interface{}, err error) {
	x, y := p.window.GetCursorPos()
	return position{x: x, y: y}, nil
}

func (p *Plugin) setSize(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	w := args["width"].(int)
	h := args["height"].(int)
	p.window.SetSize(w, h)
	return nil, nil
}

func (p *Plugin) setPosition(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	x := args["x"].(int)
	y := args["y"].(int)
	p.window.SetPos(x, y)
	return nil, nil
}

func (p *Plugin) setCursorPosition(arguments interface{}) (reply interface{}, err error) {
	args := arguments.(map[interface{}]interface{})
	x := args["x"].(float64)
	y := args["y"].(float64)
	p.window.SetCursorPos(x, y)
	return nil, nil
}

func (p *Plugin) setDropCallback(arguments interface{}) (reply interface{}, err error) {
	return nil, nil
}
