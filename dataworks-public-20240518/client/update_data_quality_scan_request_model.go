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
	ComputeResource *UpdateDataQualityScanRequestComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Description     *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks           []*UpdateDataQualityScanRequestHooks         `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 231263586109857423
	Owner      *string                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Parameters []*UpdateDataQualityScanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// 101
	ProjectId       *int64                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RuntimeResource *UpdateDataQualityScanRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	Spec            *string                                      `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Trigger         *UpdateDataQualityScanRequestTrigger         `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// auto_createAlertRule_Finished_1kUTk6
	Name    *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	// example:
	//
	// temp_237669.zip_byBwm_1734661241752
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 20315
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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
