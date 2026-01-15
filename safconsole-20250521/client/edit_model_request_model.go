// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditModelRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCustomerModuleId(v int32) *EditModelRequest
  GetCustomerModuleId() *int32 
}

type EditModelRequest struct {
  // example:
  // 
  // 456
  CustomerModuleId *int32 `json:"CustomerModuleId,omitempty" xml:"CustomerModuleId,omitempty"`
}

func (s EditModelRequest) String() string {
  return dara.Prettify(s)
}

func (s EditModelRequest) GoString() string {
  return s.String()
}

func (s *EditModelRequest) GetCustomerModuleId() *int32  {
  return s.CustomerModuleId
}

func (s *EditModelRequest) SetCustomerModuleId(v int32) *EditModelRequest {
  s.CustomerModuleId = &v
  return s
}

func (s *EditModelRequest) Validate() error {
  return dara.Validate(s)
}

