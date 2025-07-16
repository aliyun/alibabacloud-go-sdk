// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedirectTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *RedirectTaskRequest
	GetAppType() *string
	SetByManager(v string) *RedirectTaskRequest
	GetByManager() *string
	SetLanguage(v string) *RedirectTaskRequest
	GetLanguage() *string
	SetNowActionExecutorId(v string) *RedirectTaskRequest
	GetNowActionExecutorId() *string
	SetProcessInstanceId(v string) *RedirectTaskRequest
	GetProcessInstanceId() *string
	SetRemark(v string) *RedirectTaskRequest
	GetRemark() *string
	SetSystemToken(v string) *RedirectTaskRequest
	GetSystemToken() *string
	SetTaskId(v int64) *RedirectTaskRequest
	GetTaskId() *int64
}

type RedirectTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// APP_PBKxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// y
	ByManager *string `json:"ByManager,omitempty" xml:"ByManager,omitempty"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	NowActionExecutorId *string `json:"NowActionExecutorId,omitempty" xml:"NowActionExecutorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// instxxxx
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	SystemToken *string `json:"SystemToken,omitempty" xml:"SystemToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hexxxx
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RedirectTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RedirectTaskRequest) GoString() string {
	return s.String()
}

func (s *RedirectTaskRequest) GetAppType() *string {
	return s.AppType
}

func (s *RedirectTaskRequest) GetByManager() *string {
	return s.ByManager
}

func (s *RedirectTaskRequest) GetLanguage() *string {
	return s.Language
}

func (s *RedirectTaskRequest) GetNowActionExecutorId() *string {
	return s.NowActionExecutorId
}

func (s *RedirectTaskRequest) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *RedirectTaskRequest) GetRemark() *string {
	return s.Remark
}

func (s *RedirectTaskRequest) GetSystemToken() *string {
	return s.SystemToken
}

func (s *RedirectTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RedirectTaskRequest) SetAppType(v string) *RedirectTaskRequest {
	s.AppType = &v
	return s
}

func (s *RedirectTaskRequest) SetByManager(v string) *RedirectTaskRequest {
	s.ByManager = &v
	return s
}

func (s *RedirectTaskRequest) SetLanguage(v string) *RedirectTaskRequest {
	s.Language = &v
	return s
}

func (s *RedirectTaskRequest) SetNowActionExecutorId(v string) *RedirectTaskRequest {
	s.NowActionExecutorId = &v
	return s
}

func (s *RedirectTaskRequest) SetProcessInstanceId(v string) *RedirectTaskRequest {
	s.ProcessInstanceId = &v
	return s
}

func (s *RedirectTaskRequest) SetRemark(v string) *RedirectTaskRequest {
	s.Remark = &v
	return s
}

func (s *RedirectTaskRequest) SetSystemToken(v string) *RedirectTaskRequest {
	s.SystemToken = &v
	return s
}

func (s *RedirectTaskRequest) SetTaskId(v int64) *RedirectTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RedirectTaskRequest) Validate() error {
	return dara.Validate(s)
}
