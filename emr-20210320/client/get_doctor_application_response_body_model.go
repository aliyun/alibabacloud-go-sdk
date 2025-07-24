// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDoctorApplicationResponseBodyData) *GetDoctorApplicationResponseBody
	GetData() *GetDoctorApplicationResponseBodyData
	SetRequestId(v string) *GetDoctorApplicationResponseBody
	GetRequestId() *string
}

type GetDoctorApplicationResponseBody struct {
	// The details of the job.
	Data *GetDoctorApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDoctorApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBody) GetData() *GetDoctorApplicationResponseBodyData {
	return s.Data
}

func (s *GetDoctorApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDoctorApplicationResponseBody) SetData(v *GetDoctorApplicationResponseBodyData) *GetDoctorApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetDoctorApplicationResponseBody) SetRequestId(v string) *GetDoctorApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDoctorApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyData struct {
	// The job analysis result.
	Analysis *GetDoctorApplicationResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The name of the job.
	//
	// example:
	//
	// CREATE TABLE test...ranks=1 (Stage-1)
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The end time of the job. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1666213200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job IDs. Multiple job IDs are separated with commas (,).
	//
	// example:
	//
	// null
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The metric information.
	Metrics *GetDoctorApplicationResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The SQL statement of the job. This parameter is left empty for non-SQL jobs.
	//
	// example:
	//
	// SELECT id, count(1) FROM test group by id;
	QuerySql *string `json:"QuerySql,omitempty" xml:"QuerySql,omitempty"`
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
	// 1677465658275
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the compute engine.
	//
	// example:
	//
	// SPARK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username that is used to submit the job.
	//
	// example:
	//
	// DW
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetDoctorApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyData) GetAnalysis() *GetDoctorApplicationResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *GetDoctorApplicationResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *GetDoctorApplicationResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDoctorApplicationResponseBodyData) GetIds() []*string {
	return s.Ids
}

func (s *GetDoctorApplicationResponseBodyData) GetMetrics() *GetDoctorApplicationResponseBodyDataMetrics {
	return s.Metrics
}

func (s *GetDoctorApplicationResponseBodyData) GetQuerySql() *string {
	return s.QuerySql
}

func (s *GetDoctorApplicationResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *GetDoctorApplicationResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDoctorApplicationResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetDoctorApplicationResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetDoctorApplicationResponseBodyData) SetAnalysis(v *GetDoctorApplicationResponseBodyDataAnalysis) *GetDoctorApplicationResponseBodyData {
	s.Analysis = v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetAppName(v string) *GetDoctorApplicationResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetEndTime(v int64) *GetDoctorApplicationResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetIds(v []*string) *GetDoctorApplicationResponseBodyData {
	s.Ids = v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetMetrics(v *GetDoctorApplicationResponseBodyDataMetrics) *GetDoctorApplicationResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetQuerySql(v string) *GetDoctorApplicationResponseBodyData {
	s.QuerySql = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetQueue(v string) *GetDoctorApplicationResponseBodyData {
	s.Queue = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetStartTime(v int64) *GetDoctorApplicationResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetType(v string) *GetDoctorApplicationResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) SetUser(v string) *GetDoctorApplicationResponseBodyData {
	s.User = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataAnalysis struct {
	// The score of the job.
	//
	// example:
	//
	// 67
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The suggestion for running the job.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetDoctorApplicationResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataAnalysis) GetScore() *int32 {
	return s.Score
}

func (s *GetDoctorApplicationResponseBodyDataAnalysis) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetDoctorApplicationResponseBodyDataAnalysis) SetScore(v int32) *GetDoctorApplicationResponseBodyDataAnalysis {
	s.Score = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataAnalysis) SetSuggestion(v string) *GetDoctorApplicationResponseBodyDataAnalysis {
	s.Suggestion = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataMetrics struct {
	// The aggregated amount of memory that is allocated to the job multiplied by the number of seconds the job has been running.
	MemSeconds *GetDoctorApplicationResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The memory usage.
	MemUtilization *GetDoctorApplicationResponseBodyDataMetricsMemUtilization `json:"MemUtilization,omitempty" xml:"MemUtilization,omitempty" type:"Struct"`
	// The aggregated number of vCPUs that are allocated to the job multiplied by the number of seconds the job has been running.
	VcoreSeconds *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
	// The CPU utilization. The meaning is the same as that of the %CPU command in the output of the Linux top command.
	VcoreUtilization *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization `json:"VcoreUtilization,omitempty" xml:"VcoreUtilization,omitempty" type:"Struct"`
}

func (s GetDoctorApplicationResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) GetMemSeconds() *GetDoctorApplicationResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) GetMemUtilization() *GetDoctorApplicationResponseBodyDataMetricsMemUtilization {
	return s.MemUtilization
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) GetVcoreSeconds() *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) GetVcoreUtilization() *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization {
	return s.VcoreUtilization
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) SetMemSeconds(v *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) *GetDoctorApplicationResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) SetMemUtilization(v *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) *GetDoctorApplicationResponseBodyDataMetrics {
	s.MemUtilization = v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) SetVcoreSeconds(v *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) *GetDoctorApplicationResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) SetVcoreUtilization(v *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) *GetDoctorApplicationResponseBodyDataMetrics {
	s.VcoreUtilization = v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataMetricsMemSeconds struct {
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

func (s GetDoctorApplicationResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) SetDescription(v string) *GetDoctorApplicationResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) SetName(v string) *GetDoctorApplicationResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) SetUnit(v string) *GetDoctorApplicationResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) SetValue(v int64) *GetDoctorApplicationResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataMetricsMemUtilization struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of used memory to total available memory
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// memUtilization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.82
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorApplicationResponseBodyDataMetricsMemUtilization) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataMetricsMemUtilization) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) GetName() *string {
	return s.Name
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) SetDescription(v string) *GetDoctorApplicationResponseBodyDataMetricsMemUtilization {
	s.Description = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) SetName(v string) *GetDoctorApplicationResponseBodyDataMetricsMemUtilization {
	s.Name = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) SetUnit(v string) *GetDoctorApplicationResponseBodyDataMetricsMemUtilization {
	s.Unit = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) SetValue(v float32) *GetDoctorApplicationResponseBodyDataMetricsMemUtilization {
	s.Value = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsMemUtilization) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds struct {
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

func (s GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) SetName(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}

type GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization struct {
	// The description of the metric.
	//
	// example:
	//
	// Ratio of used vcore to total available cores
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// vcoreUtilization
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// “”
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 32.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) GoString() string {
	return s.String()
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) GetDescription() *string {
	return s.Description
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) GetName() *string {
	return s.Name
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) GetUnit() *string {
	return s.Unit
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) GetValue() *float32 {
	return s.Value
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) SetDescription(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization {
	s.Description = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) SetName(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization {
	s.Name = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) SetUnit(v string) *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization {
	s.Unit = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) SetValue(v float32) *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization {
	s.Value = &v
	return s
}

func (s *GetDoctorApplicationResponseBodyDataMetricsVcoreUtilization) Validate() error {
	return dara.Validate(s)
}
