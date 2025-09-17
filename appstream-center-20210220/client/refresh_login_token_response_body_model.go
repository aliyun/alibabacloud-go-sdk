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
	// v12369636c721ba6b3ddb1683341016775c3f63e4d0e78f120f9a0544ed826b7af7daf747c402f0d0730b52f451b70****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// example:
	//
	// 419F31B9-1FDF-5644-ABA3-D00026FA****
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
