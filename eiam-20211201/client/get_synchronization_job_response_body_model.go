// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSynchronizationJobResponseBody
	GetRequestId() *string
	SetSynchronizationJob(v *GetSynchronizationJobResponseBodySynchronizationJob) *GetSynchronizationJobResponseBody
	GetSynchronizationJob() *GetSynchronizationJobResponseBodySynchronizationJob
}

type GetSynchronizationJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the synchronization job.
	SynchronizationJob *GetSynchronizationJobResponseBodySynchronizationJob `json:"SynchronizationJob,omitempty" xml:"SynchronizationJob,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSynchronizationJobResponseBody) GetSynchronizationJob() *GetSynchronizationJobResponseBodySynchronizationJob {
	return s.SynchronizationJob
}

func (s *GetSynchronizationJobResponseBody) SetRequestId(v string) *GetSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSynchronizationJobResponseBody) SetSynchronizationJob(v *GetSynchronizationJobResponseBodySynchronizationJob) *GetSynchronizationJobResponseBody {
	s.SynchronizationJob = v
	return s
}

func (s *GetSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJob struct {
	// The direction of the synchronization job. Valid values:
	//
	// 	- ingress
	//
	// 	- egress
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end time of the synchronization. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The result of the synchronization job.
	Result *GetSynchronizationJobResponseBodySynchronizationJobResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The start time of the synchronization. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the synchronization job. Valid values:
	//
	// 	- pending
	//
	// 	- running
	//
	// 	- failed
	//
	// 	- partial_success
	//
	// 	- success
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the synchronization job.
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9nj9xxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// The ID of the synchronization destination.
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3nxxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the synchronization destination. Valid values:
	//
	// 	- identity_provider
	//
	// 	- application
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The trigger type of the synchronization. Valid values:
	//
	// 	- auto
	//
	// 	- manual
	//
	// example:
	//
	// auto
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJob) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJob) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetDirection() *string {
	return s.Direction
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetResult() *GetSynchronizationJobResponseBodySynchronizationJobResult {
	return s.Result
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetStatus() *string {
	return s.Status
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetTargetId() *string {
	return s.TargetId
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetTargetType() *string {
	return s.TargetType
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetDirection(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Direction = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetEndTime(v int64) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.EndTime = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetResult(v *GetSynchronizationJobResponseBodySynchronizationJobResult) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Result = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetStartTime(v int64) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.StartTime = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetStatus(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.Status = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetSynchronizationJobId(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.SynchronizationJobId = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTargetId(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TargetId = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTargetType(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TargetType = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) SetTriggerType(v string) *GetSynchronizationJobResponseBodySynchronizationJob {
	s.TriggerType = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJob) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResult struct {
	// The error code corresponding to the error message.
	//
	// example:
	//
	// ResourceNotFound. SynchronizationJob
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned in the case of an error.
	//
	// example:
	//
	// The specified SynchronizationJob resource: %s not found.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The group member synchronization result statistics.
	GroupMemberStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics `json:"GroupMemberStatistics,omitempty" xml:"GroupMemberStatistics,omitempty" type:"Struct"`
	// The group synchronization result statistics.
	GroupStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics `json:"GroupStatistics,omitempty" xml:"GroupStatistics,omitempty" type:"Struct"`
	// The organization synchronization result statistics.
	OrganizationalUnitStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics `json:"OrganizationalUnitStatistics,omitempty" xml:"OrganizationalUnitStatistics,omitempty" type:"Struct"`
	// The user synchronization result statistics.
	UserStatistics *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics `json:"UserStatistics,omitempty" xml:"UserStatistics,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResult) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetGroupMemberStatistics() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	return s.GroupMemberStatistics
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetGroupStatistics() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	return s.GroupStatistics
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetOrganizationalUnitStatistics() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	return s.OrganizationalUnitStatistics
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) GetUserStatistics() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	return s.UserStatistics
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetErrorCode(v string) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.ErrorCode = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetErrorMessage(v string) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetGroupMemberStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.GroupMemberStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetGroupStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.GroupStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetOrganizationalUnitStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.OrganizationalUnitStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) SetUserStatistics(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) *GetSynchronizationJobResponseBodySynchronizationJobResult {
	s.UserStatistics = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResult) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics struct {
	// The binding result statistics.
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// The creation result statistics.
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The notification result statistics.
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// The result statistics about identical group members.
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetBinded() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	return s.Binded
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetCreated() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	return s.Created
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetDeleted() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	return s.Deleted
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetPushed() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	return s.Pushed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetSame() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	return s.Same
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) GetUpdated() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	return s.Updated
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics {
	s.Updated = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatistics) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupMemberStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics struct {
	// The binding result statistics.
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// The creation result statistics.
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The notification result statistics.
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// The result statistics about identical groups.
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetBinded() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	return s.Binded
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetCreated() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	return s.Created
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetDeleted() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	return s.Deleted
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetPushed() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	return s.Pushed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetSame() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	return s.Same
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) GetUpdated() *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	return s.Updated
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics {
	s.Updated = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatistics) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultGroupStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics struct {
	// The binding result statistics.
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// The creation result statistics.
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The notification result statistics.
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// The result statistics about identical organizations.
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetBinded() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	return s.Binded
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetCreated() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	return s.Created
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetDeleted() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	return s.Deleted
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetPushed() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	return s.Pushed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetSame() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	return s.Same
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) GetUpdated() *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	return s.Updated
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics {
	s.Updated = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatistics) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultOrganizationalUnitStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics struct {
	// The binding result statistics.
	Binded *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// The creation result statistics.
	Created *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// The deletion result statistics.
	Deleted *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// The notification result statistics.
	Pushed *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// The result statistics about identical users.
	Same *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// The update result statistics.
	Updated *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetBinded() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	return s.Binded
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetCreated() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	return s.Created
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetDeleted() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	return s.Deleted
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetPushed() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	return s.Pushed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetSame() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	return s.Same
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) GetUpdated() *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	return s.Updated
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetBinded(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Binded = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetCreated(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Created = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetDeleted(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Deleted = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetPushed(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Pushed = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetSame(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Same = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) SetUpdated(v *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics {
	s.Updated = v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatistics) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated struct {
	// The number of failed items.
	//
	// example:
	//
	// 10
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of skipped items.
	//
	// example:
	//
	// 10
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// The number of successful items.
	//
	// example:
	//
	// 10
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetFailed(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetSkipped(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetSuccess(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) SetTotal(v int64) *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *GetSynchronizationJobResponseBodySynchronizationJobResultUserStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}
