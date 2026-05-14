// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListCdrIbAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentAnswerInTime(v int64) *ClinkListCdrIbAgentRequest
	GetAgentAnswerInTime() *int64
	SetBridgeDuration(v int64) *ClinkListCdrIbAgentRequest
	GetBridgeDuration() *int64
	SetBridgeDurationEnd(v int64) *ClinkListCdrIbAgentRequest
	GetBridgeDurationEnd() *int64
	SetBridgeTime(v int64) *ClinkListCdrIbAgentRequest
	GetBridgeTime() *int64
	SetBridgeTimeEnd(v int64) *ClinkListCdrIbAgentRequest
	GetBridgeTimeEnd() *int64
	SetClientNumber(v string) *ClinkListCdrIbAgentRequest
	GetClientNumber() *string
	SetCno(v string) *ClinkListCdrIbAgentRequest
	GetCno() *string
	SetCustomerNumber(v string) *ClinkListCdrIbAgentRequest
	GetCustomerNumber() *string
	SetEndTime(v int64) *ClinkListCdrIbAgentRequest
	GetEndTime() *int64
	SetEndTimeEnd(v int64) *ClinkListCdrIbAgentRequest
	GetEndTimeEnd() *int64
	SetEnterpriseId(v int64) *ClinkListCdrIbAgentRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkListCdrIbAgentRequest
	GetHiddenType() *int64
	SetHotlineName(v string) *ClinkListCdrIbAgentRequest
	GetHotlineName() *string
	SetLimit(v int64) *ClinkListCdrIbAgentRequest
	GetLimit() *int64
	SetMainUniqueId(v string) *ClinkListCdrIbAgentRequest
	GetMainUniqueId() *string
	SetOffset(v int64) *ClinkListCdrIbAgentRequest
	GetOffset() *int64
	SetOwnerId(v int64) *ClinkListCdrIbAgentRequest
	GetOwnerId() *int64
	SetQno(v string) *ClinkListCdrIbAgentRequest
	GetQno() *string
	SetResourceOwnerAccount(v string) *ClinkListCdrIbAgentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkListCdrIbAgentRequest
	GetResourceOwnerId() *int64
	SetScrollId(v string) *ClinkListCdrIbAgentRequest
	GetScrollId() *string
	SetStartTime(v int64) *ClinkListCdrIbAgentRequest
	GetStartTime() *int64
	SetStartTimeEnd(v int64) *ClinkListCdrIbAgentRequest
	GetStartTimeEnd() *int64
}

type ClinkListCdrIbAgentRequest struct {
	// 及时应答。取值范围如下： 0: 否 1: 是 默认全部
	//
	// example:
	//
	// 1
	AgentAnswerInTime *int64 `json:"AgentAnswerInTime,omitempty" xml:"AgentAnswerInTime,omitempty"`
	// 通话时长范围查询最小值，和bridgeDurationEnd配合使用
	//
	// example:
	//
	// 25
	BridgeDuration *int64 `json:"BridgeDuration,omitempty" xml:"BridgeDuration,omitempty"`
	// 通话时长范围查询最大值
	//
	// example:
	//
	// 90
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
	// 座席号码
	//
	// example:
	//
	// xxx
	ClientNumber *string `json:"ClientNumber,omitempty" xml:"ClientNumber,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1001
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
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
	// 是否隐藏号码。 0: 不隐藏，1: 中间四位，2: 最后八位，3: 全部号码，4: 最后四位。默认值为 0
	//
	// example:
	//
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 热线别名
	//
	// example:
	//
	// HotlineName
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
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
	// 2211
	Qno                  *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 滚动查询ID。配合scrollSearch使用，第一次查询时不需要传递此参数，后续查询可从返回列表中取值，并传参数。当scrollId为空时，代表滚动查询结束。
	//
	// example:
	//
	// ScrollId
	ScrollId *string `json:"ScrollId,omitempty" xml:"ScrollId,omitempty"`
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
}

func (s ClinkListCdrIbAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkListCdrIbAgentRequest) GoString() string {
	return s.String()
}

func (s *ClinkListCdrIbAgentRequest) GetAgentAnswerInTime() *int64 {
	return s.AgentAnswerInTime
}

func (s *ClinkListCdrIbAgentRequest) GetBridgeDuration() *int64 {
	return s.BridgeDuration
}

func (s *ClinkListCdrIbAgentRequest) GetBridgeDurationEnd() *int64 {
	return s.BridgeDurationEnd
}

func (s *ClinkListCdrIbAgentRequest) GetBridgeTime() *int64 {
	return s.BridgeTime
}

func (s *ClinkListCdrIbAgentRequest) GetBridgeTimeEnd() *int64 {
	return s.BridgeTimeEnd
}

func (s *ClinkListCdrIbAgentRequest) GetClientNumber() *string {
	return s.ClientNumber
}

func (s *ClinkListCdrIbAgentRequest) GetCno() *string {
	return s.Cno
}

func (s *ClinkListCdrIbAgentRequest) GetCustomerNumber() *string {
	return s.CustomerNumber
}

func (s *ClinkListCdrIbAgentRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ClinkListCdrIbAgentRequest) GetEndTimeEnd() *int64 {
	return s.EndTimeEnd
}

func (s *ClinkListCdrIbAgentRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkListCdrIbAgentRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkListCdrIbAgentRequest) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkListCdrIbAgentRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ClinkListCdrIbAgentRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkListCdrIbAgentRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *ClinkListCdrIbAgentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkListCdrIbAgentRequest) GetQno() *string {
	return s.Qno
}

func (s *ClinkListCdrIbAgentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkListCdrIbAgentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkListCdrIbAgentRequest) GetScrollId() *string {
	return s.ScrollId
}

func (s *ClinkListCdrIbAgentRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ClinkListCdrIbAgentRequest) GetStartTimeEnd() *int64 {
	return s.StartTimeEnd
}

func (s *ClinkListCdrIbAgentRequest) SetAgentAnswerInTime(v int64) *ClinkListCdrIbAgentRequest {
	s.AgentAnswerInTime = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetBridgeDuration(v int64) *ClinkListCdrIbAgentRequest {
	s.BridgeDuration = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetBridgeDurationEnd(v int64) *ClinkListCdrIbAgentRequest {
	s.BridgeDurationEnd = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetBridgeTime(v int64) *ClinkListCdrIbAgentRequest {
	s.BridgeTime = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetBridgeTimeEnd(v int64) *ClinkListCdrIbAgentRequest {
	s.BridgeTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetClientNumber(v string) *ClinkListCdrIbAgentRequest {
	s.ClientNumber = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetCno(v string) *ClinkListCdrIbAgentRequest {
	s.Cno = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetCustomerNumber(v string) *ClinkListCdrIbAgentRequest {
	s.CustomerNumber = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetEndTime(v int64) *ClinkListCdrIbAgentRequest {
	s.EndTime = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetEndTimeEnd(v int64) *ClinkListCdrIbAgentRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetEnterpriseId(v int64) *ClinkListCdrIbAgentRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetHiddenType(v int64) *ClinkListCdrIbAgentRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetHotlineName(v string) *ClinkListCdrIbAgentRequest {
	s.HotlineName = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetLimit(v int64) *ClinkListCdrIbAgentRequest {
	s.Limit = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetMainUniqueId(v string) *ClinkListCdrIbAgentRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetOffset(v int64) *ClinkListCdrIbAgentRequest {
	s.Offset = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetOwnerId(v int64) *ClinkListCdrIbAgentRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetQno(v string) *ClinkListCdrIbAgentRequest {
	s.Qno = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetResourceOwnerAccount(v string) *ClinkListCdrIbAgentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetResourceOwnerId(v int64) *ClinkListCdrIbAgentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetScrollId(v string) *ClinkListCdrIbAgentRequest {
	s.ScrollId = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetStartTime(v int64) *ClinkListCdrIbAgentRequest {
	s.StartTime = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) SetStartTimeEnd(v int64) *ClinkListCdrIbAgentRequest {
	s.StartTimeEnd = &v
	return s
}

func (s *ClinkListCdrIbAgentRequest) Validate() error {
	return dara.Validate(s)
}
