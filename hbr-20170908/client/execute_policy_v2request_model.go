// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecutePolicyV2Request interface {
  dara.Model
  String() string
  GoString() string
  SetDataSourceId(v string) *ExecutePolicyV2Request
  GetDataSourceId() *string 
  SetPolicyId(v string) *ExecutePolicyV2Request
  GetPolicyId() *string 
  SetRuleId(v string) *ExecutePolicyV2Request
  GetRuleId() *string 
  SetSourceType(v string) *ExecutePolicyV2Request
  GetSourceType() *string 
}

type ExecutePolicyV2Request struct {
  // Data source ID.
  // 
  // example:
  // 
  // i-bp1************dtv
  DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
  // Policy ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // po-000************hky
  PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
  // Rule ID, limited to executing rules of **RuleType*	- **BACKUP**.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // rule-0002*****ux8
  RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
  // Data source type, with the value range as follows:
  // 
  // - **UDM_ECS**: Represents ECS full machine backup.
  // 
  // example:
  // 
  // UDM_ECS
  SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ExecutePolicyV2Request) String() string {
  return dara.Prettify(s)
}

func (s ExecutePolicyV2Request) GoString() string {
  return s.String()
}

func (s *ExecutePolicyV2Request) GetDataSourceId() *string  {
  return s.DataSourceId
}

func (s *ExecutePolicyV2Request) GetPolicyId() *string  {
  return s.PolicyId
}

func (s *ExecutePolicyV2Request) GetRuleId() *string  {
  return s.RuleId
}

func (s *ExecutePolicyV2Request) GetSourceType() *string  {
  return s.SourceType
}

func (s *ExecutePolicyV2Request) SetDataSourceId(v string) *ExecutePolicyV2Request {
  s.DataSourceId = &v
  return s
}

func (s *ExecutePolicyV2Request) SetPolicyId(v string) *ExecutePolicyV2Request {
  s.PolicyId = &v
  return s
}

func (s *ExecutePolicyV2Request) SetRuleId(v string) *ExecutePolicyV2Request {
  s.RuleId = &v
  return s
}

func (s *ExecutePolicyV2Request) SetSourceType(v string) *ExecutePolicyV2Request {
  s.SourceType = &v
  return s
}

func (s *ExecutePolicyV2Request) Validate() error {
  return dara.Validate(s)
}

