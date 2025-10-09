// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v *UpdateDataQualityScanRequestComputeResource) *UpdateDataQualityScanRequest
	GetComputeResource() *UpdateDataQualityScanRequestComputeResource
	SetDescription(v string) *UpdateDataQualityScanRequest
	GetDescription() *string
	SetHooks(v []*UpdateDataQualityScanRequestHooks) *UpdateDataQualityScanRequest
	GetHooks() []*UpdateDataQualityScanRequestHooks
	SetId(v int64) *UpdateDataQualityScanRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityScanRequest
	GetName() *string
	SetOwner(v string) *UpdateDataQualityScanRequest
	GetOwner() *string
	SetParameters(v []*UpdateDataQualityScanRequestParameters) *UpdateDataQualityScanRequest
	GetParameters() []*UpdateDataQualityScanRequestParameters
	SetProjectId(v int64) *UpdateDataQualityScanRequest
	GetProjectId() *int64
	SetRuntimeResource(v *UpdateDataQualityScanRequestRuntimeResource) *UpdateDataQualityScanRequest
	GetRuntimeResource() *UpdateDataQualityScanRequestRuntimeResource
	SetSpec(v string) *UpdateDataQualityScanRequest
	GetSpec() *string
	SetTrigger(v *UpdateDataQualityScanRequestTrigger) *UpdateDataQualityScanRequest
	GetTrigger() *UpdateDataQualityScanRequestTrigger
}

type UpdateDataQualityScanRequest struct {
	// The compute engine used during execution. If it\\"s not specified, the data source connection defined in the Spec will be used.
	ComputeResource *UpdateDataQualityScanRequestComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// Description of the data quality monitor.
	//
	// example:
	//
	// Daily data quality scanning of ods tables.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook configuration after the data quality monitor stops.
	Hooks []*UpdateDataQualityScanRequestHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// 231263586109857423
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The definition of execution parameters for the data quality monitor.
	Parameters []*UpdateDataQualityScanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the DataWorks workspace where the data quality monitor resides. You can obtain the workspace ID by calling the [ListProjects](https://help.aliyun.com/document_detail/2852607.html) operation.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used during the execution of the data quality monitor.
	RuntimeResource *UpdateDataQualityScanRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The Spec code of the data quality monitoring content. For more information, see [Data quality Spec configuration description](https://help.aliyun.com/document_detail/2963394.html).
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
	// Trigger settings for the data quality monitor.
	Trigger *UpdateDataQualityScanRequestTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s UpdateDataQualityScanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequest) GetComputeResource() *UpdateDataQualityScanRequestComputeResource {
	return s.ComputeResource
}

func (s *UpdateDataQualityScanRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityScanRequest) GetHooks() []*UpdateDataQualityScanRequestHooks {
	return s.Hooks
}

func (s *UpdateDataQualityScanRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityScanRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityScanRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateDataQualityScanRequest) GetParameters() []*UpdateDataQualityScanRequestParameters {
	return s.Parameters
}

func (s *UpdateDataQualityScanRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityScanRequest) GetRuntimeResource() *UpdateDataQualityScanRequestRuntimeResource {
	return s.RuntimeResource
}

func (s *UpdateDataQualityScanRequest) GetSpec() *string {
	return s.Spec
}

func (s *UpdateDataQualityScanRequest) GetTrigger() *UpdateDataQualityScanRequestTrigger {
	return s.Trigger
}

func (s *UpdateDataQualityScanRequest) SetComputeResource(v *UpdateDataQualityScanRequestComputeResource) *UpdateDataQualityScanRequest {
	s.ComputeResource = v
	return s
}

func (s *UpdateDataQualityScanRequest) SetDescription(v string) *UpdateDataQualityScanRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetHooks(v []*UpdateDataQualityScanRequestHooks) *UpdateDataQualityScanRequest {
	s.Hooks = v
	return s
}

func (s *UpdateDataQualityScanRequest) SetId(v int64) *UpdateDataQualityScanRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetName(v string) *UpdateDataQualityScanRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetOwner(v string) *UpdateDataQualityScanRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetParameters(v []*UpdateDataQualityScanRequestParameters) *UpdateDataQualityScanRequest {
	s.Parameters = v
	return s
}

func (s *UpdateDataQualityScanRequest) SetProjectId(v int64) *UpdateDataQualityScanRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetRuntimeResource(v *UpdateDataQualityScanRequestRuntimeResource) *UpdateDataQualityScanRequest {
	s.RuntimeResource = v
	return s
}

func (s *UpdateDataQualityScanRequest) SetSpec(v string) *UpdateDataQualityScanRequest {
	s.Spec = &v
	return s
}

func (s *UpdateDataQualityScanRequest) SetTrigger(v *UpdateDataQualityScanRequestTrigger) *UpdateDataQualityScanRequest {
	s.Trigger = v
	return s
}

func (s *UpdateDataQualityScanRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestComputeResource struct {
	// Workspace environment of the compute engine. Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the compute engine, which is a unique identifier.
	//
	// example:
	//
	// auto_createAlertRule_Finished_1kUTk6
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Additional settings for the compute engine.
	Runtime *UpdateDataQualityScanRequestComputeResourceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s UpdateDataQualityScanRequestComputeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestComputeResource) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestComputeResource) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateDataQualityScanRequestComputeResource) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityScanRequestComputeResource) GetRuntime() *UpdateDataQualityScanRequestComputeResourceRuntime {
	return s.Runtime
}

func (s *UpdateDataQualityScanRequestComputeResource) SetEnvType(v string) *UpdateDataQualityScanRequestComputeResource {
	s.EnvType = &v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResource) SetName(v string) *UpdateDataQualityScanRequestComputeResource {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResource) SetRuntime(v *UpdateDataQualityScanRequestComputeResourceRuntime) *UpdateDataQualityScanRequestComputeResource {
	s.Runtime = v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestComputeResourceRuntime struct {
	// The engine type. These settings are only supported for the EMR compute engine.This setting? Valid values:
	//
	// 	- Hive: Hive SQL
	//
	// 	- Spark: Spark SQL
	//
	// 	- Kyuubi
	//
	// example:
	//
	// Hive
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Additional Hive engine parameters. Currently, only the mapreduce.job.queuename parameter is supported.
	//
	// example:
	//
	// mapreduce.job.queuename=dq_queue
	HiveConf map[string]interface{} `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// Additional Spark engine parameters. Currently, only the spark.yarn.queue parameter is supported.
	//
	// example:
	//
	// spark.yarn.queue=dq_queue
	SparkConf map[string]interface{} `json:"SparkConf,omitempty" xml:"SparkConf,omitempty"`
}

func (s UpdateDataQualityScanRequestComputeResourceRuntime) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestComputeResourceRuntime) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) GetEngine() *string {
	return s.Engine
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) GetHiveConf() map[string]interface{} {
	return s.HiveConf
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) GetSparkConf() map[string]interface{} {
	return s.SparkConf
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) SetEngine(v string) *UpdateDataQualityScanRequestComputeResourceRuntime {
	s.Engine = &v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) SetHiveConf(v map[string]interface{}) *UpdateDataQualityScanRequestComputeResourceRuntime {
	s.HiveConf = v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) SetSparkConf(v map[string]interface{}) *UpdateDataQualityScanRequestComputeResourceRuntime {
	s.SparkConf = v
	return s
}

func (s *UpdateDataQualityScanRequestComputeResourceRuntime) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestHooks struct {
	// The hook trigger condition. When this condition is met, the hook is triggered. Valid expression format:
	//
	// Specifies multiple combinations of rule severity levels and rule validation statuses, such as `results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }`. This means the hook is triggered if any executed rule has Fail with Normal severity, Error with High severity, or Warn with High severity. The severity values must match those defined in the Spec. The status values must match those in DataQualityResult.
	//
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the hook. Valid values:
	//
	// 	- BlockTaskInstance: Block the scheduling of the task instance.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityScanRequestHooks) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestHooks) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestHooks) GetCondition() *string {
	return s.Condition
}

func (s *UpdateDataQualityScanRequestHooks) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityScanRequestHooks) SetCondition(v string) *UpdateDataQualityScanRequestHooks {
	s.Condition = &v
	return s
}

func (s *UpdateDataQualityScanRequestHooks) SetType(v string) *UpdateDataQualityScanRequestHooks {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityScanRequestHooks) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestParameters struct {
	// The parameter name.
	//
	// example:
	//
	// temp_237669.zip_byBwm_1734661241752
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// smtp.qiye.aliyun.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateDataQualityScanRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestParameters) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityScanRequestParameters) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityScanRequestParameters) SetName(v string) *UpdateDataQualityScanRequestParameters {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityScanRequestParameters) SetValue(v string) *UpdateDataQualityScanRequestParameters {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityScanRequestParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestRuntimeResource struct {
	// The default number of CUs configured for task running.
	//
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 20315
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image ID of the task runtime configuration.
	//
	// example:
	//
	// i-xxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s UpdateDataQualityScanRequestRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestRuntimeResource) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *UpdateDataQualityScanRequestRuntimeResource) GetId() *string {
	return s.Id
}

func (s *UpdateDataQualityScanRequestRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *UpdateDataQualityScanRequestRuntimeResource) SetCu(v float32) *UpdateDataQualityScanRequestRuntimeResource {
	s.Cu = &v
	return s
}

func (s *UpdateDataQualityScanRequestRuntimeResource) SetId(v string) *UpdateDataQualityScanRequestRuntimeResource {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityScanRequestRuntimeResource) SetImage(v string) *UpdateDataQualityScanRequestRuntimeResource {
	s.Image = &v
	return s
}

func (s *UpdateDataQualityScanRequestRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityScanRequestTrigger struct {
	// If the trigger mode is BySchedule, the ID of the scheduling task that triggers the monitor must be configured.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger mode of the data quality monitor. Valid values:
	//
	// 	- ByManual: Manually triggered. Default setting.
	//
	// 	- BySchedule: Triggered by a scheduled task instance.
	//
	// example:
	//
	// BySchedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityScanRequestTrigger) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityScanRequestTrigger) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityScanRequestTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *UpdateDataQualityScanRequestTrigger) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityScanRequestTrigger) SetTaskIds(v []*int64) *UpdateDataQualityScanRequestTrigger {
	s.TaskIds = v
	return s
}

func (s *UpdateDataQualityScanRequestTrigger) SetType(v string) *UpdateDataQualityScanRequestTrigger {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityScanRequestTrigger) Validate() error {
	return dara.Validate(s)
}
