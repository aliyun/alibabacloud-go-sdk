// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDeploymentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeployments(v []*GetStackDeploymentsResponseBodyDeployments) *GetStackDeploymentsResponseBody
	GetDeployments() []*GetStackDeploymentsResponseBodyDeployments
	SetRequestId(v string) *GetStackDeploymentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetStackDeploymentsResponseBody
	GetTotalCount() *int32
}

type GetStackDeploymentsResponseBody struct {
	Deployments []*GetStackDeploymentsResponseBodyDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetStackDeploymentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBody) GetDeployments() []*GetStackDeploymentsResponseBodyDeployments {
	return s.Deployments
}

func (s *GetStackDeploymentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackDeploymentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetStackDeploymentsResponseBody) SetDeployments(v []*GetStackDeploymentsResponseBodyDeployments) *GetStackDeploymentsResponseBody {
	s.Deployments = v
	return s
}

func (s *GetStackDeploymentsResponseBody) SetRequestId(v string) *GetStackDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackDeploymentsResponseBody) SetTotalCount(v int32) *GetStackDeploymentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetStackDeploymentsResponseBody) Validate() error {
	if s.Deployments != nil {
		for _, item := range s.Deployments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackDeploymentsResponseBodyDeployments struct {
	Config *GetStackDeploymentsResponseBodyDeploymentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// v1
	ConfigVersion *string `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// example:
	//
	// 2026-04-01T12:10:18Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// production
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// example:
	//
	// 1
	DeploymentNo *string `json:"deploymentNo,omitempty" xml:"deploymentNo,omitempty"`
	// example:
	//
	// v1
	DeploymentVersion *string `json:"deploymentVersion,omitempty" xml:"deploymentVersion,omitempty"`
	// example:
	//
	// 38000
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// Manual
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// example:
	//
	// \\n Error: Invalid value for input variable\\n \\n   on main.tf line 17, in module \\"alb\\":\\n   17:   log_project           = var.log_project.project_name\\n \\n The given value is not suitable for module.alb.var.log_project declared at\\n modules/alb/main.tf:34,1-23: string required.\\n╵\\n
	FailedReason *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	// example:
	//
	// job-as154vldqt46mv0ixxxxx
	JobId       *string                                                  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	Outputs     []*GetStackDeploymentsResponseBodyDeploymentsOutputs     `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	Parameters  []*GetStackDeploymentsResponseBodyDeploymentsParameters  `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	PlanOutputs []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputs `json:"planOutputs,omitempty" xml:"planOutputs,omitempty" type:"Repeated"`
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// task-as1d4vld8ogb2l32xxxxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeployments) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeployments) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetConfig() *GetStackDeploymentsResponseBodyDeploymentsConfig {
	return s.Config
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetDeploymentName() *string {
	return s.DeploymentName
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetDeploymentNo() *string {
	return s.DeploymentNo
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetDeploymentVersion() *string {
	return s.DeploymentVersion
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetExecuteType() *string {
	return s.ExecuteType
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetFailedReason() *string {
	return s.FailedReason
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetJobId() *string {
	return s.JobId
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetOutputs() []*GetStackDeploymentsResponseBodyDeploymentsOutputs {
	return s.Outputs
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetParameters() []*GetStackDeploymentsResponseBodyDeploymentsParameters {
	return s.Parameters
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetPlanOutputs() []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputs {
	return s.PlanOutputs
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetStatus() *string {
	return s.Status
}

func (s *GetStackDeploymentsResponseBodyDeployments) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetConfig(v *GetStackDeploymentsResponseBodyDeploymentsConfig) *GetStackDeploymentsResponseBodyDeployments {
	s.Config = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetConfigVersion(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.ConfigVersion = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetCreateTime(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.CreateTime = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetDeploymentName(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.DeploymentName = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetDeploymentNo(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.DeploymentNo = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetDeploymentVersion(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.DeploymentVersion = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetElapsedTime(v int64) *GetStackDeploymentsResponseBodyDeployments {
	s.ElapsedTime = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetExecuteType(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.ExecuteType = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetFailedReason(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.FailedReason = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetJobId(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.JobId = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetOutputs(v []*GetStackDeploymentsResponseBodyDeploymentsOutputs) *GetStackDeploymentsResponseBodyDeployments {
	s.Outputs = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetParameters(v []*GetStackDeploymentsResponseBodyDeploymentsParameters) *GetStackDeploymentsResponseBodyDeployments {
	s.Parameters = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetPlanOutputs(v []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) *GetStackDeploymentsResponseBodyDeployments {
	s.PlanOutputs = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetStatus(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.Status = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) SetTaskId(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.TaskId = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeployments) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PlanOutputs != nil {
		for _, item := range s.PlanOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackDeploymentsResponseBodyDeploymentsConfig struct {
	// example:
	//
	// false
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// example:
	//
	// false
	IsDestroy *bool `json:"isDestroy,omitempty" xml:"isDestroy,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsConfig) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsConfig) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsConfig) GetAutoApply() *bool {
	return s.AutoApply
}

func (s *GetStackDeploymentsResponseBodyDeploymentsConfig) GetIsDestroy() *bool {
	return s.IsDestroy
}

func (s *GetStackDeploymentsResponseBodyDeploymentsConfig) SetAutoApply(v bool) *GetStackDeploymentsResponseBodyDeploymentsConfig {
	s.AutoApply = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsConfig) SetIsDestroy(v bool) *GetStackDeploymentsResponseBodyDeploymentsConfig {
	s.IsDestroy = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsConfig) Validate() error {
	return dara.Validate(s)
}

type GetStackDeploymentsResponseBodyDeploymentsOutputs struct {
	// example:
	//
	// The name of the SLS log project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// component.sls.project_name
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// log-project-xxxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsOutputs) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) GetDescription() *string {
	return s.Description
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) GetExpression() *string {
	return s.Expression
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) GetName() *string {
	return s.Name
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) GetType() *string {
	return s.Type
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) GetValue() *string {
	return s.Value
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) SetDescription(v string) *GetStackDeploymentsResponseBodyDeploymentsOutputs {
	s.Description = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) SetExpression(v string) *GetStackDeploymentsResponseBodyDeploymentsOutputs {
	s.Expression = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) SetName(v string) *GetStackDeploymentsResponseBodyDeploymentsOutputs {
	s.Name = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) SetType(v string) *GetStackDeploymentsResponseBodyDeploymentsOutputs {
	s.Type = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) SetValue(v string) *GetStackDeploymentsResponseBodyDeploymentsOutputs {
	s.Value = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsOutputs) Validate() error {
	return dara.Validate(s)
}

type GetStackDeploymentsResponseBodyDeploymentsParameters struct {
	// example:
	//
	// cn-hangzhou
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// example:
	//
	// region of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// region
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// ap-southeast-6
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsParameters) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsParameters) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetDescription() *string {
	return s.Description
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetName() *string {
	return s.Name
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetType() *string {
	return s.Type
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetValue() *string {
	return s.Value
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetDefaultValue(v string) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.DefaultValue = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetDescription(v string) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.Description = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetName(v string) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.Name = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetType(v string) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.Type = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetValue(v string) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.Value = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) Validate() error {
	return dara.Validate(s)
}

type GetStackDeploymentsResponseBodyDeploymentsPlanOutputs struct {
	// example:
	//
	// update
	ModuleAction       *string                                                                  `json:"moduleAction,omitempty" xml:"moduleAction,omitempty"`
	ModuleActionDetail *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail `json:"moduleActionDetail,omitempty" xml:"moduleActionDetail,omitempty" type:"Struct"`
	ResourceChanges    []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges  `json:"resourceChanges,omitempty" xml:"resourceChanges,omitempty" type:"Repeated"`
	// example:
	//
	// sls
	StackModuleName *string `json:"stackModuleName,omitempty" xml:"stackModuleName,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) GetModuleAction() *string {
	return s.ModuleAction
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) GetModuleActionDetail() *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail {
	return s.ModuleActionDetail
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) GetResourceChanges() []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges {
	return s.ResourceChanges
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) GetStackModuleName() *string {
	return s.StackModuleName
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) SetModuleAction(v string) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs {
	s.ModuleAction = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) SetModuleActionDetail(v *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs {
	s.ModuleActionDetail = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) SetResourceChanges(v []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs {
	s.ResourceChanges = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) SetStackModuleName(v string) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs {
	s.StackModuleName = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputs) Validate() error {
	if s.ModuleActionDetail != nil {
		if err := s.ModuleActionDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceChanges != nil {
		for _, item := range s.ResourceChanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail struct {
	// example:
	//
	// 0
	Add *int32 `json:"add,omitempty" xml:"add,omitempty"`
	// example:
	//
	// 1
	Change *int32 `json:"change,omitempty" xml:"change,omitempty"`
	// example:
	//
	// 0
	Destroy *int32 `json:"destroy,omitempty" xml:"destroy,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) GetAdd() *int32 {
	return s.Add
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) GetChange() *int32 {
	return s.Change
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) GetDestroy() *int32 {
	return s.Destroy
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) SetAdd(v int32) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail {
	s.Add = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) SetChange(v int32) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail {
	s.Change = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) SetDestroy(v int32) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail {
	s.Destroy = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail) Validate() error {
	return dara.Validate(s)
}

type GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges struct {
	// example:
	//
	// ~ resource \\"alicloud_log_store\\" \\"default\\" {\\n        id                    = \\"alb-log-project-v1-ph-xxxxx:alb-log-store-ph\\"\\n      ~ max_split_shard_count = 64 -> 32\\n        name                  = \\"alb-log-store-ph\\"\\n\\n        # (13 unchanged attributes hidden)\\n    }
	Change          *string   `json:"change,omitempty" xml:"change,omitempty"`
	ResourceActions []*string `json:"resourceActions,omitempty" xml:"resourceActions,omitempty" type:"Repeated"`
	// example:
	//
	// alicloud_log_store.default
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty" xml:"resourceIdentifier,omitempty"`
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) String() string {
	return dara.Prettify(s)
}

func (s GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) GoString() string {
	return s.String()
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) GetChange() *string {
	return s.Change
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) GetResourceActions() []*string {
	return s.ResourceActions
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) SetChange(v string) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges {
	s.Change = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) SetResourceActions(v []*string) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges {
	s.ResourceActions = v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) SetResourceIdentifier(v string) *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges {
	s.ResourceIdentifier = &v
	return s
}

func (s *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges) Validate() error {
	return dara.Validate(s)
}
