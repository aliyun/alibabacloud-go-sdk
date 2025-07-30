// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangePasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginToken(v string) *ChangePasswordResponseBody
	GetLoginToken() *string
	SetRequestId(v string) *ChangePasswordResponseBody
	GetRequestId() *string
}

type ChangePasswordResponseBody struct {
	// The logon token.
	//
	// example:
	//
	// v18101ac6a9e69c66b04a163031680463660b4b216cd758f34b60b9ad6a7c7f7334b83dd8f75eef4209c68f9f1080b****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 484256DA-D816-44D2-9D86-B6EE4D5B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) GetLoginToken() *string {
	return s.LoginToken
}

func (s *ChangePasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangePasswordResponseBody) SetLoginToken(v string) *ChangePasswordResponseBody {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangePasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
