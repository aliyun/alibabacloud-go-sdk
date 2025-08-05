// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableHanaBackupPlanResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnableHanaBackupPlanResponseBody
  GetCode() *string 
  SetMessage(v string) *EnableHanaBackupPlanResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnableHanaBackupPlanResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableHanaBackupPlanResponseBody
  GetSuccess() *bool 
}

type EnableHanaBackupPlanResponseBody struct {
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

func (s EnableHanaBackupPlanResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableHanaBackupPlanResponseBody) GoString() string {
  return s.String()
}

func (s *EnableHanaBackupPlanResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnableHanaBackupPlanResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnableHanaBackupPlanResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableHanaBackupPlanResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableHanaBackupPlanResponseBody) SetCode(v string) *EnableHanaBackupPlanResponseBody {
  s.Code = &v
  return s
}

func (s *EnableHanaBackupPlanResponseBody) SetMessage(v string) *EnableHanaBackupPlanResponseBody {
  s.Message = &v
  return s
}

func (s *EnableHanaBackupPlanResponseBody) SetRequestId(v string) *EnableHanaBackupPlanResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableHanaBackupPlanResponseBody) SetSuccess(v bool) *EnableHanaBackupPlanResponseBody {
  s.Success = &v
  return s
}

func (s *EnableHanaBackupPlanResponseBody) Validate() error {
  return dara.Validate(s)
}

