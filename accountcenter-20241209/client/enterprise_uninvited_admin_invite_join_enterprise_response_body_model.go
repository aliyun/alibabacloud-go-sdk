// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody
  GetCode() *string 
  SetData(v []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody
  GetData() []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData 
  SetMessage(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody
  GetSuccess() *bool 
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GetData() []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData  {
  return s.Data
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetCode(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetData(v []*EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetMessage(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetRequestId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) SetSuccess(v bool) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBody) Validate() error {
  if s.Data != nil {
    for _, item := range s.Data {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData struct {
  ApplicantAliyunId *string `json:"ApplicantAliyunId,omitempty" xml:"ApplicantAliyunId,omitempty"`
  ApplicantPk *string `json:"ApplicantPk,omitempty" xml:"ApplicantPk,omitempty"`
  ApplyRemark *string `json:"ApplyRemark,omitempty" xml:"ApplyRemark,omitempty"`
  ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
  AuditorAliyunId *string `json:"AuditorAliyunId,omitempty" xml:"AuditorAliyunId,omitempty"`
  AuditorPk *string `json:"AuditorPk,omitempty" xml:"AuditorPk,omitempty"`
  EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
  LeId *string `json:"LeId,omitempty" xml:"LeId,omitempty"`
  LeLicenseNo *string `json:"LeLicenseNo,omitempty" xml:"LeLicenseNo,omitempty"`
  LeName *string `json:"LeName,omitempty" xml:"LeName,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  NbId *string `json:"NbId,omitempty" xml:"NbId,omitempty"`
  OperatedAliyunId *string `json:"OperatedAliyunId,omitempty" xml:"OperatedAliyunId,omitempty"`
  OperatedPk *string `json:"OperatedPk,omitempty" xml:"OperatedPk,omitempty"`
  Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
  TimeoutTime *int64 `json:"TimeoutTime,omitempty" xml:"TimeoutTime,omitempty"`
  TodoId *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
  TodoType *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetApplicantAliyunId() *string  {
  return s.ApplicantAliyunId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetApplicantPk() *string  {
  return s.ApplicantPk
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetApplyRemark() *string  {
  return s.ApplyRemark
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetApplyTime() *int64  {
  return s.ApplyTime
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetAuditorAliyunId() *string  {
  return s.AuditorAliyunId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetAuditorPk() *string  {
  return s.AuditorPk
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetEcId() *string  {
  return s.EcId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetLeId() *string  {
  return s.LeId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetLeLicenseNo() *string  {
  return s.LeLicenseNo
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetLeName() *string  {
  return s.LeName
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetNbId() *string  {
  return s.NbId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetOperatedAliyunId() *string  {
  return s.OperatedAliyunId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetOperatedPk() *string  {
  return s.OperatedPk
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetSuccess() *string  {
  return s.Success
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetTimeoutTime() *int64  {
  return s.TimeoutTime
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetTodoId() *string  {
  return s.TodoId
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) GetTodoType() *string  {
  return s.TodoType
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplicantAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.ApplicantAliyunId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplicantPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.ApplicantPk = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplyRemark(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.ApplyRemark = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetApplyTime(v int64) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.ApplyTime = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetAuditorAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.AuditorAliyunId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetAuditorPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.AuditorPk = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetEcId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.EcId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.LeId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeLicenseNo(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.LeLicenseNo = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetLeName(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.LeName = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetMessage(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.Message = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetNbId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.NbId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetOperatedAliyunId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.OperatedAliyunId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetOperatedPk(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.OperatedPk = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetSuccess(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.Success = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTimeoutTime(v int64) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.TimeoutTime = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTodoId(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.TodoId = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) SetTodoType(v string) *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData {
  s.TodoType = &v
  return s
}

func (s *EnterpriseUninvitedAdminInviteJoinEnterpriseResponseBodyData) Validate() error {
  return dara.Validate(s)
}

