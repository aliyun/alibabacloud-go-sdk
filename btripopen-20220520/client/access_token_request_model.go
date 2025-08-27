// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecret(v string) *AccessTokenRequest
	GetAppSecret() *string
}

type AccessTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ggZADkghsadgogeDxdaD
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
}

func (s AccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s AccessTokenRequest) GoString() string {
	return s.String()
}

func (s *AccessTokenRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *AccessTokenRequest) SetAppSecret(v string) *AccessTokenRequest {
	s.AppSecret = &v
	return s
}

func (s *AccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
