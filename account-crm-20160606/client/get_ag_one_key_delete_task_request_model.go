// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgOneKeyDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *GetAgOneKeyDeleteTaskRequest
	GetAgAccountType() *string
	SetAppName(v string) *GetAgOneKeyDeleteTaskRequest
	GetAppName() *string
	SetMpk(v string) *GetAgOneKeyDeleteTaskRequest
	GetMpk() *string
	SetPk(v string) *GetAgOneKeyDeleteTaskRequest
	GetPk() *string
	SetTaskId(v string) *GetAgOneKeyDeleteTaskRequest
	GetTaskId() *string
}

type GetAgOneKeyDeleteTaskRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAgOneKeyDeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgOneKeyDeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAgOneKeyDeleteTaskRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *GetAgOneKeyDeleteTaskRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAgOneKeyDeleteTaskRequest) GetMpk() *string {
	return s.Mpk
}

func (s *GetAgOneKeyDeleteTaskRequest) GetPk() *string {
	return s.Pk
}

func (s *GetAgOneKeyDeleteTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAgOneKeyDeleteTaskRequest) SetAgAccountType(v string) *GetAgOneKeyDeleteTaskRequest {
	s.AgAccountType = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskRequest) SetAppName(v string) *GetAgOneKeyDeleteTaskRequest {
	s.AppName = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskRequest) SetMpk(v string) *GetAgOneKeyDeleteTaskRequest {
	s.Mpk = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskRequest) SetPk(v string) *GetAgOneKeyDeleteTaskRequest {
	s.Pk = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskRequest) SetTaskId(v string) *GetAgOneKeyDeleteTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetAgOneKeyDeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
