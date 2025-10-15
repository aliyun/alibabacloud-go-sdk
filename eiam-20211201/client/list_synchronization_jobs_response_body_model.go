// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSynchronizationJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSynchronizationJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSynchronizationJobsResponseBody
	GetRequestId() *string
	SetSynchronizationJobs(v []*ListSynchronizationJobsResponseBodySynchronizationJobs) *ListSynchronizationJobsResponseBody
	GetSynchronizationJobs() []*ListSynchronizationJobsResponseBodySynchronizationJobs
	SetTotalCount(v int64) *ListSynchronizationJobsResponseBody
	GetTotalCount() *int64
}

type ListSynchronizationJobsResponseBody struct {
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId           *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobs []*ListSynchronizationJobsResponseBodySynchronizationJobs `json:"SynchronizationJobs,omitempty" xml:"SynchronizationJobs,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSynchronizationJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSynchronizationJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSynchronizationJobsResponseBody) GetSynchronizationJobs() []*ListSynchronizationJobsResponseBodySynchronizationJobs {
	return s.SynchronizationJobs
}

func (s *ListSynchronizationJobsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListSynchronizationJobsResponseBody) SetNextToken(v string) *ListSynchronizationJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetRequestId(v string) *ListSynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetSynchronizationJobs(v []*ListSynchronizationJobsResponseBodySynchronizationJobs) *ListSynchronizationJobsResponseBody {
	s.SynchronizationJobs = v
	return s
}

func (s *ListSynchronizationJobsResponseBody) SetTotalCount(v int64) *ListSynchronizationJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSynchronizationJobsResponseBody) Validate() error {
	if s.SynchronizationJobs != nil {
		for _, item := range s.SynchronizationJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSynchronizationJobsResponseBodySynchronizationJobs struct {
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 同步任务方向
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// 同步结束时间
	//
	// example:
	//
	// 1649830226000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 同步任务结果
	Result *ListSynchronizationJobsResponseBodySynchronizationJobsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// 同步开始时间
	//
	// example:
	//
	// 1649830226000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 同步任务状态
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 同步任务ID
	//
	// example:
	//
	// sync_0000347vjovtcf41li0fgsd98gn24q9njxxxxx
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	// 同步目标ID
	//
	// example:
	//
	// idp_my664lwkhpicbyzirog3xxxxx
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// 同步目标类型
	//
	// example:
	//
	// identity_provider
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 同步触发类型
	//
	// example:
	//
	// auto
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobs) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobs) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetDescription() *string {
	return s.Description
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetDirection() *string {
	return s.Direction
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetResult() *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	return s.Result
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetStatus() *string {
	return s.Status
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetTargetId() *string {
	return s.TargetId
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetTargetType() *string {
	return s.TargetType
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetDescription(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Description = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetDirection(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Direction = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetEndTime(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.EndTime = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetResult(v *ListSynchronizationJobsResponseBodySynchronizationJobsResult) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Result = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetStartTime(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.StartTime = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetStatus(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.Status = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetSynchronizationJobId(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.SynchronizationJobId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTargetId(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TargetId = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTargetType(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TargetType = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) SetTriggerType(v string) *ListSynchronizationJobsResponseBodySynchronizationJobs {
	s.TriggerType = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobs) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResult struct {
	// 同步结果错误码
	//
	// example:
	//
	// ErrorCodeNotFound
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 同步结果错误信息描述
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 组成员同步结果统计
	GroupMemberStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics `json:"GroupMemberStatistics,omitempty" xml:"GroupMemberStatistics,omitempty" type:"Struct"`
	// 组同步结果统计
	GroupStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics `json:"GroupStatistics,omitempty" xml:"GroupStatistics,omitempty" type:"Struct"`
	// 组织同步结果统计
	OrganizationalUnitStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics `json:"OrganizationalUnitStatistics,omitempty" xml:"OrganizationalUnitStatistics,omitempty" type:"Struct"`
	// 用户同步结果统计
	UserStatistics *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics `json:"UserStatistics,omitempty" xml:"UserStatistics,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResult) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResult) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetGroupMemberStatistics() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	return s.GroupMemberStatistics
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetGroupStatistics() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	return s.GroupStatistics
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetOrganizationalUnitStatistics() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	return s.OrganizationalUnitStatistics
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) GetUserStatistics() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	return s.UserStatistics
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetErrorCode(v string) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.ErrorCode = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetErrorMessage(v string) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.ErrorMessage = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetGroupMemberStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.GroupMemberStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetGroupStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.GroupStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetOrganizationalUnitStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.OrganizationalUnitStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) SetUserStatistics(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) *ListSynchronizationJobsResponseBodySynchronizationJobsResult {
	s.UserStatistics = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResult) Validate() error {
	if s.GroupMemberStatistics != nil {
		if err := s.GroupMemberStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.GroupStatistics != nil {
		if err := s.GroupStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.OrganizationalUnitStatistics != nil {
		if err := s.OrganizationalUnitStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.UserStatistics != nil {
		if err := s.UserStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetBinded() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	return s.Binded
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetCreated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	return s.Created
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetDeleted() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	return s.Deleted
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetPushed() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	return s.Pushed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetSame() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	return s.Same
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) GetUpdated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	return s.Updated
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics {
	s.Updated = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatistics) Validate() error {
	if s.Binded != nil {
		if err := s.Binded.Validate(); err != nil {
			return err
		}
	}
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
	if s.Pushed != nil {
		if err := s.Pushed.Validate(); err != nil {
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

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupMemberStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetBinded() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	return s.Binded
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetCreated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	return s.Created
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetDeleted() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	return s.Deleted
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetPushed() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	return s.Pushed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetSame() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	return s.Same
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) GetUpdated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	return s.Updated
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics {
	s.Updated = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatistics) Validate() error {
	if s.Binded != nil {
		if err := s.Binded.Validate(); err != nil {
			return err
		}
	}
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
	if s.Pushed != nil {
		if err := s.Pushed.Validate(); err != nil {
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

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultGroupStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetBinded() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	return s.Binded
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetCreated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	return s.Created
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetDeleted() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	return s.Deleted
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetPushed() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	return s.Pushed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetSame() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	return s.Same
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) GetUpdated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	return s.Updated
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics {
	s.Updated = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatistics) Validate() error {
	if s.Binded != nil {
		if err := s.Binded.Validate(); err != nil {
			return err
		}
	}
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
	if s.Pushed != nil {
		if err := s.Pushed.Validate(); err != nil {
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

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultOrganizationalUnitStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics struct {
	// 绑定结果统计
	Binded *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded `json:"Binded,omitempty" xml:"Binded,omitempty" type:"Struct"`
	// 创建结果统计
	Created *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated `json:"Created,omitempty" xml:"Created,omitempty" type:"Struct"`
	// 删除结果统计
	Deleted  *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted  `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Struct"`
	Exported *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported `json:"Exported,omitempty" xml:"Exported,omitempty" type:"Struct"`
	// 推送结果统计
	Pushed *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed `json:"Pushed,omitempty" xml:"Pushed,omitempty" type:"Struct"`
	// 相同结果统计
	Same *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame `json:"Same,omitempty" xml:"Same,omitempty" type:"Struct"`
	// 更新结果统计
	Updated *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated `json:"Updated,omitempty" xml:"Updated,omitempty" type:"Struct"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetBinded() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	return s.Binded
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetCreated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	return s.Created
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetDeleted() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	return s.Deleted
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetExported() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported {
	return s.Exported
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetPushed() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	return s.Pushed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetSame() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	return s.Same
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) GetUpdated() *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	return s.Updated
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetBinded(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Binded = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetCreated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Created = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetDeleted(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Deleted = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetExported(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Exported = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetPushed(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Pushed = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetSame(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Same = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) SetUpdated(v *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics {
	s.Updated = v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatistics) Validate() error {
	if s.Binded != nil {
		if err := s.Binded.Validate(); err != nil {
			return err
		}
	}
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
	if s.Exported != nil {
		if err := s.Exported.Validate(); err != nil {
			return err
		}
	}
	if s.Pushed != nil {
		if err := s.Pushed.Validate(); err != nil {
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

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsBinded) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsCreated) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsDeleted) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported struct {
	Failed  *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total   *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsExported) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsPushed) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsSame) Validate() error {
	return dara.Validate(s)
}

type ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated struct {
	// 失败数目
	//
	// example:
	//
	// 1
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// 跳过数目
	//
	// example:
	//
	// 1
	Skipped *int64 `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	// 成功数目
	//
	// example:
	//
	// 1
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
	// 总共数目
	//
	// example:
	//
	// 3
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) String() string {
	return dara.Prettify(s)
}

func (s ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GoString() string {
	return s.String()
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GetFailed() *int64 {
	return s.Failed
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GetSkipped() *int64 {
	return s.Skipped
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GetSuccess() *int64 {
	return s.Success
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) GetTotal() *int64 {
	return s.Total
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetFailed(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Failed = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetSkipped(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Skipped = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetSuccess(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Success = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) SetTotal(v int64) *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated {
	s.Total = &v
	return s
}

func (s *ListSynchronizationJobsResponseBodySynchronizationJobsResultUserStatisticsUpdated) Validate() error {
	return dara.Validate(s)
}
