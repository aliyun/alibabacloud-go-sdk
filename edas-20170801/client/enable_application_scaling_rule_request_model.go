// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableApplicationScalingRuleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *EnableApplicationScalingRuleRequest
  GetAppId() *string 
  SetScalingRuleName(v string) *EnableApplicationScalingRuleRequest
  GetScalingRuleName() *string 
}

type EnableApplicationScalingRuleRequest struct {
  // The ID of the application. You can call the ListApplication operation to query the application ID. For more information, see [ListApplication](https://help.aliyun.com/document_detail/149390.html).
  // 
  // example:
  // 
  // 78194c76-3dca-418e-a263-cccd1ab4****
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  // The name of the auto scaling policy.
  // 
  // example:
  // 
  // cpu-trigger
  ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s EnableApplicationScalingRuleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableApplicationScalingRuleRequest) GoString() string {
  return s.String()
}

func (s *EnableApplicationScalingRuleRequest) GetAppId() *string  {
  return s.AppId
}

func (s *EnableApplicationScalingRuleRequest) GetScalingRuleName() *string  {
  return s.ScalingRuleName
}

func (s *EnableApplicationScalingRuleRequest) SetAppId(v string) *EnableApplicationScalingRuleRequest {
  s.AppId = &v
  return s
}

func (s *EnableApplicationScalingRuleRequest) SetScalingRuleName(v string) *EnableApplicationScalingRuleRequest {
  s.ScalingRuleName = &v
  return s
}

func (s *EnableApplicationScalingRuleRequest) Validate() error {
  return dara.Validate(s)
}

