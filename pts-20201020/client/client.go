// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreatePtsSceneRequest struct {
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s CreatePtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneRequest) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneRequest) SetScene(v string) *CreatePtsSceneRequest {
	s.Scene = &v
	return s
}

type CreatePtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId        *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneResponseBody) SetCode(v string) *CreatePtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetHttpStatusCode(v int32) *CreatePtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetMessage(v string) *CreatePtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetRequestId(v string) *CreatePtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetSceneId(v string) *CreatePtsSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *CreatePtsSceneResponseBody) SetSuccess(v bool) *CreatePtsSceneResponseBody {
	s.Success = &v
	return s
}

type CreatePtsSceneResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneResponse) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneResponse) SetHeaders(v map[string]*string) *CreatePtsSceneResponse {
	s.Headers = v
	return s
}

func (s *CreatePtsSceneResponse) SetBody(v *CreatePtsSceneResponseBody) *CreatePtsSceneResponse {
	s.Body = v
	return s
}

type CreatePtsSceneBaseLineFromReportRequest struct {
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	SceneId  *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s CreatePtsSceneBaseLineFromReportRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportRequest) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportRequest) SetReportId(v string) *CreatePtsSceneBaseLineFromReportRequest {
	s.ReportId = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportRequest) SetSceneId(v string) *CreatePtsSceneBaseLineFromReportRequest {
	s.SceneId = &v
	return s
}

type CreatePtsSceneBaseLineFromReportResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePtsSceneBaseLineFromReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetCode(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetHttpStatusCode(v int32) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetMessage(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetRequestId(v string) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponseBody) SetSuccess(v bool) *CreatePtsSceneBaseLineFromReportResponseBody {
	s.Success = &v
	return s
}

type CreatePtsSceneBaseLineFromReportResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePtsSceneBaseLineFromReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePtsSceneBaseLineFromReportResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePtsSceneBaseLineFromReportResponse) GoString() string {
	return s.String()
}

func (s *CreatePtsSceneBaseLineFromReportResponse) SetHeaders(v map[string]*string) *CreatePtsSceneBaseLineFromReportResponse {
	s.Headers = v
	return s
}

func (s *CreatePtsSceneBaseLineFromReportResponse) SetBody(v *CreatePtsSceneBaseLineFromReportResponseBody) *CreatePtsSceneBaseLineFromReportResponse {
	s.Body = v
	return s
}

type DeletePtsSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeletePtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneRequest) SetSceneId(v string) *DeletePtsSceneRequest {
	s.SceneId = &v
	return s
}

type DeletePtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneResponseBody) SetCode(v string) *DeletePtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetHttpStatusCode(v int32) *DeletePtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetMessage(v string) *DeletePtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetRequestId(v string) *DeletePtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetSuccess(v bool) *DeletePtsSceneResponseBody {
	s.Success = &v
	return s
}

type DeletePtsSceneResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneResponse) SetHeaders(v map[string]*string) *DeletePtsSceneResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsSceneResponse) SetBody(v *DeletePtsSceneResponseBody) *DeletePtsSceneResponse {
	s.Body = v
	return s
}

type DeletePtsSceneBaseLineRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DeletePtsSceneBaseLineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineRequest) SetSceneId(v string) *DeletePtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

type DeletePtsSceneBaseLineResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsSceneBaseLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineResponseBody) SetCode(v string) *DeletePtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *DeletePtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetMessage(v string) *DeletePtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetRequestId(v string) *DeletePtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponseBody) SetSuccess(v bool) *DeletePtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

type DeletePtsSceneBaseLineResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePtsSceneBaseLineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *DeletePtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsSceneBaseLineResponse) SetBody(v *DeletePtsSceneBaseLineResponseBody) *DeletePtsSceneBaseLineResponse {
	s.Body = v
	return s
}

type DeletePtsScenesRequest struct {
	SceneIds map[string]interface{} `json:"SceneIds,omitempty" xml:"SceneIds,omitempty"`
}

func (s DeletePtsScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsScenesRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesRequest) SetSceneIds(v map[string]interface{}) *DeletePtsScenesRequest {
	s.SceneIds = v
	return s
}

type DeletePtsScenesShrinkRequest struct {
	SceneIdsShrink *string `json:"SceneIds,omitempty" xml:"SceneIds,omitempty"`
}

func (s DeletePtsScenesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsScenesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesShrinkRequest) SetSceneIdsShrink(v string) *DeletePtsScenesShrinkRequest {
	s.SceneIdsShrink = &v
	return s
}

type DeletePtsScenesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsScenesResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesResponseBody) SetCode(v string) *DeletePtsScenesResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetHttpStatusCode(v int32) *DeletePtsScenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetMessage(v string) *DeletePtsScenesResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetRequestId(v string) *DeletePtsScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsScenesResponseBody) SetSuccess(v bool) *DeletePtsScenesResponseBody {
	s.Success = &v
	return s
}

type DeletePtsScenesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePtsScenesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePtsScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePtsScenesResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsScenesResponse) SetHeaders(v map[string]*string) *DeletePtsScenesResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsScenesResponse) SetBody(v *DeletePtsScenesResponseBody) *DeletePtsScenesResponse {
	s.Body = v
	return s
}

type GetJMeterLogsRequest struct {
	// 第几台引擎，起始为0
	AgentIndex *int32 `json:"AgentIndex,omitempty" xml:"AgentIndex,omitempty"`
	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 关键字
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 日志等级
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 报告ID
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// 线程名
	Thread *string `json:"Thread,omitempty" xml:"Thread,omitempty"`
}

func (s GetJMeterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterLogsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsRequest) SetAgentIndex(v int32) *GetJMeterLogsRequest {
	s.AgentIndex = &v
	return s
}

func (s *GetJMeterLogsRequest) SetBeginTime(v int64) *GetJMeterLogsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterLogsRequest) SetEndTime(v int64) *GetJMeterLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterLogsRequest) SetKeyword(v string) *GetJMeterLogsRequest {
	s.Keyword = &v
	return s
}

func (s *GetJMeterLogsRequest) SetLevel(v string) *GetJMeterLogsRequest {
	s.Level = &v
	return s
}

func (s *GetJMeterLogsRequest) SetPageNumber(v int32) *GetJMeterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterLogsRequest) SetPageSize(v int32) *GetJMeterLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetJMeterLogsRequest) SetReportId(v string) *GetJMeterLogsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterLogsRequest) SetThread(v string) *GetJMeterLogsRequest {
	s.Thread = &v
	return s
}

type GetJMeterLogsResponseBody struct {
	// 引擎数量，想要获得第几台引擎的日志可以根据引擎数量传值
	AgentCount *int32  `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 日志内容
	Logs       []map[string]interface{} `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	Message    *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetJMeterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsResponseBody) SetAgentCount(v int32) *GetJMeterLogsResponseBody {
	s.AgentCount = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetCode(v string) *GetJMeterLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetLogs(v []map[string]interface{}) *GetJMeterLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetJMeterLogsResponseBody) SetMessage(v string) *GetJMeterLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetPageNumber(v int32) *GetJMeterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetPageSize(v int32) *GetJMeterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetRequestId(v string) *GetJMeterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetSuccess(v bool) *GetJMeterLogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterLogsResponseBody) SetTotalCount(v int64) *GetJMeterLogsResponseBody {
	s.TotalCount = &v
	return s
}

type GetJMeterLogsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJMeterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJMeterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterLogsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterLogsResponse) SetHeaders(v map[string]*string) *GetJMeterLogsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterLogsResponse) SetBody(v *GetJMeterLogsResponseBody) *GetJMeterLogsResponse {
	s.Body = v
	return s
}

type GetJMeterSampleMetricsRequest struct {
	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 报告ID
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// 采样器索引，从0开始。-1返回全场景
	SamplerId *int32 `json:"SamplerId,omitempty" xml:"SamplerId,omitempty"`
}

func (s GetJMeterSampleMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSampleMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsRequest) SetBeginTime(v int64) *GetJMeterSampleMetricsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetEndTime(v int64) *GetJMeterSampleMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetReportId(v string) *GetJMeterSampleMetricsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterSampleMetricsRequest) SetSamplerId(v int32) *GetJMeterSampleMetricsRequest {
	s.SamplerId = &v
	return s
}

type GetJMeterSampleMetricsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 采样器聚合数据列表
	SampleMetricList []*string `json:"SampleMetricList,omitempty" xml:"SampleMetricList,omitempty" type:"Repeated"`
	// 采样器列表，可根据该列表传递需要查询的采样器
	SamplerMap map[string]interface{} `json:"SamplerMap,omitempty" xml:"SamplerMap,omitempty"`
	Success    *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJMeterSampleMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSampleMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsResponseBody) SetCode(v string) *GetJMeterSampleMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetMessage(v string) *GetJMeterSampleMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetRequestId(v string) *GetJMeterSampleMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSampleMetricList(v []*string) *GetJMeterSampleMetricsResponseBody {
	s.SampleMetricList = v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSamplerMap(v map[string]interface{}) *GetJMeterSampleMetricsResponseBody {
	s.SamplerMap = v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSuccess(v bool) *GetJMeterSampleMetricsResponseBody {
	s.Success = &v
	return s
}

type GetJMeterSampleMetricsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJMeterSampleMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJMeterSampleMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSampleMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsResponse) SetHeaders(v map[string]*string) *GetJMeterSampleMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSampleMetricsResponse) SetBody(v *GetJMeterSampleMetricsResponseBody) *GetJMeterSampleMetricsResponse {
	s.Body = v
	return s
}

type GetJMeterSamplingLogsRequest struct {
	// 压测引擎编号
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 开始时间
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 关键字
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 最大响应时间，单位ms
	MaxRT *int32 `json:"MaxRT,omitempty" xml:"MaxRT,omitempty"`
	// 最小响应时间，单位ms
	MinRT      *int32 `json:"MinRT,omitempty" xml:"MinRT,omitempty"`
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 报告ID
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ResponseCode *string `json:"ResponseCode,omitempty" xml:"ResponseCode,omitempty"`
	// 第几个采样器，从0开始
	SamplerId *int32 `json:"SamplerId,omitempty" xml:"SamplerId,omitempty"`
	// 采样结果是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 线程
	Thread *string `json:"Thread,omitempty" xml:"Thread,omitempty"`
}

func (s GetJMeterSamplingLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSamplingLogsRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsRequest) SetAgentId(v int64) *GetJMeterSamplingLogsRequest {
	s.AgentId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetBeginTime(v int64) *GetJMeterSamplingLogsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetEndTime(v int64) *GetJMeterSamplingLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetKeyword(v string) *GetJMeterSamplingLogsRequest {
	s.Keyword = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetMaxRT(v int32) *GetJMeterSamplingLogsRequest {
	s.MaxRT = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetMinRT(v int32) *GetJMeterSamplingLogsRequest {
	s.MinRT = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetPageNumber(v int32) *GetJMeterSamplingLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetPageSize(v int32) *GetJMeterSamplingLogsRequest {
	s.PageSize = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetReportId(v string) *GetJMeterSamplingLogsRequest {
	s.ReportId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetResponseCode(v string) *GetJMeterSamplingLogsRequest {
	s.ResponseCode = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetSamplerId(v int32) *GetJMeterSamplingLogsRequest {
	s.SamplerId = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetSuccess(v bool) *GetJMeterSamplingLogsRequest {
	s.Success = &v
	return s
}

func (s *GetJMeterSamplingLogsRequest) SetThread(v string) *GetJMeterSamplingLogsRequest {
	s.Thread = &v
	return s
}

type GetJMeterSamplingLogsResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 采样器的采样结果
	SampleResults []*string `json:"SampleResults,omitempty" xml:"SampleResults,omitempty" type:"Repeated"`
	Success       *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount    *int64    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetJMeterSamplingLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSamplingLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsResponseBody) SetCode(v string) *GetJMeterSamplingLogsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetHttpStatusCode(v int32) *GetJMeterSamplingLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetMessage(v string) *GetJMeterSamplingLogsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetPageNumber(v int32) *GetJMeterSamplingLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetPageSize(v int32) *GetJMeterSamplingLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetRequestId(v string) *GetJMeterSamplingLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetSampleResults(v []*string) *GetJMeterSamplingLogsResponseBody {
	s.SampleResults = v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetSuccess(v bool) *GetJMeterSamplingLogsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterSamplingLogsResponseBody) SetTotalCount(v int64) *GetJMeterSamplingLogsResponseBody {
	s.TotalCount = &v
	return s
}

type GetJMeterSamplingLogsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJMeterSamplingLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJMeterSamplingLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSamplingLogsResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSamplingLogsResponse) SetHeaders(v map[string]*string) *GetJMeterSamplingLogsResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSamplingLogsResponse) SetBody(v *GetJMeterSamplingLogsResponseBody) *GetJMeterSamplingLogsResponse {
	s.Body = v
	return s
}

type GetJMeterSceneRunningDataRequest struct {
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetJMeterSceneRunningDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSceneRunningDataRequest) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataRequest) SetSceneId(v string) *GetJMeterSceneRunningDataRequest {
	s.SceneId = &v
	return s
}

type GetJMeterSceneRunningDataResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DocumentUrl    *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 运行中的数据
	RunningData *GetJMeterSceneRunningDataResponseBodyRunningData `json:"RunningData,omitempty" xml:"RunningData,omitempty" type:"Struct"`
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJMeterSceneRunningDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponseBody) GoString() string {
	return s.String()
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

type GetJMeterSceneRunningDataResponseBodyRunningData struct {
	// 压测引擎数量
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// 压测引擎列表
	AgentIdList []*string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty" type:"Repeated"`
	// 场景整体的采样状态
	AllSampleStat map[string]interface{} `json:"AllSampleStat,omitempty" xml:"AllSampleStat,omitempty"`
	// 并发量
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// 是否生成了报告
	HasReport *bool `json:"HasReport,omitempty" xml:"HasReport,omitempty"`
	// 压测计划持续时间，单位s
	HoldFor *int32 `json:"HoldFor,omitempty" xml:"HoldFor,omitempty"`
	// 是否是调试
	IsDebugging *bool `json:"IsDebugging,omitempty" xml:"IsDebugging,omitempty"`
	// 每一个采样器的状态
	SampleStatList []map[string]interface{} `json:"SampleStatList,omitempty" xml:"SampleStatList,omitempty" type:"Repeated"`
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名称
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 当前所处阶段
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// 压测计划开始时间戳，单位ms
	StartTimeTS *int64 `json:"StartTimeTS,omitempty" xml:"StartTimeTS,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 目前消耗的vum
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetJMeterSceneRunningDataResponseBodyRunningData) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponseBodyRunningData) GoString() string {
	return s.String()
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

type GetJMeterSceneRunningDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJMeterSceneRunningDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJMeterSceneRunningDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJMeterSceneRunningDataResponse) GoString() string {
	return s.String()
}

func (s *GetJMeterSceneRunningDataResponse) SetHeaders(v map[string]*string) *GetJMeterSceneRunningDataResponse {
	s.Headers = v
	return s
}

func (s *GetJMeterSceneRunningDataResponse) SetBody(v *GetJMeterSceneRunningDataResponseBody) *GetJMeterSceneRunningDataResponse {
	s.Body = v
	return s
}

type GetOpenJMeterSceneRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetOpenJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneRequest) SetSceneId(v string) *GetOpenJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type GetOpenJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 场景详情
	Scene   *GetOpenJMeterSceneResponseBodyScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Struct"`
	Success *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpenJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBody) SetCode(v string) *GetOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *GetOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetMessage(v string) *GetOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetRequestId(v string) *GetOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetScene(v *GetOpenJMeterSceneResponseBodyScene) *GetOpenJMeterSceneResponseBody {
	s.Scene = v
	return s
}

func (s *GetOpenJMeterSceneResponseBody) SetSuccess(v bool) *GetOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type GetOpenJMeterSceneResponseBodyScene struct {
	// 施压机数量
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// 基本信息
	BaseInfo *GetOpenJMeterSceneResponseBodySceneBaseInfo `json:"BaseInfo,omitempty" xml:"BaseInfo,omitempty" type:"Struct"`
	// 最大并发
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// constantThroughputTimerType
	ConstantThroughputTimerType *string `json:"ConstantThroughputTimerType,omitempty" xml:"ConstantThroughputTimerType,omitempty"`
	// DNS配置
	DnsCacheConfig *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig `json:"DnsCacheConfig,omitempty" xml:"DnsCacheConfig,omitempty" type:"Struct"`
	// 压测持续时间，单位为s
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// 文件列表
	FileList []*GetOpenJMeterSceneResponseBodySceneFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// 是否为VPC压测
	IsVpcTest *bool `json:"IsVpcTest,omitempty" xml:"IsVpcTest,omitempty"`
	// 递增时间，单位s
	RampUp *int32 `json:"RampUp,omitempty" xml:"RampUp,omitempty"`
	// VPC压测时配置
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 安全组id，VPC压测时配置
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 递增阶梯数。预热时间和预热阶段数都不配置时 使用固定压力值 只配置预热时间，不配置阶段数时 使用均匀递增 预热时间和阶段数都配置时，并且steps<rampUp 使用阶梯递增 不能只配置steps，不配置rampUp 如果这样配置，默认使用固定压力值
	Steps *int32 `json:"Steps,omitempty" xml:"Steps,omitempty"`
	// synchronizing timer 类型
	SyncTimerType *string `json:"SyncTimerType,omitempty" xml:"SyncTimerType,omitempty"`
	// 测试文件
	TestFile *string `json:"TestFile,omitempty" xml:"TestFile,omitempty"`
	// 交换机id，VPC压测时配置
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// vpc的id，VPC压测时配置
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodyScene) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodyScene) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetAgentCount(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.AgentCount = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetBaseInfo(v *GetOpenJMeterSceneResponseBodySceneBaseInfo) *GetOpenJMeterSceneResponseBodyScene {
	s.BaseInfo = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetConcurrency(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Concurrency = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetConstantThroughputTimerType(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.ConstantThroughputTimerType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetDnsCacheConfig(v *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) *GetOpenJMeterSceneResponseBodyScene {
	s.DnsCacheConfig = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetDuration(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Duration = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetEnvironmentId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.EnvironmentId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetFileList(v []*GetOpenJMeterSceneResponseBodySceneFileList) *GetOpenJMeterSceneResponseBodyScene {
	s.FileList = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetIsVpcTest(v bool) *GetOpenJMeterSceneResponseBodyScene {
	s.IsVpcTest = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetRampUp(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.RampUp = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetRegionId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.RegionId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSceneId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SceneId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSceneName(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SceneName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSecurityGroupId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SecurityGroupId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSteps(v int32) *GetOpenJMeterSceneResponseBodyScene {
	s.Steps = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetSyncTimerType(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.SyncTimerType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetTestFile(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.TestFile = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetVSwitchId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.VSwitchId = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodyScene) SetVpcId(v string) *GetOpenJMeterSceneResponseBodyScene {
	s.VpcId = &v
	return s
}

type GetOpenJMeterSceneResponseBodySceneBaseInfo struct {
	// 创建人名
	CreateName *string `json:"CreateName,omitempty" xml:"CreateName,omitempty"`
	// 修改人名
	ModifyName *string `json:"ModifyName,omitempty" xml:"ModifyName,omitempty"`
	// 操作类型
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// 场景压测负责人
	Principal *string `json:"Principal,omitempty" xml:"Principal,omitempty"`
	// 备注
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// 场景来源
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneBaseInfo) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetCreateName(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.CreateName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetModifyName(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.ModifyName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetOperateType(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.OperateType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetPrincipal(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Principal = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetRemark(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Remark = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneBaseInfo) SetResource(v string) *GetOpenJMeterSceneResponseBodySceneBaseInfo {
	s.Resource = &v
	return s
}

type GetOpenJMeterSceneResponseBodySceneDnsCacheConfig struct {
	// 是否清除缓存
	ClearCacheEachIteration *bool `json:"ClearCacheEachIteration,omitempty" xml:"ClearCacheEachIteration,omitempty"`
	// DNS服务器
	DnsServers []*string `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	// 域名绑定
	HostTable map[string]interface{} `json:"HostTable,omitempty" xml:"HostTable,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetClearCacheEachIteration(v bool) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.ClearCacheEachIteration = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetDnsServers(v []*string) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.DnsServers = v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig) SetHostTable(v map[string]interface{}) *GetOpenJMeterSceneResponseBodySceneDnsCacheConfig {
	s.HostTable = v
	return s
}

type GetOpenJMeterSceneResponseBodySceneFileList struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件地址
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// 文件大小
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件类型
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// 文件ID
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 文件的md5值
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// csv文件是否切分
	SplitCsv *bool `json:"SplitCsv,omitempty" xml:"SplitCsv,omitempty"`
}

func (s GetOpenJMeterSceneResponseBodySceneFileList) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponseBodySceneFileList) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileName(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileName = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileOssAddress(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileOssAddress = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileSize(v int64) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileSize = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetFileType(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.FileType = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetId(v int64) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.Id = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetMd5(v string) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.Md5 = &v
	return s
}

func (s *GetOpenJMeterSceneResponseBodySceneFileList) SetSplitCsv(v bool) *GetOpenJMeterSceneResponseBodySceneFileList {
	s.SplitCsv = &v
	return s
}

type GetOpenJMeterSceneResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOpenJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *GetOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *GetOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *GetOpenJMeterSceneResponse) SetBody(v *GetOpenJMeterSceneResponseBody) *GetOpenJMeterSceneResponse {
	s.Body = v
	return s
}

type GetPtsReportDetailsRequest struct {
	PlanId  *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsReportDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsRequest) SetPlanId(v string) *GetPtsReportDetailsRequest {
	s.PlanId = &v
	return s
}

func (s *GetPtsReportDetailsRequest) SetSceneId(v string) *GetPtsReportDetailsRequest {
	s.SceneId = &v
	return s
}

type GetPtsReportDetailsResponseBody struct {
	ApiMetricsList []*GetPtsReportDetailsResponseBodyApiMetricsList `json:"ApiMetricsList,omitempty" xml:"ApiMetricsList,omitempty" type:"Repeated"`
	Code           *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	ReportOverView *GetPtsReportDetailsResponseBodyReportOverView   `json:"ReportOverView,omitempty" xml:"ReportOverView,omitempty" type:"Struct"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneMetrics   *GetPtsReportDetailsResponseBodySceneMetrics     `json:"SceneMetrics,omitempty" xml:"SceneMetrics,omitempty" type:"Struct"`
	SceneSnapShot  *GetPtsReportDetailsResponseBodySceneSnapShot    `json:"SceneSnapShot,omitempty" xml:"SceneSnapShot,omitempty" type:"Struct"`
	Success        *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsReportDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBody) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodyApiMetricsList struct {
	AllCount       *int64   `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	ApiName        *string  `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AvgRt          *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgTps         *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	FailCountBiz   *int64   `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	FailCountReq   *int64   `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	MaxRt          *float32 `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	MinRt          *float32 `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	Seg50Rt        *float32 `json:"Seg50Rt,omitempty" xml:"Seg50Rt,omitempty"`
	Seg75Rt        *float32 `json:"Seg75Rt,omitempty" xml:"Seg75Rt,omitempty"`
	Seg90Rt        *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	Seg99Rt        *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsReportDetailsResponseBodyApiMetricsList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodyApiMetricsList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodyReportOverView struct {
	AgentCount *int32  `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReportId   *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Vum        *int64  `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsReportDetailsResponseBodyReportOverView) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodyReportOverView) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneMetrics struct {
	AllCount       *int64   `json:"AllCount,omitempty" xml:"AllCount,omitempty"`
	AvgRt          *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgTps         *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	FailCountBiz   *int64   `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	FailCountReq   *int64   `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	Seg90Rt        *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	Seg99Rt        *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneMetrics) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneMetrics) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShot struct {
	AdvanceSetting      *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting        `json:"AdvanceSetting,omitempty" xml:"AdvanceSetting,omitempty" type:"Struct"`
	CreateTime          *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileParameterList   []*GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList   `json:"FileParameterList,omitempty" xml:"FileParameterList,omitempty" type:"Repeated"`
	GlobalParameterList []*GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList `json:"GlobalParameterList,omitempty" xml:"GlobalParameterList,omitempty" type:"Repeated"`
	LoadConfig          *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig            `json:"LoadConfig,omitempty" xml:"LoadConfig,omitempty" type:"Struct"`
	ModifiedTime        *string                                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RelationList        []*GetPtsReportDetailsResponseBodySceneSnapShotRelationList        `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	SceneId             *string                                                            `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName           *string                                                            `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	Status              *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShot) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShot) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting struct {
	ConnectionTimeoutInSecond *int32                                                                         `json:"ConnectionTimeoutInSecond,omitempty" xml:"ConnectionTimeoutInSecond,omitempty"`
	DomainBindingList         []*GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList `json:"DomainBindingList,omitempty" xml:"DomainBindingList,omitempty" type:"Repeated"`
	LogRate                   *int32                                                                         `json:"LogRate,omitempty" xml:"LogRate,omitempty"`
	SuccessCode               *string                                                                        `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSetting) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList struct {
	Domain *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ips    []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) SetDomain(v string) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList {
	s.Domain = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList) SetIps(v []*string) *GetPtsReportDetailsResponseBodySceneSnapShotAdvanceSettingDomainBindingList {
	s.Ips = v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList struct {
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) SetFileName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList {
	s.FileName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList) SetFileOssAddress(v string) *GetPtsReportDetailsResponseBodySceneSnapShotFileParameterList {
	s.FileOssAddress = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList struct {
	ParamName  *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) SetParamName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList {
	s.ParamName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList) SetParamValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotGlobalParameterList {
	s.ParamValue = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig struct {
	AgentCount             *int32                                                                          `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	ApiLoadConfigList      []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList      `json:"ApiLoadConfigList,omitempty" xml:"ApiLoadConfigList,omitempty" type:"Repeated"`
	Configuration          *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration            `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	MaxRunningTime         *int32                                                                          `json:"MaxRunningTime,omitempty" xml:"MaxRunningTime,omitempty"`
	RelationLoadConfigList []*GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList `json:"RelationLoadConfigList,omitempty" xml:"RelationLoadConfigList,omitempty" type:"Repeated"`
	TestMode               *string                                                                         `json:"TestMode,omitempty" xml:"TestMode,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfig) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList struct {
	RpsBegin *int32 `json:"RpsBegin,omitempty" xml:"RpsBegin,omitempty"`
	RpsLimit *int32 `json:"RpsLimit,omitempty" xml:"RpsLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) SetRpsBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList {
	s.RpsBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList) SetRpsLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigApiLoadConfigList {
	s.RpsLimit = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration struct {
	AllConcurrencyBegin *int32 `json:"AllConcurrencyBegin,omitempty" xml:"AllConcurrencyBegin,omitempty"`
	AllConcurrencyLimit *int32 `json:"AllConcurrencyLimit,omitempty" xml:"AllConcurrencyLimit,omitempty"`
	AllRpsBegin         *int32 `json:"AllRpsBegin,omitempty" xml:"AllRpsBegin,omitempty"`
	AllRpsLimit         *int32 `json:"AllRpsLimit,omitempty" xml:"AllRpsLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigConfiguration) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList struct {
	ConcurrencyBegin *int32 `json:"ConcurrencyBegin,omitempty" xml:"ConcurrencyBegin,omitempty"`
	ConcurrencyLimit *int32 `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) SetConcurrencyBegin(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList {
	s.ConcurrencyBegin = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList) SetConcurrencyLimit(v int32) *GetPtsReportDetailsResponseBodySceneSnapShotLoadConfigRelationLoadConfigList {
	s.ConcurrencyLimit = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationList struct {
	ApiList                  []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList                  `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	FileParameterExplainList []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList `json:"FileParameterExplainList,omitempty" xml:"FileParameterExplainList,omitempty" type:"Repeated"`
	RelationId               *string                                                                             `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	RelationName             *string                                                                             `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList struct {
	ApiId              *string                                                                          `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName            *string                                                                          `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Body               *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody             `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	CheckPointList     []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList `json:"CheckPointList,omitempty" xml:"CheckPointList,omitempty" type:"Repeated"`
	ExportList         []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList     `json:"ExportList,omitempty" xml:"ExportList,omitempty" type:"Repeated"`
	HeaderList         []*GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList     `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Repeated"`
	Method             *string                                                                          `json:"Method,omitempty" xml:"Method,omitempty"`
	RedirectCountLimit *int32                                                                           `json:"RedirectCountLimit,omitempty" xml:"RedirectCountLimit,omitempty"`
	TimeoutInSecond    *int32                                                                           `json:"TimeoutInSecond,omitempty" xml:"TimeoutInSecond,omitempty"`
	Url                *string                                                                          `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody struct {
	BodyValue   *string `json:"BodyValue,omitempty" xml:"BodyValue,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) SetBodyValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody {
	s.BodyValue = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody) SetContentType(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListBody {
	s.ContentType = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList struct {
	CheckPoint  *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	CheckType   *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListCheckPointList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList struct {
	Count       *string `json:"Count,omitempty" xml:"Count,omitempty"`
	ExportName  *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	ExportType  *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	ExportValue *string `json:"ExportValue,omitempty" xml:"ExportValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListExportList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList struct {
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) SetHeaderName(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList {
	s.HeaderName = &v
	return s
}

func (s *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList) SetHeaderValue(v string) *GetPtsReportDetailsResponseBodySceneSnapShotRelationListApiListHeaderList {
	s.HeaderValue = &v
	return s
}

type GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList struct {
	BaseFile      *bool   `json:"BaseFile,omitempty" xml:"BaseFile,omitempty"`
	CycleOnce     *bool   `json:"CycleOnce,omitempty" xml:"CycleOnce,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileParamName *string `json:"FileParamName,omitempty" xml:"FileParamName,omitempty"`
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponseBodySceneSnapShotRelationListFileParameterExplainList) GoString() string {
	return s.String()
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

type GetPtsReportDetailsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsReportDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsReportDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsResponse) SetHeaders(v map[string]*string) *GetPtsReportDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetPtsReportDetailsResponse) SetBody(v *GetPtsReportDetailsResponseBody) *GetPtsReportDetailsResponse {
	s.Body = v
	return s
}

type GetPtsReportsBySceneIdRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsReportsBySceneIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportsBySceneIdRequest) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdRequest) SetPageNumber(v int32) *GetPtsReportsBySceneIdRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPtsReportsBySceneIdRequest) SetPageSize(v int32) *GetPtsReportsBySceneIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetPtsReportsBySceneIdRequest) SetSceneId(v string) *GetPtsReportsBySceneIdRequest {
	s.SceneId = &v
	return s
}

type GetPtsReportsBySceneIdResponseBody struct {
	Code               *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode     *int32                                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	ReportOverViewList []*GetPtsReportsBySceneIdResponseBodyReportOverViewList `json:"ReportOverViewList,omitempty" xml:"ReportOverViewList,omitempty" type:"Repeated"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsReportsBySceneIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponseBody) SetCode(v string) *GetPtsReportsBySceneIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetHttpStatusCode(v int32) *GetPtsReportsBySceneIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetMessage(v string) *GetPtsReportsBySceneIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetReportOverViewList(v []*GetPtsReportsBySceneIdResponseBodyReportOverViewList) *GetPtsReportsBySceneIdResponseBody {
	s.ReportOverViewList = v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetRequestId(v string) *GetPtsReportsBySceneIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBody) SetSuccess(v bool) *GetPtsReportsBySceneIdResponseBody {
	s.Success = &v
	return s
}

type GetPtsReportsBySceneIdResponseBodyReportOverViewList struct {
	AgentCount *int32  `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReportId   *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Vum        *int64  `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsReportsBySceneIdResponseBodyReportOverViewList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponseBodyReportOverViewList) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetAgentCount(v int32) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.AgentCount = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetEndTime(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.EndTime = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetReportId(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.ReportId = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetReportName(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.ReportName = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetStartTime(v string) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.StartTime = &v
	return s
}

func (s *GetPtsReportsBySceneIdResponseBodyReportOverViewList) SetVum(v int64) *GetPtsReportsBySceneIdResponseBodyReportOverViewList {
	s.Vum = &v
	return s
}

type GetPtsReportsBySceneIdResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsReportsBySceneIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsReportsBySceneIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsReportsBySceneIdResponse) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdResponse) SetHeaders(v map[string]*string) *GetPtsReportsBySceneIdResponse {
	s.Headers = v
	return s
}

func (s *GetPtsReportsBySceneIdResponse) SetBody(v *GetPtsReportsBySceneIdResponseBody) *GetPtsReportsBySceneIdResponse {
	s.Body = v
	return s
}

type GetPtsSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRequest) SetSceneId(v string) *GetPtsSceneRequest {
	s.SceneId = &v
	return s
}

type GetPtsSceneResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scene          *GetPtsSceneResponseBodyScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Struct"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBody) SetCode(v string) *GetPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetMessage(v string) *GetPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetRequestId(v string) *GetPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneResponseBody) SetScene(v *GetPtsSceneResponseBodyScene) *GetPtsSceneResponseBody {
	s.Scene = v
	return s
}

func (s *GetPtsSceneResponseBody) SetSuccess(v bool) *GetPtsSceneResponseBody {
	s.Success = &v
	return s
}

type GetPtsSceneResponseBodyScene struct {
	AdvanceSetting      *GetPtsSceneResponseBodySceneAdvanceSetting        `json:"AdvanceSetting,omitempty" xml:"AdvanceSetting,omitempty" type:"Struct"`
	CreateTime          *string                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileParameterList   []*GetPtsSceneResponseBodySceneFileParameterList   `json:"FileParameterList,omitempty" xml:"FileParameterList,omitempty" type:"Repeated"`
	GlobalParameterList []*GetPtsSceneResponseBodySceneGlobalParameterList `json:"GlobalParameterList,omitempty" xml:"GlobalParameterList,omitempty" type:"Repeated"`
	LoadConfig          *GetPtsSceneResponseBodySceneLoadConfig            `json:"LoadConfig,omitempty" xml:"LoadConfig,omitempty" type:"Struct"`
	ModifiedTime        *string                                            `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RelationList        []*GetPtsSceneResponseBodySceneRelationList        `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
	SceneId             *string                                            `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName           *string                                            `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	Status              *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetPtsSceneResponseBodyScene) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodyScene) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodyScene) SetAdvanceSetting(v *GetPtsSceneResponseBodySceneAdvanceSetting) *GetPtsSceneResponseBodyScene {
	s.AdvanceSetting = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetCreateTime(v string) *GetPtsSceneResponseBodyScene {
	s.CreateTime = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetFileParameterList(v []*GetPtsSceneResponseBodySceneFileParameterList) *GetPtsSceneResponseBodyScene {
	s.FileParameterList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetGlobalParameterList(v []*GetPtsSceneResponseBodySceneGlobalParameterList) *GetPtsSceneResponseBodyScene {
	s.GlobalParameterList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetLoadConfig(v *GetPtsSceneResponseBodySceneLoadConfig) *GetPtsSceneResponseBodyScene {
	s.LoadConfig = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetModifiedTime(v string) *GetPtsSceneResponseBodyScene {
	s.ModifiedTime = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetRelationList(v []*GetPtsSceneResponseBodySceneRelationList) *GetPtsSceneResponseBodyScene {
	s.RelationList = v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetSceneId(v string) *GetPtsSceneResponseBodyScene {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetSceneName(v string) *GetPtsSceneResponseBodyScene {
	s.SceneName = &v
	return s
}

func (s *GetPtsSceneResponseBodyScene) SetStatus(v string) *GetPtsSceneResponseBodyScene {
	s.Status = &v
	return s
}

type GetPtsSceneResponseBodySceneAdvanceSetting struct {
	ConnectionTimeoutInSecond *int32                                                         `json:"ConnectionTimeoutInSecond,omitempty" xml:"ConnectionTimeoutInSecond,omitempty"`
	DomainBindingList         []*GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList `json:"DomainBindingList,omitempty" xml:"DomainBindingList,omitempty" type:"Repeated"`
	LogRate                   *int32                                                         `json:"LogRate,omitempty" xml:"LogRate,omitempty"`
	SuccessCode               *string                                                        `json:"SuccessCode,omitempty" xml:"SuccessCode,omitempty"`
}

func (s GetPtsSceneResponseBodySceneAdvanceSetting) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneAdvanceSetting) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetConnectionTimeoutInSecond(v int32) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.ConnectionTimeoutInSecond = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetDomainBindingList(v []*GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.DomainBindingList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetLogRate(v int32) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.LogRate = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSetting) SetSuccessCode(v string) *GetPtsSceneResponseBodySceneAdvanceSetting {
	s.SuccessCode = &v
	return s
}

type GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList struct {
	Domain *string   `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Ips    []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) SetDomain(v string) *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList {
	s.Domain = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList) SetIps(v []*string) *GetPtsSceneResponseBodySceneAdvanceSettingDomainBindingList {
	s.Ips = v
	return s
}

type GetPtsSceneResponseBodySceneFileParameterList struct {
	FileName       *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s GetPtsSceneResponseBodySceneFileParameterList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneFileParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) SetFileName(v string) *GetPtsSceneResponseBodySceneFileParameterList {
	s.FileName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneFileParameterList) SetFileOssAddress(v string) *GetPtsSceneResponseBodySceneFileParameterList {
	s.FileOssAddress = &v
	return s
}

type GetPtsSceneResponseBodySceneGlobalParameterList struct {
	ParamName  *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneGlobalParameterList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneGlobalParameterList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) SetParamName(v string) *GetPtsSceneResponseBodySceneGlobalParameterList {
	s.ParamName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneGlobalParameterList) SetParamValue(v string) *GetPtsSceneResponseBodySceneGlobalParameterList {
	s.ParamValue = &v
	return s
}

type GetPtsSceneResponseBodySceneLoadConfig struct {
	AgentCount             *int32                                                          `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	ApiLoadConfigList      []*GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList      `json:"ApiLoadConfigList,omitempty" xml:"ApiLoadConfigList,omitempty" type:"Repeated"`
	Configuration          *GetPtsSceneResponseBodySceneLoadConfigConfiguration            `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Struct"`
	MaxRunningTime         *int32                                                          `json:"MaxRunningTime,omitempty" xml:"MaxRunningTime,omitempty"`
	RelationLoadConfigList []*GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList `json:"RelationLoadConfigList,omitempty" xml:"RelationLoadConfigList,omitempty" type:"Repeated"`
	TestMode               *string                                                         `json:"TestMode,omitempty" xml:"TestMode,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfig) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetAgentCount(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.AgentCount = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetApiLoadConfigList(v []*GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) *GetPtsSceneResponseBodySceneLoadConfig {
	s.ApiLoadConfigList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetConfiguration(v *GetPtsSceneResponseBodySceneLoadConfigConfiguration) *GetPtsSceneResponseBodySceneLoadConfig {
	s.Configuration = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetMaxRunningTime(v int32) *GetPtsSceneResponseBodySceneLoadConfig {
	s.MaxRunningTime = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetRelationLoadConfigList(v []*GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) *GetPtsSceneResponseBodySceneLoadConfig {
	s.RelationLoadConfigList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfig) SetTestMode(v string) *GetPtsSceneResponseBodySceneLoadConfig {
	s.TestMode = &v
	return s
}

type GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList struct {
	ApiId    *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	RpsBegin *int32  `json:"RpsBegin,omitempty" xml:"RpsBegin,omitempty"`
	RpsLimit *int32  `json:"RpsLimit,omitempty" xml:"RpsLimit,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetApiId(v string) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.ApiId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetRpsBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.RpsBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList) SetRpsLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigApiLoadConfigList {
	s.RpsLimit = &v
	return s
}

type GetPtsSceneResponseBodySceneLoadConfigConfiguration struct {
	AllConcurrencyBegin *int32 `json:"AllConcurrencyBegin,omitempty" xml:"AllConcurrencyBegin,omitempty"`
	AllConcurrencyLimit *int32 `json:"AllConcurrencyLimit,omitempty" xml:"AllConcurrencyLimit,omitempty"`
	AllRpsBegin         *int32 `json:"AllRpsBegin,omitempty" xml:"AllRpsBegin,omitempty"`
	AllRpsLimit         *int32 `json:"AllRpsLimit,omitempty" xml:"AllRpsLimit,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigConfiguration) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigConfiguration) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllConcurrencyBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllConcurrencyBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllConcurrencyLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllConcurrencyLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllRpsBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllRpsBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigConfiguration) SetAllRpsLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigConfiguration {
	s.AllRpsLimit = &v
	return s
}

type GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList struct {
	ConcurrencyBegin *int32  `json:"ConcurrencyBegin,omitempty" xml:"ConcurrencyBegin,omitempty"`
	ConcurrencyLimit *int32  `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	RelationId       *string `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
}

func (s GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetConcurrencyBegin(v int32) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyBegin = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetConcurrencyLimit(v int32) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.ConcurrencyLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList) SetRelationId(v string) *GetPtsSceneResponseBodySceneLoadConfigRelationLoadConfigList {
	s.RelationId = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationList struct {
	ApiList                  []*GetPtsSceneResponseBodySceneRelationListApiList                  `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	FileParameterExplainList []*GetPtsSceneResponseBodySceneRelationListFileParameterExplainList `json:"FileParameterExplainList,omitempty" xml:"FileParameterExplainList,omitempty" type:"Repeated"`
	RelationId               *string                                                             `json:"RelationId,omitempty" xml:"RelationId,omitempty"`
	RelationName             *string                                                             `json:"RelationName,omitempty" xml:"RelationName,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetApiList(v []*GetPtsSceneResponseBodySceneRelationListApiList) *GetPtsSceneResponseBodySceneRelationList {
	s.ApiList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetFileParameterExplainList(v []*GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) *GetPtsSceneResponseBodySceneRelationList {
	s.FileParameterExplainList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetRelationId(v string) *GetPtsSceneResponseBodySceneRelationList {
	s.RelationId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationList) SetRelationName(v string) *GetPtsSceneResponseBodySceneRelationList {
	s.RelationName = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListApiList struct {
	ApiId              *string                                                          `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName            *string                                                          `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	Body               *GetPtsSceneResponseBodySceneRelationListApiListBody             `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	CheckPointList     []*GetPtsSceneResponseBodySceneRelationListApiListCheckPointList `json:"CheckPointList,omitempty" xml:"CheckPointList,omitempty" type:"Repeated"`
	ExportList         []*GetPtsSceneResponseBodySceneRelationListApiListExportList     `json:"ExportList,omitempty" xml:"ExportList,omitempty" type:"Repeated"`
	HeaderList         []*GetPtsSceneResponseBodySceneRelationListApiListHeaderList     `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Repeated"`
	Method             *string                                                          `json:"Method,omitempty" xml:"Method,omitempty"`
	RedirectCountLimit *int32                                                           `json:"RedirectCountLimit,omitempty" xml:"RedirectCountLimit,omitempty"`
	TimeoutInSecond    *int32                                                           `json:"TimeoutInSecond,omitempty" xml:"TimeoutInSecond,omitempty"`
	Url                *string                                                          `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetApiId(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ApiId = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetApiName(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ApiName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetBody(v *GetPtsSceneResponseBodySceneRelationListApiListBody) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Body = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetCheckPointList(v []*GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.CheckPointList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetExportList(v []*GetPtsSceneResponseBodySceneRelationListApiListExportList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.ExportList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetHeaderList(v []*GetPtsSceneResponseBodySceneRelationListApiListHeaderList) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.HeaderList = v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetMethod(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Method = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetRedirectCountLimit(v int32) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.RedirectCountLimit = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetTimeoutInSecond(v int32) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.TimeoutInSecond = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiList) SetUrl(v string) *GetPtsSceneResponseBodySceneRelationListApiList {
	s.Url = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListApiListBody struct {
	BodyValue   *string `json:"BodyValue,omitempty" xml:"BodyValue,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) SetBodyValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListBody {
	s.BodyValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListBody) SetContentType(v string) *GetPtsSceneResponseBodySceneRelationListApiListBody {
	s.ContentType = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListApiListCheckPointList struct {
	CheckPoint  *string `json:"CheckPoint,omitempty" xml:"CheckPoint,omitempty"`
	CheckType   *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	ExpectValue *string `json:"ExpectValue,omitempty" xml:"ExpectValue,omitempty"`
	Operator    *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetCheckPoint(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.CheckPoint = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetCheckType(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.CheckType = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetExpectValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.ExpectValue = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList) SetOperator(v string) *GetPtsSceneResponseBodySceneRelationListApiListCheckPointList {
	s.Operator = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListApiListExportList struct {
	Count       *string `json:"Count,omitempty" xml:"Count,omitempty"`
	ExportName  *string `json:"ExportName,omitempty" xml:"ExportName,omitempty"`
	ExportType  *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	ExportValue *string `json:"ExportValue,omitempty" xml:"ExportValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListExportList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListExportList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetCount(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.Count = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportName(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportType(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportType = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListExportList) SetExportValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListExportList {
	s.ExportValue = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListApiListHeaderList struct {
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListApiListHeaderList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListApiListHeaderList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) SetHeaderName(v string) *GetPtsSceneResponseBodySceneRelationListApiListHeaderList {
	s.HeaderName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListApiListHeaderList) SetHeaderValue(v string) *GetPtsSceneResponseBodySceneRelationListApiListHeaderList {
	s.HeaderValue = &v
	return s
}

type GetPtsSceneResponseBodySceneRelationListFileParameterExplainList struct {
	BaseFile      *bool   `json:"BaseFile,omitempty" xml:"BaseFile,omitempty"`
	CycleOnce     *bool   `json:"CycleOnce,omitempty" xml:"CycleOnce,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileParamName *string `json:"FileParamName,omitempty" xml:"FileParamName,omitempty"`
}

func (s GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetBaseFile(v bool) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.BaseFile = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetCycleOnce(v bool) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.CycleOnce = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetFileName(v string) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.FileName = &v
	return s
}

func (s *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList) SetFileParamName(v string) *GetPtsSceneResponseBodySceneRelationListFileParameterExplainList {
	s.FileParamName = &v
	return s
}

type GetPtsSceneResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneResponse) SetHeaders(v map[string]*string) *GetPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneResponse) SetBody(v *GetPtsSceneResponseBody) *GetPtsSceneResponse {
	s.Body = v
	return s
}

type GetPtsSceneBaseLineRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneBaseLineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineRequest) SetSceneId(v string) *GetPtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

type GetPtsSceneBaseLineResponseBody struct {
	Baseline       *GetPtsSceneBaseLineResponseBodyBaseline `json:"Baseline,omitempty" xml:"Baseline,omitempty" type:"Struct"`
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneId        *string                                  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBody) SetBaseline(v *GetPtsSceneBaseLineResponseBodyBaseline) *GetPtsSceneBaseLineResponseBody {
	s.Baseline = v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetCode(v string) *GetPtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetMessage(v string) *GetPtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetRequestId(v string) *GetPtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetSceneId(v string) *GetPtsSceneBaseLineResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBody) SetSuccess(v bool) *GetPtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

type GetPtsSceneBaseLineResponseBodyBaseline struct {
	ApiBaselines  []*GetPtsSceneBaseLineResponseBodyBaselineApiBaselines `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty" type:"Repeated"`
	Name          *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	SceneBaseline *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline  `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty" type:"Struct"`
}

func (s GetPtsSceneBaseLineResponseBodyBaseline) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaseline) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetApiBaselines(v []*GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.ApiBaselines = v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetName(v string) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.Name = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaseline) SetSceneBaseline(v *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) *GetPtsSceneBaseLineResponseBodyBaseline {
	s.SceneBaseline = v
	return s
}

type GetPtsSceneBaseLineResponseBodyBaselineApiBaselines struct {
	AvgRt          *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgTps         *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	FailCountBiz   *int64   `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	FailCountReq   *int64   `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	Id             *int64   `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxRt          *int32   `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	MinRt          *int32   `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	Name           *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Seg90Rt        *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	Seg99Rt        *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetAvgRt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.AvgRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetAvgTps(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.AvgTps = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetFailCountBiz(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetFailCountReq(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetId(v int64) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Id = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetMaxRt(v int32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.MaxRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetMinRt(v int32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.MinRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetName(v string) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Name = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSeg90Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSeg99Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSuccessRateBiz(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines) SetSuccessRateReq(v float32) *GetPtsSceneBaseLineResponseBodyBaselineApiBaselines {
	s.SuccessRateReq = &v
	return s
}

type GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline struct {
	AvgRt          *float32 `json:"AvgRt,omitempty" xml:"AvgRt,omitempty"`
	AvgTps         *float32 `json:"AvgTps,omitempty" xml:"AvgTps,omitempty"`
	FailCountBiz   *int64   `json:"FailCountBiz,omitempty" xml:"FailCountBiz,omitempty"`
	FailCountReq   *int64   `json:"FailCountReq,omitempty" xml:"FailCountReq,omitempty"`
	Seg90Rt        *float32 `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	Seg99Rt        *float32 `json:"Seg99Rt,omitempty" xml:"Seg99Rt,omitempty"`
	SuccessRateBiz *float32 `json:"SuccessRateBiz,omitempty" xml:"SuccessRateBiz,omitempty"`
	SuccessRateReq *float32 `json:"SuccessRateReq,omitempty" xml:"SuccessRateReq,omitempty"`
}

func (s GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetAvgRt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.AvgRt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetAvgTps(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.AvgTps = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetFailCountBiz(v int64) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.FailCountBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetFailCountReq(v int64) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.FailCountReq = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSeg90Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.Seg90Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSeg99Rt(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.Seg99Rt = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSuccessRateBiz(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.SuccessRateBiz = &v
	return s
}

func (s *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline) SetSuccessRateReq(v float32) *GetPtsSceneBaseLineResponseBodyBaselineSceneBaseline {
	s.SuccessRateReq = &v
	return s
}

type GetPtsSceneBaseLineResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsSceneBaseLineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *GetPtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneBaseLineResponse) SetBody(v *GetPtsSceneBaseLineResponseBody) *GetPtsSceneBaseLineResponse {
	s.Body = v
	return s
}

type GetPtsSceneRunningDataRequest struct {
	PlanId  *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRunningDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataRequest) SetPlanId(v string) *GetPtsSceneRunningDataRequest {
	s.PlanId = &v
	return s
}

func (s *GetPtsSceneRunningDataRequest) SetSceneId(v string) *GetPtsSceneRunningDataRequest {
	s.SceneId = &v
	return s
}

type GetPtsSceneRunningDataResponseBody struct {
	AgentLocation        []*GetPtsSceneRunningDataResponseBodyAgentLocation        `json:"AgentLocation,omitempty" xml:"AgentLocation,omitempty" type:"Repeated"`
	AliveAgents          *int32                                                    `json:"AliveAgents,omitempty" xml:"AliveAgents,omitempty"`
	AverageRt            *int64                                                    `json:"AverageRt,omitempty" xml:"AverageRt,omitempty"`
	BeginTime            *int64                                                    `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	ChainMonitorDataList []*GetPtsSceneRunningDataResponseBodyChainMonitorDataList `json:"ChainMonitorDataList,omitempty" xml:"ChainMonitorDataList,omitempty" type:"Repeated"`
	Code                 *string                                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Concurrency          *int32                                                    `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ConcurrencyLimit     *int32                                                    `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	FailedBusinessCount  *int64                                                    `json:"FailedBusinessCount,omitempty" xml:"FailedBusinessCount,omitempty"`
	FailedRequestCount   *int64                                                    `json:"FailedRequestCount,omitempty" xml:"FailedRequestCount,omitempty"`
	HasReport            *bool                                                     `json:"HasReport,omitempty" xml:"HasReport,omitempty"`
	HttpStatusCode       *int32                                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message              *string                                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestBps           *string                                                   `json:"RequestBps,omitempty" xml:"RequestBps,omitempty"`
	RequestId            *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseBps          *string                                                   `json:"ResponseBps,omitempty" xml:"ResponseBps,omitempty"`
	Seg90Rt              *int64                                                    `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	Status               *int32                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Success              *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalAgents          *int32                                                    `json:"TotalAgents,omitempty" xml:"TotalAgents,omitempty"`
	TotalRequestCount    *int64                                                    `json:"TotalRequestCount,omitempty" xml:"TotalRequestCount,omitempty"`
	TpsLimit             *int32                                                    `json:"TpsLimit,omitempty" xml:"TpsLimit,omitempty"`
	Vum                  *int64                                                    `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBody) GoString() string {
	return s.String()
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

type GetPtsSceneRunningDataResponseBodyAgentLocation struct {
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyAgentLocation) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyAgentLocation) GoString() string {
	return s.String()
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

type GetPtsSceneRunningDataResponseBodyChainMonitorDataList struct {
	ApiId            *string                                                                 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName          *string                                                                 `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AverageRt        *int32                                                                  `json:"AverageRt,omitempty" xml:"AverageRt,omitempty"`
	CheckPointResult *GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult `json:"CheckPointResult,omitempty" xml:"CheckPointResult,omitempty" type:"Struct"`
	Concurrency      *float32                                                                `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ConfigQps        *int32                                                                  `json:"ConfigQps,omitempty" xml:"ConfigQps,omitempty"`
	Count2XX         *int64                                                                  `json:"Count2XX,omitempty" xml:"Count2XX,omitempty"`
	FailedCount      *int64                                                                  `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	FailedQps        *float32                                                                `json:"FailedQps,omitempty" xml:"FailedQps,omitempty"`
	MaxRt            *int32                                                                  `json:"MaxRt,omitempty" xml:"MaxRt,omitempty"`
	MinRt            *int32                                                                  `json:"MinRt,omitempty" xml:"MinRt,omitempty"`
	NodeId           *int64                                                                  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	Qps2XX           *float32                                                                `json:"Qps2XX,omitempty" xml:"Qps2XX,omitempty"`
	RealQps          *float32                                                                `json:"RealQps,omitempty" xml:"RealQps,omitempty"`
	TimePoint        *int64                                                                  `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataList) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataList) GoString() string {
	return s.String()
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

type GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult struct {
	FailedBusinessCount  *int64   `json:"FailedBusinessCount,omitempty" xml:"FailedBusinessCount,omitempty"`
	FailedBusinessQps    *float32 `json:"FailedBusinessQps,omitempty" xml:"FailedBusinessQps,omitempty"`
	SucceedBusinessCount *int64   `json:"SucceedBusinessCount,omitempty" xml:"SucceedBusinessCount,omitempty"`
	SucceedBusinessQps   *float32 `json:"SucceedBusinessQps,omitempty" xml:"SucceedBusinessQps,omitempty"`
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataResponseBodyChainMonitorDataListCheckPointResult) GoString() string {
	return s.String()
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

type GetPtsSceneRunningDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsSceneRunningDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsSceneRunningDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningDataResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataResponse) SetHeaders(v map[string]*string) *GetPtsSceneRunningDataResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneRunningDataResponse) SetBody(v *GetPtsSceneRunningDataResponseBody) *GetPtsSceneRunningDataResponse {
	s.Body = v
	return s
}

type GetPtsSceneRunningStatusRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRunningStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningStatusRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusRequest) SetSceneId(v string) *GetPtsSceneRunningStatusRequest {
	s.SceneId = &v
	return s
}

type GetPtsSceneRunningStatusResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ModifiedTime   *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneName      *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPtsSceneRunningStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusResponseBody) SetCode(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetCreateTime(v string) *GetPtsSceneRunningStatusResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetHttpStatusCode(v int32) *GetPtsSceneRunningStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetMessage(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetModifiedTime(v string) *GetPtsSceneRunningStatusResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetRequestId(v string) *GetPtsSceneRunningStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetSceneName(v string) *GetPtsSceneRunningStatusResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetStatus(v string) *GetPtsSceneRunningStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetPtsSceneRunningStatusResponseBody) SetSuccess(v bool) *GetPtsSceneRunningStatusResponseBody {
	s.Success = &v
	return s
}

type GetPtsSceneRunningStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPtsSceneRunningStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPtsSceneRunningStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPtsSceneRunningStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningStatusResponse) SetHeaders(v map[string]*string) *GetPtsSceneRunningStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneRunningStatusResponse) SetBody(v *GetPtsSceneRunningStatusResponseBody) *GetPtsSceneRunningStatusResponse {
	s.Body = v
	return s
}

type ListEnvsRequest struct {
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// 环境名
	EnvName    *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEnvsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvsRequest) SetEnvId(v string) *ListEnvsRequest {
	s.EnvId = &v
	return s
}

func (s *ListEnvsRequest) SetEnvName(v string) *ListEnvsRequest {
	s.EnvName = &v
	return s
}

func (s *ListEnvsRequest) SetPageNumber(v int32) *ListEnvsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvsRequest) SetPageSize(v int32) *ListEnvsRequest {
	s.PageSize = &v
	return s
}

type ListEnvsResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 环境列表
	Envs           []*ListEnvsResponseBodyEnvs `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEnvsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBody) SetCode(v string) *ListEnvsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvsResponseBody) SetEnvs(v []*ListEnvsResponseBodyEnvs) *ListEnvsResponseBody {
	s.Envs = v
	return s
}

func (s *ListEnvsResponseBody) SetHttpStatusCode(v int32) *ListEnvsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEnvsResponseBody) SetMessage(v string) *ListEnvsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvsResponseBody) SetPageNumber(v int32) *ListEnvsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEnvsResponseBody) SetPageSize(v int32) *ListEnvsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEnvsResponseBody) SetRequestId(v string) *ListEnvsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvsResponseBody) SetSuccess(v bool) *ListEnvsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEnvsResponseBody) SetTotalCount(v int64) *ListEnvsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEnvsResponseBodyEnvs struct {
	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 环境ID
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// 环境名
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// 依赖的jmeter版本
	EnvVersion *string `json:"EnvVersion,omitempty" xml:"EnvVersion,omitempty"`
	// 包含的jar包
	Files []*ListEnvsResponseBodyEnvsFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// 修改时间
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// jmeter属性
	Properties []*ListEnvsResponseBodyEnvsProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// 关联的场景
	RelatedScenes []*string `json:"RelatedScenes,omitempty" xml:"RelatedScenes,omitempty" type:"Repeated"`
	// 关联的场景id
	RunningScenes []*string `json:"RunningScenes,omitempty" xml:"RunningScenes,omitempty" type:"Repeated"`
	// 环境的文件总大小
	UsedCapacity *int64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s ListEnvsResponseBodyEnvs) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsResponseBodyEnvs) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvs) SetCreateTime(v int64) *ListEnvsResponseBodyEnvs {
	s.CreateTime = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvId(v string) *ListEnvsResponseBodyEnvs {
	s.EnvId = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvName(v string) *ListEnvsResponseBodyEnvs {
	s.EnvName = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetEnvVersion(v string) *ListEnvsResponseBodyEnvs {
	s.EnvVersion = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetFiles(v []*ListEnvsResponseBodyEnvsFiles) *ListEnvsResponseBodyEnvs {
	s.Files = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetModifiedTime(v int64) *ListEnvsResponseBodyEnvs {
	s.ModifiedTime = &v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetProperties(v []*ListEnvsResponseBodyEnvsProperties) *ListEnvsResponseBodyEnvs {
	s.Properties = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetRelatedScenes(v []*string) *ListEnvsResponseBodyEnvs {
	s.RelatedScenes = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetRunningScenes(v []*string) *ListEnvsResponseBodyEnvs {
	s.RunningScenes = v
	return s
}

func (s *ListEnvsResponseBodyEnvs) SetUsedCapacity(v int64) *ListEnvsResponseBodyEnvs {
	s.UsedCapacity = &v
	return s
}

type ListEnvsResponseBodyEnvsFiles struct {
	// 文件ID
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件的oss地址
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// 文件大小，单位为Byte
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// jar包的md5值
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s ListEnvsResponseBodyEnvsFiles) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsResponseBodyEnvsFiles) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileId(v int64) *ListEnvsResponseBodyEnvsFiles {
	s.FileId = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileName(v string) *ListEnvsResponseBodyEnvsFiles {
	s.FileName = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileOssAddress(v string) *ListEnvsResponseBodyEnvsFiles {
	s.FileOssAddress = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetFileSize(v int64) *ListEnvsResponseBodyEnvsFiles {
	s.FileSize = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsFiles) SetMd5(v string) *ListEnvsResponseBodyEnvsFiles {
	s.Md5 = &v
	return s
}

type ListEnvsResponseBodyEnvsProperties struct {
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 属性名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 属性值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEnvsResponseBodyEnvsProperties) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsResponseBodyEnvsProperties) GoString() string {
	return s.String()
}

func (s *ListEnvsResponseBodyEnvsProperties) SetDescription(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Description = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsProperties) SetName(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Name = &v
	return s
}

func (s *ListEnvsResponseBodyEnvsProperties) SetValue(v string) *ListEnvsResponseBodyEnvsProperties {
	s.Value = &v
	return s
}

type ListEnvsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEnvsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEnvsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEnvsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvsResponse) SetHeaders(v map[string]*string) *ListEnvsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvsResponse) SetBody(v *ListEnvsResponseBody) *ListEnvsResponse {
	s.Body = v
	return s
}

type ListJMeterReportsRequest struct {
	// 报告的起始时间，单位为ms
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// 报告的结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 报告关键字
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 报告ID
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// 要查看的报告的场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListJMeterReportsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJMeterReportsRequest) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsRequest) SetBeginTime(v int64) *ListJMeterReportsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListJMeterReportsRequest) SetEndTime(v int64) *ListJMeterReportsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJMeterReportsRequest) SetKeyword(v string) *ListJMeterReportsRequest {
	s.Keyword = &v
	return s
}

func (s *ListJMeterReportsRequest) SetPageNumber(v int32) *ListJMeterReportsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJMeterReportsRequest) SetPageSize(v int32) *ListJMeterReportsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJMeterReportsRequest) SetReportId(v string) *ListJMeterReportsRequest {
	s.ReportId = &v
	return s
}

func (s *ListJMeterReportsRequest) SetSceneId(v string) *ListJMeterReportsRequest {
	s.SceneId = &v
	return s
}

type ListJMeterReportsResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Reports        []*ListJMeterReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJMeterReportsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJMeterReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponseBody) SetCode(v string) *ListJMeterReportsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetHttpStatusCode(v int32) *ListJMeterReportsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetMessage(v string) *ListJMeterReportsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetPageNumber(v int32) *ListJMeterReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetPageSize(v int32) *ListJMeterReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetReports(v []*ListJMeterReportsResponseBodyReports) *ListJMeterReportsResponseBody {
	s.Reports = v
	return s
}

func (s *ListJMeterReportsResponseBody) SetRequestId(v string) *ListJMeterReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetSuccess(v bool) *ListJMeterReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetTotalCount(v int64) *ListJMeterReportsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJMeterReportsResponseBodyReports struct {
	// 压测开始时间
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// 压测持续时间
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 报告id
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// 报告名称
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// 消耗的vum
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s ListJMeterReportsResponseBodyReports) String() string {
	return tea.Prettify(s)
}

func (s ListJMeterReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponseBodyReports) SetActualStartTime(v int64) *ListJMeterReportsResponseBodyReports {
	s.ActualStartTime = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetDuration(v string) *ListJMeterReportsResponseBodyReports {
	s.Duration = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetReportId(v string) *ListJMeterReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetReportName(v string) *ListJMeterReportsResponseBodyReports {
	s.ReportName = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetVum(v int64) *ListJMeterReportsResponseBodyReports {
	s.Vum = &v
	return s
}

type ListJMeterReportsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJMeterReportsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJMeterReportsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJMeterReportsResponse) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponse) SetHeaders(v map[string]*string) *ListJMeterReportsResponse {
	s.Headers = v
	return s
}

func (s *ListJMeterReportsResponse) SetBody(v *ListJMeterReportsResponseBody) *ListJMeterReportsResponse {
	s.Body = v
	return s
}

type ListOpenJMeterScenesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ListOpenJMeterScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpenJMeterScenesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesRequest) SetPageNumber(v int32) *ListOpenJMeterScenesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetPageSize(v int32) *ListOpenJMeterScenesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetSceneId(v string) *ListOpenJMeterScenesRequest {
	s.SceneId = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetSceneName(v string) *ListOpenJMeterScenesRequest {
	s.SceneName = &v
	return s
}

type ListOpenJMeterScenesResponseBody struct {
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JMeterScene    []*ListOpenJMeterScenesResponseBodyJMeterScene `json:"JMeterScene,omitempty" xml:"JMeterScene,omitempty" type:"Repeated"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber     *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount     *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpenJMeterScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpenJMeterScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponseBody) SetCode(v string) *ListOpenJMeterScenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetHttpStatusCode(v int32) *ListOpenJMeterScenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetJMeterScene(v []*ListOpenJMeterScenesResponseBodyJMeterScene) *ListOpenJMeterScenesResponseBody {
	s.JMeterScene = v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetMessage(v string) *ListOpenJMeterScenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetPageNumber(v int32) *ListOpenJMeterScenesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetPageSize(v int32) *ListOpenJMeterScenesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetRequestId(v string) *ListOpenJMeterScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetSuccess(v bool) *ListOpenJMeterScenesResponseBody {
	s.Success = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBody) SetTotalCount(v int64) *ListOpenJMeterScenesResponseBody {
	s.TotalCount = &v
	return s
}

type ListOpenJMeterScenesResponseBodyJMeterScene struct {
	// 压测持续时间
	DurationStr *string `json:"DurationStr,omitempty" xml:"DurationStr,omitempty"`
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ListOpenJMeterScenesResponseBodyJMeterScene) String() string {
	return tea.Prettify(s)
}

func (s ListOpenJMeterScenesResponseBodyJMeterScene) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetDurationStr(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.DurationStr = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetSceneId(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.SceneId = &v
	return s
}

func (s *ListOpenJMeterScenesResponseBodyJMeterScene) SetSceneName(v string) *ListOpenJMeterScenesResponseBodyJMeterScene {
	s.SceneName = &v
	return s
}

type ListOpenJMeterScenesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOpenJMeterScenesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOpenJMeterScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpenJMeterScenesResponse) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesResponse) SetHeaders(v map[string]*string) *ListOpenJMeterScenesResponse {
	s.Headers = v
	return s
}

func (s *ListOpenJMeterScenesResponse) SetBody(v *ListOpenJMeterScenesResponseBody) *ListOpenJMeterScenesResponse {
	s.Body = v
	return s
}

type ListPtsSceneRequest struct {
	KeyWord    *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *ListPtsSceneRequest) SetKeyWord(v string) *ListPtsSceneRequest {
	s.KeyWord = &v
	return s
}

func (s *ListPtsSceneRequest) SetPageNumber(v int32) *ListPtsSceneRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPtsSceneRequest) SetPageSize(v int32) *ListPtsSceneRequest {
	s.PageSize = &v
	return s
}

type ListPtsSceneResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SceneViewList  []*ListPtsSceneResponseBodySceneViewList `json:"SceneViewList,omitempty" xml:"SceneViewList,omitempty" type:"Repeated"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponseBody) SetCode(v string) *ListPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetHttpStatusCode(v int32) *ListPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetMessage(v string) *ListPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetRequestId(v string) *ListPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPtsSceneResponseBody) SetSceneViewList(v []*ListPtsSceneResponseBodySceneViewList) *ListPtsSceneResponseBody {
	s.SceneViewList = v
	return s
}

func (s *ListPtsSceneResponseBody) SetSuccess(v bool) *ListPtsSceneResponseBody {
	s.Success = &v
	return s
}

type ListPtsSceneResponseBodySceneViewList struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	SceneId    *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName  *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ListPtsSceneResponseBodySceneViewList) String() string {
	return tea.Prettify(s)
}

func (s ListPtsSceneResponseBodySceneViewList) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponseBodySceneViewList) SetCreateTime(v string) *ListPtsSceneResponseBodySceneViewList {
	s.CreateTime = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) SetSceneId(v string) *ListPtsSceneResponseBodySceneViewList {
	s.SceneId = &v
	return s
}

func (s *ListPtsSceneResponseBodySceneViewList) SetSceneName(v string) *ListPtsSceneResponseBodySceneViewList {
	s.SceneName = &v
	return s
}

type ListPtsSceneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *ListPtsSceneResponse) SetHeaders(v map[string]*string) *ListPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *ListPtsSceneResponse) SetBody(v *ListPtsSceneResponseBody) *ListPtsSceneResponse {
	s.Body = v
	return s
}

type ModifyPtsSceneRequest struct {
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ModifyPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneRequest) SetScene(v string) *ModifyPtsSceneRequest {
	s.Scene = &v
	return s
}

type ModifyPtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneResponseBody) SetCode(v string) *ModifyPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetHttpStatusCode(v int32) *ModifyPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetMessage(v string) *ModifyPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetRequestId(v string) *ModifyPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPtsSceneResponseBody) SetSuccess(v bool) *ModifyPtsSceneResponseBody {
	s.Success = &v
	return s
}

type ModifyPtsSceneResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *ModifyPtsSceneResponse) SetHeaders(v map[string]*string) *ModifyPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *ModifyPtsSceneResponse) SetBody(v *ModifyPtsSceneResponseBody) *ModifyPtsSceneResponse {
	s.Body = v
	return s
}

type RemoveEnvRequest struct {
	// 要删除的环境ID
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
}

func (s RemoveEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveEnvRequest) GoString() string {
	return s.String()
}

func (s *RemoveEnvRequest) SetEnvId(v string) *RemoveEnvRequest {
	s.EnvId = &v
	return s
}

type RemoveEnvResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveEnvResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEnvResponseBody) SetCode(v string) *RemoveEnvResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveEnvResponseBody) SetHttpStatusCode(v int32) *RemoveEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveEnvResponseBody) SetMessage(v string) *RemoveEnvResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveEnvResponseBody) SetRequestId(v string) *RemoveEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEnvResponseBody) SetSuccess(v bool) *RemoveEnvResponseBody {
	s.Success = &v
	return s
}

type RemoveEnvResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveEnvResponse) GoString() string {
	return s.String()
}

func (s *RemoveEnvResponse) SetHeaders(v map[string]*string) *RemoveEnvResponse {
	s.Headers = v
	return s
}

func (s *RemoveEnvResponse) SetBody(v *RemoveEnvResponseBody) *RemoveEnvResponse {
	s.Body = v
	return s
}

type RemoveOpenJMeterSceneRequest struct {
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s RemoveOpenJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneRequest) SetSceneId(v string) *RemoveOpenJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type RemoveOpenJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveOpenJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneResponseBody) SetCode(v string) *RemoveOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *RemoveOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetMessage(v string) *RemoveOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetRequestId(v string) *RemoveOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetSuccess(v bool) *RemoveOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type RemoveOpenJMeterSceneResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveOpenJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *RemoveOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *RemoveOpenJMeterSceneResponse) SetBody(v *RemoveOpenJMeterSceneResponseBody) *RemoveOpenJMeterSceneResponse {
	s.Body = v
	return s
}

type SaveEnvRequest struct {
	// 环境
	Env *SaveEnvRequestEnv `json:"Env,omitempty" xml:"Env,omitempty" type:"Struct"`
}

func (s SaveEnvRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvRequest) GoString() string {
	return s.String()
}

func (s *SaveEnvRequest) SetEnv(v *SaveEnvRequestEnv) *SaveEnvRequest {
	s.Env = v
	return s
}

type SaveEnvRequestEnv struct {
	// 环境id，不填表示新建环境，填了表示修改该环境
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// 环境名称
	EnvName *string `json:"EnvName,omitempty" xml:"EnvName,omitempty"`
	// 环境依赖的文件
	Files []*SaveEnvRequestEnvFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// jmeter插件的环境标签
	JmeterPluginLabel *string `json:"JmeterPluginLabel,omitempty" xml:"JmeterPluginLabel,omitempty"`
	// jmeter属性
	Properties []*SaveEnvRequestEnvProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s SaveEnvRequestEnv) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvRequestEnv) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnv) SetEnvId(v string) *SaveEnvRequestEnv {
	s.EnvId = &v
	return s
}

func (s *SaveEnvRequestEnv) SetEnvName(v string) *SaveEnvRequestEnv {
	s.EnvName = &v
	return s
}

func (s *SaveEnvRequestEnv) SetFiles(v []*SaveEnvRequestEnvFiles) *SaveEnvRequestEnv {
	s.Files = v
	return s
}

func (s *SaveEnvRequestEnv) SetJmeterPluginLabel(v string) *SaveEnvRequestEnv {
	s.JmeterPluginLabel = &v
	return s
}

func (s *SaveEnvRequestEnv) SetProperties(v []*SaveEnvRequestEnvProperties) *SaveEnvRequestEnv {
	s.Properties = v
	return s
}

type SaveEnvRequestEnvFiles struct {
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件oss地址，目前只支持上海region的oss地址
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
}

func (s SaveEnvRequestEnvFiles) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvRequestEnvFiles) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnvFiles) SetFileName(v string) *SaveEnvRequestEnvFiles {
	s.FileName = &v
	return s
}

func (s *SaveEnvRequestEnvFiles) SetFileOssAddress(v string) *SaveEnvRequestEnvFiles {
	s.FileOssAddress = &v
	return s
}

type SaveEnvRequestEnvProperties struct {
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 属性名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 属性值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveEnvRequestEnvProperties) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvRequestEnvProperties) GoString() string {
	return s.String()
}

func (s *SaveEnvRequestEnvProperties) SetDescription(v string) *SaveEnvRequestEnvProperties {
	s.Description = &v
	return s
}

func (s *SaveEnvRequestEnvProperties) SetName(v string) *SaveEnvRequestEnvProperties {
	s.Name = &v
	return s
}

func (s *SaveEnvRequestEnvProperties) SetValue(v string) *SaveEnvRequestEnvProperties {
	s.Value = &v
	return s
}

type SaveEnvShrinkRequest struct {
	// 环境
	EnvShrink *string `json:"Env,omitempty" xml:"Env,omitempty"`
}

func (s SaveEnvShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveEnvShrinkRequest) SetEnvShrink(v string) *SaveEnvShrinkRequest {
	s.EnvShrink = &v
	return s
}

type SaveEnvResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 操作的环境id
	EnvId          *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveEnvResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvResponseBody) GoString() string {
	return s.String()
}

func (s *SaveEnvResponseBody) SetCode(v string) *SaveEnvResponseBody {
	s.Code = &v
	return s
}

func (s *SaveEnvResponseBody) SetEnvId(v string) *SaveEnvResponseBody {
	s.EnvId = &v
	return s
}

func (s *SaveEnvResponseBody) SetHttpStatusCode(v int32) *SaveEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveEnvResponseBody) SetMessage(v string) *SaveEnvResponseBody {
	s.Message = &v
	return s
}

func (s *SaveEnvResponseBody) SetRequestId(v string) *SaveEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveEnvResponseBody) SetSuccess(v bool) *SaveEnvResponseBody {
	s.Success = &v
	return s
}

type SaveEnvResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveEnvResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveEnvResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveEnvResponse) GoString() string {
	return s.String()
}

func (s *SaveEnvResponse) SetHeaders(v map[string]*string) *SaveEnvResponse {
	s.Headers = v
	return s
}

func (s *SaveEnvResponse) SetBody(v *SaveEnvResponseBody) *SaveEnvResponse {
	s.Body = v
	return s
}

type SaveOpenJMeterSceneRequest struct {
	// 场景详情
	OpenJMeterScene *SaveOpenJMeterSceneRequestOpenJMeterScene `json:"OpenJMeterScene,omitempty" xml:"OpenJMeterScene,omitempty" type:"Struct"`
}

func (s SaveOpenJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequest) SetOpenJMeterScene(v *SaveOpenJMeterSceneRequestOpenJMeterScene) *SaveOpenJMeterSceneRequest {
	s.OpenJMeterScene = v
	return s
}

type SaveOpenJMeterSceneRequestOpenJMeterScene struct {
	// 施压引擎数量
	AgentCount *int32 `json:"AgentCount,omitempty" xml:"AgentCount,omitempty"`
	// 最大并发
	Concurrency *int32 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// constantThroughputTimerType
	ConstantThroughputTimerType *string `json:"ConstantThroughputTimerType,omitempty" xml:"ConstantThroughputTimerType,omitempty"`
	// DNS配置
	DnsCacheConfig *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig `json:"DnsCacheConfig,omitempty" xml:"DnsCacheConfig,omitempty" type:"Struct"`
	// 压测持续时间
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 关联的环境id
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// 文件列表
	FileList []*SaveOpenJMeterSceneRequestOpenJMeterSceneFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// 是否为VPC测试，默认为false表示公网测试，此值为true时VPC相关配置才生效
	IsVpcTest *bool `json:"IsVpcTest,omitempty" xml:"IsVpcTest,omitempty"`
	// Jmeter属性
	JMeterProperties []*SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties `json:"JMeterProperties,omitempty" xml:"JMeterProperties,omitempty" type:"Repeated"`
	// jmeter插件的环境标签
	JmeterPluginLabel *string `json:"JmeterPluginLabel,omitempty" xml:"JmeterPluginLabel,omitempty"`
	// 预热时间
	RampUp *int32 `json:"RampUp,omitempty" xml:"RampUp,omitempty"`
	// region的id，VPC压测时配置
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 场景ID
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// 场景名
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// 安全组id，VPC压测时配置
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 预热阶段
	Steps *int32 `json:"Steps,omitempty" xml:"Steps,omitempty"`
	// synchronizing timer 类型
	SyncTimerType *string `json:"SyncTimerType,omitempty" xml:"SyncTimerType,omitempty"`
	// 测试文件
	TestFile *string `json:"TestFile,omitempty" xml:"TestFile,omitempty"`
	// 交换机id，VPC压测时配置
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// vpc的id，VPC压测时配置
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterScene) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterScene) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetAgentCount(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.AgentCount = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetConcurrency(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Concurrency = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetConstantThroughputTimerType(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.ConstantThroughputTimerType = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetDnsCacheConfig(v *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.DnsCacheConfig = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetDuration(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Duration = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetEnvironmentId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.EnvironmentId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetFileList(v []*SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.FileList = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetIsVpcTest(v bool) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.IsVpcTest = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetJMeterProperties(v []*SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.JMeterProperties = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetJmeterPluginLabel(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.JmeterPluginLabel = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetRampUp(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.RampUp = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetRegionId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.RegionId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSceneId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SceneId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSceneName(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SceneName = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSecurityGroupId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SecurityGroupId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSteps(v int32) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.Steps = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetSyncTimerType(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.SyncTimerType = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetTestFile(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.TestFile = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetVSwitchId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.VSwitchId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterScene) SetVpcId(v string) *SaveOpenJMeterSceneRequestOpenJMeterScene {
	s.VpcId = &v
	return s
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig struct {
	ClearCacheEachIteration *bool              `json:"ClearCacheEachIteration,omitempty" xml:"ClearCacheEachIteration,omitempty"`
	DnsServers              []*string          `json:"DnsServers,omitempty" xml:"DnsServers,omitempty" type:"Repeated"`
	HostTable               map[string]*string `json:"HostTable,omitempty" xml:"HostTable,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetClearCacheEachIteration(v bool) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.ClearCacheEachIteration = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetDnsServers(v []*string) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.DnsServers = v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig) SetHostTable(v map[string]*string) *SaveOpenJMeterSceneRequestOpenJMeterSceneDnsCacheConfig {
	s.HostTable = v
	return s
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneFileList struct {
	// 文件id
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// 文件公网可访问的oss地址
	FileOssAddress *string `json:"FileOssAddress,omitempty" xml:"FileOssAddress,omitempty"`
	// 文件大小，单位byte
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// 文件的MD5
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// 是否切分，仅针对csv有效
	SplitCsv *bool `json:"SplitCsv,omitempty" xml:"SplitCsv,omitempty"`
	// 文件tag
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileId(v int64) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileId = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileName(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileName = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileOssAddress(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileOssAddress = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetFileSize(v int64) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.FileSize = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetMd5(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.Md5 = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetSplitCsv(v bool) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.SplitCsv = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList) SetTags(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneFileList {
	s.Tags = &v
	return s
}

type SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) SetName(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties {
	s.Name = &v
	return s
}

func (s *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties) SetValue(v string) *SaveOpenJMeterSceneRequestOpenJMeterSceneJMeterProperties {
	s.Value = &v
	return s
}

type SaveOpenJMeterSceneShrinkRequest struct {
	// 场景详情
	OpenJMeterSceneShrink *string `json:"OpenJMeterScene,omitempty" xml:"OpenJMeterScene,omitempty"`
}

func (s SaveOpenJMeterSceneShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneShrinkRequest) SetOpenJMeterSceneShrink(v string) *SaveOpenJMeterSceneShrinkRequest {
	s.OpenJMeterSceneShrink = &v
	return s
}

type SaveOpenJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveOpenJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneResponseBody) SetCode(v string) *SaveOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *SaveOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetMessage(v string) *SaveOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetRequestId(v string) *SaveOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetSceneId(v string) *SaveOpenJMeterSceneResponseBody {
	s.SceneId = &v
	return s
}

func (s *SaveOpenJMeterSceneResponseBody) SetSuccess(v bool) *SaveOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type SaveOpenJMeterSceneResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveOpenJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *SaveOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *SaveOpenJMeterSceneResponse) SetBody(v *SaveOpenJMeterSceneResponseBody) *SaveOpenJMeterSceneResponse {
	s.Body = v
	return s
}

type StartDebugPtsSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartDebugPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDebugPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneRequest) SetSceneId(v string) *StartDebugPtsSceneRequest {
	s.SceneId = &v
	return s
}

type StartDebugPtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PlanId         *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDebugPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDebugPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneResponseBody) SetCode(v string) *StartDebugPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetHttpStatusCode(v int32) *StartDebugPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetMessage(v string) *StartDebugPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetPlanId(v string) *StartDebugPtsSceneResponseBody {
	s.PlanId = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetRequestId(v string) *StartDebugPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDebugPtsSceneResponseBody) SetSuccess(v bool) *StartDebugPtsSceneResponseBody {
	s.Success = &v
	return s
}

type StartDebugPtsSceneResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDebugPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDebugPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDebugPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneResponse) SetHeaders(v map[string]*string) *StartDebugPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StartDebugPtsSceneResponse) SetBody(v *StartDebugPtsSceneResponseBody) *StartDebugPtsSceneResponse {
	s.Body = v
	return s
}

type StartDebuggingJMeterSceneRequest struct {
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartDebuggingJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDebuggingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneRequest) SetSceneId(v string) *StartDebuggingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type StartDebuggingJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReportId       *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDebuggingJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDebuggingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneResponseBody) SetCode(v string) *StartDebuggingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StartDebuggingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetMessage(v string) *StartDebuggingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetReportId(v string) *StartDebuggingJMeterSceneResponseBody {
	s.ReportId = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetRequestId(v string) *StartDebuggingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetSuccess(v bool) *StartDebuggingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type StartDebuggingJMeterSceneResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDebuggingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDebuggingJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDebuggingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneResponse) SetHeaders(v map[string]*string) *StartDebuggingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StartDebuggingJMeterSceneResponse) SetBody(v *StartDebuggingJMeterSceneResponseBody) *StartDebuggingJMeterSceneResponse {
	s.Body = v
	return s
}

type StartPtsSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StartPtsSceneRequest) SetSceneId(v string) *StartPtsSceneRequest {
	s.SceneId = &v
	return s
}

type StartPtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PlanId         *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartPtsSceneResponseBody) SetCode(v string) *StartPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetHttpStatusCode(v int32) *StartPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetMessage(v string) *StartPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetPlanId(v string) *StartPtsSceneResponseBody {
	s.PlanId = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetRequestId(v string) *StartPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPtsSceneResponseBody) SetSuccess(v bool) *StartPtsSceneResponseBody {
	s.Success = &v
	return s
}

type StartPtsSceneResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StartPtsSceneResponse) SetHeaders(v map[string]*string) *StartPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StartPtsSceneResponse) SetBody(v *StartPtsSceneResponseBody) *StartPtsSceneResponse {
	s.Body = v
	return s
}

type StartTestingJMeterSceneRequest struct {
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartTestingJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTestingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneRequest) SetSceneId(v string) *StartTestingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type StartTestingJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReportId       *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartTestingJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTestingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneResponseBody) SetCode(v string) *StartTestingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StartTestingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetMessage(v string) *StartTestingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetReportId(v string) *StartTestingJMeterSceneResponseBody {
	s.ReportId = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetRequestId(v string) *StartTestingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetSuccess(v bool) *StartTestingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type StartTestingJMeterSceneResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTestingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTestingJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTestingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneResponse) SetHeaders(v map[string]*string) *StartTestingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StartTestingJMeterSceneResponse) SetBody(v *StartTestingJMeterSceneResponseBody) *StartTestingJMeterSceneResponse {
	s.Body = v
	return s
}

type StopDebugPtsSceneRequest struct {
	PlanId  *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopDebugPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDebugPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneRequest) SetPlanId(v string) *StopDebugPtsSceneRequest {
	s.PlanId = &v
	return s
}

func (s *StopDebugPtsSceneRequest) SetSceneId(v string) *StopDebugPtsSceneRequest {
	s.SceneId = &v
	return s
}

type StopDebugPtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDebugPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDebugPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneResponseBody) SetCode(v string) *StopDebugPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetHttpStatusCode(v int32) *StopDebugPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetMessage(v string) *StopDebugPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetRequestId(v string) *StopDebugPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDebugPtsSceneResponseBody) SetSuccess(v bool) *StopDebugPtsSceneResponseBody {
	s.Success = &v
	return s
}

type StopDebugPtsSceneResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDebugPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDebugPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDebugPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneResponse) SetHeaders(v map[string]*string) *StopDebugPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StopDebugPtsSceneResponse) SetBody(v *StopDebugPtsSceneResponseBody) *StopDebugPtsSceneResponse {
	s.Body = v
	return s
}

type StopDebuggingJMeterSceneRequest struct {
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopDebuggingJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDebuggingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneRequest) SetSceneId(v string) *StopDebuggingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type StopDebuggingJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDebuggingJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDebuggingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneResponseBody) SetCode(v string) *StopDebuggingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StopDebuggingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetMessage(v string) *StopDebuggingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetRequestId(v string) *StopDebuggingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponseBody) SetSuccess(v bool) *StopDebuggingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type StopDebuggingJMeterSceneResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDebuggingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDebuggingJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDebuggingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneResponse) SetHeaders(v map[string]*string) *StopDebuggingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StopDebuggingJMeterSceneResponse) SetBody(v *StopDebuggingJMeterSceneResponseBody) *StopDebuggingJMeterSceneResponse {
	s.Body = v
	return s
}

type StopPtsSceneRequest struct {
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopPtsSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *StopPtsSceneRequest) SetSceneId(v string) *StopPtsSceneRequest {
	s.SceneId = &v
	return s
}

type StopPtsSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopPtsSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopPtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopPtsSceneResponseBody) SetCode(v string) *StopPtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetHttpStatusCode(v int32) *StopPtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetMessage(v string) *StopPtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetRequestId(v string) *StopPtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopPtsSceneResponseBody) SetSuccess(v bool) *StopPtsSceneResponseBody {
	s.Success = &v
	return s
}

type StopPtsSceneResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopPtsSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StopPtsSceneResponse) SetHeaders(v map[string]*string) *StopPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StopPtsSceneResponse) SetBody(v *StopPtsSceneResponseBody) *StopPtsSceneResponse {
	s.Body = v
	return s
}

type StopTestingJMeterSceneRequest struct {
	// 场景id
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopTestingJMeterSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTestingJMeterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneRequest) SetSceneId(v string) *StopTestingJMeterSceneRequest {
	s.SceneId = &v
	return s
}

type StopTestingJMeterSceneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopTestingJMeterSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTestingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneResponseBody) SetCode(v string) *StopTestingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StopTestingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetMessage(v string) *StopTestingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetRequestId(v string) *StopTestingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTestingJMeterSceneResponseBody) SetSuccess(v bool) *StopTestingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

type StopTestingJMeterSceneResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopTestingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTestingJMeterSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTestingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneResponse) SetHeaders(v map[string]*string) *StopTestingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StopTestingJMeterSceneResponse) SetBody(v *StopTestingJMeterSceneResponseBody) *StopTestingJMeterSceneResponse {
	s.Body = v
	return s
}

type UpdatePtsSceneBaseLineRequest struct {
	ApiBaselines  map[string]interface{} `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty"`
	SceneBaseline map[string]interface{} `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty"`
	SceneId       *string                `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdatePtsSceneBaseLineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePtsSceneBaseLineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineRequest) SetApiBaselines(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest {
	s.ApiBaselines = v
	return s
}

func (s *UpdatePtsSceneBaseLineRequest) SetSceneBaseline(v map[string]interface{}) *UpdatePtsSceneBaseLineRequest {
	s.SceneBaseline = v
	return s
}

func (s *UpdatePtsSceneBaseLineRequest) SetSceneId(v string) *UpdatePtsSceneBaseLineRequest {
	s.SceneId = &v
	return s
}

type UpdatePtsSceneBaseLineShrinkRequest struct {
	ApiBaselinesShrink  *string `json:"ApiBaselines,omitempty" xml:"ApiBaselines,omitempty"`
	SceneBaselineShrink *string `json:"SceneBaseline,omitempty" xml:"SceneBaseline,omitempty"`
	SceneId             *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdatePtsSceneBaseLineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePtsSceneBaseLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetApiBaselinesShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.ApiBaselinesShrink = &v
	return s
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetSceneBaselineShrink(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.SceneBaselineShrink = &v
	return s
}

func (s *UpdatePtsSceneBaseLineShrinkRequest) SetSceneId(v string) *UpdatePtsSceneBaseLineShrinkRequest {
	s.SceneId = &v
	return s
}

type UpdatePtsSceneBaseLineResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePtsSceneBaseLineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePtsSceneBaseLineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetCode(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetHttpStatusCode(v int32) *UpdatePtsSceneBaseLineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetMessage(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetRequestId(v string) *UpdatePtsSceneBaseLineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponseBody) SetSuccess(v bool) *UpdatePtsSceneBaseLineResponseBody {
	s.Success = &v
	return s
}

type UpdatePtsSceneBaseLineResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdatePtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePtsSceneBaseLineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *UpdatePtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePtsSceneBaseLineResponse) SetBody(v *UpdatePtsSceneBaseLineResponseBody) *UpdatePtsSceneBaseLineResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePtsSceneWithOptions(request *CreatePtsSceneRequest, runtime *util.RuntimeOptions) (_result *CreatePtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePtsScene(request *CreatePtsSceneRequest) (_result *CreatePtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePtsSceneResponse{}
	_body, _err := client.CreatePtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePtsSceneBaseLineFromReportWithOptions(request *CreatePtsSceneBaseLineFromReportRequest, runtime *util.RuntimeOptions) (_result *CreatePtsSceneBaseLineFromReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePtsSceneBaseLineFromReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePtsSceneBaseLineFromReport"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePtsSceneBaseLineFromReport(request *CreatePtsSceneBaseLineFromReportRequest) (_result *CreatePtsSceneBaseLineFromReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePtsSceneBaseLineFromReportResponse{}
	_body, _err := client.CreatePtsSceneBaseLineFromReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePtsSceneWithOptions(request *DeletePtsSceneRequest, runtime *util.RuntimeOptions) (_result *DeletePtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePtsScene(request *DeletePtsSceneRequest) (_result *DeletePtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePtsSceneResponse{}
	_body, _err := client.DeletePtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePtsSceneBaseLineWithOptions(request *DeletePtsSceneBaseLineRequest, runtime *util.RuntimeOptions) (_result *DeletePtsSceneBaseLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePtsSceneBaseLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePtsSceneBaseLine"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePtsSceneBaseLine(request *DeletePtsSceneBaseLineRequest) (_result *DeletePtsSceneBaseLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePtsSceneBaseLineResponse{}
	_body, _err := client.DeletePtsSceneBaseLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePtsScenesWithOptions(tmpReq *DeletePtsScenesRequest, runtime *util.RuntimeOptions) (_result *DeletePtsScenesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeletePtsScenesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SceneIds)) {
		request.SceneIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneIds, tea.String("SceneIds"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePtsScenesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePtsScenes"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePtsScenes(request *DeletePtsScenesRequest) (_result *DeletePtsScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePtsScenesResponse{}
	_body, _err := client.DeletePtsScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJMeterLogsWithOptions(request *GetJMeterLogsRequest, runtime *util.RuntimeOptions) (_result *GetJMeterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJMeterLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJMeterLogs"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJMeterLogs(request *GetJMeterLogsRequest) (_result *GetJMeterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJMeterLogsResponse{}
	_body, _err := client.GetJMeterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJMeterSampleMetricsWithOptions(request *GetJMeterSampleMetricsRequest, runtime *util.RuntimeOptions) (_result *GetJMeterSampleMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJMeterSampleMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJMeterSampleMetrics"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJMeterSampleMetrics(request *GetJMeterSampleMetricsRequest) (_result *GetJMeterSampleMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJMeterSampleMetricsResponse{}
	_body, _err := client.GetJMeterSampleMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJMeterSamplingLogsWithOptions(request *GetJMeterSamplingLogsRequest, runtime *util.RuntimeOptions) (_result *GetJMeterSamplingLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJMeterSamplingLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJMeterSamplingLogs"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJMeterSamplingLogs(request *GetJMeterSamplingLogsRequest) (_result *GetJMeterSamplingLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJMeterSamplingLogsResponse{}
	_body, _err := client.GetJMeterSamplingLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJMeterSceneRunningDataWithOptions(request *GetJMeterSceneRunningDataRequest, runtime *util.RuntimeOptions) (_result *GetJMeterSceneRunningDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetJMeterSceneRunningDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetJMeterSceneRunningData"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJMeterSceneRunningData(request *GetJMeterSceneRunningDataRequest) (_result *GetJMeterSceneRunningDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetJMeterSceneRunningDataResponse{}
	_body, _err := client.GetJMeterSceneRunningDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOpenJMeterSceneWithOptions(request *GetOpenJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *GetOpenJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetOpenJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetOpenJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOpenJMeterScene(request *GetOpenJMeterSceneRequest) (_result *GetOpenJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpenJMeterSceneResponse{}
	_body, _err := client.GetOpenJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsReportDetailsWithOptions(request *GetPtsReportDetailsRequest, runtime *util.RuntimeOptions) (_result *GetPtsReportDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsReportDetailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsReportDetails"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsReportDetails(request *GetPtsReportDetailsRequest) (_result *GetPtsReportDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsReportDetailsResponse{}
	_body, _err := client.GetPtsReportDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsReportsBySceneIdWithOptions(request *GetPtsReportsBySceneIdRequest, runtime *util.RuntimeOptions) (_result *GetPtsReportsBySceneIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsReportsBySceneIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsReportsBySceneId"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsReportsBySceneId(request *GetPtsReportsBySceneIdRequest) (_result *GetPtsReportsBySceneIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsReportsBySceneIdResponse{}
	_body, _err := client.GetPtsReportsBySceneIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsSceneWithOptions(request *GetPtsSceneRequest, runtime *util.RuntimeOptions) (_result *GetPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsScene(request *GetPtsSceneRequest) (_result *GetPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsSceneResponse{}
	_body, _err := client.GetPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsSceneBaseLineWithOptions(request *GetPtsSceneBaseLineRequest, runtime *util.RuntimeOptions) (_result *GetPtsSceneBaseLineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsSceneBaseLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsSceneBaseLine"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsSceneBaseLine(request *GetPtsSceneBaseLineRequest) (_result *GetPtsSceneBaseLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsSceneBaseLineResponse{}
	_body, _err := client.GetPtsSceneBaseLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsSceneRunningDataWithOptions(request *GetPtsSceneRunningDataRequest, runtime *util.RuntimeOptions) (_result *GetPtsSceneRunningDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsSceneRunningDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsSceneRunningData"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsSceneRunningData(request *GetPtsSceneRunningDataRequest) (_result *GetPtsSceneRunningDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsSceneRunningDataResponse{}
	_body, _err := client.GetPtsSceneRunningDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPtsSceneRunningStatusWithOptions(request *GetPtsSceneRunningStatusRequest, runtime *util.RuntimeOptions) (_result *GetPtsSceneRunningStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPtsSceneRunningStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPtsSceneRunningStatus"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPtsSceneRunningStatus(request *GetPtsSceneRunningStatusRequest) (_result *GetPtsSceneRunningStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPtsSceneRunningStatusResponse{}
	_body, _err := client.GetPtsSceneRunningStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEnvsWithOptions(request *ListEnvsRequest, runtime *util.RuntimeOptions) (_result *ListEnvsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListEnvsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListEnvs"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEnvs(request *ListEnvsRequest) (_result *ListEnvsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEnvsResponse{}
	_body, _err := client.ListEnvsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJMeterReportsWithOptions(request *ListJMeterReportsRequest, runtime *util.RuntimeOptions) (_result *ListJMeterReportsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListJMeterReportsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJMeterReports"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJMeterReports(request *ListJMeterReportsRequest) (_result *ListJMeterReportsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJMeterReportsResponse{}
	_body, _err := client.ListJMeterReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOpenJMeterScenesWithOptions(request *ListOpenJMeterScenesRequest, runtime *util.RuntimeOptions) (_result *ListOpenJMeterScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOpenJMeterScenesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOpenJMeterScenes"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOpenJMeterScenes(request *ListOpenJMeterScenesRequest) (_result *ListOpenJMeterScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpenJMeterScenesResponse{}
	_body, _err := client.ListOpenJMeterScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPtsSceneWithOptions(request *ListPtsSceneRequest, runtime *util.RuntimeOptions) (_result *ListPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPtsScene(request *ListPtsSceneRequest) (_result *ListPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPtsSceneResponse{}
	_body, _err := client.ListPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPtsSceneWithOptions(request *ModifyPtsSceneRequest, runtime *util.RuntimeOptions) (_result *ModifyPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPtsScene(request *ModifyPtsSceneRequest) (_result *ModifyPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPtsSceneResponse{}
	_body, _err := client.ModifyPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveEnvWithOptions(request *RemoveEnvRequest, runtime *util.RuntimeOptions) (_result *RemoveEnvResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveEnvResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveEnv"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveEnv(request *RemoveEnvRequest) (_result *RemoveEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveEnvResponse{}
	_body, _err := client.RemoveEnvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveOpenJMeterSceneWithOptions(request *RemoveOpenJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *RemoveOpenJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveOpenJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveOpenJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveOpenJMeterScene(request *RemoveOpenJMeterSceneRequest) (_result *RemoveOpenJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveOpenJMeterSceneResponse{}
	_body, _err := client.RemoveOpenJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveEnvWithOptions(tmpReq *SaveEnvRequest, runtime *util.RuntimeOptions) (_result *SaveEnvResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveEnvShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.Env))) {
		request.EnvShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.Env), tea.String("Env"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveEnvResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveEnv"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveEnv(request *SaveEnvRequest) (_result *SaveEnvResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveEnvResponse{}
	_body, _err := client.SaveEnvWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveOpenJMeterSceneWithOptions(tmpReq *SaveOpenJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *SaveOpenJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveOpenJMeterSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.OpenJMeterScene))) {
		request.OpenJMeterSceneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.OpenJMeterScene), tea.String("OpenJMeterScene"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveOpenJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveOpenJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveOpenJMeterScene(request *SaveOpenJMeterSceneRequest) (_result *SaveOpenJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveOpenJMeterSceneResponse{}
	_body, _err := client.SaveOpenJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDebugPtsSceneWithOptions(request *StartDebugPtsSceneRequest, runtime *util.RuntimeOptions) (_result *StartDebugPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartDebugPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartDebugPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDebugPtsScene(request *StartDebugPtsSceneRequest) (_result *StartDebugPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDebugPtsSceneResponse{}
	_body, _err := client.StartDebugPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDebuggingJMeterSceneWithOptions(request *StartDebuggingJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *StartDebuggingJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartDebuggingJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartDebuggingJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDebuggingJMeterScene(request *StartDebuggingJMeterSceneRequest) (_result *StartDebuggingJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDebuggingJMeterSceneResponse{}
	_body, _err := client.StartDebuggingJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartPtsSceneWithOptions(request *StartPtsSceneRequest, runtime *util.RuntimeOptions) (_result *StartPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartPtsScene(request *StartPtsSceneRequest) (_result *StartPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartPtsSceneResponse{}
	_body, _err := client.StartPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTestingJMeterSceneWithOptions(request *StartTestingJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *StartTestingJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartTestingJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartTestingJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTestingJMeterScene(request *StartTestingJMeterSceneRequest) (_result *StartTestingJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTestingJMeterSceneResponse{}
	_body, _err := client.StartTestingJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDebugPtsSceneWithOptions(request *StopDebugPtsSceneRequest, runtime *util.RuntimeOptions) (_result *StopDebugPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopDebugPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopDebugPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDebugPtsScene(request *StopDebugPtsSceneRequest) (_result *StopDebugPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDebugPtsSceneResponse{}
	_body, _err := client.StopDebugPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopDebuggingJMeterSceneWithOptions(request *StopDebuggingJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *StopDebuggingJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopDebuggingJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopDebuggingJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDebuggingJMeterScene(request *StopDebuggingJMeterSceneRequest) (_result *StopDebuggingJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDebuggingJMeterSceneResponse{}
	_body, _err := client.StopDebuggingJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopPtsSceneWithOptions(request *StopPtsSceneRequest, runtime *util.RuntimeOptions) (_result *StopPtsSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopPtsSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopPtsScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopPtsScene(request *StopPtsSceneRequest) (_result *StopPtsSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopPtsSceneResponse{}
	_body, _err := client.StopPtsSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTestingJMeterSceneWithOptions(request *StopTestingJMeterSceneRequest, runtime *util.RuntimeOptions) (_result *StopTestingJMeterSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopTestingJMeterSceneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopTestingJMeterScene"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTestingJMeterScene(request *StopTestingJMeterSceneRequest) (_result *StopTestingJMeterSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTestingJMeterSceneResponse{}
	_body, _err := client.StopTestingJMeterSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePtsSceneBaseLineWithOptions(tmpReq *UpdatePtsSceneBaseLineRequest, runtime *util.RuntimeOptions) (_result *UpdatePtsSceneBaseLineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePtsSceneBaseLineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApiBaselines)) {
		request.ApiBaselinesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiBaselines, tea.String("ApiBaselines"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SceneBaseline)) {
		request.SceneBaselineShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SceneBaseline, tea.String("SceneBaseline"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdatePtsSceneBaseLineResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdatePtsSceneBaseLine"), tea.String("2020-10-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePtsSceneBaseLine(request *UpdatePtsSceneBaseLineRequest) (_result *UpdatePtsSceneBaseLineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePtsSceneBaseLineResponse{}
	_body, _err := client.UpdatePtsSceneBaseLineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
