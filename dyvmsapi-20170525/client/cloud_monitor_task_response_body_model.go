// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudMonitorTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudMonitorTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudMonitorTaskResponseBody
	GetCode() *string
	SetData(v *CloudMonitorTaskResponseBodyData) *CloudMonitorTaskResponseBody
	GetData() *CloudMonitorTaskResponseBodyData
	SetMessage(v string) *CloudMonitorTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudMonitorTaskResponseBody
	GetRequestId() *string
}

type CloudMonitorTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudMonitorTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudMonitorTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudMonitorTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudMonitorTaskResponseBody) GetData() *CloudMonitorTaskResponseBodyData {
	return s.Data
}

func (s *CloudMonitorTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudMonitorTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudMonitorTaskResponseBody) SetAccessDeniedDetail(v string) *CloudMonitorTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudMonitorTaskResponseBody) SetCode(v string) *CloudMonitorTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudMonitorTaskResponseBody) SetData(v *CloudMonitorTaskResponseBodyData) *CloudMonitorTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudMonitorTaskResponseBody) SetMessage(v string) *CloudMonitorTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudMonitorTaskResponseBody) SetRequestId(v string) *CloudMonitorTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudMonitorTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudMonitorTaskResponseBodyData struct {
	// 座席状态统计
	AgentInfo *CloudMonitorTaskResponseBodyDataAgentInfo `json:"AgentInfo,omitempty" xml:"AgentInfo,omitempty" type:"Struct"`
	// 座席统计详情
	AgentStatisticList []*CloudMonitorTaskResponseBodyDataAgentStatisticList `json:"AgentStatisticList,omitempty" xml:"AgentStatisticList,omitempty" type:"Repeated"`
	// 通道监控信息
	ChannelInfo *CloudMonitorTaskResponseBodyDataChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Struct"`
	// 已移除座席统计详情，仅当 `includeRemovedAgents=1` 时返回
	RemovedAgentStatisticList []*CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList `json:"RemovedAgentStatisticList,omitempty" xml:"RemovedAgentStatisticList,omitempty" type:"Repeated"`
	// 任务监控信息
	Statistic *CloudMonitorTaskResponseBodyDataStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	// 当日任务监控信息
	TodayStatistic *CloudMonitorTaskResponseBodyDataTodayStatistic `json:"TodayStatistic,omitempty" xml:"TodayStatistic,omitempty" type:"Struct"`
}

func (s CloudMonitorTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyData) GetAgentInfo() *CloudMonitorTaskResponseBodyDataAgentInfo {
	return s.AgentInfo
}

func (s *CloudMonitorTaskResponseBodyData) GetAgentStatisticList() []*CloudMonitorTaskResponseBodyDataAgentStatisticList {
	return s.AgentStatisticList
}

func (s *CloudMonitorTaskResponseBodyData) GetChannelInfo() *CloudMonitorTaskResponseBodyDataChannelInfo {
	return s.ChannelInfo
}

func (s *CloudMonitorTaskResponseBodyData) GetRemovedAgentStatisticList() []*CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	return s.RemovedAgentStatisticList
}

func (s *CloudMonitorTaskResponseBodyData) GetStatistic() *CloudMonitorTaskResponseBodyDataStatistic {
	return s.Statistic
}

func (s *CloudMonitorTaskResponseBodyData) GetTodayStatistic() *CloudMonitorTaskResponseBodyDataTodayStatistic {
	return s.TodayStatistic
}

func (s *CloudMonitorTaskResponseBodyData) SetAgentInfo(v *CloudMonitorTaskResponseBodyDataAgentInfo) *CloudMonitorTaskResponseBodyData {
	s.AgentInfo = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) SetAgentStatisticList(v []*CloudMonitorTaskResponseBodyDataAgentStatisticList) *CloudMonitorTaskResponseBodyData {
	s.AgentStatisticList = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) SetChannelInfo(v *CloudMonitorTaskResponseBodyDataChannelInfo) *CloudMonitorTaskResponseBodyData {
	s.ChannelInfo = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) SetRemovedAgentStatisticList(v []*CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) *CloudMonitorTaskResponseBodyData {
	s.RemovedAgentStatisticList = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) SetStatistic(v *CloudMonitorTaskResponseBodyDataStatistic) *CloudMonitorTaskResponseBodyData {
	s.Statistic = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) SetTodayStatistic(v *CloudMonitorTaskResponseBodyDataTodayStatistic) *CloudMonitorTaskResponseBodyData {
	s.TodayStatistic = v
	return s
}

func (s *CloudMonitorTaskResponseBodyData) Validate() error {
	if s.AgentInfo != nil {
		if err := s.AgentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.AgentStatisticList != nil {
		for _, item := range s.AgentStatisticList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChannelInfo != nil {
		if err := s.ChannelInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RemovedAgentStatisticList != nil {
		for _, item := range s.RemovedAgentStatisticList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistic != nil {
		if err := s.Statistic.Validate(); err != nil {
			return err
		}
	}
	if s.TodayStatistic != nil {
		if err := s.TodayStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudMonitorTaskResponseBodyDataAgentInfo struct {
	// 可用座席数
	//
	// example:
	//
	// 2
	AvailableCount *int64 `json:"AvailableCount,omitempty" xml:"AvailableCount,omitempty"`
	// 通话中的座席数
	//
	// example:
	//
	// 6
	BridgeCount *int64 `json:"BridgeCount,omitempty" xml:"BridgeCount,omitempty"`
	// 当前在线座席数
	//
	// example:
	//
	// 2
	OnlineCount *int64 `json:"OnlineCount,omitempty" xml:"OnlineCount,omitempty"`
	// 响铃中的座席数
	//
	// example:
	//
	// 4
	RingingCount *int64 `json:"RingingCount,omitempty" xml:"RingingCount,omitempty"`
	// 座席总数
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 整理中的座席数
	//
	// example:
	//
	// 4
	WrapupCount *int64 `json:"WrapupCount,omitempty" xml:"WrapupCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataAgentInfo) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataAgentInfo) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetAvailableCount() *int64 {
	return s.AvailableCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetBridgeCount() *int64 {
	return s.BridgeCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetOnlineCount() *int64 {
	return s.OnlineCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetRingingCount() *int64 {
	return s.RingingCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) GetWrapupCount() *int64 {
	return s.WrapupCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetAvailableCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.AvailableCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetBridgeCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.BridgeCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetOnlineCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.OnlineCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetRingingCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.RingingCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetTotalCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.TotalCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) SetWrapupCount(v int64) *CloudMonitorTaskResponseBodyDataAgentInfo {
	s.WrapupCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentInfo) Validate() error {
	return dara.Validate(s)
}

type CloudMonitorTaskResponseBodyDataAgentStatisticList struct {
	// 座席平均进线时长,空闲时长/呼叫数
	//
	// example:
	//
	// 0
	AgentAvgIdleTime *int64 `json:"AgentAvgIdleTime,omitempty" xml:"AgentAvgIdleTime,omitempty"`
	// 座席当前状态:  stateIdle(空闲)  statePause(置忙)  stateInuse(通话） stateCalling(呼叫中)  stateWrapup(整理)   为空代表该当前时间座席在该任务中没有相关状态，如：座席未被使用等
	//
	// example:
	//
	// stateIdle
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 座席平均通话时长，单位秒数
	//
	// example:
	//
	// 116
	AvgBridgeTime *string `json:"AvgBridgeTime,omitempty" xml:"AvgBridgeTime,omitempty"`
	// 接听数
	//
	// example:
	//
	// 0
	BridgeCount *int64 `json:"BridgeCount,omitempty" xml:"BridgeCount,omitempty"`
	// 座席自动应答率=(座席接通时间-开始呼叫座席时间)<=2s的通话数/座席接通数
	//
	// example:
	//
	// 0.0
	BridgeIn2sRate *float64 `json:"BridgeIn2sRate,omitempty" xml:"BridgeIn2sRate,omitempty"`
	// 任务级别，座席单日接听上限，-1代表未配置接听上限
	//
	// example:
	//
	// -1
	CallLimit *int64 `json:"CallLimit,omitempty" xml:"CallLimit,omitempty"`
	// 呼叫数
	//
	// example:
	//
	// 0
	CalledCount *int64 `json:"CalledCount,omitempty" xml:"CalledCount,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 企业级别 座席单日接听上限，-1代表未配置接听上限
	//
	// example:
	//
	// -1
	EnterpriseCallLimit *int64 `json:"EnterpriseCallLimit,omitempty" xml:"EnterpriseCallLimit,omitempty"`
	// 企业级别 座席今日已接听数
	//
	// example:
	//
	// 3
	EnterpriseTodayAnsweredCount *int64 `json:"EnterpriseTodayAnsweredCount,omitempty" xml:"EnterpriseTodayAnsweredCount,omitempty"`
	// 座席名称
	//
	// example:
	//
	// name2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席未接听数
	//
	// example:
	//
	// 0
	NoBridgeCount *int64 `json:"NoBridgeCount,omitempty" xml:"NoBridgeCount,omitempty"`
	// 任务级别 座席今日已接听数
	//
	// example:
	//
	// 3
	TodayAnsweredCount *int64 `json:"TodayAnsweredCount,omitempty" xml:"TodayAnsweredCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataAgentStatisticList) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataAgentStatisticList) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetAgentAvgIdleTime() *int64 {
	return s.AgentAvgIdleTime
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetAvgBridgeTime() *string {
	return s.AvgBridgeTime
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetBridgeCount() *int64 {
	return s.BridgeCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetBridgeIn2sRate() *float64 {
	return s.BridgeIn2sRate
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetCallLimit() *int64 {
	return s.CallLimit
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetCalledCount() *int64 {
	return s.CalledCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetCno() *string {
	return s.Cno
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetEnterpriseCallLimit() *int64 {
	return s.EnterpriseCallLimit
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetEnterpriseTodayAnsweredCount() *int64 {
	return s.EnterpriseTodayAnsweredCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetName() *string {
	return s.Name
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetNoBridgeCount() *int64 {
	return s.NoBridgeCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) GetTodayAnsweredCount() *int64 {
	return s.TodayAnsweredCount
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetAgentAvgIdleTime(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.AgentAvgIdleTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetAgentStatus(v string) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.AgentStatus = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetAvgBridgeTime(v string) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.AvgBridgeTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetBridgeCount(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.BridgeCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetBridgeIn2sRate(v float64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.BridgeIn2sRate = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetCallLimit(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.CallLimit = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetCalledCount(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.CalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetCno(v string) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.Cno = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetEnterpriseCallLimit(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.EnterpriseCallLimit = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetEnterpriseTodayAnsweredCount(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.EnterpriseTodayAnsweredCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetName(v string) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.Name = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetNoBridgeCount(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.NoBridgeCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) SetTodayAnsweredCount(v int64) *CloudMonitorTaskResponseBodyDataAgentStatisticList {
	s.TodayAnsweredCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataAgentStatisticList) Validate() error {
	return dara.Validate(s)
}

type CloudMonitorTaskResponseBodyDataChannelInfo struct {
	// 客户接通个数
	//
	// example:
	//
	// 0
	AnswerCount *int64 `json:"AnswerCount,omitempty" xml:"AnswerCount,omitempty"`
	// 桥接个数
	//
	// example:
	//
	// 0
	BridgedCount *int64 `json:"BridgedCount,omitempty" xml:"BridgedCount,omitempty"`
	// 任务当前通道数
	//
	// example:
	//
	// 0
	ChannelCount *int64 `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	// 企业预测外呼通道数
	//
	// example:
	//
	// 2
	EnterpriseChannelCount *int64 `json:"EnterpriseChannelCount,omitempty" xml:"EnterpriseChannelCount,omitempty"`
	// 溢出到ivr数
	//
	// example:
	//
	// 0
	IvrCount *int64 `json:"IvrCount,omitempty" xml:"IvrCount,omitempty"`
	// 等待转座席的个数
	//
	// example:
	//
	// 0
	WaitAgentCount *int64 `json:"WaitAgentCount,omitempty" xml:"WaitAgentCount,omitempty"`
	// 等待客户接听的个数
	//
	// example:
	//
	// 0
	WaitAnswerCount *int64 `json:"WaitAnswerCount,omitempty" xml:"WaitAnswerCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataChannelInfo) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataChannelInfo) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetAnswerCount() *int64 {
	return s.AnswerCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetBridgedCount() *int64 {
	return s.BridgedCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetChannelCount() *int64 {
	return s.ChannelCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetEnterpriseChannelCount() *int64 {
	return s.EnterpriseChannelCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetIvrCount() *int64 {
	return s.IvrCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetWaitAgentCount() *int64 {
	return s.WaitAgentCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) GetWaitAnswerCount() *int64 {
	return s.WaitAnswerCount
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetAnswerCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.AnswerCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetBridgedCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.BridgedCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetChannelCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.ChannelCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetEnterpriseChannelCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.EnterpriseChannelCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetIvrCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.IvrCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetWaitAgentCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.WaitAgentCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) SetWaitAnswerCount(v int64) *CloudMonitorTaskResponseBodyDataChannelInfo {
	s.WaitAnswerCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataChannelInfo) Validate() error {
	return dara.Validate(s)
}

type CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList struct {
	// 座席平均进线时长,空闲时长/呼叫数
	//
	// example:
	//
	// 13
	AgentAvgIdleTime *int64 `json:"AgentAvgIdleTime,omitempty" xml:"AgentAvgIdleTime,omitempty"`
	// 座席当前状态:  stateIdle(空闲)  statePause(置忙)  stateInuse(通话） stateCalling(呼叫中)  stateWrapup(整理)   为空代表该当前时间座席在该任务中没有相关状态，如：座席未被使用等
	//
	// example:
	//
	// stateIdle
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// 座席平均通话时长，单位秒数
	//
	// example:
	//
	// 39
	AvgBridgeTime *string `json:"AvgBridgeTime,omitempty" xml:"AvgBridgeTime,omitempty"`
	// 接听数
	//
	// example:
	//
	// 3
	BridgeCount *int64 `json:"BridgeCount,omitempty" xml:"BridgeCount,omitempty"`
	// 座席自动应答率=(座席接通时间-开始呼叫座席时间)<=2s的通话数/座席接通数
	//
	// example:
	//
	// 3.55
	BridgeIn2sRate *float64 `json:"BridgeIn2sRate,omitempty" xml:"BridgeIn2sRate,omitempty"`
	// 任务级别，座席单日接听上限，-1代表未配置接听上限
	//
	// example:
	//
	// -1
	CallLimit *int64 `json:"CallLimit,omitempty" xml:"CallLimit,omitempty"`
	// 呼叫数
	//
	// example:
	//
	// 5
	CalledCount *int64 `json:"CalledCount,omitempty" xml:"CalledCount,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 企业级别 座席单日接听上限，-1代表未配置接听上限
	//
	// example:
	//
	// -1
	EnterpriseCallLimit *int64 `json:"EnterpriseCallLimit,omitempty" xml:"EnterpriseCallLimit,omitempty"`
	// 企业级别 座席今日已接听数
	//
	// example:
	//
	// 3
	EnterpriseTodayAnsweredCount *int64 `json:"EnterpriseTodayAnsweredCount,omitempty" xml:"EnterpriseTodayAnsweredCount,omitempty"`
	// 座席名称
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席未接听数
	//
	// example:
	//
	// 3
	NoBridgeCount *int64 `json:"NoBridgeCount,omitempty" xml:"NoBridgeCount,omitempty"`
	// 任务级别 座席今日已接听数
	//
	// example:
	//
	// 3
	TodayAnsweredCount *int64 `json:"TodayAnsweredCount,omitempty" xml:"TodayAnsweredCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetAgentAvgIdleTime() *int64 {
	return s.AgentAvgIdleTime
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetAvgBridgeTime() *string {
	return s.AvgBridgeTime
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetBridgeCount() *int64 {
	return s.BridgeCount
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetBridgeIn2sRate() *float64 {
	return s.BridgeIn2sRate
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetCallLimit() *int64 {
	return s.CallLimit
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetCalledCount() *int64 {
	return s.CalledCount
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetCno() *string {
	return s.Cno
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetEnterpriseCallLimit() *int64 {
	return s.EnterpriseCallLimit
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetEnterpriseTodayAnsweredCount() *int64 {
	return s.EnterpriseTodayAnsweredCount
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetName() *string {
	return s.Name
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetNoBridgeCount() *int64 {
	return s.NoBridgeCount
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) GetTodayAnsweredCount() *int64 {
	return s.TodayAnsweredCount
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetAgentAvgIdleTime(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.AgentAvgIdleTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetAgentStatus(v string) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.AgentStatus = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetAvgBridgeTime(v string) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.AvgBridgeTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetBridgeCount(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.BridgeCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetBridgeIn2sRate(v float64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.BridgeIn2sRate = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetCallLimit(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.CallLimit = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetCalledCount(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.CalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetCno(v string) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.Cno = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetEnterpriseCallLimit(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.EnterpriseCallLimit = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetEnterpriseTodayAnsweredCount(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.EnterpriseTodayAnsweredCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetName(v string) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.Name = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetNoBridgeCount(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.NoBridgeCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) SetTodayAnsweredCount(v int64) *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList {
	s.TodayAnsweredCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataRemovedAgentStatisticList) Validate() error {
	return dara.Validate(s)
}

type CloudMonitorTaskResponseBodyDataStatistic struct {
	// 座席数量
	//
	// example:
	//
	// 2
	AgentCount *int64 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// 应答数量
	//
	// example:
	//
	// 3
	AnswerCount *int64 `json:"AnswerCount,omitempty" xml:"AnswerCount,omitempty"`
	// 呼叫坐席数
	//
	// example:
	//
	// 3
	CallAgentCount *int64 `json:"CallAgentCount,omitempty" xml:"CallAgentCount,omitempty"`
	// 已呼叫数量
	//
	// example:
	//
	// 4
	CalledCount *int64 `json:"CalledCount,omitempty" xml:"CalledCount,omitempty"`
	// 任务属性
	CtiLinkTaskProperty *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty `json:"CtiLinkTaskProperty,omitempty" xml:"CtiLinkTaskProperty,omitempty" type:"Struct"`
	// 未呼叫数量
	//
	// example:
	//
	// 1
	UncalledCount *int64 `json:"UncalledCount,omitempty" xml:"UncalledCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataStatistic) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataStatistic) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetAnswerCount() *int64 {
	return s.AnswerCount
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetCallAgentCount() *int64 {
	return s.CallAgentCount
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetCalledCount() *int64 {
	return s.CalledCount
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetCtiLinkTaskProperty() *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	return s.CtiLinkTaskProperty
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) GetUncalledCount() *int64 {
	return s.UncalledCount
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetAgentCount(v int64) *CloudMonitorTaskResponseBodyDataStatistic {
	s.AgentCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetAnswerCount(v int64) *CloudMonitorTaskResponseBodyDataStatistic {
	s.AnswerCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetCallAgentCount(v int64) *CloudMonitorTaskResponseBodyDataStatistic {
	s.CallAgentCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetCalledCount(v int64) *CloudMonitorTaskResponseBodyDataStatistic {
	s.CalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetCtiLinkTaskProperty(v *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) *CloudMonitorTaskResponseBodyDataStatistic {
	s.CtiLinkTaskProperty = v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) SetUncalledCount(v int64) *CloudMonitorTaskResponseBodyDataStatistic {
	s.UncalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatistic) Validate() error {
	if s.CtiLinkTaskProperty != nil {
		if err := s.CtiLinkTaskProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty struct {
	// 外呼任务Id
	//
	// example:
	//
	// 12314
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 任务名称
	//
	// example:
	//
	// testTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 任务结束时间
	//
	// example:
	//
	// 2026-04-20 11:00:00
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// 任务开始时间
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态，0.初始 1.运行中 2.暂停 3.结束
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务状态描述
	//
	// example:
	//
	// 示例值
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetId() *int64 {
	return s.Id
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetName() *string {
	return s.Name
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetStatus() *int64 {
	return s.Status
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetId(v int64) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.Id = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetName(v string) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.Name = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetOverTime(v string) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.OverTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetStartTime(v string) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.StartTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetStatus(v int64) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.Status = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) SetStatusDescription(v string) *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty {
	s.StatusDescription = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataStatisticCtiLinkTaskProperty) Validate() error {
	return dara.Validate(s)
}

type CloudMonitorTaskResponseBodyDataTodayStatistic struct {
	// 座席数量
	//
	// example:
	//
	// 3
	AgentCount *int64 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// 应答数量
	//
	// example:
	//
	// 3
	AnswerCount *int64 `json:"AnswerCount,omitempty" xml:"AnswerCount,omitempty"`
	// 呼叫坐席数
	//
	// example:
	//
	// 3
	CallAgentCount *int64 `json:"CallAgentCount,omitempty" xml:"CallAgentCount,omitempty"`
	// 已呼叫数量
	//
	// example:
	//
	// 3
	CalledCount         *int64                                                             `json:"CalledCount,omitempty" xml:"CalledCount,omitempty"`
	CtiLinkTaskProperty *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty `json:"CtiLinkTaskProperty,omitempty" xml:"CtiLinkTaskProperty,omitempty" type:"Struct"`
	// 未呼叫数量
	//
	// example:
	//
	// 0
	UncalledCount *int64 `json:"UncalledCount,omitempty" xml:"UncalledCount,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataTodayStatistic) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataTodayStatistic) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetAgentCount() *int64 {
	return s.AgentCount
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetAnswerCount() *int64 {
	return s.AnswerCount
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetCallAgentCount() *int64 {
	return s.CallAgentCount
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetCalledCount() *int64 {
	return s.CalledCount
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetCtiLinkTaskProperty() *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	return s.CtiLinkTaskProperty
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) GetUncalledCount() *int64 {
	return s.UncalledCount
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetAgentCount(v int64) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.AgentCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetAnswerCount(v int64) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.AnswerCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetCallAgentCount(v int64) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.CallAgentCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetCalledCount(v int64) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.CalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetCtiLinkTaskProperty(v *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.CtiLinkTaskProperty = v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) SetUncalledCount(v int64) *CloudMonitorTaskResponseBodyDataTodayStatistic {
	s.UncalledCount = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatistic) Validate() error {
	if s.CtiLinkTaskProperty != nil {
		if err := s.CtiLinkTaskProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty struct {
	// 外呼任务Id
	//
	// example:
	//
	// 12314
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 任务名称
	//
	// example:
	//
	// testTask
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 任务结束时间
	//
	// example:
	//
	// 2026-04-20 11:00:00
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// 任务开始时间
	//
	// example:
	//
	// 2026-04-20 10:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 任务状态，0.初始 1.运行中 2.暂停 3.结束
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 任务状态描述
	//
	// example:
	//
	// 示例值示例值
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
}

func (s CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) String() string {
	return dara.Prettify(s)
}

func (s CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GoString() string {
	return s.String()
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetId() *int64 {
	return s.Id
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetName() *string {
	return s.Name
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetStatus() *string {
	return s.Status
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetId(v int64) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.Id = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetName(v string) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.Name = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetOverTime(v string) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.OverTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetStartTime(v string) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.StartTime = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetStatus(v string) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.Status = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) SetStatusDescription(v string) *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty {
	s.StatusDescription = &v
	return s
}

func (s *CloudMonitorTaskResponseBodyDataTodayStatisticCtiLinkTaskProperty) Validate() error {
	return dara.Validate(s)
}
