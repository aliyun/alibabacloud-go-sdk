// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoQueryAccountTodoListRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetAppName() *string 
  SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListRequest
  GetMaxResults() *int32 
  SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetNextToken() *string 
  SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetOperatePk() *string 
  SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetOrientedNbId() *string 
  SetPage(v int32) *EnterpriseTodoQueryAccountTodoListRequest
  GetPage() *int32 
  SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListRequest
  GetPageSize() *int32 
  SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListRequest
  GetShowCompleteInfo() *bool 
  SetStatus(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetStatus() *string 
  SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListRequest
  GetTodoType() *string 
}

type EnterpriseTodoQueryAccountTodoListRequest struct {
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

func (s EnterpriseTodoQueryAccountTodoListRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoQueryAccountTodoListRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetOperatePk() *string  {
  return s.OperatePk
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetPage() *int32  {
  return s.Page
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetShowCompleteInfo() *bool  {
  return s.ShowCompleteInfo
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) GetTodoType() *string  {
  return s.TodoType
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetAppName(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetMaxResults(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetNextToken(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.NextToken = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOperatePk(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.OperatePk = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedEcId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedLeId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetOrientedNbId(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetPage(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
  s.Page = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetPageSize(v int32) *EnterpriseTodoQueryAccountTodoListRequest {
  s.PageSize = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetShowCompleteInfo(v bool) *EnterpriseTodoQueryAccountTodoListRequest {
  s.ShowCompleteInfo = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetStatus(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) SetTodoType(v string) *EnterpriseTodoQueryAccountTodoListRequest {
  s.TodoType = &v
  return s
}

func (s *EnterpriseTodoQueryAccountTodoListRequest) Validate() error {
  return dara.Validate(s)
}

