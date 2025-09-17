// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTokenCode(v string) *CreateAccessTokenResponseBody
	GetAccessTokenCode() *string
	SetAccessTokenId(v string) *CreateAccessTokenResponseBody
	GetAccessTokenId() *string
	SetRequestId(v string) *CreateAccessTokenResponseBody
	GetRequestId() *string
}

type CreateAccessTokenResponseBody struct {
	// The value of the activation code. The value is returned only once after the CreateAccessToken operation is called and cannot be subsequently queried. Make sure that you properly save the returned value.
	//
	// example:
	//
	// B57QoTXEA2Tytr0uZWoNY5Aju5Jt****
	AccessTokenCode *string `json:"AccessTokenCode,omitempty" xml:"AccessTokenCode,omitempty"`
	// The ID of the activation code.
	//
	// example:
	//
	// at-bp1akz2zp67r0k6r****
	AccessTokenId *string `json:"AccessTokenId,omitempty" xml:"AccessTokenId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB4A7EA2-6FDA-5655-B067-854532FB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessTokenResponseBody) GetAccessTokenCode() *string {
	return s.AccessTokenCode
}

func (s *CreateAccessTokenResponseBody) GetAccessTokenId() *string {
	return s.AccessTokenId
}

func (s *CreateAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessTokenResponseBody) SetAccessTokenCode(v string) *CreateAccessTokenResponseBody {
	s.AccessTokenCode = &v
	return s
}

func (s *CreateAccessTokenResponseBody) SetAccessTokenId(v string) *CreateAccessTokenResponseBody {
	s.AccessTokenId = &v
	return s
}

func (s *CreateAccessTokenResponseBody) SetRequestId(v string) *CreateAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
