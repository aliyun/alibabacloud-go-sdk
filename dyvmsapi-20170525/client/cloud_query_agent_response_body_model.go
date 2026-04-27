// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudQueryAgentResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudQueryAgentResponseBody
	GetCode() *string
	SetData(v *CloudQueryAgentResponseBodyData) *CloudQueryAgentResponseBody
	GetData() *CloudQueryAgentResponseBodyData
	SetMessage(v string) *CloudQueryAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudQueryAgentResponseBody
	GetRequestId() *string
}

type CloudQueryAgentResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudQueryAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudQueryAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudQueryAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudQueryAgentResponseBody) GetData() *CloudQueryAgentResponseBodyData {
	return s.Data
}

func (s *CloudQueryAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudQueryAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudQueryAgentResponseBody) SetAccessDeniedDetail(v string) *CloudQueryAgentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudQueryAgentResponseBody) SetCode(v string) *CloudQueryAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CloudQueryAgentResponseBody) SetData(v *CloudQueryAgentResponseBodyData) *CloudQueryAgentResponseBody {
	s.Data = v
	return s
}

func (s *CloudQueryAgentResponseBody) SetMessage(v string) *CloudQueryAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CloudQueryAgentResponseBody) SetRequestId(v string) *CloudQueryAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudQueryAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudQueryAgentResponseBodyData struct {
	// 座席列表数组
	Agents []*CloudQueryAgentResponseBodyDataAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// 总数
	//
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CloudQueryAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponseBodyData) GetAgents() []*CloudQueryAgentResponseBodyDataAgents {
	return s.Agents
}

func (s *CloudQueryAgentResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *CloudQueryAgentResponseBodyData) SetAgents(v []*CloudQueryAgentResponseBodyDataAgents) *CloudQueryAgentResponseBodyData {
	s.Agents = v
	return s
}

func (s *CloudQueryAgentResponseBodyData) SetTotal(v int64) *CloudQueryAgentResponseBodyData {
	s.Total = &v
	return s
}

func (s *CloudQueryAgentResponseBodyData) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudQueryAgentResponseBodyDataAgents struct {
	// 座席信息
	Agent *CloudQueryAgentResponseBodyDataAgentsAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// 座席所属队列信息
	QueueList []*CloudQueryAgentResponseBodyDataAgentsQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
}

func (s CloudQueryAgentResponseBodyDataAgents) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponseBodyDataAgents) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponseBodyDataAgents) GetAgent() *CloudQueryAgentResponseBodyDataAgentsAgent {
	return s.Agent
}

func (s *CloudQueryAgentResponseBodyDataAgents) GetQueueList() []*CloudQueryAgentResponseBodyDataAgentsQueueList {
	return s.QueueList
}

func (s *CloudQueryAgentResponseBodyDataAgents) SetAgent(v *CloudQueryAgentResponseBodyDataAgentsAgent) *CloudQueryAgentResponseBodyDataAgents {
	s.Agent = v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgents) SetQueueList(v []*CloudQueryAgentResponseBodyDataAgentsQueueList) *CloudQueryAgentResponseBodyDataAgents {
	s.QueueList = v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgents) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	if s.QueueList != nil {
		for _, item := range s.QueueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudQueryAgentResponseBodyDataAgentsAgent struct {
	// 是否启用，0：停用，1：启用，默认启用
	//
	// example:
	//
	// 1
	Active *string `json:"Active,omitempty" xml:"Active,omitempty"`
	// 座席类型，1：电话座席，2：电脑座席，默认电话座席
	//
	// example:
	//
	// 1
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// 区号格式
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 座席绑定电话
	//
	// example:
	//
	// 41008502
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 电话类型，1:固话 2:手机 3:分机 4:软电话
	//
	// example:
	//
	// 1
	BindTelType *string `json:"BindTelType,omitempty" xml:"BindTelType,omitempty"`
	// 呼叫权限，0：无限制，1：国内长途，2：国内本市，3：内部呼叫，默认无限制
	//
	// example:
	//
	// 0
	CallPower *string `json:"CallPower,omitempty" xml:"CallPower,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
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
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 呼入是否录音，0：不录用，1：录音，默认录音
	//
	// example:
	//
	// 1
	IbRecord *string `json:"IbRecord,omitempty" xml:"IbRecord,omitempty"`
	// 座席id
	//
	// example:
	//
	// 133
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 是否开启ASR转写：0：不开启，1：开启，默认不开启
	//
	// example:
	//
	// 1
	IsAsr *string `json:"IsAsr,omitempty" xml:"IsAsr,omitempty"`
	// 是否允许外呼，0：不允许，1：可以，默认允许
	//
	// example:
	//
	// 1
	IsOb *string `json:"IsOb,omitempty" xml:"IsOb,omitempty"`
	// example:
	//
	// 1
	IsQualityCheck *string `json:"IsQualityCheck,omitempty" xml:"IsQualityCheck,omitempty"`
	// 座席姓名
	//
	// example:
	//
	// name1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 外显号码
	//
	// example:
	//
	// 41008502
	ObClid *string `json:"ObClid,omitempty" xml:"ObClid,omitempty"`
	// 外显规则 1：企业默认 2：随机 3：按区号 4：动态外显
	//
	// example:
	//
	// 1
	ObClidType *string `json:"ObClidType,omitempty" xml:"ObClidType,omitempty"`
	// 外呼是否录音，0：不录音，1：录音，默认录音
	//
	// example:
	//
	// 1
	ObRecord *string `json:"ObRecord,omitempty" xml:"ObRecord,omitempty"`
	// 1：班长席，0：普通座席，默认普通座席
	//
	// example:
	//
	// 1
	Power *string `json:"Power,omitempty" xml:"Power,omitempty"`
	// 座席状态，0:离线，1：在线
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// webrtc软电话返回地址，0：企业默认 1：公网域名 2：专线域名 3：公网IP 4：专线IP
	//
	// example:
	//
	// 0
	WebrtcUrlType *string `json:"WebrtcUrlType,omitempty" xml:"WebrtcUrlType,omitempty"`
	// 整理时间，秒数，默认10秒
	//
	// example:
	//
	// 10
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudQueryAgentResponseBodyDataAgentsAgent) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponseBodyDataAgentsAgent) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetActive() *string {
	return s.Active
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetAgentType() *string {
	return s.AgentType
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetAreaCode() *string {
	return s.AreaCode
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetBindTel() *string {
	return s.BindTel
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetBindTelType() *string {
	return s.BindTelType
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetCallPower() *string {
	return s.CallPower
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetCno() *string {
	return s.Cno
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetIbRecord() *string {
	return s.IbRecord
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetId() *string {
	return s.Id
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetIsAsr() *string {
	return s.IsAsr
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetIsOb() *string {
	return s.IsOb
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetIsQualityCheck() *string {
	return s.IsQualityCheck
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetName() *string {
	return s.Name
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetObClid() *string {
	return s.ObClid
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetObClidType() *string {
	return s.ObClidType
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetObRecord() *string {
	return s.ObRecord
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetPower() *string {
	return s.Power
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetStatus() *string {
	return s.Status
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetWebrtcUrlType() *string {
	return s.WebrtcUrlType
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetActive(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Active = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetAgentType(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.AgentType = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetAreaCode(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.AreaCode = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetBindTel(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.BindTel = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetBindTelType(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.BindTelType = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetCallPower(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.CallPower = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetCno(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Cno = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetCreateTime(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.CreateTime = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetEnterpriseId(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetIbRecord(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.IbRecord = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetId(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Id = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetIsAsr(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.IsAsr = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetIsOb(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.IsOb = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetIsQualityCheck(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.IsQualityCheck = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetName(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Name = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetObClid(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.ObClid = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetObClidType(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.ObClidType = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetObRecord(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.ObRecord = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetPower(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Power = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetStatus(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Status = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetWebrtcUrlType(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.WebrtcUrlType = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) SetWrapup(v string) *CloudQueryAgentResponseBodyDataAgentsAgent {
	s.Wrapup = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsAgent) Validate() error {
	return dara.Validate(s)
}

type CloudQueryAgentResponseBodyDataAgentsQueueList struct {
	// 位置播报 0关闭 1大于announce_position_param时播放 2小于等于announce_position_param时播放
	//
	// example:
	//
	// 1
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
	// 4
	AnnouncePositionParam *int64 `json:"AnnouncePositionParam,omitempty" xml:"AnnouncePositionParam,omitempty"`
	// 是否播报\\"您是下一位\\",0关闭，1开启
	//
	// example:
	//
	// 1
	AnnouncePositionYouarenext *int64 `json:"AnnouncePositionYouarenext,omitempty" xml:"AnnouncePositionYouarenext,omitempty"`
	// 播报固定语音 0关闭 1打开
	//
	// example:
	//
	// 1
	AnnounceSound *int64 `json:"AnnounceSound,omitempty" xml:"AnnounceSound,omitempty"`
	// 固定语音文件
	//
	// example:
	//
	// voice_filexxxxx
	AnnounceSoundFile *string `json:"AnnounceSoundFile,omitempty" xml:"AnnounceSoundFile,omitempty"`
	// 播放固定语音周期，秒数
	//
	// example:
	//
	// 17
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
	// queue_name
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
	// 76
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
	// class1
	MusicClass *string `json:"MusicClass,omitempty" xml:"MusicClass,omitempty"`
	// 队列号
	//
	// example:
	//
	// 322
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
	// 20
	Retry *int64 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// 语音报号，true:播报，false:不播报
	//
	// example:
	//
	// true
	SayAgentno *bool `json:"SayAgentno,omitempty" xml:"SayAgentno,omitempty"`
	// 服务水平秒数，低于此时间内接听的认为是高服务水平
	//
	// example:
	//
	// 5
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
	// 1
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

func (s CloudQueryAgentResponseBodyDataAgentsQueueList) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryAgentResponseBodyDataAgentsQueueList) GoString() string {
	return s.String()
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnouncePosition() *int64 {
	return s.AnnouncePosition
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnouncePositionFrequency() *int64 {
	return s.AnnouncePositionFrequency
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnouncePositionParam() *int64 {
	return s.AnnouncePositionParam
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnouncePositionYouarenext() *int64 {
	return s.AnnouncePositionYouarenext
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnounceSound() *int64 {
	return s.AnnounceSound
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnounceSoundFile() *string {
	return s.AnnounceSoundFile
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetAnnounceSoundFrequency() *int64 {
	return s.AnnounceSoundFrequency
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetDescription() *string {
	return s.Description
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetId() *int64 {
	return s.Id
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetMaxLen() *int64 {
	return s.MaxLen
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetMusicClass() *string {
	return s.MusicClass
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetQno() *string {
	return s.Qno
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetSayAgentno() *bool {
	return s.SayAgentno
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetStrategy() *string {
	return s.Strategy
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetWeight() *int64 {
	return s.Weight
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnouncePosition(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnouncePosition = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnouncePositionFrequency(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnouncePositionFrequency = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnouncePositionParam(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnouncePositionParam = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnouncePositionYouarenext(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnouncePositionYouarenext = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnounceSound(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnounceSound = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnounceSoundFile(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnounceSoundFile = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetAnnounceSoundFrequency(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.AnnounceSoundFrequency = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetCreateTime(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.CreateTime = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetDescription(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Description = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetEnterpriseId(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetId(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Id = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetJoinEmpty(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.JoinEmpty = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetMaxLen(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.MaxLen = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetMemberTimeout(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.MemberTimeout = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetMusicClass(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.MusicClass = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetQno(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Qno = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetQueueTimeout(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.QueueTimeout = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetRetry(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Retry = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetSayAgentno(v bool) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.SayAgentno = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetServiceLevel(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.ServiceLevel = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetStrategy(v string) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Strategy = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetVipSupport(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.VipSupport = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetWeight(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.Weight = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) SetWrapupTime(v int64) *CloudQueryAgentResponseBodyDataAgentsQueueList {
	s.WrapupTime = &v
	return s
}

func (s *CloudQueryAgentResponseBodyDataAgentsQueueList) Validate() error {
	return dara.Validate(s)
}
