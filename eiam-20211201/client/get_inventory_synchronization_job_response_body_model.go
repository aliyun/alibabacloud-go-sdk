// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInventorySynchronizationJob(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) *GetInventorySynchronizationJobResponseBody
	GetInventorySynchronizationJob() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob
	SetRequestId(v string) *GetInventorySynchronizationJobResponseBody
	GetRequestId() *string
}

type GetInventorySynchronizationJobResponseBody struct {
	// The inventory synchronization job.
	InventorySynchronizationJob *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob `json:"InventorySynchronizationJob,omitempty" xml:"InventorySynchronizationJob,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBody) GetInventorySynchronizationJob() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	return s.InventorySynchronizationJob
}

func (s *GetInventorySynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInventorySynchronizationJobResponseBody) SetInventorySynchronizationJob(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) *GetInventorySynchronizationJobResponseBody {
	s.InventorySynchronizationJob = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBody) SetRequestId(v string) *GetInventorySynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBody) Validate() error {
	if s.InventorySynchronizationJob != nil {
		if err := s.InventorySynchronizationJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob struct {
	// The synchronization job description.
	//
	// example:
	//
	// No Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The synchronization end time.
	//
	// example:
	//
	// 1761531556
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The synchronization target ID.
	//
	// example:
	//
	// src_xxxxxx
	InventorySourceId *string `json:"InventorySourceId,omitempty" xml:"InventorySourceId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// sync_xxxxxxxxx
	InventorySynchronizationJobId *string `json:"InventorySynchronizationJobId,omitempty" xml:"InventorySynchronizationJobId,omitempty"`
	// The synchronization node task result.
	Result *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The synchronization start time.
	//
	// example:
	//
	// 1761531256
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The synchronization job status.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The synchronization trigger type.
	//
	// example:
	//
	// manual
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetDescription() *string {
	return s.Description
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetInventorySourceId() *string {
	return s.InventorySourceId
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetInventorySynchronizationJobId() *string {
	return s.InventorySynchronizationJobId
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetResult() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult {
	return s.Result
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetStatus() *string {
	return s.Status
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetDescription(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.Description = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetEndTime(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.EndTime = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetInventorySourceId(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.InventorySourceId = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetInventorySynchronizationJobId(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.InventorySynchronizationJobId = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetResult(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.Result = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetStartTime(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.StartTime = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetStatus(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.Status = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) SetTriggerType(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob {
	s.TriggerType = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJob) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult struct {
	// The agent synchronization result statistics.
	AgentStatistics *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics `json:"AgentStatistics,omitempty" xml:"AgentStatistics,omitempty" type:"Struct"`
	// The asset synchronization result statistics.
	AssetStatistics *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics `json:"AssetStatistics,omitempty" xml:"AssetStatistics,omitempty" type:"Struct"`
	// The error code of the synchronization result.
	//
	// example:
	//
	// No ErrorCode
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message description of the synchronization result.
	//
	// example:
	//
	// No Error Message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) GetAgentStatistics() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics {
	return s.AgentStatistics
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) GetAssetStatistics() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics {
	return s.AssetStatistics
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) SetAgentStatistics(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult {
	s.AgentStatistics = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) SetAssetStatistics(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult {
	s.AssetStatistics = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) SetErrorCode(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult {
	s.ErrorCode = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) SetErrorMessage(v string) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResult) Validate() error {
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

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics struct {
	// The creation result statistics.
	Created *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The unchanged result statistics.
	Same *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) GetCreated() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated {
	return s.Created
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) GetDeleted() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted {
	return s.Deleted
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) GetSame() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame {
	return s.Same
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) GetUpdated() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated {
	return s.Updated
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) SetCreated(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics {
	s.Created = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) SetDeleted(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics {
	s.Deleted = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) SetSame(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics {
	s.Same = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) SetUpdated(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics {
	s.Updated = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatistics) Validate() error {
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

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAgentStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics struct {
	// The creation result statistics.
	Created *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The unchanged result statistics.
	Same *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) GetCreated() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated {
	return s.Created
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) GetDeleted() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted {
	return s.Deleted
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) GetSame() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame {
	return s.Same
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) GetUpdated() *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated {
	return s.Updated
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) SetCreated(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics {
	s.Created = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) SetDeleted(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics {
	s.Deleted = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) SetSame(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics {
	s.Same = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) SetUpdated(v *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics {
	s.Updated = v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatistics) Validate() error {
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

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated struct {
	// The failure count.
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The skipped count.
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The success count.
	//
	// example:
	//
	// 3
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total count.
	//
	// example:
	//
	// 5
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) SetFailed(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) SetSkipped(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) SetSuccess(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) SetTotal(v int64) *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetInventorySynchronizationJobResponseBodyInventorySynchronizationJobResultAssetStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}
