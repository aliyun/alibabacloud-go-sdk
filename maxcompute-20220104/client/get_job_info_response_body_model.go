// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody
	GetData() *GetJobInfoResponseBodyData
	SetErrorCode(v string) *GetJobInfoResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetJobInfoResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetJobInfoResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetJobInfoResponseBody
	GetRequestId() *string
}

type GetJobInfoResponseBody struct {
	// The returned result.
	Data *GetJobInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0be3e0bb16654558425251398e27a9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBody) GetData() *GetJobInfoResponseBodyData {
	return s.Data
}

func (s *GetJobInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobInfoResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetJobInfoResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobInfoResponseBody) SetData(v *GetJobInfoResponseBodyData) *GetJobInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetJobInfoResponseBody) SetErrorCode(v string) *GetJobInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobInfoResponseBody) SetErrorMsg(v string) *GetJobInfoResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobInfoResponseBody) SetHttpCode(v int32) *GetJobInfoResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobInfoResponseBody) SetRequestId(v string) *GetJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyData struct {
	// The amount of resources consumed by the job. This parameter is returned only for jobs that are complete.Unit: 100\\*Core\\*s.
	//
	// example:
	//
	// 10
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The end time of the job.
	//
	// example:
	//
	// 1672112913
	EndAtTime *int64 `json:"endAtTime,omitempty" xml:"endAtTime,omitempty"`
	// The ID of the ancestor node.
	//
	// example:
	//
	// node_4
	ExtNodeId *string `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	// The Alibaba Cloud account ID of the task owner.
	//
	// example:
	//
	// duty_2
	ExtNodeOnDuty *string `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	// The upstream platform.
	//
	// example:
	//
	// platform_3
	ExtPlantFrom *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	// The amount of data scanned by the job.
	//
	// example:
	//
	// 1234
	InputBytes *float64 `json:"inputBytes,omitempty" xml:"inputBytes,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 20230410****60gg
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The owner of the job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The substatuses of the job lifecycle.
	JobSubStatusList []*GetJobInfoResponseBodyDataJobSubStatusList `json:"jobSubStatusList,omitempty" xml:"jobSubStatusList,omitempty" type:"Repeated"`
	// The type of the job.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// The number of memory consumed by the job. This parameter is returned only for jobs that are complete.Unit: MB\\*s.
	//
	// example:
	//
	// 40
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The project name.
	//
	// example:
	//
	// dp_cdm_prod
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the computing quota that is used by the job.
	//
	// example:
	//
	// os_bigdata
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The quota type.
	//
	// example:
	//
	// subscription
	QuotaType *string `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The start time, which is the time when the job received the first batch of computing resources. For jobs that run for a short period of time or do not consume computing resources, such as the jobs that involve DDL statements, the job submission time is used instead.
	//
	// example:
	//
	// 1672112113
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The execution duration, which is the duration from the start time to the end time of the job.
	//
	// example:
	//
	// 800
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The intelligent diagnostics result.
	SceneResults []*GetJobInfoResponseBodyDataSceneResults `json:"sceneResults,omitempty" xml:"sceneResults,omitempty" type:"Repeated"`
	// The signature of the SQL job. You can use the signature to find the instances on which each time an SQL statement is executed.
	//
	// example:
	//
	// 20c1efb4a7caca1865f4aa784bb500efae74af04
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The job status.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1672112013
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 4784****5249
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The total duration from the time a job is submitted to the time the job is terminated.
	//
	// example:
	//
	// 900
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The wait time, which is the duration from the time the job is submitted to the time the job starts to run.
	//
	// example:
	//
	// 100
	WaitingTime *int64 `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s GetJobInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyData) GetCuUsage() *int64 {
	return s.CuUsage
}

func (s *GetJobInfoResponseBodyData) GetEndAtTime() *int64 {
	return s.EndAtTime
}

func (s *GetJobInfoResponseBodyData) GetExtNodeId() *string {
	return s.ExtNodeId
}

func (s *GetJobInfoResponseBodyData) GetExtNodeOnDuty() *string {
	return s.ExtNodeOnDuty
}

func (s *GetJobInfoResponseBodyData) GetExtPlantFrom() *string {
	return s.ExtPlantFrom
}

func (s *GetJobInfoResponseBodyData) GetInputBytes() *float64 {
	return s.InputBytes
}

func (s *GetJobInfoResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetJobInfoResponseBodyData) GetJobOwner() *string {
	return s.JobOwner
}

func (s *GetJobInfoResponseBodyData) GetJobSubStatusList() []*GetJobInfoResponseBodyDataJobSubStatusList {
	return s.JobSubStatusList
}

func (s *GetJobInfoResponseBodyData) GetJobType() *string {
	return s.JobType
}

func (s *GetJobInfoResponseBodyData) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *GetJobInfoResponseBodyData) GetPriority() *int64 {
	return s.Priority
}

func (s *GetJobInfoResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *GetJobInfoResponseBodyData) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *GetJobInfoResponseBodyData) GetQuotaType() *string {
	return s.QuotaType
}

func (s *GetJobInfoResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetJobInfoResponseBodyData) GetRunningAtTime() *int64 {
	return s.RunningAtTime
}

func (s *GetJobInfoResponseBodyData) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *GetJobInfoResponseBodyData) GetSceneResults() []*GetJobInfoResponseBodyDataSceneResults {
	return s.SceneResults
}

func (s *GetJobInfoResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetJobInfoResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetJobInfoResponseBodyData) GetSubmittedAtTime() *int64 {
	return s.SubmittedAtTime
}

func (s *GetJobInfoResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetJobInfoResponseBodyData) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *GetJobInfoResponseBodyData) GetWaitingTime() *int64 {
	return s.WaitingTime
}

func (s *GetJobInfoResponseBodyData) SetCuUsage(v int64) *GetJobInfoResponseBodyData {
	s.CuUsage = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetEndAtTime(v int64) *GetJobInfoResponseBodyData {
	s.EndAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtNodeId(v string) *GetJobInfoResponseBodyData {
	s.ExtNodeId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtNodeOnDuty(v string) *GetJobInfoResponseBodyData {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetExtPlantFrom(v string) *GetJobInfoResponseBodyData {
	s.ExtPlantFrom = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetInputBytes(v float64) *GetJobInfoResponseBodyData {
	s.InputBytes = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetInstanceId(v string) *GetJobInfoResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobOwner(v string) *GetJobInfoResponseBodyData {
	s.JobOwner = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobSubStatusList(v []*GetJobInfoResponseBodyDataJobSubStatusList) *GetJobInfoResponseBodyData {
	s.JobSubStatusList = v
	return s
}

func (s *GetJobInfoResponseBodyData) SetJobType(v string) *GetJobInfoResponseBodyData {
	s.JobType = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetMemoryUsage(v int64) *GetJobInfoResponseBodyData {
	s.MemoryUsage = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetPriority(v int64) *GetJobInfoResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetProject(v string) *GetJobInfoResponseBodyData {
	s.Project = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetQuotaNickname(v string) *GetJobInfoResponseBodyData {
	s.QuotaNickname = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetQuotaType(v string) *GetJobInfoResponseBodyData {
	s.QuotaType = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRegion(v string) *GetJobInfoResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRunningAtTime(v int64) *GetJobInfoResponseBodyData {
	s.RunningAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetRunningTime(v int64) *GetJobInfoResponseBodyData {
	s.RunningTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSceneResults(v []*GetJobInfoResponseBodyDataSceneResults) *GetJobInfoResponseBodyData {
	s.SceneResults = v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSignature(v string) *GetJobInfoResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetStatus(v string) *GetJobInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetSubmittedAtTime(v int64) *GetJobInfoResponseBodyData {
	s.SubmittedAtTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetTenantId(v string) *GetJobInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetTotalTime(v int64) *GetJobInfoResponseBodyData {
	s.TotalTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) SetWaitingTime(v int64) *GetJobInfoResponseBodyData {
	s.WaitingTime = &v
	return s
}

func (s *GetJobInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataJobSubStatusList struct {
	// The encoding of the substatus.
	//
	// example:
	//
	// 1010
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the substatus.
	//
	// example:
	//
	// Waiting for scheduling
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The start time of the substatus.
	//
	// example:
	//
	// 2025-03-05 00:04:15.717364 +0800
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetJobInfoResponseBodyDataJobSubStatusList) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataJobSubStatusList) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetCode() *int32 {
	return s.Code
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetDescription() *string {
	return s.Description
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetCode(v int32) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.Code = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetDescription(v string) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) SetStartTime(v string) *GetJobInfoResponseBodyDataJobSubStatusList {
	s.StartTime = &v
	return s
}

func (s *GetJobInfoResponseBodyDataJobSubStatusList) Validate() error {
	return dara.Validate(s)
}

type GetJobInfoResponseBodyDataSceneResults struct {
	// The intelligent diagnostics result description.
	//
	// example:
	//
	// This job uses annual and monthly computing resources. It may be that the job is waiting for resources due to the large amount of overall job running data, many resources requested, and low job priority. Please go to Resource Consumption to view the specific situation. You can also go to Cost Optimization to see if you need to adjust resource configuration.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Information about the nodes where data skew or data expansion is detected. This parameter is returned only when the diagnostics scenario is data skew or data expansion.
	Params map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	// The intelligent diagnostics result scenario.
	//
	// example:
	//
	// LackResource
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The intelligent diagnostics result tag.
	//
	// example:
	//
	// SubscriptionLackResource
	SceneTag *string `json:"sceneTag,omitempty" xml:"sceneTag,omitempty"`
	// The intelligent diagnostics result summary.
	//
	// example:
	//
	// Insufficient computing resources available for the job. Click to view details.
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// The intelligent diagnostics result type.
	//
	// example:
	//
	// warning
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetJobInfoResponseBodyDataSceneResults) String() string {
	return dara.Prettify(s)
}

func (s GetJobInfoResponseBodyDataSceneResults) GoString() string {
	return s.String()
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetDescription() *string {
	return s.Description
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetParams() map[string]*string {
	return s.Params
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetScene() *string {
	return s.Scene
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetSceneTag() *string {
	return s.SceneTag
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetSummary() *string {
	return s.Summary
}

func (s *GetJobInfoResponseBodyDataSceneResults) GetType() *string {
	return s.Type
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetDescription(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Description = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetParams(v map[string]*string) *GetJobInfoResponseBodyDataSceneResults {
	s.Params = v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetScene(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Scene = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetSceneTag(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.SceneTag = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetSummary(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Summary = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) SetType(v string) *GetJobInfoResponseBodyDataSceneResults {
	s.Type = &v
	return s
}

func (s *GetJobInfoResponseBodyDataSceneResults) Validate() error {
	return dara.Validate(s)
}
