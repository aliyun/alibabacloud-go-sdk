// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableNotifyPolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableNotifyPolicyResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableNotifyPolicyResponseBody
  GetSuccess() *bool 
  SetUuid(v string) *EnableNotifyPolicyResponseBody
  GetUuid() *string 
}

type EnableNotifyPolicyResponseBody struct {
  // Id of the request
  // 
  // example:
  // 
  // 0CEC5375-C554-562B-A65F-***
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
  // Indicates whether the request was successful.
  Success *bool `json:"success,omitempty" xml:"success,omitempty"`
  // The unique identifier of the notification policy.
  // 
  // example:
  // 
  // 7076c75c-c804-461e-975f-c6f9ed5af745
  Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s EnableNotifyPolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableNotifyPolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableNotifyPolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableNotifyPolicyResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableNotifyPolicyResponseBody) GetUuid() *string  {
  return s.Uuid
}

func (s *EnableNotifyPolicyResponseBody) SetRequestId(v string) *EnableNotifyPolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableNotifyPolicyResponseBody) SetSuccess(v bool) *EnableNotifyPolicyResponseBody {
  s.Success = &v
  return s
}

func (s *EnableNotifyPolicyResponseBody) SetUuid(v string) *EnableNotifyPolicyResponseBody {
  s.Uuid = &v
  return s
}

func (s *EnableNotifyPolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

