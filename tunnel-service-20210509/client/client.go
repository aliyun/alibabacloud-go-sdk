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

type DeleteSessionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSessionResponse) SetHeaders(v map[string]*string) *DeleteSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSessionResponse) SetBody(v string) *DeleteSessionResponse {
	s.Body = &v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetBody(v string) *GetInstanceResponse {
	s.Body = &v
	return s
}

type HeartBeatRequest struct {
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	InstanceType  *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	SessionStatus *string `json:"sessionStatus,omitempty" xml:"sessionStatus,omitempty"`
}

func (s HeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s HeartBeatRequest) GoString() string {
	return s.String()
}

func (s *HeartBeatRequest) SetInstanceId(v string) *HeartBeatRequest {
	s.InstanceId = &v
	return s
}

func (s *HeartBeatRequest) SetInstanceType(v string) *HeartBeatRequest {
	s.InstanceType = &v
	return s
}

func (s *HeartBeatRequest) SetSessionStatus(v string) *HeartBeatRequest {
	s.SessionStatus = &v
	return s
}

type HeartBeatResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s HeartBeatResponse) GoString() string {
	return s.String()
}

func (s *HeartBeatResponse) SetHeaders(v map[string]*string) *HeartBeatResponse {
	s.Headers = v
	return s
}

func (s *HeartBeatResponse) SetBody(v string) *HeartBeatResponse {
	s.Body = &v
	return s
}

type UnRegisterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnRegisterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnRegisterInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnRegisterInstanceResponse) SetHeaders(v map[string]*string) *UnRegisterInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnRegisterInstanceResponse) SetBody(v string) *UnRegisterInstanceResponse {
	s.Body = &v
	return s
}

type CreateSessionRequest struct {
	SessionName *string `json:"sessionName,omitempty" xml:"sessionName,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) SetSessionName(v string) *CreateSessionRequest {
	s.SessionName = &v
	return s
}

type CreateSessionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) SetHeaders(v map[string]*string) *CreateSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionResponse) SetBody(v string) *CreateSessionResponse {
	s.Body = &v
	return s
}

type RegisterInstanceRequest struct {
	InstanceId   *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	Source       *string `json:"source,omitempty" xml:"source,omitempty"`
	Description  *string `json:"description,omitempty" xml:"description,omitempty"`
	Params       *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s RegisterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterInstanceRequest) GoString() string {
	return s.String()
}

func (s *RegisterInstanceRequest) SetInstanceId(v string) *RegisterInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RegisterInstanceRequest) SetSessionId(v string) *RegisterInstanceRequest {
	s.SessionId = &v
	return s
}

func (s *RegisterInstanceRequest) SetInstanceType(v string) *RegisterInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *RegisterInstanceRequest) SetSource(v string) *RegisterInstanceRequest {
	s.Source = &v
	return s
}

func (s *RegisterInstanceRequest) SetDescription(v string) *RegisterInstanceRequest {
	s.Description = &v
	return s
}

func (s *RegisterInstanceRequest) SetParams(v string) *RegisterInstanceRequest {
	s.Params = &v
	return s
}

type RegisterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterInstanceResponse) GoString() string {
	return s.String()
}

func (s *RegisterInstanceResponse) SetHeaders(v map[string]*string) *RegisterInstanceResponse {
	s.Headers = v
	return s
}

func (s *RegisterInstanceResponse) SetBody(v string) *RegisterInstanceResponse {
	s.Body = &v
	return s
}

type GetSessionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSessionResponse) GoString() string {
	return s.String()
}

func (s *GetSessionResponse) SetHeaders(v map[string]*string) *GetSessionResponse {
	s.Headers = v
	return s
}

func (s *GetSessionResponse) SetBody(v string) *GetSessionResponse {
	s.Body = &v
	return s
}

type ListSessionsRequest struct {
	PageNum  *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionsRequest) SetPageNum(v int32) *ListSessionsRequest {
	s.PageNum = &v
	return s
}

func (s *ListSessionsRequest) SetPageSize(v int32) *ListSessionsRequest {
	s.PageSize = &v
	return s
}

type ListSessionsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListSessionsResponse) SetHeaders(v map[string]*string) *ListSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListSessionsResponse) SetBody(v string) *ListSessionsResponse {
	s.Body = &v
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("tunnel-service"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DeleteSession(sessionId *string) (_result *DeleteSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSessionResponse{}
	_body, _err := client.DeleteSessionWithOptions(sessionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSessionWithOptions(sessionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSessionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteSessionResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSession"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/v1/sessions/"+tea.StringValue(sessionId)), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstance"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/instances/"+tea.StringValue(instanceId)), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HeartBeat(request *HeartBeatRequest) (_result *HeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &HeartBeatResponse{}
	_body, _err := client.HeartBeatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HeartBeatWithOptions(request *HeartBeatRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *HeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionStatus)) {
		body["sessionStatus"] = request.SessionStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &HeartBeatResponse{}
	_body, _err := client.DoROARequest(tea.String("HeartBeat"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/sessions/"), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnRegisterInstance(instanceId *string) (_result *UnRegisterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnRegisterInstanceResponse{}
	_body, _err := client.UnRegisterInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnRegisterInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnRegisterInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &UnRegisterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("UnRegisterInstance"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/v1/instances/"+tea.StringValue(instanceId)), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionName)) {
		body["sessionName"] = request.SessionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateSession"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/sessions/"), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterInstance(request *RegisterInstanceRequest) (_result *RegisterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RegisterInstanceResponse{}
	_body, _err := client.RegisterInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterInstanceWithOptions(request *RegisterInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RegisterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		body["source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &RegisterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("RegisterInstance"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/v1/instances/"), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSession(sessionId *string) (_result *GetSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSessionResponse{}
	_body, _err := client.GetSessionWithOptions(sessionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSessionWithOptions(sessionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSessionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetSessionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSession"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/sessions/"+tea.StringValue(sessionId)), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSessions(request *ListSessionsRequest) (_result *ListSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSessionsResponse{}
	_body, _err := client.ListSessionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSessionsWithOptions(request *ListSessionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["pageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSessionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSessions"), tea.String("2021-05-09"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/v1/sessions/"), tea.String("string"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
