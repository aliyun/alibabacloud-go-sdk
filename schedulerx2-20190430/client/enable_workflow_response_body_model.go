// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableWorkflowResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v int32) *EnableWorkflowResponseBody
  GetCode() *int32 
  SetMessage(v string) *EnableWorkflowResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableWorkflowResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableWorkflowResponseBody
  GetSuccess() *bool 
}

type EnableWorkflowResponseBody struct {
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
  // Your request is denied as lack of ssl protect.
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 4F68ABED-AC31-4412-9297-D9A8F0401108
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the workflow was enabled. Valid values:
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

func (s EnableWorkflowResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableWorkflowResponseBody) GoString() string {
  return s.String()
}

func (s *EnableWorkflowResponseBody) GetCode() *int32  {
  return s.Code
}

func (s *EnableWorkflowResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableWorkflowResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableWorkflowResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableWorkflowResponseBody) SetCode(v int32) *EnableWorkflowResponseBody {
  s.Code = &v
  return s
}

func (s *EnableWorkflowResponseBody) SetMessage(v string) *EnableWorkflowResponseBody {
  s.Message = &v
  return s
}

func (s *EnableWorkflowResponseBody) SetRequestId(v string) *EnableWorkflowResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableWorkflowResponseBody) SetSuccess(v bool) *EnableWorkflowResponseBody {
  s.Success = &v
  return s
}

func (s *EnableWorkflowResponseBody) Validate() error {
  return dara.Validate(s)
}

