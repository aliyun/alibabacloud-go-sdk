// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetThreadResponseBody
	GetCreateTime() *string
	SetDigitalEmployeeName(v string) *GetThreadResponseBody
	GetDigitalEmployeeName() *string
	SetRequestId(v string) *GetThreadResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetThreadResponseBody
	GetStatus() *string
	SetThreadId(v string) *GetThreadResponseBody
	GetThreadId() *string
	SetTitle(v string) *GetThreadResponseBody
	GetTitle() *string
	SetUpdateTime(v string) *GetThreadResponseBody
	GetUpdateTime() *string
	SetVariables(v *GetThreadResponseBodyVariables) *GetThreadResponseBody
	GetVariables() *GetThreadResponseBodyVariables
	SetVersion(v int64) *GetThreadResponseBody
	GetVersion() *int64
}

type GetThreadResponseBody struct {
	// example:
	//
	// 2025-12-19T15:19:55.040403272+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// thread-t71rrw-1s7r7z9gu0v2s
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// example:
	//
	// testLive
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// 2025-12-19T15:19:55.040403272+08:00
	UpdateTime *string                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Variables  *GetThreadResponseBodyVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Struct"`
	// example:
	//
	// 1231
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetThreadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetThreadResponseBody) GoString() string {
	return s.String()
}

func (s *GetThreadResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetThreadResponseBody) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *GetThreadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetThreadResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetThreadResponseBody) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetThreadResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetThreadResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetThreadResponseBody) GetVariables() *GetThreadResponseBodyVariables {
	return s.Variables
}

func (s *GetThreadResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *GetThreadResponseBody) SetCreateTime(v string) *GetThreadResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetThreadResponseBody) SetDigitalEmployeeName(v string) *GetThreadResponseBody {
	s.DigitalEmployeeName = &v
	return s
}

func (s *GetThreadResponseBody) SetRequestId(v string) *GetThreadResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetThreadResponseBody) SetStatus(v string) *GetThreadResponseBody {
	s.Status = &v
	return s
}

func (s *GetThreadResponseBody) SetThreadId(v string) *GetThreadResponseBody {
	s.ThreadId = &v
	return s
}

func (s *GetThreadResponseBody) SetTitle(v string) *GetThreadResponseBody {
	s.Title = &v
	return s
}

func (s *GetThreadResponseBody) SetUpdateTime(v string) *GetThreadResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetThreadResponseBody) SetVariables(v *GetThreadResponseBodyVariables) *GetThreadResponseBody {
	s.Variables = v
	return s
}

func (s *GetThreadResponseBody) SetVersion(v int64) *GetThreadResponseBody {
	s.Version = &v
	return s
}

func (s *GetThreadResponseBody) Validate() error {
	if s.Variables != nil {
		if err := s.Variables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetThreadResponseBodyVariables struct {
	// example:
	//
	// kubenest
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// hd1
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetThreadResponseBodyVariables) String() string {
	return dara.Prettify(s)
}

func (s GetThreadResponseBodyVariables) GoString() string {
	return s.String()
}

func (s *GetThreadResponseBodyVariables) GetProject() *string {
	return s.Project
}

func (s *GetThreadResponseBodyVariables) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetThreadResponseBodyVariables) SetProject(v string) *GetThreadResponseBodyVariables {
	s.Project = &v
	return s
}

func (s *GetThreadResponseBodyVariables) SetWorkspace(v string) *GetThreadResponseBodyVariables {
	s.Workspace = &v
	return s
}

func (s *GetThreadResponseBodyVariables) Validate() error {
	return dara.Validate(s)
}
