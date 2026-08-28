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
