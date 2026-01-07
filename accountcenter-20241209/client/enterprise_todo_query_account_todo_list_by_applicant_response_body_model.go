// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListByApplicantResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetCode() *string 
  SetData(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetData() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData 
  SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetMaxResults() *int32 
  SetMessage(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetMessage() *string 
  SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetNextToken() *string 
  SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody
  GetSuccess() *bool 
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetData() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData  {
  return s.Data
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetCode(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetData(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetMessage(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.NextToken = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetRequestId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) SetSuccess(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData struct {
  Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
  TodoList []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList `json:"TodoList,omitempty" xml:"TodoList,omitempty" type:"Repeated"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) GetCount() *int32  {
  return s.Count
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) GetTodoList() []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList  {
  return s.TodoList
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) SetCount(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData {
  s.Count = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) SetTodoList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData {
  s.TodoList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyData) Validate() error {
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

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList struct {
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
  FromLe *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe `json:"FromLe,omitempty" xml:"FromLe,omitempty" type:"Struct"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  ProcessList []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList `json:"ProcessList,omitempty" xml:"ProcessList,omitempty" type:"Repeated"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  TimeoutTime *int64 `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
  ToLe *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe `json:"ToLe,omitempty" xml:"ToLe,omitempty" type:"Struct"`
  ToLeAudit *bool `json:"ToLeAudit,omitempty" xml:"ToLeAudit,omitempty"`
  TodoId *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
  TodoType *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
  TodoView *string `json:"TodoView,omitempty" xml:"TodoView,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetApplicantAliyunId() *string  {
  return s.ApplicantAliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetApplicantPk() *string  {
  return s.ApplicantPk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetApplyRemark() *string  {
  return s.ApplyRemark
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetApplyTime() *int64  {
  return s.ApplyTime
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetAuditorAliyunId() *string  {
  return s.AuditorAliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetAuditorPk() *string  {
  return s.AuditorPk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetAuditorStatus() *string  {
  return s.AuditorStatus
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetCanceledTime() *int64  {
  return s.CanceledTime
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetClosed() *bool  {
  return s.Closed
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetCurrAuditor() *bool  {
  return s.CurrAuditor
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetFromLe() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe  {
  return s.FromLe
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetProcessList() []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList  {
  return s.ProcessList
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetTimeoutTime() *int64  {
  return s.TimeoutTime
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetToLe() *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe  {
  return s.ToLe
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetToLeAudit() *bool  {
  return s.ToLeAudit
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetTodoId() *string  {
  return s.TodoId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetTodoType() *string  {
  return s.TodoType
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) GetTodoView() *string  {
  return s.TodoView
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplicantAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ApplicantAliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplicantPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ApplicantPk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplyRemark(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ApplyRemark = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetApplyTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ApplyTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.AuditorAliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.AuditorPk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetAuditorStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.AuditorStatus = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetCanceledTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.CanceledTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetClosed(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.Closed = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetCurrAuditor(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.CurrAuditor = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetFromLe(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.FromLe = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetProcessList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ProcessList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTimeoutTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.TimeoutTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetToLe(v *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ToLe = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetToLeAudit(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.ToLeAudit = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.TodoId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.TodoType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) SetTodoView(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList {
  s.TodoView = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoList) Validate() error {
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

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe struct {
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
  ManageableAccountCount *int64 `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
  ManagedAccountCount *int64 `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
  ManagerList []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  SubjectName *string `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
  SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetLicenseNumber() *string  {
  return s.LicenseNumber
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetManageableAccountCount() *int64  {
  return s.ManageableAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetManagedAccountCount() *int64  {
  return s.ManagedAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetManagerList() []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList  {
  return s.ManagerList
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetSubjectName() *string  {
  return s.SubjectName
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) GetSubjectType() *string  {
  return s.SubjectType
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.LicenseNumber = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.ManageableAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.ManagedAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.ManagerList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.SubjectName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe {
  s.SubjectType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLe) Validate() error {
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

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
  Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GetPkEncoded() *string  {
  return s.PkEncoded
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) GetRole() *string  {
  return s.Role
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
  s.PkEncoded = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList {
  s.Role = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListFromLeManagerList) Validate() error {
  return dara.Validate(s)
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  AuditTime *int64 `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetAuditTime() *int64  {
  return s.AuditTime
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetRemark() *string  {
  return s.Remark
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetAuditTime(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.AuditTime = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetRemark(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.Remark = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListProcessList) Validate() error {
  return dara.Validate(s)
}

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe struct {
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  LicenseNumber *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
  ManageableAccountCount *int64 `json:"ManageableAccountCount,omitempty" xml:"ManageableAccountCount,omitempty"`
  ManagedAccountCount *int64 `json:"ManagedAccountCount,omitempty" xml:"ManagedAccountCount,omitempty"`
  ManagerList []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList `json:"ManagerList,omitempty" xml:"ManagerList,omitempty" type:"Repeated"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  SubjectName *string `json:"SubjectName,omitempty" xml:"SubjectName,omitempty"`
  SubjectType *string `json:"SubjectType,omitempty" xml:"SubjectType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetLicenseNumber() *string  {
  return s.LicenseNumber
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetManageableAccountCount() *int64  {
  return s.ManageableAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetManagedAccountCount() *int64  {
  return s.ManagedAccountCount
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetManagerList() []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList  {
  return s.ManagerList
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetSubjectName() *string  {
  return s.SubjectName
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) GetSubjectType() *string  {
  return s.SubjectType
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.EcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.LeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetLicenseNumber(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.LicenseNumber = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManageableAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.ManageableAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManagedAccountCount(v int64) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.ManagedAccountCount = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetManagerList(v []*EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.ManagerList = v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.NbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetSubjectName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.SubjectName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) SetSubjectType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe {
  s.SubjectType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLe) Validate() error {
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

type EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList struct {
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  PkEncoded *string `json:"PkEncoded,omitempty" xml:"PkEncoded,omitempty"`
  Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GetPkEncoded() *string  {
  return s.PkEncoded
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) GetRole() *string  {
  return s.Role
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetAliyunId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetPk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
  s.Pk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetPkEncoded(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
  s.PkEncoded = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) SetRole(v string) *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList {
  s.Role = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantResponseBodyDataTodoListToLeManagerList) Validate() error {
  return dara.Validate(s)
}

