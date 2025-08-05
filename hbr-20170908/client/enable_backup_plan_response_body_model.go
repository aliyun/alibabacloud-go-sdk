// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableBackupPlanResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableBackupPlanResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableBackupPlanResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableBackupPlanResponseBody
  GetSuccess() *bool 
}

type EnableBackupPlanResponseBody struct {
  // The response code. The status code 200 indicates that the request was successful.
  // 
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  // The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
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

func (s EnableBackupPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupPlanResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBackupPlanResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableBackupPlanResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableBackupPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBackupPlanResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableBackupPlanResponseBody) SetCode(v string) *EnableBackupPlanResponseBody {
  s.Code = &v
  return s
}

func (s *EnableBackupPlanResponseBody) SetMessage(v string) *EnableBackupPlanResponseBody {
  s.Message = &v
  return s
}

func (s *EnableBackupPlanResponseBody) SetRequestId(v string) *EnableBackupPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBackupPlanResponseBody) SetSuccess(v bool) *EnableBackupPlanResponseBody {
  s.Success = &v
  return s
}

func (s *EnableBackupPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

