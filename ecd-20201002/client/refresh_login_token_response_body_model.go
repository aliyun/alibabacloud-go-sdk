// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshLoginTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginToken(v string) *RefreshLoginTokenResponseBody
	GetLoginToken() *string
	SetRequestId(v string) *RefreshLoginTokenResponseBody
	GetRequestId() *string
}

type RefreshLoginTokenResponseBody struct {
	// example:
	//
	// v1c27bab6c205b2fdfac916434306375722776d6aa89e30b7836d18c95ade9137f0f5ac4325260782184e96ee2b3f0****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshLoginTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponseBody) GetLoginToken() *string {
	return s.LoginToken
}

func (s *RefreshLoginTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshLoginTokenResponseBody) SetLoginToken(v string) *RefreshLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) SetRequestId(v string) *RefreshLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
