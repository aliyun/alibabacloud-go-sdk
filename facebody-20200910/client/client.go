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

type DetectIPCPedestrianOptimizedRequest struct {
	// image data
	ImageData []byte `json:"imageData,omitempty" xml:"imageData,omitempty"`
	// image width
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
	// image height
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
}

func (s DetectIPCPedestrianOptimizedRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedRequest) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedRequest) SetImageData(v []byte) *DetectIPCPedestrianOptimizedRequest {
	s.ImageData = v
	return s
}

func (s *DetectIPCPedestrianOptimizedRequest) SetWidth(v int64) *DetectIPCPedestrianOptimizedRequest {
	s.Width = &v
	return s
}

func (s *DetectIPCPedestrianOptimizedRequest) SetHeight(v int64) *DetectIPCPedestrianOptimizedRequest {
	s.Height = &v
	return s
}

type DetectIPCPedestrianOptimizedResponseBody struct {
	// data
	Data *DetectIPCPedestrianOptimizedResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectIPCPedestrianOptimizedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedResponseBody) SetData(v *DetectIPCPedestrianOptimizedResponseBodyData) *DetectIPCPedestrianOptimizedResponseBody {
	s.Data = v
	return s
}

func (s *DetectIPCPedestrianOptimizedResponseBody) SetRequestId(v string) *DetectIPCPedestrianOptimizedResponseBody {
	s.RequestId = &v
	return s
}

type DetectIPCPedestrianOptimizedResponseBodyData struct {
	// imageInfoList
	ImageInfoList []*DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList `json:"ImageInfoList,omitempty" xml:"ImageInfoList,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianOptimizedResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedResponseBodyData) SetImageInfoList(v []*DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList) *DetectIPCPedestrianOptimizedResponseBodyData {
	s.ImageInfoList = v
	return s
}

type DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList struct {
	// Elements
	Elements []*DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList) SetElements(v []*DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements) *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoList {
	s.Elements = v
	return s
}

type DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements struct {
	// box array
	Boxes []*int64 `json:"Boxes,omitempty" xml:"Boxes,omitempty" type:"Repeated"`
	// score
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements) SetBoxes(v []*int64) *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements {
	s.Boxes = v
	return s
}

func (s *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements) SetScore(v float32) *DetectIPCPedestrianOptimizedResponseBodyDataImageInfoListElements {
	s.Score = &v
	return s
}

type DetectIPCPedestrianOptimizedResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectIPCPedestrianOptimizedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectIPCPedestrianOptimizedResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectIPCPedestrianOptimizedResponse) GoString() string {
	return s.String()
}

func (s *DetectIPCPedestrianOptimizedResponse) SetHeaders(v map[string]*string) *DetectIPCPedestrianOptimizedResponse {
	s.Headers = v
	return s
}

func (s *DetectIPCPedestrianOptimizedResponse) SetBody(v *DetectIPCPedestrianOptimizedResponseBody) *DetectIPCPedestrianOptimizedResponse {
	s.Body = v
	return s
}

type ExecuteServerSideVerificationRequest struct {
	CertificateName   *string `json:"certificateName,omitempty" xml:"certificateName,omitempty"`
	CertificateNumber *string `json:"certificateNumber,omitempty" xml:"certificateNumber,omitempty"`
	FacialPictureData []byte  `json:"facialPictureData,omitempty" xml:"facialPictureData,omitempty"`
	FacialPictureUrl  *string `json:"facialPictureUrl,omitempty" xml:"facialPictureUrl,omitempty"`
	SceneType         *string `json:"sceneType,omitempty" xml:"sceneType,omitempty"`
}

func (s ExecuteServerSideVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteServerSideVerificationRequest) GoString() string {
	return s.String()
}

func (s *ExecuteServerSideVerificationRequest) SetCertificateName(v string) *ExecuteServerSideVerificationRequest {
	s.CertificateName = &v
	return s
}

func (s *ExecuteServerSideVerificationRequest) SetCertificateNumber(v string) *ExecuteServerSideVerificationRequest {
	s.CertificateNumber = &v
	return s
}

func (s *ExecuteServerSideVerificationRequest) SetFacialPictureData(v []byte) *ExecuteServerSideVerificationRequest {
	s.FacialPictureData = v
	return s
}

func (s *ExecuteServerSideVerificationRequest) SetFacialPictureUrl(v string) *ExecuteServerSideVerificationRequest {
	s.FacialPictureUrl = &v
	return s
}

func (s *ExecuteServerSideVerificationRequest) SetSceneType(v string) *ExecuteServerSideVerificationRequest {
	s.SceneType = &v
	return s
}

type ExecuteServerSideVerificationResponseBody struct {
	// Id of the request
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ExecuteServerSideVerificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ExecuteServerSideVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteServerSideVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteServerSideVerificationResponseBody) SetRequestId(v string) *ExecuteServerSideVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteServerSideVerificationResponseBody) SetData(v *ExecuteServerSideVerificationResponseBodyData) *ExecuteServerSideVerificationResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteServerSideVerificationResponseBody) SetCode(v string) *ExecuteServerSideVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteServerSideVerificationResponseBody) SetMessage(v string) *ExecuteServerSideVerificationResponseBody {
	s.Message = &v
	return s
}

type ExecuteServerSideVerificationResponseBodyData struct {
	Pass              *bool   `json:"Pass,omitempty" xml:"Pass,omitempty"`
	VerificationToken *string `json:"VerificationToken,omitempty" xml:"VerificationToken,omitempty"`
	Reason            *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ExecuteServerSideVerificationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteServerSideVerificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteServerSideVerificationResponseBodyData) SetPass(v bool) *ExecuteServerSideVerificationResponseBodyData {
	s.Pass = &v
	return s
}

func (s *ExecuteServerSideVerificationResponseBodyData) SetVerificationToken(v string) *ExecuteServerSideVerificationResponseBodyData {
	s.VerificationToken = &v
	return s
}

func (s *ExecuteServerSideVerificationResponseBodyData) SetReason(v string) *ExecuteServerSideVerificationResponseBodyData {
	s.Reason = &v
	return s
}

type ExecuteServerSideVerificationResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteServerSideVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteServerSideVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteServerSideVerificationResponse) GoString() string {
	return s.String()
}

func (s *ExecuteServerSideVerificationResponse) SetHeaders(v map[string]*string) *ExecuteServerSideVerificationResponse {
	s.Headers = v
	return s
}

func (s *ExecuteServerSideVerificationResponse) SetBody(v *ExecuteServerSideVerificationResponseBody) *ExecuteServerSideVerificationResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("facebody"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
 * 行人检测快速版
 */
func (client *Client) DetectIPCPedestrianOptimized(request *DetectIPCPedestrianOptimizedRequest) (_result *DetectIPCPedestrianOptimizedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DetectIPCPedestrianOptimizedResponse{}
	_body, _err := client.DetectIPCPedestrianOptimizedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectIPCPedestrianOptimizedWithOptions(request *DetectIPCPedestrianOptimizedRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DetectIPCPedestrianOptimizedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageData)) {
		body["imageData"] = request.ImageData
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["width"] = request.Width
	}

	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["height"] = request.Height
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &DetectIPCPedestrianOptimizedResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("DetectIPCPedestrianOptimized"), tea.String("2020-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/viapi/k8s/facebody/detect-ipc-pedestrian-optimized"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteServerSideVerification(request *ExecuteServerSideVerificationRequest) (_result *ExecuteServerSideVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteServerSideVerificationResponse{}
	_body, _err := client.ExecuteServerSideVerificationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteServerSideVerificationWithOptions(request *ExecuteServerSideVerificationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExecuteServerSideVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		body["certificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateNumber)) {
		body["certificateNumber"] = request.CertificateNumber
	}

	if !tea.BoolValue(util.IsUnset(request.FacialPictureData)) {
		body["facialPictureData"] = request.FacialPictureData
	}

	if !tea.BoolValue(util.IsUnset(request.FacialPictureUrl)) {
		body["facialPictureUrl"] = request.FacialPictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SceneType)) {
		body["sceneType"] = request.SceneType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &ExecuteServerSideVerificationResponse{}
	_body, _err := client.DoROARequestWithForm(tea.String("ExecuteServerSideVerification"), tea.String("2020-09-10"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/viapi/thirdparty/realperson/execServerSideVerification"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
