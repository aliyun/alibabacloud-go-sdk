// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveMpuTaskSeiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SetLiveMpuTaskSeiRequest
	GetAppId() *string
	SetCustomSei(v string) *SetLiveMpuTaskSeiRequest
	GetCustomSei() *string
	SetTaskId(v string) *SetLiveMpuTaskSeiRequest
	GetTaskId() *string
}

type SetLiveMpuTaskSeiRequest struct {
	// The application ID.
	//
	// >  The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The custom SEI.
	//
	// >  The value is a JSON string that can be up to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"myroom"}
	CustomSei *string `json:"CustomSei,omitempty" xml:"CustomSei,omitempty"`
	// The task ID.
	//
	// >  The ID can be up to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// AL-4bce036dd90277c50705b0599wgfffc7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SetLiveMpuTaskSeiRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLiveMpuTaskSeiRequest) GoString() string {
	return s.String()
}

func (s *SetLiveMpuTaskSeiRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetLiveMpuTaskSeiRequest) GetCustomSei() *string {
	return s.CustomSei
}

func (s *SetLiveMpuTaskSeiRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SetLiveMpuTaskSeiRequest) SetAppId(v string) *SetLiveMpuTaskSeiRequest {
	s.AppId = &v
	return s
}

func (s *SetLiveMpuTaskSeiRequest) SetCustomSei(v string) *SetLiveMpuTaskSeiRequest {
	s.CustomSei = &v
	return s
}

func (s *SetLiveMpuTaskSeiRequest) SetTaskId(v string) *SetLiveMpuTaskSeiRequest {
	s.TaskId = &v
	return s
}

func (s *SetLiveMpuTaskSeiRequest) Validate() error {
	return dara.Validate(s)
}
