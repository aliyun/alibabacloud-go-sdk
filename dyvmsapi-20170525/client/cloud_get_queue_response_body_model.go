// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudGetQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudGetQueueResponseBody
	GetCode() *string
	SetData(v *CloudGetQueueResponseBodyData) *CloudGetQueueResponseBody
	GetData() *CloudGetQueueResponseBodyData
	SetMessage(v string) *CloudGetQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudGetQueueResponseBody
	GetRequestId() *string
}

type CloudGetQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudGetQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudGetQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CloudGetQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudGetQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudGetQueueResponseBody) GetData() *CloudGetQueueResponseBodyData {
	return s.Data
}

func (s *CloudGetQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudGetQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudGetQueueResponseBody) SetAccessDeniedDetail(v string) *CloudGetQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudGetQueueResponseBody) SetCode(v string) *CloudGetQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CloudGetQueueResponseBody) SetData(v *CloudGetQueueResponseBodyData) *CloudGetQueueResponseBody {
	s.Data = v
	return s
}

func (s *CloudGetQueueResponseBody) SetMessage(v string) *CloudGetQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CloudGetQueueResponseBody) SetRequestId(v string) *CloudGetQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudGetQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudGetQueueResponseBodyData struct {
	// 队列配置信息
	Queue *CloudGetQueueResponseBodyDataQueue `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Struct"`
	// 队列所需技能数组
	QueueSkills []*CloudGetQueueResponseBodyDataQueueSkills `json:"QueueSkills,omitempty" xml:"QueueSkills,omitempty" type:"Repeated"`
}

func (s CloudGetQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudGetQueueResponseBodyData) GetQueue() *CloudGetQueueResponseBodyDataQueue {
	return s.Queue
}

func (s *CloudGetQueueResponseBodyData) GetQueueSkills() []*CloudGetQueueResponseBodyDataQueueSkills {
	return s.QueueSkills
}

func (s *CloudGetQueueResponseBodyData) SetQueue(v *CloudGetQueueResponseBodyDataQueue) *CloudGetQueueResponseBodyData {
	s.Queue = v
	return s
}

func (s *CloudGetQueueResponseBodyData) SetQueueSkills(v []*CloudGetQueueResponseBodyDataQueueSkills) *CloudGetQueueResponseBodyData {
	s.QueueSkills = v
	return s
}

func (s *CloudGetQueueResponseBodyData) Validate() error {
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

type CloudGetQueueResponseBodyDataQueue struct {
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
	// 0
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
	// ""
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
	// 2026-03-30 06:12:33
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
	// 54473
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
	// 60
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
	// 2632
	Qno *string `json:"Qno,omitempty" xml:"Qno,omitempty"`
	// 队列超时时间，取值范围5-600秒，默认600秒
	//
	// example:
	//
	// 600
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

func (s CloudGetQueueResponseBodyDataQueue) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueResponseBodyDataQueue) GoString() string {
	return s.String()
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnouncePosition() *int64 {
	return s.AnnouncePosition
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnouncePositionFrequency() *int64 {
	return s.AnnouncePositionFrequency
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnouncePositionParam() *int64 {
	return s.AnnouncePositionParam
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnouncePositionYouarenext() *int64 {
	return s.AnnouncePositionYouarenext
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnounceSound() *int64 {
	return s.AnnounceSound
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnounceSoundFile() *string {
	return s.AnnounceSoundFile
}

func (s *CloudGetQueueResponseBodyDataQueue) GetAnnounceSoundFrequency() *int64 {
	return s.AnnounceSoundFrequency
}

func (s *CloudGetQueueResponseBodyDataQueue) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudGetQueueResponseBodyDataQueue) GetDescription() *string {
	return s.Description
}

func (s *CloudGetQueueResponseBodyDataQueue) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetQueueResponseBodyDataQueue) GetId() *int64 {
	return s.Id
}

func (s *CloudGetQueueResponseBodyDataQueue) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *CloudGetQueueResponseBodyDataQueue) GetMaxLen() *int64 {
	return s.MaxLen
}

func (s *CloudGetQueueResponseBodyDataQueue) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *CloudGetQueueResponseBodyDataQueue) GetMusicClass() *string {
	return s.MusicClass
}

func (s *CloudGetQueueResponseBodyDataQueue) GetQno() *string {
	return s.Qno
}

func (s *CloudGetQueueResponseBodyDataQueue) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *CloudGetQueueResponseBodyDataQueue) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudGetQueueResponseBodyDataQueue) GetSayAgentno() *bool {
	return s.SayAgentno
}

func (s *CloudGetQueueResponseBodyDataQueue) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *CloudGetQueueResponseBodyDataQueue) GetStrategy() *string {
	return s.Strategy
}

func (s *CloudGetQueueResponseBodyDataQueue) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *CloudGetQueueResponseBodyDataQueue) GetWeight() *int64 {
	return s.Weight
}

func (s *CloudGetQueueResponseBodyDataQueue) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnouncePosition(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnouncePosition = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnouncePositionFrequency(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnouncePositionFrequency = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnouncePositionParam(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnouncePositionParam = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnouncePositionYouarenext(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnouncePositionYouarenext = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnounceSound(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnounceSound = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnounceSoundFile(v string) *CloudGetQueueResponseBodyDataQueue {
	s.AnnounceSoundFile = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetAnnounceSoundFrequency(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.AnnounceSoundFrequency = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetCreateTime(v string) *CloudGetQueueResponseBodyDataQueue {
	s.CreateTime = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetDescription(v string) *CloudGetQueueResponseBodyDataQueue {
	s.Description = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetEnterpriseId(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetId(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.Id = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetJoinEmpty(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.JoinEmpty = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetMaxLen(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.MaxLen = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetMemberTimeout(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.MemberTimeout = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetMusicClass(v string) *CloudGetQueueResponseBodyDataQueue {
	s.MusicClass = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetQno(v string) *CloudGetQueueResponseBodyDataQueue {
	s.Qno = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetQueueTimeout(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.QueueTimeout = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetRetry(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.Retry = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetSayAgentno(v bool) *CloudGetQueueResponseBodyDataQueue {
	s.SayAgentno = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetServiceLevel(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.ServiceLevel = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetStrategy(v string) *CloudGetQueueResponseBodyDataQueue {
	s.Strategy = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetVipSupport(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.VipSupport = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetWeight(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.Weight = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) SetWrapupTime(v int64) *CloudGetQueueResponseBodyDataQueue {
	s.WrapupTime = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueue) Validate() error {
	return dara.Validate(s)
}

type CloudGetQueueResponseBodyDataQueueSkills struct {
	// 创建时间，格式: yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 2026-03-30 06:12:33
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
	// 51937
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

func (s CloudGetQueueResponseBodyDataQueueSkills) String() string {
	return dara.Prettify(s)
}

func (s CloudGetQueueResponseBodyDataQueueSkills) GoString() string {
	return s.String()
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetId() *int64 {
	return s.Id
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetQueueId() *int64 {
	return s.QueueId
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetSkillId() *int64 {
	return s.SkillId
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) GetSkillLevel() *int64 {
	return s.SkillLevel
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetCreateTime(v string) *CloudGetQueueResponseBodyDataQueueSkills {
	s.CreateTime = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetEnterpriseId(v int64) *CloudGetQueueResponseBodyDataQueueSkills {
	s.EnterpriseId = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetId(v int64) *CloudGetQueueResponseBodyDataQueueSkills {
	s.Id = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetQueueId(v int64) *CloudGetQueueResponseBodyDataQueueSkills {
	s.QueueId = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetSkillId(v int64) *CloudGetQueueResponseBodyDataQueueSkills {
	s.SkillId = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) SetSkillLevel(v int64) *CloudGetQueueResponseBodyDataQueueSkills {
	s.SkillLevel = &v
	return s
}

func (s *CloudGetQueueResponseBodyDataQueueSkills) Validate() error {
	return dara.Validate(s)
}
