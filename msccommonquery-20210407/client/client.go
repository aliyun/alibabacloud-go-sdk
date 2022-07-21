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

type ListMessagesRequest struct {
	ChannelType    *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	EndTimestamp   *int64  `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Locale         *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTimestamp *int64  `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) SetChannelType(v string) *ListMessagesRequest {
	s.ChannelType = &v
	return s
}

func (s *ListMessagesRequest) SetEndTimestamp(v int64) *ListMessagesRequest {
	s.EndTimestamp = &v
	return s
}

func (s *ListMessagesRequest) SetLocale(v string) *ListMessagesRequest {
	s.Locale = &v
	return s
}

func (s *ListMessagesRequest) SetPageNo(v int32) *ListMessagesRequest {
	s.PageNo = &v
	return s
}

func (s *ListMessagesRequest) SetPageSize(v int32) *ListMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessagesRequest) SetStartTimestamp(v int64) *ListMessagesRequest {
	s.StartTimestamp = &v
	return s
}

type ListMessagesResponseBody struct {
	Code    *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*ListMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int32  `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) SetCode(v string) *ListMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMessagesResponseBody) SetData(v []*ListMessagesResponseBodyData) *ListMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBody) SetMessage(v string) *ListMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessagesResponseBody) SetSuccess(v bool) *ListMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMessagesResponseBody) SetTotal(v int32) *ListMessagesResponseBody {
	s.Total = &v
	return s
}

type ListMessagesResponseBodyData struct {
	ChannelType *string `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	MessageId   *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Receiver    *string `json:"Receiver,omitempty" xml:"Receiver,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Timestamp   *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListMessagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyData) SetChannelType(v string) *ListMessagesResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetContent(v string) *ListMessagesResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetMessageId(v string) *ListMessagesResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetReceiver(v string) *ListMessagesResponseBodyData {
	s.Receiver = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetStatus(v string) *ListMessagesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTimestamp(v int64) *ListMessagesResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTitle(v string) *ListMessagesResponseBodyData {
	s.Title = &v
	return s
}

type ListMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMessagesResponse) SetHeaders(v map[string]*string) *ListMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMessagesResponse) SetStatusCode(v int32) *ListMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessagesResponse) SetBody(v *ListMessagesResponseBody) *ListMessagesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("msccommonquery"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListMessagesWithOptions(request *ListMessagesRequest, runtime *util.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Locale)) {
		query["Locale"] = request.Locale
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimestamp)) {
		body["EndTimestamp"] = request.EndTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimestamp)) {
		body["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessages"),
		Version:     tea.String("2021-04-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMessages(request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
