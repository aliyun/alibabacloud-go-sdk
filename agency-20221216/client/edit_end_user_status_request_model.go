// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditEndUserStatusRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCreditStatus(v string) *EditEndUserStatusRequest
  GetCreditStatus() *string 
  SetUid(v int64) *EditEndUserStatusRequest
  GetUid() *int64 
}

type EditEndUserStatusRequest struct {
  // ### Shutdown status
  // 
  // - postPayFreeze: pay-as-you-go resource suspension
  // 
  // - postPayThaw: pay-as-you-go resource resumption.
  // 
  // example:
  // 
  // postPayFreeze
  CreditStatus *string `json:"CreditStatus,omitempty" xml:"CreditStatus,omitempty"`
  // The UID of the customer to be processed.
  // 
  // example:
  // 
  // 1792155717328010
  Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s EditEndUserStatusRequest) String() string {
  return dara.Prettify(s)
}

func (s EditEndUserStatusRequest) GoString() string {
  return s.String()
}

func (s *EditEndUserStatusRequest) GetCreditStatus() *string  {
  return s.CreditStatus
}

func (s *EditEndUserStatusRequest) GetUid() *int64  {
  return s.Uid
}

func (s *EditEndUserStatusRequest) SetCreditStatus(v string) *EditEndUserStatusRequest {
  s.CreditStatus = &v
  return s
}

func (s *EditEndUserStatusRequest) SetUid(v int64) *EditEndUserStatusRequest {
  s.Uid = &v
  return s
}

func (s *EditEndUserStatusRequest) Validate() error {
  return dara.Validate(s)
}

