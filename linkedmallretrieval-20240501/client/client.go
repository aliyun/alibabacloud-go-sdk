// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AISearchQuery struct {
	// example:
	//
	// active
	Card *string `json:"card,omitempty" xml:"card,omitempty"`
	// example:
	//
	// 今年五一假期放假时间表
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AISearchQuery) String() string {
	return tea.Prettify(s)
}

func (s AISearchQuery) GoString() string {
	return s.String()
}

func (s *AISearchQuery) SetCard(v string) *AISearchQuery {
	s.Card = &v
	return s
}

func (s *AISearchQuery) SetQuery(v string) *AISearchQuery {
	s.Query = &v
	return s
}

type AISearchResult struct {
	Header  *EventHeader `json:"header,omitempty" xml:"header,omitempty"`
	Payload *string      `json:"payload,omitempty" xml:"payload,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResult) String() string {
	return tea.Prettify(s)
}

func (s AISearchResult) GoString() string {
	return s.String()
}

func (s *AISearchResult) SetHeader(v *EventHeader) *AISearchResult {
	s.Header = v
	return s
}

func (s *AISearchResult) SetPayload(v string) *AISearchResult {
	s.Payload = &v
	return s
}

func (s *AISearchResult) SetRequestId(v string) *AISearchResult {
	s.RequestId = &v
	return s
}

type EventHeader struct {
	// example:
	//
	// on_common_search_stream
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// ff3de49-cedc-47ea-ba3c-8456acd345d8
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// 55c2349-cedc-47ea-ba3c-8456acd6c7d8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1403
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s EventHeader) String() string {
	return tea.Prettify(s)
}

func (s EventHeader) GoString() string {
	return s.String()
}

func (s *EventHeader) SetEvent(v string) *EventHeader {
	s.Event = &v
	return s
}

func (s *EventHeader) SetEventId(v string) *EventHeader {
	s.EventId = &v
	return s
}

func (s *EventHeader) SetRequestId(v string) *EventHeader {
	s.RequestId = &v
	return s
}

func (s *EventHeader) SetResponseTime(v int64) *EventHeader {
	s.ResponseTime = &v
	return s
}

type AISearchRequest struct {
	// example:
	//
	// active
	Card *string `json:"card,omitempty" xml:"card,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 今年五一假期放假时间表
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s AISearchRequest) String() string {
	return tea.Prettify(s)
}

func (s AISearchRequest) GoString() string {
	return s.String()
}

func (s *AISearchRequest) SetCard(v string) *AISearchRequest {
	s.Card = &v
	return s
}

func (s *AISearchRequest) SetQuery(v string) *AISearchRequest {
	s.Query = &v
	return s
}

type AISearchResponseBody struct {
	Header  *AISearchResponseBodyHeader `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *string                     `json:"payload,omitempty" xml:"payload,omitempty"`
	// example:
	//
	// D016A23D-738A-5209-A91A-6145845C5A23
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AISearchResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchResponseBody) SetHeader(v *AISearchResponseBodyHeader) *AISearchResponseBody {
	s.Header = v
	return s
}

func (s *AISearchResponseBody) SetPayload(v string) *AISearchResponseBody {
	s.Payload = &v
	return s
}

func (s *AISearchResponseBody) SetRequestId(v string) *AISearchResponseBody {
	s.RequestId = &v
	return s
}

type AISearchResponseBodyHeader struct {
	// example:
	//
	// on_common_search_stream
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// ff3de49-cedc-47ea-ba3c-8456acd345d8
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// D4F6D088-EDE9-5823-9E66-22603937A40B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1403
	ResponseTime *string `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s AISearchResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s AISearchResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *AISearchResponseBodyHeader) SetEvent(v string) *AISearchResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *AISearchResponseBodyHeader) SetEventId(v string) *AISearchResponseBodyHeader {
	s.EventId = &v
	return s
}

func (s *AISearchResponseBodyHeader) SetRequestId(v string) *AISearchResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *AISearchResponseBodyHeader) SetResponseTime(v string) *AISearchResponseBodyHeader {
	s.ResponseTime = &v
	return s
}

type AISearchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AISearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchResponse) String() string {
	return tea.Prettify(s)
}

func (s AISearchResponse) GoString() string {
	return s.String()
}

func (s *AISearchResponse) SetHeaders(v map[string]*string) *AISearchResponse {
	s.Headers = v
	return s
}

func (s *AISearchResponse) SetStatusCode(v int32) *AISearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchResponse) SetBody(v *AISearchResponseBody) *AISearchResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkedmallretrieval"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// AI搜索
//
// @param request - AISearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AISearchResponse
func (client *Client) AISearchWithOptions(request *AISearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AISearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Card)) {
		query["card"] = request.Card
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AISearch"),
		Version:     tea.String("2024-05-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-entry/v1/linkedRetrieval/commands/aiSearch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AISearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// AI搜索
//
// @param request - AISearchRequest
//
// @return AISearchResponse
func (client *Client) AISearch(request *AISearchRequest) (_result *AISearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AISearchResponse{}
	_body, _err := client.AISearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
