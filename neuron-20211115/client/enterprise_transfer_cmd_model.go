// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTransferCmd interface {
  dara.Model
  String() string
  GoString() string
  SetAccountId(v string) *EnterpriseTransferCmd
  GetAccountId() *string 
  SetId(v int64) *EnterpriseTransferCmd
  GetId() *int64 
}

type EnterpriseTransferCmd struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 121321412341
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s EnterpriseTransferCmd) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTransferCmd) GoString() string {
  return s.String()
}

func (s *EnterpriseTransferCmd) GetAccountId() *string  {
  return s.AccountId
}

func (s *EnterpriseTransferCmd) GetId() *int64  {
  return s.Id
}

func (s *EnterpriseTransferCmd) SetAccountId(v string) *EnterpriseTransferCmd {
  s.AccountId = &v
  return s
}

func (s *EnterpriseTransferCmd) SetId(v int64) *EnterpriseTransferCmd {
  s.Id = &v
  return s
}

func (s *EnterpriseTransferCmd) Validate() error {
  return dara.Validate(s)
}

