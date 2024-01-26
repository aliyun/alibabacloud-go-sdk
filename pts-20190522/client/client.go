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

type GetAliwareReportRequest struct {
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetAliwareReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAliwareReportRequest) GoString() string {
	return s.String()
}

func (s *GetAliwareReportRequest) SetReportId(v int64) *GetAliwareReportRequest {
	s.ReportId = &v
	return s
}

type GetAliwareReportResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshot  *string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Summary   *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetAliwareReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAliwareReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAliwareReportResponseBody) SetCode(v string) *GetAliwareReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetAliwareReportResponseBody) SetMessage(v string) *GetAliwareReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetAliwareReportResponseBody) SetRequestId(v string) *GetAliwareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAliwareReportResponseBody) SetSnapshot(v string) *GetAliwareReportResponseBody {
	s.Snapshot = &v
	return s
}

func (s *GetAliwareReportResponseBody) SetSuccess(v bool) *GetAliwareReportResponseBody {
	s.Success = &v
	return s
}

func (s *GetAliwareReportResponseBody) SetSummary(v string) *GetAliwareReportResponseBody {
	s.Summary = &v
	return s
}

type GetAliwareReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAliwareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliwareReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAliwareReportResponse) GoString() string {
	return s.String()
}

func (s *GetAliwareReportResponse) SetHeaders(v map[string]*string) *GetAliwareReportResponse {
	s.Headers = v
	return s
}

func (s *GetAliwareReportResponse) SetStatusCode(v int32) *GetAliwareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliwareReportResponse) SetBody(v *GetAliwareReportResponseBody) *GetAliwareReportResponse {
	s.Body = v
	return s
}

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

type StartSceneRequest struct {
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	TaskId  *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TeamId  *int64 `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	UserId  *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *StartSceneRequest) SetTaskId(v int64) *StartSceneRequest {
	s.TaskId = &v
	return s
}

func (s *StartSceneRequest) SetTeamId(v int64) *StartSceneRequest {
	s.TeamId = &v
	return s
}

func (s *StartSceneRequest) SetUserId(v int64) *StartSceneRequest {
	s.UserId = &v
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
 * @param request GetAliwareReportRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAliwareReportResponse
 */
// Deprecated
func (client *Client) GetAliwareReportWithOptions(request *GetAliwareReportRequest, runtime *util.RuntimeOptions) (_result *GetAliwareReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAliwareReport"),
		Version:     tea.String("2019-05-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAliwareReportResponse{}
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
 * @param request GetAliwareReportRequest
 * @return GetAliwareReportResponse
 */
// Deprecated
func (client *Client) GetAliwareReport(request *GetAliwareReportRequest) (_result *GetAliwareReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAliwareReportResponse{}
	_body, _err := client.GetAliwareReportWithOptions(request, runtime)
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
		Version:     tea.String("2019-05-22"),
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

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TeamId)) {
		query["TeamId"] = request.TeamId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartScene"),
		Version:     tea.String("2019-05-22"),
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
