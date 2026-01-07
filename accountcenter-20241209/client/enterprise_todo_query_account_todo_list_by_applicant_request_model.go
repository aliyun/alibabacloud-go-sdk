// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListByApplicantRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetAppName() *string 
  SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetNextToken() *string 
  SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetOperatePk() *string 
  SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetOrientedNbId() *string 
  SetPage(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetPage() *int32 
  SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetPageSize() *int32 
  SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetShowCompleteInfo() *bool 
  SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetStatus() *string 
  SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest
  GetTodoType() *string 
}

type EnterpriseTodoQueryAccountTodoListByApplicantRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  OperatePk *string `json:"OperatePk,omitempty" xml:"OperatePk,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  ShowCompleteInfo *bool `json:"ShowCompleteInfo,omitempty" xml:"ShowCompleteInfo,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  TodoType *string `json:"TodoType,omitempty" xml:"TodoType,omitempty"`
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListByApplicantRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetOperatePk() *string  {
  return s.OperatePk
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetPage() *int32  {
  return s.Page
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) GetTodoType() *string  {
  return s.TodoType
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetAppName(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.NextToken = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.OperatePk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetPage(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.Page = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.PageSize = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListByApplicantRequest {
  s.TodoType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListByApplicantRequest) Validate() error {
  return dara.Validate(s)
}

