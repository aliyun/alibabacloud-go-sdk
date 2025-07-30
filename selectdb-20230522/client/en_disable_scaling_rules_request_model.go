// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnDisableScalingRulesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *EnDisableScalingRulesRequest
  GetClusterId() *string 
  SetDbInstanceId(v string) *EnDisableScalingRulesRequest
  GetDbInstanceId() *string 
  SetProduct(v string) *EnDisableScalingRulesRequest
  GetProduct() *string 
  SetRegionId(v string) *EnDisableScalingRulesRequest
  GetRegionId() *string 
  SetResourceOwnerId(v int64) *EnDisableScalingRulesRequest
  GetResourceOwnerId() *int64 
  SetScalingRulesEnable(v bool) *EnDisableScalingRulesRequest
  GetScalingRulesEnable() *bool 
}

type EnDisableScalingRulesRequest struct {
  // The cluster ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // selectdb-cn-nwy3jv1oa02-be
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // selectdb-cn-7213cjv****
  DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
  // The cloud service.
  // 
  // example:
  // 
  // selectdb
  Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
  // The region ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // Specifies whether to enable the scheduled scaling policy.
  // 
  // Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // true
  ScalingRulesEnable *bool `json:"ScalingRulesEnable,omitempty" xml:"ScalingRulesEnable,omitempty"`
}

func (s EnDisableScalingRulesRequest) String() string {
  return dara.Prettify(s)
}

func (s EnDisableScalingRulesRequest) GoString() string {
  return s.String()
}

func (s *EnDisableScalingRulesRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnDisableScalingRulesRequest) GetDbInstanceId() *string  {
  return s.DbInstanceId
}

func (s *EnDisableScalingRulesRequest) GetProduct() *string  {
  return s.Product
}

func (s *EnDisableScalingRulesRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnDisableScalingRulesRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *EnDisableScalingRulesRequest) GetScalingRulesEnable() *bool  {
  return s.ScalingRulesEnable
}

func (s *EnDisableScalingRulesRequest) SetClusterId(v string) *EnDisableScalingRulesRequest {
  s.ClusterId = &v
  return s
}

func (s *EnDisableScalingRulesRequest) SetDbInstanceId(v string) *EnDisableScalingRulesRequest {
  s.DbInstanceId = &v
  return s
}

func (s *EnDisableScalingRulesRequest) SetProduct(v string) *EnDisableScalingRulesRequest {
  s.Product = &v
  return s
}

func (s *EnDisableScalingRulesRequest) SetRegionId(v string) *EnDisableScalingRulesRequest {
  s.RegionId = &v
  return s
}

func (s *EnDisableScalingRulesRequest) SetResourceOwnerId(v int64) *EnDisableScalingRulesRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *EnDisableScalingRulesRequest) SetScalingRulesEnable(v bool) *EnDisableScalingRulesRequest {
  s.ScalingRulesEnable = &v
  return s
}

func (s *EnDisableScalingRulesRequest) Validate() error {
  return dara.Validate(s)
}

