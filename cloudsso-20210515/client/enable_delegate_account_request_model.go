// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableDelegateAccountRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountId(v string) *EnableDelegateAccountRequest
  GetAccountId() *string 
}

type EnableDelegateAccountRequest struct {
  // The ID of the delegated administrator account of CloudSSO.
  // 
  // example:
  // 
  // 180658567986****
  AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s EnableDelegateAccountRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableDelegateAccountRequest) GoString() string {
  return s.String()
}

func (s *EnableDelegateAccountRequest) GetAccountId() *string  {
  return s.AccountId
}

func (s *EnableDelegateAccountRequest) SetAccountId(v string) *EnableDelegateAccountRequest {
  s.AccountId = &v
  return s
}

func (s *EnableDelegateAccountRequest) Validate() error {
  return dara.Validate(s)
}

