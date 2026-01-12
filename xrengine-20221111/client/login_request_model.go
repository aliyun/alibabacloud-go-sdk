// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmpId(v string) *LoginRequest
	GetEmpId() *string
	SetEmpName(v string) *LoginRequest
	GetEmpName() *string
	SetToken(v string) *LoginRequest
	GetToken() *string
	SetType(v string) *LoginRequest
	GetType() *string
}

type LoginRequest struct {
	EmpId   *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	EmpName *string `json:"EmpName,omitempty" xml:"EmpName,omitempty"`
	// This parameter is required.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LoginRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginRequest) GoString() string {
	return s.String()
}

func (s *LoginRequest) GetEmpId() *string {
	return s.EmpId
}

func (s *LoginRequest) GetEmpName() *string {
	return s.EmpName
}

func (s *LoginRequest) GetToken() *string {
	return s.Token
}

func (s *LoginRequest) GetType() *string {
	return s.Type
}

func (s *LoginRequest) SetEmpId(v string) *LoginRequest {
	s.EmpId = &v
	return s
}

func (s *LoginRequest) SetEmpName(v string) *LoginRequest {
	s.EmpName = &v
	return s
}

func (s *LoginRequest) SetToken(v string) *LoginRequest {
	s.Token = &v
	return s
}

func (s *LoginRequest) SetType(v string) *LoginRequest {
	s.Type = &v
	return s
}

func (s *LoginRequest) Validate() error {
	return dara.Validate(s)
}
