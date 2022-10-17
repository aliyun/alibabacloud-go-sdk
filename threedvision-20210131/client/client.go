// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type EstimateMonocularImageDepthRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EstimateMonocularImageDepthRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthRequest) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthRequest) SetImageURL(v string) *EstimateMonocularImageDepthRequest {
	s.ImageURL = &v
	return s
}

type EstimateMonocularImageDepthAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s EstimateMonocularImageDepthAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthAdvanceRequest) SetImageURLObject(v io.Reader) *EstimateMonocularImageDepthAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type EstimateMonocularImageDepthResponseBody struct {
	Data      *EstimateMonocularImageDepthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EstimateMonocularImageDepthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthResponseBody) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthResponseBody) SetData(v *EstimateMonocularImageDepthResponseBodyData) *EstimateMonocularImageDepthResponseBody {
	s.Data = v
	return s
}

func (s *EstimateMonocularImageDepthResponseBody) SetRequestId(v string) *EstimateMonocularImageDepthResponseBody {
	s.RequestId = &v
	return s
}

type EstimateMonocularImageDepthResponseBodyData struct {
	DepthMapUrl     *string `json:"DepthMapUrl,omitempty" xml:"DepthMapUrl,omitempty"`
	DepthToColorUrl *string `json:"DepthToColorUrl,omitempty" xml:"DepthToColorUrl,omitempty"`
}

func (s EstimateMonocularImageDepthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthResponseBodyData) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthResponseBodyData) SetDepthMapUrl(v string) *EstimateMonocularImageDepthResponseBodyData {
	s.DepthMapUrl = &v
	return s
}

func (s *EstimateMonocularImageDepthResponseBodyData) SetDepthToColorUrl(v string) *EstimateMonocularImageDepthResponseBodyData {
	s.DepthToColorUrl = &v
	return s
}

type EstimateMonocularImageDepthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EstimateMonocularImageDepthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EstimateMonocularImageDepthResponse) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthResponse) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthResponse) SetHeaders(v map[string]*string) *EstimateMonocularImageDepthResponse {
	s.Headers = v
	return s
}

func (s *EstimateMonocularImageDepthResponse) SetStatusCode(v int32) *EstimateMonocularImageDepthResponse {
	s.StatusCode = &v
	return s
}

func (s *EstimateMonocularImageDepthResponse) SetBody(v *EstimateMonocularImageDepthResponseBody) *EstimateMonocularImageDepthResponse {
	s.Body = v
	return s
}

type EstimateMonocularVideoDepthRequest struct {
	SampleRate *string `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	VideoURL   *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EstimateMonocularVideoDepthRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularVideoDepthRequest) GoString() string {
	return s.String()
}

func (s *EstimateMonocularVideoDepthRequest) SetSampleRate(v string) *EstimateMonocularVideoDepthRequest {
	s.SampleRate = &v
	return s
}

func (s *EstimateMonocularVideoDepthRequest) SetVideoURL(v string) *EstimateMonocularVideoDepthRequest {
	s.VideoURL = &v
	return s
}

type EstimateMonocularVideoDepthAdvanceRequest struct {
	SampleRate     *string   `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	VideoURLObject io.Reader `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s EstimateMonocularVideoDepthAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularVideoDepthAdvanceRequest) GoString() string {
	return s.String()
}

func (s *EstimateMonocularVideoDepthAdvanceRequest) SetSampleRate(v string) *EstimateMonocularVideoDepthAdvanceRequest {
	s.SampleRate = &v
	return s
}

func (s *EstimateMonocularVideoDepthAdvanceRequest) SetVideoURLObject(v io.Reader) *EstimateMonocularVideoDepthAdvanceRequest {
	s.VideoURLObject = v
	return s
}

type EstimateMonocularVideoDepthResponseBody struct {
	Data      *EstimateMonocularVideoDepthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EstimateMonocularVideoDepthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularVideoDepthResponseBody) GoString() string {
	return s.String()
}

func (s *EstimateMonocularVideoDepthResponseBody) SetData(v *EstimateMonocularVideoDepthResponseBodyData) *EstimateMonocularVideoDepthResponseBody {
	s.Data = v
	return s
}

func (s *EstimateMonocularVideoDepthResponseBody) SetMessage(v string) *EstimateMonocularVideoDepthResponseBody {
	s.Message = &v
	return s
}

func (s *EstimateMonocularVideoDepthResponseBody) SetRequestId(v string) *EstimateMonocularVideoDepthResponseBody {
	s.RequestId = &v
	return s
}

type EstimateMonocularVideoDepthResponseBodyData struct {
	DepthUrl    *string `json:"DepthUrl,omitempty" xml:"DepthUrl,omitempty"`
	DepthVisUrl *string `json:"DepthVisUrl,omitempty" xml:"DepthVisUrl,omitempty"`
}

func (s EstimateMonocularVideoDepthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularVideoDepthResponseBodyData) GoString() string {
	return s.String()
}

func (s *EstimateMonocularVideoDepthResponseBodyData) SetDepthUrl(v string) *EstimateMonocularVideoDepthResponseBodyData {
	s.DepthUrl = &v
	return s
}

func (s *EstimateMonocularVideoDepthResponseBodyData) SetDepthVisUrl(v string) *EstimateMonocularVideoDepthResponseBodyData {
	s.DepthVisUrl = &v
	return s
}

type EstimateMonocularVideoDepthResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EstimateMonocularVideoDepthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EstimateMonocularVideoDepthResponse) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularVideoDepthResponse) GoString() string {
	return s.String()
}

func (s *EstimateMonocularVideoDepthResponse) SetHeaders(v map[string]*string) *EstimateMonocularVideoDepthResponse {
	s.Headers = v
	return s
}

func (s *EstimateMonocularVideoDepthResponse) SetStatusCode(v int32) *EstimateMonocularVideoDepthResponse {
	s.StatusCode = &v
	return s
}

func (s *EstimateMonocularVideoDepthResponse) SetBody(v *EstimateMonocularVideoDepthResponseBody) *EstimateMonocularVideoDepthResponse {
	s.Body = v
	return s
}

type EstimateStereoImageDepthRequest struct {
	LeftImageURL  *string `json:"LeftImageURL,omitempty" xml:"LeftImageURL,omitempty"`
	RightImageURL *string `json:"RightImageURL,omitempty" xml:"RightImageURL,omitempty"`
}

func (s EstimateStereoImageDepthRequest) String() string {
	return tea.Prettify(s)
}

func (s EstimateStereoImageDepthRequest) GoString() string {
	return s.String()
}

func (s *EstimateStereoImageDepthRequest) SetLeftImageURL(v string) *EstimateStereoImageDepthRequest {
	s.LeftImageURL = &v
	return s
}

func (s *EstimateStereoImageDepthRequest) SetRightImageURL(v string) *EstimateStereoImageDepthRequest {
	s.RightImageURL = &v
	return s
}

type EstimateStereoImageDepthResponseBody struct {
	Data      *EstimateStereoImageDepthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EstimateStereoImageDepthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimateStereoImageDepthResponseBody) GoString() string {
	return s.String()
}

func (s *EstimateStereoImageDepthResponseBody) SetData(v *EstimateStereoImageDepthResponseBodyData) *EstimateStereoImageDepthResponseBody {
	s.Data = v
	return s
}

func (s *EstimateStereoImageDepthResponseBody) SetRequestId(v string) *EstimateStereoImageDepthResponseBody {
	s.RequestId = &v
	return s
}

type EstimateStereoImageDepthResponseBodyData struct {
	DisparityMapURL *string `json:"DisparityMapURL,omitempty" xml:"DisparityMapURL,omitempty"`
	DisparityVisURL *string `json:"DisparityVisURL,omitempty" xml:"DisparityVisURL,omitempty"`
}

func (s EstimateStereoImageDepthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s EstimateStereoImageDepthResponseBodyData) GoString() string {
	return s.String()
}

func (s *EstimateStereoImageDepthResponseBodyData) SetDisparityMapURL(v string) *EstimateStereoImageDepthResponseBodyData {
	s.DisparityMapURL = &v
	return s
}

func (s *EstimateStereoImageDepthResponseBodyData) SetDisparityVisURL(v string) *EstimateStereoImageDepthResponseBodyData {
	s.DisparityVisURL = &v
	return s
}

type EstimateStereoImageDepthResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EstimateStereoImageDepthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EstimateStereoImageDepthResponse) String() string {
	return tea.Prettify(s)
}

func (s EstimateStereoImageDepthResponse) GoString() string {
	return s.String()
}

func (s *EstimateStereoImageDepthResponse) SetHeaders(v map[string]*string) *EstimateStereoImageDepthResponse {
	s.Headers = v
	return s
}

func (s *EstimateStereoImageDepthResponse) SetStatusCode(v int32) *EstimateStereoImageDepthResponse {
	s.StatusCode = &v
	return s
}

func (s *EstimateStereoImageDepthResponse) SetBody(v *EstimateStereoImageDepthResponseBody) *EstimateStereoImageDepthResponse {
	s.Body = v
	return s
}

type GetAsyncJobResultRequest struct {
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetJobId(v string) *GetAsyncJobResultRequest {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponseBody struct {
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncJobResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponse) SetHeaders(v map[string]*string) *GetAsyncJobResultResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncJobResultResponse) SetStatusCode(v int32) *GetAsyncJobResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type ReconstructBodyBySingleImageRequest struct {
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ReconstructBodyBySingleImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageRequest) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageRequest) SetImageURL(v string) *ReconstructBodyBySingleImageRequest {
	s.ImageURL = &v
	return s
}

type ReconstructBodyBySingleImageAdvanceRequest struct {
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ReconstructBodyBySingleImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageAdvanceRequest) SetImageURLObject(v io.Reader) *ReconstructBodyBySingleImageAdvanceRequest {
	s.ImageURLObject = v
	return s
}

type ReconstructBodyBySingleImageResponseBody struct {
	Data      *ReconstructBodyBySingleImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReconstructBodyBySingleImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageResponseBody) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageResponseBody) SetData(v *ReconstructBodyBySingleImageResponseBodyData) *ReconstructBodyBySingleImageResponseBody {
	s.Data = v
	return s
}

func (s *ReconstructBodyBySingleImageResponseBody) SetRequestId(v string) *ReconstructBodyBySingleImageResponseBody {
	s.RequestId = &v
	return s
}

type ReconstructBodyBySingleImageResponseBodyData struct {
	DepthURL *string `json:"DepthURL,omitempty" xml:"DepthURL,omitempty"`
	MeshURL  *string `json:"MeshURL,omitempty" xml:"MeshURL,omitempty"`
}

func (s ReconstructBodyBySingleImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageResponseBodyData) SetDepthURL(v string) *ReconstructBodyBySingleImageResponseBodyData {
	s.DepthURL = &v
	return s
}

func (s *ReconstructBodyBySingleImageResponseBodyData) SetMeshURL(v string) *ReconstructBodyBySingleImageResponseBodyData {
	s.MeshURL = &v
	return s
}

type ReconstructBodyBySingleImageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReconstructBodyBySingleImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReconstructBodyBySingleImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageResponse) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageResponse) SetHeaders(v map[string]*string) *ReconstructBodyBySingleImageResponse {
	s.Headers = v
	return s
}

func (s *ReconstructBodyBySingleImageResponse) SetStatusCode(v int32) *ReconstructBodyBySingleImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ReconstructBodyBySingleImageResponse) SetBody(v *ReconstructBodyBySingleImageResponseBody) *ReconstructBodyBySingleImageResponse {
	s.Body = v
	return s
}

type ReconstructThreeDMultiViewRequest struct {
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ZipFileUrl *string `json:"ZipFileUrl,omitempty" xml:"ZipFileUrl,omitempty"`
}

func (s ReconstructThreeDMultiViewRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewRequest) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewRequest) SetMode(v string) *ReconstructThreeDMultiViewRequest {
	s.Mode = &v
	return s
}

func (s *ReconstructThreeDMultiViewRequest) SetZipFileUrl(v string) *ReconstructThreeDMultiViewRequest {
	s.ZipFileUrl = &v
	return s
}

type ReconstructThreeDMultiViewAdvanceRequest struct {
	Mode             *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ZipFileUrlObject io.Reader `json:"ZipFileUrl,omitempty" xml:"ZipFileUrl,omitempty"`
}

func (s ReconstructThreeDMultiViewAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewAdvanceRequest) SetMode(v string) *ReconstructThreeDMultiViewAdvanceRequest {
	s.Mode = &v
	return s
}

func (s *ReconstructThreeDMultiViewAdvanceRequest) SetZipFileUrlObject(v io.Reader) *ReconstructThreeDMultiViewAdvanceRequest {
	s.ZipFileUrlObject = v
	return s
}

type ReconstructThreeDMultiViewResponseBody struct {
	Data      *ReconstructThreeDMultiViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReconstructThreeDMultiViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewResponseBody) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewResponseBody) SetData(v *ReconstructThreeDMultiViewResponseBodyData) *ReconstructThreeDMultiViewResponseBody {
	s.Data = v
	return s
}

func (s *ReconstructThreeDMultiViewResponseBody) SetRequestId(v string) *ReconstructThreeDMultiViewResponseBody {
	s.RequestId = &v
	return s
}

type ReconstructThreeDMultiViewResponseBodyData struct {
	PointCloudURL *string `json:"PointCloudURL,omitempty" xml:"PointCloudURL,omitempty"`
}

func (s ReconstructThreeDMultiViewResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewResponseBodyData) SetPointCloudURL(v string) *ReconstructThreeDMultiViewResponseBodyData {
	s.PointCloudURL = &v
	return s
}

type ReconstructThreeDMultiViewResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReconstructThreeDMultiViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReconstructThreeDMultiViewResponse) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewResponse) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewResponse) SetHeaders(v map[string]*string) *ReconstructThreeDMultiViewResponse {
	s.Headers = v
	return s
}

func (s *ReconstructThreeDMultiViewResponse) SetStatusCode(v int32) *ReconstructThreeDMultiViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ReconstructThreeDMultiViewResponse) SetBody(v *ReconstructThreeDMultiViewResponseBody) *ReconstructThreeDMultiViewResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("threedvision"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) EstimateMonocularImageDepthWithOptions(request *EstimateMonocularImageDepthRequest, runtime *util.RuntimeOptions) (_result *EstimateMonocularImageDepthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EstimateMonocularImageDepth"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EstimateMonocularImageDepthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EstimateMonocularImageDepth(request *EstimateMonocularImageDepthRequest) (_result *EstimateMonocularImageDepthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EstimateMonocularImageDepthResponse{}
	_body, _err := client.EstimateMonocularImageDepthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EstimateMonocularImageDepthAdvance(request *EstimateMonocularImageDepthAdvanceRequest, runtime *util.RuntimeOptions) (_result *EstimateMonocularImageDepthResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("threedvision"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	estimateMonocularImageDepthReq := &EstimateMonocularImageDepthRequest{}
	openapiutil.Convert(request, estimateMonocularImageDepthReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		estimateMonocularImageDepthReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	estimateMonocularImageDepthResp, _err := client.EstimateMonocularImageDepthWithOptions(estimateMonocularImageDepthReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = estimateMonocularImageDepthResp
	return _result, _err
}

func (client *Client) EstimateMonocularVideoDepthWithOptions(request *EstimateMonocularVideoDepthRequest, runtime *util.RuntimeOptions) (_result *EstimateMonocularVideoDepthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SampleRate)) {
		body["SampleRate"] = request.SampleRate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoURL)) {
		body["VideoURL"] = request.VideoURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EstimateMonocularVideoDepth"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EstimateMonocularVideoDepthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EstimateMonocularVideoDepth(request *EstimateMonocularVideoDepthRequest) (_result *EstimateMonocularVideoDepthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EstimateMonocularVideoDepthResponse{}
	_body, _err := client.EstimateMonocularVideoDepthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EstimateMonocularVideoDepthAdvance(request *EstimateMonocularVideoDepthAdvanceRequest, runtime *util.RuntimeOptions) (_result *EstimateMonocularVideoDepthResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("threedvision"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	estimateMonocularVideoDepthReq := &EstimateMonocularVideoDepthRequest{}
	openapiutil.Convert(request, estimateMonocularVideoDepthReq)
	if !tea.BoolValue(util.IsUnset(request.VideoURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.VideoURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		estimateMonocularVideoDepthReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	estimateMonocularVideoDepthResp, _err := client.EstimateMonocularVideoDepthWithOptions(estimateMonocularVideoDepthReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = estimateMonocularVideoDepthResp
	return _result, _err
}

func (client *Client) EstimateStereoImageDepthWithOptions(request *EstimateStereoImageDepthRequest, runtime *util.RuntimeOptions) (_result *EstimateStereoImageDepthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LeftImageURL)) {
		body["LeftImageURL"] = request.LeftImageURL
	}

	if !tea.BoolValue(util.IsUnset(request.RightImageURL)) {
		body["RightImageURL"] = request.RightImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EstimateStereoImageDepth"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EstimateStereoImageDepthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EstimateStereoImageDepth(request *EstimateStereoImageDepthRequest) (_result *EstimateStereoImageDepthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EstimateStereoImageDepthResponse{}
	_body, _err := client.EstimateStereoImageDepthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncJobResult"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncJobResult(request *GetAsyncJobResultRequest) (_result *GetAsyncJobResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.GetAsyncJobResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReconstructBodyBySingleImageWithOptions(request *ReconstructBodyBySingleImageRequest, runtime *util.RuntimeOptions) (_result *ReconstructBodyBySingleImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageURL)) {
		body["ImageURL"] = request.ImageURL
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReconstructBodyBySingleImage"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReconstructBodyBySingleImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReconstructBodyBySingleImage(request *ReconstructBodyBySingleImageRequest) (_result *ReconstructBodyBySingleImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReconstructBodyBySingleImageResponse{}
	_body, _err := client.ReconstructBodyBySingleImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReconstructBodyBySingleImageAdvance(request *ReconstructBodyBySingleImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *ReconstructBodyBySingleImageResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("threedvision"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	reconstructBodyBySingleImageReq := &ReconstructBodyBySingleImageRequest{}
	openapiutil.Convert(request, reconstructBodyBySingleImageReq)
	if !tea.BoolValue(util.IsUnset(request.ImageURLObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ImageURLObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		reconstructBodyBySingleImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	reconstructBodyBySingleImageResp, _err := client.ReconstructBodyBySingleImageWithOptions(reconstructBodyBySingleImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = reconstructBodyBySingleImageResp
	return _result, _err
}

func (client *Client) ReconstructThreeDMultiViewWithOptions(request *ReconstructThreeDMultiViewRequest, runtime *util.RuntimeOptions) (_result *ReconstructThreeDMultiViewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		body["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.ZipFileUrl)) {
		body["ZipFileUrl"] = request.ZipFileUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReconstructThreeDMultiView"),
		Version:     tea.String("2021-01-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReconstructThreeDMultiViewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReconstructThreeDMultiView(request *ReconstructThreeDMultiViewRequest) (_result *ReconstructThreeDMultiViewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReconstructThreeDMultiViewResponse{}
	_body, _err := client.ReconstructThreeDMultiViewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReconstructThreeDMultiViewAdvance(request *ReconstructThreeDMultiViewAdvanceRequest, runtime *util.RuntimeOptions) (_result *ReconstructThreeDMultiViewResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.IsUnset(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("threedvision"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	var ossClient *oss.Client
	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	reconstructThreeDMultiViewReq := &ReconstructThreeDMultiViewRequest{}
	openapiutil.Convert(request, reconstructThreeDMultiViewReq)
	if !tea.BoolValue(util.IsUnset(request.ZipFileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.ZipFileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		reconstructThreeDMultiViewReq.ZipFileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	reconstructThreeDMultiViewResp, _err := client.ReconstructThreeDMultiViewWithOptions(reconstructThreeDMultiViewReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = reconstructThreeDMultiViewResp
	return _result, _err
}
