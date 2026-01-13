// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountCredentialDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AccountCredentialDTO
	GetAccountId() *string
	SetCallerIdentity(v string) *AccountCredentialDTO
	GetCallerIdentity() *string
	SetId(v int64) *AccountCredentialDTO
	GetId() *int64
}

type AccountCredentialDTO struct {
	AccountId      *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CallerIdentity *string `json:"callerIdentity,omitempty" xml:"callerIdentity,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
}

func (s AccountCredentialDTO) String() string {
	return dara.Prettify(s)
}

func (s AccountCredentialDTO) GoString() string {
	return s.String()
}

func (s *AccountCredentialDTO) GetAccountId() *string {
	return s.AccountId
}

func (s *AccountCredentialDTO) GetCallerIdentity() *string {
	return s.CallerIdentity
}

func (s *AccountCredentialDTO) GetId() *int64 {
	return s.Id
}

func (s *AccountCredentialDTO) SetAccountId(v string) *AccountCredentialDTO {
	s.AccountId = &v
	return s
}

func (s *AccountCredentialDTO) SetCallerIdentity(v string) *AccountCredentialDTO {
	s.CallerIdentity = &v
	return s
}

func (s *AccountCredentialDTO) SetId(v int64) *AccountCredentialDTO {
	s.Id = &v
	return s
}

func (s *AccountCredentialDTO) Validate() error {
	return dara.Validate(s)
}
