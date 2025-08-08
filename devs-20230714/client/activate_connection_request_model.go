// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v *GitAccount) *ActivateConnectionRequest
	GetAccount() *GitAccount
	SetCredential(v *OAuthCredential) *ActivateConnectionRequest
	GetCredential() *OAuthCredential
}

type ActivateConnectionRequest struct {
	Account    *GitAccount      `json:"account,omitempty" xml:"account,omitempty"`
	Credential *OAuthCredential `json:"credential,omitempty" xml:"credential,omitempty"`
}

func (s ActivateConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateConnectionRequest) GoString() string {
	return s.String()
}

func (s *ActivateConnectionRequest) GetAccount() *GitAccount {
	return s.Account
}

func (s *ActivateConnectionRequest) GetCredential() *OAuthCredential {
	return s.Credential
}

func (s *ActivateConnectionRequest) SetAccount(v *GitAccount) *ActivateConnectionRequest {
	s.Account = v
	return s
}

func (s *ActivateConnectionRequest) SetCredential(v *OAuthCredential) *ActivateConnectionRequest {
	s.Credential = v
	return s
}

func (s *ActivateConnectionRequest) Validate() error {
	return dara.Validate(s)
}
