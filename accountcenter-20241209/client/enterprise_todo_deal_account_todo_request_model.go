// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseTodoDealAccountTodoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppName(v string) *EnterpriseTodoDealAccountTodoRequest
  GetAppName() *string 
  SetOrientedEcId(v string) *EnterpriseTodoDealAccountTodoRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseTodoDealAccountTodoRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseTodoDealAccountTodoRequest
  GetOrientedNbId() *string 
  SetRemark(v string) *EnterpriseTodoDealAccountTodoRequest
  GetRemark() *string 
  SetStatus(v string) *EnterpriseTodoDealAccountTodoRequest
  GetStatus() *string 
  SetTodoId(v string) *EnterpriseTodoDealAccountTodoRequest
  GetTodoId() *string 
}

type EnterpriseTodoDealAccountTodoRequest struct {
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  TodoId *string `json:"TodoId,omitempty" xml:"TodoId,omitempty"`
}

func (s EnterpriseTodoDealAccountTodoRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseTodoDealAccountTodoRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetRemark() *string  {
  return s.Remark
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetStatus() *string  {
  return s.Status
}

func (s *EnterpriseTodoDealAccountTodoRequest) GetTodoId() *string  {
  return s.TodoId
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetAppName(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.AppName = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedEcId(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedLeId(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetOrientedNbId(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetRemark(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.Remark = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetStatus(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.Status = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) SetTodoId(v string) *EnterpriseTodoDealAccountTodoRequest {
  s.TodoId = &v
  return s
}

func (s *EnterpriseTodoDealAccountTodoRequest) Validate() error {
  return dara.Validate(s)
}

