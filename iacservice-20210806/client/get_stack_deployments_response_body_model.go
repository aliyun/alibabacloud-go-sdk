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
	// The deployment results of the stack.
	Deployments []*GetStackDeploymentsResponseBodyDeployments `json:"deployments,omitempty" xml:"deployments,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BF72A6FB-B071-5F2E-A036-9D62545B962C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
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
	// The configuration item.
	Config *GetStackDeploymentsResponseBodyDeploymentsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The configuration version, such as v1. The initial value is v1. The version number increments each time the stack is updated or refreshed and the configuration changes.
	//
	// example:
	//
	// v1
	ConfigVersion *string `json:"configVersion,omitempty" xml:"configVersion,omitempty"`
	// The creation time in UTC, in the format of YYYY-MM-DDTHH:mm:ssZ (ISO 8601).
	//
	// example:
	//
	// 2026-04-01T12:10:18Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The deployment name.
	//
	// example:
	//
	// production
	DeploymentName *string `json:"deploymentName,omitempty" xml:"deploymentName,omitempty"`
	// The deployment number. The deployment number for each stack starts from 1 and increments each time a deployment is successfully triggered.
	//
	// example:
	//
	// 1
	DeploymentNo *string `json:"deploymentNo,omitempty" xml:"deploymentNo,omitempty"`
	// Deprecated field.
	//
	// example:
	//
	// v1
	DeploymentVersion *string `json:"deploymentVersion,omitempty" xml:"deploymentVersion,omitempty"`
	// The execution duration, in milliseconds.
	//
	// example:
	//
	// 38000
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// The execution type.
	//
	// Manual: Manual execution (default).
	//
	// Auto: Automatic execution.
	//
	// example:
	//
	// Manual
	ExecuteType *string `json:"executeType,omitempty" xml:"executeType,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// \\n Error: Invalid value for input variable\\n \\n   on main.tf line 17, in module \\"alb\\":\\n   17:   log_project           = var.log_project.project_name\\n \\n The given value is not suitable for module.alb.var.log_project declared at\\n modules/alb/main.tf:34,1-23: string required.\\n╵\\n
	FailedReason *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	// The job ID.
	//
	// example:
	//
	// job-as154vldqt46mv0ixxxxx
	JobId *string `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// OSS object key prefix for deployment logs
	LogOutputPath *string `json:"logOutputPath,omitempty" xml:"logOutputPath,omitempty"`
	// The outputs.
	Outputs []*GetStackDeploymentsResponseBodyDeploymentsOutputs `json:"outputs,omitempty" xml:"outputs,omitempty" type:"Repeated"`
	// The parameter set content.
	Parameters []*GetStackDeploymentsResponseBodyDeploymentsParameters `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	// The state file output results.
	PlanOutputs []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputs `json:"planOutputs,omitempty" xml:"planOutputs,omitempty" type:"Repeated"`
	// The deployment status.
	//
	// | Name | Description |
	//
	// |------|------|
	//
	// | Pending | The initial status after the deployment is created. |
	//
	// | PriorityQueued | Priority queuing in progress. |
	//
	// | PlanQueued | The deployment is queuing because no workflow is available after creation. |
	//
	// | ApplyQueued | The deployment is queuing because no workflow is available during execution. |
	//
	// | Planning | The resource deployment is in the Plan phase. |
	//
	// | Planned | The resource deployment has completed the Plan phase. |
	//
	// | ConfigProactiveInProgress | Compliance pre-check in progress. |
	//
	// | ConfigProactiveSuccess | Compliance pre-check succeeded. |
	//
	// | DetectInProgress | Drift detection in progress. |
	//
	// | ImportQueued | The deployment is queuing because no workflow is available during Import execution. |
	//
	// | Importing | The resource deployment is in the Import phase. |
	//
	// | Imported | The resource deployment has completed the Import phase. |
	//
	// | StateQueued | The deployment is queuing because no workflow is available during state command execution. |
	//
	// | Stating | The resource deployment is executing the state command. |
	//
	// | Stated | The resource deployment has completed the state command execution. |
	//
	// | Confirmed | The resource deployment has been confirmed after the Plan phase. |
	//
	// | PlannedAndFinished | No diff was found after the Plan phase. The deployment is in a final status. |
	//
	// | Applying | The resource deployment is in the Apply phase. |
	//
	// | Applied | The resource deployment has completed the Apply phase. |
	//
	// | Discarded | The resource deployment has been discarded and is in a final status. |
	//
	// | Errored | The deployment execution encountered an error and is in a final status. |
	//
	// | ConfigProactiveFailure | Compliance pre-check failed. |
	//
	// | Canceled | The deployment execution has been canceled and is in a final status. |
	//
	// example:
	//
	// Pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task ID.
	//
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

func (s *GetStackDeploymentsResponseBodyDeployments) GetLogOutputPath() *string {
	return s.LogOutputPath
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

func (s *GetStackDeploymentsResponseBodyDeployments) SetLogOutputPath(v string) *GetStackDeploymentsResponseBodyDeployments {
	s.LogOutputPath = &v
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
	// Specifies whether to automatically execute the task. Default value: false. Valid values:
	//
	// - **false**: No.
	//
	// - **true**: Yes.
	//
	// example:
	//
	// false
	AutoApply *bool `json:"autoApply,omitempty" xml:"autoApply,omitempty"`
	// Indicates whether this is a destroy job.
	//
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
	// The description.
	//
	// example:
	//
	// The name of the SLS log project
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expression that can reference component outputs, in the format: component.{component name}.{component output name}.
	//
	// example:
	//
	// component.sls.project_name
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	// The name.
	//
	// example:
	//
	// project_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter type.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The actual value after the deployment is completed.
	//
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
	// The default value of the parameter.
	//
	// example:
	//
	// cn-hangzhou
	DefaultValue *string `json:"defaultValue,omitempty" xml:"defaultValue,omitempty"`
	// The description.
	//
	// example:
	//
	// region of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// region
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether the parameter is sensitive. Sensitive parameter values are not visible in the console or API. Valid values:
	//
	// - true: Sensitive.
	//
	// - false: Not sensitive.
	Sensitive *bool `json:"sensitive,omitempty" xml:"sensitive,omitempty"`
	// The parameter type.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The parameter value.
	//
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

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) GetSensitive() *bool {
	return s.Sensitive
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

func (s *GetStackDeploymentsResponseBodyDeploymentsParameters) SetSensitive(v bool) *GetStackDeploymentsResponseBodyDeploymentsParameters {
	s.Sensitive = &v
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
	// The change type of the component. Valid values:
	//
	// - create: All resource changes in the component are creations.
	//
	// - delete: All resource changes in the component are deletions.
	//
	// - read: All resource changes in the component are reads.
	//
	// - update: Resource changes in the component include two or more types among creation, deletion, and read.
	//
	// example:
	//
	// update
	ModuleAction *string `json:"moduleAction,omitempty" xml:"moduleAction,omitempty"`
	// The number of resources to be created, updated, and destroyed in this deployment.
	ModuleActionDetail *GetStackDeploymentsResponseBodyDeploymentsPlanOutputsModuleActionDetail `json:"moduleActionDetail,omitempty" xml:"moduleActionDetail,omitempty" type:"Struct"`
	// The resource change information.
	ResourceChanges []*GetStackDeploymentsResponseBodyDeploymentsPlanOutputsResourceChanges `json:"resourceChanges,omitempty" xml:"resourceChanges,omitempty" type:"Repeated"`
	// The component name of the stack.
	//
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
	// The number of resources to be created.
	//
	// example:
	//
	// 0
	Add *int32 `json:"add,omitempty" xml:"add,omitempty"`
	// The number of resources to be changed.
	//
	// example:
	//
	// 1
	Change *int32 `json:"change,omitempty" xml:"change,omitempty"`
	// The number of resources to be destroyed.
	//
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
	// The diff information of the resource change.
	//
	// example:
	//
	// ~ resource \\"alicloud_log_store\\" \\"default\\" {\\n        id                    = \\"alb-log-project-v1-ph-xxxxx:alb-log-store-ph\\"\\n      ~ max_split_shard_count = 64 -> 32\\n        name                  = \\"alb-log-store-ph\\"\\n\\n        # (13 unchanged attributes hidden)\\n    }
	Change *string `json:"change,omitempty" xml:"change,omitempty"`
	// The types of resource change actions included in this resource change.
	ResourceActions []*string `json:"resourceActions,omitempty" xml:"resourceActions,omitempty" type:"Repeated"`
	// The unique identifier of the resource.
	//
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
