// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPasswordFreeLoginUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMainBizType(v string) *UnbindPasswordFreeLoginUserRequest
	GetMainBizType() *string
	SetSerialNumber(v string) *UnbindPasswordFreeLoginUserRequest
	GetSerialNumber() *string
	SetUuid(v string) *UnbindPasswordFreeLoginUserRequest
	GetUuid() *string
}

type UnbindPasswordFreeLoginUserRequest struct {
	MainBizType  *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UnbindPasswordFreeLoginUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindPasswordFreeLoginUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindPasswordFreeLoginUserRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *UnbindPasswordFreeLoginUserRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UnbindPasswordFreeLoginUserRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UnbindPasswordFreeLoginUserRequest) SetMainBizType(v string) *UnbindPasswordFreeLoginUserRequest {
	s.MainBizType = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserRequest) SetSerialNumber(v string) *UnbindPasswordFreeLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserRequest) SetUuid(v string) *UnbindPasswordFreeLoginUserRequest {
	s.Uuid = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserRequest) Validate() error {
	return dara.Validate(s)
}
