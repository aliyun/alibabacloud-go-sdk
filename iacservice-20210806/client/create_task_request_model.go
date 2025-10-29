// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoApply(v bool) *CreateTaskRequest
	GetAutoApply() *bool
	SetAutoDestroy(v bool) *CreateTaskRequest
	GetAutoDestroy() *bool
	SetClientToken(v string) *CreateTaskRequest
	GetClientToken() *string
	SetDescription(v string) *CreateTaskRequest
	GetDescription() *string
	SetGroupInfo(v *CreateTaskRequestGroupInfo) *CreateTaskRequest
	GetGroupInfo() *CreateTaskRequestGroupInfo
	SetInitModuleState(v bool) *CreateTaskRequest
	GetInitModuleState() *bool
	SetModuleId(v string) *CreateTaskRequest
	GetModuleId() *string
	SetModuleVersion(v string) *CreateTaskRequest
	GetModuleVersion() *string
	SetName(v string) *CreateTaskRequest
	GetName() *string
	SetProtectionStrategy(v []*string) *CreateTaskRequest
	GetProtectionStrategy() []*string
	SetRamRole(v string) *CreateTaskRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *CreateTaskRequest
	GetSkipPropertyValidation() *bool
	SetTags(v []*CreateTaskRequestTags) *CreateTaskRequest
	GetTags() []*CreateTaskRequestTags
	SetTaskBackend(v *CreateTaskRequestTaskBackend) *CreateTaskRequest
	GetTaskBackend() *CreateTaskRequestTaskBackend
	SetTerraformVersion(v string) *CreateTaskRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *CreateTaskRequest
	GetTriggerStrategy() *string
}

type CreateTaskRequest struct {
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// demo
	Description     *string                     `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo       *CreateTaskRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	InitModuleState *bool                       `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mod-148e7853433574fff6b316f4eb737e
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Name               *string   `json:"name,omitempty" xml:"name,omitempty"`
	ProtectionStrategy []*string `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	RamRole                *string                       `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	SkipPropertyValidation *bool                         `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	Tags                   []*CreateTaskRequestTags      `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TaskBackend            *CreateTaskRequestTaskBackend `json:"taskBackend,omitempty" xml:"taskBackend,omitempty" type:"Struct"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Always
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *CreateTaskRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *CreateTaskRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateTaskRequest) GetGroupInfo() *CreateTaskRequestGroupInfo {
	return s.GroupInfo
}

func (s *CreateTaskRequest) GetInitModuleState() *bool {
	return s.InitModuleState
}

func (s *CreateTaskRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *CreateTaskRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *CreateTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateTaskRequest) GetProtectionStrategy() []*string {
	return s.ProtectionStrategy
}

func (s *CreateTaskRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateTaskRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *CreateTaskRequest) GetTags() []*CreateTaskRequestTags {
	return s.Tags
}

func (s *CreateTaskRequest) GetTaskBackend() *CreateTaskRequestTaskBackend {
	return s.TaskBackend
}

func (s *CreateTaskRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *CreateTaskRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *CreateTaskRequest) SetAutoApply(v bool) *CreateTaskRequest {
	s.AutoApply = &v
	return s
}

func (s *CreateTaskRequest) SetAutoDestroy(v bool) *CreateTaskRequest {
	s.AutoDestroy = &v
	return s
}

func (s *CreateTaskRequest) SetClientToken(v string) *CreateTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTaskRequest) SetDescription(v string) *CreateTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateTaskRequest) SetGroupInfo(v *CreateTaskRequestGroupInfo) *CreateTaskRequest {
	s.GroupInfo = v
	return s
}

func (s *CreateTaskRequest) SetInitModuleState(v bool) *CreateTaskRequest {
	s.InitModuleState = &v
	return s
}

func (s *CreateTaskRequest) SetModuleId(v string) *CreateTaskRequest {
	s.ModuleId = &v
	return s
}

func (s *CreateTaskRequest) SetModuleVersion(v string) *CreateTaskRequest {
	s.ModuleVersion = &v
	return s
}

func (s *CreateTaskRequest) SetName(v string) *CreateTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTaskRequest) SetProtectionStrategy(v []*string) *CreateTaskRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *CreateTaskRequest) SetRamRole(v string) *CreateTaskRequest {
	s.RamRole = &v
	return s
}

func (s *CreateTaskRequest) SetSkipPropertyValidation(v bool) *CreateTaskRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *CreateTaskRequest) SetTags(v []*CreateTaskRequestTags) *CreateTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateTaskRequest) SetTaskBackend(v *CreateTaskRequestTaskBackend) *CreateTaskRequest {
	s.TaskBackend = v
	return s
}

func (s *CreateTaskRequest) SetTerraformVersion(v string) *CreateTaskRequest {
	s.TerraformVersion = &v
	return s
}

func (s *CreateTaskRequest) SetTriggerStrategy(v string) *CreateTaskRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	if s.GroupInfo != nil {
		if err := s.GroupInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskBackend != nil {
		if err := s.TaskBackend.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskRequestGroupInfo struct {
	// example:
	//
	// g-5fd38c9b92d541a7083a86432e2
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// project-433aead7560572057e5d9167608
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s CreateTaskRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateTaskRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateTaskRequestGroupInfo) SetGroupId(v string) *CreateTaskRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *CreateTaskRequestGroupInfo) SetProjectId(v string) *CreateTaskRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *CreateTaskRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s CreateTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateTaskRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateTaskRequestTags) SetTagKey(v string) *CreateTaskRequestTags {
	s.TagKey = &v
	return s
}

func (s *CreateTaskRequestTags) SetTagValue(v string) *CreateTaskRequestTags {
	s.TagValue = &v
	return s
}

func (s *CreateTaskRequestTags) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestTaskBackend struct {
	BucketEndpoint *string `json:"bucketEndpoint,omitempty" xml:"bucketEndpoint,omitempty"`
	BucketName     *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	ObjectPath     *string `json:"objectPath,omitempty" xml:"objectPath,omitempty"`
}

func (s CreateTaskRequestTaskBackend) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestTaskBackend) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTaskBackend) GetBucketEndpoint() *string {
	return s.BucketEndpoint
}

func (s *CreateTaskRequestTaskBackend) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateTaskRequestTaskBackend) GetObjectPath() *string {
	return s.ObjectPath
}

func (s *CreateTaskRequestTaskBackend) SetBucketEndpoint(v string) *CreateTaskRequestTaskBackend {
	s.BucketEndpoint = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetBucketName(v string) *CreateTaskRequestTaskBackend {
	s.BucketName = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) SetObjectPath(v string) *CreateTaskRequestTaskBackend {
	s.ObjectPath = &v
	return s
}

func (s *CreateTaskRequestTaskBackend) Validate() error {
	return dara.Validate(s)
}
