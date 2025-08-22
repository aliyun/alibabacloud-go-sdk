// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteTerraformPlanResponseBody
  GetRequestId() *string 
  SetStateId(v string) *ExecuteTerraformPlanResponseBody
  GetStateId() *string 
}

type ExecuteTerraformPlanResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 0D298375-F92F-5B65-82E4-EA68F02521F1
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // task-xxxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteTerraformPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformPlanResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTerraformPlanResponseBody) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformPlanResponseBody) SetRequestId(v string) *ExecuteTerraformPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTerraformPlanResponseBody) SetStateId(v string) *ExecuteTerraformPlanResponseBody {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

