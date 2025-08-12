// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableEventRulesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableEventRulesResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableEventRulesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableEventRulesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableEventRulesResponseBody
  GetSuccess() *bool 
}

type EnableEventRulesResponseBody struct {
  // The HTTP status code.
  // 
  // >  The status code 200 indicates that the call was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // success
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 20F2896A-6684-4A04-8255-4155B1593C70
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the call was successful. The value true indicates a success. The value false indicates a failure.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableEventRulesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableEventRulesResponseBody) GoString() string {
  return s.String()
}

func (s *EnableEventRulesResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableEventRulesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableEventRulesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableEventRulesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableEventRulesResponseBody) SetCode(v string) *EnableEventRulesResponseBody {
  s.Code = &v
  return s
}

func (s *EnableEventRulesResponseBody) SetMessage(v string) *EnableEventRulesResponseBody {
  s.Message = &v
  return s
}

func (s *EnableEventRulesResponseBody) SetRequestId(v string) *EnableEventRulesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableEventRulesResponseBody) SetSuccess(v bool) *EnableEventRulesResponseBody {
  s.Success = &v
  return s
}

func (s *EnableEventRulesResponseBody) Validate() error {
  return dara.Validate(s)
}

