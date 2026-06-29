// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskDetail interface {
	dara.Model
	String() string
	GoString() string
	SetAdmins(v *CreateTaskDetailAdmins) *CreateTaskDetail
	GetAdmins() *CreateTaskDetailAdmins
	SetAllowAppendData(v bool) *CreateTaskDetail
	GetAllowAppendData() *bool
	SetAssignConfig(v *TaskAssginConfig) *CreateTaskDetail
	GetAssignConfig() *TaskAssginConfig
	SetDatasetProxyRelations(v []*DatasetProxyConfig) *CreateTaskDetail
	GetDatasetProxyRelations() []*DatasetProxyConfig
	SetExif(v map[string]interface{}) *CreateTaskDetail
	GetExif() map[string]interface{}
	SetTags(v []*string) *CreateTaskDetail
	GetTags() []*string
	SetTaskName(v string) *CreateTaskDetail
	GetTaskName() *string
	SetTaskTemplateConfig(v *TaskTemplateConfig) *CreateTaskDetail
	GetTaskTemplateConfig() *TaskTemplateConfig
	SetTaskWorkflow(v []*CreateTaskDetailTaskWorkflow) *CreateTaskDetail
	GetTaskWorkflow() []*CreateTaskDetailTaskWorkflow
	SetTemplateId(v string) *CreateTaskDetail
	GetTemplateId() *string
	SetUUID(v string) *CreateTaskDetail
	GetUUID() *string
	SetVoteConfigs(v map[string]*CreateTaskDetailVoteInfo) *CreateTaskDetail
	GetVoteConfigs() map[string]*CreateTaskDetailVoteInfo
}

type CreateTaskDetail struct {
	// Administrators; defaults to the Creator.
	Admins *CreateTaskDetailAdmins `json:"Admins,omitempty" xml:"Admins,omitempty" type:"Struct"`
	// Indicates whether appending new data is supported. Valid values:
	//
	// - true: Supported
	//
	// - false: Not supported
	//
	// example:
	//
	// true
	AllowAppendData *bool `json:"AllowAppendData,omitempty" xml:"AllowAppendData,omitempty"`
	// Job assignment mechanism.
	//
	// This parameter is required.
	AssignConfig *TaskAssginConfig `json:"AssignConfig,omitempty" xml:"AssignConfig,omitempty"`
	// List of dataset configurations.
	//
	// This parameter is required.
	DatasetProxyRelations []*DatasetProxyConfig `json:"DatasetProxyRelations,omitempty" xml:"DatasetProxyRelations,omitempty" type:"Repeated"`
	// Additional configuration. JSON format.
	//
	// example:
	//
	// {  "TaskCaptionType": "DOC_LINK"}
	Exif map[string]interface{} `json:"Exif,omitempty" xml:"Exif,omitempty"`
	// List of job labels.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Job name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 测试任务202208101424
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Template Configuration.
	TaskTemplateConfig *TaskTemplateConfig `json:"TaskTemplateConfig,omitempty" xml:"TaskTemplateConfig,omitempty"`
	// List of pipelines.
	//
	// This parameter is required.
	TaskWorkflow []*CreateTaskDetailTaskWorkflow `json:"TaskWorkflow,omitempty" xml:"TaskWorkflow,omitempty" type:"Repeated"`
	// Template ID. For more information, see [ListTemplates](https://help.aliyun.com/document_detail/454654.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 152936***8342353920
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Unique identifier (UUID), controlled by the business side.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16623***80757311
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
	// Vote configuration
	VoteConfigs map[string]*CreateTaskDetailVoteInfo `json:"VoteConfigs,omitempty" xml:"VoteConfigs,omitempty"`
}

func (s CreateTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskDetail) GoString() string {
	return s.String()
}

func (s *CreateTaskDetail) GetAdmins() *CreateTaskDetailAdmins {
	return s.Admins
}

func (s *CreateTaskDetail) GetAllowAppendData() *bool {
	return s.AllowAppendData
}

func (s *CreateTaskDetail) GetAssignConfig() *TaskAssginConfig {
	return s.AssignConfig
}

func (s *CreateTaskDetail) GetDatasetProxyRelations() []*DatasetProxyConfig {
	return s.DatasetProxyRelations
}

func (s *CreateTaskDetail) GetExif() map[string]interface{} {
	return s.Exif
}

func (s *CreateTaskDetail) GetTags() []*string {
	return s.Tags
}

func (s *CreateTaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTaskDetail) GetTaskTemplateConfig() *TaskTemplateConfig {
	return s.TaskTemplateConfig
}

func (s *CreateTaskDetail) GetTaskWorkflow() []*CreateTaskDetailTaskWorkflow {
	return s.TaskWorkflow
}

func (s *CreateTaskDetail) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateTaskDetail) GetUUID() *string {
	return s.UUID
}

func (s *CreateTaskDetail) GetVoteConfigs() map[string]*CreateTaskDetailVoteInfo {
	return s.VoteConfigs
}

func (s *CreateTaskDetail) SetAdmins(v *CreateTaskDetailAdmins) *CreateTaskDetail {
	s.Admins = v
	return s
}

func (s *CreateTaskDetail) SetAllowAppendData(v bool) *CreateTaskDetail {
	s.AllowAppendData = &v
	return s
}

func (s *CreateTaskDetail) SetAssignConfig(v *TaskAssginConfig) *CreateTaskDetail {
	s.AssignConfig = v
	return s
}

func (s *CreateTaskDetail) SetDatasetProxyRelations(v []*DatasetProxyConfig) *CreateTaskDetail {
	s.DatasetProxyRelations = v
	return s
}

func (s *CreateTaskDetail) SetExif(v map[string]interface{}) *CreateTaskDetail {
	s.Exif = v
	return s
}

func (s *CreateTaskDetail) SetTags(v []*string) *CreateTaskDetail {
	s.Tags = v
	return s
}

func (s *CreateTaskDetail) SetTaskName(v string) *CreateTaskDetail {
	s.TaskName = &v
	return s
}

func (s *CreateTaskDetail) SetTaskTemplateConfig(v *TaskTemplateConfig) *CreateTaskDetail {
	s.TaskTemplateConfig = v
	return s
}

func (s *CreateTaskDetail) SetTaskWorkflow(v []*CreateTaskDetailTaskWorkflow) *CreateTaskDetail {
	s.TaskWorkflow = v
	return s
}

func (s *CreateTaskDetail) SetTemplateId(v string) *CreateTaskDetail {
	s.TemplateId = &v
	return s
}

func (s *CreateTaskDetail) SetUUID(v string) *CreateTaskDetail {
	s.UUID = &v
	return s
}

func (s *CreateTaskDetail) SetVoteConfigs(v map[string]*CreateTaskDetailVoteInfo) *CreateTaskDetail {
	s.VoteConfigs = v
	return s
}

func (s *CreateTaskDetail) Validate() error {
	if s.Admins != nil {
		if err := s.Admins.Validate(); err != nil {
			return err
		}
	}
	if s.AssignConfig != nil {
		if err := s.AssignConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DatasetProxyRelations != nil {
		for _, item := range s.DatasetProxyRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskTemplateConfig != nil {
		if err := s.TaskTemplateConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TaskWorkflow != nil {
		for _, item := range s.TaskWorkflow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskDetailAdmins struct {
	// List of administrators.
	Users []*SimpleUser `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateTaskDetailAdmins) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskDetailAdmins) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailAdmins) GetUsers() []*SimpleUser {
	return s.Users
}

func (s *CreateTaskDetailAdmins) SetUsers(v []*SimpleUser) *CreateTaskDetailAdmins {
	s.Users = v
	return s
}

func (s *CreateTaskDetailAdmins) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskDetailTaskWorkflow struct {
	// Node name. Valid values:
	//
	// - MARK: Annotate.
	//
	// - CHECK: Check.
	//
	// - SAMPLING: Validate.
	//
	// example:
	//
	// MARK
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s CreateTaskDetailTaskWorkflow) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskDetailTaskWorkflow) GoString() string {
	return s.String()
}

func (s *CreateTaskDetailTaskWorkflow) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateTaskDetailTaskWorkflow) SetNodeName(v string) *CreateTaskDetailTaskWorkflow {
	s.NodeName = &v
	return s
}

func (s *CreateTaskDetailTaskWorkflow) Validate() error {
	return dara.Validate(s)
}
