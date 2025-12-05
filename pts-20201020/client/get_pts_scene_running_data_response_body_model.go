// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentLocation(v []*GetPtsSceneRunningDataResponseBodyAgentLocation) *GetPtsSceneRunningDataResponseBody
	GetAgentLocation() []*GetPtsSceneRunningDataResponseBodyAgentLocation
	SetAliveAgents(v int32) *GetPtsSceneRunningDataResponseBody
	GetAliveAgents() *int32
	SetAverageRt(v int64) *GetPtsSceneRunningDataResponseBody
	GetAverageRt() *int64
	SetBeginTime(v int64) *GetPtsSceneRunningDataResponseBody
	GetBeginTime() *int64
	SetChainMonitorDataList(v []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList) *GetPtsSceneRunningDataResponseBody
	GetChainMonitorDataList() []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList
	SetCode(v string) *GetPtsSceneRunningDataResponseBody
	GetCode() *string
	SetConcurrency(v int32) *GetPtsSceneRunningDataResponseBody
	GetConcurrency() *int32
	SetConcurrencyLimit(v int32) *GetPtsSceneRunningDataResponseBody
	GetConcurrencyLimit() *int32
	SetFailedBusinessCount(v int64) *GetPtsSceneRunningDataResponseBody
	GetFailedBusinessCount() *int64
	SetFailedRequestCount(v int64) *GetPtsSceneRunningDataResponseBody
	GetFailedRequestCount() *int64
	SetHasReport(v bool) *GetPtsSceneRunningDataResponseBody
	GetHasReport() *bool
	SetHttpStatusCode(v int32) *GetPtsSceneRunningDataResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsSceneRunningDataResponseBody
	GetMessage() *string
	SetRequestBps(v string) *GetPtsSceneRunningDataResponseBody
	GetRequestBps() *string
	SetRequestId(v string) *GetPtsSceneRunningDataResponseBody
	GetRequestId() *string
	SetResponseBps(v string) *GetPtsSceneRunningDataResponseBody
	GetResponseBps() *string
	SetSeg90Rt(v int64) *GetPtsSceneRunningDataResponseBody
	GetSeg90Rt() *int64
	SetStatus(v int32) *GetPtsSceneRunningDataResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *GetPtsSceneRunningDataResponseBody
	GetSuccess() *bool
	SetTotalAgents(v int32) *GetPtsSceneRunningDataResponseBody
	GetTotalAgents() *int32
	SetTotalRealQps(v int32) *GetPtsSceneRunningDataResponseBody
	GetTotalRealQps() *int32
	SetTotalRequestCount(v int64) *GetPtsSceneRunningDataResponseBody
	GetTotalRequestCount() *int64
	SetTpsLimit(v int32) *GetPtsSceneRunningDataResponseBody
	GetTpsLimit() *int32
	SetVum(v int64) *GetPtsSceneRunningDataResponseBody
	GetVum() *int64
}

type GetPtsSceneRunningDataResponseBody struct {
	// The location information of stress testers.
	AgentLocation []*GetPtsSceneRunningDataResponseBodyAgentLocation `json:"AgentLocation,omitempty" xml:"AgentLocation,omitempty" type:"Repeated"`
	// The number of healthy engines.
	//
	// example:
	//
	// 10
	AliveAgents *int32 `json:"AliveAgents,omitempty" xml:"AliveAgents,omitempty"`
	// The average RT.
	//
	// example:
	//
	// 45
	AverageRt *int64 `json:"AverageRt,omitempty" xml:"AverageRt,omitempty"`
	// The start time of the stress testing that is displayed as a timestamp. Unit: ms.
	//
	// example:
	//
	// 1651895518339
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The stress testing details of the GetPtsSceneRunningData operation.
	ChainMonitorDataList []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList `json:"ChainMonitorDataList,omitempty" xml:"ChainMonitorDataList,omitempty" type:"Repeated"`
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total concurrency.
	//
	// example:
	//
	// 10
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The maximum concurrency.
	//
	// example:
	//
	// 20
	ConcurrencyLimit *int32 `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	// The total number of failed businesses.
	//
	// example:
	//
	// 78
	FailedBusinessCount *int64 `json:"FailedBusinessCount,omitempty" xml:"FailedBusinessCount,omitempty"`
	// The number of failed requests.
	//
	// example:
	//
	// 90
	FailedRequestCount *int64 `json:"FailedRequestCount,omitempty" xml:"FailedRequestCount,omitempty"`
	// Indicates whether a report is generated.
	//
	// example:
	//
	// false
	HasReport *bool `json:"HasReport,omitempty" xml:"HasReport,omitempty"`
	// The HTTP status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// no message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The size of the request body.
	//
	// example:
	//
	// 89kb
	RequestBps *string `json:"RequestBps,omitempty" xml:"RequestBps,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC4E3177-6745-4925-B423-4E89VV34221A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The size of the response body.
	//
	// example:
	//
	// 8kb
	ResponseBps *string `json:"ResponseBps,omitempty" xml:"ResponseBps,omitempty"`
	// The 90th percentile of reaction time (RT).
	//
	// example:
	//
	// 45
	Seg90Rt *int64 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// The scenario status. The default parameter value is 7.
	//
	// example:
	//
	// 6
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of stress testers.
	//
	// example:
	//
	// 10
	TotalAgents *int32 `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	// The total number of queries per second (QPS).
	//
	// example:
	//
	// 10
	TotalRealQps *int32 `json:"TotalRealQps,omitempty" xml:"TotalRealQps,omitempty"`
	// The total number of requests.
	//
	// example:
	//
	// 8900
	TotalRequestCount *int64 `json:"TotalRequestCount,omitempty" xml:"TotalRequestCount,omitempty"`
	// The maximum transactions per second (TPS).
	//
	// example:
	//
	// 80
	TpsLimit *int32 `json:"TpsLimit,omitempty" xml:"TpsLimit,omitempty"`
	// The consumed Virtual User Minutes (VUM).
	//
	// example:
	//
	// 100
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponseBody) GetAgentLocation() []*GetPtsSceneRunningDataResponseBodyAgentLocation {
	return s.AgentLocation
}

func (s *GetPtsSceneRunningDataResponseBody) GetAliveAgents() *int32 {
	return s.AliveAgents
}

func (s *GetPtsSceneRunningDataResponseBody) GetAverageRt() *int64 {
	return s.AverageRt
}

func (s *GetPtsSceneRunningDataResponseBody) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetPtsSceneRunningDataResponseBody) GetChainMonitorDataList() []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	return s.ChainMonitorDataList
}

func (s *GetPtsSceneRunningDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsSceneRunningDataResponseBody) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetPtsSceneRunningDataResponseBody) GetConcurrencyLimit() *int32 {
	return s.ConcurrencyLimit
}

func (s *GetPtsSceneRunningDataResponseBody) GetFailedBusinessCount() *int64 {
	return s.FailedBusinessCount
}

func (s *GetPtsSceneRunningDataResponseBody) GetFailedRequestCount() *int64 {
	return s.FailedRequestCount
}

func (s *GetPtsSceneRunningDataResponseBody) GetHasReport() *bool {
	return s.HasReport
}

func (s *GetPtsSceneRunningDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsSceneRunningDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsSceneRunningDataResponseBody) GetRequestBps() *string {
	return s.RequestBps
}

func (s *GetPtsSceneRunningDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsSceneRunningDataResponseBody) GetResponseBps() *string {
	return s.ResponseBps
}

func (s *GetPtsSceneRunningDataResponseBody) GetSeg90Rt() *int64 {
	return s.Seg90Rt
}

func (s *GetPtsSceneRunningDataResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *GetPtsSceneRunningDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsSceneRunningDataResponseBody) GetTotalAgents() *int32 {
	return s.TotalAgents
}

func (s *GetPtsSceneRunningDataResponseBody) GetTotalRealQps() *int32 {
	return s.TotalRealQps
}

func (s *GetPtsSceneRunningDataResponseBody) GetTotalRequestCount() *int64 {
	return s.TotalRequestCount
}

func (s *GetPtsSceneRunningDataResponseBody) GetTpsLimit() *int32 {
	return s.TpsLimit
}

func (s *GetPtsSceneRunningDataResponseBody) GetVum() *int64 {
	return s.Vum
}

func (s *GetPtsSceneRunningDataResponseBody) SetAgentLocation(v []*GetPtsSceneRunningDataResponseBodyAgentLocation) *GetPtsSceneRunningDataResponseBody {
	s.AgentLocation = v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetAliveAgents(v int32) *GetPtsSceneRunningDataResponseBody {
	s.AliveAgents = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetAverageRt(v int64) *GetPtsSceneRunningDataResponseBody {
	s.AverageRt = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetBeginTime(v int64) *GetPtsSceneRunningDataResponseBody {
	s.BeginTime = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetChainMonitorDataList(v []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList) *GetPtsSceneRunningDataResponseBody {
	s.ChainMonitorDataList = v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetCode(v string) *GetPtsSceneRunningDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetConcurrency(v int32) *GetPtsSceneRunningDataResponseBody {
	s.Concurrency = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetConcurrencyLimit(v int32) *GetPtsSceneRunningDataResponseBody {
	s.ConcurrencyLimit = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetFailedBusinessCount(v int64) *GetPtsSceneRunningDataResponseBody {
	s.FailedBusinessCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetFailedRequestCount(v int64) *GetPtsSceneRunningDataResponseBody {
	s.FailedRequestCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetHasReport(v bool) *GetPtsSceneRunningDataResponseBody {
	s.HasReport = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneRunningDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetMessage(v string) *GetPtsSceneRunningDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetRequestBps(v string) *GetPtsSceneRunningDataResponseBody {
	s.RequestBps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetRequestId(v string) *GetPtsSceneRunningDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetResponseBps(v string) *GetPtsSceneRunningDataResponseBody {
	s.ResponseBps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetSeg90Rt(v int64) *GetPtsSceneRunningDataResponseBody {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetStatus(v int32) *GetPtsSceneRunningDataResponseBody {
	s.Status = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetSuccess(v bool) *GetPtsSceneRunningDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetTotalAgents(v int32) *GetPtsSceneRunningDataResponseBody {
	s.TotalAgents = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetTotalRealQps(v int32) *GetPtsSceneRunningDataResponseBody {
	s.TotalRealQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetTotalRequestCount(v int64) *GetPtsSceneRunningDataResponseBody {
	s.TotalRequestCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetTpsLimit(v int32) *GetPtsSceneRunningDataResponseBody {
	s.TpsLimit = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) SetVum(v int64) *GetPtsSceneRunningDataResponseBody {
	s.Vum = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBody) Validate() error {
	if s.AgentLocation != nil {
		for _, item := range s.AgentLocation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChainMonitorDataList != nil {
		for _, item := range s.ChainMonitorDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsSceneRunningDataResponseBodyAgentLocation struct {
	// The number of stress testers.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The provider of the stress tester.
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The province in which the stress tester resides.
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The region in which the stress tester resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyAgentLocation) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyAgentLocation) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) GetCount() *int32 {
	return s.Count
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) GetIsp() *string {
	return s.Isp
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) GetProvince() *string {
	return s.Province
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) GetRegion() *string {
	return s.Region
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) SetCount(v int32) *GetPtsSceneRunningDataResponseBodyAgentLocation {
	s.Count = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) SetIsp(v string) *GetPtsSceneRunningDataResponseBodyAgentLocation {
	s.Isp = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) SetProvince(v string) *GetPtsSceneRunningDataResponseBodyAgentLocation {
	s.Province = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) SetRegion(v string) *GetPtsSceneRunningDataResponseBodyAgentLocation {
	s.Region = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyAgentLocation) Validate() error {
	return dara.Validate(s)
}

type GetPtsSceneRunningDataResponseBodyChainMonitorDataList struct {
	// The API ID.
	//
	// example:
	//
	// ANBDC8B
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The average RT.
	//
	// example:
	//
	// 46
	AverageRt *int32 `json:"AverageRt,omitempty" xml:"AverageRt,omitempty"`
	// The check point results.
	CheckPointResult *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult `json:"CheckPointResult,omitempty" xml:"CheckPointResult,omitempty" type:"Struct"`
	// The concurrency.
	//
	// example:
	//
	// 100
	Concurrency *float32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The RPS of successful and failed requests.
	//
	// example:
	//
	// 78
	ConfigQps *int32 `json:"ConfigQps,omitempty" xml:"ConfigQps,omitempty"`
	// The number of successful requests.
	//
	// example:
	//
	// 7890
	Count2XX *int64 `json:"Count2XX,omitempty" xml:"Count2XX,omitempty"`
	// The total number of failed requests.
	//
	// example:
	//
	// 456
	FailedCount *int64 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The RPS of failed requests.
	//
	// example:
	//
	// 15
	FailedQps *float32 `json:"FailedQps,omitempty" xml:"FailedQps,omitempty"`
	// The maximum RT.
	//
	// example:
	//
	// 56
	MaxRt *int32 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// The minimum RT.
	//
	// example:
	//
	// 16
	MinRt *int32 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	// The API ID.
	//
	// example:
	//
	// 78509
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The Requests Per Second (RPS) of successful requests.
	//
	// example:
	//
	// 78
	Qps2XX *float32 `json:"Qps2XX,omitempty" xml:"Qps2XX,omitempty"`
	// The actual RPS.
	//
	// example:
	//
	// 23
	RealQps *float32 `json:"RealQps,omitempty" xml:"RealQps,omitempty"`
	// The point in time at which the stress testing is performed.
	//
	// example:
	//
	// 1278908899
	TimePoint *int64 `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetApiId() *string {
	return s.ApiId
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetApiName() *string {
	return s.ApiName
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetAverageRt() *int32 {
	return s.AverageRt
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetCheckPointResult() *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult {
	return s.CheckPointResult
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetConcurrency() *float32 {
	return s.Concurrency
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetConfigQps() *int32 {
	return s.ConfigQps
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetCount2XX() *int64 {
	return s.Count2XX
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetFailedCount() *int64 {
	return s.FailedCount
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetFailedQps() *float32 {
	return s.FailedQps
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetMaxRt() *int32 {
	return s.MaxRt
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetMinRt() *int32 {
	return s.MinRt
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetQps2XX() *float32 {
	return s.Qps2XX
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetRealQps() *float32 {
	return s.RealQps
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GetTimePoint() *int64 {
	return s.TimePoint
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetApiId(v string) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.ApiId = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetApiName(v string) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.ApiName = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetAverageRt(v int32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.AverageRt = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetCheckPointResult(v *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.CheckPointResult = v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetConcurrency(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.Concurrency = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetConfigQps(v int32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.ConfigQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetCount2XX(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.Count2XX = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetFailedCount(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.FailedCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetFailedQps(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.FailedQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetMaxRt(v int32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.MaxRt = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetMinRt(v int32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.MinRt = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetNodeId(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.NodeId = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetQps2XX(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.Qps2XX = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetRealQps(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.RealQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) SetTimePoint(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataList {
	s.TimePoint = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataList) Validate() error {
	if s.CheckPointResult != nil {
		if err := s.CheckPointResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult struct {
	// The number of failed businesses.
	//
	// example:
	//
	// 1000
	FailedBusinessCount *int64 `json:"FailedBusinessCount,omitempty" xml:"FailedBusinessCount,omitempty"`
	// The RPS of failed businesses.
	//
	// example:
	//
	// 78
	FailedBusinessQps *float32 `json:"FailedBusinessQps,omitempty" xml:"FailedBusinessQps,omitempty"`
	// The number of successful businesses.
	//
	// example:
	//
	// 908
	SucceedBusinessCount *int64 `json:"SucceedBusinessCount,omitempty" xml:"SucceedBusinessCount,omitempty"`
	// The RPS of the successful businesses.
	//
	// example:
	//
	// 89
	SucceedBusinessQps *float32 `json:"SucceedBusinessQps,omitempty" xml:"SucceedBusinessQps,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GetFailedBusinessCount() *int64 {
	return s.FailedBusinessCount
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GetFailedBusinessQps() *float32 {
	return s.FailedBusinessQps
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GetSucceedBusinessCount() *int64 {
	return s.SucceedBusinessCount
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GetSucceedBusinessQps() *float32 {
	return s.SucceedBusinessQps
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) SetFailedBusinessCount(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult {
	s.FailedBusinessCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) SetFailedBusinessQps(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult {
	s.FailedBusinessQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) SetSucceedBusinessCount(v int64) *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult {
	s.SucceedBusinessCount = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) SetSucceedBusinessQps(v float32) *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult {
	s.SucceedBusinessQps = &v
	return s
}

func (s *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) Validate() error {
	return dara.Validate(s)
}
