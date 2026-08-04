// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllowAgAccountLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *AllowAgAccountLoginRequest
	GetAgAccountType() *string
	SetAppName(v string) *AllowAgAccountLoginRequest
	GetAppName() *string
	SetMpk(v string) *AllowAgAccountLoginRequest
	GetMpk() *string
	SetPk(v string) *AllowAgAccountLoginRequest
	GetPk() *string
}

type AllowAgAccountLoginRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s AllowAgAccountLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s AllowAgAccountLoginRequest) GoString() string {
	return s.String()
}

func (s *AllowAgAccountLoginRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *AllowAgAccountLoginRequest) GetAppName() *string {
	return s.AppName
}

func (s *AllowAgAccountLoginRequest) GetMpk() *string {
	return s.Mpk
}

func (s *AllowAgAccountLoginRequest) GetPk() *string {
	return s.Pk
}

func (s *AllowAgAccountLoginRequest) SetAgAccountType(v string) *AllowAgAccountLoginRequest {
	s.AgAccountType = &v
	return s
}

func (s *AllowAgAccountLoginRequest) SetAppName(v string) *AllowAgAccountLoginRequest {
	s.AppName = &v
	return s
}

func (s *AllowAgAccountLoginRequest) SetMpk(v string) *AllowAgAccountLoginRequest {
	s.Mpk = &v
	return s
}

func (s *AllowAgAccountLoginRequest) SetPk(v string) *AllowAgAccountLoginRequest {
	s.Pk = &v
	return s
}

func (s *AllowAgAccountLoginRequest) Validate() error {
	return dara.Validate(s)
}
