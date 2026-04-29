// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *CloudCreateQueueRequest
	GetEnterpriseId() *int64
	SetOwnerId(v int64) *CloudCreateQueueRequest
	GetOwnerId() *int64
	SetQueue(v *CloudCreateQueueRequestQueue) *CloudCreateQueueRequest
	GetQueue() *CloudCreateQueueRequestQueue
	SetQueueSkills(v []*CloudCreateQueueRequestQueueSkills) *CloudCreateQueueRequest
	GetQueueSkills() []*CloudCreateQueueRequestQueueSkills
	SetResourceOwnerAccount(v string) *CloudCreateQueueRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloudCreateQueueRequest
	GetResourceOwnerId() *int64
}

type CloudCreateQueueRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 队列信息
	//
	// This parameter is required.
	Queue *CloudCreateQueueRequestQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
	// 队列技能
	//
	// This parameter is required.
	QueueSkills          []*CloudCreateQueueRequestQueueSkills `json:"QueueSkills,omitempty" xml:"QueueSkills,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CloudCreateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloudCreateQueueRequest) GetQueue() *CloudCreateQueueRequestQueue {
	return s.Queue
}

func (s *CloudCreateQueueRequest) GetQueueSkills() []*CloudCreateQueueRequestQueueSkills {
	return s.QueueSkills
}

func (s *CloudCreateQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloudCreateQueueRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloudCreateQueueRequest) SetEnterpriseId(v int64) *CloudCreateQueueRequest {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueRequest) SetOwnerId(v int64) *CloudCreateQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CloudCreateQueueRequest) SetQueue(v *CloudCreateQueueRequestQueue) *CloudCreateQueueRequest {
	s.Queue = v
	return s
}

func (s *CloudCreateQueueRequest) SetQueueSkills(v []*CloudCreateQueueRequestQueueSkills) *CloudCreateQueueRequest {
	s.QueueSkills = v
	return s
}

func (s *CloudCreateQueueRequest) SetResourceOwnerAccount(v string) *CloudCreateQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloudCreateQueueRequest) SetResourceOwnerId(v int64) *CloudCreateQueueRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloudCreateQueueRequest) Validate() error {
	if s.Queue != nil {
		if err := s.Queue.Validate(); err != nil {
			return err
		}
	}
	if s.QueueSkills != nil {
		for _, item := range s.QueueSkills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudCreateQueueRequestQueue struct {
	// 位置播报 0关闭 1大于announce_position_param时播放 2小于等于announce_position_param时播放
	//
	// example:
	//
	// 0
	AnnouncePosition *int64 `json:"AnnouncePosition,omitempty" xml:"AnnouncePosition,omitempty"`
	// 位置播报周期，秒数
	//
	// example:
	//
	// 0
	AnnouncePositionFrequency *int64 `json:"AnnouncePositionFrequency,omitempty" xml:"AnnouncePositionFrequency,omitempty"`
	// 多余少余n个时播报，必须大于等于2
	//
	// example:
	//
	// 2
	AnnouncePositionParam *int64 `json:"AnnouncePositionParam,omitempty" xml:"AnnouncePositionParam,omitempty"`
	// 是否播报\\"您是下一位\\",0关闭，1开启
	//
	// example:
	//
	// 0
	AnnouncePositionYouarenext *int64 `json:"AnnouncePositionYouarenext,omitempty" xml:"AnnouncePositionYouarenext,omitempty"`
	// 播报固定语音 0关闭 1打开
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	AnnounceSound *int64 `json:"AnnounceSound,omitempty" xml:"AnnounceSound,omitempty"`
	// 固定语音文件
	//
	// example:
	//
	// announceSoundFile3
	AnnounceSoundFile *string `json:"AnnounceSoundFile,omitempty" xml:"AnnounceSoundFile,omitempty"`
	// 播放固定语音周期，秒数
	//
	// example:
	//
	// 10
	AnnounceSoundFrequency *int64 `json:"AnnounceSoundFrequency,omitempty" xml:"AnnounceSoundFrequency,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 队列名
	//
	// This parameter is required.
	//
	// example:
	//
	// queueName
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 队列中为空时是否可以join，取值：1：置忙 2：通话中 4：振铃 8：无效 16：整理
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	JoinEmpty *int64 `json:"JoinEmpty,omitempty" xml:"JoinEmpty,omitempty"`
	// 最大等待数，设置范围0-999，0表示不做任何限制
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MaxLen *int64 `json:"MaxLen,omitempty" xml:"MaxLen,omitempty"`
	// 座席超时时间，取值范围20-60秒，默认25秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	MemberTimeout *int64 `json:"MemberTimeout,omitempty" xml:"MemberTimeout,omitempty"`
	// 等待语音class
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	MusicClass *string `json:"MusicClass,omitempty" xml:"MusicClass,omitempty"`
	// 队列号
	//
	// This parameter is required.
	//
	// example:
	//
	// 3003
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列超时时间，取值范围20-600秒，默认600秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	QueueTimeout *int64 `json:"QueueTimeout,omitempty" xml:"QueueTimeout,omitempty"`
	// 座席超时无应答,呼叫下一座席的延迟秒数
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 语音报号，true:播报，false:不播报
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	SayAgentno *bool `json:"SayAgentno,omitempty" xml:"SayAgentno,omitempty"`
	// 服务水平秒数，低于此时间内接听的认为是高服务水平
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ServiceLevel *int64 `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
	// 呼叫策略：rrordered:技能优先 rrmemory:轮选 fewestcalls:平均 random:随机 linear:顺序 leastrecent:最长空闲时间
	//
	// This parameter is required.
	//
	// example:
	//
	// rrordered
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// 是否支持vip，1:支持，0:不支持
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	VipSupport *int64 `json:"VipSupport,omitempty" xml:"VipSupport,omitempty"`
	// 队列优先级，取值范围1-10，数值越高，优先级越高
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时间，取值范围3-3600秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s CloudCreateQueueRequestQueue) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueRequestQueue) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueRequestQueue) GetAnnouncePosition() *int64 {
	return s.AnnouncePosition
}

func (s *CloudCreateQueueRequestQueue) GetAnnouncePositionFrequency() *int64 {
	return s.AnnouncePositionFrequency
}

func (s *CloudCreateQueueRequestQueue) GetAnnouncePositionParam() *int64 {
	return s.AnnouncePositionParam
}

func (s *CloudCreateQueueRequestQueue) GetAnnouncePositionYouarenext() *int64 {
	return s.AnnouncePositionYouarenext
}

func (s *CloudCreateQueueRequestQueue) GetAnnounceSound() *int64 {
	return s.AnnounceSound
}

func (s *CloudCreateQueueRequestQueue) GetAnnounceSoundFile() *string {
	return s.AnnounceSoundFile
}

func (s *CloudCreateQueueRequestQueue) GetAnnounceSoundFrequency() *int64 {
	return s.AnnounceSoundFrequency
}

func (s *CloudCreateQueueRequestQueue) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateQueueRequestQueue) GetDescription() *string {
	return s.Description
}

func (s *CloudCreateQueueRequestQueue) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *CloudCreateQueueRequestQueue) GetMaxLen() *int64 {
	return s.MaxLen
}

func (s *CloudCreateQueueRequestQueue) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *CloudCreateQueueRequestQueue) GetMusicClass() *string {
	return s.MusicClass
}

func (s *CloudCreateQueueRequestQueue) GetQno() *string {
	return s.Qno
}

func (s *CloudCreateQueueRequestQueue) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *CloudCreateQueueRequestQueue) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudCreateQueueRequestQueue) GetSayAgentno() *bool {
	return s.SayAgentno
}

func (s *CloudCreateQueueRequestQueue) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *CloudCreateQueueRequestQueue) GetStrategy() *string {
	return s.Strategy
}

func (s *CloudCreateQueueRequestQueue) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *CloudCreateQueueRequestQueue) GetWeight() *int64 {
	return s.Weight
}

func (s *CloudCreateQueueRequestQueue) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *CloudCreateQueueRequestQueue) SetAnnouncePosition(v int64) *CloudCreateQueueRequestQueue {
	s.AnnouncePosition = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnouncePositionFrequency(v int64) *CloudCreateQueueRequestQueue {
	s.AnnouncePositionFrequency = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnouncePositionParam(v int64) *CloudCreateQueueRequestQueue {
	s.AnnouncePositionParam = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnouncePositionYouarenext(v int64) *CloudCreateQueueRequestQueue {
	s.AnnouncePositionYouarenext = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnounceSound(v int64) *CloudCreateQueueRequestQueue {
	s.AnnounceSound = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnounceSoundFile(v string) *CloudCreateQueueRequestQueue {
	s.AnnounceSoundFile = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetAnnounceSoundFrequency(v int64) *CloudCreateQueueRequestQueue {
	s.AnnounceSoundFrequency = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetCreateTime(v string) *CloudCreateQueueRequestQueue {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetDescription(v string) *CloudCreateQueueRequestQueue {
	s.Description = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetJoinEmpty(v int64) *CloudCreateQueueRequestQueue {
	s.JoinEmpty = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetMaxLen(v int64) *CloudCreateQueueRequestQueue {
	s.MaxLen = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetMemberTimeout(v int64) *CloudCreateQueueRequestQueue {
	s.MemberTimeout = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetMusicClass(v string) *CloudCreateQueueRequestQueue {
	s.MusicClass = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetQno(v string) *CloudCreateQueueRequestQueue {
	s.Qno = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetQueueTimeout(v int64) *CloudCreateQueueRequestQueue {
	s.QueueTimeout = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetRetry(v int64) *CloudCreateQueueRequestQueue {
	s.Retry = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetSayAgentno(v bool) *CloudCreateQueueRequestQueue {
	s.SayAgentno = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetServiceLevel(v int64) *CloudCreateQueueRequestQueue {
	s.ServiceLevel = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetStrategy(v string) *CloudCreateQueueRequestQueue {
	s.Strategy = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetVipSupport(v int64) *CloudCreateQueueRequestQueue {
	s.VipSupport = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetWeight(v int64) *CloudCreateQueueRequestQueue {
	s.Weight = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) SetWrapupTime(v int64) *CloudCreateQueueRequestQueue {
	s.WrapupTime = &v
	return s
}

func (s *CloudCreateQueueRequestQueue) Validate() error {
	return dara.Validate(s)
}

type CloudCreateQueueRequestQueueSkills struct {
	// skill的id
	//
	// This parameter is required.
	//
	// example:
	//
	// 66
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudCreateQueueRequestQueueSkills) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueRequestQueueSkills) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueRequestQueueSkills) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudCreateQueueRequestQueueSkills) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudCreateQueueRequestQueueSkills) SetSkillId(v int64) *CloudCreateQueueRequestQueueSkills {
	s.SkillId = &v
	return s
}

func (s *CloudCreateQueueRequestQueueSkills) SetSkillLevel(v int64) *CloudCreateQueueRequestQueueSkills {
	s.SkillLevel = &v
	return s
}

func (s *CloudCreateQueueRequestQueueSkills) Validate() error {
	return dara.Validate(s)
}
