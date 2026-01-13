// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExpendExpiredTimeRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccountId(v string) *ExpendExpiredTimeRequest
  GetAccountId() *string 
  SetTestStartTime(v string) *ExpendExpiredTimeRequest
  GetTestStartTime() *string 
}

type ExpendExpiredTimeRequest struct {
  AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
  TestStartTime *string `json:"testStartTime,omitempty" xml:"testStartTime,omitempty"`
}

func (s ExpendExpiredTimeRequest) String() string {
  return dara.Prettify(s)
}

func (s ExpendExpiredTimeRequest) GoString() string {
  return s.String()
}

func (s *ExpendExpiredTimeRequest) GetAccountId() *string  {
  return s.AccountId
}

func (s *ExpendExpiredTimeRequest) GetTestStartTime() *string  {
  return s.TestStartTime
}

func (s *ExpendExpiredTimeRequest) SetAccountId(v string) *ExpendExpiredTimeRequest {
  s.AccountId = &v
  return s
}

func (s *ExpendExpiredTimeRequest) SetTestStartTime(v string) *ExpendExpiredTimeRequest {
  s.TestStartTime = &v
  return s
}

func (s *ExpendExpiredTimeRequest) Validate() error {
  return dara.Validate(s)
}

