// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPasswordFreeLoginUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *BindPasswordFreeLoginUserRequest
	GetEndUserId() *string
	SetMainBizType(v string) *BindPasswordFreeLoginUserRequest
	GetMainBizType() *string
	SetSerialNumber(v string) *BindPasswordFreeLoginUserRequest
	GetSerialNumber() *string
	SetUuid(v string) *BindPasswordFreeLoginUserRequest
	GetUuid() *string
}

type BindPasswordFreeLoginUserRequest struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	MainBizType  *string `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s BindPasswordFreeLoginUserRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPasswordFreeLoginUserRequest) GoString() string {
	return s.String()
}

func (s *BindPasswordFreeLoginUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *BindPasswordFreeLoginUserRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *BindPasswordFreeLoginUserRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *BindPasswordFreeLoginUserRequest) GetUuid() *string {
	return s.Uuid
}

func (s *BindPasswordFreeLoginUserRequest) SetEndUserId(v string) *BindPasswordFreeLoginUserRequest {
	s.EndUserId = &v
	return s
}

func (s *BindPasswordFreeLoginUserRequest) SetMainBizType(v string) *BindPasswordFreeLoginUserRequest {
	s.MainBizType = &v
	return s
}

func (s *BindPasswordFreeLoginUserRequest) SetSerialNumber(v string) *BindPasswordFreeLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindPasswordFreeLoginUserRequest) SetUuid(v string) *BindPasswordFreeLoginUserRequest {
	s.Uuid = &v
	return s
}

func (s *BindPasswordFreeLoginUserRequest) Validate() error {
	return dara.Validate(s)
}
