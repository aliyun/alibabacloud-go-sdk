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
	DefaultExecutor *ResetWorkflowNodeConfigurationResponseBodyDefaultExecutor `json:"defaultExecutor,omitempty" xml:"defaultExecutor,omitempty" type:"Struct"`
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
	EffectiveExecutor *ResetWorkflowNodeConfigurationResponseBodyEffectiveExecutor `json:"effectiveExecutor,omitempty" xml:"effectiveExecutor,omitempty" type:"Struct"`
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
	Overrides *ResetWorkflowNodeConfigurationResponseBodyOverrides `json:"overrides,omitempty" xml:"overrides,omitempty" type:"Struct"`
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
	// 覆盖执行主体；null为默认
	//
	// example:
	//
	// {"type":"AGENT","id":"agent_1001"}
	Executor *ResetWorkflowNodeConfigurationResponseBodyOverridesExecutor `json:"executor,omitempty" xml:"executor,omitempty" type:"Struct"`
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
