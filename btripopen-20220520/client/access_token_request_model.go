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
	// The API secret of the application. For information about how to obtain the secret, see [Application credentials](/#/document/server/application-of-basic-information?handbookId=development-support).
	//
	// This parameter is required.
	//
	// example:
	//
	// ggZADk********eDxdaD
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
