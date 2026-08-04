// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ChangeAgSecurityEmailRequest
	GetAppName() *string
	SetMpk(v string) *ChangeAgSecurityEmailRequest
	GetMpk() *string
	SetPk(v string) *ChangeAgSecurityEmailRequest
	GetPk() *string
	SetSecurityEmail(v string) *ChangeAgSecurityEmailRequest
	GetSecurityEmail() *string
}

type ChangeAgSecurityEmailRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	SecurityEmail *string `json:"SecurityEmail,omitempty" xml:"SecurityEmail,omitempty"`
}

func (s ChangeAgSecurityEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityEmailRequest) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityEmailRequest) GetAppName() *string {
	return s.AppName
}

func (s *ChangeAgSecurityEmailRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ChangeAgSecurityEmailRequest) GetPk() *string {
	return s.Pk
}

func (s *ChangeAgSecurityEmailRequest) GetSecurityEmail() *string {
	return s.SecurityEmail
}

func (s *ChangeAgSecurityEmailRequest) SetAppName(v string) *ChangeAgSecurityEmailRequest {
	s.AppName = &v
	return s
}

func (s *ChangeAgSecurityEmailRequest) SetMpk(v string) *ChangeAgSecurityEmailRequest {
	s.Mpk = &v
	return s
}

func (s *ChangeAgSecurityEmailRequest) SetPk(v string) *ChangeAgSecurityEmailRequest {
	s.Pk = &v
	return s
}

func (s *ChangeAgSecurityEmailRequest) SetSecurityEmail(v string) *ChangeAgSecurityEmailRequest {
	s.SecurityEmail = &v
	return s
}

func (s *ChangeAgSecurityEmailRequest) Validate() error {
	return dara.Validate(s)
}
