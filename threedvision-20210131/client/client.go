// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type ReconstructBodyBySingleImageRequest struct {
	// A short description of struct
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	// Id of the request
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ReconstructBodyBySingleImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ReconstructBodyBySingleImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReconstructBodyBySingleImageResponseBody) GoString() string {
	return s.String()
}

func (s *ReconstructBodyBySingleImageResponseBody) SetRequestId(v string) *ReconstructBodyBySingleImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReconstructBodyBySingleImageResponseBody) SetData(v *ReconstructBodyBySingleImageResponseBodyData) *ReconstructBodyBySingleImageResponseBody {
	s.Data = v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReconstructBodyBySingleImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReconstructBodyBySingleImageResponse) SetBody(v *ReconstructBodyBySingleImageResponseBody) *ReconstructBodyBySingleImageResponse {
	s.Body = v
	return s
}

type ReconstructThreeDMultiViewRequest struct {
	// A short description of struct
	ZipFileUrl *string `json:"ZipFileUrl,omitempty" xml:"ZipFileUrl,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ReconstructThreeDMultiViewRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewRequest) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewRequest) SetZipFileUrl(v string) *ReconstructThreeDMultiViewRequest {
	s.ZipFileUrl = &v
	return s
}

func (s *ReconstructThreeDMultiViewRequest) SetMode(v string) *ReconstructThreeDMultiViewRequest {
	s.Mode = &v
	return s
}

type ReconstructThreeDMultiViewAdvanceRequest struct {
	ZipFileUrlObject io.Reader `json:"ZipFileUrlObject,omitempty" xml:"ZipFileUrlObject,omitempty" require:"true"`
	Mode             *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s ReconstructThreeDMultiViewAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewAdvanceRequest) SetZipFileUrlObject(v io.Reader) *ReconstructThreeDMultiViewAdvanceRequest {
	s.ZipFileUrlObject = v
	return s
}

func (s *ReconstructThreeDMultiViewAdvanceRequest) SetMode(v string) *ReconstructThreeDMultiViewAdvanceRequest {
	s.Mode = &v
	return s
}

type ReconstructThreeDMultiViewResponseBody struct {
	// Id of the request
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ReconstructThreeDMultiViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ReconstructThreeDMultiViewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReconstructThreeDMultiViewResponseBody) GoString() string {
	return s.String()
}

func (s *ReconstructThreeDMultiViewResponseBody) SetRequestId(v string) *ReconstructThreeDMultiViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReconstructThreeDMultiViewResponseBody) SetData(v *ReconstructThreeDMultiViewResponseBodyData) *ReconstructThreeDMultiViewResponseBody {
	s.Data = v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReconstructThreeDMultiViewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReconstructThreeDMultiViewResponse) SetBody(v *ReconstructThreeDMultiViewResponseBody) *ReconstructThreeDMultiViewResponse {
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
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GetAsyncJobResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GetAsyncJobResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBody) SetRequestId(v string) *GetAsyncJobResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncJobResultResponseBody) SetData(v *GetAsyncJobResultResponseBodyData) *GetAsyncJobResultResponseBody {
	s.Data = v
	return s
}

type GetAsyncJobResultResponseBodyData struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Result       *string `json:"Result,omitempty" xml:"Result,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultResponseBodyData) SetStatus(v string) *GetAsyncJobResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorMessage(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetResult(v string) *GetAsyncJobResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetErrorCode(v string) *GetAsyncJobResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncJobResultResponseBodyData) SetJobId(v string) *GetAsyncJobResultResponseBodyData {
	s.JobId = &v
	return s
}

type GetAsyncJobResultResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAsyncJobResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetAsyncJobResultResponse) SetBody(v *GetAsyncJobResultResponseBody) *GetAsyncJobResultResponse {
	s.Body = v
	return s
}

type EstimateMonocularImageDepthRequest struct {
	// A short description of struct
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
	ImageURLObject io.Reader `json:"ImageURLObject,omitempty" xml:"ImageURLObject,omitempty" require:"true"`
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
	// Id of the request
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EstimateMonocularImageDepthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EstimateMonocularImageDepthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimateMonocularImageDepthResponseBody) GoString() string {
	return s.String()
}

func (s *EstimateMonocularImageDepthResponseBody) SetRequestId(v string) *EstimateMonocularImageDepthResponseBody {
	s.RequestId = &v
	return s
}

func (s *EstimateMonocularImageDepthResponseBody) SetData(v *EstimateMonocularImageDepthResponseBodyData) *EstimateMonocularImageDepthResponseBody {
	s.Data = v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EstimateMonocularImageDepthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EstimateMonocularImageDepthResponse) SetBody(v *EstimateMonocularImageDepthResponseBody) *EstimateMonocularImageDepthResponse {
	s.Body = v
	return s
}

type EstimateStereoImageDepthRequest struct {
	// A short description of struct
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
	// Id of the request
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *EstimateStereoImageDepthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s EstimateStereoImageDepthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EstimateStereoImageDepthResponseBody) GoString() string {
	return s.String()
}

func (s *EstimateStereoImageDepthResponseBody) SetRequestId(v string) *EstimateStereoImageDepthResponseBody {
	s.RequestId = &v
	return s
}

func (s *EstimateStereoImageDepthResponseBody) SetData(v *EstimateStereoImageDepthResponseBodyData) *EstimateStereoImageDepthResponseBody {
	s.Data = v
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EstimateStereoImageDepthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EstimateStereoImageDepthResponse) SetBody(v *EstimateStereoImageDepthResponseBody) *EstimateStereoImageDepthResponse {
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

func (client *Client) ReconstructBodyBySingleImageWithOptions(request *ReconstructBodyBySingleImageRequest, runtime *util.RuntimeOptions) (_result *ReconstructBodyBySingleImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReconstructBodyBySingleImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReconstructBodyBySingleImage"), tea.String("2021-01-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	reconstructBodyBySingleImageReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReconstructThreeDMultiViewResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReconstructThreeDMultiView"), tea.String("2021-01-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ZipFileUrlObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	reconstructThreeDMultiViewReq.ZipFileUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	reconstructThreeDMultiViewResp, _err := client.ReconstructThreeDMultiViewWithOptions(reconstructThreeDMultiViewReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = reconstructThreeDMultiViewResp
	return _result, _err
}

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncJobResult"), tea.String("2021-01-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) EstimateMonocularImageDepthWithOptions(request *EstimateMonocularImageDepthRequest, runtime *util.RuntimeOptions) (_result *EstimateMonocularImageDepthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EstimateMonocularImageDepthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EstimateMonocularImageDepth"), tea.String("2021-01-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

	authConfig := &rpc.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Endpoint:        tea.String("openplatform.aliyuncs.com"),
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
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
	ossClient, _err = oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj = &fileform.FileField{
		Filename:    authResponse.ObjectKey,
		Content:     request.ImageURLObject,
		ContentType: tea.String(""),
	}
	ossHeader = &oss.PostObjectRequestHeader{
		AccessKeyId:         authResponse.AccessKeyId,
		Policy:              authResponse.EncodedPolicy,
		Signature:           authResponse.Signature,
		Key:                 authResponse.ObjectKey,
		File:                fileObj,
		SuccessActionStatus: tea.String("201"),
	}
	uploadRequest = &oss.PostObjectRequest{
		BucketName: authResponse.Bucket,
		Header:     ossHeader,
	}
	_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
	if _err != nil {
		return _result, _err
	}
	estimateMonocularImageDepthReq.ImageURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	estimateMonocularImageDepthResp, _err := client.EstimateMonocularImageDepthWithOptions(estimateMonocularImageDepthReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = estimateMonocularImageDepthResp
	return _result, _err
}

func (client *Client) EstimateStereoImageDepthWithOptions(request *EstimateStereoImageDepthRequest, runtime *util.RuntimeOptions) (_result *EstimateStereoImageDepthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EstimateStereoImageDepthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EstimateStereoImageDepth"), tea.String("2021-01-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
