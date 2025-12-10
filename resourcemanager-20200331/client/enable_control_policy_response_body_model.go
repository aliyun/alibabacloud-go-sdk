// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableControlPolicyResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetEnablementStatus(v string) *EnableControlPolicyResponseBody
  GetEnablementStatus() *string 
  SetRequestId(v string) *EnableControlPolicyResponseBody
  GetRequestId() *string 
}

type EnableControlPolicyResponseBody struct {
  // The status of the Control Policy feature. Valid values:
  // 
  // 	- Enabled: The Control Policy feature is enabled.
  // 
  // 	- PendingEnable: The Control Policy feature is being enabled.
  // 
  // 	- Disabled: The Control Policy feature is disabled.
  // 
  // 	- PendingDisable: The Control Policy feature is being disabled.
  // 
  // example:
  // 
  // PendingEnable
  EnablementStatus *string `json:"EnablementStatus,omitempty" xml:"EnablementStatus,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // 8CE7BD95-EFFA-4911-A1E0-BD4412697FEB
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableControlPolicyResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableControlPolicyResponseBody) GoString() string {
  return s.String()
}

func (s *EnableControlPolicyResponseBody) GetEnablementStatus() *string  {
  return s.EnablementStatus
}

func (s *EnableControlPolicyResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableControlPolicyResponseBody) SetEnablementStatus(v string) *EnableControlPolicyResponseBody {
  s.EnablementStatus = &v
  return s
}

func (s *EnableControlPolicyResponseBody) SetRequestId(v string) *EnableControlPolicyResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableControlPolicyResponseBody) Validate() error {
  return dara.Validate(s)
}

