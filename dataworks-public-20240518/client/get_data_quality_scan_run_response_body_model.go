// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityScanRun(v *GetDataQualityScanRunResponseBodyDataQualityScanRun) *GetDataQualityScanRunResponseBody
	GetDataQualityScanRun() *GetDataQualityScanRunResponseBodyDataQualityScanRun
	SetRequestId(v string) *GetDataQualityScanRunResponseBody
	GetRequestId() *string
}

type GetDataQualityScanRunResponseBody struct {
	DataQualityScanRun *GetDataQualityScanRunResponseBodyDataQualityScanRun `json:"DataQualityScanRun,omitempty" xml:"DataQualityScanRun,omitempty" type:"Struct"`
	// example:
	//
	// 0bc14115****159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityScanRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBody) GetDataQualityScanRun() *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	return s.DataQualityScanRun
}

func (s *GetDataQualityScanRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityScanRunResponseBody) SetDataQualityScanRun(v *GetDataQualityScanRunResponseBodyDataQualityScanRun) *GetDataQualityScanRunResponseBody {
	s.DataQualityScanRun = v
	return s
}

func (s *GetDataQualityScanRunResponseBody) SetRequestId(v string) *GetDataQualityScanRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityScanRunResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRun struct {
	// example:
	//
	// 1706247622000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1706247622000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// 1016440997
	Id         *int64                                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	Parameters []*GetDataQualityScanRunResponseBodyDataQualityScanRunParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	Results    []*GetDataQualityScanRunResponseBodyDataQualityScanRunResults    `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Scan       *GetDataQualityScanRunResponseBodyDataQualityScanRunScan         `json:"Scan,omitempty" xml:"Scan,omitempty" type:"Struct"`
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRun) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRun) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetParameters() []*GetDataQualityScanRunResponseBodyDataQualityScanRunParameters {
	return s.Parameters
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetResults() []*GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	return s.Results
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetScan() *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	return s.Scan
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetCreateTime(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetFinishTime(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.FinishTime = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetId(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetParameters(v []*GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.Parameters = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetResults(v []*GetDataQualityScanRunResponseBodyDataQualityScanRunResults) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.Results = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetScan(v *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.Scan = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) SetStatus(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRun {
	s.Status = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRun) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunParameters struct {
	// example:
	//
	// dt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// $[yyyy-mm-dd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) SetName(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) SetValue(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters {
	s.Value = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunParameters) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunResults struct {
	// example:
	//
	// 1725506795000
	CreateTime *int64                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Details    []*GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	Rule       *string                                                              `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// {
	//
	//   "value": "100.0"
	//
	// }
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunResults) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GetDetails() []*GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails {
	return s.Details
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GetRule() *string {
	return s.Rule
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GetSample() *string {
	return s.Sample
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) SetCreateTime(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) SetDetails(v []*GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) *GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	s.Details = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) SetRule(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	s.Rule = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) SetSample(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	s.Sample = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) SetStatus(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResults {
	s.Status = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResults) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails struct {
	// example:
	//
	// 100.0
	CheckValue *string `json:"CheckValue,omitempty" xml:"CheckValue,omitempty"`
	// example:
	//
	// 0.0
	ReferenceValue *string `json:"ReferenceValue,omitempty" xml:"ReferenceValue,omitempty"`
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) GetCheckValue() *string {
	return s.CheckValue
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) GetReferenceValue() *string {
	return s.ReferenceValue
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) GetStatus() *string {
	return s.Status
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) SetCheckValue(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails {
	s.CheckValue = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) SetReferenceValue(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails {
	s.ReferenceValue = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) SetStatus(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails {
	s.Status = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunResultsDetails) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScan struct {
	ComputeResource *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Struct"`
	// example:
	//
	// 1706247622000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 7892346529452
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// This is a hourly run data quality evaluation plan.
	Description *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Hooks       []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks `json:"Hooks,omitempty" xml:"Hooks,omitempty" type:"Repeated"`
	// example:
	//
	// 21077
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1706247622000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 7892346529452
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 7892346529452
	Owner      *string                                                              `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Parameters []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// example:
	//
	// 164024
	ProjectId       *int64                                                                  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RuntimeResource *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	Spec            *string                                                                 `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Trigger         *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger         `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScan) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetComputeResource() *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource {
	return s.ComputeResource
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetHooks() []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks {
	return s.Hooks
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetOwner() *string {
	return s.Owner
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetParameters() []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters {
	return s.Parameters
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetRuntimeResource() *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource {
	return s.RuntimeResource
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetSpec() *string {
	return s.Spec
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) GetTrigger() *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger {
	return s.Trigger
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetComputeResource(v *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.ComputeResource = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetCreateTime(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.CreateTime = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetCreateUser(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.CreateUser = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetDescription(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Description = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetHooks(v []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Hooks = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetId(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetModifyTime(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.ModifyTime = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetModifyUser(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.ModifyUser = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetName(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetOwner(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Owner = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetParameters(v []*GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Parameters = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetProjectId(v int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetRuntimeResource(v *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.RuntimeResource = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetSpec(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Spec = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) SetTrigger(v *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) *GetDataQualityScanRunResponseBodyDataQualityScanRunScan {
	s.Trigger = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScan) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource struct {
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// emr_cluster_001
	Name    *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Runtime *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) GetEnvType() *string {
	return s.EnvType
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) GetRuntime() *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime {
	return s.Runtime
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) SetEnvType(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource {
	s.EnvType = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) SetName(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) SetRuntime(v *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource {
	s.Runtime = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResource) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime struct {
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

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) GetEngine() *string {
	return s.Engine
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) GetHiveConf() map[string]interface{} {
	return s.HiveConf
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) GetSparkConf() map[string]interface{} {
	return s.SparkConf
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) SetEngine(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime {
	s.Engine = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) SetHiveConf(v map[string]interface{}) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime {
	s.HiveConf = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) SetSparkConf(v map[string]interface{}) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime {
	s.SparkConf = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanComputeResourceRuntime) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks struct {
	// example:
	//
	// results.any { r -> r.status == \\"fail\\" && r.rule.severity == \\"High\\" }
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// BlockTaskInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) GetCondition() *string {
	return s.Condition
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) GetType() *string {
	return s.Type
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) SetCondition(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks {
	s.Condition = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) SetType(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks {
	s.Type = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanHooks) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters struct {
	// example:
	//
	// dt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// $[yyyy-mm-dd-1]
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) GetName() *string {
	return s.Name
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) SetName(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters {
	s.Name = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) SetValue(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters {
	s.Value = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanParameters) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource struct {
	// example:
	//
	// 1
	Cu *float32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 60597
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// i-xxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) GetCu() *float32 {
	return s.Cu
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) GetId() *string {
	return s.Id
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) GetImage() *string {
	return s.Image
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) SetCu(v float32) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource {
	s.Cu = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) SetId(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) SetImage(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource {
	s.Image = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger struct {
	TaskIds []*int64 `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// example:
	//
	// BySchedule
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) GetType() *string {
	return s.Type
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) SetTaskIds(v []*int64) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger {
	s.TaskIds = v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) SetType(v string) *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger {
	s.Type = &v
	return s
}

func (s *GetDataQualityScanRunResponseBodyDataQualityScanRunScanTrigger) Validate() error {
	return dara.Validate(s)
}
