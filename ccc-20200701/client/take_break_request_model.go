// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeBreakRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TakeBreakRequest
	GetCode() *string
	SetDeviceId(v string) *TakeBreakRequest
	GetDeviceId() *string
	SetInstanceId(v string) *TakeBreakRequest
	GetInstanceId() *string
	SetUserId(v string) *TakeBreakRequest
	GetUserId() *string
}

type TakeBreakRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// lunchtime
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s TakeBreakRequest) String() string {
	return dara.Prettify(s)
}

func (s TakeBreakRequest) GoString() string {
	return s.String()
}

func (s *TakeBreakRequest) GetCode() *string {
	return s.Code
}

func (s *TakeBreakRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *TakeBreakRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TakeBreakRequest) GetUserId() *string {
	return s.UserId
}

func (s *TakeBreakRequest) SetCode(v string) *TakeBreakRequest {
	s.Code = &v
	return s
}

func (s *TakeBreakRequest) SetDeviceId(v string) *TakeBreakRequest {
	s.DeviceId = &v
	return s
}

func (s *TakeBreakRequest) SetInstanceId(v string) *TakeBreakRequest {
	s.InstanceId = &v
	return s
}

func (s *TakeBreakRequest) SetUserId(v string) *TakeBreakRequest {
	s.UserId = &v
	return s
}

func (s *TakeBreakRequest) Validate() error {
	return dara.Validate(s)
}
