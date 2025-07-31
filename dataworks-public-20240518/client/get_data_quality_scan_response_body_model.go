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
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScan struct {
	ComputeResource *GetDataQualityScanResponseBodyDataQualityScanComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// example:
	//
	// 1731550150000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2374924198591250
	CreateUser  *string                                               `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Description *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks       []*GetDataQualityScanResponseBodyDataQualityScanHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1731550150000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 23482597582479
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 231263586109857423
	Owner      *string                                                    `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Parameters []*GetDataQualityScanResponseBodyDataQualityScanParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// 101
	ProjectId       *int64                                                        `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RuntimeResource *GetDataQualityScanResponseBodyDataQualityScanRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	Spec            *string                                                       `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Trigger         *GetDataQualityScanResponseBodyDataQualityScanTrigger         `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanComputeResource struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// polardb_to_holo
	Name    *string                                                              `json:"Name,omitempty" xml:"Name,omitempty"`
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
	return dara.Validate(s)
}

type GetDataQualityScanResponseBodyDataQualityScanComputeResourceRuntime struct {
	// example:
	//
	// Hive
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// mapreduce.job.queuename=dq_queue
	HiveConf map[string]interface{} `json:"HiveConf,omitempty" xml:"HiveConf,omitempty"`
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
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	// example:
	//
	// e2e_autolabel
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 10
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 122878
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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
