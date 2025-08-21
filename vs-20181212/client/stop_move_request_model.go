// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMoveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopMoveRequest
	GetId() *string
	SetOwnerId(v int64) *StopMoveRequest
	GetOwnerId() *int64
	SetPan(v bool) *StopMoveRequest
	GetPan() *bool
	SetTilt(v bool) *StopMoveRequest
	GetTilt() *bool
	SetZoom(v bool) *StopMoveRequest
	GetZoom() *bool
}

type StopMoveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// true
	Pan *bool `json:"Pan,omitempty" xml:"Pan,omitempty"`
	// example:
	//
	// true
	Tilt *bool `json:"Tilt,omitempty" xml:"Tilt,omitempty"`
	// example:
	//
	// true
	Zoom *bool `json:"Zoom,omitempty" xml:"Zoom,omitempty"`
}

func (s StopMoveRequest) String() string {
	return dara.Prettify(s)
}

func (s StopMoveRequest) GoString() string {
	return s.String()
}

func (s *StopMoveRequest) GetId() *string {
	return s.Id
}

func (s *StopMoveRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopMoveRequest) GetPan() *bool {
	return s.Pan
}

func (s *StopMoveRequest) GetTilt() *bool {
	return s.Tilt
}

func (s *StopMoveRequest) GetZoom() *bool {
	return s.Zoom
}

func (s *StopMoveRequest) SetId(v string) *StopMoveRequest {
	s.Id = &v
	return s
}

func (s *StopMoveRequest) SetOwnerId(v int64) *StopMoveRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMoveRequest) SetPan(v bool) *StopMoveRequest {
	s.Pan = &v
	return s
}

func (s *StopMoveRequest) SetTilt(v bool) *StopMoveRequest {
	s.Tilt = &v
	return s
}

func (s *StopMoveRequest) SetZoom(v bool) *StopMoveRequest {
	s.Zoom = &v
	return s
}

func (s *StopMoveRequest) Validate() error {
	return dara.Validate(s)
}
