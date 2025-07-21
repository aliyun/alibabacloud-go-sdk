// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAccountLessLoginUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndUserId(v string) *BindAccountLessLoginUserRequest
	GetEndUserId() *string
	SetSerialNumber(v string) *BindAccountLessLoginUserRequest
	GetSerialNumber() *string
	SetUuid(v string) *BindAccountLessLoginUserRequest
	GetUuid() *string
}

type BindAccountLessLoginUserRequest struct {
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s BindAccountLessLoginUserRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAccountLessLoginUserRequest) GoString() string {
	return s.String()
}

func (s *BindAccountLessLoginUserRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *BindAccountLessLoginUserRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *BindAccountLessLoginUserRequest) GetUuid() *string {
	return s.Uuid
}

func (s *BindAccountLessLoginUserRequest) SetEndUserId(v string) *BindAccountLessLoginUserRequest {
	s.EndUserId = &v
	return s
}

func (s *BindAccountLessLoginUserRequest) SetSerialNumber(v string) *BindAccountLessLoginUserRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindAccountLessLoginUserRequest) SetUuid(v string) *BindAccountLessLoginUserRequest {
	s.Uuid = &v
	return s
}

func (s *BindAccountLessLoginUserRequest) Validate() error {
	return dara.Validate(s)
}
