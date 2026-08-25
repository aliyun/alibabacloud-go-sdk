// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDelegateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DisableDelegateAccountRequest
	GetAccountId() *string
}

type DisableDelegateAccountRequest struct {
	// The ID of the delegated administrator account of CloudSSO.
	//
	// example:
	//
	// 1200971777065046
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DisableDelegateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDelegateAccountRequest) GoString() string {
	return s.String()
}

func (s *DisableDelegateAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DisableDelegateAccountRequest) SetAccountId(v string) *DisableDelegateAccountRequest {
	s.AccountId = &v
	return s
}

func (s *DisableDelegateAccountRequest) Validate() error {
	return dara.Validate(s)
}
