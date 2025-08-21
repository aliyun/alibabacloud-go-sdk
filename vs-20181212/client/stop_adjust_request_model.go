// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdjustRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFocus(v bool) *StopAdjustRequest
	GetFocus() *bool
	SetId(v string) *StopAdjustRequest
	GetId() *string
	SetIris(v bool) *StopAdjustRequest
	GetIris() *bool
	SetOwnerId(v int64) *StopAdjustRequest
	GetOwnerId() *int64
}

type StopAdjustRequest struct {
	// example:
	//
	// true
	Focus *bool `json:"Focus,omitempty" xml:"Focus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	Iris    *bool  `json:"Iris,omitempty" xml:"Iris,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StopAdjustRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAdjustRequest) GoString() string {
	return s.String()
}

func (s *StopAdjustRequest) GetFocus() *bool {
	return s.Focus
}

func (s *StopAdjustRequest) GetId() *string {
	return s.Id
}

func (s *StopAdjustRequest) GetIris() *bool {
	return s.Iris
}

func (s *StopAdjustRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopAdjustRequest) SetFocus(v bool) *StopAdjustRequest {
	s.Focus = &v
	return s
}

func (s *StopAdjustRequest) SetId(v string) *StopAdjustRequest {
	s.Id = &v
	return s
}

func (s *StopAdjustRequest) SetIris(v bool) *StopAdjustRequest {
	s.Iris = &v
	return s
}

func (s *StopAdjustRequest) SetOwnerId(v int64) *StopAdjustRequest {
	s.OwnerId = &v
	return s
}

func (s *StopAdjustRequest) Validate() error {
	return dara.Validate(s)
}
