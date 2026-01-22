// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEcIdAccountId interface {
  dara.Model
  String() string
  GoString() string
  SetAccountIds(v []*int64) *EcIdAccountId
  GetAccountIds() []*int64 
  SetEcId(v string) *EcIdAccountId
  GetEcId() *string 
}

type EcIdAccountId struct {
  AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s EcIdAccountId) String() string {
  return dara.Prettify(s)
}

func (s EcIdAccountId) GoString() string {
  return s.String()
}

func (s *EcIdAccountId) GetAccountIds() []*int64  {
  return s.AccountIds
}

func (s *EcIdAccountId) GetEcId() *string  {
  return s.EcId
}

func (s *EcIdAccountId) SetAccountIds(v []*int64) *EcIdAccountId {
  s.AccountIds = v
  return s
}

func (s *EcIdAccountId) SetEcId(v string) *EcIdAccountId {
  s.EcId = &v
  return s
}

func (s *EcIdAccountId) Validate() error {
  return dara.Validate(s)
}

