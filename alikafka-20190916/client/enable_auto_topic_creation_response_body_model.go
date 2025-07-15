// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableAutoTopicCreationResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableAutoTopicCreationResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnableAutoTopicCreationResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableAutoTopicCreationResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableAutoTopicCreationResponseBody
  GetSuccess() *bool 
}

type EnableAutoTopicCreationResponseBody struct {
  // The returned status code. If the request is successful, 200 is returned.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message.
  // 
  // example:
  // 
  // operation success.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 9E3B3592-5994-4F65-A61E-E62A77A7***
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAutoTopicCreationResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableAutoTopicCreationResponseBody) GoString() string {
  return s.String()
}

func (s *EnableAutoTopicCreationResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableAutoTopicCreationResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableAutoTopicCreationResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableAutoTopicCreationResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableAutoTopicCreationResponseBody) SetCode(v int32) *EnableAutoTopicCreationResponseBody {
  s.Code = &v
  return s
}

func (s *EnableAutoTopicCreationResponseBody) SetMessage(v string) *EnableAutoTopicCreationResponseBody {
  s.Message = &v
  return s
}

func (s *EnableAutoTopicCreationResponseBody) SetRequestId(v string) *EnableAutoTopicCreationResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableAutoTopicCreationResponseBody) SetSuccess(v bool) *EnableAutoTopicCreationResponseBody {
  s.Success = &v
  return s
}

func (s *EnableAutoTopicCreationResponseBody) Validate() error {
  return dara.Validate(s)
}

