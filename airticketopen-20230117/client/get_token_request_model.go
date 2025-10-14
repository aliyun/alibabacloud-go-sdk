// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *GetTokenRequest
	GetAppKey() *string
	SetAppSecret(v string) *GetTokenRequest
	GetAppSecret() *string
}

type GetTokenRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// fu1bltcu3400iurywuri
	AppKey *string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// appSecret
	//
	// This parameter is required.
	//
	// example:
	//
	// ZzQ3MW1mb3E1ODAwI2ldUjYlWUdJn5YI
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *GetTokenRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GetTokenRequest) SetAppKey(v string) *GetTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetTokenRequest) SetAppSecret(v string) *GetTokenRequest {
	s.AppSecret = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
