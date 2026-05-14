// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkStatIbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkStatIbResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkStatIbResponseBody
	GetCode() *string
	SetData(v *ClinkStatIbResponseBodyData) *ClinkStatIbResponseBody
	GetData() *ClinkStatIbResponseBodyData
	SetMessage(v string) *ClinkStatIbResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkStatIbResponseBody
	GetRequestId() *string
}

type ClinkStatIbResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkStatIbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkStatIbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkStatIbResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkStatIbResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkStatIbResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkStatIbResponseBody) GetData() *ClinkStatIbResponseBodyData {
	return s.Data
}

func (s *ClinkStatIbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkStatIbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkStatIbResponseBody) SetAccessDeniedDetail(v string) *ClinkStatIbResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkStatIbResponseBody) SetCode(v string) *ClinkStatIbResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkStatIbResponseBody) SetData(v *ClinkStatIbResponseBodyData) *ClinkStatIbResponseBody {
	s.Data = v
	return s
}

func (s *ClinkStatIbResponseBody) SetMessage(v string) *ClinkStatIbResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkStatIbResponseBody) SetRequestId(v string) *ClinkStatIbResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkStatIbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkStatIbResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 报表数据（注：时长字段单位均为秒）
	StatIb []*ClinkStatIbResponseBodyDataStatIb `json:"StatIb,omitempty" xml:"StatIb,omitempty" type:"Repeated"`
}

func (s ClinkStatIbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkStatIbResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkStatIbResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkStatIbResponseBodyData) GetStatIb() []*ClinkStatIbResponseBodyDataStatIb {
	return s.StatIb
}

func (s *ClinkStatIbResponseBodyData) SetClinkRequestId(v string) *ClinkStatIbResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkStatIbResponseBodyData) SetStatIb(v []*ClinkStatIbResponseBodyDataStatIb) *ClinkStatIbResponseBodyData {
	s.StatIb = v
	return s
}

func (s *ClinkStatIbResponseBodyData) Validate() error {
	if s.StatIb != nil {
		for _, item := range s.StatIb {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkStatIbResponseBodyDataStatIb struct {
	// 接听电话座席数
	//
	// example:
	//
	// 96
	AnsweredAgentCount *int64 `json:"AnsweredAgentCount,omitempty" xml:"AnsweredAgentCount,omitempty"`
	// 人均接听量
	//
	// example:
	//
	// 6
	AvgAnsweredAgentCount *int64 `json:"AvgAnsweredAgentCount,omitempty" xml:"AvgAnsweredAgentCount,omitempty"`
	// 日期（统计方式2为同步日期；统计方式3为分时信息，例 0-1时）
	//
	// example:
	//
	// 20240203
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// 热线号码
	//
	// example:
	//
	// Hotline
	Hotline *string `json:"Hotline,omitempty" xml:"Hotline,omitempty"`
	// 热线别名
	//
	// example:
	//
	// HotlineName
	HotlineName *string `json:"HotlineName,omitempty" xml:"HotlineName,omitempty"`
	// 呼叫座席来电总数
	//
	// example:
	//
	// 5
	IbAgentCount *int64 `json:"IbAgentCount,omitempty" xml:"IbAgentCount,omitempty"`
	// 人工接听数
	//
	// example:
	//
	// 56
	IbAnsweredCount *int64 `json:"IbAnsweredCount,omitempty" xml:"IbAnsweredCount,omitempty"`
	// 人工接听排队总时长
	//
	// example:
	//
	// 90
	IbAnsweredQueueDuration *int64 `json:"IbAnsweredQueueDuration,omitempty" xml:"IbAnsweredQueueDuration,omitempty"`
	// 呼入人工接听率
	//
	// example:
	//
	// 0.28
	IbAnsweredRate *float64 `json:"IbAnsweredRate,omitempty" xml:"IbAnsweredRate,omitempty"`
	// 人工接听平均排队时长
	//
	// example:
	//
	// 93
	IbAvgAnsweredQueueDuration *int64 `json:"IbAvgAnsweredQueueDuration,omitempty" xml:"IbAvgAnsweredQueueDuration,omitempty"`
	// 呼入平均通话时长
	//
	// example:
	//
	// 19
	IbAvgBridgeDuration *int64 `json:"IbAvgBridgeDuration,omitempty" xml:"IbAvgBridgeDuration,omitempty"`
	// 呼入平均时长
	//
	// example:
	//
	// 98
	IbAvgDuration *int64 `json:"IbAvgDuration,omitempty" xml:"IbAvgDuration,omitempty"`
	// 进入队列平均排队时长
	//
	// example:
	//
	// 40
	IbAvgQueueDuration *int64 `json:"IbAvgQueueDuration,omitempty" xml:"IbAvgQueueDuration,omitempty"`
	// 机器人应答数
	//
	// example:
	//
	// 36
	IbBotAnsweredCount *int64 `json:"IbBotAnsweredCount,omitempty" xml:"IbBotAnsweredCount,omitempty"`
	// 机器人未应答数
	//
	// example:
	//
	// 79
	IbBotUnansweredCount *int64 `json:"IbBotUnansweredCount,omitempty" xml:"IbBotUnansweredCount,omitempty"`
	// 呼入通话总时长（客户侧）
	//
	// example:
	//
	// 28
	IbBridgeDuration *int64 `json:"IbBridgeDuration,omitempty" xml:"IbBridgeDuration,omitempty"`
	// 直转座席数
	//
	// example:
	//
	// 61
	IbDirectAgentCount *int64 `json:"IbDirectAgentCount,omitempty" xml:"IbDirectAgentCount,omitempty"`
	// 直转队列数
	//
	// example:
	//
	// 18
	IbDirectQueueCount *int64 `json:"IbDirectQueueCount,omitempty" xml:"IbDirectQueueCount,omitempty"`
	// 直转电话/IP话机数
	//
	// example:
	//
	// 81
	IbDirectTelCount *int64 `json:"IbDirectTelCount,omitempty" xml:"IbDirectTelCount,omitempty"`
	// 进入IVR人工接听数
	//
	// example:
	//
	// 17
	IbIvrAnsweredCount *int64 `json:"IbIvrAnsweredCount,omitempty" xml:"IbIvrAnsweredCount,omitempty"`
	// 进入IVR数
	//
	// example:
	//
	// 45
	IbIvrCount *int64 `json:"IbIvrCount,omitempty" xml:"IbIvrCount,omitempty"`
	// 进入IVR未转队列数
	//
	// example:
	//
	// 73
	IbIvrNoQueueCount *int64 `json:"IbIvrNoQueueCount,omitempty" xml:"IbIvrNoQueueCount,omitempty"`
	// 进入IVR转队列数
	//
	// example:
	//
	// 94
	IbIvrQueueCount *int64 `json:"IbIvrQueueCount,omitempty" xml:"IbIvrQueueCount,omitempty"`
	// 进入IVR客户速挂
	//
	// example:
	//
	// 18
	IbIvrQuickUnlinkCount *int64 `json:"IbIvrQuickUnlinkCount,omitempty" xml:"IbIvrQuickUnlinkCount,omitempty"`
	// 进入IVR系统应答数
	//
	// example:
	//
	// 12
	IbIvrSystemAnsweredCount *int64 `json:"IbIvrSystemAnsweredCount,omitempty" xml:"IbIvrSystemAnsweredCount,omitempty"`
	// 进入IVR人工未接听数
	//
	// example:
	//
	// 70
	IbIvrUnAnsweredCount *int64 `json:"IbIvrUnAnsweredCount,omitempty" xml:"IbIvrUnAnsweredCount,omitempty"`
	// 未进入IVR数
	//
	// example:
	//
	// 12
	IbNoIvrCount *int64 `json:"IbNoIvrCount,omitempty" xml:"IbNoIvrCount,omitempty"`
	// 进入队列来电通数
	//
	// example:
	//
	// 91
	IbQueueCount *int64 `json:"IbQueueCount,omitempty" xml:"IbQueueCount,omitempty"`
	// 进入队列排队总时长
	//
	// example:
	//
	// 79
	IbQueueDuration *int64 `json:"IbQueueDuration,omitempty" xml:"IbQueueDuration,omitempty"`
	// 黑名单来电数
	//
	// example:
	//
	// 90
	IbRestrictCount *int64 `json:"IbRestrictCount,omitempty" xml:"IbRestrictCount,omitempty"`
	// X秒接听数
	//
	// example:
	//
	// 1
	IbRingingRangeCount *int64 `json:"IbRingingRangeCount,omitempty" xml:"IbRingingRangeCount,omitempty"`
	// X秒接听率
	//
	// example:
	//
	// 0.43
	IbRingingRangeRate *float64 `json:"IbRingingRangeRate,omitempty" xml:"IbRingingRangeRate,omitempty"`
	// 系统应答数
	//
	// example:
	//
	// 25
	IbSystemAnsweredCount *int64 `json:"IbSystemAnsweredCount,omitempty" xml:"IbSystemAnsweredCount,omitempty"`
	// 系统未应答数
	//
	// example:
	//
	// 2
	IbSystemUnAnsweredCount *int64 `json:"IbSystemUnAnsweredCount,omitempty" xml:"IbSystemUnAnsweredCount,omitempty"`
	// 来电总通数
	//
	// example:
	//
	// 40
	IbTotalCount *int64 `json:"IbTotalCount,omitempty" xml:"IbTotalCount,omitempty"`
	// 人工未接听数
	//
	// example:
	//
	// 70
	IbUnansweredCount *int64 `json:"IbUnansweredCount,omitempty" xml:"IbUnansweredCount,omitempty"`
	// VIP呼入数
	//
	// example:
	//
	// 77
	IbVipCount *int64 `json:"IbVipCount,omitempty" xml:"IbVipCount,omitempty"`
	// 网上400呼入数
	//
	// example:
	//
	// 53
	IbWebCount *int64 `json:"IbWebCount,omitempty" xml:"IbWebCount,omitempty"`
	// 呼入总时长
	//
	// example:
	//
	// 69
	IbtotalDuration *int64 `json:"IbtotalDuration,omitempty" xml:"IbtotalDuration,omitempty"`
	// 进入IVR直转电话/IP话机人工接听数
	//
	// example:
	//
	// 95
	IvrDirectTelAnsweredCount *int64 `json:"IvrDirectTelAnsweredCount,omitempty" xml:"IvrDirectTelAnsweredCount,omitempty"`
	// 进入IVR直转电话/IP话机数
	//
	// example:
	//
	// 88
	IvrDirectTelCount *int64 `json:"IvrDirectTelCount,omitempty" xml:"IvrDirectTelCount,omitempty"`
	// 进入IVR直转电话/IP话机系统应答数
	//
	// example:
	//
	// 10
	IvrDirectTelSysAnsweredCount *int64 `json:"IvrDirectTelSysAnsweredCount,omitempty" xml:"IvrDirectTelSysAnsweredCount,omitempty"`
	// 进入IVR直转电话/IP话机人工未接听数
	//
	// example:
	//
	// 75
	IvrDirectTelUnansweredCount *int64 `json:"IvrDirectTelUnansweredCount,omitempty" xml:"IvrDirectTelUnansweredCount,omitempty"`
	// 客户速挂数
	//
	// example:
	//
	// 67
	QuickUnlinkCount *int64 `json:"QuickUnlinkCount,omitempty" xml:"QuickUnlinkCount,omitempty"`
	// 热线重复来电数
	//
	// example:
	//
	// 14
	RepeatHotlineCount *int64 `json:"RepeatHotlineCount,omitempty" xml:"RepeatHotlineCount,omitempty"`
}

func (s ClinkStatIbResponseBodyDataStatIb) String() string {
	return dara.Prettify(s)
}

func (s ClinkStatIbResponseBodyDataStatIb) GoString() string {
	return s.String()
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetAnsweredAgentCount() *int64 {
	return s.AnsweredAgentCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetAvgAnsweredAgentCount() *int64 {
	return s.AvgAnsweredAgentCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetDate() *string {
	return s.Date
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetHotline() *string {
	return s.Hotline
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetHotlineName() *string {
	return s.HotlineName
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAgentCount() *int64 {
	return s.IbAgentCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAnsweredCount() *int64 {
	return s.IbAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAnsweredQueueDuration() *int64 {
	return s.IbAnsweredQueueDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAnsweredRate() *float64 {
	return s.IbAnsweredRate
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAvgAnsweredQueueDuration() *int64 {
	return s.IbAvgAnsweredQueueDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAvgBridgeDuration() *int64 {
	return s.IbAvgBridgeDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAvgDuration() *int64 {
	return s.IbAvgDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbAvgQueueDuration() *int64 {
	return s.IbAvgQueueDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbBotAnsweredCount() *int64 {
	return s.IbBotAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbBotUnansweredCount() *int64 {
	return s.IbBotUnansweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbBridgeDuration() *int64 {
	return s.IbBridgeDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbDirectAgentCount() *int64 {
	return s.IbDirectAgentCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbDirectQueueCount() *int64 {
	return s.IbDirectQueueCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbDirectTelCount() *int64 {
	return s.IbDirectTelCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrAnsweredCount() *int64 {
	return s.IbIvrAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrCount() *int64 {
	return s.IbIvrCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrNoQueueCount() *int64 {
	return s.IbIvrNoQueueCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrQueueCount() *int64 {
	return s.IbIvrQueueCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrQuickUnlinkCount() *int64 {
	return s.IbIvrQuickUnlinkCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrSystemAnsweredCount() *int64 {
	return s.IbIvrSystemAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbIvrUnAnsweredCount() *int64 {
	return s.IbIvrUnAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbNoIvrCount() *int64 {
	return s.IbNoIvrCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbQueueCount() *int64 {
	return s.IbQueueCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbQueueDuration() *int64 {
	return s.IbQueueDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbRestrictCount() *int64 {
	return s.IbRestrictCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbRingingRangeCount() *int64 {
	return s.IbRingingRangeCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbRingingRangeRate() *float64 {
	return s.IbRingingRangeRate
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbSystemAnsweredCount() *int64 {
	return s.IbSystemAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbSystemUnAnsweredCount() *int64 {
	return s.IbSystemUnAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbTotalCount() *int64 {
	return s.IbTotalCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbUnansweredCount() *int64 {
	return s.IbUnansweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbVipCount() *int64 {
	return s.IbVipCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbWebCount() *int64 {
	return s.IbWebCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIbtotalDuration() *int64 {
	return s.IbtotalDuration
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIvrDirectTelAnsweredCount() *int64 {
	return s.IvrDirectTelAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIvrDirectTelCount() *int64 {
	return s.IvrDirectTelCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIvrDirectTelSysAnsweredCount() *int64 {
	return s.IvrDirectTelSysAnsweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetIvrDirectTelUnansweredCount() *int64 {
	return s.IvrDirectTelUnansweredCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetQuickUnlinkCount() *int64 {
	return s.QuickUnlinkCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) GetRepeatHotlineCount() *int64 {
	return s.RepeatHotlineCount
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetAnsweredAgentCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.AnsweredAgentCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetAvgAnsweredAgentCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.AvgAnsweredAgentCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetDate(v string) *ClinkStatIbResponseBodyDataStatIb {
	s.Date = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetHotline(v string) *ClinkStatIbResponseBodyDataStatIb {
	s.Hotline = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetHotlineName(v string) *ClinkStatIbResponseBodyDataStatIb {
	s.HotlineName = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAgentCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAgentCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAnsweredQueueDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAnsweredQueueDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAnsweredRate(v float64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAnsweredRate = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAvgAnsweredQueueDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAvgAnsweredQueueDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAvgBridgeDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAvgBridgeDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAvgDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAvgDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbAvgQueueDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbAvgQueueDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbBotAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbBotAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbBotUnansweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbBotUnansweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbBridgeDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbBridgeDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbDirectAgentCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbDirectAgentCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbDirectQueueCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbDirectQueueCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbDirectTelCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbDirectTelCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrNoQueueCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrNoQueueCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrQueueCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrQueueCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrQuickUnlinkCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrQuickUnlinkCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrSystemAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrSystemAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbIvrUnAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbIvrUnAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbNoIvrCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbNoIvrCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbQueueCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbQueueCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbQueueDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbQueueDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbRestrictCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbRestrictCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbRingingRangeCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbRingingRangeCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbRingingRangeRate(v float64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbRingingRangeRate = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbSystemAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbSystemAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbSystemUnAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbSystemUnAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbTotalCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbTotalCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbUnansweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbUnansweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbVipCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbVipCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbWebCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbWebCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIbtotalDuration(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IbtotalDuration = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIvrDirectTelAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IvrDirectTelAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIvrDirectTelCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IvrDirectTelCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIvrDirectTelSysAnsweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IvrDirectTelSysAnsweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetIvrDirectTelUnansweredCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.IvrDirectTelUnansweredCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetQuickUnlinkCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.QuickUnlinkCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) SetRepeatHotlineCount(v int64) *ClinkStatIbResponseBodyDataStatIb {
	s.RepeatHotlineCount = &v
	return s
}

func (s *ClinkStatIbResponseBodyDataStatIb) Validate() error {
	return dara.Validate(s)
}
