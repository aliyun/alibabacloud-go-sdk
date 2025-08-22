// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteTerraformDestroyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *ExecuteTerraformDestroyRequest
  GetClientToken() *string 
  SetStateId(v string) *ExecuteTerraformDestroyRequest
  GetStateId() *string 
}

type ExecuteTerraformDestroyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // a65451293e64979ba7a4b573950217fe
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // task-xxx
  StateId *string `json:"stateId,omitempty" xml:"stateId,omitempty"`
}

func (s ExecuteTerraformDestroyRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteTerraformDestroyRequest) GoString() string {
  return s.String()
}

func (s *ExecuteTerraformDestroyRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteTerraformDestroyRequest) GetStateId() *string  {
  return s.StateId
}

func (s *ExecuteTerraformDestroyRequest) SetClientToken(v string) *ExecuteTerraformDestroyRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteTerraformDestroyRequest) SetStateId(v string) *ExecuteTerraformDestroyRequest {
  s.StateId = &v
  return s
}

func (s *ExecuteTerraformDestroyRequest) Validate() error {
  return dara.Validate(s)
}

