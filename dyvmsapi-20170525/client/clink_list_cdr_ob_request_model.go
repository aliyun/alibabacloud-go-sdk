// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrObRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeDuration(v int64) *ClinkListCdrObRequest
	GetBridgeDuration() *int64
	SetBridgeDurationEnd(v int64) *ClinkListCdrObRequest
	GetBridgeDurationEnd() *int64
	SetBridgeTime(v int64) *ClinkListCdrObRequest
	GetBridgeTime() *int64
	SetBridgeTimeEnd(v int64) *ClinkListCdrObRequest
	GetBridgeTimeEnd() *int64
	SetCallType(v int64) *ClinkListCdrObRequest
	GetCallType() *int64
	SetClientNumber(v string) *ClinkListCdrObRequest
	GetClientNumber() *string
	SetCnos(v string) *ClinkListCdrObRequest
	GetCnos() *string
	SetCustomerNumber(v string) *ClinkListCdrObRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *ClinkListCdrObRequest
	GetEndTime() *int64
	SetEndTimeEnd(v int64) *ClinkListCdrObRequest
	GetEndTimeEnd() *int64
	SetEnterpriseId(v int64) *ClinkListCdrObRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkListCdrObRequest
	GetHiddenType() *int64
	SetLeftClid(v string) *ClinkListCdrObRequest
	GetLeftClid() *string
	SetLimit(v int64) *ClinkListCdrObRequest
	GetLimit() *int64
	SetMainUniqueId(v string) *ClinkListCdrObRequest
	GetMainUniqueId() *string
	SetOffset(v int64) *ClinkListCdrObRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListCdrObRequest
	GetOwnerId() *int64
	SetQnos(v string) *ClinkListCdrObRequest
	GetQnos() *string
	SetResourceOwnerAccount(v string) *ClinkListCdrObRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListCdrObRequest
	GetResourceOwnerId() *int64
	SetScrollId(v string) *ClinkListCdrObRequest
	GetScrollId() *string
	SetScrollSearch(v string) *ClinkListCdrObRequest
	GetScrollSearch() *string
	SetStartTime(v int64) *ClinkListCdrObRequest
	GetStartTime() *int64
	SetStartTimeEnd(v int64) *ClinkListCdrObRequest
	GetStartTimeEnd() *int64
	SetStatus(v int64) *ClinkListCdrObRequest
	GetStatus() *int64
	SetTaskName(v string) *ClinkListCdrObRequest
	GetTaskName() *string
	SetTotalDuration(v int64) *ClinkListCdrObRequest
	GetTotalDuration() *int64
	SetTotalDurationEnd(v int64) *ClinkListCdrObRequest
	GetTotalDurationEnd() *int64
	SetXnumber(v string) *ClinkListCdrObRequest
	GetXnumber() *string
}

type ClinkListCdrObRequest struct {
	// 通话时长范围查询最小值，和bridgeDurationEnd配合使用
	//
	// example:
	//
	// 10
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 通话时长范围查询最大值
	//
	// example:
	//
	// 11
	BridgeDurationEnd *int64 `json:"BridgeDurationEnd,omitempty" xml:"BridgeDurationEnd,omitempty"`
	// 接听时间，时间戳格式精确到秒。
	//
	// example:
	//
	// 1775024775
	BridgeTime *int64 `json:"BridgeTime,omitempty" xml:"BridgeTime,omitempty"`
	// 接听时间查询截止，时间戳格式精确到秒。
	//
	// example:
	//
	// 1775024775
	BridgeTimeEnd *int64 `json:"BridgeTimeEnd,omitempty" xml:"BridgeTimeEnd,omitempty"`
	// 通话类型  4: 预览外呼，6: 主叫外呼， 默认全部
	//
	// example:
	//
	// 4
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// 客户号码
	//
	// example:
	//
	// xxx
	CustomerNumber *string `json:"CustomerNumber,omitempty" xml:"CustomerNumber,omitempty"`
	// 结束时间，时间戳格式精确到秒。
	//
	// example:
	//
	// 1775024775
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 结束时间查询截止，时间戳格式精确到秒。
	//
	// example:
	//
	// 1775024775
	EndTimeEnd *int64 `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否隐藏号码。  0: 不隐藏，1: 中间四位，2: 最后八位，3: 全部号码，4: 最后四位。默认值为 0
	//
	// example:
	//
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 外显号码
	//
	// example:
	//
	// xxx
	LeftClid *string `json:"LeftClid,omitempty" xml:"LeftClid,omitempty"`
	// 查询条数，范围 10-100。默认值为 10。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 19
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// MainUniqueId
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 偏移量，范围 0-99990。默认值为 0。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 0
	Offset  *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 外呼组
	//
	// example:
	//
	// 2222
	Qnos                 *string `json:"Qnos,omitempty" xml:"Qnos,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 滚动查询ID。配合scrollSearch使用，第一次查询时不需要传递此参数，后续查询可从返回列表中取值，并传参数。当scrollId为空时，代表滚动查询结束。
	//
	// example:
	//
	// ScrollId
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
	// 是否滚动查询。默认false，true代表使用滚动查询。
	//
	// example:
	//
	// false
	ScrollSearch *string `json:"ScrollSearch,omitempty" xml:"ScrollSearch,omitempty"`
	// 开始时间，时间戳格式精确到秒。默认值取当前月份第一天
	//
	// example:
	//
	// 1775024775
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 开始时间查询截止，时间戳格式精确到秒，开始时间和结束时间跨度不能超过一个月。默认值取当前时间
	//
	// example:
	//
	// 1775024775
	StartTimeEnd *int64 `json:"StartTimeEnd,omitempty" xml:"StartTimeEnd,omitempty"`
	// 接听状态。取值范围如下：  1: 客户未接听  2: 座席未接听  3: 客户接听  4: 座席接听  默认全部
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 呼叫任务名称
	//
	// example:
	//
	// TaskName
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// 总时长范围查询最小值，和totalDurationEnd配合使用
	//
	// example:
	//
	// 80
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 总时长范围查询最大值
	//
	// example:
	//
	// 86
	TotalDurationEnd *int64 `json:"TotalDurationEnd,omitempty" xml:"TotalDurationEnd,omitempty"`
	// 虚拟号码
	//
	// example:
	//
	// xxx
	Xnumber *string `json:"Xnumber,omitempty" xml:"Xnumber,omitempty"`
}

func (s ClinkListCdrObRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrObRequest) GoString() string {
	return s.String()
}

func (s *ClinkListCdrObRequest) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrObRequest) GetBridgeDurationEnd() *int64 {
	return s.BridgeDurationEnd
}

func (s *ClinkListCdrObRequest) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrObRequest) GetBridgeTimeEnd() *int64 {
	return s.BridgeTimeEnd
}

func (s *ClinkListCdrObRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *ClinkListCdrObRequest) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrObRequest) GetCnos() *string {
	return s.Cnos
}

func (s *ClinkListCdrObRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrObRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrObRequest) GetEndTimeEnd() *int64 {
	return s.EndTimeEnd
}

func (s *ClinkListCdrObRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrObRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkListCdrObRequest) GetLeftClid() *string {
	return s.LeftClid
}

func (s *ClinkListCdrObRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListCdrObRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrObRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListCdrObRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListCdrObRequest) GetQnos() *string {
	return s.Qnos
}

func (s *ClinkListCdrObRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListCdrObRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListCdrObRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrObRequest) GetScrollSearch() *string {
	return s.ScrollSearch
}

func (s *ClinkListCdrObRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrObRequest) GetStartTimeEnd() *int64 {
	return s.StartTimeEnd
}

func (s *ClinkListCdrObRequest) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkListCdrObRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ClinkListCdrObRequest) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrObRequest) GetTotalDurationEnd() *int64 {
	return s.TotalDurationEnd
}

func (s *ClinkListCdrObRequest) GetXnumber() *string {
	return s.Xnumber
}

func (s *ClinkListCdrObRequest) SetBridgeDuration(v int64) *ClinkListCdrObRequest {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrObRequest) SetBridgeDurationEnd(v int64) *ClinkListCdrObRequest {
	s.BridgeDurationEnd = &v
	return s
}

func (s *ClinkListCdrObRequest) SetBridgeTime(v int64) *ClinkListCdrObRequest {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrObRequest) SetBridgeTimeEnd(v int64) *ClinkListCdrObRequest {
	s.BridgeTimeEnd = &v
	return s
}

func (s *ClinkListCdrObRequest) SetCallType(v int64) *ClinkListCdrObRequest {
	s.CallType = &v
	return s
}

func (s *ClinkListCdrObRequest) SetClientNumber(v string) *ClinkListCdrObRequest {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrObRequest) SetCnos(v string) *ClinkListCdrObRequest {
	s.Cnos = &v
	return s
}

func (s *ClinkListCdrObRequest) SetCustomerNumber(v string) *ClinkListCdrObRequest {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrObRequest) SetEndTime(v int64) *ClinkListCdrObRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrObRequest) SetEndTimeEnd(v int64) *ClinkListCdrObRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *ClinkListCdrObRequest) SetEnterpriseId(v int64) *ClinkListCdrObRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrObRequest) SetHiddenType(v int64) *ClinkListCdrObRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkListCdrObRequest) SetLeftClid(v string) *ClinkListCdrObRequest {
	s.LeftClid = &v
	return s
}

func (s *ClinkListCdrObRequest) SetLimit(v int64) *ClinkListCdrObRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListCdrObRequest) SetMainUniqueId(v string) *ClinkListCdrObRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrObRequest) SetOffset(v int64) *ClinkListCdrObRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListCdrObRequest) SetOwnerId(v int64) *ClinkListCdrObRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListCdrObRequest) SetQnos(v string) *ClinkListCdrObRequest {
	s.Qnos = &v
	return s
}

func (s *ClinkListCdrObRequest) SetResourceOwnerAccount(v string) *ClinkListCdrObRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListCdrObRequest) SetResourceOwnerId(v int64) *ClinkListCdrObRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListCdrObRequest) SetScrollId(v string) *ClinkListCdrObRequest {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrObRequest) SetScrollSearch(v string) *ClinkListCdrObRequest {
	s.ScrollSearch = &v
	return s
}

func (s *ClinkListCdrObRequest) SetStartTime(v int64) *ClinkListCdrObRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrObRequest) SetStartTimeEnd(v int64) *ClinkListCdrObRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *ClinkListCdrObRequest) SetStatus(v int64) *ClinkListCdrObRequest {
	s.Status = &v
	return s
}

func (s *ClinkListCdrObRequest) SetTaskName(v string) *ClinkListCdrObRequest {
	s.TaskName = &v
	return s
}

func (s *ClinkListCdrObRequest) SetTotalDuration(v int64) *ClinkListCdrObRequest {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrObRequest) SetTotalDurationEnd(v int64) *ClinkListCdrObRequest {
	s.TotalDurationEnd = &v
	return s
}

func (s *ClinkListCdrObRequest) SetXnumber(v string) *ClinkListCdrObRequest {
	s.Xnumber = &v
	return s
}

func (s *ClinkListCdrObRequest) Validate() error {
	return dara.Validate(s)
}
