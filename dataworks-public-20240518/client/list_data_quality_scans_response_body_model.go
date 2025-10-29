// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListDataQualityScansResponseBodyPageInfo) *ListDataQualityScansResponseBody
	GetPageInfo() *ListDataQualityScansResponseBodyPageInfo
	SetRequestId(v string) *ListDataQualityScansResponseBody
	GetRequestId() *string
}

type ListDataQualityScansResponseBody struct {
	// The page information.
	PageInfo *ListDataQualityScansResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The API request ID, which is generated as a UUID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityScansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBody) GetPageInfo() *ListDataQualityScansResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListDataQualityScansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityScansResponseBody) SetPageInfo(v *ListDataQualityScansResponseBodyPageInfo) *ListDataQualityScansResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListDataQualityScansResponseBody) SetRequestId(v string) *ListDataQualityScansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityScansResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityScansResponseBodyPageInfo struct {
	// The list of data quality monitors.
	DataQualityScans []*ListDataQualityScansResponseBodyPageInfoDataQualityScans `json:"DataQualityScans,omitempty" xml:"DataQualityScans,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityScansResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfo) GetDataQualityScans() []*ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	return s.DataQualityScans
}

func (s *ListDataQualityScansResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityScansResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityScansResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityScansResponseBodyPageInfo) SetDataQualityScans(v []*ListDataQualityScansResponseBodyPageInfoDataQualityScans) *ListDataQualityScansResponseBodyPageInfo {
	s.DataQualityScans = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfo) SetPageNumber(v int32) *ListDataQualityScansResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfo) SetPageSize(v int32) *ListDataQualityScansResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfo) SetTotalCount(v int32) *ListDataQualityScansResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfo) Validate() error {
	if s.DataQualityScans != nil {
		for _, item := range s.DataQualityScans {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScans struct {
	// The compute engine used during execution. If it is not specified, the data source connection defined in the Spec will be used.
	ComputeResource *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// The creation time of the data quality monitor.
	//
	// example:
	//
	// 1694512304000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the data quality monitor.
	//
	// example:
	//
	// 7892346529452
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the data quality scan task. Maximum length: 65,535 characters.
	//
	// example:
	//
	// This is a hourly run data quality evaluation plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The hook configuration after the data quality monitor stops.
	Hooks []*ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 26433
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Last update time of the data quality monitor.
	//
	// example:
	//
	// 17236236472
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The user ID of the last person who updated the data quality monitor.
	//
	// example:
	//
	// 23782382795249
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the data quality scan task. Can include digits, letters, Chinese characters, and both half-width and full-width punctuation marks. Maximum length: 255 characters.
	//
	// example:
	//
	// Hourly partition quality monitoring
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the owner responsible for the data quality monitor.
	//
	// example:
	//
	// 23782382795249
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Execution parameter definitions for the data quality monitor.
	Parameters []*ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// 59094
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource group used during the execution of the data quality monitor.
	RuntimeResource *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// Trigger settings for the data quality monitor.
	Trigger *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScans) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScans) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetComputeResource() *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource {
	return s.ComputeResource
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetDescription() *string {
	return s.Description
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetHooks() []*ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks {
	return s.Hooks
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetName() *string {
	return s.Name
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetOwner() *string {
	return s.Owner
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetParameters() []*ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters {
	return s.Parameters
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetRuntimeResource() *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource {
	return s.RuntimeResource
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) GetTrigger() *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger {
	return s.Trigger
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetComputeResource(v *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.ComputeResource = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetCreateTime(v int64) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetCreateUser(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.CreateUser = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetDescription(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Description = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetHooks(v []*ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Hooks = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetId(v int64) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Id = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetModifyTime(v int64) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.ModifyTime = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetModifyUser(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.ModifyUser = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetName(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Name = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetOwner(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Owner = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetParameters(v []*ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Parameters = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetProjectId(v int64) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetRuntimeResource(v *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.RuntimeResource = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) SetTrigger(v *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) *ListDataQualityScansResponseBodyPageInfoDataQualityScans {
	s.Trigger = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScans) Validate() error {
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

type ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource struct {
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
	// The name of the computing engine. Uniquely identifies the engine.
	//
	// example:
	//
	// emr_cluster_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Additional runtime settings for the data quality monitor.
	Runtime *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) GetEnvType() *string {
	return s.EnvType
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) GetName() *string {
	return s.Name
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) GetRuntime() *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime {
	return s.Runtime
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) SetEnvType(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource {
	s.EnvType = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) SetName(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource {
	s.Name = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) SetRuntime(v *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource {
	s.Runtime = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResource) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime struct {
	// The engine type. These settings are only supported for the EMR compute engine. Valid values:
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
	// Additional parameters for the Hive engine. Currently, only mapreduce.job.queuename is supported to set the queue.
	//
	// example:
	//
	// mapreduce.job.queuename=dq_queue
	HiveConf *string `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
	// Additional parameters for the Spark engine. Currently, only spark.yarn.queue is supported to set the queue.
	//
	// example:
	//
	// spark.yarn.queue=dq_queue
	SparkConf *string `json:"SparkConf,omitempty" xml:"SparkConf,omitempty"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) GetEngine() *string {
	return s.Engine
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) GetHiveConf() *string {
	return s.HiveConf
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) GetSparkConf() *string {
	return s.SparkConf
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) SetEngine(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime {
	s.Engine = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) SetHiveConf(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime {
	s.HiveConf = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) SetSparkConf(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime {
	s.SparkConf = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansComputeResourceRuntime) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks struct {
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
	// 	- BlockTaskInstance: Blocks the scheduling of the task instance.
	//
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) GetCondition() *string {
	return s.Condition
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) GetType() *string {
	return s.Type
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) SetCondition(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks {
	s.Condition = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) SetType(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks {
	s.Type = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansHooks) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters struct {
	// The parameter name.
	//
	// example:
	//
	// dt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// $[yyyy-mm-dd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) GetName() *string {
	return s.Name
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) SetName(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters {
	s.Name = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) SetValue(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters {
	s.Value = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource struct {
	// CU consumption for task running.
	//
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// Serverless_resource_group_xxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the image configured for task running.
	//
	// example:
	//
	// i-xxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) GetId() *string {
	return s.Id
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) SetCu(v float32) *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource {
	s.Cu = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) SetId(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource {
	s.Id = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) SetImage(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource {
	s.Image = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger struct {
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

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) GoString() string {
	return s.String()
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) GetType() *string {
	return s.Type
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) SetTaskIds(v []*int64) *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger {
	s.TaskIds = v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) SetType(v string) *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger {
	s.Type = &v
	return s
}

func (s *ListDataQualityScansResponseBodyPageInfoDataQualityScansTrigger) Validate() error {
	return dara.Validate(s)
}
