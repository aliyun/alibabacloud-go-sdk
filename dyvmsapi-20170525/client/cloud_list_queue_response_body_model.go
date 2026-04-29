// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudListQueueResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudListQueueResponseBody
	GetCode() *string
	SetData(v *CloudListQueueResponseBodyData) *CloudListQueueResponseBody
	GetData() *CloudListQueueResponseBodyData
	SetMessage(v string) *CloudListQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudListQueueResponseBody
	GetRequestId() *string
}

type CloudListQueueResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudListQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudListQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CloudListQueueResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudListQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudListQueueResponseBody) GetData() *CloudListQueueResponseBodyData {
	return s.Data
}

func (s *CloudListQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudListQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudListQueueResponseBody) SetAccessDeniedDetail(v string) *CloudListQueueResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudListQueueResponseBody) SetCode(v string) *CloudListQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CloudListQueueResponseBody) SetData(v *CloudListQueueResponseBodyData) *CloudListQueueResponseBody {
	s.Data = v
	return s
}

func (s *CloudListQueueResponseBody) SetMessage(v string) *CloudListQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CloudListQueueResponseBody) SetRequestId(v string) *CloudListQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudListQueueResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudListQueueResponseBodyData struct {
	// 队列列表
	List []*CloudListQueueResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// 总条数
	//
	// example:
	//
	// 23
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CloudListQueueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudListQueueResponseBodyData) GetList() []*CloudListQueueResponseBodyDataList {
	return s.List
}

func (s *CloudListQueueResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *CloudListQueueResponseBodyData) SetList(v []*CloudListQueueResponseBodyDataList) *CloudListQueueResponseBodyData {
	s.List = v
	return s
}

func (s *CloudListQueueResponseBodyData) SetTotal(v int64) *CloudListQueueResponseBodyData {
	s.Total = &v
	return s
}

func (s *CloudListQueueResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudListQueueResponseBodyDataList struct {
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
	// 2025-09-06 15:50:18
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
	// 48175
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 队列中为空时是否可以join，取值：1：置忙 2：通话中 4：振铃 8：无效 16：整理
	//
	// example:
	//
	// 0
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
	// 33
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
	// 1
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
	// rrmemory
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
	// 1
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// 整理时间，取值范围3-3600秒
	//
	// example:
	//
	// 10
	WrapupTime *int64 `json:"WrapupTime,omitempty" xml:"WrapupTime,omitempty"`
}

func (s CloudListQueueResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudListQueueResponseBodyDataList) GetAnnouncePosition() *int64 {
	return s.AnnouncePosition
}

func (s *CloudListQueueResponseBodyDataList) GetAnnouncePositionFrequency() *int64 {
	return s.AnnouncePositionFrequency
}

func (s *CloudListQueueResponseBodyDataList) GetAnnouncePositionParam() *int64 {
	return s.AnnouncePositionParam
}

func (s *CloudListQueueResponseBodyDataList) GetAnnouncePositionYouarenext() *int64 {
	return s.AnnouncePositionYouarenext
}

func (s *CloudListQueueResponseBodyDataList) GetAnnounceSound() *int64 {
	return s.AnnounceSound
}

func (s *CloudListQueueResponseBodyDataList) GetAnnounceSoundFile() *string {
	return s.AnnounceSoundFile
}

func (s *CloudListQueueResponseBodyDataList) GetAnnounceSoundFrequency() *int64 {
	return s.AnnounceSoundFrequency
}

func (s *CloudListQueueResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudListQueueResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *CloudListQueueResponseBodyDataList) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *CloudListQueueResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *CloudListQueueResponseBodyDataList) GetJoinEmpty() *int64 {
	return s.JoinEmpty
}

func (s *CloudListQueueResponseBodyDataList) GetMaxLen() *int64 {
	return s.MaxLen
}

func (s *CloudListQueueResponseBodyDataList) GetMemberTimeout() *int64 {
	return s.MemberTimeout
}

func (s *CloudListQueueResponseBodyDataList) GetMusicClass() *string {
	return s.MusicClass
}

func (s *CloudListQueueResponseBodyDataList) GetQno() *string {
	return s.Qno
}

func (s *CloudListQueueResponseBodyDataList) GetQueueTimeout() *int64 {
	return s.QueueTimeout
}

func (s *CloudListQueueResponseBodyDataList) GetRetry() *int64 {
	return s.Retry
}

func (s *CloudListQueueResponseBodyDataList) GetSayAgentno() *bool {
	return s.SayAgentno
}

func (s *CloudListQueueResponseBodyDataList) GetServiceLevel() *int64 {
	return s.ServiceLevel
}

func (s *CloudListQueueResponseBodyDataList) GetStrategy() *string {
	return s.Strategy
}

func (s *CloudListQueueResponseBodyDataList) GetVipSupport() *int64 {
	return s.VipSupport
}

func (s *CloudListQueueResponseBodyDataList) GetWeight() *int64 {
	return s.Weight
}

func (s *CloudListQueueResponseBodyDataList) GetWrapupTime() *int64 {
	return s.WrapupTime
}

func (s *CloudListQueueResponseBodyDataList) SetAnnouncePosition(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnouncePosition = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnouncePositionFrequency(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnouncePositionFrequency = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnouncePositionParam(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnouncePositionParam = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnouncePositionYouarenext(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnouncePositionYouarenext = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnounceSound(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnounceSound = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnounceSoundFile(v string) *CloudListQueueResponseBodyDataList {
	s.AnnounceSoundFile = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetAnnounceSoundFrequency(v int64) *CloudListQueueResponseBodyDataList {
	s.AnnounceSoundFrequency = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetCreateTime(v string) *CloudListQueueResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetDescription(v string) *CloudListQueueResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetEnterpriseId(v int64) *CloudListQueueResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetId(v int64) *CloudListQueueResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetJoinEmpty(v int64) *CloudListQueueResponseBodyDataList {
	s.JoinEmpty = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetMaxLen(v int64) *CloudListQueueResponseBodyDataList {
	s.MaxLen = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetMemberTimeout(v int64) *CloudListQueueResponseBodyDataList {
	s.MemberTimeout = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetMusicClass(v string) *CloudListQueueResponseBodyDataList {
	s.MusicClass = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetQno(v string) *CloudListQueueResponseBodyDataList {
	s.Qno = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetQueueTimeout(v int64) *CloudListQueueResponseBodyDataList {
	s.QueueTimeout = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetRetry(v int64) *CloudListQueueResponseBodyDataList {
	s.Retry = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetSayAgentno(v bool) *CloudListQueueResponseBodyDataList {
	s.SayAgentno = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetServiceLevel(v int64) *CloudListQueueResponseBodyDataList {
	s.ServiceLevel = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetStrategy(v string) *CloudListQueueResponseBodyDataList {
	s.Strategy = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetVipSupport(v int64) *CloudListQueueResponseBodyDataList {
	s.VipSupport = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetWeight(v int64) *CloudListQueueResponseBodyDataList {
	s.Weight = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) SetWrapupTime(v int64) *CloudListQueueResponseBodyDataList {
	s.WrapupTime = &v
	return s
}

func (s *CloudListQueueResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
