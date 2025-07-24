// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorApplicationsResponseBodyData) *ListDoctorApplicationsResponseBody
	GetData() []*ListDoctorApplicationsResponseBodyData
	SetMaxResults(v int32) *ListDoctorApplicationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorApplicationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorApplicationsResponseBody
	GetTotalCount() *int32
}

type ListDoctorApplicationsResponseBody struct {
	// The details of jobs.
	Data []*ListDoctorApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListDoctorApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBody) GetData() []*ListDoctorApplicationsResponseBodyData {
	return s.Data
}

func (s *ListDoctorApplicationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorApplicationsResponseBody) SetData(v []*ListDoctorApplicationsResponseBodyData) *ListDoctorApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorApplicationsResponseBody) SetMaxResults(v int32) *ListDoctorApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorApplicationsResponseBody) SetNextToken(v string) *ListDoctorApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorApplicationsResponseBody) SetRequestId(v string) *ListDoctorApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorApplicationsResponseBody) SetTotalCount(v int32) *ListDoctorApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyData struct {
	// The analysis results of the jobs.
	Analysis *ListDoctorApplicationsResponseBodyDataAnalysis `json:"Analysis,omitempty" xml:"Analysis,omitempty" type:"Struct"`
	// The job ID in YARN. The value of QueryID or SessionID is returned for a Hive job.
	//
	// example:
	//
	// application_1665056147236_*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the job.
	//
	// example:
	//
	// data-upload
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the job ended. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1682049088086
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The job IDs. Multiple job IDs are separated with commas (,).
	//
	// example:
	//
	// null
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The data about metrics.
	Metrics *ListDoctorApplicationsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The query statement. This parameter is left empty for non-SQL jobs.
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
	// TEZ
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username that was used to submit the job.
	//
	// example:
	//
	// DW
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDoctorApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyData) GetAnalysis() *ListDoctorApplicationsResponseBodyDataAnalysis {
	return s.Analysis
}

func (s *ListDoctorApplicationsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListDoctorApplicationsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListDoctorApplicationsResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDoctorApplicationsResponseBodyData) GetIds() []*string {
	return s.Ids
}

func (s *ListDoctorApplicationsResponseBodyData) GetMetrics() *ListDoctorApplicationsResponseBodyDataMetrics {
	return s.Metrics
}

func (s *ListDoctorApplicationsResponseBodyData) GetQuerySql() *string {
	return s.QuerySql
}

func (s *ListDoctorApplicationsResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *ListDoctorApplicationsResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDoctorApplicationsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListDoctorApplicationsResponseBodyData) GetUser() *string {
	return s.User
}

func (s *ListDoctorApplicationsResponseBodyData) SetAnalysis(v *ListDoctorApplicationsResponseBodyDataAnalysis) *ListDoctorApplicationsResponseBodyData {
	s.Analysis = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetAppId(v string) *ListDoctorApplicationsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetAppName(v string) *ListDoctorApplicationsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetEndTime(v int64) *ListDoctorApplicationsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetIds(v []*string) *ListDoctorApplicationsResponseBodyData {
	s.Ids = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetMetrics(v *ListDoctorApplicationsResponseBodyDataMetrics) *ListDoctorApplicationsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetQuerySql(v string) *ListDoctorApplicationsResponseBodyData {
	s.QuerySql = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetQueue(v string) *ListDoctorApplicationsResponseBodyData {
	s.Queue = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetStartTime(v int64) *ListDoctorApplicationsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetType(v string) *ListDoctorApplicationsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) SetUser(v string) *ListDoctorApplicationsResponseBodyData {
	s.User = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataAnalysis struct {
	// The score of the job.
	//
	// example:
	//
	// 23
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The suggestion on executing the job.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s ListDoctorApplicationsResponseBodyDataAnalysis) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataAnalysis) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataAnalysis) GetScore() *int32 {
	return s.Score
}

func (s *ListDoctorApplicationsResponseBodyDataAnalysis) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListDoctorApplicationsResponseBodyDataAnalysis) SetScore(v int32) *ListDoctorApplicationsResponseBodyDataAnalysis {
	s.Score = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataAnalysis) SetSuggestion(v string) *ListDoctorApplicationsResponseBodyDataAnalysis {
	s.Suggestion = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataAnalysis) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataMetrics struct {
	// The amount of memory consumed.
	MemSeconds *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The memory usage
	MemUtilization *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization `json:"MemUtilization,omitempty" xml:"MemUtilization,omitempty" type:"Struct"`
	// The CPU usage.
	VcoreSeconds *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
	// The CPU utilization. This parameter has the same meaning as %CPU in the Linux top command.
	VcoreUtilization *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization `json:"VcoreUtilization,omitempty" xml:"VcoreUtilization,omitempty" type:"Struct"`
}

func (s ListDoctorApplicationsResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) GetMemSeconds() *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds {
	return s.MemSeconds
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) GetMemUtilization() *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization {
	return s.MemUtilization
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) GetVcoreSeconds() *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds {
	return s.VcoreSeconds
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) GetVcoreUtilization() *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization {
	return s.VcoreUtilization
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) SetMemSeconds(v *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) *ListDoctorApplicationsResponseBodyDataMetrics {
	s.MemSeconds = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) SetMemUtilization(v *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) *ListDoctorApplicationsResponseBodyDataMetrics {
	s.MemUtilization = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) SetVcoreSeconds(v *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) *ListDoctorApplicationsResponseBodyDataMetrics {
	s.VcoreSeconds = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) SetVcoreUtilization(v *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) *ListDoctorApplicationsResponseBodyDataMetrics {
	s.VcoreUtilization = v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataMetricsMemSeconds struct {
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

func (s ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) SetDescription(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) SetName(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) SetUnit(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) SetValue(v int64) *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataMetricsMemUtilization struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 0.23
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) GetName() *string {
	return s.Name
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) SetDescription(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization {
	s.Description = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) SetName(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization {
	s.Name = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) SetUnit(v string) *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization {
	s.Unit = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) SetValue(v float32) *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization {
	s.Value = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsMemUtilization) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds struct {
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

func (s ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) SetDescription(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) SetName(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) SetUnit(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) SetValue(v int64) *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization struct {
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
	// ""
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 23.1
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) GoString() string {
	return s.String()
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) GetName() *string {
	return s.Name
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) GetValue() *float32 {
	return s.Value
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) SetDescription(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization {
	s.Description = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) SetName(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization {
	s.Name = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) SetUnit(v string) *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization {
	s.Unit = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) SetValue(v float32) *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization {
	s.Value = &v
	return s
}

func (s *ListDoctorApplicationsResponseBodyDataMetricsVcoreUtilization) Validate() error {
	return dara.Validate(s)
}
