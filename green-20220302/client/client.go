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

type ImageModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageModerationRequest) SetService(v string) *ImageModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageModerationRequest) SetServiceParameters(v string) *ImageModerationRequest {
	s.ServiceParameters = &v
	return s
}

type ImageModerationResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ImageModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Msg       *string                          `json:"Msg,omitempty" xml:"Msg,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBody) SetCode(v int32) *ImageModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageModerationResponseBody) SetData(v *ImageModerationResponseBodyData) *ImageModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageModerationResponseBody) SetMsg(v string) *ImageModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageModerationResponseBody) SetRequestId(v string) *ImageModerationResponseBody {
	s.RequestId = &v
	return s
}

type ImageModerationResponseBodyData struct {
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Result []*ImageModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ImageModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyData) SetDataId(v string) *ImageModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageModerationResponseBodyData) SetResult(v []*ImageModerationResponseBodyDataResult) *ImageModerationResponseBodyData {
	s.Result = v
	return s
}

type ImageModerationResponseBodyDataResult struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Label      *string  `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ImageModerationResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageModerationResponseBodyDataResult) SetConfidence(v float32) *ImageModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageModerationResponseBodyDataResult) SetLabel(v string) *ImageModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

type ImageModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImageModerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImageModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageModerationResponse) GoString() string {
	return s.String()
}

func (s *ImageModerationResponse) SetHeaders(v map[string]*string) *ImageModerationResponse {
	s.Headers = v
	return s
}

func (s *ImageModerationResponse) SetStatusCode(v int32) *ImageModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageModerationResponse) SetBody(v *ImageModerationResponseBody) *ImageModerationResponse {
	s.Body = v
	return s
}

type TextModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s TextModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s TextModerationRequest) GoString() string {
	return s.String()
}

func (s *TextModerationRequest) SetService(v string) *TextModerationRequest {
	s.Service = &v
	return s
}

func (s *TextModerationRequest) SetServiceParameters(v string) *TextModerationRequest {
	s.ServiceParameters = &v
	return s
}

type TextModerationResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *TextModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBody) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBody) SetCode(v int32) *TextModerationResponseBody {
	s.Code = &v
	return s
}

func (s *TextModerationResponseBody) SetData(v *TextModerationResponseBodyData) *TextModerationResponseBody {
	s.Data = v
	return s
}

func (s *TextModerationResponseBody) SetMessage(v string) *TextModerationResponseBody {
	s.Message = &v
	return s
}

func (s *TextModerationResponseBody) SetRequestId(v string) *TextModerationResponseBody {
	s.RequestId = &v
	return s
}

type TextModerationResponseBodyData struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	DeviceId  *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	Labels    *string `json:"labels,omitempty" xml:"labels,omitempty"`
	Reason    *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s TextModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextModerationResponseBodyData) SetAccountId(v string) *TextModerationResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetDeviceId(v string) *TextModerationResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *TextModerationResponseBodyData) SetLabels(v string) *TextModerationResponseBodyData {
	s.Labels = &v
	return s
}

func (s *TextModerationResponseBodyData) SetReason(v string) *TextModerationResponseBodyData {
	s.Reason = &v
	return s
}

type TextModerationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TextModerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TextModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s TextModerationResponse) GoString() string {
	return s.String()
}

func (s *TextModerationResponse) SetHeaders(v map[string]*string) *TextModerationResponse {
	s.Headers = v
	return s
}

func (s *TextModerationResponse) SetStatusCode(v int32) *TextModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *TextModerationResponse) SetBody(v *TextModerationResponseBody) *TextModerationResponse {
	s.Body = v
	return s
}

type VoiceModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationRequest) SetService(v string) *VoiceModerationRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationRequest) SetServiceParameters(v string) *VoiceModerationRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationResponseBody struct {
	Code    *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VoiceModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBody) SetCode(v int32) *VoiceModerationResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResponseBody) SetData(v *VoiceModerationResponseBodyData) *VoiceModerationResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResponseBody) SetMessage(v string) *VoiceModerationResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResponseBody) SetRequestId(v string) *VoiceModerationResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VoiceModerationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponseBodyData) SetTaskId(v string) *VoiceModerationResponseBodyData {
	s.TaskId = &v
	return s
}

type VoiceModerationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VoiceModerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VoiceModerationResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResponse) SetHeaders(v map[string]*string) *VoiceModerationResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResponse) SetStatusCode(v int32) *VoiceModerationResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResponse) SetBody(v *VoiceModerationResponseBody) *VoiceModerationResponse {
	s.Body = v
	return s
}

type VoiceModerationCancelRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationCancelRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelRequest) SetService(v string) *VoiceModerationCancelRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationCancelRequest) SetServiceParameters(v string) *VoiceModerationCancelRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationCancelResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationCancelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponseBody) SetCode(v int32) *VoiceModerationCancelResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetMessage(v string) *VoiceModerationCancelResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationCancelResponseBody) SetRequestId(v string) *VoiceModerationCancelResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VoiceModerationCancelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VoiceModerationCancelResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationCancelResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponse) SetHeaders(v map[string]*string) *VoiceModerationCancelResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationCancelResponse) SetStatusCode(v int32) *VoiceModerationCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationCancelResponse) SetBody(v *VoiceModerationCancelResponseBody) *VoiceModerationCancelResponse {
	s.Body = v
	return s
}

type VoiceModerationResultRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VoiceModerationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultRequest) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultRequest) SetService(v string) *VoiceModerationResultRequest {
	s.Service = &v
	return s
}

func (s *VoiceModerationResultRequest) SetServiceParameters(v string) *VoiceModerationResultRequest {
	s.ServiceParameters = &v
	return s
}

type VoiceModerationResultResponseBody struct {
	Code    *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *VoiceModerationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceModerationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBody) SetCode(v int32) *VoiceModerationResultResponseBody {
	s.Code = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetData(v *VoiceModerationResultResponseBodyData) *VoiceModerationResultResponseBody {
	s.Data = v
	return s
}

func (s *VoiceModerationResultResponseBody) SetMessage(v string) *VoiceModerationResultResponseBody {
	s.Message = &v
	return s
}

func (s *VoiceModerationResultResponseBody) SetRequestId(v string) *VoiceModerationResultResponseBody {
	s.RequestId = &v
	return s
}

type VoiceModerationResultResponseBodyData struct {
	LiveId       *string                                              `json:"LiveId,omitempty" xml:"LiveId,omitempty"`
	SliceDetails []*VoiceModerationResultResponseBodyDataSliceDetails `json:"SliceDetails,omitempty" xml:"SliceDetails,omitempty" type:"Repeated"`
	TaskId       *string                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Url          *string                                              `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyData) SetLiveId(v string) *VoiceModerationResultResponseBodyData {
	s.LiveId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetSliceDetails(v []*VoiceModerationResultResponseBodyDataSliceDetails) *VoiceModerationResultResponseBodyData {
	s.SliceDetails = v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetTaskId(v string) *VoiceModerationResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VoiceModerationResultResponseBodyData) SetUrl(v string) *VoiceModerationResultResponseBodyData {
	s.Url = &v
	return s
}

type VoiceModerationResultResponseBodyDataSliceDetails struct {
	EndTime          *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimestamp     *int64                 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	Extend           *string                `json:"Extend,omitempty" xml:"Extend,omitempty"`
	Labels           *string                `json:"Labels,omitempty" xml:"Labels,omitempty"`
	OriginAlgoResult map[string]interface{} `json:"OriginAlgoResult,omitempty" xml:"OriginAlgoResult,omitempty"`
	RiskTips         *string                `json:"RiskTips,omitempty" xml:"RiskTips,omitempty"`
	RiskWords        *string                `json:"RiskWords,omitempty" xml:"RiskWords,omitempty"`
	Score            *float32               `json:"Score,omitempty" xml:"Score,omitempty"`
	StartTime        *int64                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StartTimestamp   *int64                 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	Text             *string                `json:"Text,omitempty" xml:"Text,omitempty"`
	Url              *string                `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponseBodyDataSliceDetails) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetEndTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.EndTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetExtend(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Extend = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetLabels(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Labels = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetOriginAlgoResult(v map[string]interface{}) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.OriginAlgoResult = v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskTips(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskTips = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetRiskWords(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.RiskWords = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetScore(v float32) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Score = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTime(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTime = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetStartTimestamp(v int64) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.StartTimestamp = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetText(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Text = &v
	return s
}

func (s *VoiceModerationResultResponseBodyDataSliceDetails) SetUrl(v string) *VoiceModerationResultResponseBodyDataSliceDetails {
	s.Url = &v
	return s
}

type VoiceModerationResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VoiceModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VoiceModerationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceModerationResultResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationResultResponse) SetHeaders(v map[string]*string) *VoiceModerationResultResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationResultResponse) SetStatusCode(v int32) *VoiceModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationResultResponse) SetBody(v *VoiceModerationResultResponseBody) *VoiceModerationResultResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("green.aliyuncs.com"),
		"cn-hongkong":           tea.String("green.aliyuncs.com"),
		"cn-huhehaote":          tea.String("green.aliyuncs.com"),
		"cn-qingdao":            tea.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("green.aliyuncs.com"),
		"eu-central-1":          tea.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ImageModerationWithOptions(request *ImageModerationRequest, runtime *util.RuntimeOptions) (_result *ImageModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImageModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageModeration(request *ImageModerationRequest) (_result *ImageModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageModerationResponse{}
	_body, _err := client.ImageModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextModerationWithOptions(request *TextModerationRequest, runtime *util.RuntimeOptions) (_result *TextModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TextModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TextModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextModeration(request *TextModerationRequest) (_result *TextModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TextModerationResponse{}
	_body, _err := client.TextModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceModerationWithOptions(request *VoiceModerationRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModeration"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceModeration(request *VoiceModerationRequest) (_result *VoiceModerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationResponse{}
	_body, _err := client.VoiceModerationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceModerationCancelWithOptions(request *VoiceModerationCancelRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationCancelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModerationCancel"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceModerationCancel(request *VoiceModerationCancelRequest) (_result *VoiceModerationCancelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationCancelResponse{}
	_body, _err := client.VoiceModerationCancelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceModerationResultWithOptions(request *VoiceModerationResultRequest, runtime *util.RuntimeOptions) (_result *VoiceModerationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Service)) {
		body["Service"] = request.Service
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceModerationResult"),
		Version:     tea.String("2022-03-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceModerationResult(request *VoiceModerationResultRequest) (_result *VoiceModerationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoiceModerationResultResponse{}
	_body, _err := client.VoiceModerationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
