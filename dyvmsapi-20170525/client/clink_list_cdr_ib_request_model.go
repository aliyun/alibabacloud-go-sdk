// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeDuration(v int64) *ClinkListCdrIbRequest
	GetBridgeDuration() *int64
	SetBridgeDurationEnd(v int64) *ClinkListCdrIbRequest
	GetBridgeDurationEnd() *int64
	SetBridgeTime(v int64) *ClinkListCdrIbRequest
	GetBridgeTime() *int64
	SetBridgeTimeEnd(v int64) *ClinkListCdrIbRequest
	GetBridgeTimeEnd() *int64
	SetClientNumber(v string) *ClinkListCdrIbRequest
	GetClientNumber() *string
	SetCnos(v string) *ClinkListCdrIbRequest
	GetCnos() *string
	SetCustomerNumber(v string) *ClinkListCdrIbRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *ClinkListCdrIbRequest
	GetEndTime() *int64
	SetEndTimeEnd(v int64) *ClinkListCdrIbRequest
	GetEndTimeEnd() *int64
	SetEnterpriseId(v int64) *ClinkListCdrIbRequest
	GetEnterpriseId() *int64
	SetFirstCallCno(v string) *ClinkListCdrIbRequest
	GetFirstCallCno() *string
	SetFirstCallNumber(v string) *ClinkListCdrIbRequest
	GetFirstCallNumber() *string
	SetFirstCallQno(v string) *ClinkListCdrIbRequest
	GetFirstCallQno() *string
	SetHiddenType(v int64) *ClinkListCdrIbRequest
	GetHiddenType() *int64
	SetLimit(v int64) *ClinkListCdrIbRequest
	GetLimit() *int64
	SetMainUniqueId(v string) *ClinkListCdrIbRequest
	GetMainUniqueId() *string
	SetOffset(v int64) *ClinkListCdrIbRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListCdrIbRequest
	GetOwnerId() *int64
	SetQnos(v string) *ClinkListCdrIbRequest
	GetQnos() *string
	SetResourceOwnerAccount(v string) *ClinkListCdrIbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListCdrIbRequest
	GetResourceOwnerId() *int64
	SetScrollId(v string) *ClinkListCdrIbRequest
	GetScrollId() *string
	SetScrollSearch(v string) *ClinkListCdrIbRequest
	GetScrollSearch() *string
	SetStartTime(v int64) *ClinkListCdrIbRequest
	GetStartTime() *int64
	SetStartTimeEnd(v int64) *ClinkListCdrIbRequest
	GetStartTimeEnd() *int64
	SetStatuses(v int64) *ClinkListCdrIbRequest
	GetStatuses() *int64
	SetTotalDuration(v int64) *ClinkListCdrIbRequest
	GetTotalDuration() *int64
	SetTotalDurationEnd(v int64) *ClinkListCdrIbRequest
	GetTotalDurationEnd() *int64
}

type ClinkListCdrIbRequest struct {
	// 通话时长范围查询最小值，和bridgeDurationEnd配合使用
	//
	// example:
	//
	// 40
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 通话时长范围查询最大值
	//
	// example:
	//
	// 60
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
	// 座席电话
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席号
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
	// 首呼座席
	//
	// example:
	//
	// 1122
	FirstCallCno *string `json:"FirstCallCno,omitempty" xml:"FirstCallCno,omitempty"`
	// 首呼座席电话
	//
	// example:
	//
	// xxx
	FirstCallNumber *string `json:"FirstCallNumber,omitempty" xml:"FirstCallNumber,omitempty"`
	// 首呼队列
	//
	// example:
	//
	// 2222
	FirstCallQno *string `json:"FirstCallQno,omitempty" xml:"FirstCallQno,omitempty"`
	// 是否隐藏号码。  0: 不隐藏，1: 中间四位，2: 最后八位，3: 全部号码，4: 最后四位。默认值为 0
	//
	// example:
	//
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 查询条数，范围 10-100。默认值为 10。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	// 偏移量，范围 0-99990。默认值为 0。注：limit + offset 不允许超过100000
	//
	// example:
	//
	// 0
	Offset  *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列号
	//
	// example:
	//
	// 1122
	Qnos                 *string `json:"Qnos,omitempty" xml:"Qnos,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 滚动查询ID。配合scrollSearch使用，第一次查询时不需要传递此参数，后续查询可从返回列表中取值，并传参数。当scrollId为空时，代表滚动查询结束。
	//
	// example:
	//
	// xxx
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
	// 接听状态。取值范围如下：  1: 人工接听  2: 人工未接听  3: 系统应答  4: 系统未应答  默认全部
	//
	// example:
	//
	// 1
	Statuses *int64 `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	// 总时长范围查询最小值，和totalDurationEnd配合使用
	//
	// example:
	//
	// 69
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// 总时长范围查询最大值
	//
	// example:
	//
	// 47
	TotalDurationEnd *int64 `json:"TotalDurationEnd,omitempty" xml:"TotalDurationEnd,omitempty"`
}

func (s ClinkListCdrIbRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbRequest) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbRequest) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrIbRequest) GetBridgeDurationEnd() *int64 {
	return s.BridgeDurationEnd
}

func (s *ClinkListCdrIbRequest) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrIbRequest) GetBridgeTimeEnd() *int64 {
	return s.BridgeTimeEnd
}

func (s *ClinkListCdrIbRequest) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrIbRequest) GetCnos() *string {
	return s.Cnos
}

func (s *ClinkListCdrIbRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbRequest) GetEndTimeEnd() *int64 {
	return s.EndTimeEnd
}

func (s *ClinkListCdrIbRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrIbRequest) GetFirstCallCno() *string {
	return s.FirstCallCno
}

func (s *ClinkListCdrIbRequest) GetFirstCallNumber() *string {
	return s.FirstCallNumber
}

func (s *ClinkListCdrIbRequest) GetFirstCallQno() *string {
	return s.FirstCallQno
}

func (s *ClinkListCdrIbRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkListCdrIbRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListCdrIbRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrIbRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListCdrIbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListCdrIbRequest) GetQnos() *string {
	return s.Qnos
}

func (s *ClinkListCdrIbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListCdrIbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListCdrIbRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrIbRequest) GetScrollSearch() *string {
	return s.ScrollSearch
}

func (s *ClinkListCdrIbRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbRequest) GetStartTimeEnd() *int64 {
	return s.StartTimeEnd
}

func (s *ClinkListCdrIbRequest) GetStatuses() *int64 {
	return s.Statuses
}

func (s *ClinkListCdrIbRequest) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *ClinkListCdrIbRequest) GetTotalDurationEnd() *int64 {
	return s.TotalDurationEnd
}

func (s *ClinkListCdrIbRequest) SetBridgeDuration(v int64) *ClinkListCdrIbRequest {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetBridgeDurationEnd(v int64) *ClinkListCdrIbRequest {
	s.BridgeDurationEnd = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetBridgeTime(v int64) *ClinkListCdrIbRequest {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetBridgeTimeEnd(v int64) *ClinkListCdrIbRequest {
	s.BridgeTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetClientNumber(v string) *ClinkListCdrIbRequest {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetCnos(v string) *ClinkListCdrIbRequest {
	s.Cnos = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetCustomerNumber(v string) *ClinkListCdrIbRequest {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetEndTime(v int64) *ClinkListCdrIbRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetEndTimeEnd(v int64) *ClinkListCdrIbRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetEnterpriseId(v int64) *ClinkListCdrIbRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetFirstCallCno(v string) *ClinkListCdrIbRequest {
	s.FirstCallCno = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetFirstCallNumber(v string) *ClinkListCdrIbRequest {
	s.FirstCallNumber = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetFirstCallQno(v string) *ClinkListCdrIbRequest {
	s.FirstCallQno = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetHiddenType(v int64) *ClinkListCdrIbRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetLimit(v int64) *ClinkListCdrIbRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetMainUniqueId(v string) *ClinkListCdrIbRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetOffset(v int64) *ClinkListCdrIbRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetOwnerId(v int64) *ClinkListCdrIbRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetQnos(v string) *ClinkListCdrIbRequest {
	s.Qnos = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetResourceOwnerAccount(v string) *ClinkListCdrIbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetResourceOwnerId(v int64) *ClinkListCdrIbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetScrollId(v string) *ClinkListCdrIbRequest {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetScrollSearch(v string) *ClinkListCdrIbRequest {
	s.ScrollSearch = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetStartTime(v int64) *ClinkListCdrIbRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetStartTimeEnd(v int64) *ClinkListCdrIbRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetStatuses(v int64) *ClinkListCdrIbRequest {
	s.Statuses = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetTotalDuration(v int64) *ClinkListCdrIbRequest {
	s.TotalDuration = &v
	return s
}

func (s *ClinkListCdrIbRequest) SetTotalDurationEnd(v int64) *ClinkListCdrIbRequest {
	s.TotalDurationEnd = &v
	return s
}

func (s *ClinkListCdrIbRequest) Validate() error {
	return dara.Validate(s)
}
