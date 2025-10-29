// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityScan(v *GetDataQualityScanResponseBodyDataQualityScan) *GetDataQualityScanResponseBody
	GetDataQualityScan() *GetDataQualityScanResponseBodyDataQualityScan
	SetRequestId(v string) *GetDataQualityScanResponseBody
	GetRequestId() *string
}

type GetDataQualityScanResponseBody struct {
	// Data quality monitoring details.
	DataQualityScan *GetDataQualityScanResponseBodyDataQualityScan `json:"DataQualityScan,omitempty" xml:"DataQualityScan,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBody) GetDataQualityScan() *GetDataQualityScanResponseBodyDataQualityScan {
	return s.DataQualityScan
}

func (s *GetDataQualityScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityScanResponseBody) SetDataQualityScan(v *GetDataQualityScanResponseBodyDataQualityScan) *GetDataQualityScanResponseBody {
	s.DataQualityScan = v
	return s
}

func (s *GetDataQualityScanResponseBody) SetRequestId(v string) *GetDataQualityScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityScanResponseBody) Validate() error {
	if s.DataQualityScan != nil {
		if err := s.DataQualityScan.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityScanResponseBodyDataQualityScan struct {
	// The compute engine used at runtime. Optional. If not specified, the data source defined in the Spec is used.
	ComputeResource *GetDataQualityScanResponseBodyDataQualityScanComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// The creation time of the data quality monitor.
	//
	// example:
	//
	// 1731550150000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who creates the data quality monitor.
	//
	// example:
	//
	// 2374924198591250
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The data quality monitor description.
	//
	// example:
	//
	// aily data quality scanning of ods tables.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Hook configurations after the data quality monitoring run ends.
	Hooks []*GetDataQualityScanResponseBodyDataQualityScanHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The data quality monitoring ID.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Last modified time of the data quality monitor.
	//
	// example:
	//
	// 1731550150000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the user who last modifies the data quality monitor.
	//
	// example:
	//
	// 23482597582479
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The data quality monitor name.
	//
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the user who owns the data quality monitor.
	//
	// example:
	//
	// 231263586109857423
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The definition of execution parameters for the data quality monitor.
	Parameters []*GetDataQualityScanResponseBodyDataQualityScanParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The workspace ID where the data quality monitor resides. You can obtain the workspace ID by calling the [ListProjects](https://help.aliyun.com/document_detail/2780068.html) operation.
	//
	// example:
	//
	// 101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used during the running of the data quality monitor.
	RuntimeResource *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// Spec code for the content of the data quality monitoring.
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
	// The trigger configurations of the data quality monitoring task.
	Trigger *GetDataQualityScanResponseBodyDataQualityScanTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s GetDataQualityScanResponseBodyDataQualityScan) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScan) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetComputeResource() *GetDataQualityScanResponseBodyDataQualityScanComputeResource {
	return s.ComputeResource
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetHooks() []*GetDataQualityScanResponseBodyDataQualityScanHooks {
	return s.Hooks
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetOwner() *string {
	return s.Owner
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetParameters() []*GetDataQualityScanResponseBodyDataQualityScanParameters {
	return s.Parameters
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetRuntimeResource() *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource {
	return s.RuntimeResource
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetSpec() *string {
	return s.Spec
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) GetTrigger() *GetDataQualityScanResponseBodyDataQualityScanTrigger {
	return s.Trigger
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetComputeResource(v *GetDataQualityScanResponseBodyDataQualityScanComputeResource) *GetDataQualityScanResponseBodyDataQualityScan {
	s.ComputeResource = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetCreateTime(v int64) *GetDataQualityScanResponseBodyDataQualityScan {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetCreateUser(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.CreateUser = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetDescription(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Description = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetHooks(v []*GetDataQualityScanResponseBodyDataQualityScanHooks) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Hooks = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetId(v int64) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetModifyTime(v int64) *GetDataQualityScanResponseBodyDataQualityScan {
	s.ModifyTime = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetModifyUser(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.ModifyUser = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetName(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetOwner(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Owner = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetParameters(v []*GetDataQualityScanResponseBodyDataQualityScanParameters) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Parameters = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetProjectId(v int64) *GetDataQualityScanResponseBodyDataQualityScan {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetRuntimeResource(v *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) *GetDataQualityScanResponseBodyDataQualityScan {
	s.RuntimeResource = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetSpec(v string) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Spec = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) SetTrigger(v *GetDataQualityScanResponseBodyDataQualityScanTrigger) *GetDataQualityScanResponseBodyDataQualityScan {
	s.Trigger = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScan) Validate() error {
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

type GetDataQualityScanResponseBodyDataQualityScanComputeResource struct {
	// The workspace environment to which the compute engine belongs.
	//
	// Valid values:
	//
	// 	- Prod: production environment .
	//
	// 	- Dev: development environment.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the compute engine, which is a unique identifier.
	//
	// example:
	//
	// polardb_to_holo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// More settings for data quality monitor at runtime.
	Runtime *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s GetDataQualityScanResponseBodyDataQualityScanComputeResource) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanComputeResource) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) GetRuntime() *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime {
	return s.Runtime
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) SetEnvType(v string) *GetDataQualityScanResponseBodyDataQualityScanComputeResource {
	s.EnvType = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) SetName(v string) *GetDataQualityScanResponseBodyDataQualityScanComputeResource {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) SetRuntime(v *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) *GetDataQualityScanResponseBodyDataQualityScanComputeResource {
	s.Runtime = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResource) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime struct {
	// The type of the compute engine. Only EMR compute engines support these settings.
	//
	// Valid values:
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

func (s GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) GetEngine() *string {
	return s.Engine
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) GetHiveConf() map[string]interface{} {
	return s.HiveConf
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) GetSparkConf() map[string]interface{} {
	return s.SparkConf
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) SetEngine(v string) *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime {
	s.Engine = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) SetHiveConf(v map[string]interface{}) *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime {
	s.HiveConf = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) SetSparkConf(v map[string]interface{}) *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime {
	s.SparkConf = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanHooks struct {
	// The Hook trigger condition. The hook will run if the condition is met. Currently, only one type of expression syntax is supported:
	//
	// 	- You can specify multiple combinations of rule severity levels and validation statuses using an expression such as `results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }`. This expression means the condition is met if any executed rule has a result of Fail with severity Normal, Error with severity High, or Warn with severity High. In the condition expression, the values of severity and status are predefined enums. The values of severity must match those defined in the Spec, and the values of status must match those in DataQualityResult.
	//
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The type of the Hook.
	//
	// Valid values:
	//
	// 	- BlockTaskInstance: BlockTaskInstance: Blocks the scheduling of the task instance.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityScanResponseBodyDataQualityScanHooks) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanHooks) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanHooks) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityScanResponseBodyDataQualityScanHooks) GetType() *string {
	return s.Type
}

func (s *GetDataQualityScanResponseBodyDataQualityScanHooks) SetCondition(v string) *GetDataQualityScanResponseBodyDataQualityScanHooks {
	s.Condition = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanHooks) SetType(v string) *GetDataQualityScanResponseBodyDataQualityScanHooks {
	s.Type = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanHooks) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanParameters struct {
	// The parameter value.
	//
	// example:
	//
	// e2e_autolabel
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityScanResponseBodyDataQualityScanParameters) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanParameters) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanParameters) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanResponseBodyDataQualityScanParameters) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityScanResponseBodyDataQualityScanParameters) SetName(v string) *GetDataQualityScanResponseBodyDataQualityScanParameters {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanParameters) SetValue(v string) *GetDataQualityScanResponseBodyDataQualityScanParameters {
	s.Value = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanParameters) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanRuntimeResource struct {
	// Reserved compute units (CU) for the resource group.
	//
	// example:
	//
	// 10
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 122878
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image ID used in the runtime configuration.
	//
	// example:
	//
	// hp-tlp-e2e-repo-registry-vpc.cn-heyuan-acdr-1.cr.aliyuncs.com/hp-service/worker:9b28b6d-202506091008
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) GetId() *string {
	return s.Id
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) SetCu(v float32) *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) SetId(v string) *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) SetImage(v string) *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanTrigger struct {
	// If the trigger mode is set to BySchedule, the scheduling task ID must be specified.
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The trigger mode of the monitoring task.
	//
	// Valid values:
	//
	// 	- ByManual: Manual trigger. This is the default setting.
	//
	// 	- BySchedule: Triggered by a scheduled task instance.
	//
	// example:
	//
	// BySchedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityScanResponseBodyDataQualityScanTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanResponseBodyDataQualityScanTrigger) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanResponseBodyDataQualityScanTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *GetDataQualityScanResponseBodyDataQualityScanTrigger) GetType() *string {
	return s.Type
}

func (s *GetDataQualityScanResponseBodyDataQualityScanTrigger) SetTaskIds(v []*int64) *GetDataQualityScanResponseBodyDataQualityScanTrigger {
	s.TaskIds = v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanTrigger) SetType(v string) *GetDataQualityScanResponseBodyDataQualityScanTrigger {
	s.Type = &v
	return s
}

func (s *GetDataQualityScanResponseBodyDataQualityScanTrigger) Validate() error {
	return dara.Validate(s)
}
