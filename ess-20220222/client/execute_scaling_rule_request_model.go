// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteScalingRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBreachThreshold(v float32) *ExecuteScalingRuleRequest
  GetBreachThreshold() *float32 
  SetClientToken(v string) *ExecuteScalingRuleRequest
  GetClientToken() *string 
  SetMetricValue(v float32) *ExecuteScalingRuleRequest
  GetMetricValue() *float32 
  SetOwnerAccount(v string) *ExecuteScalingRuleRequest
  GetOwnerAccount() *string 
  SetOwnerId(v int64) *ExecuteScalingRuleRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *ExecuteScalingRuleRequest
  GetRegionId() *string 
  SetResourceOwnerAccount(v string) *ExecuteScalingRuleRequest
  GetResourceOwnerAccount() *string 
  SetResourceOwnerId(v int64) *ExecuteScalingRuleRequest
  GetResourceOwnerId() *int64 
  SetScalingRuleAri(v string) *ExecuteScalingRuleRequest
  GetScalingRuleAri() *string 
}

type ExecuteScalingRuleRequest struct {
  // The threshold specified when the step scaling rule is executed. Valid values: -9.999999E18 to 9.999999E18.
  // 
  // example:
  // 
  // 1.0
  BreachThreshold *float32 `json:"BreachThreshold,omitempty" xml:"BreachThreshold,omitempty"`
  // The client token that is used to ensure the idempotence of the request.
  // 
  // You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25965.html).
  // 
  // example:
  // 
  // 123e4567-e89b-12d3-a456-426655440000
  ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
  // The metric value specified when the step scaling rule is executed. Valid values: -9.999999E18 to 9.999999E18.
  // 
  // example:
  // 
  // 1.0
  MetricValue *float32 `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
  OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID of the scaling group.
  // 
  // example:
  // 
  // cn-qingdao
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
  ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
  // The unique identifier of the scaling rule.
  // 
  // >  You can call this operation to execute simple scaling rules and step scaling rules. If you want to call this operation to execute a step scaling rule, you must specify `BreachThreshold` and `MetricValue`.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ari:acs:ess:cn-hangzhou:140692647406****:scalingrule/asr-bp1dvirgwkoowxk7****
  ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
}

func (s ExecuteScalingRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteScalingRuleRequest) GoString() string {
  return s.String()
}

func (s *ExecuteScalingRuleRequest) GetBreachThreshold() *float32  {
  return s.BreachThreshold
}

func (s *ExecuteScalingRuleRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteScalingRuleRequest) GetMetricValue() *float32  {
  return s.MetricValue
}

func (s *ExecuteScalingRuleRequest) GetOwnerAccount() *string  {
  return s.OwnerAccount
}

func (s *ExecuteScalingRuleRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *ExecuteScalingRuleRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *ExecuteScalingRuleRequest) GetResourceOwnerAccount() *string  {
  return s.ResourceOwnerAccount
}

func (s *ExecuteScalingRuleRequest) GetResourceOwnerId() *int64  {
  return s.ResourceOwnerId
}

func (s *ExecuteScalingRuleRequest) GetScalingRuleAri() *string  {
  return s.ScalingRuleAri
}

func (s *ExecuteScalingRuleRequest) SetBreachThreshold(v float32) *ExecuteScalingRuleRequest {
  s.BreachThreshold = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetClientToken(v string) *ExecuteScalingRuleRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetMetricValue(v float32) *ExecuteScalingRuleRequest {
  s.MetricValue = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerAccount(v string) *ExecuteScalingRuleRequest {
  s.OwnerAccount = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerId(v int64) *ExecuteScalingRuleRequest {
  s.OwnerId = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetRegionId(v string) *ExecuteScalingRuleRequest {
  s.RegionId = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerAccount(v string) *ExecuteScalingRuleRequest {
  s.ResourceOwnerAccount = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerId(v int64) *ExecuteScalingRuleRequest {
  s.ResourceOwnerId = &v
  return s
}

func (s *ExecuteScalingRuleRequest) SetScalingRuleAri(v string) *ExecuteScalingRuleRequest {
  s.ScalingRuleAri = &v
  return s
}

func (s *ExecuteScalingRuleRequest) Validate() error {
  return dara.Validate(s)
}

