// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditNewBuyStatusRequest interface {
  dara.Model
  String() string
  GoString() string
  SetNewBuyStatus(v string) *EditNewBuyStatusRequest
  GetNewBuyStatus() *string 
  SetUid(v int64) *EditNewBuyStatusRequest
  GetUid() *int64 
}

type EditNewBuyStatusRequest struct {
  // The new purchase status. Valid values:
  // 
  // - cancelBan: allows new purchases
  // 
  // - ban: prohibits new purchases.
  // 
  // example:
  // 
  // cancelBan
  NewBuyStatus *string `json:"NewBuyStatus,omitempty" xml:"NewBuyStatus,omitempty"`
  // The UID of the customer.
  // 
  // example:
  // 
  // 1133166938931507
  Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditNewBuyStatusRequest) String() string {
  return dara.Prettify(s)
}

func (s EditNewBuyStatusRequest) GoString() string {
  return s.String()
}

func (s *EditNewBuyStatusRequest) GetNewBuyStatus() *string  {
  return s.NewBuyStatus
}

func (s *EditNewBuyStatusRequest) GetUid() *int64  {
  return s.Uid
}

func (s *EditNewBuyStatusRequest) SetNewBuyStatus(v string) *EditNewBuyStatusRequest {
  s.NewBuyStatus = &v
  return s
}

func (s *EditNewBuyStatusRequest) SetUid(v int64) *EditNewBuyStatusRequest {
  s.Uid = &v
  return s
}

func (s *EditNewBuyStatusRequest) Validate() error {
  return dara.Validate(s)
}

