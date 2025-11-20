// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v float32) *GenerateVideoAdvanceRequest
	GetDuration() *float32
	SetDurationAdaption(v bool) *GenerateVideoAdvanceRequest
	GetDurationAdaption() *bool
	SetFileList(v []*GenerateVideoAdvanceRequestFileList) *GenerateVideoAdvanceRequest
	GetFileList() []*GenerateVideoAdvanceRequestFileList
	SetHeight(v int32) *GenerateVideoAdvanceRequest
	GetHeight() *int32
	SetMute(v bool) *GenerateVideoAdvanceRequest
	GetMute() *bool
	SetPuzzleEffect(v bool) *GenerateVideoAdvanceRequest
	GetPuzzleEffect() *bool
	SetScene(v string) *GenerateVideoAdvanceRequest
	GetScene() *string
	SetSmartEffect(v bool) *GenerateVideoAdvanceRequest
	GetSmartEffect() *bool
	SetStyle(v string) *GenerateVideoAdvanceRequest
	GetStyle() *string
	SetTransitionStyle(v string) *GenerateVideoAdvanceRequest
	GetTransitionStyle() *string
	SetWidth(v int32) *GenerateVideoAdvanceRequest
	GetWidth() *int32
}

type GenerateVideoAdvanceRequest struct {
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
	FileList []*GenerateVideoAdvanceRequestFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
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

func (s GenerateVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoAdvanceRequest) GetDuration() *float32 {
	return s.Duration
}

func (s *GenerateVideoAdvanceRequest) GetDurationAdaption() *bool {
	return s.DurationAdaption
}

func (s *GenerateVideoAdvanceRequest) GetFileList() []*GenerateVideoAdvanceRequestFileList {
	return s.FileList
}

func (s *GenerateVideoAdvanceRequest) GetHeight() *int32 {
	return s.Height
}

func (s *GenerateVideoAdvanceRequest) GetMute() *bool {
	return s.Mute
}

func (s *GenerateVideoAdvanceRequest) GetPuzzleEffect() *bool {
	return s.PuzzleEffect
}

func (s *GenerateVideoAdvanceRequest) GetScene() *string {
	return s.Scene
}

func (s *GenerateVideoAdvanceRequest) GetSmartEffect() *bool {
	return s.SmartEffect
}

func (s *GenerateVideoAdvanceRequest) GetStyle() *string {
	return s.Style
}

func (s *GenerateVideoAdvanceRequest) GetTransitionStyle() *string {
	return s.TransitionStyle
}

func (s *GenerateVideoAdvanceRequest) GetWidth() *int32 {
	return s.Width
}

func (s *GenerateVideoAdvanceRequest) SetDuration(v float32) *GenerateVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetDurationAdaption(v bool) *GenerateVideoAdvanceRequest {
	s.DurationAdaption = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetFileList(v []*GenerateVideoAdvanceRequestFileList) *GenerateVideoAdvanceRequest {
	s.FileList = v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetHeight(v int32) *GenerateVideoAdvanceRequest {
	s.Height = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetMute(v bool) *GenerateVideoAdvanceRequest {
	s.Mute = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetPuzzleEffect(v bool) *GenerateVideoAdvanceRequest {
	s.PuzzleEffect = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetScene(v string) *GenerateVideoAdvanceRequest {
	s.Scene = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetSmartEffect(v bool) *GenerateVideoAdvanceRequest {
	s.SmartEffect = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetStyle(v string) *GenerateVideoAdvanceRequest {
	s.Style = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetTransitionStyle(v string) *GenerateVideoAdvanceRequest {
	s.TransitionStyle = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) SetWidth(v int32) *GenerateVideoAdvanceRequest {
	s.Width = &v
	return s
}

func (s *GenerateVideoAdvanceRequest) Validate() error {
	if s.FileList != nil {
		for _, item := range s.FileList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateVideoAdvanceRequestFileList struct {
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// video
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GenerateVideoAdvanceRequestFileList) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoAdvanceRequestFileList) GoString() string {
	return s.String()
}

func (s *GenerateVideoAdvanceRequestFileList) GetFileName() *string {
	return s.FileName
}

func (s *GenerateVideoAdvanceRequestFileList) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *GenerateVideoAdvanceRequestFileList) GetType() *string {
	return s.Type
}

func (s *GenerateVideoAdvanceRequestFileList) SetFileName(v string) *GenerateVideoAdvanceRequestFileList {
	s.FileName = &v
	return s
}

func (s *GenerateVideoAdvanceRequestFileList) SetFileUrlObject(v io.Reader) *GenerateVideoAdvanceRequestFileList {
	s.FileUrlObject = v
	return s
}

func (s *GenerateVideoAdvanceRequestFileList) SetType(v string) *GenerateVideoAdvanceRequestFileList {
	s.Type = &v
	return s
}

func (s *GenerateVideoAdvanceRequestFileList) Validate() error {
	return dara.Validate(s)
}
