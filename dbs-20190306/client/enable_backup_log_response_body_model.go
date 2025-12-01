// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableBackupLogResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetBackupPlanId(v string) *EnableBackupLogResponseBody
  GetBackupPlanId() *string 
  SetErrCode(v string) *EnableBackupLogResponseBody
  GetErrCode() *string 
  SetErrMessage(v string) *EnableBackupLogResponseBody
  GetErrMessage() *string 
  SetHttpStatusCode(v int32) *EnableBackupLogResponseBody
  GetHttpStatusCode() *int32 
  SetNeedPrecheck(v bool) *EnableBackupLogResponseBody
  GetNeedPrecheck() *bool 
  SetRequestId(v string) *EnableBackupLogResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnableBackupLogResponseBody
  GetSuccess() *bool 
}

type EnableBackupLogResponseBody struct {
  // The backup schedule ID.
  // 
  // example:
  // 
  // dbstooi01xxxx
  BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
  // The error code.
  // 
  // example:
  // 
  // Param.NotFound
  ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
  // The error message.
  // 
  // example:
  // 
  // findValidDBSJob error
  ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
  // The HTTP status code.
  // 
  // example:
  // 
  // 200
  HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
  // Indicates whether a precheck is triggered. Valid values:
  // 
  // 	- true: A precheck is triggered. You must call the [StartBackupPlan](https://help.aliyun.com/document_detail/2869816.html) operation to start the backup schedule.
  // 
  // 	- false: No precheck is triggered.
  // 
  // example:
  // 
  // false
  NeedPrecheck *bool `json:"NeedPrecheck,omitempty" xml:"NeedPrecheck,omitempty"`
  // The ID of the request.
  // 
  // example:
  // 
  // D6E068C3-25BC-455A-85FE-45F0B22ECB1F
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // Indicates whether the request is successful. Valid values:
  // 
  // 	- true: The request is successful.
  // 
  // 	- false: The request fails.
  // 
  // example:
  // 
  // true
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableBackupLogResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableBackupLogResponseBody) GoString() string {
  return s.String()
}

func (s *EnableBackupLogResponseBody) GetBackupPlanId() *string  {
  return s.BackupPlanId
}

func (s *EnableBackupLogResponseBody) GetErrCode() *string  {
  return s.ErrCode
}

func (s *EnableBackupLogResponseBody) GetErrMessage() *string  {
  return s.ErrMessage
}

func (s *EnableBackupLogResponseBody) GetHttpStatusCode() *int32  {
  return s.HttpStatusCode
}

func (s *EnableBackupLogResponseBody) GetNeedPrecheck() *bool  {
  return s.NeedPrecheck
}

func (s *EnableBackupLogResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableBackupLogResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnableBackupLogResponseBody) SetBackupPlanId(v string) *EnableBackupLogResponseBody {
  s.BackupPlanId = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetErrCode(v string) *EnableBackupLogResponseBody {
  s.ErrCode = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetErrMessage(v string) *EnableBackupLogResponseBody {
  s.ErrMessage = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetHttpStatusCode(v int32) *EnableBackupLogResponseBody {
  s.HttpStatusCode = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetNeedPrecheck(v bool) *EnableBackupLogResponseBody {
  s.NeedPrecheck = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetRequestId(v string) *EnableBackupLogResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableBackupLogResponseBody) SetSuccess(v bool) *EnableBackupLogResponseBody {
  s.Success = &v
  return s
}

func (s *EnableBackupLogResponseBody) Validate() error {
  return dara.Validate(s)
}

