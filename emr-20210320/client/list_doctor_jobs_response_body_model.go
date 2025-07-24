// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorJobsResponseBodyData) *ListDoctorJobsResponseBody
	GetData() []*ListDoctorJobsResponseBodyData
	SetMaxResults(v int32) *ListDoctorJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorJobsResponseBody
	GetTotalCount() *int32
}

type ListDoctorJobsResponseBody struct {
	// The information about the jobs.
	Data []*ListDoctorJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDoctorJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponseBody) GetData() []*ListDoctorJobsResponseBodyData {
	return s.Data
}

func (s *ListDoctorJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorJobsResponseBody) SetData(v []*ListDoctorJobsResponseBodyData) *ListDoctorJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorJobsResponseBody) SetMaxResults(v int32) *ListDoctorJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorJobsResponseBody) SetNextToken(v string) *ListDoctorJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorJobsResponseBody) SetRequestId(v string) *ListDoctorJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorJobsResponseBody) SetTotalCount(v int32) *ListDoctorJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorJobsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsResponseBodyData struct {
	// The ID of the job that was submitted to YARN.
	//
	// example:
	//
	// application_1607584549220_*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// tpcds
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The total running time of the job. Unit: milliseconds.
	//
	// example:
	//
	// 242
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	// The final state of the job. Valid values:
	//
	// 	- SUCCEEDED
	//
	// 	- FAILED
	//
	// 	- KILLED
	//
	// 	- ENDED
	//
	// 	- UNDEFINED
	//
	// example:
	//
	// KILLED
	FinalStatus *string `json:"FinalStatus,omitempty" xml:"FinalStatus,omitempty"`
	// The end time of the job. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1683690929000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The time when the job was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1672148400000
	LaunchTime *int64 `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The data about the metrics.
	Metrics *ListDoctorJobsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The YARN queue to which the job was submitted.
	//
	// example:
	//
	// DW
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// The time when the job was submitted. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1675180800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The running state of the job. Valid values:
	//
	// 	- FINISHED
	//
	// 	- FAILED
	//
	// 	- KILLED
	//
	// example:
	//
	// FINISHED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The type of the compute engine.
	//
	// example:
	//
	// SPARK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username that was used to submit the job.
	//
	// example:
	//
	// DW
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDoctorJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListDoctorJobsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListDoctorJobsResponseBodyData) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListDoctorJobsResponseBodyData) GetFinalStatus() *string {
	return s.FinalStatus
}

func (s *ListDoctorJobsResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListDoctorJobsResponseBodyData) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *ListDoctorJobsResponseBodyData) GetMetrics() *ListDoctorJobsResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorJobsResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *ListDoctorJobsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorJobsResponseBodyData) GetState() *string {
	return s.State
}

func (s *ListDoctorJobsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListDoctorJobsResponseBodyData) GetUser() *string {
	return s.User
}

func (s *ListDoctorJobsResponseBodyData) SetAppId(v string) *ListDoctorJobsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetAppName(v string) *ListDoctorJobsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetElapsedTime(v int64) *ListDoctorJobsResponseBodyData {
	s.ElapsedTime = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetFinalStatus(v string) *ListDoctorJobsResponseBodyData {
	s.FinalStatus = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetFinishTime(v int64) *ListDoctorJobsResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetLaunchTime(v int64) *ListDoctorJobsResponseBodyData {
	s.LaunchTime = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetMetrics(v *ListDoctorJobsResponseBodyDataMetrics) *ListDoctorJobsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetQueue(v string) *ListDoctorJobsResponseBodyData {
	s.Queue = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetStartTime(v int64) *ListDoctorJobsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetState(v string) *ListDoctorJobsResponseBodyData {
	s.State = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetType(v string) *ListDoctorJobsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) SetUser(v string) *ListDoctorJobsResponseBodyData {
	s.User = &v
	return s
}

func (s *ListDoctorJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsResponseBodyDataMetrics struct {
	// The amount of memory consumed.
	MemSeconds *ListDoctorJobsResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The CPU usage.
	VcoreSeconds *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
}

func (s ListDoctorJobsResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponseBodyDataMetrics) GetMemSeconds() *ListDoctorJobsResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *ListDoctorJobsResponseBodyDataMetrics) GetVcoreSeconds() *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *ListDoctorJobsResponseBodyDataMetrics) SetMemSeconds(v *ListDoctorJobsResponseBodyDataMetricsMemSeconds) *ListDoctorJobsResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetrics) SetVcoreSeconds(v *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) *ListDoctorJobsResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsResponseBodyDataMetricsMemSeconds struct {
	// The description of the metric.
	//
	// example:
	//
	// Total memory usage over time in seconds
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memSeconds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// MB 	- Sec
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 12312312
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorJobsResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) SetDescription(v string) *ListDoctorJobsResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) SetName(v string) *ListDoctorJobsResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) SetUnit(v string) *ListDoctorJobsResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) SetValue(v int64) *ListDoctorJobsResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsResponseBodyDataMetricsVcoreSeconds struct {
	// The description of the metric.
	//
	// example:
	//
	// Total vcore usage over time in seconds
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// vcoreSeconds
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// VCores 	- Sec
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 11123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) SetName(v string) *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorJobsResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}
