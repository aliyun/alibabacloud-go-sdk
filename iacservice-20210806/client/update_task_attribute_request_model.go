// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoApply(v bool) *UpdateTaskAttributeRequest
	GetAutoApply() *bool
	SetAutoDestroy(v bool) *UpdateTaskAttributeRequest
	GetAutoDestroy() *bool
	SetClientToken(v string) *UpdateTaskAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTaskAttributeRequest
	GetDescription() *string
	SetGroupInfo(v *UpdateTaskAttributeRequestGroupInfo) *UpdateTaskAttributeRequest
	GetGroupInfo() *UpdateTaskAttributeRequestGroupInfo
	SetInitModuleState(v bool) *UpdateTaskAttributeRequest
	GetInitModuleState() *bool
	SetModuleVersion(v string) *UpdateTaskAttributeRequest
	GetModuleVersion() *string
	SetName(v string) *UpdateTaskAttributeRequest
	GetName() *string
	SetProtectionStrategy(v []*string) *UpdateTaskAttributeRequest
	GetProtectionStrategy() []*string
	SetRamRole(v string) *UpdateTaskAttributeRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *UpdateTaskAttributeRequest
	GetSkipPropertyValidation() *bool
	SetTags(v []*UpdateTaskAttributeRequestTags) *UpdateTaskAttributeRequest
	GetTags() []*UpdateTaskAttributeRequestTags
	SetTerraformVersion(v string) *UpdateTaskAttributeRequest
	GetTerraformVersion() *string
	SetTriggerStrategy(v string) *UpdateTaskAttributeRequest
	GetTriggerStrategy() *string
}

type UpdateTaskAttributeRequest struct {
	// example:
	//
	// true
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// true
	AutoDestroy *bool `json:"autoDestroy,omitempty" xml:"autoDestroy,omitempty"`
	// This parameter is required.
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// demo
	Description     *string                              `json:"description,omitempty" xml:"description,omitempty"`
	GroupInfo       *UpdateTaskAttributeRequestGroupInfo `json:"groupInfo,omitempty" xml:"groupInfo,omitempty" type:"Struct"`
	InitModuleState *bool                                `json:"initModuleState,omitempty" xml:"initModuleState,omitempty"`
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// test
	Name               *string   `json:"name,omitempty" xml:"name,omitempty"`
	ProtectionStrategy []*string `json:"protectionStrategy,omitempty" xml:"protectionStrategy,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	RamRole                *string                           `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	SkipPropertyValidation *bool                             `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	Tags                   []*UpdateTaskAttributeRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
}

func (s UpdateTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequest) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *UpdateTaskAttributeRequest) GetAutoDestroy() *bool {
	return s.AutoDestroy
}

func (s *UpdateTaskAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTaskAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTaskAttributeRequest) GetGroupInfo() *UpdateTaskAttributeRequestGroupInfo {
	return s.GroupInfo
}

func (s *UpdateTaskAttributeRequest) GetInitModuleState() *bool {
	return s.InitModuleState
}

func (s *UpdateTaskAttributeRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *UpdateTaskAttributeRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTaskAttributeRequest) GetProtectionStrategy() []*string {
	return s.ProtectionStrategy
}

func (s *UpdateTaskAttributeRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateTaskAttributeRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *UpdateTaskAttributeRequest) GetTags() []*UpdateTaskAttributeRequestTags {
	return s.Tags
}

func (s *UpdateTaskAttributeRequest) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *UpdateTaskAttributeRequest) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *UpdateTaskAttributeRequest) SetAutoApply(v bool) *UpdateTaskAttributeRequest {
	s.AutoApply = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetAutoDestroy(v bool) *UpdateTaskAttributeRequest {
	s.AutoDestroy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetClientToken(v string) *UpdateTaskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetDescription(v string) *UpdateTaskAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetGroupInfo(v *UpdateTaskAttributeRequestGroupInfo) *UpdateTaskAttributeRequest {
	s.GroupInfo = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetInitModuleState(v bool) *UpdateTaskAttributeRequest {
	s.InitModuleState = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetModuleVersion(v string) *UpdateTaskAttributeRequest {
	s.ModuleVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetName(v string) *UpdateTaskAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetProtectionStrategy(v []*string) *UpdateTaskAttributeRequest {
	s.ProtectionStrategy = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetRamRole(v string) *UpdateTaskAttributeRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetSkipPropertyValidation(v bool) *UpdateTaskAttributeRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTags(v []*UpdateTaskAttributeRequestTags) *UpdateTaskAttributeRequest {
	s.Tags = v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTerraformVersion(v string) *UpdateTaskAttributeRequest {
	s.TerraformVersion = &v
	return s
}

func (s *UpdateTaskAttributeRequest) SetTriggerStrategy(v string) *UpdateTaskAttributeRequest {
	s.TriggerStrategy = &v
	return s
}

func (s *UpdateTaskAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAttributeRequestGroupInfo struct {
	// example:
	//
	// g-433aead7560571e66e31274ffd3
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// p-433aead75605713865c386cb9d
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s UpdateTaskAttributeRequestGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequestGroupInfo) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequestGroupInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateTaskAttributeRequestGroupInfo) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetGroupId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.GroupId = &v
	return s
}

func (s *UpdateTaskAttributeRequestGroupInfo) SetProjectId(v string) *UpdateTaskAttributeRequestGroupInfo {
	s.ProjectId = &v
	return s
}

func (s *UpdateTaskAttributeRequestGroupInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskAttributeRequestTags struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s UpdateTaskAttributeRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskAttributeRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateTaskAttributeRequestTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateTaskAttributeRequestTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateTaskAttributeRequestTags) SetTagKey(v string) *UpdateTaskAttributeRequestTags {
	s.TagKey = &v
	return s
}

func (s *UpdateTaskAttributeRequestTags) SetTagValue(v string) *UpdateTaskAttributeRequestTags {
	s.TagValue = &v
	return s
}

func (s *UpdateTaskAttributeRequestTags) Validate() error {
	return dara.Validate(s)
}
