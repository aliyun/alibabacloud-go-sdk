// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousMoveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ContinuousMoveRequest
	GetId() *string
	SetOwnerId(v int64) *ContinuousMoveRequest
	GetOwnerId() *int64
	SetPan(v string) *ContinuousMoveRequest
	GetPan() *string
	SetTilt(v string) *ContinuousMoveRequest
	GetTilt() *string
	SetZoom(v string) *ContinuousMoveRequest
	GetZoom() *string
}

type ContinuousMoveRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 0.5
	Pan *string `json:"Pan,omitempty" xml:"Pan,omitempty"`
	// example:
	//
	// 0.5
	Tilt *string `json:"Tilt,omitempty" xml:"Tilt,omitempty"`
	// example:
	//
	// 0.5
	Zoom *string `json:"Zoom,omitempty" xml:"Zoom,omitempty"`
}

func (s ContinuousMoveRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinuousMoveRequest) GoString() string {
	return s.String()
}

func (s *ContinuousMoveRequest) GetId() *string {
	return s.Id
}

func (s *ContinuousMoveRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ContinuousMoveRequest) GetPan() *string {
	return s.Pan
}

func (s *ContinuousMoveRequest) GetTilt() *string {
	return s.Tilt
}

func (s *ContinuousMoveRequest) GetZoom() *string {
	return s.Zoom
}

func (s *ContinuousMoveRequest) SetId(v string) *ContinuousMoveRequest {
	s.Id = &v
	return s
}

func (s *ContinuousMoveRequest) SetOwnerId(v int64) *ContinuousMoveRequest {
	s.OwnerId = &v
	return s
}

func (s *ContinuousMoveRequest) SetPan(v string) *ContinuousMoveRequest {
	s.Pan = &v
	return s
}

func (s *ContinuousMoveRequest) SetTilt(v string) *ContinuousMoveRequest {
	s.Tilt = &v
	return s
}

func (s *ContinuousMoveRequest) SetZoom(v string) *ContinuousMoveRequest {
	s.Zoom = &v
	return s
}

func (s *ContinuousMoveRequest) Validate() error {
	return dara.Validate(s)
}
