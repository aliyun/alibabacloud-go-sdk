// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDataComponentPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMetaDataComponentPageResponseBodyData) *ListMetaDataComponentPageResponseBody
	GetData() []*ListMetaDataComponentPageResponseBodyData
	SetErrCode(v string) *ListMetaDataComponentPageResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListMetaDataComponentPageResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *ListMetaDataComponentPageResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListMetaDataComponentPageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMetaDataComponentPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaDataComponentPageResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListMetaDataComponentPageResponseBody
	GetTotalCount() *int32
}

type ListMetaDataComponentPageResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child parameters.
	Data []*ListMetaDataComponentPageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which is the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed. Check errCode and errMessage for details.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the query conditions. This parameter is used for pagination.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMetaDataComponentPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDataComponentPageResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaDataComponentPageResponseBody) GetData() []*ListMetaDataComponentPageResponseBodyData {
	return s.Data
}

func (s *ListMetaDataComponentPageResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListMetaDataComponentPageResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListMetaDataComponentPageResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListMetaDataComponentPageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaDataComponentPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaDataComponentPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaDataComponentPageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMetaDataComponentPageResponseBody) SetData(v []*ListMetaDataComponentPageResponseBodyData) *ListMetaDataComponentPageResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetErrCode(v string) *ListMetaDataComponentPageResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetErrMessage(v string) *ListMetaDataComponentPageResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetPageIndex(v int32) *ListMetaDataComponentPageResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetPageSize(v int32) *ListMetaDataComponentPageResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetRequestId(v string) *ListMetaDataComponentPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetSuccess(v bool) *ListMetaDataComponentPageResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) SetTotalCount(v int32) *ListMetaDataComponentPageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaDataComponentPageResponseBodyData struct {
	// The entry component type. In some operations, this parameter is used as a backward compatible field for version 1.1.0. Valid values:
	//
	// - 0: source
	//
	// - 1: destination
	//
	// example:
	//
	// 0
	ComponentType *int64 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The creation time of the table or partition.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The datasource config in JSON string format. The structure is defined by each dsType. Parse the JSON string before use. Sensitive fields such as tokens are masked in the response.
	//
	// example:
	//
	// {"endpoint":"...","token":"******"}
	DsConfig *string `json:"dsConfig,omitempty" xml:"dsConfig,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// Data source description
	DsDesc *string `json:"dsDesc,omitempty" xml:"dsDesc,omitempty"`
	// The business ID of the data source (external ID, which may be the same as the primary key ID).
	//
	// example:
	//
	// 290
	DsId *string `json:"dsId,omitempty" xml:"dsId,omitempty"`
	// The data source name. Exact match and fuzzy match are supported.
	//
	// example:
	//
	// test_ds318_hangzhou_0428
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
	// The connectivity status of the data source. Valid values:
	//
	// - 0: Not tested.
	//
	// - 1: Connected.
	//
	// - 2: Connection failed.
	//
	// - -1: Connectivity test not supported.
	//
	// example:
	//
	// 1
	DsStatus *int32 `json:"dsStatus,omitempty" xml:"dsStatus,omitempty"`
	// The data source type, such as Hive or MaxCompute.
	//
	// example:
	//
	// Hive
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The version number of the data source.
	//
	// example:
	//
	// 3.2.0
	DsVersion *string `json:"dsVersion,omitempty" xml:"dsVersion,omitempty"`
	// Indicates whether the data source has expired. Valid values:
	//
	// - true: Expired.
	//
	// - false: Not expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"expired,omitempty" xml:"expired,omitempty"`
	// The primary key ID that uniquely identifies a record.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The profiling task information, including the task status, scheduling ID, profiling rule, and profiling type. This field is empty if the data source is not associated with a profiling task.
	ProfilingJob *ListMetaDataComponentPageResponseBodyDataProfilingJob `json:"profilingJob,omitempty" xml:"profilingJob,omitempty" type:"Struct"`
}

func (s ListMetaDataComponentPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDataComponentPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaDataComponentPageResponseBodyData) GetComponentType() *int64 {
	return s.ComponentType
}

func (s *ListMetaDataComponentPageResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsConfig() *string {
	return s.DsConfig
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsDesc() *string {
	return s.DsDesc
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsId() *string {
	return s.DsId
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsName() *string {
	return s.DsName
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsStatus() *int32 {
	return s.DsStatus
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsType() *string {
	return s.DsType
}

func (s *ListMetaDataComponentPageResponseBodyData) GetDsVersion() *string {
	return s.DsVersion
}

func (s *ListMetaDataComponentPageResponseBodyData) GetExpired() *bool {
	return s.Expired
}

func (s *ListMetaDataComponentPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMetaDataComponentPageResponseBodyData) GetProfilingJob() *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	return s.ProfilingJob
}

func (s *ListMetaDataComponentPageResponseBodyData) SetComponentType(v int64) *ListMetaDataComponentPageResponseBodyData {
	s.ComponentType = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetCreateTime(v string) *ListMetaDataComponentPageResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsConfig(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsConfig = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsDesc(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsDesc = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsId(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsId = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsName(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsName = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsStatus(v int32) *ListMetaDataComponentPageResponseBodyData {
	s.DsStatus = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsType(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsType = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetDsVersion(v string) *ListMetaDataComponentPageResponseBodyData {
	s.DsVersion = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetExpired(v bool) *ListMetaDataComponentPageResponseBodyData {
	s.Expired = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetId(v int64) *ListMetaDataComponentPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) SetProfilingJob(v *ListMetaDataComponentPageResponseBodyDataProfilingJob) *ListMetaDataComponentPageResponseBodyData {
	s.ProfilingJob = v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyData) Validate() error {
	if s.ProfilingJob != nil {
		if err := s.ProfilingJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetaDataComponentPageResponseBodyDataProfilingJob struct {
	// The component ID, which is the primary key of the data source component.
	//
	// example:
	//
	// 12345
	ComponentId *int64 `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// The creation time of the table or partition.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The primary key ID that uniquely identifies a record.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The description of the profiling task.
	//
	// example:
	//
	// Profiling task description
	JobDesc *string `json:"jobDesc,omitempty" xml:"jobDesc,omitempty"`
	// The name of the profiling task.
	//
	// example:
	//
	// job_name
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The ID of the most recent profiling task batch.
	//
	// example:
	//
	// 20001
	LastBatchId *string `json:"lastBatchId,omitempty" xml:"lastBatchId,omitempty"`
	// The profiling task status. Valid values:
	//
	// - 0: Not started.
	//
	// - 1: Running.
	//
	// - 2: Stopped.
	//
	// example:
	//
	// 1
	ProfilingEnable *int32 `json:"profilingEnable,omitempty" xml:"profilingEnable,omitempty"`
	// The profiling permission. Valid values:
	//
	// - 0: read-only link
	//
	// - 1: client
	//
	// example:
	//
	// 0
	ProfilingPermission *int32 `json:"profilingPermission,omitempty" xml:"profilingPermission,omitempty"`
	// The cron expression for scheduled profiling. This parameter takes effect only when profilingType is set to CRON.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	ProfilingRule *string `json:"profilingRule,omitempty" xml:"profilingRule,omitempty"`
	// The profiling policy (scheduling type). Valid values:
	//
	// - 0: daily
	//
	// - 1: CRON
	//
	// example:
	//
	// 0
	ProfilingType *int32 `json:"profilingType,omitempty" xml:"profilingType,omitempty"`
	// The scheduling ID, which uniquely identifies the profiling task on the scheduling side.
	//
	// example:
	//
	// 12345
	SchedulerToken *string `json:"schedulerToken,omitempty" xml:"schedulerToken,omitempty"`
}

func (s ListMetaDataComponentPageResponseBodyDataProfilingJob) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDataComponentPageResponseBodyDataProfilingJob) GoString() string {
	return s.String()
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetComponentId() *int64 {
	return s.ComponentId
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetId() *int64 {
	return s.Id
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetJobDesc() *string {
	return s.JobDesc
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetJobName() *string {
	return s.JobName
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetLastBatchId() *string {
	return s.LastBatchId
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetProfilingEnable() *int32 {
	return s.ProfilingEnable
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetProfilingPermission() *int32 {
	return s.ProfilingPermission
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetProfilingRule() *string {
	return s.ProfilingRule
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetProfilingType() *int32 {
	return s.ProfilingType
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) GetSchedulerToken() *string {
	return s.SchedulerToken
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetComponentId(v int64) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.ComponentId = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetCreateTime(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.CreateTime = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetId(v int64) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.Id = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetJobDesc(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.JobDesc = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetJobName(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.JobName = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetLastBatchId(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.LastBatchId = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetProfilingEnable(v int32) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.ProfilingEnable = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetProfilingPermission(v int32) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.ProfilingPermission = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetProfilingRule(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.ProfilingRule = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetProfilingType(v int32) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.ProfilingType = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) SetSchedulerToken(v string) *ListMetaDataComponentPageResponseBodyDataProfilingJob {
	s.SchedulerToken = &v
	return s
}

func (s *ListMetaDataComponentPageResponseBodyDataProfilingJob) Validate() error {
	return dara.Validate(s)
}
