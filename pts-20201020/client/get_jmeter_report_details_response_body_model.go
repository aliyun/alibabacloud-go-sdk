// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterReportDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetJMeterReportDetailsResponseBody
	GetCode() *string
	SetCodeKey(v string) *GetJMeterReportDetailsResponseBody
	GetCodeKey() *string
	SetDocumentUrl(v string) *GetJMeterReportDetailsResponseBody
	GetDocumentUrl() *string
	SetDynamicCtx(v string) *GetJMeterReportDetailsResponseBody
	GetDynamicCtx() *string
	SetHttpStatusCode(v int32) *GetJMeterReportDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetJMeterReportDetailsResponseBody
	GetMessage() *string
	SetReportOverView(v *GetJMeterReportDetailsResponseBodyReportOverView) *GetJMeterReportDetailsResponseBody
	GetReportOverView() *GetJMeterReportDetailsResponseBodyReportOverView
	SetRequestId(v string) *GetJMeterReportDetailsResponseBody
	GetRequestId() *string
	SetSamplerMetricsList(v []*GetJMeterReportDetailsResponseBodySamplerMetricsList) *GetJMeterReportDetailsResponseBody
	GetSamplerMetricsList() []*GetJMeterReportDetailsResponseBodySamplerMetricsList
	SetSceneMetrics(v *GetJMeterReportDetailsResponseBodySceneMetrics) *GetJMeterReportDetailsResponseBody
	GetSceneMetrics() *GetJMeterReportDetailsResponseBodySceneMetrics
	SetSuccess(v bool) *GetJMeterReportDetailsResponseBody
	GetSuccess() *bool
}

type GetJMeterReportDetailsResponseBody struct {
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The code key that corresponds to the key in Medusa. If no code key is available, or if the content corresponding to the code key fails to be obtained or is empty, the returned message is displayed as the default information.
	CodeKey *string `json:"CodeKey,omitempty" xml:"CodeKey,omitempty"`
	// The URL used to access the document.
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The returned dynamic contents that are separated by the && operator.
	DynamicCtx *string `json:"DynamicCtx,omitempty" xml:"DynamicCtx,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The details of the report.
	ReportOverView *GetJMeterReportDetailsResponseBodyReportOverView `json:"ReportOverView,omitempty" xml:"ReportOverView,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The dimensions of APIs.
	SamplerMetricsList []*GetJMeterReportDetailsResponseBodySamplerMetricsList `json:"SamplerMetricsList,omitempty" xml:"SamplerMetricsList,omitempty" type:"Repeated"`
	// The dimensions of the whole scenario.
	SceneMetrics *GetJMeterReportDetailsResponseBodySceneMetrics `json:"SceneMetrics,omitempty" xml:"SceneMetrics,omitempty" type:"Struct"`
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
}

func (s GetJMeterReportDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJMeterReportDetailsResponseBody) GetCodeKey() *string {
	return s.CodeKey
}

func (s *GetJMeterReportDetailsResponseBody) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *GetJMeterReportDetailsResponseBody) GetDynamicCtx() *string {
	return s.DynamicCtx
}

func (s *GetJMeterReportDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetJMeterReportDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJMeterReportDetailsResponseBody) GetReportOverView() *GetJMeterReportDetailsResponseBodyReportOverView {
	return s.ReportOverView
}

func (s *GetJMeterReportDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJMeterReportDetailsResponseBody) GetSamplerMetricsList() []*GetJMeterReportDetailsResponseBodySamplerMetricsList {
	return s.SamplerMetricsList
}

func (s *GetJMeterReportDetailsResponseBody) GetSceneMetrics() *GetJMeterReportDetailsResponseBodySceneMetrics {
	return s.SceneMetrics
}

func (s *GetJMeterReportDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterReportDetailsResponseBody) SetCode(v string) *GetJMeterReportDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetCodeKey(v string) *GetJMeterReportDetailsResponseBody {
	s.CodeKey = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetDocumentUrl(v string) *GetJMeterReportDetailsResponseBody {
	s.DocumentUrl = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetDynamicCtx(v string) *GetJMeterReportDetailsResponseBody {
	s.DynamicCtx = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetHttpStatusCode(v int32) *GetJMeterReportDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetMessage(v string) *GetJMeterReportDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetReportOverView(v *GetJMeterReportDetailsResponseBodyReportOverView) *GetJMeterReportDetailsResponseBody {
	s.ReportOverView = v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetRequestId(v string) *GetJMeterReportDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetSamplerMetricsList(v []*GetJMeterReportDetailsResponseBodySamplerMetricsList) *GetJMeterReportDetailsResponseBody {
	s.SamplerMetricsList = v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetSceneMetrics(v *GetJMeterReportDetailsResponseBodySceneMetrics) *GetJMeterReportDetailsResponseBody {
	s.SceneMetrics = v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) SetSuccess(v bool) *GetJMeterReportDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBody) Validate() error {
	if s.ReportOverView != nil {
		if err := s.ReportOverView.Validate(); err != nil {
			return err
		}
	}
	if s.SamplerMetricsList != nil {
		for _, item := range s.SamplerMetricsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneMetrics != nil {
		if err := s.SceneMetrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJMeterReportDetailsResponseBodyReportOverView struct {
	// The number of used engines.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The end of the queried time range.
	//
	// example:
	//
	// 2023-05-03 10:45:11
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The report ID.
	//
	// example:
	//
	// GHB56VD
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The report name.
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The beginning of the queried time range.
	//
	// example:
	//
	// 2023-05-03 10:35:11
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The consumed Virtual User Minutes (VUM).
	//
	// example:
	//
	// 4452
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetJMeterReportDetailsResponseBodyReportOverView) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsResponseBodyReportOverView) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetReportName() *string {
	return s.ReportName
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) GetVum() *int64 {
	return s.Vum
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetAgentCount(v int32) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.AgentCount = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetEndTime(v string) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.EndTime = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetReportId(v string) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.ReportId = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetReportName(v string) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.ReportName = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetStartTime(v string) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.StartTime = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) SetVum(v int64) *GetJMeterReportDetailsResponseBodyReportOverView {
	s.Vum = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodyReportOverView) Validate() error {
	return dara.Validate(s)
}

type GetJMeterReportDetailsResponseBodySamplerMetricsList struct {
	// The total number of requests.
	//
	// example:
	//
	// 731
	AllCount *int64 `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	// The API name.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The average RT. Unit: milliseconds.
	//
	// example:
	//
	// 44.2
	AvgRt *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average TPS.
	//
	// example:
	//
	// 12
	AvgTps *float64 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// The request failure rate.
	//
	// example:
	//
	// 10
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// The maximum RT. Unit: milliseconds.
	//
	// example:
	//
	// 78
	MaxRt *float64 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// The minimum RT. Unit: milliseconds.
	//
	// example:
	//
	// 11
	MinRt *float64 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	// The 75th percentile of RT. Unit: milliseconds.
	//
	// example:
	//
	// 22.4
	Seg75Rt *float64 `json:"Seg75Rt,omitempty" xml:"Seg75Rt,omitempty"`
	// The 90th percentile of RT. Unit: milliseconds.
	//
	// example:
	//
	// 65
	Seg90Rt *float64 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// The 99th percentile of RT. Unit: milliseconds.
	//
	// example:
	//
	// 77
	Seg99Rt *float64 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// The request success rate. The parameter value must be a non-negative number less than or equal to 100.
	//
	// example:
	//
	// 100
	SuccessRateReq *float64 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetJMeterReportDetailsResponseBodySamplerMetricsList) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsResponseBodySamplerMetricsList) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetAllCount() *int64 {
	return s.AllCount
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetApiName() *string {
	return s.ApiName
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetAvgTps() *float64 {
	return s.AvgTps
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetMaxRt() *float64 {
	return s.MaxRt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetMinRt() *float64 {
	return s.MinRt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetSeg75Rt() *float64 {
	return s.Seg75Rt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetSeg90Rt() *float64 {
	return s.Seg90Rt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetSeg99Rt() *float64 {
	return s.Seg99Rt
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) GetSuccessRateReq() *float64 {
	return s.SuccessRateReq
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetAllCount(v int64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.AllCount = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetApiName(v string) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.ApiName = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetAvgRt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.AvgRt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetAvgTps(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.AvgTps = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetFailCountReq(v int64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.FailCountReq = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetMaxRt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.MaxRt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetMinRt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.MinRt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetSeg75Rt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.Seg75Rt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetSeg90Rt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.Seg90Rt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetSeg99Rt(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.Seg99Rt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) SetSuccessRateReq(v float64) *GetJMeterReportDetailsResponseBodySamplerMetricsList {
	s.SuccessRateReq = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySamplerMetricsList) Validate() error {
	return dara.Validate(s)
}

type GetJMeterReportDetailsResponseBodySceneMetrics struct {
	// The total number of requests.
	//
	// example:
	//
	// 717
	AllCount *int64 `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	// The average response time (RT). Unit: milliseconds.
	//
	// example:
	//
	// 23
	AvgRt *float64 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average transactions per second (TPS).
	//
	// example:
	//
	// 78
	AvgTps *float64 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// The request failure rate.
	//
	// example:
	//
	// 34
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// The 90th percentile of RT. Unit: milliseconds.
	//
	// example:
	//
	// 35
	Seg90Rt *float64 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// The 99th percentile of RT. Unit: milliseconds.
	//
	// example:
	//
	// 56
	Seg99Rt *float64 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// The request success rate. The parameter value must be a non-negative number less than or equal to 100.
	//
	// example:
	//
	// 0.99
	SuccessRateReq *float64 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetJMeterReportDetailsResponseBodySceneMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterReportDetailsResponseBodySceneMetrics) GoString() string {
	return s.String()
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetAllCount() *int64 {
	return s.AllCount
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetAvgRt() *float64 {
	return s.AvgRt
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetAvgTps() *float64 {
	return s.AvgTps
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetSeg90Rt() *float64 {
	return s.Seg90Rt
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetSeg99Rt() *float64 {
	return s.Seg99Rt
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) GetSuccessRateReq() *float64 {
	return s.SuccessRateReq
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetAllCount(v int64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.AllCount = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetAvgRt(v float64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.AvgRt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetAvgTps(v float64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.AvgTps = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetFailCountReq(v int64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.FailCountReq = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetSeg90Rt(v float64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.Seg90Rt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetSeg99Rt(v float64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.Seg99Rt = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) SetSuccessRateReq(v float64) *GetJMeterReportDetailsResponseBodySceneMetrics {
	s.SuccessRateReq = &v
	return s
}

func (s *GetJMeterReportDetailsResponseBodySceneMetrics) Validate() error {
	return dara.Validate(s)
}
