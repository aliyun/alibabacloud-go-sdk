// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateQueueResponseBody
	GetCode() *string
	SetData(v *CloudCreateQueueResponseBodyData) *CloudCreateQueueResponseBody
	GetData() *CloudCreateQueueResponseBodyData
	SetMessage(v string) *CloudCreateQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateQueueResponseBody
	GetRequestId() *string
}

type CloudCreateQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateQueueResponseBody) GetData() *CloudCreateQueueResponseBodyData {
	return s.Data
}

func (s *CloudCreateQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateQueueResponseBody) SetAccessDeniedDetail(v string) *CloudCreateQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateQueueResponseBody) SetCode(v string) *CloudCreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateQueueResponseBody) SetData(v *CloudCreateQueueResponseBodyData) *CloudCreateQueueResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateQueueResponseBody) SetMessage(v string) *CloudCreateQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateQueueResponseBody) SetRequestId(v string) *CloudCreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateQueueResponseBodyData struct {
	// 队列配置信息
	Queue *CloudCreateQueueResponseBodyDataQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
	// 队列所需技能数组
	QueueSkills []*CloudCreateQueueResponseBodyDataQueueSkills `json:"QueueSkills,omitempty" xml:"QueueSkills,omitempty" type:"Repeated"`
}

func (s CloudCreateQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueResponseBodyData) GetQueue() *CloudCreateQueueResponseBodyDataQueue {
	return s.Queue
}

func (s *CloudCreateQueueResponseBodyData) GetQueueSkills() []*CloudCreateQueueResponseBodyDataQueueSkills {
	return s.QueueSkills
}

func (s *CloudCreateQueueResponseBodyData) SetQueue(v *CloudCreateQueueResponseBodyDataQueue) *CloudCreateQueueResponseBodyData {
	s.Queue = v
	return s
}

func (s *CloudCreateQueueResponseBodyData) SetQueueSkills(v []*CloudCreateQueueResponseBodyDataQueueSkills) *CloudCreateQueueResponseBodyData {
	s.QueueSkills = v
	return s
}

func (s *CloudCreateQueueResponseBodyData) Validate() error {
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

type CloudCreateQueueResponseBodyDataQueue struct {
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
	// 10
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
	// example:
	//
	// 0
	AnnounceSound *int64 `json:"AnnounceSound,omitempty" xml:"AnnounceSound,omitempty"`
	// 固定语音文件
	//
	// example:
	//
	// sound_file3
	AnnounceSoundFile *string `json:"AnnounceSoundFile,omitempty" xml:"AnnounceSoundFile,omitempty"`
	// 播放固定语音周期，秒数
	//
	// example:
	//
	// 0
	AnnounceSoundFrequency *int64 `json:"AnnounceSoundFrequency,omitempty" xml:"AnnounceSoundFrequency,omitempty"`
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 队列名
	//
	// example:
	//
	// queueName
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 队列id
	//
	// example:
	//
	// 80
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列中为空时是否可以join，取值：1：置忙 2：通话中 4：振铃 8：无效 16：整理
	//
	// example:
	//
	// 1
	JoinEmpty *int64 `json:"JoinEmpty,omitempty" xml:"JoinEmpty,omitempty"`
	// 最大等待数，设置范围0-999，0表示不做任何限制
	//
	// example:
	//
	// 0
	MaxLen *int64 `json:"MaxLen,omitempty" xml:"MaxLen,omitempty"`
	// 座席超时时间，取值范围20-60秒，默认25秒
	//
	// example:
	//
	// 25
	MemberTimeout *int64 `json:"MemberTimeout,omitempty" xml:"MemberTimeout,omitempty"`
	// 等待语音class
	//
	// example:
	//
	// default
	MusicClass *string `json:"MusicClass,omitempty" xml:"MusicClass,omitempty"`
	// 队列号
	//
	// example:
	//
	// 12331
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列超时时间，取值范围5-600秒，默认600秒
	//
	// example:
	//
	// 60
	QueueTimeout *int64 `json:"QueueTimeout,omitempty" xml:"QueueTimeout,omitempty"`
	// 座席超时无应答,呼叫下一座席的延迟秒数
	//
	// example:
	//
	// 5
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 语音报号，true:播报，false:不播报
	//
	// example:
	//
	// false
	SayAgentno *bool `json:"SayAgentno,omitempty" xml:"SayAgentno,omitempty"`
	// 服务水平秒数，低于此时间内接听的认为是高服务水平
	//
	// example:
	//
	// 10
	ServiceLevel *int64 `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
	// 呼叫策略：rrordered:技能优先 rrmemory:轮选 fewestcalls:平均 random:随机 linear:顺序 leastrecent:最长空闲时间
	//
	// example:
	//
	// rrordered
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// 是否支持vip，1:支持，0:不支持
	//
	// example:
	//
	// 0
	VipSupport *int64 `json:"VipSupport,omitempty" xml:"VipSupport,omitempty"`
	// 队列优先级，取值范围1-10，数值越高，优先级越高
	//
	// example:
	//
	// 10
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时间，取值范围3-3600秒
	//
	// example:
	//
	// 30
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s CloudCreateQueueResponseBodyDataQueue) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueResponseBodyDataQueue) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnouncePosition() *int64 {
	return s.AnnouncePosition
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnouncePositionFrequency() *int64 {
	return s.AnnouncePositionFrequency
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnouncePositionParam() *int64 {
	return s.AnnouncePositionParam
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnouncePositionYouarenext() *int64 {
	return s.AnnouncePositionYouarenext
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnounceSound() *int64 {
	return s.AnnounceSound
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnounceSoundFile() *string {
	return s.AnnounceSoundFile
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetAnnounceSoundFrequency() *int64 {
	return s.AnnounceSoundFrequency
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetDescription() *string {
	return s.Description
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetId() *int64 {
	return s.Id
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetMaxLen() *int64 {
	return s.MaxLen
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetMusicClass() *string {
	return s.MusicClass
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetQno() *string {
	return s.Qno
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetSayAgentno() *bool {
	return s.SayAgentno
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetStrategy() *string {
	return s.Strategy
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetWeight() *int64 {
	return s.Weight
}

func (s *CloudCreateQueueResponseBodyDataQueue) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnouncePosition(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnouncePosition = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnouncePositionFrequency(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnouncePositionFrequency = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnouncePositionParam(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnouncePositionParam = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnouncePositionYouarenext(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnouncePositionYouarenext = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnounceSound(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnounceSound = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnounceSoundFile(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnounceSoundFile = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetAnnounceSoundFrequency(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.AnnounceSoundFrequency = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetCreateTime(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetDescription(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.Description = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetEnterpriseId(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetId(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.Id = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetJoinEmpty(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.JoinEmpty = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetMaxLen(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.MaxLen = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetMemberTimeout(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.MemberTimeout = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetMusicClass(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.MusicClass = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetQno(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.Qno = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetQueueTimeout(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.QueueTimeout = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetRetry(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.Retry = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetSayAgentno(v bool) *CloudCreateQueueResponseBodyDataQueue {
	s.SayAgentno = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetServiceLevel(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.ServiceLevel = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetStrategy(v string) *CloudCreateQueueResponseBodyDataQueue {
	s.Strategy = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetVipSupport(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.VipSupport = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetWeight(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.Weight = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) SetWrapupTime(v int64) *CloudCreateQueueResponseBodyDataQueue {
	s.WrapupTime = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueue) Validate() error {
	return dara.Validate(s)
}

type CloudCreateQueueResponseBodyDataQueueSkills struct {
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 企业编号
	//
	// example:
	//
	// 7000002
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// queueSkill关系表中id
	//
	// example:
	//
	// 54474
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列id
	//
	// example:
	//
	// 51937
	QueueId *int64 `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// skill的id
	//
	// example:
	//
	// 48734
	SkillId *int64 `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// 技能值
	//
	// example:
	//
	// 5
	SkillLevel *int64 `json:"SkillLevel,omitempty" xml:"SkillLevel,omitempty"`
}

func (s CloudCreateQueueResponseBodyDataQueueSkills) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueResponseBodyDataQueueSkills) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetId() *int64 {
	return s.Id
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetQueueId() *int64 {
	return s.QueueId
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetCreateTime(v string) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetEnterpriseId(v int64) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetId(v int64) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.Id = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetQueueId(v int64) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.QueueId = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetSkillId(v int64) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.SkillId = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) SetSkillLevel(v int64) *CloudCreateQueueResponseBodyDataQueueSkills {
	s.SkillLevel = &v
	return s
}

func (s *CloudCreateQueueResponseBodyDataQueueSkills) Validate() error {
	return dara.Validate(s)
}
