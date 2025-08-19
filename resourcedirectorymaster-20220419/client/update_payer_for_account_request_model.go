// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePayerForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *UpdatePayerForAccountRequest
	GetAccountId() *string
	SetPayerAccountId(v string) *UpdatePayerForAccountRequest
	GetPayerAccountId() *string
}

type UpdatePayerForAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the billing account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12323344****
	PayerAccountId *string `json:"PayerAccountId,omitempty" xml:"PayerAccountId,omitempty"`
}

func (s UpdatePayerForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePayerForAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdatePayerForAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdatePayerForAccountRequest) GetPayerAccountId() *string {
	return s.PayerAccountId
}

func (s *UpdatePayerForAccountRequest) SetAccountId(v string) *UpdatePayerForAccountRequest {
	s.AccountId = &v
	return s
}

func (s *UpdatePayerForAccountRequest) SetPayerAccountId(v string) *UpdatePayerForAccountRequest {
	s.PayerAccountId = &v
	return s
}

func (s *UpdatePayerForAccountRequest) Validate() error {
	return dara.Validate(s)
}
