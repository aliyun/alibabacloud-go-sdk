// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowNodeConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanEdit(v bool) *GetWorkflowNodeConfigurationResponseBody
	GetCanEdit() *bool
	SetCode(v string) *GetWorkflowNodeConfigurationResponseBody
	GetCode() *string
	SetConfigVersion(v int64) *GetWorkflowNodeConfigurationResponseBody
	GetConfigVersion() *int64
	SetDefaultExecutor(v *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) *GetWorkflowNodeConfigurationResponseBody
	GetDefaultExecutor() *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor
	SetDefaultModel(v string) *GetWorkflowNodeConfigurationResponseBody
	GetDefaultModel() *string
	SetDefaultResponsibleUserIds(v []*string) *GetWorkflowNodeConfigurationResponseBody
	GetDefaultResponsibleUserIds() []*string
	SetEffectiveExecutor(v *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *GetWorkflowNodeConfigurationResponseBody
	GetEffectiveExecutor() *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor
	SetEffectiveModel(v string) *GetWorkflowNodeConfigurationResponseBody
	GetEffectiveModel() *string
	SetEffectiveResponsibleUserIds(v []*string) *GetWorkflowNodeConfigurationResponseBody
	GetEffectiveResponsibleUserIds() []*string
	SetInstanceId(v string) *GetWorkflowNodeConfigurationResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetWorkflowNodeConfigurationResponseBody
	GetMessage() *string
	SetNodeId(v string) *GetWorkflowNodeConfigurationResponseBody
	GetNodeId() *string
	SetOverrides(v *GetWorkflowNodeConfigurationResponseBodyOverrides) *GetWorkflowNodeConfigurationResponseBody
	GetOverrides() *GetWorkflowNodeConfigurationResponseBodyOverrides
	SetRequestId(v string) *GetWorkflowNodeConfigurationResponseBody
	GetRequestId() *string
	SetStateVersion(v int64) *GetWorkflowNodeConfigurationResponseBody
	GetStateVersion() *int64
	SetUpdatedAt(v string) *GetWorkflowNodeConfigurationResponseBody
	GetUpdatedAt() *string
	SetUpdatedBy(v int64) *GetWorkflowNodeConfigurationResponseBody
	GetUpdatedBy() *int64
}

type GetWorkflowNodeConfigurationResponseBody struct {
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
	// The configuration version. It is used as the expectedVersion when updating or restoring the configuration.
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
	DefaultExecutor *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor `json:"defaultExecutor,omitempty" xml:"defaultExecutor,omitempty" type:"Struct"`
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
	EffectiveExecutor *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor `json:"effectiveExecutor,omitempty" xml:"effectiveExecutor,omitempty" type:"Struct"`
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
	// The override values for the instance. If a field is null, the default value of the template is used.
	//
	// example:
	//
	// {"model":"qwen-max-plus"}
	Overrides *GetWorkflowNodeConfigurationResponseBodyOverrides `json:"overrides,omitempty" xml:"overrides,omitempty" type:"Struct"`
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

func (s GetWorkflowNodeConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetConfigVersion() *int64 {
	return s.ConfigVersion
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetDefaultExecutor() *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	return s.DefaultExecutor
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetDefaultModel() *string {
	return s.DefaultModel
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetDefaultResponsibleUserIds() []*string {
	return s.DefaultResponsibleUserIds
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetEffectiveExecutor() *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	return s.EffectiveExecutor
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetEffectiveModel() *string {
	return s.EffectiveModel
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetEffectiveResponsibleUserIds() []*string {
	return s.EffectiveResponsibleUserIds
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetOverrides() *GetWorkflowNodeConfigurationResponseBodyOverrides {
	return s.Overrides
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetStateVersion() *int64 {
	return s.StateVersion
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetWorkflowNodeConfigurationResponseBody) GetUpdatedBy() *int64 {
	return s.UpdatedBy
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetCanEdit(v bool) *GetWorkflowNodeConfigurationResponseBody {
	s.CanEdit = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetCode(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetConfigVersion(v int64) *GetWorkflowNodeConfigurationResponseBody {
	s.ConfigVersion = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetDefaultExecutor(v *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) *GetWorkflowNodeConfigurationResponseBody {
	s.DefaultExecutor = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetDefaultModel(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.DefaultModel = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetDefaultResponsibleUserIds(v []*string) *GetWorkflowNodeConfigurationResponseBody {
	s.DefaultResponsibleUserIds = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetEffectiveExecutor(v *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) *GetWorkflowNodeConfigurationResponseBody {
	s.EffectiveExecutor = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetEffectiveModel(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.EffectiveModel = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetEffectiveResponsibleUserIds(v []*string) *GetWorkflowNodeConfigurationResponseBody {
	s.EffectiveResponsibleUserIds = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetInstanceId(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetMessage(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetNodeId(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.NodeId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetOverrides(v *GetWorkflowNodeConfigurationResponseBodyOverrides) *GetWorkflowNodeConfigurationResponseBody {
	s.Overrides = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetRequestId(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetStateVersion(v int64) *GetWorkflowNodeConfigurationResponseBody {
	s.StateVersion = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetUpdatedAt(v string) *GetWorkflowNodeConfigurationResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) SetUpdatedBy(v int64) *GetWorkflowNodeConfigurationResponseBody {
	s.UpdatedBy = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBody) Validate() error {
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

type GetWorkflowNodeConfigurationResponseBodyDefaultExecutor struct {
	// The ID of the executor. If the type is USER, this is a positive integer string representing the platform user ID. If the type is AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the executor. Valid values: USER for a real person, AGENT for a digital employee, and SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetId() *string {
	return s.Id
}

func (s *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) GetType() *string {
	return s.Type
}

func (s *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetId(v string) *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Id = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) SetType(v string) *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor {
	s.Type = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyDefaultExecutor) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor struct {
	// The ID of the executor. If the type is USER, this is a positive integer string representing the platform user ID. If the type is AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the executor. Valid values: USER for a real person, AGENT for a digital employee, and SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetId() *string {
	return s.Id
}

func (s *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) GetType() *string {
	return s.Type
}

func (s *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetId(v string) *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Id = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) SetType(v string) *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor {
	s.Type = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyEffectiveExecutor) Validate() error {
	return dara.Validate(s)
}

type GetWorkflowNodeConfigurationResponseBodyOverrides struct {
	// The override executor. A null value indicates that the default is used.
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	Executor *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// The override model runtime profile. A null value indicates that the default is used.
	//
	// example:
	//
	// qwen-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The list of override responsible user IDs. A null value indicates that the default is used.
	//
	// example:
	//
	// [10001]
	ResponsibleUserIds []*int64 `json:"responsibleUserIds,omitempty" xml:"responsibleUserIds,omitempty" type:"Repeated"`
}

func (s GetWorkflowNodeConfigurationResponseBodyOverrides) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponseBodyOverrides) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) GetExecutor() *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	return s.Executor
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) GetModel() *string {
	return s.Model
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) GetResponsibleUserIds() []*int64 {
	return s.ResponsibleUserIds
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) SetExecutor(v *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) *GetWorkflowNodeConfigurationResponseBodyOverrides {
	s.Executor = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) SetModel(v string) *GetWorkflowNodeConfigurationResponseBodyOverrides {
	s.Model = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) SetResponsibleUserIds(v []*int64) *GetWorkflowNodeConfigurationResponseBodyOverrides {
	s.ResponsibleUserIds = v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverrides) Validate() error {
	if s.Executor != nil {
		if err := s.Executor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkflowNodeConfigurationResponseBodyOverridesExecutor struct {
	// The ID of the executor. If the type is USER, this is a positive integer string representing the platform user ID. If the type is AGENT, this is the technical name of the digital employee in the workspace.
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type of the executor. Valid values: USER for a real person, AGENT for a digital employee, and SYSTEM for a legacy system executor. The SYSTEM value is used only in responses.
	//
	// example:
	//
	// AGENT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GoString() string {
	return s.String()
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetId() *string {
	return s.Id
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) GetType() *string {
	return s.Type
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetId(v string) *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Id = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) SetType(v string) *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor {
	s.Type = &v
	return s
}

func (s *GetWorkflowNodeConfigurationResponseBodyOverridesExecutor) Validate() error {
	return dara.Validate(s)
}
