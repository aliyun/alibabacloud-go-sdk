// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSchemaPropertyRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableSchemaPropertyRequest
  GetInstanceId() *string 
  SetPropertyName(v string) *EnableSchemaPropertyRequest
  GetPropertyName() *string 
  SetRequestId(v string) *EnableSchemaPropertyRequest
  GetRequestId() *string 
  SetSchemaId(v string) *EnableSchemaPropertyRequest
  GetSchemaId() *string 
}

type EnableSchemaPropertyRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // name
  PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
  // example:
  // 
  // 03C67DAD-EB26-41D8-949D-9B0C470FB716
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // schema id
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // profile
  SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s EnableSchemaPropertyRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSchemaPropertyRequest) GoString() string {
  return s.String()
}

func (s *EnableSchemaPropertyRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableSchemaPropertyRequest) GetPropertyName() *string  {
  return s.PropertyName
}

func (s *EnableSchemaPropertyRequest) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableSchemaPropertyRequest) GetSchemaId() *string  {
  return s.SchemaId
}

func (s *EnableSchemaPropertyRequest) SetInstanceId(v string) *EnableSchemaPropertyRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableSchemaPropertyRequest) SetPropertyName(v string) *EnableSchemaPropertyRequest {
  s.PropertyName = &v
  return s
}

func (s *EnableSchemaPropertyRequest) SetRequestId(v string) *EnableSchemaPropertyRequest {
  s.RequestId = &v
  return s
}

func (s *EnableSchemaPropertyRequest) SetSchemaId(v string) *EnableSchemaPropertyRequest {
  s.SchemaId = &v
  return s
}

func (s *EnableSchemaPropertyRequest) Validate() error {
  return dara.Validate(s)
}

