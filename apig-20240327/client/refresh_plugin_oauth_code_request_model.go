// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshPluginOAuthCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshPluginOAuthCodeRequest
	GetCode() *string
}

type RefreshPluginOAuthCodeRequest struct {
	// The authorization code returned by the Codeup OAuth authorization callback (the code parameter carried in the callback after authorization at https://codeup.aliyun.com/oauth/authorize?app_id=xxx&state=apigw). If this parameter is empty, the ErrInvalidParameterWithEmpty("oauth code") error is returned.
	//
	// example:
	//
	// 4/0AX4xxxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s RefreshPluginOAuthCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshPluginOAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *RefreshPluginOAuthCodeRequest) GetCode() *string {
	return s.Code
}

func (s *RefreshPluginOAuthCodeRequest) SetCode(v string) *RefreshPluginOAuthCodeRequest {
	s.Code = &v
	return s
}

func (s *RefreshPluginOAuthCodeRequest) Validate() error {
	return dara.Validate(s)
}
