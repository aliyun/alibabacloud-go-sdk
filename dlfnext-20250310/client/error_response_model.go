// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iErrorResponse interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *ErrorResponse
  GetCode() *int32 
  SetMessage(v string) *ErrorResponse
  GetMessage() *string 
  SetResourceName(v string) *ErrorResponse
  GetResourceName() *string 
  SetResourceType(v string) *ErrorResponse
  GetResourceType() *string 
}

type ErrorResponse struct {
  // The error code.
  // 
  // example:
  // 
  // 404
  Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // Resource not found
  Message *string `json:"message,omitempty" xml:"message,omitempty"`
  // The name of the resource involved in the error.
  // 
  // example:
  // 
  // my_table
  ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
  // The type of the resource involved in the error.
  // 
  // example:
  // 
  // table
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ErrorResponse) String() string {
  return dara.Prettify(s)
}

func (s ErrorResponse) GoString() string {
  return s.String()
}

func (s *ErrorResponse) GetCode() *int32  {
  return s.Code
}

func (s *ErrorResponse) GetMessage() *string  {
  return s.Message
}

func (s *ErrorResponse) GetResourceName() *string  {
  return s.ResourceName
}

func (s *ErrorResponse) GetResourceType() *string  {
  return s.ResourceType
}

func (s *ErrorResponse) SetCode(v int32) *ErrorResponse {
  s.Code = &v
  return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
  s.Message = &v
  return s
}

func (s *ErrorResponse) SetResourceName(v string) *ErrorResponse {
  s.ResourceName = &v
  return s
}

func (s *ErrorResponse) SetResourceType(v string) *ErrorResponse {
  s.ResourceType = &v
  return s
}

func (s *ErrorResponse) Validate() error {
  return dara.Validate(s)
}

