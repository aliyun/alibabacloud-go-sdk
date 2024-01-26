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

type ListEventInProgressRequest struct {
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
}

func (s ListEventInProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressRequest) GoString() string {
	return s.String()
}

func (s *ListEventInProgressRequest) SetRegionIds(v []*string) *ListEventInProgressRequest {
	s.RegionIds = v
	return s
}

type ListEventInProgressShrinkRequest struct {
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
}

func (s ListEventInProgressShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEventInProgressShrinkRequest) SetRegionIdsShrink(v string) *ListEventInProgressShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

type ListEventInProgressResponseBody struct {
	Data      []*ListEventInProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEventInProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBody) SetData(v []*ListEventInProgressResponseBodyData) *ListEventInProgressResponseBody {
	s.Data = v
	return s
}

func (s *ListEventInProgressResponseBody) SetRequestId(v string) *ListEventInProgressResponseBody {
	s.RequestId = &v
	return s
}

type ListEventInProgressResponseBodyData struct {
	EventUpdates []*ListEventInProgressResponseBodyDataEventUpdates `json:"EventUpdates,omitempty" xml:"EventUpdates,omitempty" type:"Repeated"`
	Id           *int64                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Impacts      []*ListEventInProgressResponseBodyDataImpacts      `json:"Impacts,omitempty" xml:"Impacts,omitempty" type:"Repeated"`
	StartTime    *int64                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Title        *string                                            `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListEventInProgressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBodyData) SetEventUpdates(v []*ListEventInProgressResponseBodyDataEventUpdates) *ListEventInProgressResponseBodyData {
	s.EventUpdates = v
	return s
}

func (s *ListEventInProgressResponseBodyData) SetId(v int64) *ListEventInProgressResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListEventInProgressResponseBodyData) SetImpacts(v []*ListEventInProgressResponseBodyDataImpacts) *ListEventInProgressResponseBodyData {
	s.Impacts = v
	return s
}

func (s *ListEventInProgressResponseBodyData) SetStartTime(v int64) *ListEventInProgressResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListEventInProgressResponseBodyData) SetTitle(v string) *ListEventInProgressResponseBodyData {
	s.Title = &v
	return s
}

type ListEventInProgressResponseBodyDataEventUpdates struct {
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	PublishTime *int64  `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
}

func (s ListEventInProgressResponseBodyDataEventUpdates) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBodyDataEventUpdates) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBodyDataEventUpdates) SetContent(v string) *ListEventInProgressResponseBodyDataEventUpdates {
	s.Content = &v
	return s
}

func (s *ListEventInProgressResponseBodyDataEventUpdates) SetPublishTime(v int64) *ListEventInProgressResponseBodyDataEventUpdates {
	s.PublishTime = &v
	return s
}

type ListEventInProgressResponseBodyDataImpacts struct {
	Product      *ListEventInProgressResponseBodyDataImpactsProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Struct"`
	RecoveryTime *int64                                             `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	Region       *ListEventInProgressResponseBodyDataImpactsRegion  `json:"Region,omitempty" xml:"Region,omitempty" type:"Struct"`
	StartTime    *int64                                             `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListEventInProgressResponseBodyDataImpacts) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBodyDataImpacts) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBodyDataImpacts) SetProduct(v *ListEventInProgressResponseBodyDataImpactsProduct) *ListEventInProgressResponseBodyDataImpacts {
	s.Product = v
	return s
}

func (s *ListEventInProgressResponseBodyDataImpacts) SetRecoveryTime(v int64) *ListEventInProgressResponseBodyDataImpacts {
	s.RecoveryTime = &v
	return s
}

func (s *ListEventInProgressResponseBodyDataImpacts) SetRegion(v *ListEventInProgressResponseBodyDataImpactsRegion) *ListEventInProgressResponseBodyDataImpacts {
	s.Region = v
	return s
}

func (s *ListEventInProgressResponseBodyDataImpacts) SetStartTime(v int64) *ListEventInProgressResponseBodyDataImpacts {
	s.StartTime = &v
	return s
}

type ListEventInProgressResponseBodyDataImpactsProduct struct {
	ProductId   *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s ListEventInProgressResponseBodyDataImpactsProduct) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBodyDataImpactsProduct) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBodyDataImpactsProduct) SetProductId(v string) *ListEventInProgressResponseBodyDataImpactsProduct {
	s.ProductId = &v
	return s
}

func (s *ListEventInProgressResponseBodyDataImpactsProduct) SetProductName(v string) *ListEventInProgressResponseBodyDataImpactsProduct {
	s.ProductName = &v
	return s
}

type ListEventInProgressResponseBodyDataImpactsRegion struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListEventInProgressResponseBodyDataImpactsRegion) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponseBodyDataImpactsRegion) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponseBodyDataImpactsRegion) SetRegionId(v string) *ListEventInProgressResponseBodyDataImpactsRegion {
	s.RegionId = &v
	return s
}

func (s *ListEventInProgressResponseBodyDataImpactsRegion) SetRegionName(v string) *ListEventInProgressResponseBodyDataImpactsRegion {
	s.RegionName = &v
	return s
}

type ListEventInProgressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventInProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventInProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventInProgressResponse) GoString() string {
	return s.String()
}

func (s *ListEventInProgressResponse) SetHeaders(v map[string]*string) *ListEventInProgressResponse {
	s.Headers = v
	return s
}

func (s *ListEventInProgressResponse) SetStatusCode(v int32) *ListEventInProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventInProgressResponse) SetBody(v *ListEventInProgressResponseBody) *ListEventInProgressResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("status"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ListEventInProgressWithOptions(tmpReq *ListEventInProgressRequest, runtime *util.RuntimeOptions) (_result *ListEventInProgressResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListEventInProgressShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		body["RegionIds"] = request.RegionIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventInProgress"),
		Version:     tea.String("2020-01-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventInProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEventInProgress(request *ListEventInProgressRequest) (_result *ListEventInProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventInProgressResponse{}
	_body, _err := client.ListEventInProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
