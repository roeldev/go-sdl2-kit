// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by go generate; DO NOT EDIT.

package event

import (
	"github.com/go-pogo/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioDeviceEventHandler interface {
	HandleAudioDeviceEvent(*sdl.AudioDeviceEvent) error
}

type AudioDeviceAddedEventHandler interface {
	HandleAudioDeviceAddedEvent(*sdl.AudioDeviceEvent) error
}

type AudioDeviceRemovedEventHandler interface {
	HandleAudioDeviceRemovedEvent(*sdl.AudioDeviceEvent) error
}

type ClipboardEventHandler interface {
	HandleClipboardEvent(*sdl.ClipboardEvent) error
}

type ControllerAxisEventHandler interface {
	HandleControllerAxisEvent(*sdl.ControllerAxisEvent) error
}

type ControllerButtonEventHandler interface {
	HandleControllerButtonEvent(*sdl.ControllerButtonEvent) error
}

type ControllerButtonDownEventHandler interface {
	HandleControllerButtonDownEvent(*sdl.ControllerButtonEvent) error
}

type ControllerButtonUpEventHandler interface {
	HandleControllerButtonUpEvent(*sdl.ControllerButtonEvent) error
}

type ControllerDeviceEventHandler interface {
	HandleControllerDeviceEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceAddedEventHandler interface {
	HandleControllerDeviceAddedEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceMappedEventHandler interface {
	HandleControllerDeviceMappedEvent(*sdl.ControllerDeviceEvent) error
}

type ControllerDeviceRemovedEventHandler interface {
	HandleControllerDeviceRemovedEvent(*sdl.ControllerDeviceEvent) error
}

type DisplayEventHandler interface {
	HandleDisplayEvent(*sdl.DisplayEvent) error
}

type DollarGestureEventHandler interface {
	HandleDollarGestureEvent(*sdl.DollarGestureEvent) error
}

type DropEventHandler interface {
	HandleDropEvent(*sdl.DropEvent) error
}

type JoyAxisEventHandler interface {
	HandleJoyAxisEvent(*sdl.JoyAxisEvent) error
}

type JoyBallEventHandler interface {
	HandleJoyBallEvent(*sdl.JoyBallEvent) error
}

type JoyButtonEventHandler interface {
	HandleJoyButtonEvent(*sdl.JoyButtonEvent) error
}

type JoyButtonDownEventHandler interface {
	HandleJoyButtonDownEvent(*sdl.JoyButtonEvent) error
}

type JoyButtonUpEventHandler interface {
	HandleJoyButtonUpEvent(*sdl.JoyButtonEvent) error
}

type JoyDeviceAddedEventHandler interface {
	HandleJoyDeviceAddedEvent(*sdl.JoyDeviceAddedEvent) error
}

type JoyDeviceRemovedEventHandler interface {
	HandleJoyDeviceRemovedEvent(*sdl.JoyDeviceRemovedEvent) error
}

type JoyHatEventHandler interface {
	HandleJoyHatEvent(*sdl.JoyHatEvent) error
}

type KeyDownEventHandler interface {
	HandleKeyDownEvent(*sdl.KeyboardEvent) error
}

type KeyUpEventHandler interface {
	HandleKeyUpEvent(*sdl.KeyboardEvent) error
}

type KeyboardEventHandler interface {
	HandleKeyboardEvent(*sdl.KeyboardEvent) error
}

type MouseButtonEventHandler interface {
	HandleMouseButtonEvent(*sdl.MouseButtonEvent) error
}

type MouseButtonDownEventHandler interface {
	HandleMouseButtonDownEvent(*sdl.MouseButtonEvent) error
}

type MouseButtonUpEventHandler interface {
	HandleMouseButtonUpEvent(*sdl.MouseButtonEvent) error
}

type MouseMotionEventHandler interface {
	HandleMouseMotionEvent(*sdl.MouseMotionEvent) error
}

type MouseWheelEventHandler interface {
	HandleMouseWheelEvent(*sdl.MouseWheelEvent) error
}

type MultiGestureEventHandler interface {
	HandleMultiGestureEvent(*sdl.MultiGestureEvent) error
}

type OSEventHandler interface {
	HandleOSEvent(*sdl.OSEvent) error
}

type RenderEventHandler interface {
	HandleRenderEvent(*sdl.RenderEvent) error
}

type SensorEventHandler interface {
	HandleSensorEvent(*sdl.SensorEvent) error
}

type SysWMEventHandler interface {
	HandleSysWMEvent(*sdl.SysWMEvent) error
}

type TextEditingEventHandler interface {
	HandleTextEditingEvent(*sdl.TextEditingEvent) error
}

type TextInputEventHandler interface {
	HandleTextInputEvent(*sdl.TextInputEvent) error
}

type TouchFingerEventHandler interface {
	HandleTouchFingerEvent(*sdl.TouchFingerEvent) error
}

type TouchFingerDownEventHandler interface {
	HandleTouchFingerDownEvent(*sdl.TouchFingerEvent) error
}

type TouchFingerMotionEventHandler interface {
	HandleTouchFingerMotionEvent(*sdl.TouchFingerEvent) error
}

type TouchFingerUpEventHandler interface {
	HandleTouchFingerUpEvent(*sdl.TouchFingerEvent) error
}

type UserEventHandler interface {
	HandleUserEvent(*sdl.UserEvent) error
}

type WindowEventHandler interface {
	HandleWindowEvent(*sdl.WindowEvent) error
}

type WindowCloseEventHandler interface {
	HandleWindowCloseEvent(*sdl.WindowEvent) error
}

type WindowEnterEventHandler interface {
	HandleWindowEnterEvent(*sdl.WindowEvent) error
}

type WindowExposedEventHandler interface {
	HandleWindowExposedEvent(*sdl.WindowEvent) error
}

type WindowFocusGainedEventHandler interface {
	HandleWindowFocusGainedEvent(*sdl.WindowEvent) error
}

type WindowFocusLostEventHandler interface {
	HandleWindowFocusLostEvent(*sdl.WindowEvent) error
}

type WindowHiddenEventHandler interface {
	HandleWindowHiddenEvent(*sdl.WindowEvent) error
}

type WindowHitTestEventHandler interface {
	HandleWindowHitTestEvent(*sdl.WindowEvent) error
}

type WindowLeaveEventHandler interface {
	HandleWindowLeaveEvent(*sdl.WindowEvent) error
}

type WindowMaximizedEventHandler interface {
	HandleWindowMaximizedEvent(*sdl.WindowEvent) error
}

type WindowMinimizedEventHandler interface {
	HandleWindowMinimizedEvent(*sdl.WindowEvent) error
}

type WindowMovedEventHandler interface {
	HandleWindowMovedEvent(*sdl.WindowEvent) error
}

type WindowResizedEventHandler interface {
	HandleWindowResizedEvent(*sdl.WindowEvent) error
}

type WindowRestoredEventHandler interface {
	HandleWindowRestoredEvent(*sdl.WindowEvent) error
}

type WindowShownEventHandler interface {
	HandleWindowShownEvent(*sdl.WindowEvent) error
}

type WindowSizeChangedEventHandler interface {
	HandleWindowSizeChangedEvent(*sdl.WindowEvent) error
}

type WindowTakeFocusEventHandler interface {
	HandleWindowTakeFocusEvent(*sdl.WindowEvent) error
}

type handlers struct {
	AudioDevice             []AudioDeviceEventHandler
	AudioDeviceAdded        []AudioDeviceAddedEventHandler
	AudioDeviceRemoved      []AudioDeviceRemovedEventHandler
	Clipboard               []ClipboardEventHandler
	ControllerAxis          []ControllerAxisEventHandler
	ControllerButton        []ControllerButtonEventHandler
	ControllerButtonDown    []ControllerButtonDownEventHandler
	ControllerButtonUp      []ControllerButtonUpEventHandler
	ControllerDevice        []ControllerDeviceEventHandler
	ControllerDeviceAdded   []ControllerDeviceAddedEventHandler
	ControllerDeviceMapped  []ControllerDeviceMappedEventHandler
	ControllerDeviceRemoved []ControllerDeviceRemovedEventHandler
	Display                 []DisplayEventHandler
	DollarGesture           []DollarGestureEventHandler
	Drop                    []DropEventHandler
	JoyAxis                 []JoyAxisEventHandler
	JoyBall                 []JoyBallEventHandler
	JoyButton               []JoyButtonEventHandler
	JoyButtonDown           []JoyButtonDownEventHandler
	JoyButtonUp             []JoyButtonUpEventHandler
	JoyDeviceAdded          []JoyDeviceAddedEventHandler
	JoyDeviceRemoved        []JoyDeviceRemovedEventHandler
	JoyHat                  []JoyHatEventHandler
	KeyDown                 []KeyDownEventHandler
	KeyUp                   []KeyUpEventHandler
	Keyboard                []KeyboardEventHandler
	MouseButton             []MouseButtonEventHandler
	MouseButtonDown         []MouseButtonDownEventHandler
	MouseButtonUp           []MouseButtonUpEventHandler
	MouseMotion             []MouseMotionEventHandler
	MouseWheel              []MouseWheelEventHandler
	MultiGesture            []MultiGestureEventHandler
	OS                      []OSEventHandler
	Render                  []RenderEventHandler
	Sensor                  []SensorEventHandler
	SysWM                   []SysWMEventHandler
	TextEditing             []TextEditingEventHandler
	TextInput               []TextInputEventHandler
	TouchFinger             []TouchFingerEventHandler
	TouchFingerDown         []TouchFingerDownEventHandler
	TouchFingerMotion       []TouchFingerMotionEventHandler
	TouchFingerUp           []TouchFingerUpEventHandler
	User                    []UserEventHandler
	Window                  []WindowEventHandler
	WindowClose             []WindowCloseEventHandler
	WindowEnter             []WindowEnterEventHandler
	WindowExposed           []WindowExposedEventHandler
	WindowFocusGained       []WindowFocusGainedEventHandler
	WindowFocusLost         []WindowFocusLostEventHandler
	WindowHidden            []WindowHiddenEventHandler
	WindowHitTest           []WindowHitTestEventHandler
	WindowLeave             []WindowLeaveEventHandler
	WindowMaximized         []WindowMaximizedEventHandler
	WindowMinimized         []WindowMinimizedEventHandler
	WindowMoved             []WindowMovedEventHandler
	WindowResized           []WindowResizedEventHandler
	WindowRestored          []WindowRestoredEventHandler
	WindowShown             []WindowShownEventHandler
	WindowSizeChanged       []WindowSizeChangedEventHandler
	WindowTakeFocus         []WindowTakeFocusEventHandler
}

func (h *handlers) register(handler interface{}) (n uint) {

	if v, ok := handler.(AudioDeviceEventHandler); ok {
		h.AudioDevice = append(h.AudioDevice, v)
		n++
	}
	if v, ok := handler.(AudioDeviceAddedEventHandler); ok {
		h.AudioDeviceAdded = append(h.AudioDeviceAdded, v)
		n++
	}
	if v, ok := handler.(AudioDeviceRemovedEventHandler); ok {
		h.AudioDeviceRemoved = append(h.AudioDeviceRemoved, v)
		n++
	}
	if v, ok := handler.(ClipboardEventHandler); ok {
		h.Clipboard = append(h.Clipboard, v)
		n++
	}
	if v, ok := handler.(ControllerAxisEventHandler); ok {
		h.ControllerAxis = append(h.ControllerAxis, v)
		n++
	}
	if v, ok := handler.(ControllerButtonEventHandler); ok {
		h.ControllerButton = append(h.ControllerButton, v)
		n++
	}
	if v, ok := handler.(ControllerButtonDownEventHandler); ok {
		h.ControllerButtonDown = append(h.ControllerButtonDown, v)
		n++
	}
	if v, ok := handler.(ControllerButtonUpEventHandler); ok {
		h.ControllerButtonUp = append(h.ControllerButtonUp, v)
		n++
	}
	if v, ok := handler.(ControllerDeviceEventHandler); ok {
		h.ControllerDevice = append(h.ControllerDevice, v)
		n++
	}
	if v, ok := handler.(ControllerDeviceAddedEventHandler); ok {
		h.ControllerDeviceAdded = append(h.ControllerDeviceAdded, v)
		n++
	}
	if v, ok := handler.(ControllerDeviceMappedEventHandler); ok {
		h.ControllerDeviceMapped = append(h.ControllerDeviceMapped, v)
		n++
	}
	if v, ok := handler.(ControllerDeviceRemovedEventHandler); ok {
		h.ControllerDeviceRemoved = append(h.ControllerDeviceRemoved, v)
		n++
	}
	if v, ok := handler.(DisplayEventHandler); ok {
		h.Display = append(h.Display, v)
		n++
	}
	if v, ok := handler.(DollarGestureEventHandler); ok {
		h.DollarGesture = append(h.DollarGesture, v)
		n++
	}
	if v, ok := handler.(DropEventHandler); ok {
		h.Drop = append(h.Drop, v)
		n++
	}
	if v, ok := handler.(JoyAxisEventHandler); ok {
		h.JoyAxis = append(h.JoyAxis, v)
		n++
	}
	if v, ok := handler.(JoyBallEventHandler); ok {
		h.JoyBall = append(h.JoyBall, v)
		n++
	}
	if v, ok := handler.(JoyButtonEventHandler); ok {
		h.JoyButton = append(h.JoyButton, v)
		n++
	}
	if v, ok := handler.(JoyButtonDownEventHandler); ok {
		h.JoyButtonDown = append(h.JoyButtonDown, v)
		n++
	}
	if v, ok := handler.(JoyButtonUpEventHandler); ok {
		h.JoyButtonUp = append(h.JoyButtonUp, v)
		n++
	}
	if v, ok := handler.(JoyDeviceAddedEventHandler); ok {
		h.JoyDeviceAdded = append(h.JoyDeviceAdded, v)
		n++
	}
	if v, ok := handler.(JoyDeviceRemovedEventHandler); ok {
		h.JoyDeviceRemoved = append(h.JoyDeviceRemoved, v)
		n++
	}
	if v, ok := handler.(JoyHatEventHandler); ok {
		h.JoyHat = append(h.JoyHat, v)
		n++
	}
	if v, ok := handler.(KeyDownEventHandler); ok {
		h.KeyDown = append(h.KeyDown, v)
		n++
	}
	if v, ok := handler.(KeyUpEventHandler); ok {
		h.KeyUp = append(h.KeyUp, v)
		n++
	}
	if v, ok := handler.(KeyboardEventHandler); ok {
		h.Keyboard = append(h.Keyboard, v)
		n++
	}
	if v, ok := handler.(MouseButtonEventHandler); ok {
		h.MouseButton = append(h.MouseButton, v)
		n++
	}
	if v, ok := handler.(MouseButtonDownEventHandler); ok {
		h.MouseButtonDown = append(h.MouseButtonDown, v)
		n++
	}
	if v, ok := handler.(MouseButtonUpEventHandler); ok {
		h.MouseButtonUp = append(h.MouseButtonUp, v)
		n++
	}
	if v, ok := handler.(MouseMotionEventHandler); ok {
		h.MouseMotion = append(h.MouseMotion, v)
		n++
	}
	if v, ok := handler.(MouseWheelEventHandler); ok {
		h.MouseWheel = append(h.MouseWheel, v)
		n++
	}
	if v, ok := handler.(MultiGestureEventHandler); ok {
		h.MultiGesture = append(h.MultiGesture, v)
		n++
	}
	if v, ok := handler.(OSEventHandler); ok {
		h.OS = append(h.OS, v)
		n++
	}
	if v, ok := handler.(RenderEventHandler); ok {
		h.Render = append(h.Render, v)
		n++
	}
	if v, ok := handler.(SensorEventHandler); ok {
		h.Sensor = append(h.Sensor, v)
		n++
	}
	if v, ok := handler.(SysWMEventHandler); ok {
		h.SysWM = append(h.SysWM, v)
		n++
	}
	if v, ok := handler.(TextEditingEventHandler); ok {
		h.TextEditing = append(h.TextEditing, v)
		n++
	}
	if v, ok := handler.(TextInputEventHandler); ok {
		h.TextInput = append(h.TextInput, v)
		n++
	}
	if v, ok := handler.(TouchFingerEventHandler); ok {
		h.TouchFinger = append(h.TouchFinger, v)
		n++
	}
	if v, ok := handler.(TouchFingerDownEventHandler); ok {
		h.TouchFingerDown = append(h.TouchFingerDown, v)
		n++
	}
	if v, ok := handler.(TouchFingerMotionEventHandler); ok {
		h.TouchFingerMotion = append(h.TouchFingerMotion, v)
		n++
	}
	if v, ok := handler.(TouchFingerUpEventHandler); ok {
		h.TouchFingerUp = append(h.TouchFingerUp, v)
		n++
	}
	if v, ok := handler.(UserEventHandler); ok {
		h.User = append(h.User, v)
		n++
	}
	if v, ok := handler.(WindowEventHandler); ok {
		h.Window = append(h.Window, v)
		n++
	}
	if v, ok := handler.(WindowCloseEventHandler); ok {
		h.WindowClose = append(h.WindowClose, v)
		n++
	}
	if v, ok := handler.(WindowEnterEventHandler); ok {
		h.WindowEnter = append(h.WindowEnter, v)
		n++
	}
	if v, ok := handler.(WindowExposedEventHandler); ok {
		h.WindowExposed = append(h.WindowExposed, v)
		n++
	}
	if v, ok := handler.(WindowFocusGainedEventHandler); ok {
		h.WindowFocusGained = append(h.WindowFocusGained, v)
		n++
	}
	if v, ok := handler.(WindowFocusLostEventHandler); ok {
		h.WindowFocusLost = append(h.WindowFocusLost, v)
		n++
	}
	if v, ok := handler.(WindowHiddenEventHandler); ok {
		h.WindowHidden = append(h.WindowHidden, v)
		n++
	}
	if v, ok := handler.(WindowHitTestEventHandler); ok {
		h.WindowHitTest = append(h.WindowHitTest, v)
		n++
	}
	if v, ok := handler.(WindowLeaveEventHandler); ok {
		h.WindowLeave = append(h.WindowLeave, v)
		n++
	}
	if v, ok := handler.(WindowMaximizedEventHandler); ok {
		h.WindowMaximized = append(h.WindowMaximized, v)
		n++
	}
	if v, ok := handler.(WindowMinimizedEventHandler); ok {
		h.WindowMinimized = append(h.WindowMinimized, v)
		n++
	}
	if v, ok := handler.(WindowMovedEventHandler); ok {
		h.WindowMoved = append(h.WindowMoved, v)
		n++
	}
	if v, ok := handler.(WindowResizedEventHandler); ok {
		h.WindowResized = append(h.WindowResized, v)
		n++
	}
	if v, ok := handler.(WindowRestoredEventHandler); ok {
		h.WindowRestored = append(h.WindowRestored, v)
		n++
	}
	if v, ok := handler.(WindowShownEventHandler); ok {
		h.WindowShown = append(h.WindowShown, v)
		n++
	}
	if v, ok := handler.(WindowSizeChangedEventHandler); ok {
		h.WindowSizeChanged = append(h.WindowSizeChanged, v)
		n++
	}
	if v, ok := handler.(WindowTakeFocusEventHandler); ok {
		h.WindowTakeFocus = append(h.WindowTakeFocus, v)
		n++
	}
	return
}

func (h *handlers) handle(event sdl.Event) error {
	var err error
	switch e := event.(type) {

	case *sdl.DisplayEvent:
		for _, x := range h.Display {
			errors.Append(&err, x.HandleDisplayEvent(e))
		}

	case *sdl.WindowEvent:
		for _, x := range h.Window {
			errors.Append(&err, x.HandleWindowEvent(e))
		}

		switch e.Event {
		case sdl.WINDOWEVENT_CLOSE:
			for _, x := range h.WindowClose {
				errors.Append(&err, x.HandleWindowCloseEvent(e))
			}
		case sdl.WINDOWEVENT_ENTER:
			for _, x := range h.WindowEnter {
				errors.Append(&err, x.HandleWindowEnterEvent(e))
			}
		case sdl.WINDOWEVENT_EXPOSED:
			for _, x := range h.WindowExposed {
				errors.Append(&err, x.HandleWindowExposedEvent(e))
			}
		case sdl.WINDOWEVENT_FOCUS_GAINED:
			for _, x := range h.WindowFocusGained {
				errors.Append(&err, x.HandleWindowFocusGainedEvent(e))
			}
		case sdl.WINDOWEVENT_FOCUS_LOST:
			for _, x := range h.WindowFocusLost {
				errors.Append(&err, x.HandleWindowFocusLostEvent(e))
			}
		case sdl.WINDOWEVENT_HIDDEN:
			for _, x := range h.WindowHidden {
				errors.Append(&err, x.HandleWindowHiddenEvent(e))
			}
		case sdl.WINDOWEVENT_HIT_TEST:
			for _, x := range h.WindowHitTest {
				errors.Append(&err, x.HandleWindowHitTestEvent(e))
			}
		case sdl.WINDOWEVENT_LEAVE:
			for _, x := range h.WindowLeave {
				errors.Append(&err, x.HandleWindowLeaveEvent(e))
			}
		case sdl.WINDOWEVENT_MAXIMIZED:
			for _, x := range h.WindowMaximized {
				errors.Append(&err, x.HandleWindowMaximizedEvent(e))
			}
		case sdl.WINDOWEVENT_MINIMIZED:
			for _, x := range h.WindowMinimized {
				errors.Append(&err, x.HandleWindowMinimizedEvent(e))
			}
		case sdl.WINDOWEVENT_MOVED:
			for _, x := range h.WindowMoved {
				errors.Append(&err, x.HandleWindowMovedEvent(e))
			}
		case sdl.WINDOWEVENT_RESIZED:
			for _, x := range h.WindowResized {
				errors.Append(&err, x.HandleWindowResizedEvent(e))
			}
		case sdl.WINDOWEVENT_RESTORED:
			for _, x := range h.WindowRestored {
				errors.Append(&err, x.HandleWindowRestoredEvent(e))
			}
		case sdl.WINDOWEVENT_SHOWN:
			for _, x := range h.WindowShown {
				errors.Append(&err, x.HandleWindowShownEvent(e))
			}
		case sdl.WINDOWEVENT_SIZE_CHANGED:
			for _, x := range h.WindowSizeChanged {
				errors.Append(&err, x.HandleWindowSizeChangedEvent(e))
			}
		case sdl.WINDOWEVENT_TAKE_FOCUS:
			for _, x := range h.WindowTakeFocus {
				errors.Append(&err, x.HandleWindowTakeFocusEvent(e))
			}
		}

	case *sdl.KeyboardEvent:
		for _, x := range h.Keyboard {
			errors.Append(&err, x.HandleKeyboardEvent(e))
		}

		switch e.Type {
		case sdl.KEYDOWN:
			for _, x := range h.KeyDown {
				errors.Append(&err, x.HandleKeyDownEvent(e))
			}
		case sdl.KEYUP:
			for _, x := range h.KeyUp {
				errors.Append(&err, x.HandleKeyUpEvent(e))
			}
		}

	case *sdl.TextEditingEvent:
		for _, x := range h.TextEditing {
			errors.Append(&err, x.HandleTextEditingEvent(e))
		}

	case *sdl.TextInputEvent:
		for _, x := range h.TextInput {
			errors.Append(&err, x.HandleTextInputEvent(e))
		}

	case *sdl.MouseMotionEvent:
		for _, x := range h.MouseMotion {
			errors.Append(&err, x.HandleMouseMotionEvent(e))
		}

	case *sdl.MouseButtonEvent:
		for _, x := range h.MouseButton {
			errors.Append(&err, x.HandleMouseButtonEvent(e))
		}

		switch e.Type {
		case sdl.MOUSEBUTTONDOWN:
			for _, x := range h.MouseButtonDown {
				errors.Append(&err, x.HandleMouseButtonDownEvent(e))
			}
		case sdl.MOUSEBUTTONUP:
			for _, x := range h.MouseButtonUp {
				errors.Append(&err, x.HandleMouseButtonUpEvent(e))
			}
		}

	case *sdl.MouseWheelEvent:
		for _, x := range h.MouseWheel {
			errors.Append(&err, x.HandleMouseWheelEvent(e))
		}

	case *sdl.JoyAxisEvent:
		for _, x := range h.JoyAxis {
			errors.Append(&err, x.HandleJoyAxisEvent(e))
		}

	case *sdl.JoyBallEvent:
		for _, x := range h.JoyBall {
			errors.Append(&err, x.HandleJoyBallEvent(e))
		}

	case *sdl.JoyHatEvent:
		for _, x := range h.JoyHat {
			errors.Append(&err, x.HandleJoyHatEvent(e))
		}

	case *sdl.JoyButtonEvent:
		for _, x := range h.JoyButton {
			errors.Append(&err, x.HandleJoyButtonEvent(e))
		}

		switch e.Type {
		case sdl.JOYBUTTONDOWN:
			for _, x := range h.JoyButtonDown {
				errors.Append(&err, x.HandleJoyButtonDownEvent(e))
			}
		case sdl.JOYBUTTONUP:
			for _, x := range h.JoyButtonUp {
				errors.Append(&err, x.HandleJoyButtonUpEvent(e))
			}
		}

	case *sdl.JoyDeviceAddedEvent:
		for _, x := range h.JoyDeviceAdded {
			errors.Append(&err, x.HandleJoyDeviceAddedEvent(e))
		}

	case *sdl.JoyDeviceRemovedEvent:
		for _, x := range h.JoyDeviceRemoved {
			errors.Append(&err, x.HandleJoyDeviceRemovedEvent(e))
		}

	case *sdl.ControllerAxisEvent:
		for _, x := range h.ControllerAxis {
			errors.Append(&err, x.HandleControllerAxisEvent(e))
		}

	case *sdl.ControllerButtonEvent:
		for _, x := range h.ControllerButton {
			errors.Append(&err, x.HandleControllerButtonEvent(e))
		}

		switch e.Type {
		case sdl.CONTROLLERBUTTONDOWN:
			for _, x := range h.ControllerButtonDown {
				errors.Append(&err, x.HandleControllerButtonDownEvent(e))
			}
		case sdl.CONTROLLERBUTTONUP:
			for _, x := range h.ControllerButtonUp {
				errors.Append(&err, x.HandleControllerButtonUpEvent(e))
			}
		}

	case *sdl.ControllerDeviceEvent:
		for _, x := range h.ControllerDevice {
			errors.Append(&err, x.HandleControllerDeviceEvent(e))
		}

		switch e.Type {
		case sdl.CONTROLLERDEVICEADDED:
			for _, x := range h.ControllerDeviceAdded {
				errors.Append(&err, x.HandleControllerDeviceAddedEvent(e))
			}
		case sdl.CONTROLLERDEVICEREMAPPED:
			for _, x := range h.ControllerDeviceMapped {
				errors.Append(&err, x.HandleControllerDeviceMappedEvent(e))
			}
		case sdl.CONTROLLERDEVICEREMOVED:
			for _, x := range h.ControllerDeviceRemoved {
				errors.Append(&err, x.HandleControllerDeviceRemovedEvent(e))
			}
		}

	case *sdl.AudioDeviceEvent:
		for _, x := range h.AudioDevice {
			errors.Append(&err, x.HandleAudioDeviceEvent(e))
		}

		switch e.Type {
		case sdl.AUDIODEVICEADDED:
			for _, x := range h.AudioDeviceAdded {
				errors.Append(&err, x.HandleAudioDeviceAddedEvent(e))
			}
		case sdl.AUDIODEVICEREMOVED:
			for _, x := range h.AudioDeviceRemoved {
				errors.Append(&err, x.HandleAudioDeviceRemovedEvent(e))
			}
		}

	case *sdl.TouchFingerEvent:
		for _, x := range h.TouchFinger {
			errors.Append(&err, x.HandleTouchFingerEvent(e))
		}

		switch e.Type {
		case sdl.FINGERDOWN:
			for _, x := range h.TouchFingerDown {
				errors.Append(&err, x.HandleTouchFingerDownEvent(e))
			}
		case sdl.FINGERMOTION:
			for _, x := range h.TouchFingerMotion {
				errors.Append(&err, x.HandleTouchFingerMotionEvent(e))
			}
		case sdl.FINGERUP:
			for _, x := range h.TouchFingerUp {
				errors.Append(&err, x.HandleTouchFingerUpEvent(e))
			}
		}

	case *sdl.MultiGestureEvent:
		for _, x := range h.MultiGesture {
			errors.Append(&err, x.HandleMultiGestureEvent(e))
		}

	case *sdl.DollarGestureEvent:
		for _, x := range h.DollarGesture {
			errors.Append(&err, x.HandleDollarGestureEvent(e))
		}

	case *sdl.DropEvent:
		for _, x := range h.Drop {
			errors.Append(&err, x.HandleDropEvent(e))
		}

	case *sdl.SensorEvent:
		for _, x := range h.Sensor {
			errors.Append(&err, x.HandleSensorEvent(e))
		}

	case *sdl.RenderEvent:
		for _, x := range h.Render {
			errors.Append(&err, x.HandleRenderEvent(e))
		}

	case *sdl.OSEvent:
		for _, x := range h.OS {
			errors.Append(&err, x.HandleOSEvent(e))
		}

	case *sdl.ClipboardEvent:
		for _, x := range h.Clipboard {
			errors.Append(&err, x.HandleClipboardEvent(e))
		}

	case *sdl.UserEvent:
		for _, x := range h.User {
			errors.Append(&err, x.HandleUserEvent(e))
		}

	case *sdl.SysWMEvent:
		for _, x := range h.SysWM {
			errors.Append(&err, x.HandleSysWMEvent(e))
		}

	}

	return err
}
