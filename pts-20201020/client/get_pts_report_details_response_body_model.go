// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiMetricsList(v []*GetPtsReportDetailsResponseBodyApiMetricsList) *GetPtsReportDetailsResponseBody
	GetApiMetricsList() []*GetPtsReportDetailsResponseBodyApiMetricsList
	SetCode(v string) *GetPtsReportDetailsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetPtsReportDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetPtsReportDetailsResponseBody
	GetMessage() *string
	SetReportOverView(v *GetPtsReportDetailsResponseBodyReportOverView) *GetPtsReportDetailsResponseBody
	GetReportOverView() *GetPtsReportDetailsResponseBodyReportOverView
	SetRequestId(v string) *GetPtsReportDetailsResponseBody
	GetRequestId() *string
	SetSceneMetrics(v *GetPtsReportDetailsResponseBodySceneMetrics) *GetPtsReportDetailsResponseBody
	GetSceneMetrics() *GetPtsReportDetailsResponseBodySceneMetrics
	SetSceneSnapShot(v *GetPtsReportDetailsResponseBodySceneSnapShot) *GetPtsReportDetailsResponseBody
	GetSceneSnapShot() *GetPtsReportDetailsResponseBodySceneSnapShot
	SetSuccess(v bool) *GetPtsReportDetailsResponseBody
	GetSuccess() *bool
}

type GetPtsReportDetailsResponseBody struct {
	// The metrics for API operations in the PTS scenario
	ApiMetricsList []*GetPtsReportDetailsResponseBodyApiMetricsList `json:"ApiMetricsList,omitempty" xml:"ApiMetricsList,omitempty" type:"Repeated"`
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The summary of the report.
	ReportOverView *GetPtsReportDetailsResponseBodyReportOverView `json:"ReportOverView,omitempty" xml:"ReportOverView,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC4E3177-6745-4925-B423-4E89VV34221A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metrics of the scenario.
	SceneMetrics *GetPtsReportDetailsResponseBodySceneMetrics `json:"SceneMetrics,omitempty" xml:"SceneMetrics,omitempty" type:"Struct"`
	// The snapshot of the scenario.
	SceneSnapShot *GetPtsReportDetailsResponseBodySceneSnapShot `json:"SceneSnapShot,omitempty" xml:"SceneSnapShot,omitempty" type:"Struct"`
	// Indicates whether the operation is successful. Valid values:
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

func (s GetPtsReportDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBody) GetApiMetricsList() []*GetPtsReportDetailsResponseBodyApiMetricsList {
	return s.ApiMetricsList
}

func (s *GetPtsReportDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPtsReportDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetPtsReportDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPtsReportDetailsResponseBody) GetReportOverView() *GetPtsReportDetailsResponseBodyReportOverView {
	return s.ReportOverView
}

func (s *GetPtsReportDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPtsReportDetailsResponseBody) GetSceneMetrics() *GetPtsReportDetailsResponseBodySceneMetrics {
	return s.SceneMetrics
}

func (s *GetPtsReportDetailsResponseBody) GetSceneSnapShot() *GetPtsReportDetailsResponseBodySceneSnapShot {
	return s.SceneSnapShot
}

func (s *GetPtsReportDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPtsReportDetailsResponseBody) SetApiMetricsList(v []*GetPtsReportDetailsResponseBodyApiMetricsList) *GetPtsReportDetailsResponseBody {
	s.ApiMetricsList = v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetCode(v string) *GetPtsReportDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetHttpStatusCode(v int32) *GetPtsReportDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetMessage(v string) *GetPtsReportDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetReportOverView(v *GetPtsReportDetailsResponseBodyReportOverView) *GetPtsReportDetailsResponseBody {
	s.ReportOverView = v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetRequestId(v string) *GetPtsReportDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetSceneMetrics(v *GetPtsReportDetailsResponseBodySceneMetrics) *GetPtsReportDetailsResponseBody {
	s.SceneMetrics = v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetSceneSnapShot(v *GetPtsReportDetailsResponseBodySceneSnapShot) *GetPtsReportDetailsResponseBody {
	s.SceneSnapShot = v
	return s
}

func (s *GetPtsReportDetailsResponseBody) SetSuccess(v bool) *GetPtsReportDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetPtsReportDetailsResponseBody) Validate() error {
	if s.ApiMetricsList != nil {
		for _, item := range s.ApiMetricsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ReportOverView != nil {
		if err := s.ReportOverView.Validate(); err != nil {
			return err
		}
	}
	if s.SceneMetrics != nil {
		if err := s.SceneMetrics.Validate(); err != nil {
			return err
		}
	}
	if s.SceneSnapShot != nil {
		if err := s.SceneSnapShot.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodyApiMetricsList struct {
	// The total number of requests.
	//
	// example:
	//
	// 11872
	AllCount *int64 `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// Test-API
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The average response time. Unit: ms.
	//
	// example:
	//
	// 170.49
	AvgRt *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average TPS.
	//
	// example:
	//
	// 100.61
	AvgTps *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// The number of business-related failures. If a checkpoint is defined, a failure occurs when the conditions for the checkpoint are not satisfied.
	//
	// example:
	//
	// 0
	FailCountBiz *int64 `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	// The number of failed requests.
	//
	// example:
	//
	// 0
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// The maximum response time. Unit: ms.
	//
	// example:
	//
	// 600
	MaxRt *float32 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	// The minimum response time. Unit: ms.
	//
	// example:
	//
	// 162
	MinRt *float32 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	// The 50th percentile response time.
	//
	// example:
	//
	// 168
	Seg50Rt *float32 `json:"Seg50Rt,omitempty" xml:"Seg50Rt,omitempty"`
	// The 75th percentile response time.
	//
	// example:
	//
	// 169
	Seg75Rt *float32 `json:"Seg75Rt,omitempty" xml:"Seg75Rt,omitempty"`
	// The 90th percentile response time.
	//
	// example:
	//
	// 170
	Seg90Rt *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// The 99th percentile response time.
	//
	// example:
	//
	// 284
	Seg99Rt *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// The business success rate. The value is the ratio of the number of successful business to the total number of business.
	//
	// example:
	//
	// 0
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	// The request success rate. The value is the ratio of the number of successful requests to the total number of requests.
	//
	// example:
	//
	// 1
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsReportDetailsResponseBodyApiMetricsList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodyApiMetricsList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetAllCount() *int64 {
	return s.AllCount
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetApiName() *string {
	return s.ApiName
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetAvgRt() *float32 {
	return s.AvgRt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetAvgTps() *float32 {
	return s.AvgTps
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetFailCountBiz() *int64 {
	return s.FailCountBiz
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetMaxRt() *float32 {
	return s.MaxRt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetMinRt() *float32 {
	return s.MinRt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSeg50Rt() *float32 {
	return s.Seg50Rt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSeg75Rt() *float32 {
	return s.Seg75Rt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSeg90Rt() *float32 {
	return s.Seg90Rt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSeg99Rt() *float32 {
	return s.Seg99Rt
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSuccessRateBiz() *float32 {
	return s.SuccessRateBiz
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) GetSuccessRateReq() *float32 {
	return s.SuccessRateReq
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetAllCount(v int64) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.AllCount = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetApiName(v string) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.ApiName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetAvgRt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.AvgRt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetAvgTps(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.AvgTps = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetFailCountBiz(v int64) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetFailCountReq(v int64) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetMaxRt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.MaxRt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetMinRt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.MinRt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSeg50Rt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.Seg50Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSeg75Rt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.Seg75Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSeg90Rt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSeg99Rt(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSuccessRateBiz(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) SetSuccessRateReq(v float32) *GetPtsReportDetailsResponseBodyApiMetricsList {
	s.SuccessRateReq = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyApiMetricsList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodyReportOverView struct {
	// The number of load generators. Each load generator has an IP address.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The end time of the performance testing task.
	//
	// example:
	//
	// 2024-09-20 10:41:33
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the report.
	//
	// example:
	//
	// GHB56VD
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The name of the report.
	//
	// example:
	//
	// PTS-TEST
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The start time of the performance testing task.
	//
	// example:
	//
	// 2024-09-20 10:39:33
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The virtual user minutes (VUM).
	//
	// example:
	//
	// 1012
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsReportDetailsResponseBodyReportOverView) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodyReportOverView) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetReportId() *string {
	return s.ReportId
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetReportName() *string {
	return s.ReportName
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetStartTime() *string {
	return s.StartTime
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) GetVum() *int64 {
	return s.Vum
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetAgentCount(v int32) *GetPtsReportDetailsResponseBodyReportOverView {
	s.AgentCount = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetEndTime(v string) *GetPtsReportDetailsResponseBodyReportOverView {
	s.EndTime = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetReportId(v string) *GetPtsReportDetailsResponseBodyReportOverView {
	s.ReportId = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetReportName(v string) *GetPtsReportDetailsResponseBodyReportOverView {
	s.ReportName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetStartTime(v string) *GetPtsReportDetailsResponseBodyReportOverView {
	s.StartTime = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) SetVum(v int64) *GetPtsReportDetailsResponseBodyReportOverView {
	s.Vum = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodyReportOverView) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneMetrics struct {
	// The number of requests in the scenario.
	//
	// example:
	//
	// 11872
	AllCount *int64 `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	// The average response time in the scenario.
	//
	// example:
	//
	// 170.49
	AvgRt *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	// The average transactions per second (TPS) in the scenario.
	//
	// example:
	//
	// 100.61
	AvgTps *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	// The number of business failures in the scenario.
	//
	// example:
	//
	// 0
	FailCountBiz *int64 `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	// The number of failed requests in the scenario.
	//
	// example:
	//
	// 0
	FailCountReq *int64 `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	// The 90th percentile response time.
	//
	// example:
	//
	// 170
	Seg90Rt *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	// The 99th percentile response time.
	//
	// example:
	//
	// 284
	Seg99Rt *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	// The business success rate in the scenario.
	//
	// example:
	//
	// 0
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	// The request success rate in the scenario.
	//
	// example:
	//
	// 1
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneMetrics) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetAllCount() *int64 {
	return s.AllCount
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetAvgRt() *float32 {
	return s.AvgRt
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetAvgTps() *float32 {
	return s.AvgTps
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetFailCountBiz() *int64 {
	return s.FailCountBiz
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetFailCountReq() *int64 {
	return s.FailCountReq
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetSeg90Rt() *float32 {
	return s.Seg90Rt
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetSeg99Rt() *float32 {
	return s.Seg99Rt
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetSuccessRateBiz() *float32 {
	return s.SuccessRateBiz
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) GetSuccessRateReq() *float32 {
	return s.SuccessRateReq
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetAllCount(v int64) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.AllCount = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetAvgRt(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.AvgRt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetAvgTps(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.AvgTps = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetFailCountBiz(v int64) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetFailCountReq(v int64) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetSeg90Rt(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetSeg99Rt(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetSuccessRateBiz(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) SetSuccessRateReq(v float32) *GetPtsReportDetailsResponseBodySceneMetrics {
	s.SuccessRateReq = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneMetrics) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShot struct {
	// The advanced settings of the scenario.
	AdvanceSetting *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting `json:"AdvanceSetting,omitempty" xml:"AdvanceSetting,omitempty" type:"Struct"`
	// The time when the scenario was created.
	//
	// example:
	//
	// 2024-09-20 09:28:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The file used in the scenario.
	FileParameterList []*GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList `json:"FileParameterList,omitempty" xml:"FileParameterList,omitempty" type:"Repeated"`
	// The global parameters.
	GlobalParameterList []*GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList `json:"GlobalParameterList,omitempty" xml:"GlobalParameterList,omitempty" type:"Repeated"`
	// The load settings.
	LoadConfig *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig `json:"LoadConfig,omitempty" xml:"LoadConfig,omitempty" type:"Struct"`
	// The last modification time of the scenario.
	//
	// example:
	//
	// 2020-10-10 10:10:10
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The sessions.
	RelationList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	// The ID of the scenario.
	//
	// example:
	//
	// 7HBNS3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The name of the scenario.
	//
	// example:
	//
	// PTS-TEST
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The status of the scenario.
	//
	// example:
	//
	// STOPPED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShot) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShot) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetAdvanceSetting() *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting {
	return s.AdvanceSetting
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetFileParameterList() []*GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList {
	return s.FileParameterList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetGlobalParameterList() []*GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList {
	return s.GlobalParameterList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetLoadConfig() *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	return s.LoadConfig
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetRelationList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationList {
	return s.RelationList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetSceneName() *string {
	return s.SceneName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) GetStatus() *string {
	return s.Status
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetAdvanceSetting(v *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.AdvanceSetting = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetCreateTime(v string) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.CreateTime = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetFileParameterList(v []*GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.FileParameterList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetGlobalParameterList(v []*GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.GlobalParameterList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetLoadConfig(v *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.LoadConfig = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetModifiedTime(v string) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.ModifiedTime = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetRelationList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationList) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.RelationList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetSceneId(v string) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.SceneId = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetSceneName(v string) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.SceneName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) SetStatus(v string) *GetPtsReportDetailsResponseBodySceneSnapShot {
	s.Status = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShot) Validate() error {
	if s.AdvanceSetting != nil {
		if err := s.AdvanceSetting.Validate(); err != nil {
			return err
		}
	}
	if s.FileParameterList != nil {
		for _, item := range s.FileParameterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GlobalParameterList != nil {
		for _, item := range s.GlobalParameterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LoadConfig != nil {
		if err := s.LoadConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting struct {
	// The timeout period of the scenario.
	//
	// example:
	//
	// 5
	ConnectionTimeoutInSecond *int32 `json:"ConnectionTimeoutInSecond,omitempty" xml:"ConnectionTimeoutInSecond,omitempty"`
	// The domain name-IP address binding relationships.
	DomainBindingList []*GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList `json:"DomainBindingList,omitempty" xml:"DomainBindingList,omitempty" type:"Repeated"`
	// The log sampling rate.
	//
	// example:
	//
	// 1
	LogRate *int32 `json:"LogRate,omitempty" xml:"LogRate,omitempty"`
	// The custom success code.
	//
	// example:
	//
	// 429,404
	SuccessCode *string `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GetConnectionTimeoutInSecond() *int32 {
	return s.ConnectionTimeoutInSecond
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GetDomainBindingList() []*GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList {
	return s.DomainBindingList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GetLogRate() *int32 {
	return s.LogRate
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GetSuccessCode() *string {
	return s.SuccessCode
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) SetConnectionTimeoutInSecond(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting {
	s.ConnectionTimeoutInSecond = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) SetDomainBindingList(v []*GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting {
	s.DomainBindingList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) SetLogRate(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting {
	s.LogRate = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) SetSuccessCode(v string) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting {
	s.SuccessCode = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) Validate() error {
	if s.DomainBindingList != nil {
		for _, item := range s.DomainBindingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP addresses bound to the domain name.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) GetDomain() *string {
	return s.Domain
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) GetIps() []*string {
	return s.Ips
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) SetDomain(v string) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList {
	s.Domain = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) SetIps(v []*string) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList {
	s.Ips = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList struct {
	// The name of the file.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// https://test-bucket.oss-cn-shanghai.aliyuncs.com/test.csv
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) GetFileName() *string {
	return s.FileName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) GetFileOssAddress() *string {
	return s.FileOssAddress
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) SetFileName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList {
	s.FileName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) SetFileOssAddress(v string) *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList {
	s.FileOssAddress = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList struct {
	// The name of the parameter.
	//
	// example:
	//
	// username
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// user01
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) GetParamName() *string {
	return s.ParamName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) GetParamValue() *string {
	return s.ParamValue
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) SetParamName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList {
	s.ParamName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) SetParamValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList {
	s.ParamValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig struct {
	// The number of load generators.
	//
	// example:
	//
	// 1
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The API request load settings.
	ApiLoadConfigList []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList `json:"ApiLoadConfigList,omitempty" xml:"ApiLoadConfigList,omitempty" type:"Repeated"`
	// The concurrency and RPS limits in the scenario.
	Configuration *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	// The maximum running time. Unit: minutes.
	//
	// example:
	//
	// 2
	MaxRunningTime *int32 `json:"MaxRunningTime,omitempty" xml:"MaxRunningTime,omitempty"`
	// The settings of the session.
	RelationLoadConfigList []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList `json:"RelationLoadConfigList,omitempty" xml:"RelationLoadConfigList,omitempty" type:"Repeated"`
	// The load application mode.
	//
	// example:
	//
	// TPS
	TestMode *string `json:"TestMode,omitempty" xml:"TestMode,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetApiLoadConfigList() []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList {
	return s.ApiLoadConfigList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetConfiguration() *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration {
	return s.Configuration
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetMaxRunningTime() *int32 {
	return s.MaxRunningTime
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetRelationLoadConfigList() []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList {
	return s.RelationLoadConfigList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GetTestMode() *string {
	return s.TestMode
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetAgentCount(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.AgentCount = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetApiLoadConfigList(v []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.ApiLoadConfigList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetConfiguration(v *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.Configuration = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetMaxRunningTime(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.MaxRunningTime = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetRelationLoadConfigList(v []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.RelationLoadConfigList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) SetTestMode(v string) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig {
	s.TestMode = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) Validate() error {
	if s.ApiLoadConfigList != nil {
		for _, item := range s.ApiLoadConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.RelationLoadConfigList != nil {
		for _, item := range s.RelationLoadConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList struct {
	// The starting requests per second (RPS).
	//
	// example:
	//
	// 10
	RpsBegin *int32 `json:"RpsBegin,omitempty" xml:"RpsBegin,omitempty"`
	// The maximum RPS.
	//
	// example:
	//
	// 10
	RpsLimit *int32 `json:"RpsLimit,omitempty" xml:"RpsLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) GetRpsBegin() *int32 {
	return s.RpsBegin
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) GetRpsLimit() *int32 {
	return s.RpsLimit
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) SetRpsBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList {
	s.RpsBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) SetRpsLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList {
	s.RpsLimit = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration struct {
	// The starting number of concurrent virtual users in the scenario.
	//
	// example:
	//
	// 0
	AllConcurrencyBegin *int32 `json:"AllConcurrencyBegin,omitempty" xml:"AllConcurrencyBegin,omitempty"`
	// The maximum number of concurrent virtual users in the scenario.
	//
	// example:
	//
	// 500
	AllConcurrencyLimit *int32 `json:"AllConcurrencyLimit,omitempty" xml:"AllConcurrencyLimit,omitempty"`
	// The starting RPS in the scenario.
	//
	// example:
	//
	// 100
	AllRpsBegin *int32 `json:"AllRpsBegin,omitempty" xml:"AllRpsBegin,omitempty"`
	// The maximum RPS in the scenario.
	//
	// example:
	//
	// 1000
	AllRpsLimit *int32 `json:"AllRpsLimit,omitempty" xml:"AllRpsLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GetAllConcurrencyBegin() *int32 {
	return s.AllConcurrencyBegin
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GetAllConcurrencyLimit() *int32 {
	return s.AllConcurrencyLimit
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GetAllRpsBegin() *int32 {
	return s.AllRpsBegin
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GetAllRpsLimit() *int32 {
	return s.AllRpsLimit
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) SetAllConcurrencyBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration {
	s.AllConcurrencyBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) SetAllConcurrencyLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration {
	s.AllConcurrencyLimit = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) SetAllRpsBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration {
	s.AllRpsBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) SetAllRpsLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration {
	s.AllRpsLimit = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList struct {
	// The starting number of concurrent virtual users.
	//
	// example:
	//
	// 10
	ConcurrencyBegin *int32 `json:"ConcurrencyBegin,omitempty" xml:"ConcurrencyBegin,omitempty"`
	// The maximum number of concurrent virtual users.
	//
	// example:
	//
	// 20
	ConcurrencyLimit *int32 `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) GetConcurrencyBegin() *int32 {
	return s.ConcurrencyBegin
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) GetConcurrencyLimit() *int32 {
	return s.ConcurrencyLimit
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) SetConcurrencyBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList {
	s.ConcurrencyBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) SetConcurrencyLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList {
	s.ConcurrencyLimit = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationList struct {
	// The settings of the API operation.
	ApiList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	// The file parameters used by the session.
	FileParameterExplainList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList `json:"FileParameterExplainList,omitempty" xml:"FileParameterExplainList,omitempty" type:"Repeated"`
	// The ID of the session.
	//
	// example:
	//
	// HGBN4D
	RelationId *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	// The name of the session.
	//
	// example:
	//
	// Test-session-1
	RelationName *string `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GetApiList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	return s.ApiList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GetFileParameterExplainList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList {
	return s.FileParameterExplainList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GetRelationId() *string {
	return s.RelationId
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GetRelationName() *string {
	return s.RelationName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) SetApiList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) *GetPtsReportDetailsResponseBodySceneSnapShotRelationList {
	s.ApiList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) SetFileParameterExplainList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) *GetPtsReportDetailsResponseBodySceneSnapShotRelationList {
	s.FileParameterExplainList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) SetRelationId(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationList {
	s.RelationId = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) SetRelationName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationList {
	s.RelationName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationList) Validate() error {
	if s.ApiList != nil {
		for _, item := range s.ApiList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FileParameterExplainList != nil {
		for _, item := range s.FileParameterExplainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList struct {
	// The ID of the API operation.
	//
	// example:
	//
	// MNB45
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API operation.
	//
	// example:
	//
	// Test-API
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The request body.
	Body *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The checkpoints of the API operation.
	CheckPointList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList `json:"CheckPointList,omitempty" xml:"CheckPointList,omitempty" type:"Repeated"`
	// The export parameters.
	ExportList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList `json:"ExportList,omitempty" xml:"ExportList,omitempty" type:"Repeated"`
	// The headers.
	HeaderList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Repeated"`
	// The method of the request.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The number of redirections.
	//
	// example:
	//
	// 5
	RedirectCountLimit *int32 `json:"RedirectCountLimit,omitempty" xml:"RedirectCountLimit,omitempty"`
	// The timeout period.
	//
	// example:
	//
	// 5
	TimeoutInSecond *int32 `json:"TimeoutInSecond,omitempty" xml:"TimeoutInSecond,omitempty"`
	// The URL to which the API request is sent.
	//
	// example:
	//
	// http://www.example.com/
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetApiId() *string {
	return s.ApiId
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetApiName() *string {
	return s.ApiName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetBody() *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody {
	return s.Body
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetCheckPointList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList {
	return s.CheckPointList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetExportList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList {
	return s.ExportList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetHeaderList() []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList {
	return s.HeaderList
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetMethod() *string {
	return s.Method
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetRedirectCountLimit() *int32 {
	return s.RedirectCountLimit
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetTimeoutInSecond() *int32 {
	return s.TimeoutInSecond
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GetUrl() *string {
	return s.Url
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetApiId(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.ApiId = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetApiName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.ApiName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetBody(v *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.Body = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetCheckPointList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.CheckPointList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetExportList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.ExportList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetHeaderList(v []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.HeaderList = v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetMethod(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.Method = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetRedirectCountLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.RedirectCountLimit = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetTimeoutInSecond(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.TimeoutInSecond = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) SetUrl(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList {
	s.Url = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.CheckPointList != nil {
		for _, item := range s.CheckPointList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExportList != nil {
		for _, item := range s.ExportList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HeaderList != nil {
		for _, item := range s.HeaderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody struct {
	// The content of the request body.
	//
	// example:
	//
	// {key:value}
	BodyValue *string `json:"BodyValue,omitempty" xml:"BodyValue,omitempty"`
	// The type of the request body.
	//
	// example:
	//
	// application/x-www-form-urlencoded
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) GetBodyValue() *string {
	return s.BodyValue
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) GetContentType() *string {
	return s.ContentType
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) SetBodyValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody {
	s.BodyValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) SetContentType(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody {
	s.ContentType = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList struct {
	// The checked item.
	//
	// example:
	//
	// userId
	CheckPoint *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	// The check type.
	//
	// example:
	//
	// EXPORTED_PARAM
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The expected value.
	//
	// example:
	//
	// 111
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	// The check operator.
	//
	// example:
	//
	// ctn
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GetCheckPoint() *string {
	return s.CheckPoint
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GetCheckType() *string {
	return s.CheckType
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GetExpectValue() *string {
	return s.ExpectValue
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GetOperator() *string {
	return s.Operator
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) SetCheckPoint(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList {
	s.CheckPoint = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) SetCheckType(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList {
	s.CheckType = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) SetExpectValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList {
	s.ExpectValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) SetOperator(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList {
	s.Operator = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList struct {
	// The index of the export parameter.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the export parameter.
	//
	// example:
	//
	// userId
	ExportName *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	// The source of the export parameter.
	//
	// example:
	//
	// BODY_JSON
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The actual path from which you want to extract the export parameter values.
	//
	// example:
	//
	// data.userId
	ExportValue *string `json:"ExportValue,omitempty" xml:"ExportValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GetCount() *string {
	return s.Count
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GetExportName() *string {
	return s.ExportName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GetExportType() *string {
	return s.ExportType
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GetExportValue() *string {
	return s.ExportValue
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) SetCount(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList {
	s.Count = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) SetExportName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList {
	s.ExportName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) SetExportType(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList {
	s.ExportType = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) SetExportValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList {
	s.ExportValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList struct {
	// The name of the header.
	//
	// example:
	//
	// User-Agent
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The value of the header.
	//
	// example:
	//
	// PTS
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) GetHeaderName() *string {
	return s.HeaderName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) SetHeaderName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList {
	s.HeaderName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) SetHeaderValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList {
	s.HeaderValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) Validate() error {
	return dara.Validate(s)
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList struct {
	// Indicates whether the file is used as the baseline file.
	//
	// example:
	//
	// true
	BaseFile *bool `json:"BaseFile,omitempty" xml:"BaseFile,omitempty"`
	// Indicates whether the parameters are used once.
	//
	// example:
	//
	// false
	CycleOnce *bool `json:"CycleOnce,omitempty" xml:"CycleOnce,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The parameters in the file.
	//
	// example:
	//
	// username
	FileParamName *string `json:"FileParamName,omitempty" xml:"FileParamName,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GetBaseFile() *bool {
	return s.BaseFile
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GetCycleOnce() *bool {
	return s.CycleOnce
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GetFileName() *string {
	return s.FileName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GetFileParamName() *string {
	return s.FileParamName
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) SetBaseFile(v bool) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList {
	s.BaseFile = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) SetCycleOnce(v bool) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList {
	s.CycleOnce = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) SetFileName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList {
	s.FileName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) SetFileParamName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList {
	s.FileParamName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) Validate() error {
	return dara.Validate(s)
}
