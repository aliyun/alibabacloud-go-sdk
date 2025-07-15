// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ResetAppCodeRequest
	GetAppCode() *string
	SetNewAppCode(v string) *ResetAppCodeRequest
	GetNewAppCode() *string
	SetSecurityToken(v string) *ResetAppCodeRequest
	GetSecurityToken() *string
}

type ResetAppCodeRequest struct {
	// The AppCode of the app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3aaf905a0a1f4f0eabc6d891dfa08afc
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The new AppCode of the app.
	//
	// example:
	//
	// 6f0e7ab2aa5f4b8fb18421e6edf4fb6c2
	NewAppCode    *string `json:"NewAppCode,omitempty" xml:"NewAppCode,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAppCodeRequest) GoString() string {
	return s.String()
}

func (s *ResetAppCodeRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *ResetAppCodeRequest) GetNewAppCode() *string {
	return s.NewAppCode
}

func (s *ResetAppCodeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResetAppCodeRequest) SetAppCode(v string) *ResetAppCodeRequest {
	s.AppCode = &v
	return s
}

func (s *ResetAppCodeRequest) SetNewAppCode(v string) *ResetAppCodeRequest {
	s.NewAppCode = &v
	return s
}

func (s *ResetAppCodeRequest) SetSecurityToken(v string) *ResetAppCodeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetAppCodeRequest) Validate() error {
	return dara.Validate(s)
}
