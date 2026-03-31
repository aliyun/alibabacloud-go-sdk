// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAccountAliasResponseBody
	GetRequestId() *string
}

type SetAccountAliasResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccountAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccountAliasResponseBody) SetRequestId(v string) *SetAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
