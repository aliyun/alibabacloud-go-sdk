// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateDataQualityScanRequest
	GetClientToken() *string
	SetComputeResource(v *CreateDataQualityScanRequestComputeResource) *CreateDataQualityScanRequest
	GetComputeResource() *CreateDataQualityScanRequestComputeResource
	SetDescription(v string) *CreateDataQualityScanRequest
	GetDescription() *string
	SetHooks(v []*CreateDataQualityScanRequestHooks) *CreateDataQualityScanRequest
	GetHooks() []*CreateDataQualityScanRequestHooks
	SetName(v string) *CreateDataQualityScanRequest
	GetName() *string
	SetOwner(v string) *CreateDataQualityScanRequest
	GetOwner() *string
	SetParameters(v []*CreateDataQualityScanRequestParameters) *CreateDataQualityScanRequest
	GetParameters() []*CreateDataQualityScanRequestParameters
	SetProjectId(v int64) *CreateDataQualityScanRequest
	GetProjectId() *int64
	SetRuntimeResource(v *CreateDataQualityScanRequestRuntimeResource) *CreateDataQualityScanRequest
	GetRuntimeResource() *CreateDataQualityScanRequestRuntimeResource
	SetSpec(v string) *CreateDataQualityScanRequest
	GetSpec() *string
	SetTrigger(v *CreateDataQualityScanRequestTrigger) *CreateDataQualityScanRequest
	GetTrigger() *CreateDataQualityScanRequestTrigger
}

type CreateDataQualityScanRequest struct {
	// The idempotency token.
	//
	// This parameter is required.
	//
	// example:
	//
	// a-customized-uuid
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The compute engine used when the data quality monitor runs. This parameter is optional. If it is not specified, the data source connection specified in the Spec is used.
	ComputeResource *CreateDataQualityScanRequestComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// The description of the data quality monitor.
	//
	// example:
	//
	// Daily data quality scanning of ods tables.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration for hooks that run after the data quality monitor finishes.
	Hooks []*CreateDataQualityScanRequestHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The name of the data quality monitor.
	//
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner of the data quality monitor.
	//
	// example:
	//
	// 95279527****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The execution parameter definitions of the data quality monitor.
	Parameters []*CreateDataQualityScanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/overview) and go to the Workspace Management page to obtain the workspace ID. This parameter specifies the DataWorks workspace for the API call.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used when the data quality monitor runs.
	RuntimeResource *CreateDataQualityScanRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The Spec code that defines the data quality monitor. For more information, see [Data quality spec configuration guide](https://help.aliyun.com/document_detail/2963394.html).
	//
	// example:
	//
	// {
	//
	//     "datasets": [
	//
	//         {
	//
	//             "type": "Table",
	//
	//             "dataSource": {
	//
	//                 "name": "odps_first",
	//
	//                 "envType": "Prod"
	//
	//             },
	//
	//             "tables": [
	//
	//                 "ods_d_user_info"
	//
	//             ],
	//
	//             "filter": "pt = $[yyyymmdd-1]"
	//
	//         }
	//
	//     ],
	//
	//     "rules": [
	//
	//         {
	//
	//             "assertion": "row_count > 0"
	//
	//         }, {
	//
	//             "templateId": "SYSTEM:field:null_value:fixed",
	//
	//             "pass": "when = 0",
	//
	//             "name": "The id cannot be empty.",
	//
	//             "severity": "High",
	//
	//              "identity": "a-customized-data-quality-rule-uuid"
	//
	//         }
	//
	//     ]
	//
	// }
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The trigger settings of the data quality monitor.
	Trigger *CreateDataQualityScanRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s CreateDataQualityScanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataQualityScanRequest) GetComputeResource() *CreateDataQualityScanRequestComputeResource {
	return s.ComputeResource
}

func (s *CreateDataQualityScanRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityScanRequest) GetHooks() []*CreateDataQualityScanRequestHooks {
	return s.Hooks
}

func (s *CreateDataQualityScanRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityScanRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateDataQualityScanRequest) GetParameters() []*CreateDataQualityScanRequestParameters {
	return s.Parameters
}

func (s *CreateDataQualityScanRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityScanRequest) GetRuntimeResource() *CreateDataQualityScanRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *CreateDataQualityScanRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateDataQualityScanRequest) GetTrigger() *CreateDataQualityScanRequestTrigger {
	return s.Trigger
}

func (s *CreateDataQualityScanRequest) SetClientToken(v string) *CreateDataQualityScanRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetComputeResource(v *CreateDataQualityScanRequestComputeResource) *CreateDataQualityScanRequest {
	s.ComputeResource = v
	return s
}

func (s *CreateDataQualityScanRequest) SetDescription(v string) *CreateDataQualityScanRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetHooks(v []*CreateDataQualityScanRequestHooks) *CreateDataQualityScanRequest {
	s.Hooks = v
	return s
}

func (s *CreateDataQualityScanRequest) SetName(v string) *CreateDataQualityScanRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetOwner(v string) *CreateDataQualityScanRequest {
	s.Owner = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetParameters(v []*CreateDataQualityScanRequestParameters) *CreateDataQualityScanRequest {
	s.Parameters = v
	return s
}

func (s *CreateDataQualityScanRequest) SetProjectId(v int64) *CreateDataQualityScanRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetRuntimeResource(v *CreateDataQualityScanRequestRuntimeResource) *CreateDataQualityScanRequest {
	s.RuntimeResource = v
	return s
}

func (s *CreateDataQualityScanRequest) SetSpec(v string) *CreateDataQualityScanRequest {
	s.Spec = &v
	return s
}

func (s *CreateDataQualityScanRequest) SetTrigger(v *CreateDataQualityScanRequestTrigger) *CreateDataQualityScanRequest {
	s.Trigger = v
	return s
}

func (s *CreateDataQualityScanRequest) Validate() error {
	if s.ComputeResource != nil {
		if err := s.ComputeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Hooks != nil {
		for _, item := range s.Hooks {
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
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityScanRequestComputeResource struct {
	// The workspace environment to which the compute engine belongs. Valid values:
	//
	// - Prod: production environment
	//
	// - Dev: development environment
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the compute engine, which serves as a unique identifier.
	//
	// example:
	//
	// emr_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The additional settings for the data quality monitor at runtime.
	Runtime *CreateDataQualityScanRequestComputeResourceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s CreateDataQualityScanRequestComputeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestComputeResource) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestComputeResource) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateDataQualityScanRequestComputeResource) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityScanRequestComputeResource) GetRuntime() *CreateDataQualityScanRequestComputeResourceRuntime {
	return s.Runtime
}

func (s *CreateDataQualityScanRequestComputeResource) SetEnvType(v string) *CreateDataQualityScanRequestComputeResource {
	s.EnvType = &v
	return s
}

func (s *CreateDataQualityScanRequestComputeResource) SetName(v string) *CreateDataQualityScanRequestComputeResource {
	s.Name = &v
	return s
}

func (s *CreateDataQualityScanRequestComputeResource) SetRuntime(v *CreateDataQualityScanRequestComputeResourceRuntime) *CreateDataQualityScanRequestComputeResource {
	s.Runtime = v
	return s
}

func (s *CreateDataQualityScanRequestComputeResource) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityScanRequestComputeResourceRuntime struct {
	// The type of the engine used. Only the E-MapReduce (EMR) compute engine supports this setting. Valid values:
	//
	// - Hive: Hive SQL
	//
	// - Spark: Spark SQL
	//
	// - Kyuubi
	//
	// example:
	//
	// Hive
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The additional parameter settings when the Hive engine is used. Currently, only the queue can be set by using mapreduce.job.queuename.
	//
	// example:
	//
	// {"mapreduce.job.queuename": "dq_queue"}
	HiveConf map[string]interface{} `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// The additional parameter settings when the Spark engine is used. Currently, only the queue can be set by using spark.yarn.queue.
	//
	// example:
	//
	// {"spark.yarn.queue": "dq_queue"}
	SparkConf map[string]interface{} `json:"SparkConf,omitempty" xml:"SparkConf,omitempty"`
}

func (s CreateDataQualityScanRequestComputeResourceRuntime) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestComputeResourceRuntime) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) GetEngine() *string {
	return s.Engine
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) GetHiveConf() map[string]interface{} {
	return s.HiveConf
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) GetSparkConf() map[string]interface{} {
	return s.SparkConf
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) SetEngine(v string) *CreateDataQualityScanRequestComputeResourceRuntime {
	s.Engine = &v
	return s
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) SetHiveConf(v map[string]interface{}) *CreateDataQualityScanRequestComputeResourceRuntime {
	s.HiveConf = v
	return s
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) SetSparkConf(v map[string]interface{}) *CreateDataQualityScanRequestComputeResourceRuntime {
	s.SparkConf = v
	return s
}

func (s *CreateDataQualityScanRequestComputeResourceRuntime) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityScanRequestHooks struct {
	// The condition that triggers the hook. When this condition is met, the hook is triggered. Currently, only one expression syntax is supported:
	//
	// Specify multiple groups of rule severity levels and rule check statuses. For example, `results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }` indicates that among the executed rules, if the check result of a rule with the severity of Normal is Fail, the check result of a rule with the severity of High is Error, or the check result of a rule with the severity of High is Warn, the condition is met. The enumerations of severity in the conditional expression are consistent with those in the spec, and the enumerations of status are consistent with those in DataQualityResult.
	//
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the hook. Valid values:
	//
	// - BlockTaskInstance: blocks task instance scheduling.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityScanRequestHooks) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestHooks) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestHooks) GetCondition() *string {
	return s.Condition
}

func (s *CreateDataQualityScanRequestHooks) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityScanRequestHooks) SetCondition(v string) *CreateDataQualityScanRequestHooks {
	s.Condition = &v
	return s
}

func (s *CreateDataQualityScanRequestHooks) SetType(v string) *CreateDataQualityScanRequestHooks {
	s.Type = &v
	return s
}

func (s *CreateDataQualityScanRequestHooks) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityScanRequestParameters struct {
	// The name of the parameter.
	//
	// example:
	//
	// triggerTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// $[yyyymmdd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDataQualityScanRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestParameters) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityScanRequestParameters) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityScanRequestParameters) SetName(v string) *CreateDataQualityScanRequestParameters {
	s.Name = &v
	return s
}

func (s *CreateDataQualityScanRequestParameters) SetValue(v string) *CreateDataQualityScanRequestParameters {
	s.Value = &v
	return s
}

func (s *CreateDataQualityScanRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityScanRequestRuntimeResource struct {
	// The compute unit (CU) consumption of the runtime configuration.
	//
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 0525242e-d0ee-4bda-bc73-765d82f6a34a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image ID of the runtime configuration.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s CreateDataQualityScanRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *CreateDataQualityScanRequestRuntimeResource) GetId() *string {
	return s.Id
}

func (s *CreateDataQualityScanRequestRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *CreateDataQualityScanRequestRuntimeResource) SetCu(v float32) *CreateDataQualityScanRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *CreateDataQualityScanRequestRuntimeResource) SetId(v string) *CreateDataQualityScanRequestRuntimeResource {
	s.Id = &v
	return s
}

func (s *CreateDataQualityScanRequestRuntimeResource) SetImage(v string) *CreateDataQualityScanRequestRuntimeResource {
	s.Image = &v
	return s
}

func (s *CreateDataQualityScanRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityScanRequestTrigger struct {
	// The IDs of the scheduled tasks that trigger the data quality monitor when the trigger method is set to BySchedule.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger method of the data quality monitor. Valid values:
	//
	// - ByManual: manual trigger. This is the default value.
	//
	// - BySchedule: scheduled task instance trigger.
	//
	// example:
	//
	// BySchedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityScanRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRequestTrigger) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRequestTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *CreateDataQualityScanRequestTrigger) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityScanRequestTrigger) SetTaskIds(v []*int64) *CreateDataQualityScanRequestTrigger {
	s.TaskIds = v
	return s
}

func (s *CreateDataQualityScanRequestTrigger) SetType(v string) *CreateDataQualityScanRequestTrigger {
	s.Type = &v
	return s
}

func (s *CreateDataQualityScanRequestTrigger) Validate() error {
	return dara.Validate(s)
}
