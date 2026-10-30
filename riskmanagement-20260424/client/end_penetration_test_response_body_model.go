// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEndPenetrationTestResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EndPenetrationTestResponseBody
  GetCode() *string 
  SetMessage(v string) *EndPenetrationTestResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EndPenetrationTestResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EndPenetrationTestResponseBody
  GetSuccess() *bool 
}

type EndPenetrationTestResponseBody struct {
  // The status code.
  // 
  // > 200: Success. Other values (such as 500 and 400): Error codes.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The message.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 6B48613E-86DE-5411-BDBE-429C80B45F3F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the call was successful. Valid values:
  // 
  // - **true**: The call was successful.                                
  // 
  // - **false**: The call failed.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EndPenetrationTestResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EndPenetrationTestResponseBody) GoString() string {
  return s.String()
}

func (s *EndPenetrationTestResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EndPenetrationTestResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EndPenetrationTestResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EndPenetrationTestResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EndPenetrationTestResponseBody) SetCode(v string) *EndPenetrationTestResponseBody {
  s.Code = &v
  return s
}

func (s *EndPenetrationTestResponseBody) SetMessage(v string) *EndPenetrationTestResponseBody {
  s.Message = &v
  return s
}

func (s *EndPenetrationTestResponseBody) SetRequestId(v string) *EndPenetrationTestResponseBody {
  s.RequestId = &v
  return s
}

func (s *EndPenetrationTestResponseBody) SetSuccess(v bool) *EndPenetrationTestResponseBody {
  s.Success = &v
  return s
}

func (s *EndPenetrationTestResponseBody) Validate() error {
  return dara.Validate(s)
}

