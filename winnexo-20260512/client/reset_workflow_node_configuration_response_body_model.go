// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetWorkflowNodeConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanEdit(v bool) *ResetWorkflowNodeConfigurationResponseBody
	GetCanEdit() *bool
	SetCode(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetCode() *string
	SetConfigVersion(v int64) *ResetWorkflowNodeConfigurationResponseBody
	GetConfigVersion() *int64
	SetDefaultExecutor(v *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) *ResetWorkflowNodeConfigurationResponseBody
	GetDefaultExecutor() *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor
	SetDefaultModel(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetDefaultModel() *string
	SetDefaultResponsibleUserIds(v []*string) *ResetWorkflowNodeConfigurationResponseBody
	GetDefaultResponsibleUserIds() []*string
	SetEffectiveExecutor(v *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *ResetWorkflowNodeConfigurationResponseBody
	GetEffectiveExecutor() *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor
	SetEffectiveModel(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetEffectiveModel() *string
	SetEffectiveResponsibleUserIds(v []*string) *ResetWorkflowNodeConfigurationResponseBody
	GetEffectiveResponsibleUserIds() []*string
	SetInstanceId(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetInstanceId() *string
	SetMessage(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetMessage() *string
	SetNodeId(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetNodeId() *string
	SetOverrides(v *ResetWorkflowNodeConfigurationResponseBodyOverrides) *ResetWorkflowNodeConfigurationResponseBody
	GetOverrides() *ResetWorkflowNodeConfigurationResponseBodyOverrides
	SetRequestId(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetRequestId() *string
	SetStateVersion(v int64) *ResetWorkflowNodeConfigurationResponseBody
	GetStateVersion() *int64
	SetUpdatedAt(v string) *ResetWorkflowNodeConfigurationResponseBody
	GetUpdatedAt() *string
	SetUpdatedBy(v int64) *ResetWorkflowNodeConfigurationResponseBody
	GetUpdatedBy() *int64
}

type ResetWorkflowNodeConfigurationResponseBody struct {
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
	// The default executor of the template.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	DefaultExecutor *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor `json:"defaultExecutor,omitempty" xml:"defaultExecutor,omitempty" type:"Struct"`
	// The default model runtime profile of the template.
	//
	// example:
	//
	// qwen-max
	DefaultModel *string `json:"defaultModel,omitempty" xml:"defaultModel,omitempty"`
	// The list of user IDs for the default responsible users of the template.
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
	EffectiveExecutor *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor `json:"effectiveExecutor,omitempty" xml:"effectiveExecutor,omitempty" type:"Struct"`
	// The currently effective model runtime profile.
	//
	// example:
	//
	// qwen-max-plus
	EffectiveModel *string `json:"effectiveModel,omitempty" xml:"effectiveModel,omitempty"`
	// The list of user IDs for the currently effective responsible users.
	//
	// example:
	//
	// [10001]
	EffectiveResponsibleUserIds []*string `json:"effectiveResponsibleUserIds,omitempty" xml:"effectiveResponsibleUserIds,omitempty" type:"Repeated"`
	// The current physical plan ID.
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
	// The business node key.
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// The instance override values. If a field is null, the default value from the template is used.
	//
	// example:
	//
	// {"model":"qwen-max-plus"}
	Overrides *ResetWorkflowNodeConfigurationResponseBodyOverrides `json:"overrides,omitempty" xml:"overrides,omitempty" type:"Struct"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The state refresh cursor. This cannot be used as a configuration version.
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

func (s ResetWorkflowNodeConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetConfigVersion() *int64 {
	return s.ConfigVersion
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetDefaultExecutor() *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	return s.DefaultExecutor
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetDefaultModel() *string {
	return s.DefaultModel
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetDefaultResponsibleUserIds() []*string {
	return s.DefaultResponsibleUserIds
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetEffectiveExecutor() *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	return s.EffectiveExecutor
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetEffectiveModel() *string {
	return s.EffectiveModel
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetEffectiveResponsibleUserIds() []*string {
	return s.EffectiveResponsibleUserIds
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetOverrides() *ResetWorkflowNodeConfigurationResponseBodyOverrides {
	return s.Overrides
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ResetWorkflowNodeConfigurationResponseBody) GetUpdatedBy() *int64 {
	return s.UpdatedBy
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetCanEdit(v bool) *ResetWorkflowNodeConfigurationResponseBody {
	s.CanEdit = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetCode(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetConfigVersion(v int64) *ResetWorkflowNodeConfigurationResponseBody {
	s.ConfigVersion = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetDefaultExecutor(v *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) *ResetWorkflowNodeConfigurationResponseBody {
	s.DefaultExecutor = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetDefaultModel(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.DefaultModel = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetDefaultResponsibleUserIds(v []*string) *ResetWorkflowNodeConfigurationResponseBody {
	s.DefaultResponsibleUserIds = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetEffectiveExecutor(v *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *ResetWorkflowNodeConfigurationResponseBody {
	s.EffectiveExecutor = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetEffectiveModel(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.EffectiveModel = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetEffectiveResponsibleUserIds(v []*string) *ResetWorkflowNodeConfigurationResponseBody {
	s.EffectiveResponsibleUserIds = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetInstanceId(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetMessage(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetNodeId(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.NodeId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetOverrides(v *ResetWorkflowNodeConfigurationResponseBodyOverrides) *ResetWorkflowNodeConfigurationResponseBody {
	s.Overrides = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetRequestId(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetStateVersion(v int64) *ResetWorkflowNodeConfigurationResponseBody {
	s.StateVersion = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetUpdatedAt(v string) *ResetWorkflowNodeConfigurationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) SetUpdatedBy(v int64) *ResetWorkflowNodeConfigurationResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBody) Validate() error {
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

type ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor struct {
	// The executor ID. For USER, this is a positive integer string representing the platform user ID. For AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER for a real user, AGENT for a digital employee, or SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetId() *string {
	return s.Id
}

func (s *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetType() *string {
	return s.Type
}

func (s *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetId(v string) *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Id = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetType(v string) *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Type = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor) Validate() error {
	return dara.Validate(s)
}

type ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor struct {
	// The executor ID. For USER, this is a positive integer string representing the platform user ID. For AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER for a real user, AGENT for a digital employee, or SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetId() *string {
	return s.Id
}

func (s *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetType() *string {
	return s.Type
}

func (s *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetId(v string) *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Id = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetType(v string) *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Type = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) Validate() error {
	return dara.Validate(s)
}

type ResetWorkflowNodeConfigurationResponseBodyOverrides struct {
	// The overridden executor. A null value indicates the default.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	Executor *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// The overridden model runtime profile. A null value indicates the default.
	//
	// example:
	//
	// qwen-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The list of overridden responsible user IDs. A null value indicates the default.
	//
	// example:
	//
	// [10001]
	ResponsibleUserIds []*int64 `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty" type:"Repeated"`
}

func (s ResetWorkflowNodeConfigurationResponseBodyOverrides) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponseBodyOverrides) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) GetExecutor() *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	return s.Executor
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) GetModel() *string {
	return s.Model
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) GetResponsibleUserIds() []*int64 {
	return s.ResponsibleUserIds
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) SetExecutor(v *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) *ResetWorkflowNodeConfigurationResponseBodyOverrides {
	s.Executor = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) SetModel(v string) *ResetWorkflowNodeConfigurationResponseBodyOverrides {
	s.Model = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) SetResponsibleUserIds(v []*int64) *ResetWorkflowNodeConfigurationResponseBodyOverrides {
	s.ResponsibleUserIds = v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverrides) Validate() error {
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor struct {
	// The executor ID. For USER, this is a positive integer string representing the platform user ID. For AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The executor type. Valid values: USER for a real user, AGENT for a digital employee, or SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) String() string {
	return dara.Prettify(s)
}

func (s ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GoString() string {
	return s.String()
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetId() *string {
	return s.Id
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetType() *string {
	return s.Type
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetId(v string) *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Id = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetType(v string) *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Type = &v
	return s
}

func (s *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor) Validate() error {
	return dara.Validate(s)
}
