// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableJobResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableJobResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnableJobResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableJobResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableJobResponseBody
  GetSuccess() *bool 
}

type EnableJobResponseBody struct {
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
  // The error message that is returned only if the corresponding error occurs.
  // 
  // example:
  // 
  // jobid: 92583 not match groupId: testSchedulerx.defaultGroup
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request was successful. Valid values:
  // 
  // 	- **true**
  // 
  // 	- **false**
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableJobResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableJobResponseBody) GoString() string {
  return s.String()
}

func (s *EnableJobResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableJobResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableJobResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableJobResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableJobResponseBody) SetCode(v int32) *EnableJobResponseBody {
  s.Code = &v
  return s
}

func (s *EnableJobResponseBody) SetMessage(v string) *EnableJobResponseBody {
  s.Message = &v
  return s
}

func (s *EnableJobResponseBody) SetRequestId(v string) *EnableJobResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableJobResponseBody) SetSuccess(v bool) *EnableJobResponseBody {
  s.Success = &v
  return s
}

func (s *EnableJobResponseBody) Validate() error {
  return dara.Validate(s)
}

