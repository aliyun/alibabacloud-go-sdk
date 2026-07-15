// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryV2Request interface {
  dara.Model
  String() string
  GoString() string
  SetBizType(v string) *EstimatedPriceQueryV2Request
  GetBizType() *string 
  SetDepartDate(v string) *EstimatedPriceQueryV2Request
  GetDepartDate() *string 
  SetFromCity(v string) *EstimatedPriceQueryV2Request
  GetFromCity() *string 
  SetLeaveDate(v string) *EstimatedPriceQueryV2Request
  GetLeaveDate() *string 
  SetToCity(v string) *EstimatedPriceQueryV2Request
  GetToCity() *string 
  SetUserId(v string) *EstimatedPriceQueryV2Request
  GetUserId() *string 
}

type EstimatedPriceQueryV2Request struct {
  // This parameter is required.
  BizType *string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
  // This parameter is required.
  DepartDate *string `json:"depart_date,omitempty" xml:"depart_date,omitempty"`
  // This parameter is required.
  FromCity *string `json:"from_city,omitempty" xml:"from_city,omitempty"`
  // This parameter is required.
  LeaveDate *string `json:"leave_date,omitempty" xml:"leave_date,omitempty"`
  // This parameter is required.
  ToCity *string `json:"to_city,omitempty" xml:"to_city,omitempty"`
  UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s EstimatedPriceQueryV2Request) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryV2Request) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryV2Request) GetBizType() *string  {
  return s.BizType
}

func (s *EstimatedPriceQueryV2Request) GetDepartDate() *string  {
  return s.DepartDate
}

func (s *EstimatedPriceQueryV2Request) GetFromCity() *string  {
  return s.FromCity
}

func (s *EstimatedPriceQueryV2Request) GetLeaveDate() *string  {
  return s.LeaveDate
}

func (s *EstimatedPriceQueryV2Request) GetToCity() *string  {
  return s.ToCity
}

func (s *EstimatedPriceQueryV2Request) GetUserId() *string  {
  return s.UserId
}

func (s *EstimatedPriceQueryV2Request) SetBizType(v string) *EstimatedPriceQueryV2Request {
  s.BizType = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) SetDepartDate(v string) *EstimatedPriceQueryV2Request {
  s.DepartDate = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) SetFromCity(v string) *EstimatedPriceQueryV2Request {
  s.FromCity = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) SetLeaveDate(v string) *EstimatedPriceQueryV2Request {
  s.LeaveDate = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) SetToCity(v string) *EstimatedPriceQueryV2Request {
  s.ToCity = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) SetUserId(v string) *EstimatedPriceQueryV2Request {
  s.UserId = &v
  return s
}

func (s *EstimatedPriceQueryV2Request) Validate() error {
  return dara.Validate(s)
}

