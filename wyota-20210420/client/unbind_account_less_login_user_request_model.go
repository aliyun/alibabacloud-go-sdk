// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAccountLessLoginUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSerialNumber(v string) *UnbindAccountLessLoginUserRequest
	GetSerialNumber() *string
	SetUuid(v string) *UnbindAccountLessLoginUserRequest
	GetUuid() *string
}

type UnbindAccountLessLoginUserRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s UnbindAccountLessLoginUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAccountLessLoginUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindAccountLessLoginUserRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *UnbindAccountLessLoginUserRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UnbindAccountLessLoginUserRequest) SetSerialNumber(v string) *UnbindAccountLessLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *UnbindAccountLessLoginUserRequest) SetUuid(v string) *UnbindAccountLessLoginUserRequest {
	s.Uuid = &v
	return s
}

func (s *UnbindAccountLessLoginUserRequest) Validate() error {
	return dara.Validate(s)
}
