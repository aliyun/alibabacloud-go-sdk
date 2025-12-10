// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedServicesForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListDelegatedServicesForAccountRequest
	GetAccountId() *string
}

type ListDelegatedServicesForAccountRequest struct {
	// The ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ListDelegatedServicesForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDelegatedServicesForAccountRequest) SetAccountId(v string) *ListDelegatedServicesForAccountRequest {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedServicesForAccountRequest) Validate() error {
	return dara.Validate(s)
}
