// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTaskConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationCode(v string) *QueryTaskConcurrencyRequest
	GetApplicationCode() *string
	SetCallerUacAccountId(v string) *QueryTaskConcurrencyRequest
	GetCallerUacAccountId() *string
	SetCurrentWorkspaceId(v string) *QueryTaskConcurrencyRequest
	GetCurrentWorkspaceId() *string
	SetTaskId(v int64) *QueryTaskConcurrencyRequest
	GetTaskId() *int64
}

type QueryTaskConcurrencyRequest struct {
	// example:
	//
	// B9191F0E57
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// abc123***
	CallerUacAccountId *string `json:"CallerUacAccountId,omitempty" xml:"CallerUacAccountId,omitempty"`
	// example:
	//
	// abc123***
	CurrentWorkspaceId *string `json:"CurrentWorkspaceId,omitempty" xml:"CurrentWorkspaceId,omitempty"`
	// example:
	//
	// 12345
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryTaskConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTaskConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *QueryTaskConcurrencyRequest) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *QueryTaskConcurrencyRequest) GetCallerUacAccountId() *string {
	return s.CallerUacAccountId
}

func (s *QueryTaskConcurrencyRequest) GetCurrentWorkspaceId() *string {
	return s.CurrentWorkspaceId
}

func (s *QueryTaskConcurrencyRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryTaskConcurrencyRequest) SetApplicationCode(v string) *QueryTaskConcurrencyRequest {
	s.ApplicationCode = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) SetCallerUacAccountId(v string) *QueryTaskConcurrencyRequest {
	s.CallerUacAccountId = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) SetCurrentWorkspaceId(v string) *QueryTaskConcurrencyRequest {
	s.CurrentWorkspaceId = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) SetTaskId(v int64) *QueryTaskConcurrencyRequest {
	s.TaskId = &v
	return s
}

func (s *QueryTaskConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
