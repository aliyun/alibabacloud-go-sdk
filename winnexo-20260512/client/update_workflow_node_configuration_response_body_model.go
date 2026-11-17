// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowNodeConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanEdit(v bool) *UpdateWorkflowNodeConfigurationResponseBody
	GetCanEdit() *bool
	SetCode(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetCode() *string
	SetConfigVersion(v int64) *UpdateWorkflowNodeConfigurationResponseBody
	GetConfigVersion() *int64
	SetDefaultExecutor(v *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) *UpdateWorkflowNodeConfigurationResponseBody
	GetDefaultExecutor() *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor
	SetDefaultModel(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetDefaultModel() *string
	SetDefaultResponsibleUserIds(v []*string) *UpdateWorkflowNodeConfigurationResponseBody
	GetDefaultResponsibleUserIds() []*string
	SetEffectiveExecutor(v *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *UpdateWorkflowNodeConfigurationResponseBody
	GetEffectiveExecutor() *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor
	SetEffectiveModel(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetEffectiveModel() *string
	SetEffectiveResponsibleUserIds(v []*string) *UpdateWorkflowNodeConfigurationResponseBody
	GetEffectiveResponsibleUserIds() []*string
	SetInstanceId(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetInstanceId() *string
	SetMessage(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetMessage() *string
	SetNodeId(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetNodeId() *string
	SetOverrides(v *UpdateWorkflowNodeConfigurationResponseBodyOverrides) *UpdateWorkflowNodeConfigurationResponseBody
	GetOverrides() *UpdateWorkflowNodeConfigurationResponseBodyOverrides
	SetRequestId(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetRequestId() *string
	SetStateVersion(v int64) *UpdateWorkflowNodeConfigurationResponseBody
	GetStateVersion() *int64
	SetUpdatedAt(v string) *UpdateWorkflowNodeConfigurationResponseBody
	GetUpdatedAt() *string
	SetUpdatedBy(v int64) *UpdateWorkflowNodeConfigurationResponseBody
	GetUpdatedBy() *int64
}

type UpdateWorkflowNodeConfigurationResponseBody struct {
	// Indicates whether the current user can edit this node.
	//
	// example:
	//
	// true
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// The common business status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The configuration version. Use this value as the expectedVersion when updating or restoring.
	//
	// example:
	//
	// 3
	ConfigVersion *int64 `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// The default executor defined in the template.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	DefaultExecutor *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor `json:"defaultExecutor,omitempty" xml:"defaultExecutor,omitempty" type:"Struct"`
	// The default model profile defined in the template.
	//
	// example:
	//
	// qwen-max
	DefaultModel *string `json:"defaultModel,omitempty" xml:"defaultModel,omitempty"`
	// The list of user IDs for the default owners defined in the template.
	//
	// example:
	//
	// [10001]
	DefaultResponsibleUserIds []*string `json:"defaultResponsibleUserIds,omitempty" xml:"defaultResponsibleUserIds,omitempty" type:"Repeated"`
	// The currently effective executor.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	EffectiveExecutor *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor `json:"effectiveExecutor,omitempty" xml:"effectiveExecutor,omitempty" type:"Struct"`
	// The currently effective model profile.
	//
	// example:
	//
	// qwen-max-plus
	EffectiveModel *string `json:"effectiveModel,omitempty" xml:"effectiveModel,omitempty"`
	// The list of user IDs for the currently effective owners.
	//
	// example:
	//
	// [10001]
	EffectiveResponsibleUserIds []*string `json:"effectiveResponsibleUserIds,omitempty" xml:"effectiveResponsibleUserIds,omitempty" type:"Repeated"`
	// The ID of the current physical plan.
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The common error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The key of the business node.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The instance override values. If a field is null, the default value defined in the template is used.
	//
	// example:
	//
	// {"model":"qwen-max-plus"}
	Overrides *UpdateWorkflowNodeConfigurationResponseBodyOverrides `json:"overrides,omitempty" xml:"overrides,omitempty" type:"Struct"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The state refresh cursor. This value cannot be used as the configuration version.
	//
	// example:
	//
	// 8
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
	// The time when the configuration was last updated, in ISO 8601 format. This field is empty if the configuration has never been updated.
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The platform user ID of the user who last updated the configuration. This field is empty if the configuration has never been updated.
	//
	// example:
	//
	// 10001
	UpdatedBy *int64 `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetConfigVersion() *int64 {
	return s.ConfigVersion
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetDefaultExecutor() *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	return s.DefaultExecutor
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetDefaultModel() *string {
	return s.DefaultModel
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetDefaultResponsibleUserIds() []*string {
	return s.DefaultResponsibleUserIds
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetEffectiveExecutor() *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	return s.EffectiveExecutor
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetEffectiveModel() *string {
	return s.EffectiveModel
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetEffectiveResponsibleUserIds() []*string {
	return s.EffectiveResponsibleUserIds
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetOverrides() *UpdateWorkflowNodeConfigurationResponseBodyOverrides {
	return s.Overrides
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) GetUpdatedBy() *int64 {
	return s.UpdatedBy
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetCanEdit(v bool) *UpdateWorkflowNodeConfigurationResponseBody {
	s.CanEdit = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetCode(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetConfigVersion(v int64) *UpdateWorkflowNodeConfigurationResponseBody {
	s.ConfigVersion = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetDefaultExecutor(v *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) *UpdateWorkflowNodeConfigurationResponseBody {
	s.DefaultExecutor = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetDefaultModel(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.DefaultModel = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetDefaultResponsibleUserIds(v []*string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.DefaultResponsibleUserIds = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetEffectiveExecutor(v *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *UpdateWorkflowNodeConfigurationResponseBody {
	s.EffectiveExecutor = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetEffectiveModel(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.EffectiveModel = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetEffectiveResponsibleUserIds(v []*string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.EffectiveResponsibleUserIds = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetInstanceId(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.InstanceId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetMessage(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetNodeId(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.NodeId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetOverrides(v *UpdateWorkflowNodeConfigurationResponseBodyOverrides) *UpdateWorkflowNodeConfigurationResponseBody {
	s.Overrides = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetRequestId(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetStateVersion(v int64) *UpdateWorkflowNodeConfigurationResponseBody {
	s.StateVersion = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetUpdatedAt(v string) *UpdateWorkflowNodeConfigurationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) SetUpdatedBy(v int64) *UpdateWorkflowNodeConfigurationResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBody) Validate() error {
	if s.DefaultExecutor != nil {
		if err := s.DefaultExecutor.Validate(); err != nil {
			return err
		}
	}
	if s.EffectiveExecutor != nil {
		if err := s.EffectiveExecutor.Validate(); err != nil {
			return err
		}
	}
	if s.Overrides != nil {
		if err := s.Overrides.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor struct {
	// The executor ID. If the type is USER, the value is a string representing a positive integer platform user ID. If the type is AGENT, the value is the technical name of the digital worker in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER (a human user), AGENT (a digital worker), and SYSTEM (a legacy system executor). The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetId() *string {
	return s.Id
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetId(v string) *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetType(v string) *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor struct {
	// The executor ID. If the type is USER, the value is a string representing a positive integer platform user ID. If the type is AGENT, the value is the technical name of the digital worker in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER (a human user), AGENT (a digital worker), and SYSTEM (a legacy system executor). The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetId() *string {
	return s.Id
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetId(v string) *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetType(v string) *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor) Validate() error {
	return dara.Validate(s)
}

type UpdateWorkflowNodeConfigurationResponseBodyOverrides struct {
	// The overridden executor. A value of null indicates that the default value is used.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	Executor *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// The overridden model profile. A value of null indicates that the default value is used.
	//
	// example:
	//
	// qwen-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The list of overridden owner IDs. A value of null indicates that the default value is used.
	//
	// example:
	//
	// [10001]
	ResponsibleUserIds []*int64 `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty" type:"Repeated"`
}

func (s UpdateWorkflowNodeConfigurationResponseBodyOverrides) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponseBodyOverrides) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) GetExecutor() *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	return s.Executor
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) GetModel() *string {
	return s.Model
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) GetResponsibleUserIds() []*int64 {
	return s.ResponsibleUserIds
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) SetExecutor(v *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) *UpdateWorkflowNodeConfigurationResponseBodyOverrides {
	s.Executor = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) SetModel(v string) *UpdateWorkflowNodeConfigurationResponseBodyOverrides {
	s.Model = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) SetResponsibleUserIds(v []*int64) *UpdateWorkflowNodeConfigurationResponseBodyOverrides {
	s.ResponsibleUserIds = v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverrides) Validate() error {
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor struct {
	// The executor ID. If the type is USER, the value is a string representing a positive integer platform user ID. If the type is AGENT, the value is the technical name of the digital worker in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER (a human user), AGENT (a digital worker), and SYSTEM (a legacy system executor). The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetId() *string {
	return s.Id
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetType() *string {
	return s.Type
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetId(v string) *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetType(v string) *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Type = &v
	return s
}

func (s *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor) Validate() error {
	return dara.Validate(s)
}
