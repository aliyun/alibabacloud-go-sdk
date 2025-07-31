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
	// This parameter is required.
	//
	// example:
	//
	// a-customized-uuid
	ClientToken     *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ComputeResource *CreateDataQualityScanRequestComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	Description     *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks           []*CreateDataQualityScanRequestHooks         `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// data_quality_scan_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 95279527****
	Owner      *string                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Parameters []*CreateDataQualityScanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// 101
	ProjectId       *int64                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RuntimeResource *CreateDataQualityScanRequestRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	Spec            *string                                      `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Trigger         *CreateDataQualityScanRequestTrigger         `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateDataQualityScanRequestComputeResource struct {
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// emr_first
	Name    *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
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
	return dara.Validate(s)
}

type CreateDataQualityScanRequestComputeResourceRuntime struct {
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
	// example:
	//
	// results.any { r -> r.status == \\"Fail\\" && r.rule.severity == \\"Normal\\" || r.status == \\"Error\\" && r.rule.severity == \\"High\\" || r.status == \\"Warn\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
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
	// example:
	//
	// triggerTime
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 0.25
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 0525242e-d0ee-4bda-bc73-765d82f6a34a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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
