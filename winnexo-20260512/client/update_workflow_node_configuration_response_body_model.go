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
	// 当前用户是否可编辑此节点
	//
	// example:
	//
	// true
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// 公共业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 配置版本；更新或恢复时作为expectedVersion
	//
	// example:
	//
	// 3
	ConfigVersion *int64 `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// 模板默认执行主体
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	DefaultExecutor *UpdateWorkflowNodeConfigurationResponseBodyDefaultExecutor `json:"defaultExecutor,omitempty" xml:"defaultExecutor,omitempty" type:"Struct"`
	// 模板默认模型运行档
	//
	// example:
	//
	// qwen-max
	DefaultModel *string `json:"defaultModel,omitempty" xml:"defaultModel,omitempty"`
	// 模板默认负责人用户ID列表
	//
	// example:
	//
	// [10001]
	DefaultResponsibleUserIds []*string `json:"defaultResponsibleUserIds,omitempty" xml:"defaultResponsibleUserIds,omitempty" type:"Repeated"`
	// 当前生效的执行主体
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	EffectiveExecutor *UpdateWorkflowNodeConfigurationResponseBodyEffectiveExecutor `json:"effectiveExecutor,omitempty" xml:"effectiveExecutor,omitempty" type:"Struct"`
	// 当前生效的模型运行档
	//
	// example:
	//
	// qwen-max-plus
	EffectiveModel *string `json:"effectiveModel,omitempty" xml:"effectiveModel,omitempty"`
	// 当前生效的负责人用户ID列表
	//
	// example:
	//
	// [10001]
	EffectiveResponsibleUserIds []*string `json:"effectiveResponsibleUserIds,omitempty" xml:"effectiveResponsibleUserIds,omitempty" type:"Repeated"`
	// 当前物理计划ID
	//
	// example:
	//
	// inst_2081
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// 公共错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务节点key
	//
	// example:
	//
	// node_7f3a2c
	NodeId *string `json:"nodeId,omitempty" xml:"nodeId,omitempty"`
	// 实例覆盖值；字段为null时使用模板默认值
	//
	// example:
	//
	// {"model":"qwen-max-plus"}
	Overrides *UpdateWorkflowNodeConfigurationResponseBodyOverrides `json:"overrides,omitempty" xml:"overrides,omitempty" type:"Struct"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 状态刷新游标，不能作为配置版本
	//
	// example:
	//
	// 8
	StateVersion *int64 `json:"stateVersion,omitempty" xml:"stateVersion,omitempty"`
	// 最后配置更新时间ISO8601；从未更新时为空
	//
	// example:
	//
	// 2026-11-05T08:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// 最后更新配置的平台用户ID；从未更新时为空
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
	// USER 为正整数平台用户ID字符串；AGENT 为空间数字员工技术名称
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// USER 真人、AGENT 数字员工或历史 SYSTEM 系统执行人；SYSTEM仅用于响应
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
	// USER 为正整数平台用户ID字符串；AGENT 为空间数字员工技术名称
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// USER 真人、AGENT 数字员工或历史 SYSTEM 系统执行人；SYSTEM仅用于响应
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
	// 覆盖执行主体；null为默认
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	Executor *UpdateWorkflowNodeConfigurationResponseBodyOverridesExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
	// 覆盖模型运行档；null为默认
	//
	// example:
	//
	// qwen-max
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// 覆盖负责人ID列表；null为默认
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
	// USER 为正整数平台用户ID字符串；AGENT 为空间数字员工技术名称
	//
	// example:
	//
	// agent_1001
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// USER 真人、AGENT 数字员工或历史 SYSTEM 系统执行人；SYSTEM仅用于响应
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
