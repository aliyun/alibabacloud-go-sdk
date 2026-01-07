// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetCode() *string 
  SetData(v *EnterpriseTodoQueryAccountTodoListResponseBodyData) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetData() *EnterpriseTodoQueryAccountTodoListResponseBodyData 
  SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetMaxResults() *int32 
  SetMessage(v string) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetMessage() *string 
  SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetNextToken() *string 
  SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListResponseBody
  GetSuccess() *bool 
}

type EnterpriseTodoQueryAccountTodoListResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseTodoQueryAccountTodoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetData() *EnterpriseTodoQueryAccountTodoListResponseBodyData  {
  return s.Data
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetCode(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetData(v *EnterpriseTodoQueryAccountTodoListResponseBodyData) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetMessage(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.NextToken = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListResponseBodyData struct {
  Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
  TodoList []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Repeated"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) GetCount() *int32  {
  return s.Count
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) GetTodoList() []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList  {
  return s.TodoList
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) SetCount(v int32) *EnterpriseTodoQueryAccountTodoListResponseBodyData {
  s.Count = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) SetTodoList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) *EnterpriseTodoQueryAccountTodoListResponseBodyData {
  s.TodoList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyData) Validate() error {
  if s.TodoList != nil {
    for _, item := range s.TodoList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  ApplicantAliyunId *string `json:"ApplicantAliyunId,omitempty" xml:"ApplicantAliyunId,omitempty"`
  ApplicantPk *string `json:"ApplicantPk,omitempty" xml:"ApplicantPk,omitempty"`
  ApplyRemark *string `json:"ApplyRemark,omitempty" xml:"ApplyRemark,omitempty"`
  ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
  AuditorAliyunId *string `json:"AuditorAliyunId,omitempty" xml:"AuditorAliyunId,omitempty"`
  AuditorPk *string `json:"AuditorPk,omitempty" xml:"AuditorPk,omitempty"`
  AuditorStatus *string `json:"AuditorStatus,omitempty" xml:"AuditorStatus,omitempty"`
  CanceledTime *int64 `json:"CanceledTime,omitempty" xml:"CanceledTime,omitempty"`
  Closed *bool `json:"Closed,omitempty" xml:"Closed,omitempty"`
  CurrAuditor *bool `json:"CurrAuditor,omitempty" xml:"CurrAuditor,omitempty"`
  FromLe *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe `json:"FromLe,omitempty" xml:"FromLe,omitempty" type:"Struct"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  ProcessList []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  TimeoutTime *int64 `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
  ToLe *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe `json:"ToLe,omitempty" xml:"ToLe,omitempty" type:"Struct"`
  ToLeAudit *bool `json:"ToLeAudit,omitempty" xml:"ToLeAudit,omitempty"`
  TodoId *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
  TodoType *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
  TodoView *string `json:"TodoView,omitempty" xml:"TodoView,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetApplicantAliyunId() *string  {
  return s.ApplicantAliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetApplicantPk() *string  {
  return s.ApplicantPk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetApplyRemark() *string  {
  return s.ApplyRemark
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetApplyTime() *int64  {
  return s.ApplyTime
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetAuditorAliyunId() *string  {
  return s.AuditorAliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetAuditorPk() *string  {
  return s.AuditorPk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetAuditorStatus() *string  {
  return s.AuditorStatus
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetCanceledTime() *int64  {
  return s.CanceledTime
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetClosed() *bool  {
  return s.Closed
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetCurrAuditor() *bool  {
  return s.CurrAuditor
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetFromLe() *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe  {
  return s.FromLe
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetProcessList() []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList  {
  return s.ProcessList
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetTimeoutTime() *int64  {
  return s.TimeoutTime
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetToLe() *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe  {
  return s.ToLe
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetToLeAudit() *bool  {
  return s.ToLeAudit
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetTodoId() *string  {
  return s.TodoId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetTodoType() *string  {
  return s.TodoType
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) GetTodoView() *string  {
  return s.TodoView
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplicantAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ApplicantAliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplicantPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ApplicantPk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplyRemark(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ApplyRemark = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetApplyTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ApplyTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.AuditorAliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.AuditorPk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetAuditorStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.AuditorStatus = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetCanceledTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.CanceledTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetClosed(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.Closed = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetCurrAuditor(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.CurrAuditor = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetFromLe(v *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.FromLe = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetProcessList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ProcessList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTimeoutTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.TimeoutTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetToLe(v *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ToLe = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetToLeAudit(v bool) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.ToLeAudit = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.TodoId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.TodoType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) SetTodoView(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList {
  s.TodoView = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoList) Validate() error {
  if s.FromLe != nil {
    if err := s.FromLe.Validate(); err != nil {
      return err
    }
  }
  if s.ProcessList != nil {
    for _, item := range s.ProcessList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.ToLe != nil {
    if err := s.ToLe.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe struct {
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
  ManageableAccountCount *int64 `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
  ManagedAccountCount *int64 `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
  ManagerList []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  SubjectName *string `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
  SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetLicenseNumber() *string  {
  return s.LicenseNumber
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetManageableAccountCount() *int64  {
  return s.ManageableAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetManagedAccountCount() *int64  {
  return s.ManagedAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetManagerList() []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList  {
  return s.ManagerList
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetSubjectName() *string  {
  return s.SubjectName
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) GetSubjectType() *string  {
  return s.SubjectType
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.LicenseNumber = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.ManageableAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.ManagedAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.ManagerList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.SubjectName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe {
  s.SubjectType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLe) Validate() error {
  if s.ManagerList != nil {
    for _, item := range s.ManagerList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
  Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GetPkEncoded() *string  {
  return s.PkEncoded
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) GetRole() *string  {
  return s.Role
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
  s.PkEncoded = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList {
  s.Role = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListFromLeManagerList) Validate() error {
  return dara.Validate(s)
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  AuditTime *int64 `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetAuditTime() *int64  {
  return s.AuditTime
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetRemark() *string  {
  return s.Remark
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetAuditTime(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.AuditTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetRemark(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.Remark = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListProcessList) Validate() error {
  return dara.Validate(s)
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe struct {
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
  ManageableAccountCount *int64 `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
  ManagedAccountCount *int64 `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
  ManagerList []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  SubjectName *string `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
  SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetLicenseNumber() *string  {
  return s.LicenseNumber
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetManageableAccountCount() *int64  {
  return s.ManageableAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetManagedAccountCount() *int64  {
  return s.ManagedAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetManagerList() []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList  {
  return s.ManagerList
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetSubjectName() *string  {
  return s.SubjectName
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) GetSubjectType() *string  {
  return s.SubjectType
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.LicenseNumber = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.ManageableAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.ManagedAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.ManagerList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.SubjectName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe {
  s.SubjectType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLe) Validate() error {
  if s.ManagerList != nil {
    for _, item := range s.ManagerList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
  Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GetPkEncoded() *string  {
  return s.PkEncoded
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) GetRole() *string  {
  return s.Role
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
  s.PkEncoded = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList {
  s.Role = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListResponseBodyDataTodoListToLeManagerList) Validate() error {
  return dara.Validate(s)
}

