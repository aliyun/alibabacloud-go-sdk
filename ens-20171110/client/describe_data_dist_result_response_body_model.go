// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDistResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistResults(v *DescribeDataDistResultResponseBodyDistResults) *DescribeDataDistResultResponseBody
	GetDistResults() *DescribeDataDistResultResponseBodyDistResults
	SetPageNumber(v int32) *DescribeDataDistResultResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataDistResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataDistResultResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataDistResultResponseBody
	GetTotalCount() *int32
}

type DescribeDataDistResultResponseBody struct {
	// The distribution status of data files on edge instances.
	DistResults *DescribeDataDistResultResponseBodyDistResults `json:"DistResults,omitempty" xml:"DistResults,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 50373E71-7710-4620-8AAB-133CCE49451C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataDistResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBody) GetDistResults() *DescribeDataDistResultResponseBodyDistResults {
	return s.DistResults
}

func (s *DescribeDataDistResultResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataDistResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataDistResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataDistResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataDistResultResponseBody) SetDistResults(v *DescribeDataDistResultResponseBodyDistResults) *DescribeDataDistResultResponseBody {
	s.DistResults = v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetPageNumber(v int32) *DescribeDataDistResultResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetPageSize(v int32) *DescribeDataDistResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetRequestId(v string) *DescribeDataDistResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) SetTotalCount(v int32) *DescribeDataDistResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataDistResultResponseBody) Validate() error {
	if s.DistResults != nil {
		if err := s.DistResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResults struct {
	DistResult []*DescribeDataDistResultResponseBodyDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResults) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResults) GetDistResult() []*DescribeDataDistResultResponseBodyDistResultsDistResult {
	return s.DistResult
}

func (s *DescribeDataDistResultResponseBodyDistResults) SetDistResult(v []*DescribeDataDistResultResponseBodyDistResultsDistResult) *DescribeDataDistResultResponseBodyDistResults {
	s.DistResult = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResults) Validate() error {
	if s.DistResult != nil {
		for _, item := range s.DistResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResultsDistResult struct {
	// The name of the data file.
	//
	// example:
	//
	// gcs-prod-websocket-eip-unicom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The distribution status statistics.
	StatusStats *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats `json:"StatusStats,omitempty" xml:"StatusStats,omitempty" type:"Struct"`
	// The version number of the data file.
	//
	// example:
	//
	// standard
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) GetName() *string {
	return s.Name
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) GetStatusStats() *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats {
	return s.StatusStats
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) GetVersion() *string {
	return s.Version
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetName(v string) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.Name = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetStatusStats(v *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.StatusStats = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) SetVersion(v string) *DescribeDataDistResultResponseBodyDistResultsDistResult {
	s.Version = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResult) Validate() error {
	if s.StatusStats != nil {
		if err := s.StatusStats.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats struct {
	StatusStat []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat `json:"StatusStat,omitempty" xml:"StatusStat,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) GetStatusStat() []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	return s.StatusStat
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) SetStatusStat(v []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats {
	s.StatusStat = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStats) Validate() error {
	if s.StatusStat != nil {
		for _, item := range s.StatusStat {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat struct {
	// The number of associated edge instances.
	//
	// example:
	//
	// 10
	InstanceCount *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The distribution status of the edge instance.
	Instances *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The distribution status. The value is of the enumeration type. Valid values:
	//
	// 	- SUCCESS: The distribution is successful.
	//
	// 	- FAILED: The distribution failed.
	//
	// 	- DISTING: The data is being distributed.
	//
	// 	- POD_RESTARTING: The idle pod is being restarted.
	//
	// 	- DELETED: The data is cleared or removed.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) GetInstanceCount() *string {
	return s.InstanceCount
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) GetInstances() *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances {
	return s.Instances
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetInstanceCount(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.InstanceCount = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetInstances(v *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.Instances = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) SetStatus(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat {
	s.Status = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStat) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances struct {
	Instance []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) GetInstance() []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	return s.Instance
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) SetInstance(v []*DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances {
	s.Instance = v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance struct {
	// The ID of the instance
	//
	// example:
	//
	// i-5qzje8f5un1wmi341m2yetaxv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the distribution. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-19T07:24:52Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The description of the distribution status.
	//
	// example:
	//
	// SUCCESS
	StatusDescrip *string `json:"StatusDescrip,omitempty" xml:"StatusDescrip,omitempty"`
	// The time when the distribution status was last updated. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-11T02:38:19Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GetStatusDescrip() *string {
	return s.StatusDescrip
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetInstanceId(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStartTime(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetStatusDescrip(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.StatusDescrip = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) SetUpdateTime(v string) *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDataDistResultResponseBodyDistResultsDistResultStatusStatsStatusStatInstancesInstance) Validate() error {
	return dara.Validate(s)
}
