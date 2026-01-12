// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmpId(v string) *LoginAppRequest
	GetEmpId() *string
	SetEmpName(v string) *LoginAppRequest
	GetEmpName() *string
	SetToken(v string) *LoginAppRequest
	GetToken() *string
	SetType(v string) *LoginAppRequest
	GetType() *string
}

type LoginAppRequest struct {
	EmpId   *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	EmpName *string `json:"EmpName,omitempty" xml:"EmpName,omitempty"`
	Token   *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LoginAppRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginAppRequest) GoString() string {
	return s.String()
}

func (s *LoginAppRequest) GetEmpId() *string {
	return s.EmpId
}

func (s *LoginAppRequest) GetEmpName() *string {
	return s.EmpName
}

func (s *LoginAppRequest) GetToken() *string {
	return s.Token
}

func (s *LoginAppRequest) GetType() *string {
	return s.Type
}

func (s *LoginAppRequest) SetEmpId(v string) *LoginAppRequest {
	s.EmpId = &v
	return s
}

func (s *LoginAppRequest) SetEmpName(v string) *LoginAppRequest {
	s.EmpName = &v
	return s
}

func (s *LoginAppRequest) SetToken(v string) *LoginAppRequest {
	s.Token = &v
	return s
}

func (s *LoginAppRequest) SetType(v string) *LoginAppRequest {
	s.Type = &v
	return s
}

func (s *LoginAppRequest) Validate() error {
	return dara.Validate(s)
}
