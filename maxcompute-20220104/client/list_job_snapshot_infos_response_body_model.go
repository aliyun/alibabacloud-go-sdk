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
	// 异常信息
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// - 1xx: Informational - The request has been received and the process is continuing.
	//
	// - 2xx: Success - The request was successfully received, understood, and accepted.
	//
	// - 3xx: Redirection - Further action must be taken to complete the request.
	//
	// - 4xx: Client Error - The request contains bad syntax or cannot be fulfilled.
	//
	// - 5xx: Server Error - The server failed to fulfill an apparently valid request.
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
	// The list of job snapshots.
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
	// The total number of entries returned.
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
	// The amount of CPU cores requested by the job at the snapshot time.
	//
	// example:
	//
	// 200
	CpuRequest *int64 `json:"cpuRequest,omitempty" xml:"cpuRequest,omitempty"`
	// The CPU usage of the job at the snapshot time. Unit: cores.
	//
	// example:
	//
	// 100
	CpuUsage *int64 `json:"cpuUsage,omitempty" xml:"cpuUsage,omitempty"`
	// The CPU fulfillment ratio of the job at the snapshot time. This is calculated by dividing the CPU usage by the CPU request.
	//
	// example:
	//
	// 0.5
	CpuUsageToRequestRatio *float64 `json:"cpuUsageToRequestRatio,omitempty" xml:"cpuUsageToRequestRatio,omitempty"`
	// The upstream node ID.
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
	ExtPlantFrom  *string `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	ExtPlatformId *string `json:"extPlatformId,omitempty" xml:"extPlatformId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 20241028****jkl
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The job owner.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The job type.
	//
	// example:
	//
	// SQL
	JobType *string `json:"jobType,omitempty" xml:"jobType,omitempty"`
	// This parameter is not used.
	//
	// example:
	//
	// -1
	MaxCpuPct *float64 `json:"maxCpuPct,omitempty" xml:"maxCpuPct,omitempty"`
	// This parameter is not used.
	//
	// example:
	//
	// -1
	MaxMemoryPct *float64 `json:"maxMemoryPct,omitempty" xml:"maxMemoryPct,omitempty"`
	// The amount of memory requested by the job at the snapshot time, in MB.
	//
	// example:
	//
	// 409600
	MemoryRequest *int64 `json:"memoryRequest,omitempty" xml:"memoryRequest,omitempty"`
	// The memory usage of the job at the snapshot time. Unit: MB.
	//
	// example:
	//
	// 409600
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The memory fulfillment ratio of the job at the snapshot time. This is calculated by dividing the memory usage by the memory request.
	//
	// example:
	//
	// 1
	MemoryUsageToRequestRatio *float64 `json:"memoryUsageToRequestRatio,omitempty" xml:"memoryUsageToRequestRatio,omitempty"`
	// The CPU usage percentage of a subscription job at the snapshot time. This value is calculated by dividing the CPU usage by the sum of the reserved CPU guarantee and the elastic reserved CPU. This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinCpuPct *float64 `json:"minCpuPct,omitempty" xml:"minCpuPct,omitempty"`
	// The memory usage percentage of a subscription job at the observation time. This value is calculated by dividing the memory usage by the sum of the reserved memory guarantee and the elastic reserved memory. This parameter is not available for pay-as-you-go jobs.
	//
	// example:
	//
	// 0.6
	MinMemoryPct *float64 `json:"minMemoryPct,omitempty" xml:"minMemoryPct,omitempty"`
	// The job priority.
	//
	// example:
	//
	// 9
	Priority *int64 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The project name.
	//
	// example:
	//
	// projectA
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the computing quota that the job uses.
	//
	// example:
	//
	// quota_A
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
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The time when the job started running.
	//
	// > The time when the job acquired its first computing resource.
	//
	// example:
	//
	// 1736821805
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The runtime duration, in seconds. This is the duration from when the job started running to the snapshot time. If the job has not started, this parameter is empty.
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
	// The job status.
	//
	// > The status of a snapshot job can only be \\`running\\`.
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the job was submitted.
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
	// The total duration, in seconds. This is the duration from when the job was submitted to the snapshot time.
	//
	// example:
	//
	// 63
	TotalTime *int64 `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	// The waiting duration, in seconds. This is the duration from when the job was submitted to when it started running. If the job has not started, this is the duration from the submission time to the snapshot time.
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

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) GetExtPlatformId() *string {
	return s.ExtPlatformId
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

func (s *ListJobSnapshotInfosResponseBodyDataJobInfoList) SetExtPlatformId(v string) *ListJobSnapshotInfosResponseBodyDataJobInfoList {
	s.ExtPlatformId = &v
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
