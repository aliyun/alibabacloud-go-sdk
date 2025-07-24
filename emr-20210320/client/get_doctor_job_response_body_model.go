// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorJobResponseBodyData) *GetDoctorJobResponseBody
	GetData() *GetDoctorJobResponseBodyData
	SetRequestId(v string) *GetDoctorJobResponseBody
	GetRequestId() *string
}

type GetDoctorJobResponseBody struct {
	// The information about the job.
	Data *GetDoctorJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponseBody) GetData() *GetDoctorJobResponseBodyData {
	return s.Data
}

func (s *GetDoctorJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorJobResponseBody) SetData(v *GetDoctorJobResponseBodyData) *GetDoctorJobResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorJobResponseBody) SetRequestId(v string) *GetDoctorJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorJobResponseBodyData struct {
	// The ID of the job that was submitted to YARN.
	//
	// example:
	//
	// application_1542620905989_****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// CREATE TABLE test...ranks=1 (Stage-1)
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The total running time of the job. Unit: milliseconds.
	//
	// example:
	//
	// 278
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
	// The time when the job ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1658388322000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The time when the job was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1665014400000
	LaunchTime *int64 `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	// The data about metrics.
	Metrics *GetDoctorJobResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
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
	// 1673946000000
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

func (s GetDoctorJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetDoctorJobResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetDoctorJobResponseBodyData) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *GetDoctorJobResponseBodyData) GetFinalStatus() *string {
	return s.FinalStatus
}

func (s *GetDoctorJobResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetDoctorJobResponseBodyData) GetLaunchTime() *int64 {
	return s.LaunchTime
}

func (s *GetDoctorJobResponseBodyData) GetMetrics() *GetDoctorJobResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorJobResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *GetDoctorJobResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDoctorJobResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetDoctorJobResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetDoctorJobResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetDoctorJobResponseBodyData) SetAppId(v string) *GetDoctorJobResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetAppName(v string) *GetDoctorJobResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetElapsedTime(v int64) *GetDoctorJobResponseBodyData {
	s.ElapsedTime = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetFinalStatus(v string) *GetDoctorJobResponseBodyData {
	s.FinalStatus = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetFinishTime(v int64) *GetDoctorJobResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetLaunchTime(v int64) *GetDoctorJobResponseBodyData {
	s.LaunchTime = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetMetrics(v *GetDoctorJobResponseBodyDataMetrics) *GetDoctorJobResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetQueue(v string) *GetDoctorJobResponseBodyData {
	s.Queue = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetStartTime(v int64) *GetDoctorJobResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetState(v string) *GetDoctorJobResponseBodyData {
	s.State = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetType(v string) *GetDoctorJobResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) SetUser(v string) *GetDoctorJobResponseBodyData {
	s.User = &v
	return s
}

func (s *GetDoctorJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorJobResponseBodyDataMetrics struct {
	// The amount of memory consumed.
	MemSeconds *GetDoctorJobResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The CPU usage.
	VcoreSeconds *GetDoctorJobResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
}

func (s GetDoctorJobResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponseBodyDataMetrics) GetMemSeconds() *GetDoctorJobResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *GetDoctorJobResponseBodyDataMetrics) GetVcoreSeconds() *GetDoctorJobResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *GetDoctorJobResponseBodyDataMetrics) SetMemSeconds(v *GetDoctorJobResponseBodyDataMetricsMemSeconds) *GetDoctorJobResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetrics) SetVcoreSeconds(v *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) *GetDoctorJobResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorJobResponseBodyDataMetricsMemSeconds struct {
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

func (s GetDoctorJobResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) SetDescription(v string) *GetDoctorJobResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) SetName(v string) *GetDoctorJobResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) SetUnit(v string) *GetDoctorJobResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) SetValue(v int64) *GetDoctorJobResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type GetDoctorJobResponseBodyDataMetricsVcoreSeconds struct {
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

func (s GetDoctorJobResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorJobResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *GetDoctorJobResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) SetName(v string) *GetDoctorJobResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *GetDoctorJobResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *GetDoctorJobResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorJobResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}
