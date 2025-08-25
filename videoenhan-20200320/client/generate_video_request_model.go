// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v float32) *GenerateVideoRequest
	GetDuration() *float32
	SetDurationAdaption(v bool) *GenerateVideoRequest
	GetDurationAdaption() *bool
	SetFileList(v []*GenerateVideoRequestFileList) *GenerateVideoRequest
	GetFileList() []*GenerateVideoRequestFileList
	SetHeight(v int32) *GenerateVideoRequest
	GetHeight() *int32
	SetMute(v bool) *GenerateVideoRequest
	GetMute() *bool
	SetPuzzleEffect(v bool) *GenerateVideoRequest
	GetPuzzleEffect() *bool
	SetScene(v string) *GenerateVideoRequest
	GetScene() *string
	SetSmartEffect(v bool) *GenerateVideoRequest
	GetSmartEffect() *bool
	SetStyle(v string) *GenerateVideoRequest
	GetStyle() *string
	SetTransitionStyle(v string) *GenerateVideoRequest
	GetTransitionStyle() *string
	SetWidth(v int32) *GenerateVideoRequest
	GetWidth() *int32
}

type GenerateVideoRequest struct {
	// example:
	//
	// 10
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	DurationAdaption *bool `json:"DurationAdaption,omitempty" xml:"DurationAdaption,omitempty"`
	// 1
	//
	// This parameter is required.
	FileList []*GenerateVideoRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// 640
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// true
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// example:
	//
	// true
	PuzzleEffect *bool `json:"PuzzleEffect,omitempty" xml:"PuzzleEffect,omitempty"`
	// example:
	//
	// costume
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// true
	SmartEffect *bool `json:"SmartEffect,omitempty" xml:"SmartEffect,omitempty"`
	// example:
	//
	// fast
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// example:
	//
	// brush
	TransitionStyle *string `json:"TransitionStyle,omitempty" xml:"TransitionStyle,omitempty"`
	// example:
	//
	// 640
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s GenerateVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequest) GetDuration() *float32 {
	return s.Duration
}

func (s *GenerateVideoRequest) GetDurationAdaption() *bool {
	return s.DurationAdaption
}

func (s *GenerateVideoRequest) GetFileList() []*GenerateVideoRequestFileList {
	return s.FileList
}

func (s *GenerateVideoRequest) GetHeight() *int32 {
	return s.Height
}

func (s *GenerateVideoRequest) GetMute() *bool {
	return s.Mute
}

func (s *GenerateVideoRequest) GetPuzzleEffect() *bool {
	return s.PuzzleEffect
}

func (s *GenerateVideoRequest) GetScene() *string {
	return s.Scene
}

func (s *GenerateVideoRequest) GetSmartEffect() *bool {
	return s.SmartEffect
}

func (s *GenerateVideoRequest) GetStyle() *string {
	return s.Style
}

func (s *GenerateVideoRequest) GetTransitionStyle() *string {
	return s.TransitionStyle
}

func (s *GenerateVideoRequest) GetWidth() *int32 {
	return s.Width
}

func (s *GenerateVideoRequest) SetDuration(v float32) *GenerateVideoRequest {
	s.Duration = &v
	return s
}

func (s *GenerateVideoRequest) SetDurationAdaption(v bool) *GenerateVideoRequest {
	s.DurationAdaption = &v
	return s
}

func (s *GenerateVideoRequest) SetFileList(v []*GenerateVideoRequestFileList) *GenerateVideoRequest {
	s.FileList = v
	return s
}

func (s *GenerateVideoRequest) SetHeight(v int32) *GenerateVideoRequest {
	s.Height = &v
	return s
}

func (s *GenerateVideoRequest) SetMute(v bool) *GenerateVideoRequest {
	s.Mute = &v
	return s
}

func (s *GenerateVideoRequest) SetPuzzleEffect(v bool) *GenerateVideoRequest {
	s.PuzzleEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetScene(v string) *GenerateVideoRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoRequest) SetSmartEffect(v bool) *GenerateVideoRequest {
	s.SmartEffect = &v
	return s
}

func (s *GenerateVideoRequest) SetStyle(v string) *GenerateVideoRequest {
	s.Style = &v
	return s
}

func (s *GenerateVideoRequest) SetTransitionStyle(v string) *GenerateVideoRequest {
	s.TransitionStyle = &v
	return s
}

func (s *GenerateVideoRequest) SetWidth(v int32) *GenerateVideoRequest {
	s.Width = &v
	return s
}

func (s *GenerateVideoRequest) Validate() error {
	return dara.Validate(s)
}

type GenerateVideoRequestFileList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1-video1.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/GenerateVideo/1-video1.mp4
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GenerateVideoRequestFileList) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoRequestFileList) GetFileName() *string {
	return s.FileName
}

func (s *GenerateVideoRequestFileList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GenerateVideoRequestFileList) GetType() *string {
	return s.Type
}

func (s *GenerateVideoRequestFileList) SetFileName(v string) *GenerateVideoRequestFileList {
	s.FileName = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetFileUrl(v string) *GenerateVideoRequestFileList {
	s.FileUrl = &v
	return s
}

func (s *GenerateVideoRequestFileList) SetType(v string) *GenerateVideoRequestFileList {
	s.Type = &v
	return s
}

func (s *GenerateVideoRequestFileList) Validate() error {
	return dara.Validate(s)
}
