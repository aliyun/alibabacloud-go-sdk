// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudOutboundPreviewObReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudOutboundPreviewObReportResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudOutboundPreviewObReportResponseBody
	GetCode() *string
	SetData(v *CloudOutboundPreviewObReportResponseBodyData) *CloudOutboundPreviewObReportResponseBody
	GetData() *CloudOutboundPreviewObReportResponseBodyData
	SetMessage(v string) *CloudOutboundPreviewObReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudOutboundPreviewObReportResponseBody
	GetRequestId() *string
}

type CloudOutboundPreviewObReportResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudOutboundPreviewObReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudOutboundPreviewObReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundPreviewObReportResponseBody) GoString() string {
	return s.String()
}

func (s *CloudOutboundPreviewObReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudOutboundPreviewObReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudOutboundPreviewObReportResponseBody) GetData() *CloudOutboundPreviewObReportResponseBodyData {
	return s.Data
}

func (s *CloudOutboundPreviewObReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudOutboundPreviewObReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudOutboundPreviewObReportResponseBody) SetAccessDeniedDetail(v string) *CloudOutboundPreviewObReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBody) SetCode(v string) *CloudOutboundPreviewObReportResponseBody {
	s.Code = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBody) SetData(v *CloudOutboundPreviewObReportResponseBodyData) *CloudOutboundPreviewObReportResponseBody {
	s.Data = v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBody) SetMessage(v string) *CloudOutboundPreviewObReportResponseBody {
	s.Message = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBody) SetRequestId(v string) *CloudOutboundPreviewObReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudOutboundPreviewObReportResponseBodyData struct {
	List []*CloudOutboundPreviewObReportResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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

func (s CloudOutboundPreviewObReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundPreviewObReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudOutboundPreviewObReportResponseBodyData) GetList() []*CloudOutboundPreviewObReportResponseBodyDataList {
	return s.List
}

func (s *CloudOutboundPreviewObReportResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *CloudOutboundPreviewObReportResponseBodyData) GetTotalCount() *string {
	return s.TotalCount
}

func (s *CloudOutboundPreviewObReportResponseBodyData) GetTotalPageCount() *string {
	return s.TotalPageCount
}

func (s *CloudOutboundPreviewObReportResponseBodyData) SetList(v []*CloudOutboundPreviewObReportResponseBodyDataList) *CloudOutboundPreviewObReportResponseBodyData {
	s.List = v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyData) SetPageSize(v string) *CloudOutboundPreviewObReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyData) SetTotalCount(v string) *CloudOutboundPreviewObReportResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyData) SetTotalPageCount(v string) *CloudOutboundPreviewObReportResponseBodyData {
	s.TotalPageCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyData) Validate() error {
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

type CloudOutboundPreviewObReportResponseBodyDataList struct {
	// 座席接听数
	//
	// example:
	//
	// 3
	AgentAnsweredCount *string `json:"AgentAnsweredCount,omitempty" xml:"AgentAnsweredCount,omitempty"`
	// 座席姓名
	//
	// example:
	//
	// name3
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// 座席接听率
	//
	// example:
	//
	// 100%
	AgentRate *string `json:"AgentRate,omitempty" xml:"AgentRate,omitempty"`
	// 平均通话时长
	//
	// example:
	//
	// 00:00:12
	AvgBridgeTime *string `json:"AvgBridgeTime,omitempty" xml:"AvgBridgeTime,omitempty"`
	// AXB双方接听数
	//
	// example:
	//
	// 0
	AxbBothAnsweredCount *string `json:"AxbBothAnsweredCount,omitempty" xml:"AxbBothAnsweredCount,omitempty"`
	// AXB双方接听率 = AXB双方接听数 / AXB外呼总数
	//
	// example:
	//
	// 0%
	AxbBothAnsweredRate *string `json:"AxbBothAnsweredRate,omitempty" xml:"AxbBothAnsweredRate,omitempty"`
	// AXB外呼双方接听分钟数
	//
	// example:
	//
	// 0
	AxbBothAnsweredTime *string `json:"AxbBothAnsweredTime,omitempty" xml:"AxbBothAnsweredTime,omitempty"`
	// AXB外呼总数
	//
	// example:
	//
	// 0
	AxbObTotalCount *string `json:"AxbObTotalCount,omitempty" xml:"AxbObTotalCount,omitempty"`
	// 双方接听数
	//
	// example:
	//
	// 3
	BothAnsweredCount *string `json:"BothAnsweredCount,omitempty" xml:"BothAnsweredCount,omitempty"`
	// 双方接听时长
	//
	// example:
	//
	// 2
	BothAnsweredTime *string `json:"BothAnsweredTime,omitempty" xml:"BothAnsweredTime,omitempty"`
	// 呼叫接通率
	//
	// example:
	//
	// 100%
	CallBridgedRate *string `json:"CallBridgedRate,omitempty" xml:"CallBridgedRate,omitempty"`
	// 发起呼叫次数
	//
	// example:
	//
	// 3
	CallTotalCount *string `json:"CallTotalCount,omitempty" xml:"CallTotalCount,omitempty"`
	// 座席工号
	//
	// example:
	//
	// 3008
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 客户接听数
	//
	// example:
	//
	// 3
	CustomerAnsweredCount *string `json:"CustomerAnsweredCount,omitempty" xml:"CustomerAnsweredCount,omitempty"`
	// 客户接听率
	//
	// example:
	//
	// 100%
	CustomerRate *string `json:"CustomerRate,omitempty" xml:"CustomerRate,omitempty"`
	// 统计周期
	//
	// example:
	//
	// 示例值
	DateTimeRange *string `json:"DateTimeRange,omitempty" xml:"DateTimeRange,omitempty"`
	// 企业Id
	//
	// example:
	//
	// 6000001
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 拦截次数
	//
	// example:
	//
	// 0
	InterceptionCount *string `json:"InterceptionCount,omitempty" xml:"InterceptionCount,omitempty"`
	// 拦截率
	//
	// example:
	//
	// 0%
	InterceptionRate *string `json:"InterceptionRate,omitempty" xml:"InterceptionRate,omitempty"`
	// 固话双方接听数
	//
	// example:
	//
	// 3
	LandlineBothAnsweredCount *string `json:"LandlineBothAnsweredCount,omitempty" xml:"LandlineBothAnsweredCount,omitempty"`
	// 固话双方接听率 = 固话双方接听数 / 固话外呼总数
	//
	// example:
	//
	// 100%
	LandlineBothAnsweredRate *string `json:"LandlineBothAnsweredRate,omitempty" xml:"LandlineBothAnsweredRate,omitempty"`
	// 固话双方接听分钟数
	//
	// example:
	//
	// 0
	LandlineBothAnsweredTime *string `json:"LandlineBothAnsweredTime,omitempty" xml:"LandlineBothAnsweredTime,omitempty"`
	// 固话外呼总数
	//
	// example:
	//
	// 0
	LandlineObTotalCount *string `json:"LandlineObTotalCount,omitempty" xml:"LandlineObTotalCount,omitempty"`
	// 最长通话时长
	//
	// example:
	//
	// 00:00:14
	MaxBridgeTime *string `json:"MaxBridgeTime,omitempty" xml:"MaxBridgeTime,omitempty"`
	// 最短通话时长
	//
	// example:
	//
	// 00:00:11
	MinBridgeTime *string `json:"MinBridgeTime,omitempty" xml:"MinBridgeTime,omitempty"`
	// 总通话时长
	//
	// example:
	//
	// 00:00:38
	TotalBridgeTime *string `json:"TotalBridgeTime,omitempty" xml:"TotalBridgeTime,omitempty"`
	// 外呼总数
	//
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 通话占比
	//
	// example:
	//
	// 0%
	VadRate *string `json:"VadRate,omitempty" xml:"VadRate,omitempty"`
	// 有效通话平均通话时长
	//
	// example:
	//
	// 00:00:12
	ValidAvgBridgeTime *string `json:"ValidAvgBridgeTime,omitempty" xml:"ValidAvgBridgeTime,omitempty"`
	// 有效通话次数
	//
	// example:
	//
	// 3
	ValidCalls *string `json:"ValidCalls,omitempty" xml:"ValidCalls,omitempty"`
	// 有效通话总通话时长
	//
	// example:
	//
	// 00:00:38
	ValidTotalBridgeTime *string `json:"ValidTotalBridgeTime,omitempty" xml:"ValidTotalBridgeTime,omitempty"`
}

func (s CloudOutboundPreviewObReportResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s CloudOutboundPreviewObReportResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAgentAnsweredCount() *string {
	return s.AgentAnsweredCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAgentRate() *string {
	return s.AgentRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAvgBridgeTime() *string {
	return s.AvgBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAxbBothAnsweredCount() *string {
	return s.AxbBothAnsweredCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAxbBothAnsweredRate() *string {
	return s.AxbBothAnsweredRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAxbBothAnsweredTime() *string {
	return s.AxbBothAnsweredTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetAxbObTotalCount() *string {
	return s.AxbObTotalCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetBothAnsweredCount() *string {
	return s.BothAnsweredCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetBothAnsweredTime() *string {
	return s.BothAnsweredTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetCallBridgedRate() *string {
	return s.CallBridgedRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetCallTotalCount() *string {
	return s.CallTotalCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetCno() *string {
	return s.Cno
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetCustomerAnsweredCount() *string {
	return s.CustomerAnsweredCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetCustomerRate() *string {
	return s.CustomerRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetDateTimeRange() *string {
	return s.DateTimeRange
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetInterceptionCount() *string {
	return s.InterceptionCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetInterceptionRate() *string {
	return s.InterceptionRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetLandlineBothAnsweredCount() *string {
	return s.LandlineBothAnsweredCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetLandlineBothAnsweredRate() *string {
	return s.LandlineBothAnsweredRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetLandlineBothAnsweredTime() *string {
	return s.LandlineBothAnsweredTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetLandlineObTotalCount() *string {
	return s.LandlineObTotalCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetMaxBridgeTime() *string {
	return s.MaxBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetMinBridgeTime() *string {
	return s.MinBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetTotalBridgeTime() *string {
	return s.TotalBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetTotalCount() *string {
	return s.TotalCount
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetVadRate() *string {
	return s.VadRate
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetValidAvgBridgeTime() *string {
	return s.ValidAvgBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetValidCalls() *string {
	return s.ValidCalls
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) GetValidTotalBridgeTime() *string {
	return s.ValidTotalBridgeTime
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAgentAnsweredCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AgentAnsweredCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAgentName(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAgentRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AgentRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAvgBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AvgBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAxbBothAnsweredCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AxbBothAnsweredCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAxbBothAnsweredRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AxbBothAnsweredRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAxbBothAnsweredTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AxbBothAnsweredTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetAxbObTotalCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.AxbObTotalCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetBothAnsweredCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.BothAnsweredCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetBothAnsweredTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.BothAnsweredTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetCallBridgedRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.CallBridgedRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetCallTotalCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.CallTotalCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetCno(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.Cno = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetCustomerAnsweredCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.CustomerAnsweredCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetCustomerRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.CustomerRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetDateTimeRange(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.DateTimeRange = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetEnterpriseId(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.EnterpriseId = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetInterceptionCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.InterceptionCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetInterceptionRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.InterceptionRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetLandlineBothAnsweredCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.LandlineBothAnsweredCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetLandlineBothAnsweredRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.LandlineBothAnsweredRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetLandlineBothAnsweredTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.LandlineBothAnsweredTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetLandlineObTotalCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.LandlineObTotalCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetMaxBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.MaxBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetMinBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.MinBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetTotalBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.TotalBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetTotalCount(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.TotalCount = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetVadRate(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.VadRate = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetValidAvgBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.ValidAvgBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetValidCalls(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.ValidCalls = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) SetValidTotalBridgeTime(v string) *CloudOutboundPreviewObReportResponseBodyDataList {
	s.ValidTotalBridgeTime = &v
	return s
}

func (s *CloudOutboundPreviewObReportResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
