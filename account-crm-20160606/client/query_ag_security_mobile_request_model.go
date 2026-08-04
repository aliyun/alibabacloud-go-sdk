// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgSecurityMobileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *QueryAgSecurityMobileRequest
	GetAgAccountType() *string
	SetAppName(v string) *QueryAgSecurityMobileRequest
	GetAppName() *string
	SetMpk(v string) *QueryAgSecurityMobileRequest
	GetMpk() *string
	SetPk(v string) *QueryAgSecurityMobileRequest
	GetPk() *string
}

type QueryAgSecurityMobileRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QueryAgSecurityMobileRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAgSecurityMobileRequest) GoString() string {
	return s.String()
}

func (s *QueryAgSecurityMobileRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *QueryAgSecurityMobileRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryAgSecurityMobileRequest) GetMpk() *string {
	return s.Mpk
}

func (s *QueryAgSecurityMobileRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryAgSecurityMobileRequest) SetAgAccountType(v string) *QueryAgSecurityMobileRequest {
	s.AgAccountType = &v
	return s
}

func (s *QueryAgSecurityMobileRequest) SetAppName(v string) *QueryAgSecurityMobileRequest {
	s.AppName = &v
	return s
}

func (s *QueryAgSecurityMobileRequest) SetMpk(v string) *QueryAgSecurityMobileRequest {
	s.Mpk = &v
	return s
}

func (s *QueryAgSecurityMobileRequest) SetPk(v string) *QueryAgSecurityMobileRequest {
	s.Pk = &v
	return s
}

func (s *QueryAgSecurityMobileRequest) Validate() error {
	return dara.Validate(s)
}
