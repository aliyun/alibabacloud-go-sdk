// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectVerticalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCountDetectDoor(v int32) *RectVerticalRequest
	GetCountDetectDoor() *int32
	SetDetectDoor(v bool) *RectVerticalRequest
	GetDetectDoor() *bool
	SetSubSceneId(v string) *RectVerticalRequest
	GetSubSceneId() *string
	SetVerticalRect(v string) *RectVerticalRequest
	GetVerticalRect() *string
}

type RectVerticalRequest struct {
	// example:
	//
	// 2
	CountDetectDoor *int32 `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
	// example:
	//
	// true
	DetectDoor *bool `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"844946777965268992":[[0.42418407210144654,0.33625107620738004,0.42620819117478337,0.635753199572695],[0.5158627587152769,0.3071978991900134,0.5177513758740194,0.6312118011104786],[0.582693212445534,0.3733969265933281,0.5807612760319687,0.6139402811250833]]}
	VerticalRect *string `json:"VerticalRect,omitempty" xml:"VerticalRect,omitempty"`
}

func (s RectVerticalRequest) String() string {
	return dara.Prettify(s)
}

func (s RectVerticalRequest) GoString() string {
	return s.String()
}

func (s *RectVerticalRequest) GetCountDetectDoor() *int32 {
	return s.CountDetectDoor
}

func (s *RectVerticalRequest) GetDetectDoor() *bool {
	return s.DetectDoor
}

func (s *RectVerticalRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *RectVerticalRequest) GetVerticalRect() *string {
	return s.VerticalRect
}

func (s *RectVerticalRequest) SetCountDetectDoor(v int32) *RectVerticalRequest {
	s.CountDetectDoor = &v
	return s
}

func (s *RectVerticalRequest) SetDetectDoor(v bool) *RectVerticalRequest {
	s.DetectDoor = &v
	return s
}

func (s *RectVerticalRequest) SetSubSceneId(v string) *RectVerticalRequest {
	s.SubSceneId = &v
	return s
}

func (s *RectVerticalRequest) SetVerticalRect(v string) *RectVerticalRequest {
	s.VerticalRect = &v
	return s
}

func (s *RectVerticalRequest) Validate() error {
	return dara.Validate(s)
}
