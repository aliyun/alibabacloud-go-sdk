// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentWorkloadReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentWorkloadReportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentWorkloadReportResponseBody
	GetCode() *string
	SetData(v *CloudAgentWorkloadReportResponseBodyData) *CloudAgentWorkloadReportResponseBody
	GetData() *CloudAgentWorkloadReportResponseBodyData
	SetMessage(v string) *CloudAgentWorkloadReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentWorkloadReportResponseBody
	GetRequestId() *string
}

type CloudAgentWorkloadReportResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentWorkloadReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 756BCB70-492E-5877-ABEB-03BB5CA28540
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentWorkloadReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentWorkloadReportResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentWorkloadReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentWorkloadReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentWorkloadReportResponseBody) GetData() *CloudAgentWorkloadReportResponseBodyData {
	return s.Data
}

func (s *CloudAgentWorkloadReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentWorkloadReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentWorkloadReportResponseBody) SetAccessDeniedDetail(v string) *CloudAgentWorkloadReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBody) SetCode(v string) *CloudAgentWorkloadReportResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBody) SetData(v *CloudAgentWorkloadReportResponseBodyData) *CloudAgentWorkloadReportResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBody) SetMessage(v string) *CloudAgentWorkloadReportResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBody) SetRequestId(v string) *CloudAgentWorkloadReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentWorkloadReportResponseBodyData struct {
	List []*CloudAgentWorkloadReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPageCount *string `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s CloudAgentWorkloadReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentWorkloadReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentWorkloadReportResponseBodyData) GetList() []*CloudAgentWorkloadReportResponseBodyDataList {
	return s.List
}

func (s *CloudAgentWorkloadReportResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *CloudAgentWorkloadReportResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyData) GetTotalPageCount() *string {
	return s.TotalPageCount
}

func (s *CloudAgentWorkloadReportResponseBodyData) SetList(v []*CloudAgentWorkloadReportResponseBodyDataList) *CloudAgentWorkloadReportResponseBodyData {
	s.List = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyData) SetPageSize(v string) *CloudAgentWorkloadReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyData) SetTotalCount(v string) *CloudAgentWorkloadReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyData) SetTotalPageCount(v string) *CloudAgentWorkloadReportResponseBodyData {
	s.TotalPageCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyData) Validate() error {
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

type CloudAgentWorkloadReportResponseBodyDataList struct {
	// 座席来电接听率
	//
	// example:
	//
	// 80%
	AgentAnswerRate *string `json:"AgentAnswerRate,omitempty" xml:"AgentAnswerRate,omitempty"`
	// 座席创建时间
	//
	// example:
	//
	// 2019-04-20
	AgentCreateTime *string `json:"AgentCreateTime,omitempty" xml:"AgentCreateTime,omitempty"`
	// 座席名称
	//
	// example:
	//
	// test-3
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 座席呼叫接听率
	//
	// example:
	//
	// 0%
	AgentRate *string `json:"AgentRate,omitempty" xml:"AgentRate,omitempty"`
	// 座席停用时间
	//
	// example:
	//
	// 00:00:01
	AgentStopTime *string `json:"AgentStopTime,omitempty" xml:"AgentStopTime,omitempty"`
	// 平均空闲时长
	//
	// example:
	//
	// 00:00:01
	AvgIdleTime *string `json:"AvgIdleTime,omitempty" xml:"AvgIdleTime,omitempty"`
	// 平均置忙时长
	//
	// example:
	//
	// 00:00:00
	AvgPauseTime *string `json:"AvgPauseTime,omitempty" xml:"AvgPauseTime,omitempty"`
	// 平均等待时长
	//
	// example:
	//
	// 00:00:00
	AvgPreviewObWaitTime *string `json:"AvgPreviewObWaitTime,omitempty" xml:"AvgPreviewObWaitTime,omitempty"`
	// 平均休息时长
	//
	// example:
	//
	// 00:00:03
	AvgRestTime *string `json:"AvgRestTime,omitempty" xml:"AvgRestTime,omitempty"`
	// 通话利用率
	//
	// example:
	//
	// 66%
	CallUtilization *string `json:"CallUtilization,omitempty" xml:"CallUtilization,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2026-04-20 10:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 时间范围
	//
	// example:
	//
	// 示例值
	DateTimeRange *string `json:"DateTimeRange,omitempty" xml:"DateTimeRange,omitempty"`
	// 统计日期
	//
	// example:
	//
	// 2026-04-20
	Day *string `json:"Day,omitempty" xml:"Day,omitempty"`
	// 主叫外呼客户接听数
	//
	// example:
	//
	// 4
	DirectObAnsweredCount *string `json:"DirectObAnsweredCount,omitempty" xml:"DirectObAnsweredCount,omitempty"`
	// 主叫外呼接听量分布
	//
	// example:
	//
	// 0
	DirectObAnsweredDurationArray *string `json:"DirectObAnsweredDurationArray,omitempty" xml:"DirectObAnsweredDurationArray,omitempty"`
	// 主叫外呼接听量分布描述
	DirectObAnsweredDurationDscArray []*string `json:"DirectObAnsweredDurationDscArray,omitempty" xml:"DirectObAnsweredDurationDscArray,omitempty" type:"Repeated"`
	// 主叫外呼总时长
	//
	// example:
	//
	// 00:00:00
	DirectObBridgeDuration *string `json:"DirectObBridgeDuration,omitempty" xml:"DirectObBridgeDuration,omitempty"`
	// 主叫外呼双方接听分钟数
	//
	// example:
	//
	// 00:00:00
	DirectObBridgeTime *string `json:"DirectObBridgeTime,omitempty" xml:"DirectObBridgeTime,omitempty"`
	// 主叫外呼咨询总次数
	//
	// example:
	//
	// 0
	DirectObConsultCount *string `json:"DirectObConsultCount,omitempty" xml:"DirectObConsultCount,omitempty"`
	// 主叫外呼强插总次数
	//
	// example:
	//
	// 0
	DirectObMonitorBargeCount *string `json:"DirectObMonitorBargeCount,omitempty" xml:"DirectObMonitorBargeCount,omitempty"`
	// 主叫外呼总监控强拆次数
	//
	// example:
	//
	// 0
	DirectObMonitorDisconnectCount *string `json:"DirectObMonitorDisconnectCount,omitempty" xml:"DirectObMonitorDisconnectCount,omitempty"`
	// 主叫外呼监听总次数
	//
	// example:
	//
	// 0
	DirectObMonitorSpyCount *string `json:"DirectObMonitorSpyCount,omitempty" xml:"DirectObMonitorSpyCount,omitempty"`
	// 主叫外呼监控三方总次数
	//
	// example:
	//
	// 0
	DirectObMonitorThreewayCount *string `json:"DirectObMonitorThreewayCount,omitempty" xml:"DirectObMonitorThreewayCount,omitempty"`
	// 主叫外呼耳语总次数
	//
	// example:
	//
	// 0
	DirectObMonitorWhisperCount *string `json:"DirectObMonitorWhisperCount,omitempty" xml:"DirectObMonitorWhisperCount,omitempty"`
	// 主叫外呼总次数
	//
	// example:
	//
	// 0
	DirectObTotalCount *string `json:"DirectObTotalCount,omitempty" xml:"DirectObTotalCount,omitempty"`
	// 主叫外呼总分钟数
	//
	// example:
	//
	// 00:00:22
	DirectObTotalTime *string `json:"DirectObTotalTime,omitempty" xml:"DirectObTotalTime,omitempty"`
	// 主叫外呼转移总次数
	//
	// example:
	//
	// 0
	DirectObTransferCount *string `json:"DirectObTransferCount,omitempty" xml:"DirectObTransferCount,omitempty"`
	// 主叫外呼有效接听数
	//
	// example:
	//
	// 0
	DirectObValidCalls *string `json:"DirectObValidCalls,omitempty" xml:"DirectObValidCalls,omitempty"`
	// 主叫外呼有效接听时长
	//
	// example:
	//
	// 00:00:00
	DirectObValidTotalBridgeTime *string `json:"DirectObValidTotalBridgeTime,omitempty" xml:"DirectObValidTotalBridgeTime,omitempty"`
	// 企业id
	//
	// example:
	//
	// 7000002
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 工时利用率
	//
	// example:
	//
	// 80%
	HoursUtilization *string `json:"HoursUtilization,omitempty" xml:"HoursUtilization,omitempty"`
	// 座席挂机数
	//
	// example:
	//
	// 0
	IbAgentHangupCount *string `json:"IbAgentHangupCount,omitempty" xml:"IbAgentHangupCount,omitempty"`
	// 来电座席拒接数
	//
	// example:
	//
	// 0
	IbAgentRefuseCount *string `json:"IbAgentRefuseCount,omitempty" xml:"IbAgentRefuseCount,omitempty"`
	// 呼入座席接听数
	//
	// example:
	//
	// 0
	IbAnsweredCount *string `json:"IbAnsweredCount,omitempty" xml:"IbAnsweredCount,omitempty"`
	// 接听量分布
	//
	// example:
	//
	// ""
	IbAnsweredDurationArray []*string `json:"IbAnsweredDurationArray,omitempty" xml:"IbAnsweredDurationArray,omitempty" type:"Repeated"`
	// 呼入接听量分布描述
	IbAnsweredDurationDscArray []*string `json:"IbAnsweredDurationDscArray,omitempty" xml:"IbAnsweredDurationDscArray,omitempty" type:"Repeated"`
	// 呼入接听平均延迟时间
	//
	// example:
	//
	// 00:00:00
	IbAvgAnswerDelayTime *string `json:"IbAvgAnswerDelayTime,omitempty" xml:"IbAvgAnswerDelayTime,omitempty"`
	// 呼入平均通话时长
	//
	// example:
	//
	// 00:00:00
	IbAvgBridgeTime *string `json:"IbAvgBridgeTime,omitempty" xml:"IbAvgBridgeTime,omitempty"`
	// 呼入平均呼叫时长
	//
	// example:
	//
	// 00:00:00
	IbAvgCallingTime *string `json:"IbAvgCallingTime,omitempty" xml:"IbAvgCallingTime,omitempty"`
	// 呼入平均保持时长
	//
	// example:
	//
	// 00:00:00
	IbAvgHoldTime *string `json:"IbAvgHoldTime,omitempty" xml:"IbAvgHoldTime,omitempty"`
	// 呼入平均整理时长
	//
	// example:
	//
	// 00:00:10
	IbAvgWrapupTime *string `json:"IbAvgWrapupTime,omitempty" xml:"IbAvgWrapupTime,omitempty"`
	// 呼入通话次数
	//
	// example:
	//
	// 2
	IbBridgeCount *string `json:"IbBridgeCount,omitempty" xml:"IbBridgeCount,omitempty"`
	// 呼入通话分钟数
	//
	// example:
	//
	// 00:00:00
	IbBridgeMinute *string `json:"IbBridgeMinute,omitempty" xml:"IbBridgeMinute,omitempty"`
	// 呼入双方接听时长
	//
	// example:
	//
	// 00:00:00
	IbBridgeTime *string `json:"IbBridgeTime,omitempty" xml:"IbBridgeTime,omitempty"`
	// 呼入呼叫次数
	//
	// example:
	//
	// 0
	IbCallingCount *string `json:"IbCallingCount,omitempty" xml:"IbCallingCount,omitempty"`
	// 呼入总呼叫时长
	//
	// example:
	//
	// 00:00:00
	IbCallingTime *string `json:"IbCallingTime,omitempty" xml:"IbCallingTime,omitempty"`
	// 呼入咨询数
	//
	// example:
	//
	// 0
	IbConsultCount *string `json:"IbConsultCount,omitempty" xml:"IbConsultCount,omitempty"`
	// 客户挂机数
	//
	// example:
	//
	// 3
	IbCustomerHangupCount *string `json:"IbCustomerHangupCount,omitempty" xml:"IbCustomerHangupCount,omitempty"`
	// 呼叫保持次数
	//
	// example:
	//
	// 0
	IbHoldCount *string `json:"IbHoldCount,omitempty" xml:"IbHoldCount,omitempty"`
	// 呼入总保持时长
	//
	// example:
	//
	// 00:00:00
	IbHoldTime *string `json:"IbHoldTime,omitempty" xml:"IbHoldTime,omitempty"`
	// 呼入接听最大延迟时间
	//
	// example:
	//
	// 00:00:00
	IbMaxAnswerDelayTime *string `json:"IbMaxAnswerDelayTime,omitempty" xml:"IbMaxAnswerDelayTime,omitempty"`
	// 呼入最大通话时长
	//
	// example:
	//
	// 00:00:00
	IbMaxBridgeTime *string `json:"IbMaxBridgeTime,omitempty" xml:"IbMaxBridgeTime,omitempty"`
	// 呼入最大呼叫时长
	//
	// example:
	//
	// 00:00:10
	IbMaxCallingTime *string `json:"IbMaxCallingTime,omitempty" xml:"IbMaxCallingTime,omitempty"`
	// 呼入最大保持时长
	//
	// example:
	//
	// 00:00:00
	IbMaxHoldTime *string `json:"IbMaxHoldTime,omitempty" xml:"IbMaxHoldTime,omitempty"`
	// 呼入最大整理时长
	//
	// example:
	//
	// 00:00:00
	IbMaxWrapupTime *string `json:"IbMaxWrapupTime,omitempty" xml:"IbMaxWrapupTime,omitempty"`
	// 呼入接听最小延迟时间
	//
	// example:
	//
	// 00:00:00
	IbMinAnswerDelayTime *string `json:"IbMinAnswerDelayTime,omitempty" xml:"IbMinAnswerDelayTime,omitempty"`
	// 呼入最小通话时长
	//
	// example:
	//
	// 00:00:10
	IbMinBridgeTime *string `json:"IbMinBridgeTime,omitempty" xml:"IbMinBridgeTime,omitempty"`
	// 呼入最小呼叫时长
	//
	// example:
	//
	// 00:00:00
	IbMinCallingTime *string `json:"IbMinCallingTime,omitempty" xml:"IbMinCallingTime,omitempty"`
	// 呼入最小保持时长
	//
	// example:
	//
	// 00:00:00
	IbMinHoldTime *string `json:"IbMinHoldTime,omitempty" xml:"IbMinHoldTime,omitempty"`
	// 呼入最小整理时长
	//
	// example:
	//
	// 00:00:00
	IbMinWrapupTime *string `json:"IbMinWrapupTime,omitempty" xml:"IbMinWrapupTime,omitempty"`
	// 呼入监控强插数
	//
	// example:
	//
	// 0
	IbMonitorBargeCount *string `json:"IbMonitorBargeCount,omitempty" xml:"IbMonitorBargeCount,omitempty"`
	// 呼入强插时长
	//
	// example:
	//
	// 00:00:00
	IbMonitorBargeDuration *string `json:"IbMonitorBargeDuration,omitempty" xml:"IbMonitorBargeDuration,omitempty"`
	// 呼入监控强拆数
	//
	// example:
	//
	// 0
	IbMonitorDisconnectCount *string `json:"IbMonitorDisconnectCount,omitempty" xml:"IbMonitorDisconnectCount,omitempty"`
	// 呼入监控抢线数
	//
	// example:
	//
	// 0
	IbMonitorPickupCount *string `json:"IbMonitorPickupCount,omitempty" xml:"IbMonitorPickupCount,omitempty"`
	// 呼入抢线时长
	//
	// example:
	//
	// 0
	IbMonitorPickupDuration *string `json:"IbMonitorPickupDuration,omitempty" xml:"IbMonitorPickupDuration,omitempty"`
	// 呼入监控监听数
	//
	// example:
	//
	// 3
	IbMonitorSpyCount *string `json:"IbMonitorSpyCount,omitempty" xml:"IbMonitorSpyCount,omitempty"`
	// 呼入监听时长
	//
	// example:
	//
	// 00:00:00
	IbMonitorSpyDuration *string `json:"IbMonitorSpyDuration,omitempty" xml:"IbMonitorSpyDuration,omitempty"`
	// 呼入监控三方数
	//
	// example:
	//
	// 0
	IbMonitorThreewayCount *string `json:"IbMonitorThreewayCount,omitempty" xml:"IbMonitorThreewayCount,omitempty"`
	// 呼入监控耳语数
	//
	// example:
	//
	// 0
	IbMonitorWhisperCount *string `json:"IbMonitorWhisperCount,omitempty" xml:"IbMonitorWhisperCount,omitempty"`
	// 呼入耳语时长
	//
	// example:
	//
	// 00:00:00
	IbMonitorWhisperDuration *string `json:"IbMonitorWhisperDuration,omitempty" xml:"IbMonitorWhisperDuration,omitempty"`
	// 呼入接听总延迟时间
	//
	// example:
	//
	// 00:00:00
	IbTotalAnswerDelayTime *string `json:"IbTotalAnswerDelayTime,omitempty" xml:"IbTotalAnswerDelayTime,omitempty"`
	// 总呼入座席数
	//
	// example:
	//
	// 6
	IbTotalCount *string `json:"IbTotalCount,omitempty" xml:"IbTotalCount,omitempty"`
	// 呼入总分钟数
	//
	// example:
	//
	// 0
	IbTotalMinute *string `json:"IbTotalMinute,omitempty" xml:"IbTotalMinute,omitempty"`
	// 呼入总通话时长
	//
	// example:
	//
	// 3
	IbTotalTime *string `json:"IbTotalTime,omitempty" xml:"IbTotalTime,omitempty"`
	// 呼入转移数
	//
	// example:
	//
	// 2
	IbTransferCount *string `json:"IbTransferCount,omitempty" xml:"IbTransferCount,omitempty"`
	// 呼入座席未接听数
	//
	// example:
	//
	// 2
	IbUnansweredCount *string `json:"IbUnansweredCount,omitempty" xml:"IbUnansweredCount,omitempty"`
	// 座席来电接听数
	//
	// example:
	//
	// 3
	IbUniqueAnsweredCount *string `json:"IbUniqueAnsweredCount,omitempty" xml:"IbUniqueAnsweredCount,omitempty"`
	// 座席来电数
	//
	// example:
	//
	// 0
	IbUniqueTotalCount *string `json:"IbUniqueTotalCount,omitempty" xml:"IbUniqueTotalCount,omitempty"`
	// 座席来电未接听数
	//
	// example:
	//
	// 0
	IbUniqueUnansweredCount *string `json:"IbUniqueUnansweredCount,omitempty" xml:"IbUniqueUnansweredCount,omitempty"`
	// 呼入有效通话次数
	//
	// example:
	//
	// 0
	IbValidCalls *string `json:"IbValidCalls,omitempty" xml:"IbValidCalls,omitempty"`
	// 呼入有效通话时长
	//
	// example:
	//
	// 00:00:10
	IbValidTotalBridgeTime *string `json:"IbValidTotalBridgeTime,omitempty" xml:"IbValidTotalBridgeTime,omitempty"`
	// 呼入整理次数
	//
	// example:
	//
	// 0
	IbWrapupCount *string `json:"IbWrapupCount,omitempty" xml:"IbWrapupCount,omitempty"`
	// 呼入总整理时长
	//
	// example:
	//
	// 00:00:00
	IbWrapupTime *string `json:"IbWrapupTime,omitempty" xml:"IbWrapupTime,omitempty"`
	// id
	//
	// example:
	//
	// 123124
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 空闲次数
	//
	// example:
	//
	// 0
	IdleCount *string `json:"IdleCount,omitempty" xml:"IdleCount,omitempty"`
	// 总空闲时长
	//
	// example:
	//
	// 00:00:00
	IdleTime *string `json:"IdleTime,omitempty" xml:"IdleTime,omitempty"`
	// 预览外呼IVR转移座席接听数
	//
	// example:
	//
	// 0
	IvrTransferBridgeCount *string `json:"IvrTransferBridgeCount,omitempty" xml:"IvrTransferBridgeCount,omitempty"`
	// 预览外呼IVR转移双方接听时长
	//
	// example:
	//
	// 00:00:01
	IvrTransferBridgeDuration *string `json:"IvrTransferBridgeDuration,omitempty" xml:"IvrTransferBridgeDuration,omitempty"`
	// 预览外呼IVR转移座席数
	//
	// example:
	//
	// 0
	IvrTransferCount *string `json:"IvrTransferCount,omitempty" xml:"IvrTransferCount,omitempty"`
	// 总登录时长
	//
	// example:
	//
	// 00:00:01
	LoginTime *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	// 最大空闲时长
	//
	// example:
	//
	// 00:00:00
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// 最大置忙时长
	//
	// example:
	//
	// 00:00:00
	MaxPauseTime *string `json:"MaxPauseTime,omitempty" xml:"MaxPauseTime,omitempty"`
	// 最大休息时长
	//
	// example:
	//
	// 00:00:00
	MaxRestTime *string `json:"MaxRestTime,omitempty" xml:"MaxRestTime,omitempty"`
	// 最小空闲时长
	//
	// example:
	//
	// 00:00:01
	MinIdleTime *string `json:"MinIdleTime,omitempty" xml:"MinIdleTime,omitempty"`
	// 最小置忙时长
	//
	// example:
	//
	// 00:00:00
	MinPauseTime *string `json:"MinPauseTime,omitempty" xml:"MinPauseTime,omitempty"`
	// 最小休息时长
	//
	// example:
	//
	// 00:00:00
	MinRestTime *string `json:"MinRestTime,omitempty" xml:"MinRestTime,omitempty"`
	// 外呼平均通话时长
	//
	// example:
	//
	// 00:00:00
	ObAvgBridgeTime *string `json:"ObAvgBridgeTime,omitempty" xml:"ObAvgBridgeTime,omitempty"`
	// 外呼平均呼叫时长
	//
	// example:
	//
	// 00:00:11
	ObAvgCallingTime *string `json:"ObAvgCallingTime,omitempty" xml:"ObAvgCallingTime,omitempty"`
	// 外呼平均保持时长
	//
	// example:
	//
	// 00:00:50
	ObAvgHoldTime *string `json:"ObAvgHoldTime,omitempty" xml:"ObAvgHoldTime,omitempty"`
	// 外呼平均整理时长
	//
	// example:
	//
	// 00:00:05
	ObAvgWrapupTime *string `json:"ObAvgWrapupTime,omitempty" xml:"ObAvgWrapupTime,omitempty"`
	// 外呼通话次数
	//
	// example:
	//
	// 3
	ObBridgeCount *string `json:"ObBridgeCount,omitempty" xml:"ObBridgeCount,omitempty"`
	// 外呼总通话时长
	//
	// example:
	//
	// ObBridgeTime
	ObBridgeTime *string `json:"ObBridgeTime,omitempty" xml:"ObBridgeTime,omitempty"`
	// 外呼呼叫次数
	//
	// example:
	//
	// 1
	ObCallingCount *string `json:"ObCallingCount,omitempty" xml:"ObCallingCount,omitempty"`
	// 外呼总呼叫时长
	//
	// example:
	//
	// 00:00:00
	ObCallingTime *string `json:"ObCallingTime,omitempty" xml:"ObCallingTime,omitempty"`
	// 外呼保持次数
	//
	// example:
	//
	// 1
	ObHoldCount *string `json:"ObHoldCount,omitempty" xml:"ObHoldCount,omitempty"`
	// 外呼总保持时长
	//
	// example:
	//
	// 00:00:10
	ObHoldTime *string `json:"ObHoldTime,omitempty" xml:"ObHoldTime,omitempty"`
	// 外呼最大通话时长
	//
	// example:
	//
	// 00:00:00
	ObMaxBridgeTime *string `json:"ObMaxBridgeTime,omitempty" xml:"ObMaxBridgeTime,omitempty"`
	// 外呼最大呼叫时长
	//
	// example:
	//
	// 00:00:03
	ObMaxCallingTime *string `json:"ObMaxCallingTime,omitempty" xml:"ObMaxCallingTime,omitempty"`
	// 外呼最大保持时长
	//
	// example:
	//
	// 00:00:01
	ObMaxHoldTime *string `json:"ObMaxHoldTime,omitempty" xml:"ObMaxHoldTime,omitempty"`
	// 外呼最大整理时长
	//
	// example:
	//
	// 00:00:22
	ObMaxWrapupTime *string `json:"ObMaxWrapupTime,omitempty" xml:"ObMaxWrapupTime,omitempty"`
	// 外呼最小通话时长
	//
	// example:
	//
	// 00:00:02
	ObMinBridgeTime *string `json:"ObMinBridgeTime,omitempty" xml:"ObMinBridgeTime,omitempty"`
	// 外呼最小呼叫时长
	//
	// example:
	//
	// 0
	ObMinCallingTime *string `json:"ObMinCallingTime,omitempty" xml:"ObMinCallingTime,omitempty"`
	// 外呼最小保持时长
	//
	// example:
	//
	// 00:00:00
	ObMinHoldTime *string `json:"ObMinHoldTime,omitempty" xml:"ObMinHoldTime,omitempty"`
	// 外呼最小整理时长
	//
	// example:
	//
	// 00:00:00
	ObMinWrapupTime *string `json:"ObMinWrapupTime,omitempty" xml:"ObMinWrapupTime,omitempty"`
	// 外呼双方接听时长
	//
	// example:
	//
	// 00:00:10
	ObRealBridgeTime *string `json:"ObRealBridgeTime,omitempty" xml:"ObRealBridgeTime,omitempty"`
	// 外呼整理次数
	//
	// example:
	//
	// 3
	ObWrapupCount *string `json:"ObWrapupCount,omitempty" xml:"ObWrapupCount,omitempty"`
	// 外呼总整理时长
	//
	// example:
	//
	// 00:00:00
	ObWrapupTime *string `json:"ObWrapupTime,omitempty" xml:"ObWrapupTime,omitempty"`
	// 置忙次数
	//
	// example:
	//
	// 1
	PauseCount *string `json:"PauseCount,omitempty" xml:"PauseCount,omitempty"`
	// 总置忙时长
	//
	// example:
	//
	// 00:00:10
	PauseTime *string `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// 预测外呼座席接听通话总小时
	//
	// example:
	//
	// 1
	PredictObAgentAnsweredHour *string `json:"PredictObAgentAnsweredHour,omitempty" xml:"PredictObAgentAnsweredHour,omitempty"`
	// 预测外呼座席接听通话总时长
	//
	// example:
	//
	// 72
	PredictObAgentAnsweredTime *int64 `json:"PredictObAgentAnsweredTime,omitempty" xml:"PredictObAgentAnsweredTime,omitempty"`
	// 预测外呼座席接听数
	//
	// example:
	//
	// 27
	PredictObAnsweredCount *int64 `json:"PredictObAnsweredCount,omitempty" xml:"PredictObAnsweredCount,omitempty"`
	// 预测外呼接听量分布
	//
	// example:
	//
	// ""
	PredictObAnsweredDurationArray *string `json:"PredictObAnsweredDurationArray,omitempty" xml:"PredictObAnsweredDurationArray,omitempty"`
	// 预测外呼咨询数
	//
	// example:
	//
	// 0
	PredictObConsultCount *string `json:"PredictObConsultCount,omitempty" xml:"PredictObConsultCount,omitempty"`
	// 预测外呼监控强插数
	//
	// example:
	//
	// 0
	PredictObMonitorBargeCount *string `json:"PredictObMonitorBargeCount,omitempty" xml:"PredictObMonitorBargeCount,omitempty"`
	// 预测外呼强插时长
	//
	// example:
	//
	// 00:00:00
	PredictObMonitorBargeDuration *string `json:"PredictObMonitorBargeDuration,omitempty" xml:"PredictObMonitorBargeDuration,omitempty"`
	// 预测外呼监控强拆数
	//
	// example:
	//
	// 0
	PredictObMonitorDisconnectCount *string `json:"PredictObMonitorDisconnectCount,omitempty" xml:"PredictObMonitorDisconnectCount,omitempty"`
	// 预测外呼监控监听数
	//
	// example:
	//
	// 0
	PredictObMonitorSpyCount *string `json:"PredictObMonitorSpyCount,omitempty" xml:"PredictObMonitorSpyCount,omitempty"`
	// 预测外呼监听时长
	//
	// example:
	//
	// 00:00:00
	PredictObMonitorSpyDuration *string `json:"PredictObMonitorSpyDuration,omitempty" xml:"PredictObMonitorSpyDuration,omitempty"`
	// 预测外呼监控三方数
	//
	// example:
	//
	// 1
	PredictObMonitorThreewayCount *string `json:"PredictObMonitorThreewayCount,omitempty" xml:"PredictObMonitorThreewayCount,omitempty"`
	// 预测外呼监控耳语数
	//
	// example:
	//
	// 0
	PredictObMonitorWhisperCount *string `json:"PredictObMonitorWhisperCount,omitempty" xml:"PredictObMonitorWhisperCount,omitempty"`
	// 预测外呼耳语时长
	//
	// example:
	//
	// 00:00:00
	PredictObMonitorWhisperDuration *string `json:"PredictObMonitorWhisperDuration,omitempty" xml:"PredictObMonitorWhisperDuration,omitempty"`
	// 预测外呼总数
	//
	// example:
	//
	// 39
	PredictObTotalCount *int64 `json:"PredictObTotalCount,omitempty" xml:"PredictObTotalCount,omitempty"`
	// 预测外呼转移数
	//
	// example:
	//
	// 0
	PredictObTransferCount *string `json:"PredictObTransferCount,omitempty" xml:"PredictObTransferCount,omitempty"`
	// 预测外呼有效通话数
	//
	// example:
	//
	// 0
	PredictObValidCalls *int64 `json:"PredictObValidCalls,omitempty" xml:"PredictObValidCalls,omitempty"`
	// 预测外呼有效通话总小时
	//
	// example:
	//
	// 0
	PredictObValidTotalBridgeHour *string `json:"PredictObValidTotalBridgeHour,omitempty" xml:"PredictObValidTotalBridgeHour,omitempty"`
	// 预测外呼有效通话时长
	//
	// example:
	//
	// 33
	PredictObValidTotalBridgeTime *int64 `json:"PredictObValidTotalBridgeTime,omitempty" xml:"PredictObValidTotalBridgeTime,omitempty"`
	// 预览外呼座席接听率
	//
	// example:
	//
	// 33%
	PreviewAgentAnswerdRate *string `json:"PreviewAgentAnswerdRate,omitempty" xml:"PreviewAgentAnswerdRate,omitempty"`
	// 预览外呼客户接听率
	//
	// example:
	//
	// 0%
	PreviewCustomerAnswerdRate *string `json:"PreviewCustomerAnswerdRate,omitempty" xml:"PreviewCustomerAnswerdRate,omitempty"`
	// 预览外呼座席接听数
	//
	// example:
	//
	// 0
	PreviewObAgentAnsweredCount *string `json:"PreviewObAgentAnsweredCount,omitempty" xml:"PreviewObAgentAnsweredCount,omitempty"`
	// 预览外呼座席接听时长
	//
	// example:
	//
	// 00:00:01
	PreviewObAgentAnsweredTime *string `json:"PreviewObAgentAnsweredTime,omitempty" xml:"PreviewObAgentAnsweredTime,omitempty"`
	// 预览外呼双方接听数
	//
	// example:
	//
	// 0
	PreviewObAnsweredCount *string `json:"PreviewObAnsweredCount,omitempty" xml:"PreviewObAnsweredCount,omitempty"`
	// 预览外呼接听量分布
	PreviewObAnsweredDurationArray []*string `json:"PreviewObAnsweredDurationArray,omitempty" xml:"PreviewObAnsweredDurationArray,omitempty" type:"Repeated"`
	// 预览外呼接听量分布描述
	PreviewObAnsweredDurationDscArray []*string `json:"PreviewObAnsweredDurationDscArray,omitempty" xml:"PreviewObAnsweredDurationDscArray,omitempty" type:"Repeated"`
	// 预览外呼咨询数
	//
	// example:
	//
	// 0
	PreviewObConsultCount *string `json:"PreviewObConsultCount,omitempty" xml:"PreviewObConsultCount,omitempty"`
	// 预览外呼客户接听时长
	//
	// example:
	//
	// 00:00:00
	PreviewObCustomerAnsweredTime *string `json:"PreviewObCustomerAnsweredTime,omitempty" xml:"PreviewObCustomerAnsweredTime,omitempty"`
	// 响铃数
	//
	// example:
	//
	// 1
	PreviewObCustomerRingingCount *string `json:"PreviewObCustomerRingingCount,omitempty" xml:"PreviewObCustomerRingingCount,omitempty"`
	// 响铃率
	//
	// example:
	//
	// 100%
	PreviewObCustomerRingingRate *string `json:"PreviewObCustomerRingingRate,omitempty" xml:"PreviewObCustomerRingingRate,omitempty"`
	// 预览外呼监控强插数
	//
	// example:
	//
	// 0
	PreviewObMonitorBargeCount *string `json:"PreviewObMonitorBargeCount,omitempty" xml:"PreviewObMonitorBargeCount,omitempty"`
	// 预览外呼强插时长
	//
	// example:
	//
	// 00:00:00
	PreviewObMonitorBargeDuration *string `json:"PreviewObMonitorBargeDuration,omitempty" xml:"PreviewObMonitorBargeDuration,omitempty"`
	// 预览外呼监控强拆数
	//
	// example:
	//
	// 0
	PreviewObMonitorDisconnectCount *string `json:"PreviewObMonitorDisconnectCount,omitempty" xml:"PreviewObMonitorDisconnectCount,omitempty"`
	// 预览外呼监控监听数
	//
	// example:
	//
	// 0
	PreviewObMonitorSpyCount *string `json:"PreviewObMonitorSpyCount,omitempty" xml:"PreviewObMonitorSpyCount,omitempty"`
	// 预览外呼监听时长
	//
	// example:
	//
	// 0
	PreviewObMonitorSpyDuration *string `json:"PreviewObMonitorSpyDuration,omitempty" xml:"PreviewObMonitorSpyDuration,omitempty"`
	// 预览外呼监控三方数
	//
	// example:
	//
	// 0
	PreviewObMonitorThreewayCount *string `json:"PreviewObMonitorThreewayCount,omitempty" xml:"PreviewObMonitorThreewayCount,omitempty"`
	// 预览外呼监控耳语数
	//
	// example:
	//
	// 0
	PreviewObMonitorWhisperCount *string `json:"PreviewObMonitorWhisperCount,omitempty" xml:"PreviewObMonitorWhisperCount,omitempty"`
	// 预览外呼耳语时长
	//
	// example:
	//
	// 00:00:00
	PreviewObMonitorWhisperDuration *string `json:"PreviewObMonitorWhisperDuration,omitempty" xml:"PreviewObMonitorWhisperDuration,omitempty"`
	// 预览外呼双方接听通话小时
	//
	// example:
	//
	// 0
	PreviewObTotalBridgeHour *string `json:"PreviewObTotalBridgeHour,omitempty" xml:"PreviewObTotalBridgeHour,omitempty"`
	// 预览外呼双方接听通话时长
	//
	// example:
	//
	// 00:00:00
	PreviewObTotalBridgeTime *string `json:"PreviewObTotalBridgeTime,omitempty" xml:"PreviewObTotalBridgeTime,omitempty"`
	// 总预览外呼数
	//
	// example:
	//
	// 0
	PreviewObTotalCount *string `json:"PreviewObTotalCount,omitempty" xml:"PreviewObTotalCount,omitempty"`
	// 预览外呼被转接听数
	//
	// example:
	//
	// 1
	PreviewObTransferBridgeCount *string `json:"PreviewObTransferBridgeCount,omitempty" xml:"PreviewObTransferBridgeCount,omitempty"`
	// 预览外呼被转接听时长
	//
	// example:
	//
	// 00:00:00
	PreviewObTransferBridgeDuration *string `json:"PreviewObTransferBridgeDuration,omitempty" xml:"PreviewObTransferBridgeDuration,omitempty"`
	// 预览外呼被转数
	//
	// example:
	//
	// 1
	PreviewObTransferCount *string `json:"PreviewObTransferCount,omitempty" xml:"PreviewObTransferCount,omitempty"`
	// 预览外呼有效通话次数
	//
	// example:
	//
	// 0
	PreviewObValidCalls *string `json:"PreviewObValidCalls,omitempty" xml:"PreviewObValidCalls,omitempty"`
	// 预览外呼有效通话总小时
	//
	// example:
	//
	// 1
	PreviewObValidTotalBridgeHour *string `json:"PreviewObValidTotalBridgeHour,omitempty" xml:"PreviewObValidTotalBridgeHour,omitempty"`
	// 预览外呼有效总通话时长
	//
	// example:
	//
	// 00:00:22
	PreviewObValidTotalBridgeTime *string `json:"PreviewObValidTotalBridgeTime,omitempty" xml:"PreviewObValidTotalBridgeTime,omitempty"`
	// 队列名
	//
	// example:
	//
	// name3
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// 休息次数
	//
	// example:
	//
	// 0
	RestCount *string `json:"RestCount,omitempty" xml:"RestCount,omitempty"`
	// 总休息时长
	//
	// example:
	//
	// 00:00:00
	RestTime *string `json:"RestTime,omitempty" xml:"RestTime,omitempty"`
	// 行名称
	//
	// example:
	//
	// 示例值
	RowName *string `json:"RowName,omitempty" xml:"RowName,omitempty"`
	// 总双方接听时长
	//
	// example:
	//
	// 00:00:00
	TotalBridgeTime *string `json:"TotalBridgeTime,omitempty" xml:"TotalBridgeTime,omitempty"`
	// webcall座席接听分钟数
	//
	// example:
	//
	// 00:00:00
	WebcallObAgentAnsweredTime *string `json:"WebcallObAgentAnsweredTime,omitempty" xml:"WebcallObAgentAnsweredTime,omitempty"`
	// webcall双方接听数
	//
	// example:
	//
	// 0
	WebcallObAnsweredCount *string `json:"WebcallObAnsweredCount,omitempty" xml:"WebcallObAnsweredCount,omitempty"`
	// webcall接听量分布
	//
	// example:
	//
	// ""
	WebcallObAnsweredDurationArray *string `json:"WebcallObAnsweredDurationArray,omitempty" xml:"WebcallObAnsweredDurationArray,omitempty"`
	// webcall咨询数
	//
	// example:
	//
	// 0
	WebcallObConsultCount *string `json:"WebcallObConsultCount,omitempty" xml:"WebcallObConsultCount,omitempty"`
	// webcall监控强插数
	//
	// example:
	//
	// 0
	WebcallObMonitorBargeCount *string `json:"WebcallObMonitorBargeCount,omitempty" xml:"WebcallObMonitorBargeCount,omitempty"`
	// webcall强插时长
	//
	// example:
	//
	// 00:00:00
	WebcallObMonitorBargeDuration *string `json:"WebcallObMonitorBargeDuration,omitempty" xml:"WebcallObMonitorBargeDuration,omitempty"`
	// webcall监控强拆数
	//
	// example:
	//
	// 0
	WebcallObMonitorDisconnectCount *string `json:"WebcallObMonitorDisconnectCount,omitempty" xml:"WebcallObMonitorDisconnectCount,omitempty"`
	// webcall监控监听数
	//
	// example:
	//
	// 0
	WebcallObMonitorSpyCount *string `json:"WebcallObMonitorSpyCount,omitempty" xml:"WebcallObMonitorSpyCount,omitempty"`
	// webcall监听时长
	//
	// example:
	//
	// 00:00:00
	WebcallObMonitorSpyDuration *string `json:"WebcallObMonitorSpyDuration,omitempty" xml:"WebcallObMonitorSpyDuration,omitempty"`
	// webcall监控三方数
	//
	// example:
	//
	// 0
	WebcallObMonitorThreewayCount *string `json:"WebcallObMonitorThreewayCount,omitempty" xml:"WebcallObMonitorThreewayCount,omitempty"`
	// webcall监控耳语数
	//
	// example:
	//
	// 0
	WebcallObMonitorWhisperCount *string `json:"WebcallObMonitorWhisperCount,omitempty" xml:"WebcallObMonitorWhisperCount,omitempty"`
	// webcall耳语时长
	//
	// example:
	//
	// 00:00:00
	WebcallObMonitorWhisperDuration *string `json:"WebcallObMonitorWhisperDuration,omitempty" xml:"WebcallObMonitorWhisperDuration,omitempty"`
	// 总webcall数
	//
	// example:
	//
	// 0
	WebcallObTotalCount *string `json:"WebcallObTotalCount,omitempty" xml:"WebcallObTotalCount,omitempty"`
	// webcall转移数
	//
	// example:
	//
	// 0
	WebcallObTransferCount *string `json:"WebcallObTransferCount,omitempty" xml:"WebcallObTransferCount,omitempty"`
}

func (s CloudAgentWorkloadReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentWorkloadReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAgentAnswerRate() *string {
	return s.AgentAnswerRate
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAgentCreateTime() *string {
	return s.AgentCreateTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAgentRate() *string {
	return s.AgentRate
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAgentStopTime() *string {
	return s.AgentStopTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAvgIdleTime() *string {
	return s.AvgIdleTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAvgPauseTime() *string {
	return s.AvgPauseTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAvgPreviewObWaitTime() *string {
	return s.AvgPreviewObWaitTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetAvgRestTime() *string {
	return s.AvgRestTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetCallUtilization() *string {
	return s.CallUtilization
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDateTimeRange() *string {
	return s.DateTimeRange
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDay() *string {
	return s.Day
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObAnsweredCount() *string {
	return s.DirectObAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObAnsweredDurationArray() *string {
	return s.DirectObAnsweredDurationArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObAnsweredDurationDscArray() []*string {
	return s.DirectObAnsweredDurationDscArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObBridgeDuration() *string {
	return s.DirectObBridgeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObBridgeTime() *string {
	return s.DirectObBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObConsultCount() *string {
	return s.DirectObConsultCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObMonitorBargeCount() *string {
	return s.DirectObMonitorBargeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObMonitorDisconnectCount() *string {
	return s.DirectObMonitorDisconnectCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObMonitorSpyCount() *string {
	return s.DirectObMonitorSpyCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObMonitorThreewayCount() *string {
	return s.DirectObMonitorThreewayCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObMonitorWhisperCount() *string {
	return s.DirectObMonitorWhisperCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObTotalCount() *string {
	return s.DirectObTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObTotalTime() *string {
	return s.DirectObTotalTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObTransferCount() *string {
	return s.DirectObTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObValidCalls() *string {
	return s.DirectObValidCalls
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetDirectObValidTotalBridgeTime() *string {
	return s.DirectObValidTotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetHoursUtilization() *string {
	return s.HoursUtilization
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAgentHangupCount() *string {
	return s.IbAgentHangupCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAgentRefuseCount() *string {
	return s.IbAgentRefuseCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAnsweredCount() *string {
	return s.IbAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAnsweredDurationArray() []*string {
	return s.IbAnsweredDurationArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAnsweredDurationDscArray() []*string {
	return s.IbAnsweredDurationDscArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAvgAnswerDelayTime() *string {
	return s.IbAvgAnswerDelayTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAvgBridgeTime() *string {
	return s.IbAvgBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAvgCallingTime() *string {
	return s.IbAvgCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAvgHoldTime() *string {
	return s.IbAvgHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbAvgWrapupTime() *string {
	return s.IbAvgWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbBridgeCount() *string {
	return s.IbBridgeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbBridgeMinute() *string {
	return s.IbBridgeMinute
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbBridgeTime() *string {
	return s.IbBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbCallingCount() *string {
	return s.IbCallingCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbCallingTime() *string {
	return s.IbCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbConsultCount() *string {
	return s.IbConsultCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbCustomerHangupCount() *string {
	return s.IbCustomerHangupCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbHoldCount() *string {
	return s.IbHoldCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbHoldTime() *string {
	return s.IbHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMaxAnswerDelayTime() *string {
	return s.IbMaxAnswerDelayTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMaxBridgeTime() *string {
	return s.IbMaxBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMaxCallingTime() *string {
	return s.IbMaxCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMaxHoldTime() *string {
	return s.IbMaxHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMaxWrapupTime() *string {
	return s.IbMaxWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMinAnswerDelayTime() *string {
	return s.IbMinAnswerDelayTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMinBridgeTime() *string {
	return s.IbMinBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMinCallingTime() *string {
	return s.IbMinCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMinHoldTime() *string {
	return s.IbMinHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMinWrapupTime() *string {
	return s.IbMinWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorBargeCount() *string {
	return s.IbMonitorBargeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorBargeDuration() *string {
	return s.IbMonitorBargeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorDisconnectCount() *string {
	return s.IbMonitorDisconnectCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorPickupCount() *string {
	return s.IbMonitorPickupCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorPickupDuration() *string {
	return s.IbMonitorPickupDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorSpyCount() *string {
	return s.IbMonitorSpyCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorSpyDuration() *string {
	return s.IbMonitorSpyDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorThreewayCount() *string {
	return s.IbMonitorThreewayCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorWhisperCount() *string {
	return s.IbMonitorWhisperCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbMonitorWhisperDuration() *string {
	return s.IbMonitorWhisperDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbTotalAnswerDelayTime() *string {
	return s.IbTotalAnswerDelayTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbTotalCount() *string {
	return s.IbTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbTotalMinute() *string {
	return s.IbTotalMinute
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbTotalTime() *string {
	return s.IbTotalTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbTransferCount() *string {
	return s.IbTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbUnansweredCount() *string {
	return s.IbUnansweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbUniqueAnsweredCount() *string {
	return s.IbUniqueAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbUniqueTotalCount() *string {
	return s.IbUniqueTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbUniqueUnansweredCount() *string {
	return s.IbUniqueUnansweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbValidCalls() *string {
	return s.IbValidCalls
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbValidTotalBridgeTime() *string {
	return s.IbValidTotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbWrapupCount() *string {
	return s.IbWrapupCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIbWrapupTime() *string {
	return s.IbWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIdleCount() *string {
	return s.IdleCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIdleTime() *string {
	return s.IdleTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIvrTransferBridgeCount() *string {
	return s.IvrTransferBridgeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIvrTransferBridgeDuration() *string {
	return s.IvrTransferBridgeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetIvrTransferCount() *string {
	return s.IvrTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetLoginTime() *string {
	return s.LoginTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMaxIdleTime() *string {
	return s.MaxIdleTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMaxPauseTime() *string {
	return s.MaxPauseTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMaxRestTime() *string {
	return s.MaxRestTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMinIdleTime() *string {
	return s.MinIdleTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMinPauseTime() *string {
	return s.MinPauseTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetMinRestTime() *string {
	return s.MinRestTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObAvgBridgeTime() *string {
	return s.ObAvgBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObAvgCallingTime() *string {
	return s.ObAvgCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObAvgHoldTime() *string {
	return s.ObAvgHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObAvgWrapupTime() *string {
	return s.ObAvgWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObBridgeCount() *string {
	return s.ObBridgeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObBridgeTime() *string {
	return s.ObBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObCallingCount() *string {
	return s.ObCallingCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObCallingTime() *string {
	return s.ObCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObHoldCount() *string {
	return s.ObHoldCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObHoldTime() *string {
	return s.ObHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMaxBridgeTime() *string {
	return s.ObMaxBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMaxCallingTime() *string {
	return s.ObMaxCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMaxHoldTime() *string {
	return s.ObMaxHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMaxWrapupTime() *string {
	return s.ObMaxWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMinBridgeTime() *string {
	return s.ObMinBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMinCallingTime() *string {
	return s.ObMinCallingTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMinHoldTime() *string {
	return s.ObMinHoldTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObMinWrapupTime() *string {
	return s.ObMinWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObRealBridgeTime() *string {
	return s.ObRealBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObWrapupCount() *string {
	return s.ObWrapupCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetObWrapupTime() *string {
	return s.ObWrapupTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPauseCount() *string {
	return s.PauseCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPauseTime() *string {
	return s.PauseTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObAgentAnsweredHour() *string {
	return s.PredictObAgentAnsweredHour
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObAgentAnsweredTime() *int64 {
	return s.PredictObAgentAnsweredTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObAnsweredCount() *int64 {
	return s.PredictObAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObAnsweredDurationArray() *string {
	return s.PredictObAnsweredDurationArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObConsultCount() *string {
	return s.PredictObConsultCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorBargeCount() *string {
	return s.PredictObMonitorBargeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorBargeDuration() *string {
	return s.PredictObMonitorBargeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorDisconnectCount() *string {
	return s.PredictObMonitorDisconnectCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorSpyCount() *string {
	return s.PredictObMonitorSpyCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorSpyDuration() *string {
	return s.PredictObMonitorSpyDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorThreewayCount() *string {
	return s.PredictObMonitorThreewayCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorWhisperCount() *string {
	return s.PredictObMonitorWhisperCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObMonitorWhisperDuration() *string {
	return s.PredictObMonitorWhisperDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObTotalCount() *int64 {
	return s.PredictObTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObTransferCount() *string {
	return s.PredictObTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObValidCalls() *int64 {
	return s.PredictObValidCalls
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObValidTotalBridgeHour() *string {
	return s.PredictObValidTotalBridgeHour
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPredictObValidTotalBridgeTime() *int64 {
	return s.PredictObValidTotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewAgentAnswerdRate() *string {
	return s.PreviewAgentAnswerdRate
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewCustomerAnswerdRate() *string {
	return s.PreviewCustomerAnswerdRate
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObAgentAnsweredCount() *string {
	return s.PreviewObAgentAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObAgentAnsweredTime() *string {
	return s.PreviewObAgentAnsweredTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObAnsweredCount() *string {
	return s.PreviewObAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObAnsweredDurationArray() []*string {
	return s.PreviewObAnsweredDurationArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObAnsweredDurationDscArray() []*string {
	return s.PreviewObAnsweredDurationDscArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObConsultCount() *string {
	return s.PreviewObConsultCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObCustomerAnsweredTime() *string {
	return s.PreviewObCustomerAnsweredTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObCustomerRingingCount() *string {
	return s.PreviewObCustomerRingingCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObCustomerRingingRate() *string {
	return s.PreviewObCustomerRingingRate
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorBargeCount() *string {
	return s.PreviewObMonitorBargeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorBargeDuration() *string {
	return s.PreviewObMonitorBargeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorDisconnectCount() *string {
	return s.PreviewObMonitorDisconnectCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorSpyCount() *string {
	return s.PreviewObMonitorSpyCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorSpyDuration() *string {
	return s.PreviewObMonitorSpyDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorThreewayCount() *string {
	return s.PreviewObMonitorThreewayCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorWhisperCount() *string {
	return s.PreviewObMonitorWhisperCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObMonitorWhisperDuration() *string {
	return s.PreviewObMonitorWhisperDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTotalBridgeHour() *string {
	return s.PreviewObTotalBridgeHour
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTotalBridgeTime() *string {
	return s.PreviewObTotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTotalCount() *string {
	return s.PreviewObTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTransferBridgeCount() *string {
	return s.PreviewObTransferBridgeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTransferBridgeDuration() *string {
	return s.PreviewObTransferBridgeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObTransferCount() *string {
	return s.PreviewObTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObValidCalls() *string {
	return s.PreviewObValidCalls
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObValidTotalBridgeHour() *string {
	return s.PreviewObValidTotalBridgeHour
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetPreviewObValidTotalBridgeTime() *string {
	return s.PreviewObValidTotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetQueueName() *string {
	return s.QueueName
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetRestCount() *string {
	return s.RestCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetRestTime() *string {
	return s.RestTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetRowName() *string {
	return s.RowName
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetTotalBridgeTime() *string {
	return s.TotalBridgeTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObAgentAnsweredTime() *string {
	return s.WebcallObAgentAnsweredTime
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObAnsweredCount() *string {
	return s.WebcallObAnsweredCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObAnsweredDurationArray() *string {
	return s.WebcallObAnsweredDurationArray
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObConsultCount() *string {
	return s.WebcallObConsultCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorBargeCount() *string {
	return s.WebcallObMonitorBargeCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorBargeDuration() *string {
	return s.WebcallObMonitorBargeDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorDisconnectCount() *string {
	return s.WebcallObMonitorDisconnectCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorSpyCount() *string {
	return s.WebcallObMonitorSpyCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorSpyDuration() *string {
	return s.WebcallObMonitorSpyDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorThreewayCount() *string {
	return s.WebcallObMonitorThreewayCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorWhisperCount() *string {
	return s.WebcallObMonitorWhisperCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObMonitorWhisperDuration() *string {
	return s.WebcallObMonitorWhisperDuration
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObTotalCount() *string {
	return s.WebcallObTotalCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) GetWebcallObTransferCount() *string {
	return s.WebcallObTransferCount
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAgentAnswerRate(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AgentAnswerRate = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAgentCreateTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AgentCreateTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAgentName(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAgentRate(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AgentRate = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAgentStopTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AgentStopTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAvgIdleTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AvgIdleTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAvgPauseTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AvgPauseTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAvgPreviewObWaitTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AvgPreviewObWaitTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetAvgRestTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.AvgRestTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetCallUtilization(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.CallUtilization = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetCno(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetCreateTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDateTimeRange(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DateTimeRange = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDay(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.Day = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObAnsweredDurationArray(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObAnsweredDurationArray = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObAnsweredDurationDscArray(v []*string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObAnsweredDurationDscArray = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObBridgeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObBridgeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObConsultCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObConsultCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObMonitorBargeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObMonitorBargeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObMonitorDisconnectCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObMonitorDisconnectCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObMonitorSpyCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObMonitorSpyCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObMonitorThreewayCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObMonitorThreewayCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObMonitorWhisperCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObMonitorWhisperCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObTotalCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObTotalTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObTotalTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObValidCalls(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObValidCalls = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetDirectObValidTotalBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.DirectObValidTotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetEnterpriseId(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetHoursUtilization(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.HoursUtilization = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAgentHangupCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAgentHangupCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAgentRefuseCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAgentRefuseCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAnsweredDurationArray(v []*string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAnsweredDurationArray = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAnsweredDurationDscArray(v []*string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAnsweredDurationDscArray = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAvgAnswerDelayTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAvgAnswerDelayTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAvgBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAvgBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAvgCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAvgCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAvgHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAvgHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbAvgWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbAvgWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbBridgeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbBridgeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbBridgeMinute(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbBridgeMinute = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbCallingCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbCallingCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbConsultCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbConsultCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbCustomerHangupCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbCustomerHangupCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbHoldCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbHoldCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMaxAnswerDelayTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMaxAnswerDelayTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMaxBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMaxBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMaxCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMaxCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMaxHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMaxHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMaxWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMaxWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMinAnswerDelayTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMinAnswerDelayTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMinBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMinBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMinCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMinCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMinHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMinHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMinWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMinWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorBargeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorBargeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorBargeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorBargeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorDisconnectCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorDisconnectCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorPickupCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorPickupCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorPickupDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorPickupDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorSpyCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorSpyCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorSpyDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorSpyDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorThreewayCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorThreewayCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorWhisperCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorWhisperCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbMonitorWhisperDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbMonitorWhisperDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbTotalAnswerDelayTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbTotalAnswerDelayTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbTotalCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbTotalMinute(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbTotalMinute = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbTotalTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbTotalTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbUnansweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbUnansweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbUniqueAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbUniqueAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbUniqueTotalCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbUniqueTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbUniqueUnansweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbUniqueUnansweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbValidCalls(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbValidCalls = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbValidTotalBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbValidTotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbWrapupCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbWrapupCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIbWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IbWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetId(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIdleCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IdleCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIdleTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IdleTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIvrTransferBridgeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IvrTransferBridgeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIvrTransferBridgeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IvrTransferBridgeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetIvrTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.IvrTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetLoginTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.LoginTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMaxIdleTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MaxIdleTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMaxPauseTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MaxPauseTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMaxRestTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MaxRestTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMinIdleTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MinIdleTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMinPauseTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MinPauseTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetMinRestTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.MinRestTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObAvgBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObAvgBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObAvgCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObAvgCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObAvgHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObAvgHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObAvgWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObAvgWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObBridgeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObBridgeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObCallingCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObCallingCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObHoldCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObHoldCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMaxBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMaxBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMaxCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMaxCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMaxHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMaxHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMaxWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMaxWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMinBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMinBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMinCallingTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMinCallingTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMinHoldTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMinHoldTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObMinWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObMinWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObRealBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObRealBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObWrapupCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObWrapupCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetObWrapupTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.ObWrapupTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPauseCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PauseCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPauseTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PauseTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObAgentAnsweredHour(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObAgentAnsweredHour = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObAgentAnsweredTime(v int64) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObAgentAnsweredTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObAnsweredCount(v int64) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObAnsweredDurationArray(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObAnsweredDurationArray = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObConsultCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObConsultCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorBargeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorBargeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorBargeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorBargeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorDisconnectCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorDisconnectCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorSpyCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorSpyCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorSpyDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorSpyDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorThreewayCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorThreewayCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorWhisperCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorWhisperCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObMonitorWhisperDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObMonitorWhisperDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObTotalCount(v int64) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObValidCalls(v int64) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObValidCalls = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObValidTotalBridgeHour(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObValidTotalBridgeHour = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPredictObValidTotalBridgeTime(v int64) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PredictObValidTotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewAgentAnswerdRate(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewAgentAnswerdRate = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewCustomerAnswerdRate(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewCustomerAnswerdRate = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObAgentAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObAgentAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObAgentAnsweredTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObAgentAnsweredTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObAnsweredDurationArray(v []*string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObAnsweredDurationArray = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObAnsweredDurationDscArray(v []*string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObAnsweredDurationDscArray = v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObConsultCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObConsultCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObCustomerAnsweredTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObCustomerAnsweredTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObCustomerRingingCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObCustomerRingingCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObCustomerRingingRate(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObCustomerRingingRate = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorBargeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorBargeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorBargeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorBargeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorDisconnectCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorDisconnectCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorSpyCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorSpyCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorSpyDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorSpyDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorThreewayCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorThreewayCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorWhisperCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorWhisperCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObMonitorWhisperDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObMonitorWhisperDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTotalBridgeHour(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTotalBridgeHour = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTotalBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTotalCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTransferBridgeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTransferBridgeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTransferBridgeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTransferBridgeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObValidCalls(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObValidCalls = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObValidTotalBridgeHour(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObValidTotalBridgeHour = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetPreviewObValidTotalBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.PreviewObValidTotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetQueueName(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.QueueName = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetRestCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.RestCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetRestTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.RestTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetRowName(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.RowName = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetTotalBridgeTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.TotalBridgeTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObAgentAnsweredTime(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObAgentAnsweredTime = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObAnsweredCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObAnsweredCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObAnsweredDurationArray(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObAnsweredDurationArray = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObConsultCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObConsultCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorBargeCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorBargeCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorBargeDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorBargeDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorDisconnectCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorDisconnectCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorSpyCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorSpyCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorSpyDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorSpyDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorThreewayCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorThreewayCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorWhisperCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorWhisperCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObMonitorWhisperDuration(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObMonitorWhisperDuration = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObTotalCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObTotalCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) SetWebcallObTransferCount(v string) *CloudAgentWorkloadReportResponseBodyDataList {
	s.WebcallObTransferCount = &v
	return s
}

func (s *CloudAgentWorkloadReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
