// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformDestroyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteTerraformDestroyResponseBody
  GetRequestId() *string 
  SetStateId(v string) *ExecuteTerraformDestroyResponseBody
  GetStateId() *string 
}

type ExecuteTerraformDestroyResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // BF72A6FB-B071-5F2E-A036-9D62545B962C
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // task-xxxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteTerraformDestroyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformDestroyResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformDestroyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteTerraformDestroyResponseBody) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformDestroyResponseBody) SetRequestId(v string) *ExecuteTerraformDestroyResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteTerraformDestroyResponseBody) SetStateId(v string) *ExecuteTerraformDestroyResponseBody {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformDestroyResponseBody) Validate() error {
  return dara.Validate(s)
}

