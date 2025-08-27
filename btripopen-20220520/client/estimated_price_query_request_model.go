// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEstimatedPriceQueryRequest interface {
  dara.Model
  String() string
  GoString() string
  SetArrCity(v string) *EstimatedPriceQueryRequest
  GetArrCity() *string 
  SetCategory(v string) *EstimatedPriceQueryRequest
  GetCategory() *string 
  SetDepCity(v string) *EstimatedPriceQueryRequest
  GetDepCity() *string 
  SetEndTime(v int64) *EstimatedPriceQueryRequest
  GetEndTime() *int64 
  SetItineraryId(v string) *EstimatedPriceQueryRequest
  GetItineraryId() *string 
  SetStartTime(v int64) *EstimatedPriceQueryRequest
  GetStartTime() *int64 
  SetSubCorpId(v string) *EstimatedPriceQueryRequest
  GetSubCorpId() *string 
  SetUserId(v string) *EstimatedPriceQueryRequest
  GetUserId() *string 
}

type EstimatedPriceQueryRequest struct {
  // This parameter is required.
  ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // flight
  Category *string `json:"category,omitempty" xml:"category,omitempty"`
  // This parameter is required.
  DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1670601600000
  EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
  // example:
  // 
  // 1245
  ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1670428800000
  StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
  // example:
  // 
  // btrip123
  SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 12345678910
  UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s EstimatedPriceQueryRequest) String() string {
  return dara.Prettify(s)
}

func (s EstimatedPriceQueryRequest) GoString() string {
  return s.String()
}

func (s *EstimatedPriceQueryRequest) GetArrCity() *string  {
  return s.ArrCity
}

func (s *EstimatedPriceQueryRequest) GetCategory() *string  {
  return s.Category
}

func (s *EstimatedPriceQueryRequest) GetDepCity() *string  {
  return s.DepCity
}

func (s *EstimatedPriceQueryRequest) GetEndTime() *int64  {
  return s.EndTime
}

func (s *EstimatedPriceQueryRequest) GetItineraryId() *string  {
  return s.ItineraryId
}

func (s *EstimatedPriceQueryRequest) GetStartTime() *int64  {
  return s.StartTime
}

func (s *EstimatedPriceQueryRequest) GetSubCorpId() *string  {
  return s.SubCorpId
}

func (s *EstimatedPriceQueryRequest) GetUserId() *string  {
  return s.UserId
}

func (s *EstimatedPriceQueryRequest) SetArrCity(v string) *EstimatedPriceQueryRequest {
  s.ArrCity = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetCategory(v string) *EstimatedPriceQueryRequest {
  s.Category = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetDepCity(v string) *EstimatedPriceQueryRequest {
  s.DepCity = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetEndTime(v int64) *EstimatedPriceQueryRequest {
  s.EndTime = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetItineraryId(v string) *EstimatedPriceQueryRequest {
  s.ItineraryId = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetStartTime(v int64) *EstimatedPriceQueryRequest {
  s.StartTime = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetSubCorpId(v string) *EstimatedPriceQueryRequest {
  s.SubCorpId = &v
  return s
}

func (s *EstimatedPriceQueryRequest) SetUserId(v string) *EstimatedPriceQueryRequest {
  s.UserId = &v
  return s
}

func (s *EstimatedPriceQueryRequest) Validate() error {
  return dara.Validate(s)
}

