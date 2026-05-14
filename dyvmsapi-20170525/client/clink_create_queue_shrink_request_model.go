// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatLocation(v int64) *ClinkCreateQueueShrinkRequest
	GetChatLocation() *int64
	SetChatMaxWait(v int64) *ClinkCreateQueueShrinkRequest
	GetChatMaxWait() *int64
	SetChatStrategy(v int64) *ClinkCreateQueueShrinkRequest
	GetChatStrategy() *int64
	SetEnterpriseId(v int64) *ClinkCreateQueueShrinkRequest
	GetEnterpriseId() *int64
	SetIbAllowed(v int64) *ClinkCreateQueueShrinkRequest
	GetIbAllowed() *int64
	SetJoinEmpty(v int64) *ClinkCreateQueueShrinkRequest
	GetJoinEmpty() *int64
	SetMaxPauseClientFlag(v int64) *ClinkCreateQueueShrinkRequest
	GetMaxPauseClientFlag() *int64
	SetMaxPauseClientType(v int64) *ClinkCreateQueueShrinkRequest
	GetMaxPauseClientType() *int64
	SetMaxPauseClientValue(v int64) *ClinkCreateQueueShrinkRequest
	GetMaxPauseClientValue() *int64
	SetMaxWait(v int64) *ClinkCreateQueueShrinkRequest
	GetMaxWait() *int64
	SetMemberTimeout(v int64) *ClinkCreateQueueShrinkRequest
	GetMemberTimeout() *int64
	SetName(v string) *ClinkCreateQueueShrinkRequest
	GetName() *string
	SetOwnerId(v int64) *ClinkCreateQueueShrinkRequest
	GetOwnerId() *int64
	SetQno(v string) *ClinkCreateQueueShrinkRequest
	GetQno() *string
	SetQueueMembersShrink(v string) *ClinkCreateQueueShrinkRequest
	GetQueueMembersShrink() *string
	SetQueueTimeout(v int64) *ClinkCreateQueueShrinkRequest
	GetQueueTimeout() *int64
	SetResourceOwnerAccount(v string) *ClinkCreateQueueShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCreateQueueShrinkRequest
	GetResourceOwnerId() *int64
	SetSayCno(v int64) *ClinkCreateQueueShrinkRequest
	GetSayCno() *int64
	SetServiceLevel(v int64) *ClinkCreateQueueShrinkRequest
	GetServiceLevel() *int64
	SetStrategy(v int64) *ClinkCreateQueueShrinkRequest
	GetStrategy() *int64
	SetVipSupport(v int64) *ClinkCreateQueueShrinkRequest
	GetVipSupport() *int64
	SetWeight(v int64) *ClinkCreateQueueShrinkRequest
	GetWeight() *int64
	SetWrapupTime(v int64) *ClinkCreateQueueShrinkRequest
	GetWrapupTime() *int64
}

type ClinkCreateQueueShrinkRequest struct {
	// 在线客服排队位置推送(小于该位置则推送)；默认值为 10
	//
	// example:
	//
	// 10
	ChatLocation *int64 `json:"ChatLocation,omitempty" xml:"ChatLocation,omitempty"`
	// 在线客服最大排队数；默认值为 20
	//
	// example:
	//
	// 20
	ChatMaxWait *int64 `json:"ChatMaxWait,omitempty" xml:"ChatMaxWait,omitempty"`
	// 在线客服分配策略，2：轮选；4：随机；5：技能优先；默认值为 2
	//
	// example:
	//
	// 2
	ChatStrategy *int64 `json:"ChatStrategy,omitempty" xml:"ChatStrategy,omitempty"`
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否允许呼入队列，0: 否；1: 是；默认值为 1
	//
	// example:
	//
	// 1
	IbAllowed *int64 `json:"IbAllowed,omitempty" xml:"IbAllowed,omitempty"`
	// 队列中为空时是否可以join，1: 置忙；2: 通话中；4: 振铃；8: 无效；16: 整理 如选多种状态则传对应数值的和，如选置忙和通话中，值为 3。默认值为 0，即都未选中
	//
	// example:
	//
	// 1
	JoinEmpty *int64 `json:"JoinEmpty,omitempty" xml:"JoinEmpty,omitempty"`
	// 最大并发置忙座席数，开关配置，0:关闭，1:开启；默认值为 0
	//
	// example:
	//
	// 0
	MaxPauseClientFlag *int64 `json:"MaxPauseClientFlag,omitempty" xml:"MaxPauseClientFlag,omitempty"`
	// 最大并发置忙座席数，配置类型，0:百分比，1:数值；默认值为 0
	//
	// example:
	//
	// 0
	MaxPauseClientType *int64 `json:"MaxPauseClientType,omitempty" xml:"MaxPauseClientType,omitempty"`
	// 最大并发置忙座席数，配置值，当配置类型为百分比时最大值为100，默认值为 0
	//
	// example:
	//
	// 0
	MaxPauseClientValue *int64 `json:"MaxPauseClientValue,omitempty" xml:"MaxPauseClientValue,omitempty"`
	// 最大等待数，设置范围 0-999，0 表示不做任何限制。默认值为 5
	//
	// example:
	//
	// 5
	MaxWait *int64 `json:"MaxWait,omitempty" xml:"MaxWait,omitempty"`
	// 座席超时时间，取值范围 20-60 秒。默认值为 25
	//
	// example:
	//
	// 26
	MemberTimeout *int64 `json:"MemberTimeout,omitempty" xml:"MemberTimeout,omitempty"`
	// 队列名称，只允许输入中文，字母，数字，下划线，长度不超过 20 个字符
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列工号，4 位数字，要求唯一
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列内座席及技能值设置。技能值 1-10，技能值 1 的优先级为最高 cno 队列下座席号 penalty该座席在队列中的优先级
	QueueMembersShrink *string `json:"QueueMembers,omitempty" xml:"QueueMembers,omitempty"`
	// 队列超时时间，取值范围 30-600 秒。默认值为 600
	//
	// example:
	//
	// 33
	QueueTimeout         *int64  `json:"QueueTimeout,omitempty" xml:"QueueTimeout,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 是否报座席号，0: 否；1:是；默认值为 0
	//
	// example:
	//
	// 0
	SayCno *int64 `json:"SayCno,omitempty" xml:"SayCno,omitempty"`
	// 服务水平秒数，取值范围 10-30 秒，低于此时间内接听的认为是高服务水平。默认值为 15
	//
	// example:
	//
	// 15
	ServiceLevel *int64 `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
	// 呼叫策略， 1：顺序，2：轮选，3：平均，4：随机，5：技能优先，6：距上一通呼入空闲时间累计最长，7：技能优先+距上一通呼入空闲时间累计最长，8：距上一通呼入空闲时间累计最长(全队列)，9：技能优先+距上一通呼入空闲时间累计最长(全队列)，10：全队列平均，11：技能优先+轮选，12：最长空闲时长，13：技能优先+当前空闲状态时长最长，14：距上一通通话空闲时间累计最长，默认值为 2
	//
	// example:
	//
	// 2
	Strategy *int64 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// VIP客户优先，0:不支持；1:支持；默认值为 0
	//
	// example:
	//
	// 0
	VipSupport *int64 `json:"VipSupport,omitempty" xml:"VipSupport,omitempty"`
	// 队列优先级，取值范围 1-10。默认值为 1 座席属于多个队列时，优先级高的队列的来电将优先接听；数值越高，优先级越高
	//
	// example:
	//
	// 4
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时长，取值范围 3-300 秒，整理期间座席不接受新的呼叫。默认值为 15
	//
	// example:
	//
	// 15
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s ClinkCreateQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *ClinkCreateQueueShrinkRequest) GetChatLocation() *int64 {
	return s.ChatLocation
}

func (s *ClinkCreateQueueShrinkRequest) GetChatMaxWait() *int64 {
	return s.ChatMaxWait
}

func (s *ClinkCreateQueueShrinkRequest) GetChatStrategy() *int64 {
	return s.ChatStrategy
}

func (s *ClinkCreateQueueShrinkRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCreateQueueShrinkRequest) GetIbAllowed() *int64 {
	return s.IbAllowed
}

func (s *ClinkCreateQueueShrinkRequest) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *ClinkCreateQueueShrinkRequest) GetMaxPauseClientFlag() *int64 {
	return s.MaxPauseClientFlag
}

func (s *ClinkCreateQueueShrinkRequest) GetMaxPauseClientType() *int64 {
	return s.MaxPauseClientType
}

func (s *ClinkCreateQueueShrinkRequest) GetMaxPauseClientValue() *int64 {
	return s.MaxPauseClientValue
}

func (s *ClinkCreateQueueShrinkRequest) GetMaxWait() *int64 {
	return s.MaxWait
}

func (s *ClinkCreateQueueShrinkRequest) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *ClinkCreateQueueShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ClinkCreateQueueShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCreateQueueShrinkRequest) GetQno() *string {
	return s.Qno
}

func (s *ClinkCreateQueueShrinkRequest) GetQueueMembersShrink() *string {
	return s.QueueMembersShrink
}

func (s *ClinkCreateQueueShrinkRequest) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *ClinkCreateQueueShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCreateQueueShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCreateQueueShrinkRequest) GetSayCno() *int64 {
	return s.SayCno
}

func (s *ClinkCreateQueueShrinkRequest) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *ClinkCreateQueueShrinkRequest) GetStrategy() *int64 {
	return s.Strategy
}

func (s *ClinkCreateQueueShrinkRequest) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *ClinkCreateQueueShrinkRequest) GetWeight() *int64 {
	return s.Weight
}

func (s *ClinkCreateQueueShrinkRequest) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *ClinkCreateQueueShrinkRequest) SetChatLocation(v int64) *ClinkCreateQueueShrinkRequest {
	s.ChatLocation = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetChatMaxWait(v int64) *ClinkCreateQueueShrinkRequest {
	s.ChatMaxWait = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetChatStrategy(v int64) *ClinkCreateQueueShrinkRequest {
	s.ChatStrategy = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetEnterpriseId(v int64) *ClinkCreateQueueShrinkRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetIbAllowed(v int64) *ClinkCreateQueueShrinkRequest {
	s.IbAllowed = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetJoinEmpty(v int64) *ClinkCreateQueueShrinkRequest {
	s.JoinEmpty = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetMaxPauseClientFlag(v int64) *ClinkCreateQueueShrinkRequest {
	s.MaxPauseClientFlag = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetMaxPauseClientType(v int64) *ClinkCreateQueueShrinkRequest {
	s.MaxPauseClientType = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetMaxPauseClientValue(v int64) *ClinkCreateQueueShrinkRequest {
	s.MaxPauseClientValue = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetMaxWait(v int64) *ClinkCreateQueueShrinkRequest {
	s.MaxWait = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetMemberTimeout(v int64) *ClinkCreateQueueShrinkRequest {
	s.MemberTimeout = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetName(v string) *ClinkCreateQueueShrinkRequest {
	s.Name = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetOwnerId(v int64) *ClinkCreateQueueShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetQno(v string) *ClinkCreateQueueShrinkRequest {
	s.Qno = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetQueueMembersShrink(v string) *ClinkCreateQueueShrinkRequest {
	s.QueueMembersShrink = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetQueueTimeout(v int64) *ClinkCreateQueueShrinkRequest {
	s.QueueTimeout = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetResourceOwnerAccount(v string) *ClinkCreateQueueShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetResourceOwnerId(v int64) *ClinkCreateQueueShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetSayCno(v int64) *ClinkCreateQueueShrinkRequest {
	s.SayCno = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetServiceLevel(v int64) *ClinkCreateQueueShrinkRequest {
	s.ServiceLevel = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetStrategy(v int64) *ClinkCreateQueueShrinkRequest {
	s.Strategy = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetVipSupport(v int64) *ClinkCreateQueueShrinkRequest {
	s.VipSupport = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetWeight(v int64) *ClinkCreateQueueShrinkRequest {
	s.Weight = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) SetWrapupTime(v int64) *ClinkCreateQueueShrinkRequest {
	s.WrapupTime = &v
	return s
}

func (s *ClinkCreateQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
