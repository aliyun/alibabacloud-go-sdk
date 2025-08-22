// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteRegistryModuleResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *ExecuteRegistryModuleResponseBody
  GetRequestId() *string 
  SetStateId(v string) *ExecuteRegistryModuleResponseBody
  GetStateId() *string 
}

type ExecuteRegistryModuleResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 79284133-D4BA-56B3-954C-D538256F7EAA
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // example:
  // 
  // task-xxxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteRegistryModuleResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExecuteRegistryModuleResponseBody) GoString() string {
  return s.String()
}

func (s *ExecuteRegistryModuleResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExecuteRegistryModuleResponseBody) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteRegistryModuleResponseBody) SetRequestId(v string) *ExecuteRegistryModuleResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExecuteRegistryModuleResponseBody) SetStateId(v string) *ExecuteRegistryModuleResponseBody {
  s.StateId = &v
  return s
}

func (s *ExecuteRegistryModuleResponseBody) Validate() error {
  return dara.Validate(s)
}

