// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskExecError interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TaskExecError
	GetCode() *string
	SetExtraInfo(v string) *TaskExecError
	GetExtraInfo() *string
	SetMessage(v string) *TaskExecError
	GetMessage() *string
	SetRequestId(v string) *TaskExecError
	GetRequestId() *string
	SetTitle(v string) *TaskExecError
	GetTitle() *string
}

type TaskExecError struct {
	// example:
	//
	// AccessDenied
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 部署辅助函数权限不足，需要添加额外的权限以解决问题。https://help.aliyun.com
	ExtraInfo *string `json:"extraInfo,omitempty" xml:"extraInfo,omitempty"`
	// example:
	//
	// 部署服务[_appcenter-xxx]失败，权限不足
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1-26d1287xxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 权限不足错误
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s TaskExecError) String() string {
	return dara.Prettify(s)
}

func (s TaskExecError) GoString() string {
	return s.String()
}

func (s *TaskExecError) GetCode() *string {
	return s.Code
}

func (s *TaskExecError) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *TaskExecError) GetMessage() *string {
	return s.Message
}

func (s *TaskExecError) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskExecError) GetTitle() *string {
	return s.Title
}

func (s *TaskExecError) SetCode(v string) *TaskExecError {
	s.Code = &v
	return s
}

func (s *TaskExecError) SetExtraInfo(v string) *TaskExecError {
	s.ExtraInfo = &v
	return s
}

func (s *TaskExecError) SetMessage(v string) *TaskExecError {
	s.Message = &v
	return s
}

func (s *TaskExecError) SetRequestId(v string) *TaskExecError {
	s.RequestId = &v
	return s
}

func (s *TaskExecError) SetTitle(v string) *TaskExecError {
	s.Title = &v
	return s
}

func (s *TaskExecError) Validate() error {
	return dara.Validate(s)
}
