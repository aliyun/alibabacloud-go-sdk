// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformApplyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteTerraformApplyResponseBody
  GetRequestId() *string 
  SetStateId(v string) *ExecuteTerraformApplyResponseBody
  GetStateId() *string 
}

type ExecuteTerraformApplyResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // C896FE0A-1BEA-5D01-BFF4-B03B82B9CA3D
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // task-xxxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteTerraformApplyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformApplyResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformApplyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTerraformApplyResponseBody) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformApplyResponseBody) SetRequestId(v string) *ExecuteTerraformApplyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTerraformApplyResponseBody) SetStateId(v string) *ExecuteTerraformApplyResponseBody {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformApplyResponseBody) Validate() error {
  return dara.Validate(s)
}

