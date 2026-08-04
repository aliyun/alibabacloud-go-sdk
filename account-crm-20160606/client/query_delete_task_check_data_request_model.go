// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeleteTaskCheckDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *QueryDeleteTaskCheckDataRequest
	GetAgAccountType() *string
	SetAppName(v string) *QueryDeleteTaskCheckDataRequest
	GetAppName() *string
	SetLongLang(v string) *QueryDeleteTaskCheckDataRequest
	GetLongLang() *string
	SetMpk(v string) *QueryDeleteTaskCheckDataRequest
	GetMpk() *string
	SetPk(v string) *QueryDeleteTaskCheckDataRequest
	GetPk() *string
	SetTaskId(v string) *QueryDeleteTaskCheckDataRequest
	GetTaskId() *string
	SetTaskType(v string) *QueryDeleteTaskCheckDataRequest
	GetTaskType() *string
}

type QueryDeleteTaskCheckDataRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	LongLang      *string `json:"LongLang,omitempty" xml:"LongLang,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
	// This parameter is required.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s QueryDeleteTaskCheckDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDeleteTaskCheckDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDeleteTaskCheckDataRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *QueryDeleteTaskCheckDataRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryDeleteTaskCheckDataRequest) GetLongLang() *string {
	return s.LongLang
}

func (s *QueryDeleteTaskCheckDataRequest) GetMpk() *string {
	return s.Mpk
}

func (s *QueryDeleteTaskCheckDataRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryDeleteTaskCheckDataRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryDeleteTaskCheckDataRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *QueryDeleteTaskCheckDataRequest) SetAgAccountType(v string) *QueryDeleteTaskCheckDataRequest {
	s.AgAccountType = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetAppName(v string) *QueryDeleteTaskCheckDataRequest {
	s.AppName = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetLongLang(v string) *QueryDeleteTaskCheckDataRequest {
	s.LongLang = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetMpk(v string) *QueryDeleteTaskCheckDataRequest {
	s.Mpk = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetPk(v string) *QueryDeleteTaskCheckDataRequest {
	s.Pk = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetTaskId(v string) *QueryDeleteTaskCheckDataRequest {
	s.TaskId = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) SetTaskType(v string) *QueryDeleteTaskCheckDataRequest {
	s.TaskType = &v
	return s
}

func (s *QueryDeleteTaskCheckDataRequest) Validate() error {
	return dara.Validate(s)
}
