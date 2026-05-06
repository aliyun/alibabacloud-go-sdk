// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateTaskGroupRequest
	GetLang() *string
	SetRegId(v string) *CreateTaskGroupRequest
	GetRegId() *string
	SetSampleIds(v string) *CreateTaskGroupRequest
	GetSampleIds() *string
	SetScenes(v string) *CreateTaskGroupRequest
	GetScenes() *string
	SetServiceCodes(v string) *CreateTaskGroupRequest
	GetServiceCodes() *string
	SetServiceNames(v string) *CreateTaskGroupRequest
	GetServiceNames() *string
	SetTab(v string) *CreateTaskGroupRequest
	GetTab() *string
	SetTaskGroupName(v string) *CreateTaskGroupRequest
	GetTaskGroupName() *string
	SetType(v string) *CreateTaskGroupRequest
	GetType() *string
}

type CreateTaskGroupRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// example:
	//
	// 1,2
	SampleIds *string `json:"SampleIds,omitempty" xml:"SampleIds,omitempty"`
	// example:
	//
	// [\\"porn\\",\\"terrorism\\",\\"ad\\"]
	Scenes *string `json:"Scenes,omitempty" xml:"Scenes,omitempty"`
	// example:
	//
	// oss
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
	// example:
	//
	// Test
	ServiceNames *string `json:"ServiceNames,omitempty" xml:"ServiceNames,omitempty"`
	// example:
	//
	// INTERNET
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// TeskGroupTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
	// example:
	//
	// SAF_CONSOLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateTaskGroupRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateTaskGroupRequest) GetSampleIds() *string {
	return s.SampleIds
}

func (s *CreateTaskGroupRequest) GetScenes() *string {
	return s.Scenes
}

func (s *CreateTaskGroupRequest) GetServiceCodes() *string {
	return s.ServiceCodes
}

func (s *CreateTaskGroupRequest) GetServiceNames() *string {
	return s.ServiceNames
}

func (s *CreateTaskGroupRequest) GetTab() *string {
	return s.Tab
}

func (s *CreateTaskGroupRequest) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *CreateTaskGroupRequest) GetType() *string {
	return s.Type
}

func (s *CreateTaskGroupRequest) SetLang(v string) *CreateTaskGroupRequest {
	s.Lang = &v
	return s
}

func (s *CreateTaskGroupRequest) SetRegId(v string) *CreateTaskGroupRequest {
	s.RegId = &v
	return s
}

func (s *CreateTaskGroupRequest) SetSampleIds(v string) *CreateTaskGroupRequest {
	s.SampleIds = &v
	return s
}

func (s *CreateTaskGroupRequest) SetScenes(v string) *CreateTaskGroupRequest {
	s.Scenes = &v
	return s
}

func (s *CreateTaskGroupRequest) SetServiceCodes(v string) *CreateTaskGroupRequest {
	s.ServiceCodes = &v
	return s
}

func (s *CreateTaskGroupRequest) SetServiceNames(v string) *CreateTaskGroupRequest {
	s.ServiceNames = &v
	return s
}

func (s *CreateTaskGroupRequest) SetTab(v string) *CreateTaskGroupRequest {
	s.Tab = &v
	return s
}

func (s *CreateTaskGroupRequest) SetTaskGroupName(v string) *CreateTaskGroupRequest {
	s.TaskGroupName = &v
	return s
}

func (s *CreateTaskGroupRequest) SetType(v string) *CreateTaskGroupRequest {
	s.Type = &v
	return s
}

func (s *CreateTaskGroupRequest) Validate() error {
	return dara.Validate(s)
}
