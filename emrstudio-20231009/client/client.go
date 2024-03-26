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

type ListProjectsRequest struct {
	MaxResults  *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken   *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchVal   *string `json:"searchVal,omitempty" xml:"searchVal,omitempty"`
	WorkspaceId *int64  `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetMaxResults(v int32) *ListProjectsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProjectsRequest) SetNextToken(v string) *ListProjectsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProjectsRequest) SetSearchVal(v string) *ListProjectsRequest {
	s.SearchVal = &v
	return s
}

func (s *ListProjectsRequest) SetWorkspaceId(v int64) *ListProjectsRequest {
	s.WorkspaceId = &v
	return s
}

type ListProjectsResponseBody struct {
	Data      []*ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	NextToken *string                         `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetData(v []*ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetNextToken(v string) *ListProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyData struct {
	Code        *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProjectId   *int64  `json:"projectId,omitempty" xml:"projectId,omitempty"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetCode(v int64) *ListProjectsResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetDescription(v string) *ListProjectsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetName(v string) *ListProjectsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjectId(v int64) *ListProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("emrstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SearchVal)) {
		query["searchVal"] = request.SearchVal
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2023-10-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/dolphinscheduler/v3/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
