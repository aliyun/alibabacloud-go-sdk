// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentMonitorStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentMonitorStatisticsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentMonitorStatisticsResponseBody
	GetCode() *string
	SetData(v *CloudAgentMonitorStatisticsResponseBodyData) *CloudAgentMonitorStatisticsResponseBody
	GetData() *CloudAgentMonitorStatisticsResponseBodyData
	SetMessage(v string) *CloudAgentMonitorStatisticsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentMonitorStatisticsResponseBody
	GetRequestId() *string
}

type CloudAgentMonitorStatisticsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentMonitorStatisticsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9FF70B74-1B3C-44C1-ACDF-8DF962988F0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentMonitorStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentMonitorStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentMonitorStatisticsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentMonitorStatisticsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentMonitorStatisticsResponseBody) GetData() *CloudAgentMonitorStatisticsResponseBodyData {
	return s.Data
}

func (s *CloudAgentMonitorStatisticsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentMonitorStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentMonitorStatisticsResponseBody) SetAccessDeniedDetail(v string) *CloudAgentMonitorStatisticsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBody) SetCode(v string) *CloudAgentMonitorStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBody) SetData(v *CloudAgentMonitorStatisticsResponseBodyData) *CloudAgentMonitorStatisticsResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBody) SetMessage(v string) *CloudAgentMonitorStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBody) SetRequestId(v string) *CloudAgentMonitorStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentMonitorStatisticsResponseBodyData struct {
	// 座席统计项数组
	AgentStatistics []*CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics `json:"AgentStatistics,omitempty" xml:"AgentStatistics,omitempty" type:"Repeated"`
	// 日期，格式为：yyyy-MM-dd
	//
	// example:
	//
	// 2026-04-06
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// 总计
	//
	// example:
	//
	// 26
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CloudAgentMonitorStatisticsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentMonitorStatisticsResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) GetAgentStatistics() []*CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	return s.AgentStatistics
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) GetDate() *string {
	return s.Date
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) SetAgentStatistics(v []*CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) *CloudAgentMonitorStatisticsResponseBodyData {
	s.AgentStatistics = v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) SetDate(v string) *CloudAgentMonitorStatisticsResponseBodyData {
	s.Date = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) SetTotal(v string) *CloudAgentMonitorStatisticsResponseBodyData {
	s.Total = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyData) Validate() error {
	if s.AgentStatistics != nil {
		for _, item := range s.AgentStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics struct {
	// 座席工号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// status: 代表座席状态，状态每个值对应的含义: stateIdle: 置闲; statePause: 置忙; stateCalling: 呼叫中; stateInuse: 通话中; stateWrapup: 整理, loginTime: 代表座席登录时间, startTime: 代表开始通话的时间
	//
	// example:
	//
	// {\\"state\\":\\"stateIdle\\",\\"startTime\\":\\"1491462675\\"}
	CurrentState *string `json:"CurrentState,omitempty" xml:"CurrentState,omitempty"`
	// 外呼组编号
	//
	// example:
	//
	// WH123
	Gno *string `json:"Gno,omitempty" xml:"Gno,omitempty"`
	// 外呼组名称
	//
	// example:
	//
	// gpname
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 总空闲次数
	//
	// example:
	//
	// 5
	IdleCount *string `json:"Idle_count,omitempty" xml:"Idle_count,omitempty"`
	// 总空闲时长
	//
	// example:
	//
	// 54344
	IdleTime *string `json:"Idle_time,omitempty" xml:"Idle_time,omitempty"`
	// 最大空闲时长
	//
	// example:
	//
	// 52375
	MaxIdleTime *string `json:"Max_idle_time,omitempty" xml:"Max_idle_time,omitempty"`
	// 外呼通话最大时长
	//
	// example:
	//
	// 41
	MaxObBridgeTime *string `json:"Max_ob_bridge_time,omitempty" xml:"Max_ob_bridge_time,omitempty"`
	// 外呼呼叫最大时长
	//
	// example:
	//
	// 31
	MaxObCallingTime *string `json:"Max_ob_calling_time,omitempty" xml:"Max_ob_calling_time,omitempty"`
	// 外呼整理最大时长
	//
	// example:
	//
	// 30
	MaxObWrapupTime *string `json:"Max_ob_wrapup_time,omitempty" xml:"Max_ob_wrapup_time,omitempty"`
	// 最小空闲时长
	//
	// example:
	//
	// 9
	MinIdleTime *string `json:"Min_idle_time,omitempty" xml:"Min_idle_time,omitempty"`
	// 外呼通话最小时长
	//
	// example:
	//
	// 29
	MinObBridgeTime *string `json:"Min_ob_bridge_time,omitempty" xml:"Min_ob_bridge_time,omitempty"`
	// 外呼呼叫最小时长
	//
	// example:
	//
	// 8
	MinObCallingTime *string `json:"Min_ob_calling_time,omitempty" xml:"Min_ob_calling_time,omitempty"`
	// 外呼整理最小时长
	//
	// example:
	//
	// 30
	MinObWrapupTime *string `json:"Min_ob_wrapup_time,omitempty" xml:"Min_ob_wrapup_time,omitempty"`
	// 总外呼通话次数
	//
	// example:
	//
	// 4
	ObBridgeCount *string `json:"Ob_bridge_count,omitempty" xml:"Ob_bridge_count,omitempty"`
	// 外呼通话时长
	//
	// example:
	//
	// 144
	ObBridgeTime *string `json:"Ob_bridge_time,omitempty" xml:"Ob_bridge_time,omitempty"`
	// 总外呼呼叫次数
	//
	// example:
	//
	// 5
	ObCallingCount *string `json:"Ob_calling_count,omitempty" xml:"Ob_calling_count,omitempty"`
	// 外呼呼叫时长
	//
	// example:
	//
	// 67
	ObCallingTime *string `json:"Ob_calling_time,omitempty" xml:"Ob_calling_time,omitempty"`
	// 总外呼整理次数
	//
	// example:
	//
	// 4
	ObWrapupCount *string `json:"Ob_wrapup_count,omitempty" xml:"Ob_wrapup_count,omitempty"`
	// 外呼整理时长
	//
	// example:
	//
	// 120
	ObWrapupTime *string `json:"Ob_wrapup_time,omitempty" xml:"Ob_wrapup_time,omitempty"`
	// 预览外呼座席接听数
	//
	// example:
	//
	// 4
	PreviewObAgentAnsweredCount *string `json:"Preview_ob_agent_answered_count,omitempty" xml:"Preview_ob_agent_answered_count,omitempty"`
	// example:
	//
	// 4
	PreviewObBothAnsweredCount *string `json:"Preview_ob_both_answered_count,omitempty" xml:"Preview_ob_both_answered_count,omitempty"`
	// 双方接听最大通话时长
	//
	// example:
	//
	// 41
	PreviewObMaxBridgeTime *string `json:"Preview_ob_max_bridge_time,omitempty" xml:"Preview_ob_max_bridge_time,omitempty"`
	// 双方接听最小通话时长
	//
	// example:
	//
	// 29
	PreviewObMinBridgeTime *string `json:"Preview_ob_min_bridge_time,omitempty" xml:"Preview_ob_min_bridge_time,omitempty"`
	// 双方接听总通话时长
	//
	// example:
	//
	// 144
	PreviewObTotalBridgeTime *string `json:"Preview_ob_total_bridge_time,omitempty" xml:"Preview_ob_total_bridge_time,omitempty"`
	// 预览外呼数
	//
	// example:
	//
	// 5
	PreviewObTotalCount *string `json:"Preview_ob_total_count,omitempty" xml:"Preview_ob_total_count,omitempty"`
	// 预览外呼有效通话次数
	//
	// example:
	//
	// 4
	PreviewObValidCalls *string `json:"Preview_ob_valid_calls,omitempty" xml:"Preview_ob_valid_calls,omitempty"`
	// 预览外呼有效通话总时长
	//
	// example:
	//
	// 144
	PreviewObValidTotalBridgeTime *string `json:"Preview_ob_valid_total_bridge_time,omitempty" xml:"Preview_ob_valid_total_bridge_time,omitempty"`
	// 队列名
	//
	// example:
	//
	// 示例值
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GoString() string {
	return s.String()
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetCno() *string {
	return s.Cno
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetCurrentState() *string {
	return s.CurrentState
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetGno() *string {
	return s.Gno
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetGroupName() *string {
	return s.GroupName
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetIdleCount() *string {
	return s.IdleCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetIdleTime() *string {
	return s.IdleTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMaxIdleTime() *string {
	return s.MaxIdleTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMaxObBridgeTime() *string {
	return s.MaxObBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMaxObCallingTime() *string {
	return s.MaxObCallingTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMaxObWrapupTime() *string {
	return s.MaxObWrapupTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMinIdleTime() *string {
	return s.MinIdleTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMinObBridgeTime() *string {
	return s.MinObBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMinObCallingTime() *string {
	return s.MinObCallingTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetMinObWrapupTime() *string {
	return s.MinObWrapupTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObBridgeCount() *string {
	return s.ObBridgeCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObBridgeTime() *string {
	return s.ObBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObCallingCount() *string {
	return s.ObCallingCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObCallingTime() *string {
	return s.ObCallingTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObWrapupCount() *string {
	return s.ObWrapupCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetObWrapupTime() *string {
	return s.ObWrapupTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObAgentAnsweredCount() *string {
	return s.PreviewObAgentAnsweredCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObBothAnsweredCount() *string {
	return s.PreviewObBothAnsweredCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObMaxBridgeTime() *string {
	return s.PreviewObMaxBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObMinBridgeTime() *string {
	return s.PreviewObMinBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObTotalBridgeTime() *string {
	return s.PreviewObTotalBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObTotalCount() *string {
	return s.PreviewObTotalCount
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObValidCalls() *string {
	return s.PreviewObValidCalls
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetPreviewObValidTotalBridgeTime() *string {
	return s.PreviewObValidTotalBridgeTime
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) GetQueueName() *string {
	return s.QueueName
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetCno(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.Cno = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetCurrentState(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.CurrentState = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetGno(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.Gno = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetGroupName(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.GroupName = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetIdleCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.IdleCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetIdleTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.IdleTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMaxIdleTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MaxIdleTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMaxObBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MaxObBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMaxObCallingTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MaxObCallingTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMaxObWrapupTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MaxObWrapupTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMinIdleTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MinIdleTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMinObBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MinObBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMinObCallingTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MinObCallingTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetMinObWrapupTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.MinObWrapupTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObBridgeCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObBridgeCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObCallingCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObCallingCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObCallingTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObCallingTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObWrapupCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObWrapupCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetObWrapupTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.ObWrapupTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObAgentAnsweredCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObAgentAnsweredCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObBothAnsweredCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObBothAnsweredCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObMaxBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObMaxBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObMinBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObMinBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObTotalBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObTotalBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObTotalCount(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObTotalCount = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObValidCalls(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObValidCalls = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetPreviewObValidTotalBridgeTime(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.PreviewObValidTotalBridgeTime = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) SetQueueName(v string) *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics {
	s.QueueName = &v
	return s
}

func (s *CloudAgentMonitorStatisticsResponseBodyDataAgentStatistics) Validate() error {
	return dara.Validate(s)
}
