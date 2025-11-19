// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *CredentialConfiguration
	GetCredentialName() *string
}

type CredentialConfiguration struct {
	// 凭证的唯一标识符
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
}

func (s CredentialConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CredentialConfiguration) GoString() string {
	return s.String()
}

func (s *CredentialConfiguration) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CredentialConfiguration) SetCredentialName(v string) *CredentialConfiguration {
	s.CredentialName = &v
	return s
}

func (s *CredentialConfiguration) Validate() error {
	return dara.Validate(s)
}
