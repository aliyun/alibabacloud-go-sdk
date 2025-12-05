// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSceneRunningDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetJMeterSceneRunningDataResponseBody
	GetCode() *string
	SetDocumentUrl(v string) *GetJMeterSceneRunningDataResponseBody
	GetDocumentUrl() *string
	SetHttpStatusCode(v int32) *GetJMeterSceneRunningDataResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetJMeterSceneRunningDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJMeterSceneRunningDataResponseBody
	GetRequestId() *string
	SetRunningData(v *GetJMeterSceneRunningDataResponseBodyRunningData) *GetJMeterSceneRunningDataResponseBody
	GetRunningData() *GetJMeterSceneRunningDataResponseBodyRunningData
	SetSuccess(v bool) *GetJMeterSceneRunningDataResponseBody
	GetSuccess() *bool
}

type GetJMeterSceneRunningDataResponseBody struct {
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The URL that is used to access the document.
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The HTTP status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The running data.
	RunningData *GetJMeterSceneRunningDataResponseBodyRunningData `json:"RunningData,omitempty" xml:"RunningData,omitempty" type:"Struct"`
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

func (s GetJMeterSceneRunningDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJMeterSceneRunningDataResponseBody) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *GetJMeterSceneRunningDataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetJMeterSceneRunningDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJMeterSceneRunningDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJMeterSceneRunningDataResponseBody) GetRunningData() *GetJMeterSceneRunningDataResponseBodyRunningData {
	return s.RunningData
}

func (s *GetJMeterSceneRunningDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterSceneRunningDataResponseBody) SetCode(v string) *GetJMeterSceneRunningDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetDocumentUrl(v string) *GetJMeterSceneRunningDataResponseBody {
	s.DocumentUrl = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetHttpStatusCode(v int32) *GetJMeterSceneRunningDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetMessage(v string) *GetJMeterSceneRunningDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetRequestId(v string) *GetJMeterSceneRunningDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetRunningData(v *GetJMeterSceneRunningDataResponseBodyRunningData) *GetJMeterSceneRunningDataResponseBody {
	s.RunningData = v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) SetSuccess(v bool) *GetJMeterSceneRunningDataResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBody) Validate() error {
	if s.RunningData != nil {
		if err := s.RunningData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJMeterSceneRunningDataResponseBodyRunningData struct {
	// The number of stress testing engines.
	//
	// example:
	//
	// 2
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// The stress testing engines.
	AgentIdList []*string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty" type:"Repeated"`
	// The sampling status of the scenario.
	//
	// example:
	//
	// { "failTps":0,"successRtAvg":33,"successRtMin":29,"successRtSum":99407,         "rtAvg":33.459104678559406,"rtMin":29,"failRtMax":0,"duration":997,         "samplerId":-1, "successRtMax":133,"fullStat":{       "requestBytesSum":629926,"successRtAvg":33,"successRtMin":29,"successRtSum":174551,"count":5206,"rtSeg99":53, "successTps":2397.9732842008293,"rtSeg90":36, "rtSeg50":32,            "rtSeg99Sum":53,"rtAvg":33.528812908182864, "rtMin":29,             "failRtMax":0,"duration":2171, "successCount":5206,            "rtSegStatCount":1,"tps":2397.9732842008293 }, "successCount":2971, "failRtSum":0,"failCount":0,"count":2971,"concurrency":100, "successTps":2979.939819458375,"tps":2979.939819458375,"failRtAvg":0,         "failRtMin":0,  "rtMax":133}
	AllSampleStat map[string]interface{} `json:"AllSampleStat,omitempty" xml:"AllSampleStat,omitempty"`
	// The concurrency.
	//
	// example:
	//
	// 1000
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// The error message returned for the stress testing process. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// Engine lease failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether an error occurs in the stress testing process.
	//
	// example:
	//
	// false
	HasError *bool `json:"HasError,omitempty" xml:"HasError,omitempty"`
	// Indicates whether the report is generated.
	//
	// example:
	//
	// false
	HasReport *bool `json:"HasReport,omitempty" xml:"HasReport,omitempty"`
	// The duration of the stress testing plan. Unit: seconds.
	//
	// example:
	//
	// 600
	HoldFor *int32 `json:"HoldFor,omitempty" xml:"HoldFor,omitempty"`
	// Indicates whether a debugging is performed.
	//
	// example:
	//
	// false
	IsDebugging *bool `json:"IsDebugging,omitempty" xml:"IsDebugging,omitempty"`
	// The stress testing task ID. This ID also means the report ID.
	//
	// example:
	//
	// DYYPLDKS
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The status of samplers.
	SampleStatList []map[string]interface{} `json:"SampleStatList,omitempty" xml:"SampleStatList,omitempty" type:"Repeated"`
	// The scenario ID.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The current stage.
	//
	// example:
	//
	// 任务执行
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The timestamp when the stress testing is scheduled to start. Unit: ms.
	//
	// example:
	//
	// 1639970040000
	StartTimeTS *int64 `json:"StartTimeTS,omitempty" xml:"StartTimeTS,omitempty"`
	// The stress testing status of the scenario.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The consumed Virtual User Minutes (VUM).
	//
	// example:
	//
	// 100
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetJMeterSceneRunningDataResponseBodyRunningData) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponseBodyRunningData) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetAgentCount() *int32 {
	return s.AgentCount
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetAgentIdList() []*string {
	return s.AgentIdList
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetAllSampleStat() map[string]interface{} {
	return s.AllSampleStat
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetConcurrency() *int32 {
	return s.Concurrency
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetHasError() *bool {
	return s.HasError
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetHasReport() *bool {
	return s.HasReport
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetHoldFor() *int32 {
	return s.HoldFor
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetIsDebugging() *bool {
	return s.IsDebugging
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetReportId() *string {
	return s.ReportId
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetSampleStatList() []map[string]interface{} {
	return s.SampleStatList
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetSceneId() *string {
	return s.SceneId
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetSceneName() *string {
	return s.SceneName
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetStageName() *string {
	return s.StageName
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetStartTimeTS() *int64 {
	return s.StartTimeTS
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetStatus() *string {
	return s.Status
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) GetVum() *int64 {
	return s.Vum
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetAgentCount(v int32) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.AgentCount = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetAgentIdList(v []*string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.AgentIdList = v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetAllSampleStat(v map[string]interface{}) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.AllSampleStat = v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetConcurrency(v int32) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.Concurrency = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetErrorMessage(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.ErrorMessage = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetHasError(v bool) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.HasError = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetHasReport(v bool) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.HasReport = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetHoldFor(v int32) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.HoldFor = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetIsDebugging(v bool) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.IsDebugging = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetReportId(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.ReportId = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetSampleStatList(v []map[string]interface{}) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.SampleStatList = v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetSceneId(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.SceneId = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetSceneName(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.SceneName = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetStageName(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.StageName = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetStartTimeTS(v int64) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.StartTimeTS = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetStatus(v string) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.Status = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) SetVum(v int64) *GetJMeterSceneRunningDataResponseBodyRunningData {
	s.Vum = &v
	return s
}

func (s *GetJMeterSceneRunningDataResponseBodyRunningData) Validate() error {
	return dara.Validate(s)
}
