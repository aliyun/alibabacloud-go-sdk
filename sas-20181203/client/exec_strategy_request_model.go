// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecStrategyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetExecAction(v string) *ExecStrategyRequest
  GetExecAction() *string 
  SetLang(v string) *ExecStrategyRequest
  GetLang() *string 
  SetStrategyId(v int32) *ExecStrategyRequest
  GetStrategyId() *int32 
}

type ExecStrategyRequest struct {
  ExecAction *string `json:"ExecAction,omitempty" xml:"ExecAction,omitempty"`
  // The language of the content within the request and response. Default value: **zh**. Valid values:
  // 
  // 	- **zh**: Chinese
  // 
  // 	- **en**: English
  // 
  // example:
  // 
  // zh
  Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
  // The ID of the baseline check policy.
  // 
  // >  You can call the [DescribeStrategy](~~DescribeStrategy~~) operation to query the IDs of baseline check policies.
  // 
  // example:
  // 
  // 215421
  StrategyId *int32 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s ExecStrategyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecStrategyRequest) GoString() string {
  return s.String()
}

func (s *ExecStrategyRequest) GetExecAction() *string  {
  return s.ExecAction
}

func (s *ExecStrategyRequest) GetLang() *string  {
  return s.Lang
}

func (s *ExecStrategyRequest) GetStrategyId() *int32  {
  return s.StrategyId
}

func (s *ExecStrategyRequest) SetExecAction(v string) *ExecStrategyRequest {
  s.ExecAction = &v
  return s
}

func (s *ExecStrategyRequest) SetLang(v string) *ExecStrategyRequest {
  s.Lang = &v
  return s
}

func (s *ExecStrategyRequest) SetStrategyId(v int32) *ExecStrategyRequest {
  s.StrategyId = &v
  return s
}

func (s *ExecStrategyRequest) Validate() error {
  return dara.Validate(s)
}

