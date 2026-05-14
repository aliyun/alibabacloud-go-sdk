// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkDescribeQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkDescribeQueueResponseBody
	GetCode() *string
	SetData(v *ClinkDescribeQueueResponseBodyData) *ClinkDescribeQueueResponseBody
	GetData() *ClinkDescribeQueueResponseBodyData
	SetMessage(v string) *ClinkDescribeQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkDescribeQueueResponseBody
	GetRequestId() *string
}

type ClinkDescribeQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkDescribeQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkDescribeQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkDescribeQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkDescribeQueueResponseBody) GetData() *ClinkDescribeQueueResponseBodyData {
	return s.Data
}

func (s *ClinkDescribeQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkDescribeQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkDescribeQueueResponseBody) SetAccessDeniedDetail(v string) *ClinkDescribeQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkDescribeQueueResponseBody) SetCode(v string) *ClinkDescribeQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkDescribeQueueResponseBody) SetData(v *ClinkDescribeQueueResponseBodyData) *ClinkDescribeQueueResponseBody {
	s.Data = v
	return s
}

func (s *ClinkDescribeQueueResponseBody) SetMessage(v string) *ClinkDescribeQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkDescribeQueueResponseBody) SetRequestId(v string) *ClinkDescribeQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkDescribeQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeQueueResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string                                   `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	Result         *ClinkDescribeQueueResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ClinkDescribeQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkDescribeQueueResponseBodyData) GetResult() *ClinkDescribeQueueResponseBodyDataResult {
	return s.Result
}

func (s *ClinkDescribeQueueResponseBodyData) SetClinkRequestId(v string) *ClinkDescribeQueueResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyData) SetResult(v *ClinkDescribeQueueResponseBodyDataResult) *ClinkDescribeQueueResponseBodyData {
	s.Result = v
	return s
}

func (s *ClinkDescribeQueueResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkDescribeQueueResponseBodyDataResult struct {
	// 在线客服排队位置推送(小于该位置则推送)
	//
	// example:
	//
	// 20
	ChatLocation *int64 `json:"ChatLocation,omitempty" xml:"ChatLocation,omitempty"`
	// 在线客服最大排队数
	//
	// example:
	//
	// 20
	ChatMaxWait *int64 `json:"ChatMaxWait,omitempty" xml:"ChatMaxWait,omitempty"`
	// 在线客服分配策略，2：轮选；4：随机；5：技能优先
	//
	// example:
	//
	// 2
	ChatStrategy *int64 `json:"ChatStrategy,omitempty" xml:"ChatStrategy,omitempty"`
	// 企业id
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 允许呼入队列 0：不允许 1：允许
	//
	// example:
	//
	// 1
	IbAllowed *int64 `json:"IbAllowed,omitempty" xml:"IbAllowed,omitempty"`
	// 队列中为空时是否可以join（设置了队列等待数，来电进入队列后处于等待状态，等待超时才溢出；如果勾选了这几个选项，认为座席不可用，直接溢出，减少客户等待时间）， 1: 置忙；2: 通话中；4: 振铃；8: 无效；16: 整理。 如选多种状态则传对应数值的和，如选置忙和通话中，值为 3
	//
	// example:
	//
	// 3
	JoinEmpty *int64 `json:"JoinEmpty,omitempty" xml:"JoinEmpty,omitempty"`
	// 最大并发置忙座席数，开关配置
	//
	// example:
	//
	// 1
	MaxPauseClientFlag *int64 `json:"MaxPauseClientFlag,omitempty" xml:"MaxPauseClientFlag,omitempty"`
	// 最大并发置忙座席数，配置类型
	//
	// example:
	//
	// 1
	MaxPauseClientType *int64 `json:"MaxPauseClientType,omitempty" xml:"MaxPauseClientType,omitempty"`
	// 最大并发置忙座席数，配置值
	//
	// example:
	//
	// 10
	MaxPauseClientValue *int64 `json:"MaxPauseClientValue,omitempty" xml:"MaxPauseClientValue,omitempty"`
	// 呼叫中心最大排队数
	//
	// example:
	//
	// 20
	MaxWait *int64 `json:"MaxWait,omitempty" xml:"MaxWait,omitempty"`
	// 座席超时时长
	//
	// example:
	//
	// 51
	MemberTimeout *int64 `json:"MemberTimeout,omitempty" xml:"MemberTimeout,omitempty"`
	// 队列名称
	//
	// example:
	//
	// 41
	Name *int64 `json:"Name,omitempty" xml:"Name,omitempty"`
	// 队列号
	//
	// example:
	//
	// 2201
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// cno 座席号 penalty 优先级 type 座席类型 1：全渠道、2：呼叫中心、3：在线客服
	QueueMembers []*ClinkDescribeQueueResponseBodyDataResultQueueMembers `json:"QueueMembers,omitempty" xml:"QueueMembers,omitempty" type:"Repeated"`
	// 队列超时时长
	//
	// example:
	//
	// 51
	QueueTimeout *int64 `json:"QueueTimeout,omitempty" xml:"QueueTimeout,omitempty"`
	// 座席超时无应答,呼叫下一座席的延迟秒数,⇐0时采用默认值5。
	//
	// example:
	//
	// 5
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 语音报号(是否报座席号)(1：是 0：否)
	//
	// example:
	//
	// 0
	SayCno *int64 `json:"SayCno,omitempty" xml:"SayCno,omitempty"`
	// 服务水平秒数，取值范围 10-30秒，低于此时间内接听的认为是高服务水平
	//
	// example:
	//
	// 10
	ServiceLevel *int64 `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
	// 呼叫策略，1：顺序，2：轮选，3：平均，4：随机，5：技能优先，6：距上一通呼入空闲时间累计最长，7：技能优先+距上一通呼入空闲时间累计最长，8：距上一通呼入空闲时间累计最长(全队列)，9：技能优先+距上一通呼入空闲时间累计最长(全队列)，10：全队列平均，11：技能优先+轮选，12：最长空闲时长，13：技能优先+当前空闲状态时长最长，14：距上一通通话空闲时间累计最长
	//
	// example:
	//
	// 2
	Strategy *int64 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// 支持VIP级别(VIP客户优先)(0：不支持 1：支持)
	//
	// example:
	//
	// 0
	VipSupport *int64 `json:"VipSupport,omitempty" xml:"VipSupport,omitempty"`
	// 队列优先级，取值范围 1-10。默认值为 1，座席属于多个队列时，优先级高的队列的来电将优先接听；数值越高，优先级越高
	//
	// example:
	//
	// 6
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时间
	//
	// example:
	//
	// 45
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkDescribeQueueResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetChatLocation() *int64 {
	return s.ChatLocation
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetChatMaxWait() *int64 {
	return s.ChatMaxWait
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetChatStrategy() *int64 {
	return s.ChatStrategy
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetIbAllowed() *int64 {
	return s.IbAllowed
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetMaxPauseClientFlag() *int64 {
	return s.MaxPauseClientFlag
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetMaxPauseClientType() *int64 {
	return s.MaxPauseClientType
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetMaxPauseClientValue() *int64 {
	return s.MaxPauseClientValue
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetMaxWait() *int64 {
	return s.MaxWait
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetName() *int64 {
	return s.Name
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetQno() *string {
	return s.Qno
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetQueueMembers() []*ClinkDescribeQueueResponseBodyDataResultQueueMembers {
	return s.QueueMembers
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetRetry() *int64 {
	return s.Retry
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetSayCno() *int64 {
	return s.SayCno
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetStrategy() *int64 {
	return s.Strategy
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetWeight() *int64 {
	return s.Weight
}

func (s *ClinkDescribeQueueResponseBodyDataResult) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetChatLocation(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.ChatLocation = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetChatMaxWait(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.ChatMaxWait = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetChatStrategy(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.ChatStrategy = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetEnterpriseId(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetIbAllowed(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.IbAllowed = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetJoinEmpty(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.JoinEmpty = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetMaxPauseClientFlag(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.MaxPauseClientFlag = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetMaxPauseClientType(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.MaxPauseClientType = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetMaxPauseClientValue(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.MaxPauseClientValue = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetMaxWait(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.MaxWait = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetMemberTimeout(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.MemberTimeout = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetName(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetQno(v string) *ClinkDescribeQueueResponseBodyDataResult {
	s.Qno = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetQueueMembers(v []*ClinkDescribeQueueResponseBodyDataResultQueueMembers) *ClinkDescribeQueueResponseBodyDataResult {
	s.QueueMembers = v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetQueueTimeout(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.QueueTimeout = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetRetry(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.Retry = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetSayCno(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.SayCno = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetServiceLevel(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.ServiceLevel = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetStrategy(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.Strategy = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetVipSupport(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.VipSupport = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetWeight(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.Weight = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) SetWrapupTime(v int64) *ClinkDescribeQueueResponseBodyDataResult {
	s.WrapupTime = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResult) Validate() error {
	if s.QueueMembers != nil {
		for _, item := range s.QueueMembers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkDescribeQueueResponseBodyDataResultQueueMembers struct {
	// 座席号
	//
	// example:
	//
	// 2233
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 优先级
	//
	// example:
	//
	// 3
	Penalty *int64 `json:"Penalty,omitempty" xml:"Penalty,omitempty"`
	// 座席类型
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkDescribeQueueResponseBodyDataResultQueueMembers) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeQueueResponseBodyDataResultQueueMembers) GoString() string {
	return s.String()
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) GetCno() *string {
	return s.Cno
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) GetPenalty() *int64 {
	return s.Penalty
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) GetType() *int64 {
	return s.Type
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) SetCno(v string) *ClinkDescribeQueueResponseBodyDataResultQueueMembers {
	s.Cno = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) SetPenalty(v int64) *ClinkDescribeQueueResponseBodyDataResultQueueMembers {
	s.Penalty = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) SetType(v int64) *ClinkDescribeQueueResponseBodyDataResultQueueMembers {
	s.Type = &v
	return s
}

func (s *ClinkDescribeQueueResponseBodyDataResultQueueMembers) Validate() error {
	return dara.Validate(s)
}
