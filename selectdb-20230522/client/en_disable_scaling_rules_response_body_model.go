// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnDisableScalingRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v *EnDisableScalingRulesResponseBodyData) *EnDisableScalingRulesResponseBody
  GetData() *EnDisableScalingRulesResponseBodyData 
  SetRequestId(v string) *EnDisableScalingRulesResponseBody
  GetRequestId() *string 
}

type EnDisableScalingRulesResponseBody struct {
  // The data returned.
  Data *EnDisableScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // The request ID.
  // 
  // example:
  // 
  // 4773E4EC-025D-509F-AEA9-D53123FDFB0F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnDisableScalingRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnDisableScalingRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EnDisableScalingRulesResponseBody) GetData() *EnDisableScalingRulesResponseBodyData  {
  return s.Data
}

func (s *EnDisableScalingRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnDisableScalingRulesResponseBody) SetData(v *EnDisableScalingRulesResponseBodyData) *EnDisableScalingRulesResponseBody {
  s.Data = v
  return s
}

func (s *EnDisableScalingRulesResponseBody) SetRequestId(v string) *EnDisableScalingRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnDisableScalingRulesResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnDisableScalingRulesResponseBodyData struct {
  // The cluster ID.
  // 
  // example:
  // 
  // selectdb-cn-pe33jc1nd01-be
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // The instance ID.
  // 
  // example:
  // 
  // selectdb-cn-7213cjv****
  DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
  // Indicates whether the scheduled scaling policy is enabled.
  // 
  // Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  ScalingRulesEnable *bool `json:"ScalingRulesEnable,omitempty" xml:"ScalingRulesEnable,omitempty"`
}

func (s EnDisableScalingRulesResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnDisableScalingRulesResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnDisableScalingRulesResponseBodyData) GetClusterId() *string  {
  return s.ClusterId
}

func (s *EnDisableScalingRulesResponseBodyData) GetDbInstanceId() *string  {
  return s.DbInstanceId
}

func (s *EnDisableScalingRulesResponseBodyData) GetScalingRulesEnable() *bool  {
  return s.ScalingRulesEnable
}

func (s *EnDisableScalingRulesResponseBodyData) SetClusterId(v string) *EnDisableScalingRulesResponseBodyData {
  s.ClusterId = &v
  return s
}

func (s *EnDisableScalingRulesResponseBodyData) SetDbInstanceId(v string) *EnDisableScalingRulesResponseBodyData {
  s.DbInstanceId = &v
  return s
}

func (s *EnDisableScalingRulesResponseBodyData) SetScalingRulesEnable(v bool) *EnDisableScalingRulesResponseBodyData {
  s.ScalingRulesEnable = &v
  return s
}

func (s *EnDisableScalingRulesResponseBodyData) Validate() error {
  return dara.Validate(s)
}

