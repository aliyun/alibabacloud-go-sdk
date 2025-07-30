// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSynchronizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *RunSynchronizationJobRequest
	GetDescription() *string
	SetInstanceId(v string) *RunSynchronizationJobRequest
	GetInstanceId() *string
	SetPasswordInitialization(v bool) *RunSynchronizationJobRequest
	GetPasswordInitialization() *bool
	SetSynchronizationScopeConfig(v *RunSynchronizationJobRequestSynchronizationScopeConfig) *RunSynchronizationJobRequest
	GetSynchronizationScopeConfig() *RunSynchronizationJobRequestSynchronizationScopeConfig
	SetTargetId(v string) *RunSynchronizationJobRequest
	GetTargetId() *string
	SetTargetType(v string) *RunSynchronizationJobRequest
	GetTargetType() *string
	SetUserIdentityTypes(v []*string) *RunSynchronizationJobRequest
	GetUserIdentityTypes() []*string
}

type RunSynchronizationJobRequest struct {
	// Synchronization task description
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Whether initialize password
	//
	// example:
	//
	// true
	PasswordInitialization *bool `json:"PasswordInitialization,omitempty" xml:"PasswordInitialization,omitempty"`
	// Synchronization scope
	SynchronizationScopeConfig *RunSynchronizationJobRequestSynchronizationScopeConfig `json:"SynchronizationScopeConfig,omitempty" xml:"SynchronizationScopeConfig,omitempty" type:"Struct"`
	// The ID of the synchronization destination.
	//
	// This parameter is required.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3ngxxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the synchronization destination. Valid values:
	//
	// 	- identity_provider
	//
	// 	- application
	//
	// This parameter is required.
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// User identity types
	UserIdentityTypes []*string `json:"UserIdentityTypes,omitempty" xml:"UserIdentityTypes,omitempty" type:"Repeated"`
}

func (s RunSynchronizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobRequest) GetDescription() *string {
	return s.Description
}

func (s *RunSynchronizationJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RunSynchronizationJobRequest) GetPasswordInitialization() *bool {
	return s.PasswordInitialization
}

func (s *RunSynchronizationJobRequest) GetSynchronizationScopeConfig() *RunSynchronizationJobRequestSynchronizationScopeConfig {
	return s.SynchronizationScopeConfig
}

func (s *RunSynchronizationJobRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *RunSynchronizationJobRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *RunSynchronizationJobRequest) GetUserIdentityTypes() []*string {
	return s.UserIdentityTypes
}

func (s *RunSynchronizationJobRequest) SetDescription(v string) *RunSynchronizationJobRequest {
	s.Description = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetInstanceId(v string) *RunSynchronizationJobRequest {
	s.InstanceId = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetPasswordInitialization(v bool) *RunSynchronizationJobRequest {
	s.PasswordInitialization = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetSynchronizationScopeConfig(v *RunSynchronizationJobRequestSynchronizationScopeConfig) *RunSynchronizationJobRequest {
	s.SynchronizationScopeConfig = v
	return s
}

func (s *RunSynchronizationJobRequest) SetTargetId(v string) *RunSynchronizationJobRequest {
	s.TargetId = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetTargetType(v string) *RunSynchronizationJobRequest {
	s.TargetType = &v
	return s
}

func (s *RunSynchronizationJobRequest) SetUserIdentityTypes(v []*string) *RunSynchronizationJobRequest {
	s.UserIdentityTypes = v
	return s
}

func (s *RunSynchronizationJobRequest) Validate() error {
	return dara.Validate(s)
}

type RunSynchronizationJobRequestSynchronizationScopeConfig struct {
	// The group IDs.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The IDs of organizational units.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// UserIds
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s RunSynchronizationJobRequestSynchronizationScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s RunSynchronizationJobRequestSynchronizationScopeConfig) GoString() string {
	return s.String()
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) GetUserIds() []*string {
	return s.UserIds
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) SetGroupIds(v []*string) *RunSynchronizationJobRequestSynchronizationScopeConfig {
	s.GroupIds = v
	return s
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) SetOrganizationalUnitIds(v []*string) *RunSynchronizationJobRequestSynchronizationScopeConfig {
	s.OrganizationalUnitIds = v
	return s
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) SetUserIds(v []*string) *RunSynchronizationJobRequestSynchronizationScopeConfig {
	s.UserIds = v
	return s
}

func (s *RunSynchronizationJobRequestSynchronizationScopeConfig) Validate() error {
	return dara.Validate(s)
}
