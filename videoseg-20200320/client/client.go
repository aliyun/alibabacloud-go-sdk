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

type GetAsyncJobResultRequest struct {
	Async *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAsyncJobResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncJobResultRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncJobResultRequest) SetAsync(v bool) *GetAsyncJobResultRequest {
	s.Async = &v
	return s
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

type SegmentHalfBodyRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s SegmentHalfBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHalfBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentHalfBodyRequest) SetVideoUrl(v string) *SegmentHalfBodyRequest {
	s.VideoUrl = &v
	return s
}

func (s *SegmentHalfBodyRequest) SetAsync(v bool) *SegmentHalfBodyRequest {
	s.Async = &v
	return s
}

type SegmentHalfBodyAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s SegmentHalfBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentHalfBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentHalfBodyAdvanceRequest) SetVideoUrlObject(v io.Reader) *SegmentHalfBodyAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SegmentHalfBodyAdvanceRequest) SetAsync(v bool) *SegmentHalfBodyAdvanceRequest {
	s.Async = &v
	return s
}

type SegmentHalfBodyResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SegmentHalfBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SegmentHalfBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentHalfBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentHalfBodyResponseBody) SetRequestId(v string) *SegmentHalfBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentHalfBodyResponseBody) SetData(v *SegmentHalfBodyResponseBodyData) *SegmentHalfBodyResponseBody {
	s.Data = v
	return s
}

type SegmentHalfBodyResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SegmentHalfBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentHalfBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentHalfBodyResponseBodyData) SetVideoUrl(v string) *SegmentHalfBodyResponseBodyData {
	s.VideoUrl = &v
	return s
}

type SegmentHalfBodyResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SegmentHalfBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentHalfBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentHalfBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHalfBodyResponse) SetHeaders(v map[string]*string) *SegmentHalfBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentHalfBodyResponse) SetBody(v *SegmentHalfBodyResponseBody) *SegmentHalfBodyResponse {
	s.Body = v
	return s
}

type SegmentVideoBodyRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s SegmentVideoBodyRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVideoBodyRequest) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyRequest) SetVideoUrl(v string) *SegmentVideoBodyRequest {
	s.VideoUrl = &v
	return s
}

func (s *SegmentVideoBodyRequest) SetAsync(v bool) *SegmentVideoBodyRequest {
	s.Async = &v
	return s
}

type SegmentVideoBodyAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s SegmentVideoBodyAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SegmentVideoBodyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyAdvanceRequest) SetVideoUrlObject(v io.Reader) *SegmentVideoBodyAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SegmentVideoBodyAdvanceRequest) SetAsync(v bool) *SegmentVideoBodyAdvanceRequest {
	s.Async = &v
	return s
}

type SegmentVideoBodyResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SegmentVideoBodyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SegmentVideoBodyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SegmentVideoBodyResponseBody) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponseBody) SetRequestId(v string) *SegmentVideoBodyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SegmentVideoBodyResponseBody) SetData(v *SegmentVideoBodyResponseBodyData) *SegmentVideoBodyResponseBody {
	s.Data = v
	return s
}

type SegmentVideoBodyResponseBodyData struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SegmentVideoBodyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SegmentVideoBodyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponseBodyData) SetVideoUrl(v string) *SegmentVideoBodyResponseBodyData {
	s.VideoUrl = &v
	return s
}

type SegmentVideoBodyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SegmentVideoBodyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SegmentVideoBodyResponse) String() string {
	return tea.Prettify(s)
}

func (s SegmentVideoBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponse) SetHeaders(v map[string]*string) *SegmentVideoBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentVideoBodyResponse) SetBody(v *SegmentVideoBodyResponseBody) *SegmentVideoBodyResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("videoseg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetAsyncJobResultWithOptions(request *GetAsyncJobResultRequest, runtime *util.RuntimeOptions) (_result *GetAsyncJobResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAsyncJobResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAsyncJobResult"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SegmentHalfBodyWithOptions(request *SegmentHalfBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentHalfBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SegmentHalfBodyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SegmentHalfBody"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentHalfBody(request *SegmentHalfBodyRequest) (_result *SegmentHalfBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentHalfBodyResponse{}
	_body, _err := client.SegmentHalfBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentHalfBodyAdvance(request *SegmentHalfBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentHalfBodyResponse, _err error) {
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
		Product:  tea.String("videoseg"),
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
	segmentHalfBodyReq := &SegmentHalfBodyRequest{}
	openapiutil.Convert(request, segmentHalfBodyReq)
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
		Content:     request.VideoUrlObject,
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
	segmentHalfBodyReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentHalfBodyResp, _err := client.SegmentHalfBodyWithOptions(segmentHalfBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentHalfBodyResp
	return _result, _err
}

func (client *Client) SegmentVideoBodyWithOptions(request *SegmentVideoBodyRequest, runtime *util.RuntimeOptions) (_result *SegmentVideoBodyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SegmentVideoBodyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SegmentVideoBody"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SegmentVideoBody(request *SegmentVideoBodyRequest) (_result *SegmentVideoBodyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SegmentVideoBodyResponse{}
	_body, _err := client.SegmentVideoBodyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SegmentVideoBodyAdvance(request *SegmentVideoBodyAdvanceRequest, runtime *util.RuntimeOptions) (_result *SegmentVideoBodyResponse, _err error) {
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
		Product:  tea.String("videoseg"),
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
	segmentVideoBodyReq := &SegmentVideoBodyRequest{}
	openapiutil.Convert(request, segmentVideoBodyReq)
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
		Content:     request.VideoUrlObject,
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
	segmentVideoBodyReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	segmentVideoBodyResp, _err := client.SegmentVideoBodyWithOptions(segmentVideoBodyReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = segmentVideoBodyResp
	return _result, _err
}
