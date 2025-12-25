// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorrectVertical(v bool) *PredImageRequest
	GetCorrectVertical() *bool
	SetCountDetectDoor(v int64) *PredImageRequest
	GetCountDetectDoor() *int64
	SetDetectDoor(v bool) *PredImageRequest
	GetDetectDoor() *bool
	SetSubSceneId(v string) *PredImageRequest
	GetSubSceneId() *string
}

type PredImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	CorrectVertical *bool `json:"CorrectVertical,omitempty" xml:"CorrectVertical,omitempty"`
	// example:
	//
	// 2
	CountDetectDoor *int64 `json:"CountDetectDoor,omitempty" xml:"CountDetectDoor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	DetectDoor *bool `json:"DetectDoor,omitempty" xml:"DetectDoor,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2345****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s PredImageRequest) String() string {
	return dara.Prettify(s)
}

func (s PredImageRequest) GoString() string {
	return s.String()
}

func (s *PredImageRequest) GetCorrectVertical() *bool {
	return s.CorrectVertical
}

func (s *PredImageRequest) GetCountDetectDoor() *int64 {
	return s.CountDetectDoor
}

func (s *PredImageRequest) GetDetectDoor() *bool {
	return s.DetectDoor
}

func (s *PredImageRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *PredImageRequest) SetCorrectVertical(v bool) *PredImageRequest {
	s.CorrectVertical = &v
	return s
}

func (s *PredImageRequest) SetCountDetectDoor(v int64) *PredImageRequest {
	s.CountDetectDoor = &v
	return s
}

func (s *PredImageRequest) SetDetectDoor(v bool) *PredImageRequest {
	s.DetectDoor = &v
	return s
}

func (s *PredImageRequest) SetSubSceneId(v string) *PredImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *PredImageRequest) Validate() error {
	return dara.Validate(s)
}
