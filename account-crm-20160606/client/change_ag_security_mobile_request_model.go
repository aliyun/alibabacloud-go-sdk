// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgSecurityMobileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ChangeAgSecurityMobileRequest
	GetAppName() *string
	SetMpk(v string) *ChangeAgSecurityMobileRequest
	GetMpk() *string
	SetPk(v string) *ChangeAgSecurityMobileRequest
	GetPk() *string
	SetSecurityMobile(v string) *ChangeAgSecurityMobileRequest
	GetSecurityMobile() *string
}

type ChangeAgSecurityMobileRequest struct {
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	SecurityMobile *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
}

func (s ChangeAgSecurityMobileRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgSecurityMobileRequest) GoString() string {
	return s.String()
}

func (s *ChangeAgSecurityMobileRequest) GetAppName() *string {
	return s.AppName
}

func (s *ChangeAgSecurityMobileRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ChangeAgSecurityMobileRequest) GetPk() *string {
	return s.Pk
}

func (s *ChangeAgSecurityMobileRequest) GetSecurityMobile() *string {
	return s.SecurityMobile
}

func (s *ChangeAgSecurityMobileRequest) SetAppName(v string) *ChangeAgSecurityMobileRequest {
	s.AppName = &v
	return s
}

func (s *ChangeAgSecurityMobileRequest) SetMpk(v string) *ChangeAgSecurityMobileRequest {
	s.Mpk = &v
	return s
}

func (s *ChangeAgSecurityMobileRequest) SetPk(v string) *ChangeAgSecurityMobileRequest {
	s.Pk = &v
	return s
}

func (s *ChangeAgSecurityMobileRequest) SetSecurityMobile(v string) *ChangeAgSecurityMobileRequest {
	s.SecurityMobile = &v
	return s
}

func (s *ChangeAgSecurityMobileRequest) Validate() error {
	return dara.Validate(s)
}
