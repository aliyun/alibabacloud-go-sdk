// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAligenieUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginStateAccessToken(v string) *UnbindAligenieUserRequest
	GetLoginStateAccessToken() *string
}

type UnbindAligenieUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 99ce8a70c23a94f8569e1a525bef6e85
	LoginStateAccessToken *string `json:"LoginStateAccessToken,omitempty" xml:"LoginStateAccessToken,omitempty"`
}

func (s UnbindAligenieUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAligenieUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindAligenieUserRequest) GetLoginStateAccessToken() *string {
	return s.LoginStateAccessToken
}

func (s *UnbindAligenieUserRequest) SetLoginStateAccessToken(v string) *UnbindAligenieUserRequest {
	s.LoginStateAccessToken = &v
	return s
}

func (s *UnbindAligenieUserRequest) Validate() error {
	return dara.Validate(s)
}
