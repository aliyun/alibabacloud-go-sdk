// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAppSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *ResetAppSecretRequest
	GetAppKey() *string
	SetNewAppKey(v string) *ResetAppSecretRequest
	GetNewAppKey() *string
	SetNewAppSecret(v string) *ResetAppSecretRequest
	GetNewAppSecret() *string
	SetSecurityToken(v string) *ResetAppSecretRequest
	GetSecurityToken() *string
}

type ResetAppSecretRequest struct {
	// The key of the application that is used to make an API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60030986
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The new AppKey that you set must be globally unique.
	//
	// example:
	//
	// testAppKey
	NewAppKey *string `json:"NewAppKey,omitempty" xml:"NewAppKey,omitempty"`
	// The new key of the application. To improve compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// test***
	NewAppSecret  *string `json:"NewAppSecret,omitempty" xml:"NewAppSecret,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *ResetAppSecretRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *ResetAppSecretRequest) GetNewAppKey() *string {
	return s.NewAppKey
}

func (s *ResetAppSecretRequest) GetNewAppSecret() *string {
	return s.NewAppSecret
}

func (s *ResetAppSecretRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ResetAppSecretRequest) SetAppKey(v string) *ResetAppSecretRequest {
	s.AppKey = &v
	return s
}

func (s *ResetAppSecretRequest) SetNewAppKey(v string) *ResetAppSecretRequest {
	s.NewAppKey = &v
	return s
}

func (s *ResetAppSecretRequest) SetNewAppSecret(v string) *ResetAppSecretRequest {
	s.NewAppSecret = &v
	return s
}

func (s *ResetAppSecretRequest) SetSecurityToken(v string) *ResetAppSecretRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetAppSecretRequest) Validate() error {
	return dara.Validate(s)
}
