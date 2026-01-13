// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AccountCreateCmd
	GetAccountId() *string
}

type AccountCreateCmd struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AccountCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s AccountCreateCmd) GoString() string {
	return s.String()
}

func (s *AccountCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *AccountCreateCmd) SetAccountId(v string) *AccountCreateCmd {
	s.AccountId = &v
	return s
}

func (s *AccountCreateCmd) Validate() error {
	return dara.Validate(s)
}
