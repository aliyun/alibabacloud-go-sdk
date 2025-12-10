// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayerForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetPayerForAccountRequest
	GetAccountId() *string
}

type GetPayerForAccountRequest struct {
	// The ID of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s GetPayerForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPayerForAccountRequest) GoString() string {
	return s.String()
}

func (s *GetPayerForAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *GetPayerForAccountRequest) SetAccountId(v string) *GetPayerForAccountRequest {
	s.AccountId = &v
	return s
}

func (s *GetPayerForAccountRequest) Validate() error {
	return dara.Validate(s)
}
