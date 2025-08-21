// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAligenieUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginStateAccessToken(v string) *GetAligenieUserInfoRequest
	GetLoginStateAccessToken() *string
}

type GetAligenieUserInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 99ce8a70c23a94f8569e1a525bef6e85
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s GetAligenieUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAligenieUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAligenieUserInfoRequest) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *GetAligenieUserInfoRequest) SetLoginStateAccessToken(v string) *GetAligenieUserInfoRequest {
	s.LoginStateAccessToken = &v
	return s
}

func (s *GetAligenieUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
