// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElasticSpotSpec interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceType(v string) *ElasticSpotSpec
  GetInstanceType() *string 
  SetSpotDiscountLimit(v float64) *ElasticSpotSpec
  GetSpotDiscountLimit() *float64 
  SetSpotPriceLimit(v float64) *ElasticSpotSpec
  GetSpotPriceLimit() *float64 
  SetSpotStrategy(v string) *ElasticSpotSpec
  GetSpotStrategy() *string 
}

type ElasticSpotSpec struct {
  InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
  SpotDiscountLimit *float64 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
  SpotPriceLimit *float64 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
  SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s ElasticSpotSpec) String() string {
  return dara.Prettify(s)
}

func (s ElasticSpotSpec) GoString() string {
  return s.String()
}

func (s *ElasticSpotSpec) GetInstanceType() *string  {
  return s.InstanceType
}

func (s *ElasticSpotSpec) GetSpotDiscountLimit() *float64  {
  return s.SpotDiscountLimit
}

func (s *ElasticSpotSpec) GetSpotPriceLimit() *float64  {
  return s.SpotPriceLimit
}

func (s *ElasticSpotSpec) GetSpotStrategy() *string  {
  return s.SpotStrategy
}

func (s *ElasticSpotSpec) SetInstanceType(v string) *ElasticSpotSpec {
  s.InstanceType = &v
  return s
}

func (s *ElasticSpotSpec) SetSpotDiscountLimit(v float64) *ElasticSpotSpec {
  s.SpotDiscountLimit = &v
  return s
}

func (s *ElasticSpotSpec) SetSpotPriceLimit(v float64) *ElasticSpotSpec {
  s.SpotPriceLimit = &v
  return s
}

func (s *ElasticSpotSpec) SetSpotStrategy(v string) *ElasticSpotSpec {
  s.SpotStrategy = &v
  return s
}

func (s *ElasticSpotSpec) Validate() error {
  return dara.Validate(s)
}

