// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AuthorizeAndroidInstanceResponseBody
	GetRequestId() *string
}

type AuthorizeAndroidInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AuthorizeAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeAndroidInstanceResponseBody) SetRequestId(v string) *AuthorizeAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeAndroidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
