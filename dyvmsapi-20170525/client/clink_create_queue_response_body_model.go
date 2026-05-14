// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkCreateQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkCreateQueueResponseBody
	GetCode() *string
	SetData(v *ClinkCreateQueueResponseBodyData) *ClinkCreateQueueResponseBody
	GetData() *ClinkCreateQueueResponseBodyData
	SetMessage(v string) *ClinkCreateQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkCreateQueueResponseBody
	GetRequestId() *string
}

type ClinkCreateQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkCreateQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkCreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkCreateQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkCreateQueueResponseBody) GetData() *ClinkCreateQueueResponseBodyData {
	return s.Data
}

func (s *ClinkCreateQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkCreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkCreateQueueResponseBody) SetAccessDeniedDetail(v string) *ClinkCreateQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkCreateQueueResponseBody) SetCode(v string) *ClinkCreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkCreateQueueResponseBody) SetData(v *ClinkCreateQueueResponseBodyData) *ClinkCreateQueueResponseBody {
	s.Data = v
	return s
}

func (s *ClinkCreateQueueResponseBody) SetMessage(v string) *ClinkCreateQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkCreateQueueResponseBody) SetRequestId(v string) *ClinkCreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkCreateQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateQueueResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// 示例值示例值示例值
	ClinkRequestId *string                                `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	Queue          *ClinkCreateQueueResponseBodyDataQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
}

func (s ClinkCreateQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkCreateQueueResponseBodyData) GetQueue() *ClinkCreateQueueResponseBodyDataQueue {
	return s.Queue
}

func (s *ClinkCreateQueueResponseBodyData) SetClinkRequestId(v string) *ClinkCreateQueueResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyData) SetQueue(v *ClinkCreateQueueResponseBodyDataQueue) *ClinkCreateQueueResponseBodyData {
	s.Queue = v
	return s
}

func (s *ClinkCreateQueueResponseBodyData) Validate() error {
	if s.Queue != nil {
		if err := s.Queue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateQueueResponseBodyDataQueue struct {
	// 在线客服排队位置推送(小于该位置则推送)
	//
	// example:
	//
	// 10
	ChatLocation *int64 `json:"ChatLocation,omitempty" xml:"ChatLocation,omitempty"`
	// 在线客服最大排队数
	//
	// example:
	//
	// 10
	ChatMaxWait *int64 `json:"ChatMaxWait,omitempty" xml:"ChatMaxWait,omitempty"`
	// 在线客服分配策略，2：轮选；4：随机；5：技能优先
	//
	// example:
	//
	// 2
	ChatStrategy *int64 `json:"ChatStrategy,omitempty" xml:"ChatStrategy,omitempty"`
	// 是否允许呼入队列，0: 否；1: 是
	//
	// example:
	//
	// 1
	IbAllowed *int64 `json:"IbAllowed,omitempty" xml:"IbAllowed,omitempty"`
	// 队列中为空时是否可以join，1: 置忙，2: 通话中，4: 振铃，8: 无效，16: 整理 如选多种状态则传对应数值的和，如选置忙和通话中，值为 3。默认值为 0 即都未选中
	//
	// example:
	//
	// 0
	JoinEmpty *int64 `json:"JoinEmpty,omitempty" xml:"JoinEmpty,omitempty"`
	// 最大并发置忙座席数，开关配置
	//
	// example:
	//
	// 19
	MaxPauseClientFlag *int64 `json:"MaxPauseClientFlag,omitempty" xml:"MaxPauseClientFlag,omitempty"`
	// 最大并发置忙座席数，配置类型
	//
	// example:
	//
	// 73
	MaxPauseClientType *int64 `json:"MaxPauseClientType,omitempty" xml:"MaxPauseClientType,omitempty"`
	// 最大并发置忙座席数，配置值
	//
	// example:
	//
	// 38
	MaxPauseClientValue *int64 `json:"MaxPauseClientValue,omitempty" xml:"MaxPauseClientValue,omitempty"`
	// 最大等待数，设置范围 0-999，0 表示不做任何限制
	//
	// example:
	//
	// 0
	MaxWait *int64 `json:"MaxWait,omitempty" xml:"MaxWait,omitempty"`
	// 座席超时时间，取值范围 20-60 秒
	//
	// example:
	//
	// 20
	MemberTimeout *int64 `json:"MemberTimeout,omitempty" xml:"MemberTimeout,omitempty"`
	// 座席名称，只允许输入中文，字母，数字，下划线，长度不超过 20 个字符
	//
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 队列工号，4位数字，要求唯一
	//
	// example:
	//
	// 1122
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列内座席及技能值设置。技能值 1-10，技能值 1 的优先级为最高 cno 队列下座席号 penalty该座席在队列中的优先级
	QueueMembers []*ClinkCreateQueueResponseBodyDataQueueQueueMembers `json:"QueueMembers,omitempty" xml:"QueueMembers,omitempty" type:"Repeated"`
	// 队列超时时间，取值范围 30-600 秒
	//
	// example:
	//
	// 33
	QueueTimeout *int64 `json:"QueueTimeout,omitempty" xml:"QueueTimeout,omitempty"`
	// 语音报号(是否报座席号)(1:是 0: 否)
	//
	// example:
	//
	// 1
	SayCno *int64 `json:"SayCno,omitempty" xml:"SayCno,omitempty"`
	// 服务水平秒数，取值范围 10-30 秒，低于此时间内接听的认为是高服务水平
	//
	// example:
	//
	// 27
	ServiceLevel *int64 `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
	// 呼叫策略，1：顺序，2：轮选，3：平均，4：随机，5：技能优先，6：距上一通呼入空闲时间累计最长，7：技能优先+距上一通呼入空闲时间累计最长，8：距上一通呼入空闲时间累计最长(全队列)，9：技能优先+距上一通呼入空闲时间累计最长(全队列)，10：全队列平均，11：技能优先+轮选，12：最长空闲时长，13：技能优先+当前空闲状态时长最长，14：距上一通通话空闲时间累计最长
	//
	// example:
	//
	// 1
	Strategy *int64 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// 支持VIP级别(VIP客户优先)(0:不支持 1:支持)
	//
	// example:
	//
	// 0
	VipSupport *int64 `json:"VipSupport,omitempty" xml:"VipSupport,omitempty"`
	// 队列优先级，取值范围 1-10 座席属于多个队列时，优先级高的队列的来电将优先接听；数值越高，优先级越高
	//
	// example:
	//
	// 3
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时长，取值范围 3-300 秒，整理期间座席不接受新的呼叫
	//
	// example:
	//
	// 86
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkCreateQueueResponseBodyDataQueue) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueResponseBodyDataQueue) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetChatLocation() *int64 {
	return s.ChatLocation
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetChatMaxWait() *int64 {
	return s.ChatMaxWait
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetChatStrategy() *int64 {
	return s.ChatStrategy
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetIbAllowed() *int64 {
	return s.IbAllowed
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetMaxPauseClientFlag() *int64 {
	return s.MaxPauseClientFlag
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetMaxPauseClientType() *int64 {
	return s.MaxPauseClientType
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetMaxPauseClientValue() *int64 {
	return s.MaxPauseClientValue
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetMaxWait() *int64 {
	return s.MaxWait
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetName() *string {
	return s.Name
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetQno() *string {
	return s.Qno
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetQueueMembers() []*ClinkCreateQueueResponseBodyDataQueueQueueMembers {
	return s.QueueMembers
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetSayCno() *int64 {
	return s.SayCno
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetStrategy() *int64 {
	return s.Strategy
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetWeight() *int64 {
	return s.Weight
}

func (s *ClinkCreateQueueResponseBodyDataQueue) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetChatLocation(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.ChatLocation = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetChatMaxWait(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.ChatMaxWait = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetChatStrategy(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.ChatStrategy = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetIbAllowed(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.IbAllowed = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetJoinEmpty(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.JoinEmpty = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetMaxPauseClientFlag(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.MaxPauseClientFlag = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetMaxPauseClientType(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.MaxPauseClientType = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetMaxPauseClientValue(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.MaxPauseClientValue = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetMaxWait(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.MaxWait = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetMemberTimeout(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.MemberTimeout = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetName(v string) *ClinkCreateQueueResponseBodyDataQueue {
	s.Name = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetQno(v string) *ClinkCreateQueueResponseBodyDataQueue {
	s.Qno = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetQueueMembers(v []*ClinkCreateQueueResponseBodyDataQueueQueueMembers) *ClinkCreateQueueResponseBodyDataQueue {
	s.QueueMembers = v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetQueueTimeout(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.QueueTimeout = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetSayCno(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.SayCno = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetServiceLevel(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.ServiceLevel = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetStrategy(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.Strategy = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetVipSupport(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.VipSupport = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetWeight(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.Weight = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) SetWrapupTime(v int64) *ClinkCreateQueueResponseBodyDataQueue {
	s.WrapupTime = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueue) Validate() error {
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

type ClinkCreateQueueResponseBodyDataQueueQueueMembers struct {
	// 队列下座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 该座席在队列中的优先级
	//
	// example:
	//
	// 7
	Penalty *int64 `json:"Penalty,omitempty" xml:"Penalty,omitempty"`
}

func (s ClinkCreateQueueResponseBodyDataQueueQueueMembers) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueResponseBodyDataQueueQueueMembers) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueResponseBodyDataQueueQueueMembers) GetCno() *string {
	return s.Cno
}

func (s *ClinkCreateQueueResponseBodyDataQueueQueueMembers) GetPenalty() *int64 {
	return s.Penalty
}

func (s *ClinkCreateQueueResponseBodyDataQueueQueueMembers) SetCno(v string) *ClinkCreateQueueResponseBodyDataQueueQueueMembers {
	s.Cno = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueueQueueMembers) SetPenalty(v int64) *ClinkCreateQueueResponseBodyDataQueueQueueMembers {
	s.Penalty = &v
	return s
}

func (s *ClinkCreateQueueResponseBodyDataQueueQueueMembers) Validate() error {
	return dara.Validate(s)
}
