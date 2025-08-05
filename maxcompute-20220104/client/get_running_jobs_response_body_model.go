// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetRunningJobsResponseBodyData) *GetRunningJobsResponseBody
	GetData() *GetRunningJobsResponseBodyData
	SetErrorCode(v string) *GetRunningJobsResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetRunningJobsResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetRunningJobsResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetRunningJobsResponseBody
	GetRequestId() *string
}

type GetRunningJobsResponseBody struct {
	// The returned data.
	Data *GetRunningJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 0A3B1FD2006A24C8D8BE65CDAC028298
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ab16684833172127321e2c25
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRunningJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBody) GetData() *GetRunningJobsResponseBodyData {
	return s.Data
}

func (s *GetRunningJobsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetRunningJobsResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetRunningJobsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetRunningJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRunningJobsResponseBody) SetData(v *GetRunningJobsResponseBodyData) *GetRunningJobsResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorCode(v string) *GetRunningJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorMsg(v string) *GetRunningJobsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetHttpCode(v int32) *GetRunningJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetRequestId(v string) *GetRunningJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRunningJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRunningJobsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of jobs in the running state.
	RunningJobInfoList []*GetRunningJobsResponseBodyDataRunningJobInfoList `json:"runningJobInfoList,omitempty" xml:"runningJobInfoList,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 64
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetRunningJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetRunningJobsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRunningJobsResponseBodyData) GetRunningJobInfoList() []*GetRunningJobsResponseBodyDataRunningJobInfoList {
	return s.RunningJobInfoList
}

func (s *GetRunningJobsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetRunningJobsResponseBodyData) SetPageNumber(v int64) *GetRunningJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetPageSize(v int64) *GetRunningJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetRunningJobInfoList(v []*GetRunningJobsResponseBodyDataRunningJobInfoList) *GetRunningJobsResponseBodyData {
	s.RunningJobInfoList = v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetTotalCount(v int64) *GetRunningJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRunningJobsResponseBodyDataRunningJobInfoList struct {
	// The compute unit (CU) snapshot proportion of the job.
	//
	// example:
	//
	// 0.45
	CuSnapshot *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 2023050206371544gomgtp3ljcr4
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that submits the job.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The memory snapshot proportion of the job.
	//
	// example:
	//
	// 0.45
	MemorySnapshot *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	// The progress of the job.
	//
	// example:
	//
	// 0
	Progress *float64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The name of the MaxCompute project.
	//
	// example:
	//
	// project_20221021123044_981b
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the quota that is used by the job.
	//
	// example:
	//
	// my_quota
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The time when the job starts to run.
	//
	// example:
	//
	// 1689746864
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The time when the job is submitted.
	//
	// example:
	//
	// 1689746864
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetCuSnapshot() *float64 {
	return s.CuSnapshot
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetJobOwner() *string {
	return s.JobOwner
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetMemorySnapshot() *float64 {
	return s.MemorySnapshot
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetProgress() *float64 {
	return s.Progress
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetProject() *string {
	return s.Project
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetQuotaNickname() *string {
	return s.QuotaNickname
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetRunningAtTime() *int64 {
	return s.RunningAtTime
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) GetSubmittedAtTime() *int64 {
	return s.SubmittedAtTime
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetCuSnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetInstanceId(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetJobOwner(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetMemorySnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProgress(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Progress = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProject(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Project = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetQuotaNickname(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetRunningAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetSubmittedAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) Validate() error {
	return dara.Validate(s)
}
