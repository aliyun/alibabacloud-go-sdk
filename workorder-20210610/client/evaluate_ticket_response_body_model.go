// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEvaluateTicketResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccessDeniedDetail(v string) *EvaluateTicketResponseBody
  GetAccessDeniedDetail() *string 
  SetCode(v int32) *EvaluateTicketResponseBody
  GetCode() *int32 
  SetMessage(v string) *EvaluateTicketResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EvaluateTicketResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EvaluateTicketResponseBody
  GetSuccess() *bool 
}

type EvaluateTicketResponseBody struct {
  AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
  // The response code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message. If success is set to false, the message is returned.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // C1DA4C6F-963E-5741-AB57-67A554D102FD
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values: The value true indicates a success. The value false indicates a failure.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EvaluateTicketResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EvaluateTicketResponseBody) GoString() string {
  return s.String()
}

func (s *EvaluateTicketResponseBody) GetAccessDeniedDetail() *string  {
  return s.AccessDeniedDetail
}

func (s *EvaluateTicketResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EvaluateTicketResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EvaluateTicketResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EvaluateTicketResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EvaluateTicketResponseBody) SetAccessDeniedDetail(v string) *EvaluateTicketResponseBody {
  s.AccessDeniedDetail = &v
  return s
}

func (s *EvaluateTicketResponseBody) SetCode(v int32) *EvaluateTicketResponseBody {
  s.Code = &v
  return s
}

func (s *EvaluateTicketResponseBody) SetMessage(v string) *EvaluateTicketResponseBody {
  s.Message = &v
  return s
}

func (s *EvaluateTicketResponseBody) SetRequestId(v string) *EvaluateTicketResponseBody {
  s.RequestId = &v
  return s
}

func (s *EvaluateTicketResponseBody) SetSuccess(v bool) *EvaluateTicketResponseBody {
  s.Success = &v
  return s
}

func (s *EvaluateTicketResponseBody) Validate() error {
  return dara.Validate(s)
}

