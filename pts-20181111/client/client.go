// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetReportRequest struct {
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReportRequest) GoString() string {
	return s.String()
}

func (s *GetReportRequest) SetReportId(v int64) *GetReportRequest {
	s.ReportId = &v
	return s
}

type GetReportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshot  *string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetReportResponseBody) SetCode(v string) *GetReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetReportResponseBody) SetMessage(v string) *GetReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetReportResponseBody) SetRequestId(v string) *GetReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReportResponseBody) SetSnapshot(v string) *GetReportResponseBody {
	s.Snapshot = &v
	return s
}

func (s *GetReportResponseBody) SetSuccess(v bool) *GetReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetReportResponseBody) SetSummary(v string) *GetReportResponseBody {
	s.Summary = &v
	return s
}

type GetReportResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReportResponse) GoString() string {
	return s.String()
}

func (s *GetReportResponse) SetHeaders(v map[string]*string) *GetReportResponse {
	s.Headers = v
	return s
}

func (s *GetReportResponse) SetStatusCode(v int32) *GetReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReportResponse) SetBody(v *GetReportResponseBody) *GetReportResponse {
	s.Body = v
	return s
}

type ListRunnableScenesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRunnableScenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunnableScenesRequest) GoString() string {
	return s.String()
}

func (s *ListRunnableScenesRequest) SetPageNumber(v int32) *ListRunnableScenesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRunnableScenesRequest) SetPageSize(v int32) *ListRunnableScenesRequest {
	s.PageSize = &v
	return s
}

type ListRunnableScenesResponseBody struct {
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenes     *ListRunnableScenesResponseBodyScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Struct"`
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount *int64                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRunnableScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunnableScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunnableScenesResponseBody) SetCode(v string) *ListRunnableScenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetMessage(v string) *ListRunnableScenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetPageNumber(v int32) *ListRunnableScenesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetPageSize(v int32) *ListRunnableScenesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetRequestId(v string) *ListRunnableScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetScenes(v *ListRunnableScenesResponseBodyScenes) *ListRunnableScenesResponseBody {
	s.Scenes = v
	return s
}

func (s *ListRunnableScenesResponseBody) SetSuccess(v bool) *ListRunnableScenesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRunnableScenesResponseBody) SetTotalCount(v int64) *ListRunnableScenesResponseBody {
	s.TotalCount = &v
	return s
}

type ListRunnableScenesResponseBodyScenes struct {
	Scene []*ListRunnableScenesResponseBodyScenesScene `json:"Scene,omitempty" xml:"Scene,omitempty" type:"Repeated"`
}

func (s ListRunnableScenesResponseBodyScenes) String() string {
	return tea.Prettify(s)
}

func (s ListRunnableScenesResponseBodyScenes) GoString() string {
	return s.String()
}

func (s *ListRunnableScenesResponseBodyScenes) SetScene(v []*ListRunnableScenesResponseBodyScenesScene) *ListRunnableScenesResponseBodyScenes {
	s.Scene = v
	return s
}

type ListRunnableScenesResponseBodyScenesScene struct {
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SceneId      *int64  `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	SceneName    *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ListRunnableScenesResponseBodyScenesScene) String() string {
	return tea.Prettify(s)
}

func (s ListRunnableScenesResponseBodyScenesScene) GoString() string {
	return s.String()
}

func (s *ListRunnableScenesResponseBodyScenesScene) SetDuration(v int32) *ListRunnableScenesResponseBodyScenesScene {
	s.Duration = &v
	return s
}

func (s *ListRunnableScenesResponseBodyScenesScene) SetModifiedTime(v int64) *ListRunnableScenesResponseBodyScenesScene {
	s.ModifiedTime = &v
	return s
}

func (s *ListRunnableScenesResponseBodyScenesScene) SetSceneId(v int64) *ListRunnableScenesResponseBodyScenesScene {
	s.SceneId = &v
	return s
}

func (s *ListRunnableScenesResponseBodyScenesScene) SetSceneName(v string) *ListRunnableScenesResponseBodyScenesScene {
	s.SceneName = &v
	return s
}

type ListRunnableScenesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRunnableScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRunnableScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunnableScenesResponse) GoString() string {
	return s.String()
}

func (s *ListRunnableScenesResponse) SetHeaders(v map[string]*string) *ListRunnableScenesResponse {
	s.Headers = v
	return s
}

func (s *ListRunnableScenesResponse) SetStatusCode(v int32) *ListRunnableScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunnableScenesResponse) SetBody(v *ListRunnableScenesResponseBody) *ListRunnableScenesResponse {
	s.Body = v
	return s
}

type QueryPlanStatusRequest struct {
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	SceneId  *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s QueryPlanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPlanStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryPlanStatusRequest) SetReportId(v int64) *QueryPlanStatusRequest {
	s.ReportId = &v
	return s
}

func (s *QueryPlanStatusRequest) SetSceneId(v int64) *QueryPlanStatusRequest {
	s.SceneId = &v
	return s
}

type QueryPlanStatusResponseBody struct {
	AgentLocations      *QueryPlanStatusResponseBodyAgentLocations `json:"AgentLocations,omitempty" xml:"AgentLocations,omitempty" type:"Struct"`
	AliveAgentCount     *int32                                     `json:"AliveAgentCount,omitempty" xml:"AliveAgentCount,omitempty"`
	AverageRt           *int32                                     `json:"AverageRt,omitempty" xml:"AverageRt,omitempty"`
	BpsRequest          *string                                    `json:"BpsRequest,omitempty" xml:"BpsRequest,omitempty"`
	BpsResponse         *string                                    `json:"BpsResponse,omitempty" xml:"BpsResponse,omitempty"`
	Code                *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Concurrency         *int32                                     `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	ConcurrencyLimit    *int32                                     `json:"ConcurrencyLimit,omitempty" xml:"ConcurrencyLimit,omitempty"`
	CurrentTime         *int64                                     `json:"CurrentTime,omitempty" xml:"CurrentTime,omitempty"`
	FailedBusinessCount *int32                                     `json:"FailedBusinessCount,omitempty" xml:"FailedBusinessCount,omitempty"`
	FailedRequestCount  *int32                                     `json:"FailedRequestCount,omitempty" xml:"FailedRequestCount,omitempty"`
	Message             *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	MonitorData         *QueryPlanStatusResponseBodyMonitorData    `json:"MonitorData,omitempty" xml:"MonitorData,omitempty" type:"Struct"`
	ReportId            *int64                                     `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestCount        *string                                    `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	RequestId           *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Seg90Rt             *int32                                     `json:"Seg90Rt,omitempty" xml:"Seg90Rt,omitempty"`
	StartTime           *int64                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Success             *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Tips                *string                                    `json:"Tips,omitempty" xml:"Tips,omitempty"`
	TotalAgentCount     *int32                                     `json:"TotalAgentCount,omitempty" xml:"TotalAgentCount,omitempty"`
	Tps                 *int32                                     `json:"Tps,omitempty" xml:"Tps,omitempty"`
	TpsLimit            *int32                                     `json:"TpsLimit,omitempty" xml:"TpsLimit,omitempty"`
	Vum                 *int32                                     `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s QueryPlanStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryPlanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPlanStatusResponseBody) SetAgentLocations(v *QueryPlanStatusResponseBodyAgentLocations) *QueryPlanStatusResponseBody {
	s.AgentLocations = v
	return s
}

func (s *QueryPlanStatusResponseBody) SetAliveAgentCount(v int32) *QueryPlanStatusResponseBody {
	s.AliveAgentCount = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetAverageRt(v int32) *QueryPlanStatusResponseBody {
	s.AverageRt = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetBpsRequest(v string) *QueryPlanStatusResponseBody {
	s.BpsRequest = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetBpsResponse(v string) *QueryPlanStatusResponseBody {
	s.BpsResponse = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetCode(v string) *QueryPlanStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetConcurrency(v int32) *QueryPlanStatusResponseBody {
	s.Concurrency = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetConcurrencyLimit(v int32) *QueryPlanStatusResponseBody {
	s.ConcurrencyLimit = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetCurrentTime(v int64) *QueryPlanStatusResponseBody {
	s.CurrentTime = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetFailedBusinessCount(v int32) *QueryPlanStatusResponseBody {
	s.FailedBusinessCount = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetFailedRequestCount(v int32) *QueryPlanStatusResponseBody {
	s.FailedRequestCount = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetMessage(v string) *QueryPlanStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetMonitorData(v *QueryPlanStatusResponseBodyMonitorData) *QueryPlanStatusResponseBody {
	s.MonitorData = v
	return s
}

func (s *QueryPlanStatusResponseBody) SetReportId(v int64) *QueryPlanStatusResponseBody {
	s.ReportId = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetRequestCount(v string) *QueryPlanStatusResponseBody {
	s.RequestCount = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetRequestId(v string) *QueryPlanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetSeg90Rt(v int32) *QueryPlanStatusResponseBody {
	s.Seg90Rt = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetStartTime(v int64) *QueryPlanStatusResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetSuccess(v bool) *QueryPlanStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetTips(v string) *QueryPlanStatusResponseBody {
	s.Tips = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetTotalAgentCount(v int32) *QueryPlanStatusResponseBody {
	s.TotalAgentCount = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetTps(v int32) *QueryPlanStatusResponseBody {
	s.Tps = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetTpsLimit(v int32) *QueryPlanStatusResponseBody {
	s.TpsLimit = &v
	return s
}

func (s *QueryPlanStatusResponseBody) SetVum(v int32) *QueryPlanStatusResponseBody {
	s.Vum = &v
	return s
}

type QueryPlanStatusResponseBodyAgentLocations struct {
	AgentLocation []map[string]interface{} `json:"AgentLocation,omitempty" xml:"AgentLocation,omitempty" type:"Repeated"`
}

func (s QueryPlanStatusResponseBodyAgentLocations) String() string {
	return tea.Prettify(s)
}

func (s QueryPlanStatusResponseBodyAgentLocations) GoString() string {
	return s.String()
}

func (s *QueryPlanStatusResponseBodyAgentLocations) SetAgentLocation(v []map[string]interface{}) *QueryPlanStatusResponseBodyAgentLocations {
	s.AgentLocation = v
	return s
}

type QueryPlanStatusResponseBodyMonitorData struct {
	Data []map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s QueryPlanStatusResponseBodyMonitorData) String() string {
	return tea.Prettify(s)
}

func (s QueryPlanStatusResponseBodyMonitorData) GoString() string {
	return s.String()
}

func (s *QueryPlanStatusResponseBodyMonitorData) SetData(v []map[string]interface{}) *QueryPlanStatusResponseBodyMonitorData {
	s.Data = v
	return s
}

type QueryPlanStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPlanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPlanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPlanStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryPlanStatusResponse) SetHeaders(v map[string]*string) *QueryPlanStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryPlanStatusResponse) SetStatusCode(v int32) *QueryPlanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPlanStatusResponse) SetBody(v *QueryPlanStatusResponseBody) *QueryPlanStatusResponse {
	s.Body = v
	return s
}

type StartSceneRequest struct {
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSceneRequest) GoString() string {
	return s.String()
}

func (s *StartSceneRequest) SetSceneId(v int64) *StartSceneRequest {
	s.SceneId = &v
	return s
}

type StartSceneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReportId  *int64  `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartSceneResponseBody) SetCode(v string) *StartSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartSceneResponseBody) SetMessage(v string) *StartSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartSceneResponseBody) SetReportId(v int64) *StartSceneResponseBody {
	s.ReportId = &v
	return s
}

func (s *StartSceneResponseBody) SetRequestId(v string) *StartSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSceneResponseBody) SetSuccess(v bool) *StartSceneResponseBody {
	s.Success = &v
	return s
}

type StartSceneResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSceneResponse) GoString() string {
	return s.String()
}

func (s *StartSceneResponse) SetHeaders(v map[string]*string) *StartSceneResponse {
	s.Headers = v
	return s
}

func (s *StartSceneResponse) SetStatusCode(v int32) *StartSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSceneResponse) SetBody(v *StartSceneResponseBody) *StartSceneResponse {
	s.Body = v
	return s
}

type StopSceneRequest struct {
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s StopSceneRequest) GoString() string {
	return s.String()
}

func (s *StopSceneRequest) SetSceneId(v int64) *StopSceneRequest {
	s.SceneId = &v
	return s
}

type StopSceneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StopSceneResponseBody) SetCode(v string) *StopSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StopSceneResponseBody) SetMessage(v string) *StopSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StopSceneResponseBody) SetRequestId(v string) *StopSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSceneResponseBody) SetSuccess(v bool) *StopSceneResponseBody {
	s.Success = &v
	return s
}

type StopSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s StopSceneResponse) GoString() string {
	return s.String()
}

func (s *StopSceneResponse) SetHeaders(v map[string]*string) *StopSceneResponse {
	s.Headers = v
	return s
}

func (s *StopSceneResponse) SetStatusCode(v int32) *StopSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSceneResponse) SetBody(v *StopSceneResponseBody) *StopSceneResponse {
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

/**
 * @deprecated
 *
 * @param request GetReportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetReportResponse
 */
// Deprecated
func (client *Client) GetReportWithOptions(request *GetReportRequest, runtime *util.RuntimeOptions) (_result *GetReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReport"),
		Version:     tea.String("2018-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetReportRequest
 * @return GetReportResponse
 */
// Deprecated
func (client *Client) GetReport(request *GetReportRequest) (_result *GetReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReportResponse{}
	_body, _err := client.GetReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ListRunnableScenesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRunnableScenesResponse
 */
// Deprecated
func (client *Client) ListRunnableScenesWithOptions(request *ListRunnableScenesRequest, runtime *util.RuntimeOptions) (_result *ListRunnableScenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRunnableScenes"),
		Version:     tea.String("2018-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunnableScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ListRunnableScenesRequest
 * @return ListRunnableScenesResponse
 */
// Deprecated
func (client *Client) ListRunnableScenes(request *ListRunnableScenesRequest) (_result *ListRunnableScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRunnableScenesResponse{}
	_body, _err := client.ListRunnableScenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request QueryPlanStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryPlanStatusResponse
 */
// Deprecated
func (client *Client) QueryPlanStatusWithOptions(request *QueryPlanStatusRequest, runtime *util.RuntimeOptions) (_result *QueryPlanStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryPlanStatus"),
		Version:     tea.String("2018-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryPlanStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request QueryPlanStatusRequest
 * @return QueryPlanStatusResponse
 */
// Deprecated
func (client *Client) QueryPlanStatus(request *QueryPlanStatusRequest) (_result *QueryPlanStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryPlanStatusResponse{}
	_body, _err := client.QueryPlanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StartSceneRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartSceneResponse
 */
// Deprecated
func (client *Client) StartSceneWithOptions(request *StartSceneRequest, runtime *util.RuntimeOptions) (_result *StartSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartScene"),
		Version:     tea.String("2018-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StartSceneRequest
 * @return StartSceneResponse
 */
// Deprecated
func (client *Client) StartScene(request *StartSceneRequest) (_result *StartSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSceneResponse{}
	_body, _err := client.StartSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StopSceneRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopSceneResponse
 */
// Deprecated
func (client *Client) StopSceneWithOptions(request *StopSceneRequest, runtime *util.RuntimeOptions) (_result *StopSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SceneId)) {
		query["SceneId"] = request.SceneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopScene"),
		Version:     tea.String("2018-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StopSceneRequest
 * @return StopSceneResponse
 */
// Deprecated
func (client *Client) StopScene(request *StopSceneRequest) (_result *StopSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopSceneResponse{}
	_body, _err := client.StopSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
