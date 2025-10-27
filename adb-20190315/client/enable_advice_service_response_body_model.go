// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAdviceServiceResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetMessage(v string) *EnableAdviceServiceResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableAdviceServiceResponseBody
  GetRequestId() *string 
}

type EnableAdviceServiceResponseBody struct {
  // The message returned for the operation. Valid values:
  // 
  // 	- **Success*	- is returned if the operation is successful.
  // 
  // 	- An error message is returned if the operation fails.
  // 
  // example:
  // 
  // Success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // E1745C03-7CCE-55CF-932E-08121AAFA6AF
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAdviceServiceResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAdviceServiceResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAdviceServiceResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableAdviceServiceResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAdviceServiceResponseBody) SetMessage(v string) *EnableAdviceServiceResponseBody {
  s.Message = &v
  return s
}

func (s *EnableAdviceServiceResponseBody) SetRequestId(v string) *EnableAdviceServiceResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAdviceServiceResponseBody) Validate() error {
  return dara.Validate(s)
}

