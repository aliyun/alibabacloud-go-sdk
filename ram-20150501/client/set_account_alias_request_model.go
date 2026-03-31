// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAlias(v string) *SetAccountAliasRequest
	GetAccountAlias() *string
}

type SetAccountAliasRequest struct {
	// The alias of the Alibaba Cloud account.
	//
	// The alias must be 3 to 32 characters in length, and can contain lowercase letters, digits, and hyphens (-).
	//
	// >  It cannot start or end with a hyphen (-), and cannot contain consecutive hyphens (-).
	//
	// example:
	//
	// myalias
	AccountAlias *string `json:"AccountAlias,omitempty" xml:"AccountAlias,omitempty"`
}

func (s SetAccountAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAliasRequest) GoString() string {
	return s.String()
}

func (s *SetAccountAliasRequest) GetAccountAlias() *string {
	return s.AccountAlias
}

func (s *SetAccountAliasRequest) SetAccountAlias(v string) *SetAccountAliasRequest {
	s.AccountAlias = &v
	return s
}

func (s *SetAccountAliasRequest) Validate() error {
	return dara.Validate(s)
}
