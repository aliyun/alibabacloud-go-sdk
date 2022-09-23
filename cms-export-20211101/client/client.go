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

type BatchGetRequest struct {
	CompressionType    *string `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	Cursor             *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	Length             *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Metric             *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Namespace          *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RecordKeyWhiteList *string `json:"RecordKeyWhiteList,omitempty" xml:"RecordKeyWhiteList,omitempty"`
}

func (s BatchGetRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetRequest) GoString() string {
	return s.String()
}

func (s *BatchGetRequest) SetCompressionType(v string) *BatchGetRequest {
	s.CompressionType = &v
	return s
}

func (s *BatchGetRequest) SetCursor(v string) *BatchGetRequest {
	s.Cursor = &v
	return s
}

func (s *BatchGetRequest) SetLength(v int32) *BatchGetRequest {
	s.Length = &v
	return s
}

func (s *BatchGetRequest) SetMetric(v string) *BatchGetRequest {
	s.Metric = &v
	return s
}

func (s *BatchGetRequest) SetNamespace(v string) *BatchGetRequest {
	s.Namespace = &v
	return s
}

func (s *BatchGetRequest) SetRecordKeyWhiteList(v string) *BatchGetRequest {
	s.RecordKeyWhiteList = &v
	return s
}

type BatchGetResponseBody struct {
	Code      *int32                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *BatchGetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchGetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetResponseBody) SetCode(v int32) *BatchGetResponseBody {
	s.Code = &v
	return s
}

func (s *BatchGetResponseBody) SetData(v *BatchGetResponseBodyData) *BatchGetResponseBody {
	s.Data = v
	return s
}

func (s *BatchGetResponseBody) SetMessage(v string) *BatchGetResponseBody {
	s.Message = &v
	return s
}

func (s *BatchGetResponseBody) SetRequestId(v string) *BatchGetResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetResponseBody) SetSuccess(v bool) *BatchGetResponseBody {
	s.Success = &v
	return s
}

type BatchGetResponseBodyData struct {
	CompressionKeys   []*string                          `json:"CompressionKeys,omitempty" xml:"CompressionKeys,omitempty" type:"Repeated"`
	CompressionValues [][]*string                        `json:"CompressionValues,omitempty" xml:"CompressionValues,omitempty" type:"Repeated"`
	Cursor            *string                            `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	Length            *int64                             `json:"Length,omitempty" xml:"Length,omitempty"`
	Records           []*BatchGetResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	AnchorTs          *int64                             `json:"__anchorTs__,omitempty" xml:"__anchorTs__,omitempty"`
}

func (s BatchGetResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s BatchGetResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchGetResponseBodyData) SetCompressionKeys(v []*string) *BatchGetResponseBodyData {
	s.CompressionKeys = v
	return s
}

func (s *BatchGetResponseBodyData) SetCompressionValues(v [][]*string) *BatchGetResponseBodyData {
	s.CompressionValues = v
	return s
}

func (s *BatchGetResponseBodyData) SetCursor(v string) *BatchGetResponseBodyData {
	s.Cursor = &v
	return s
}

func (s *BatchGetResponseBodyData) SetLength(v int64) *BatchGetResponseBodyData {
	s.Length = &v
	return s
}

func (s *BatchGetResponseBodyData) SetRecords(v []*BatchGetResponseBodyDataRecords) *BatchGetResponseBodyData {
	s.Records = v
	return s
}

func (s *BatchGetResponseBodyData) SetAnchorTs(v int64) *BatchGetResponseBodyData {
	s.AnchorTs = &v
	return s
}

type BatchGetResponseBodyDataRecords struct {
	LabelValues   []*string `json:"LabelValues,omitempty" xml:"LabelValues,omitempty" type:"Repeated"`
	Labels        []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	MeasureLabels []*string `json:"MeasureLabels,omitempty" xml:"MeasureLabels,omitempty" type:"Repeated"`
	MeasureValues []*string `json:"MeasureValues,omitempty" xml:"MeasureValues,omitempty" type:"Repeated"`
	Metric        *string   `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Namespace     *string   `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period        *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	TagValues     []*string `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	Tags          []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Timestamp     *int64    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s BatchGetResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s BatchGetResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *BatchGetResponseBodyDataRecords) SetLabelValues(v []*string) *BatchGetResponseBodyDataRecords {
	s.LabelValues = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetLabels(v []*string) *BatchGetResponseBodyDataRecords {
	s.Labels = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetMeasureLabels(v []*string) *BatchGetResponseBodyDataRecords {
	s.MeasureLabels = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetMeasureValues(v []*string) *BatchGetResponseBodyDataRecords {
	s.MeasureValues = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetMetric(v string) *BatchGetResponseBodyDataRecords {
	s.Metric = &v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetNamespace(v string) *BatchGetResponseBodyDataRecords {
	s.Namespace = &v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetPeriod(v int32) *BatchGetResponseBodyDataRecords {
	s.Period = &v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetTagValues(v []*string) *BatchGetResponseBodyDataRecords {
	s.TagValues = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetTags(v []*string) *BatchGetResponseBodyDataRecords {
	s.Tags = v
	return s
}

func (s *BatchGetResponseBodyDataRecords) SetTimestamp(v int64) *BatchGetResponseBodyDataRecords {
	s.Timestamp = &v
	return s
}

type BatchGetResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetResponse) GoString() string {
	return s.String()
}

func (s *BatchGetResponse) SetHeaders(v map[string]*string) *BatchGetResponse {
	s.Headers = v
	return s
}

func (s *BatchGetResponse) SetStatusCode(v int32) *BatchGetResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetResponse) SetBody(v *BatchGetResponseBody) *BatchGetResponse {
	s.Body = v
	return s
}

type CursorRequest struct {
	EndTime   *int64                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Matchers  []*CursorRequestMatchers `json:"Matchers,omitempty" xml:"Matchers,omitempty" type:"Repeated"`
	Metric    *string                  `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Namespace *string                  `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period    *int32                   `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime *int64                   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CursorRequest) String() string {
	return tea.Prettify(s)
}

func (s CursorRequest) GoString() string {
	return s.String()
}

func (s *CursorRequest) SetEndTime(v int64) *CursorRequest {
	s.EndTime = &v
	return s
}

func (s *CursorRequest) SetMatchers(v []*CursorRequestMatchers) *CursorRequest {
	s.Matchers = v
	return s
}

func (s *CursorRequest) SetMetric(v string) *CursorRequest {
	s.Metric = &v
	return s
}

func (s *CursorRequest) SetNamespace(v string) *CursorRequest {
	s.Namespace = &v
	return s
}

func (s *CursorRequest) SetPeriod(v int32) *CursorRequest {
	s.Period = &v
	return s
}

func (s *CursorRequest) SetStartTime(v int64) *CursorRequest {
	s.StartTime = &v
	return s
}

type CursorRequestMatchers struct {
	Label    *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CursorRequestMatchers) String() string {
	return tea.Prettify(s)
}

func (s CursorRequestMatchers) GoString() string {
	return s.String()
}

func (s *CursorRequestMatchers) SetLabel(v string) *CursorRequestMatchers {
	s.Label = &v
	return s
}

func (s *CursorRequestMatchers) SetOperator(v string) *CursorRequestMatchers {
	s.Operator = &v
	return s
}

func (s *CursorRequestMatchers) SetValue(v string) *CursorRequestMatchers {
	s.Value = &v
	return s
}

type CursorShrinkRequest struct {
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MatchersShrink *string `json:"Matchers,omitempty" xml:"Matchers,omitempty"`
	Metric         *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Namespace      *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Period         *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CursorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CursorShrinkRequest) GoString() string {
	return s.String()
}

func (s *CursorShrinkRequest) SetEndTime(v int64) *CursorShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *CursorShrinkRequest) SetMatchersShrink(v string) *CursorShrinkRequest {
	s.MatchersShrink = &v
	return s
}

func (s *CursorShrinkRequest) SetMetric(v string) *CursorShrinkRequest {
	s.Metric = &v
	return s
}

func (s *CursorShrinkRequest) SetNamespace(v string) *CursorShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CursorShrinkRequest) SetPeriod(v int32) *CursorShrinkRequest {
	s.Period = &v
	return s
}

func (s *CursorShrinkRequest) SetStartTime(v int64) *CursorShrinkRequest {
	s.StartTime = &v
	return s
}

type CursorResponseBody struct {
	Code      *int32                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CursorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CursorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CursorResponseBody) GoString() string {
	return s.String()
}

func (s *CursorResponseBody) SetCode(v int32) *CursorResponseBody {
	s.Code = &v
	return s
}

func (s *CursorResponseBody) SetData(v *CursorResponseBodyData) *CursorResponseBody {
	s.Data = v
	return s
}

func (s *CursorResponseBody) SetMessage(v string) *CursorResponseBody {
	s.Message = &v
	return s
}

func (s *CursorResponseBody) SetRequestId(v string) *CursorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CursorResponseBody) SetSuccess(v bool) *CursorResponseBody {
	s.Success = &v
	return s
}

type CursorResponseBodyData struct {
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
}

func (s CursorResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CursorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CursorResponseBodyData) SetCursor(v string) *CursorResponseBodyData {
	s.Cursor = &v
	return s
}

type CursorResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CursorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CursorResponse) String() string {
	return tea.Prettify(s)
}

func (s CursorResponse) GoString() string {
	return s.String()
}

func (s *CursorResponse) SetHeaders(v map[string]*string) *CursorResponse {
	s.Headers = v
	return s
}

func (s *CursorResponse) SetStatusCode(v int32) *CursorResponse {
	s.StatusCode = &v
	return s
}

func (s *CursorResponse) SetBody(v *CursorResponseBody) *CursorResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cms-export"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BatchGetWithOptions(request *BatchGetRequest, runtime *util.RuntimeOptions) (_result *BatchGetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompressionType)) {
		body["CompressionType"] = request.CompressionType
	}

	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["Cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		body["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		body["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.RecordKeyWhiteList)) {
		body["RecordKeyWhiteList"] = request.RecordKeyWhiteList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGet"),
		Version:     tea.String("2021-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGet(request *BatchGetRequest) (_result *BatchGetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchGetResponse{}
	_body, _err := client.BatchGetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CursorWithOptions(tmpReq *CursorRequest, runtime *util.RuntimeOptions) (_result *CursorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CursorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Matchers)) {
		request.MatchersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Matchers, tea.String("Matchers"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MatchersShrink)) {
		body["Matchers"] = request.MatchersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		body["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		body["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Cursor"),
		Version:     tea.String("2021-11-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CursorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Cursor(request *CursorRequest) (_result *CursorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CursorResponse{}
	_body, _err := client.CursorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
