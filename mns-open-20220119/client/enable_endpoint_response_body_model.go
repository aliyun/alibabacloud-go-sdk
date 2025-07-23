// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEndpointResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int64) *EnableEndpointResponseBody
  GetCode() *int64 
  SetMessage(v string) *EnableEndpointResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableEndpointResponseBody
  GetRequestId() *string 
  SetStatus(v string) *EnableEndpointResponseBody
  GetStatus() *string 
  SetSuccess(v bool) *EnableEndpointResponseBody
  GetSuccess() *bool 
}

type EnableEndpointResponseBody struct {
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // operation success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 06273500-249F-5863-121D-74D51123****
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The response status.
  // 
  // example:
  // 
  // Success
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  // Indicates whether the request was successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableEndpointResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableEndpointResponseBody) GoString() string {
  return s.String()
}

func (s *EnableEndpointResponseBody) GetCode() *int64  {
  return s.Code
}

func (s *EnableEndpointResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableEndpointResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableEndpointResponseBody) GetStatus() *string  {
  return s.Status
}

func (s *EnableEndpointResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableEndpointResponseBody) SetCode(v int64) *EnableEndpointResponseBody {
  s.Code = &v
  return s
}

func (s *EnableEndpointResponseBody) SetMessage(v string) *EnableEndpointResponseBody {
  s.Message = &v
  return s
}

func (s *EnableEndpointResponseBody) SetRequestId(v string) *EnableEndpointResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableEndpointResponseBody) SetStatus(v string) *EnableEndpointResponseBody {
  s.Status = &v
  return s
}

func (s *EnableEndpointResponseBody) SetSuccess(v bool) *EnableEndpointResponseBody {
  s.Success = &v
  return s
}

func (s *EnableEndpointResponseBody) Validate() error {
  return dara.Validate(s)
}

