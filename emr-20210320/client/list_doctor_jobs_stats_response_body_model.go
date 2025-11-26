// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorJobsStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListDoctorJobsStatsResponseBodyData) *ListDoctorJobsStatsResponseBody
	GetData() []*ListDoctorJobsStatsResponseBodyData
	SetMaxResults(v int32) *ListDoctorJobsStatsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorJobsStatsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDoctorJobsStatsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDoctorJobsStatsResponseBody
	GetTotalCount() *int32
}

type ListDoctorJobsStatsResponseBody struct {
	// The summary of job information.
	Data []*ListDoctorJobsStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListDoctorJobsStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponseBody) GetData() []*ListDoctorJobsStatsResponseBodyData {
	return s.Data
}

func (s *ListDoctorJobsStatsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorJobsStatsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorJobsStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDoctorJobsStatsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDoctorJobsStatsResponseBody) SetData(v []*ListDoctorJobsStatsResponseBodyData) *ListDoctorJobsStatsResponseBody {
	s.Data = v
	return s
}

func (s *ListDoctorJobsStatsResponseBody) SetMaxResults(v int32) *ListDoctorJobsStatsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBody) SetNextToken(v string) *ListDoctorJobsStatsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBody) SetRequestId(v string) *ListDoctorJobsStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBody) SetTotalCount(v int32) *ListDoctorJobsStatsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBody) Validate() error {
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

type ListDoctorJobsStatsResponseBodyData struct {
	// The total number of jobs.
	AppsCount *ListDoctorJobsStatsResponseBodyDataAppsCount `json:"AppsCount,omitempty" xml:"AppsCount,omitempty" type:"Struct"`
	// The aggregated amount of memory that is allocated to the job multiplied by the number of seconds the job has been running.
	MemSeconds *ListDoctorJobsStatsResponseBodyDataMemSeconds `json:"MemSeconds,omitempty" xml:"MemSeconds,omitempty" type:"Struct"`
	// The YARN queue to which the job was submitted.
	//
	// example:
	//
	// DW
	Queue *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
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
	// The aggregated number of vCPUs that are allocated to the job multiplied by the number of seconds the job has been running.
	VcoreSeconds *ListDoctorJobsStatsResponseBodyDataVcoreSeconds `json:"VcoreSeconds,omitempty" xml:"VcoreSeconds,omitempty" type:"Struct"`
}

func (s ListDoctorJobsStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponseBodyData) GetAppsCount() *ListDoctorJobsStatsResponseBodyDataAppsCount {
	return s.AppsCount
}

func (s *ListDoctorJobsStatsResponseBodyData) GetMemSeconds() *ListDoctorJobsStatsResponseBodyDataMemSeconds {
	return s.MemSeconds
}

func (s *ListDoctorJobsStatsResponseBodyData) GetQueue() *string {
	return s.Queue
}

func (s *ListDoctorJobsStatsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListDoctorJobsStatsResponseBodyData) GetUser() *string {
	return s.User
}

func (s *ListDoctorJobsStatsResponseBodyData) GetVcoreSeconds() *ListDoctorJobsStatsResponseBodyDataVcoreSeconds {
	return s.VcoreSeconds
}

func (s *ListDoctorJobsStatsResponseBodyData) SetAppsCount(v *ListDoctorJobsStatsResponseBodyDataAppsCount) *ListDoctorJobsStatsResponseBodyData {
	s.AppsCount = v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) SetMemSeconds(v *ListDoctorJobsStatsResponseBodyDataMemSeconds) *ListDoctorJobsStatsResponseBodyData {
	s.MemSeconds = v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) SetQueue(v string) *ListDoctorJobsStatsResponseBodyData {
	s.Queue = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) SetType(v string) *ListDoctorJobsStatsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) SetUser(v string) *ListDoctorJobsStatsResponseBodyData {
	s.User = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) SetVcoreSeconds(v *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) *ListDoctorJobsStatsResponseBodyData {
	s.VcoreSeconds = v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyData) Validate() error {
	if s.AppsCount != nil {
		if err := s.AppsCount.Validate(); err != nil {
			return err
		}
	}
	if s.MemSeconds != nil {
		if err := s.MemSeconds.Validate(); err != nil {
			return err
		}
	}
	if s.VcoreSeconds != nil {
		if err := s.VcoreSeconds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDoctorJobsStatsResponseBodyDataAppsCount struct {
	// The description of the metric.
	//
	// example:
	//
	// Total apps count
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric.
	//
	// example:
	//
	// appsCount
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
	// 123
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDoctorJobsStatsResponseBodyDataAppsCount) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponseBodyDataAppsCount) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) GetName() *string {
	return s.Name
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) SetDescription(v string) *ListDoctorJobsStatsResponseBodyDataAppsCount {
	s.Description = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) SetName(v string) *ListDoctorJobsStatsResponseBodyDataAppsCount {
	s.Name = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) SetUnit(v string) *ListDoctorJobsStatsResponseBodyDataAppsCount {
	s.Unit = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) SetValue(v int64) *ListDoctorJobsStatsResponseBodyDataAppsCount {
	s.Value = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataAppsCount) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsStatsResponseBodyDataMemSeconds struct {
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

func (s ListDoctorJobsStatsResponseBodyDataMemSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponseBodyDataMemSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) SetDescription(v string) *ListDoctorJobsStatsResponseBodyDataMemSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) SetName(v string) *ListDoctorJobsStatsResponseBodyDataMemSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) SetUnit(v string) *ListDoctorJobsStatsResponseBodyDataMemSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) SetValue(v int64) *ListDoctorJobsStatsResponseBodyDataMemSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataMemSeconds) Validate() error {
	return dara.Validate(s)
}

type ListDoctorJobsStatsResponseBodyDataVcoreSeconds struct {
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

func (s ListDoctorJobsStatsResponseBodyDataVcoreSeconds) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorJobsStatsResponseBodyDataVcoreSeconds) GoString() string {
	return s.String()
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) GetDescription() *string {
	return s.Description
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) GetName() *string {
	return s.Name
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) GetUnit() *string {
	return s.Unit
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) GetValue() *int64 {
	return s.Value
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) SetDescription(v string) *ListDoctorJobsStatsResponseBodyDataVcoreSeconds {
	s.Description = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) SetName(v string) *ListDoctorJobsStatsResponseBodyDataVcoreSeconds {
	s.Name = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) SetUnit(v string) *ListDoctorJobsStatsResponseBodyDataVcoreSeconds {
	s.Unit = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) SetValue(v int64) *ListDoctorJobsStatsResponseBodyDataVcoreSeconds {
	s.Value = &v
	return s
}

func (s *ListDoctorJobsStatsResponseBodyDataVcoreSeconds) Validate() error {
	return dara.Validate(s)
}
