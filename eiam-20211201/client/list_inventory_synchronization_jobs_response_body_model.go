// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventorySynchronizationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInventorySynchronizationJobs(v []*ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) *ListInventorySynchronizationJobsResponseBody
	GetInventorySynchronizationJobs() []*ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs
	SetMaxResults(v int32) *ListInventorySynchronizationJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInventorySynchronizationJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInventorySynchronizationJobsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInventorySynchronizationJobsResponseBody
	GetTotalCount() *int32
}

type ListInventorySynchronizationJobsResponseBody struct {
	// The inventory synchronization jobs.
	InventorySynchronizationJobs []*ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs `json:"InventorySynchronizationJobs,omitempty" xml:"InventorySynchronizationJobs,omitempty" type:"Repeated"`
	// The maximum number of returned results.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned by the current call, which is used to retrieve the next page.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBody) GetInventorySynchronizationJobs() []*ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	return s.InventorySynchronizationJobs
}

func (s *ListInventorySynchronizationJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInventorySynchronizationJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInventorySynchronizationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInventorySynchronizationJobsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInventorySynchronizationJobsResponseBody) SetInventorySynchronizationJobs(v []*ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) *ListInventorySynchronizationJobsResponseBody {
	s.InventorySynchronizationJobs = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBody) SetMaxResults(v int32) *ListInventorySynchronizationJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBody) SetNextToken(v string) *ListInventorySynchronizationJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBody) SetRequestId(v string) *ListInventorySynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBody) SetTotalCount(v int32) *ListInventorySynchronizationJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBody) Validate() error {
	if s.InventorySynchronizationJobs != nil {
		for _, item := range s.InventorySynchronizationJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs struct {
	// The description of the synchronization job.
	//
	// example:
	//
	// No Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the synchronization job.
	//
	// example:
	//
	// 1761533214
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the synchronization targets.
	//
	// example:
	//
	// src_xxxxx
	InventorySourceId *string `json:"InventorySourceId,omitempty" xml:"InventorySourceId,omitempty"`
	// The ID of the synchronization job.
	//
	// example:
	//
	// sync_xxxxxxx
	InventorySynchronizationJobId *string `json:"InventorySynchronizationJobId,omitempty" xml:"InventorySynchronizationJobId,omitempty"`
	// The task result of the synchronization node.
	Result *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The start time of the synchronization job.
	//
	// example:
	//
	// 1761531256
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the synchronization job.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trigger type of the synchronization job.
	//
	// example:
	//
	// manual
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetDescription() *string {
	return s.Description
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetInventorySourceId() *string {
	return s.InventorySourceId
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetInventorySynchronizationJobId() *string {
	return s.InventorySynchronizationJobId
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetResult() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult {
	return s.Result
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetStatus() *string {
	return s.Status
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetDescription(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.Description = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetEndTime(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.EndTime = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetInventorySourceId(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.InventorySourceId = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetInventorySynchronizationJobId(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.InventorySynchronizationJobId = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetResult(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.Result = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetStartTime(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.StartTime = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetStatus(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.Status = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) SetTriggerType(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs {
	s.TriggerType = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobs) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult struct {
	// The statistics of the agent synchronization results.
	AgentStatistics *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics `json:"AgentStatistics,omitempty" xml:"AgentStatistics,omitempty" type:"Struct"`
	// The statistics of the asset synchronization results.
	AssetStatistics *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics `json:"AssetStatistics,omitempty" xml:"AssetStatistics,omitempty" type:"Struct"`
	// The error code of the synchronization result.
	//
	// example:
	//
	// No ErrorCode
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message of the synchronization result.
	//
	// example:
	//
	// No Error Message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) GetAgentStatistics() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics {
	return s.AgentStatistics
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) GetAssetStatistics() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics {
	return s.AssetStatistics
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) SetAgentStatistics(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult {
	s.AgentStatistics = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) SetAssetStatistics(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult {
	s.AssetStatistics = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) SetErrorCode(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult {
	s.ErrorCode = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) SetErrorMessage(v string) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResult) Validate() error {
	if s.AgentStatistics != nil {
		if err := s.AgentStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.AssetStatistics != nil {
		if err := s.AssetStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics struct {
	// The statistics of the creation results.
	Created *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The statistics of the deletion results.
	Deleted *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The statistics of the unchanged results.
	Same *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The statistics of the update results.
	Updated *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) GetCreated() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated {
	return s.Created
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) GetDeleted() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted {
	return s.Deleted
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) GetSame() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame {
	return s.Same
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) GetUpdated() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated {
	return s.Updated
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) SetCreated(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics {
	s.Created = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) SetDeleted(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics {
	s.Deleted = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) SetSame(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics {
	s.Same = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) SetUpdated(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics {
	s.Updated = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatistics) Validate() error {
	if s.Created != nil {
		if err := s.Created.Validate(); err != nil {
			return err
		}
	}
	if s.Deleted != nil {
		if err := s.Deleted.Validate(); err != nil {
			return err
		}
	}
	if s.Same != nil {
		if err := s.Same.Validate(); err != nil {
			return err
		}
	}
	if s.Updated != nil {
		if err := s.Updated.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAgentStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics struct {
	// The statistics of the creation results.
	Created *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The statistics of the deletion results.
	Deleted *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The statistics of the unchanged results.
	Same *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The statistics of the update results.
	Updated *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) GetCreated() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated {
	return s.Created
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) GetDeleted() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted {
	return s.Deleted
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) GetSame() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame {
	return s.Same
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) GetUpdated() *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated {
	return s.Updated
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) SetCreated(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics {
	s.Created = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) SetDeleted(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics {
	s.Deleted = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) SetSame(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics {
	s.Same = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) SetUpdated(v *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics {
	s.Updated = v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatistics) Validate() error {
	if s.Created != nil {
		if err := s.Created.Validate(); err != nil {
			return err
		}
	}
	if s.Deleted != nil {
		if err := s.Deleted.Validate(); err != nil {
			return err
		}
	}
	if s.Same != nil {
		if err := s.Same.Validate(); err != nil {
			return err
		}
	}
	if s.Updated != nil {
		if err := s.Updated.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated struct {
	// The number of failed creations.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped creations.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful creations.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) SetFailed(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) SetSkipped(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) SetSuccess(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) SetTotal(v int64) *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListInventorySynchronizationJobsResponseBodyInventorySynchronizationJobsResultAssetStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}
