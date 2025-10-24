// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobSnapshotInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListJobSnapshotInfosResponseBodyData) *ListJobSnapshotInfosResponseBody
	GetData() *ListJobSnapshotInfosResponseBodyData
	SetErrorCode(v string) *ListJobSnapshotInfosResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListJobSnapshotInfosResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListJobSnapshotInfosResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListJobSnapshotInfosResponseBody
	GetRequestId() *string
}

type ListJobSnapshotInfosResponseBody struct {
	// The returned data.
	Data *ListJobSnapshotInfosResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: informational response. The request is received and is being processed.
	//
	// - 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// - 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// - 4xx: client error. The request contains invalid request parameters or syntaxes, or specific request conditions cannot be met.
	//
	// - 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7e716665825896565060e87a4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobSnapshotInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBody) GetData() *ListJobSnapshotInfosResponseBodyData {
	return s.Data
}

func (s *ListJobSnapshotInfosResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListJobSnapshotInfosResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListJobSnapshotInfosResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListJobSnapshotInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobSnapshotInfosResponseBody) SetData(v *ListJobSnapshotInfosResponseBodyData) *ListJobSnapshotInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetErrorCode(v string) *ListJobSnapshotInfosResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetErrorMsg(v string) *ListJobSnapshotInfosResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetHttpCode(v int32) *ListJobSnapshotInfosResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) SetRequestId(v string) *ListJobSnapshotInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListJobSnapshotInfosResponseBodyData struct {
	// The job snapshots.
	JobInfoList []*ListJobSnapshotInfosResponseBodyDataJobInfoList `json:"jobInfoList,omitempty" xml:"jobInfoList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned results.
	//
	// example:
	//
	// 123
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobSnapshotInfosResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBodyData) GetJobInfoList() []*ListJobSnapshotInfosResponseBodyDataJobInfoList {
	return s.JobInfoList
}

func (s *ListJobSnapshotInfosResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListJobSnapshotInfosResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListJobSnapshotInfosResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListJobSnapshotInfosResponseBodyData) SetJobInfoList(v []*ListJobSnapshotInfosResponseBodyDataJobInfoList) *ListJobSnapshotInfosResponseBodyData {
	s.JobInfoList = v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetPageNumber(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetPageSize(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) SetTotalCount(v int64) *ListJobSnapshotInfosResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyData) Validate() error {
	if s.JobInfoList != nil {
		for _, item := range s.JobInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJobSnapshotInfosResponseBodyDataJobInfoList struct {
	// The CPU request amount of the job at the snapshot time point. Unit: Core.
	//
	// example:
	//
	// 200
	CpuRequest *int64 `json:"cpuRequest,omitempty" xml:"cpuRequest,omitempty"`
	// CPU usage of the job at the snapshot time. Unit: Core.
	//
	// example:
	//
	// 100
	CpuUsage *int64 `json:"cpuUsage,omitempty" xml:"cpuUsage,omitempty"`
	// The CPU satisfaction ratio of the job at the snapshot time point (cpuUsage/cpuRequest).
	//
	// example:
	//
	// 0.5
	CpuUsageToRequestRatio *float64 `json:"cpuUsageToRequestRatio,omitempty" xml:"cpuUsageToRequestRatio,omitempty"`
	// The ID of the upstream node.
	//
	// example:
	//
	// 76358164
	ExtNodeId *string `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	// The account ID of the task owner.
	//
	// example:
	//
	// duty_2
	ExtNodeOnDuty *string `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	// The upstream platform.
	//
	// example:
	//
	// Dataworks
	ExtPlantFrom *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 20241028****jkl
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that commits the job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The type of the job.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// Not applicable.
	//
	// example:
	//
	// -1
	MaxCpuPct *float64 `json:"maxCpuPct,omitempty" xml:"maxCpuPct,omitempty"`
	// Not applicable.
	//
	// example:
	//
	// -1
	MaxMemoryPct *float64 `json:"maxMemoryPct,omitempty" xml:"maxMemoryPct,omitempty"`
	// The Memory request amount of the job at the snapshot time point. Unit: MB.
	//
	// example:
	//
	// 409600
	MemoryRequest *int64 `json:"memoryRequest,omitempty" xml:"memoryRequest,omitempty"`
	// Memory usage of the job at the snapshot time. Unit: MB.
	//
	// example:
	//
	// 409600
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The Memory satisfaction ratio of the job at the snapshot time point (memoryUsage/memoryRequest).
	//
	// example:
	//
	// 1
	MemoryUsageToRequestRatio *float64 `json:"memoryUsageToRequestRatio,omitempty" xml:"memoryUsageToRequestRatio,omitempty"`
	// The CPU usage ratio of the annual or monthly subscription job at the snapshot time (CPU usage / (reserved CPU guarantee + elastic reserved CPU)). This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinCpuPct *float64 `json:"minCpuPct,omitempty" xml:"minCpuPct,omitempty"`
	// The memory usage ratio of the annual or monthly subscription job at the observation time (memory usage / (reserved memory guarantee + elastic reserved memory)). This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinMemoryPct *float64 `json:"minMemoryPct,omitempty" xml:"minMemoryPct,omitempty"`
	// The priority of the job.
	//
	// example:
	//
	// 9
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// projectA
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the computing Quota used by the job.
	//
	// example:
	//
	// quota_A
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The type of the quota.
	//
	// example:
	//
	// subscription
	QuotaType *string `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The start time of the job.
	//
	// > The time when the job received the first batch of computing resources.
	//
	// example:
	//
	// 1736821805
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The running duration, which is the duration from the runningAtTime to the snapshotTime of the job.  Unit: seconds (s).
	//
	// example:
	//
	// 43
	RunningTime *int64 `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// signatureabc
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// The snapshot time.
	//
	// example:
	//
	// 1736821848
	SnapshotTime *int64 `json:"snapshotTime,omitempty" xml:"snapshotTime,omitempty"`
	// The snapshot status of the job.
	//
	// > The snapshot status is only running.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the job was committed.
	//
	// example:
	//
	// 1736821785
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 213065738244354
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The interval from the time when the job was submitted to the snapshotTime .Unit: seconds (s).
	//
	// example:
	//
	// 63
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The duration from the time the job is submitted to the time the job starts to run. Unit: seconds (s).
	//
	// example:
	//
	// 20
	WaitingTime *int64 `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s ListJobSnapshotInfosResponseBodyDataJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListJobSnapshotInfosResponseBodyDataJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetCpuRequest() *int64 {
	return s.CpuRequest
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetCpuUsage() *int64 {
	return s.CpuUsage
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetCpuUsageToRequestRatio() *float64 {
	return s.CpuUsageToRequestRatio
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetExtNodeId() *string {
	return s.ExtNodeId
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetExtNodeOnDuty() *string {
	return s.ExtNodeOnDuty
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetExtPlantFrom() *string {
	return s.ExtPlantFrom
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetJobType() *string {
	return s.JobType
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMaxCpuPct() *float64 {
	return s.MaxCpuPct
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMaxMemoryPct() *float64 {
	return s.MaxMemoryPct
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMemoryRequest() *int64 {
	return s.MemoryRequest
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMemoryUsage() *int64 {
	return s.MemoryUsage
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMemoryUsageToRequestRatio() *float64 {
	return s.MemoryUsageToRequestRatio
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMinCpuPct() *float64 {
	return s.MinCpuPct
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetMinMemoryPct() *float64 {
	return s.MinMemoryPct
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetPriority() *int64 {
	return s.Priority
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetProject() *string {
	return s.Project
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetQuotaType() *string {
	return s.QuotaType
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetRegion() *string {
	return s.Region
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetRunningAtTime() *int64 {
	return s.RunningAtTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetRunningTime() *int64 {
	return s.RunningTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetSignature() *string {
	return s.Signature
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetSnapshotTime() *int64 {
	return s.SnapshotTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetStatus() *string {
	return s.Status
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetSubmittedAtTime() *int64 {
	return s.SubmittedAtTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetTotalTime() *int64 {
	return s.TotalTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetWaitingTime() *int64 {
	return s.WaitingTime
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuRequest(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuRequest = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuUsage(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuUsage = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetCpuUsageToRequestRatio(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.CpuUsageToRequestRatio = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtNodeId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtNodeId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtNodeOnDuty(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtPlantFrom(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtPlantFrom = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetInstanceId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetJobOwner(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetJobType(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.JobType = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMaxCpuPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MaxCpuPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMaxMemoryPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MaxMemoryPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryRequest(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryRequest = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryUsage(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryUsage = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMemoryUsageToRequestRatio(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MemoryUsageToRequestRatio = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMinCpuPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MinCpuPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetMinMemoryPct(v float64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.MinMemoryPct = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetPriority(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Priority = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetProject(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Project = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetQuotaNickname(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetQuotaType(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.QuotaType = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRegion(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Region = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRunningAtTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetRunningTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.RunningTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSignature(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Signature = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSnapshotTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.SnapshotTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetStatus(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetSubmittedAtTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetTenantId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.TenantId = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetTotalTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.TotalTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetWaitingTime(v int64) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.WaitingTime = &v
	return s
}

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) Validate() error {
	return dara.Validate(s)
}
