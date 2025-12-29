// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPayloadShrink(v string) *QueryDeviceStatusShrinkRequest
	GetPayloadShrink() *string
	SetUserInfoShrink(v string) *QueryDeviceStatusShrinkRequest
	GetUserInfoShrink() *string
}

type QueryDeviceStatusShrinkRequest struct {
	PayloadShrink  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s QueryDeviceStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusShrinkRequest) GetPayloadShrink() *string {
	return s.PayloadShrink
}

func (s *QueryDeviceStatusShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *QueryDeviceStatusShrinkRequest) SetPayloadShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.PayloadShrink = &v
	return s
}

func (s *QueryDeviceStatusShrinkRequest) SetUserInfoShrink(v string) *QueryDeviceStatusShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *QueryDeviceStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
