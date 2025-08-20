// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPullCashierShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceInfoShrink(v string) *PullCashierShrinkRequest
	GetDeviceInfoShrink() *string
	SetPayloadShrink(v string) *PullCashierShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *PullCashierShrinkRequest
	GetUserInfoShrink() *string
}

type PullCashierShrinkRequest struct {
	DeviceInfoShrink *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	PayloadShrink    *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s PullCashierShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PullCashierShrinkRequest) GoString() string {
	return s.String()
}

func (s *PullCashierShrinkRequest) GetDeviceInfoShrink() *string {
	return s.DeviceInfoShrink
}

func (s *PullCashierShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *PullCashierShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *PullCashierShrinkRequest) SetDeviceInfoShrink(v string) *PullCashierShrinkRequest {
	s.DeviceInfoShrink = &v
	return s
}

func (s *PullCashierShrinkRequest) SetPayloadShrink(v string) *PullCashierShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *PullCashierShrinkRequest) SetUserInfoShrink(v string) *PullCashierShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *PullCashierShrinkRequest) Validate() error {
	return dara.Validate(s)
}
