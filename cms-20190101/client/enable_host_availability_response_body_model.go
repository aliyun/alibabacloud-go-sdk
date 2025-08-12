// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHostAvailabilityResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableHostAvailabilityResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableHostAvailabilityResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableHostAvailabilityResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableHostAvailabilityResponseBody
  GetSuccess() *bool 
}

type EnableHostAvailabilityResponseBody struct {
  // The status code.
  // 
  // >  The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // User not authorized to operate on the specified resource.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // ACBDBB40-DFB6-4F4C-8957-51FFB233969C
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- true
  // 
  // 	- false
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHostAvailabilityResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHostAvailabilityResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHostAvailabilityResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableHostAvailabilityResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableHostAvailabilityResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHostAvailabilityResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableHostAvailabilityResponseBody) SetCode(v string) *EnableHostAvailabilityResponseBody {
  s.Code = &v
  return s
}

func (s *EnableHostAvailabilityResponseBody) SetMessage(v string) *EnableHostAvailabilityResponseBody {
  s.Message = &v
  return s
}

func (s *EnableHostAvailabilityResponseBody) SetRequestId(v string) *EnableHostAvailabilityResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHostAvailabilityResponseBody) SetSuccess(v bool) *EnableHostAvailabilityResponseBody {
  s.Success = &v
  return s
}

func (s *EnableHostAvailabilityResponseBody) Validate() error {
  return dara.Validate(s)
}

