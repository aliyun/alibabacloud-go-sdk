// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAlias(v string) *GetAccountAliasResponseBody
	GetAccountAlias() *string
	SetRequestId(v string) *GetAccountAliasResponseBody
	GetRequestId() *string
}

type GetAccountAliasResponseBody struct {
	// The alias of the Alibaba Cloud account.
	//
	// example:
	//
	// myalias
	AccountAlias *string `json:"AccountAlias,omitempty" xml:"AccountAlias,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountAliasResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountAliasResponseBody) GetAccountAlias() *string {
	return s.AccountAlias
}

func (s *GetAccountAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountAliasResponseBody) SetAccountAlias(v string) *GetAccountAliasResponseBody {
	s.AccountAlias = &v
	return s
}

func (s *GetAccountAliasResponseBody) SetRequestId(v string) *GetAccountAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
