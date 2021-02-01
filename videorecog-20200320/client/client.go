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

type DetectVideoShotRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s DetectVideoShotRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoShotRequest) SetVideoUrl(v string) *DetectVideoShotRequest {
	s.VideoUrl = &v
	return s
}

func (s *DetectVideoShotRequest) SetAsync(v bool) *DetectVideoShotRequest {
	s.Async = &v
	return s
}

type DetectVideoShotAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s DetectVideoShotAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVideoShotAdvanceRequest) SetVideoUrlObject(v io.Reader) *DetectVideoShotAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *DetectVideoShotAdvanceRequest) SetAsync(v bool) *DetectVideoShotAdvanceRequest {
	s.Async = &v
	return s
}

type DetectVideoShotResponseBody struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DetectVideoShotResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DetectVideoShotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBody) SetRequestId(v string) *DetectVideoShotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVideoShotResponseBody) SetData(v *DetectVideoShotResponseBodyData) *DetectVideoShotResponseBody {
	s.Data = v
	return s
}

type DetectVideoShotResponseBodyData struct {
	ShotFrameIds []*int32 `json:"ShotFrameIds,omitempty" xml:"ShotFrameIds,omitempty" type:"Repeated"`
}

func (s DetectVideoShotResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponseBodyData) SetShotFrameIds(v []*int32) *DetectVideoShotResponseBodyData {
	s.ShotFrameIds = v
	return s
}

type DetectVideoShotResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetectVideoShotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectVideoShotResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectVideoShotResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoShotResponse) SetHeaders(v map[string]*string) *DetectVideoShotResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoShotResponse) SetBody(v *DetectVideoShotResponseBody) *DetectVideoShotResponse {
	s.Body = v
	return s
}

type GenerateVideoCoverRequest struct {
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	IsGif    *bool   `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
}

func (s GenerateVideoCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverRequest) SetVideoUrl(v string) *GenerateVideoCoverRequest {
	s.VideoUrl = &v
	return s
}

func (s *GenerateVideoCoverRequest) SetAsync(v bool) *GenerateVideoCoverRequest {
	s.Async = &v
	return s
}

func (s *GenerateVideoCoverRequest) SetIsGif(v bool) *GenerateVideoCoverRequest {
	s.IsGif = &v
	return s
}

type GenerateVideoCoverAdvanceRequest struct {
	VideoUrlObject io.Reader `json:"VideoUrlObject,omitempty" xml:"VideoUrlObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
	IsGif          *bool     `json:"IsGif,omitempty" xml:"IsGif,omitempty"`
}

func (s GenerateVideoCoverAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverAdvanceRequest) SetVideoUrlObject(v io.Reader) *GenerateVideoCoverAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *GenerateVideoCoverAdvanceRequest) SetAsync(v bool) *GenerateVideoCoverAdvanceRequest {
	s.Async = &v
	return s
}

func (s *GenerateVideoCoverAdvanceRequest) SetIsGif(v bool) *GenerateVideoCoverAdvanceRequest {
	s.IsGif = &v
	return s
}

type GenerateVideoCoverResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *GenerateVideoCoverResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s GenerateVideoCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBody) SetRequestId(v string) *GenerateVideoCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateVideoCoverResponseBody) SetData(v *GenerateVideoCoverResponseBodyData) *GenerateVideoCoverResponseBody {
	s.Data = v
	return s
}

type GenerateVideoCoverResponseBodyData struct {
	Outputs []*GenerateVideoCoverResponseBodyDataOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s GenerateVideoCoverResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyData) SetOutputs(v []*GenerateVideoCoverResponseBodyDataOutputs) *GenerateVideoCoverResponseBodyData {
	s.Outputs = v
	return s
}

type GenerateVideoCoverResponseBodyDataOutputs struct {
	ImageURL   *string  `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s GenerateVideoCoverResponseBodyDataOutputs) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetImageURL(v string) *GenerateVideoCoverResponseBodyDataOutputs {
	s.ImageURL = &v
	return s
}

func (s *GenerateVideoCoverResponseBodyDataOutputs) SetConfidence(v float32) *GenerateVideoCoverResponseBodyDataOutputs {
	s.Confidence = &v
	return s
}

type GenerateVideoCoverResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GenerateVideoCoverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateVideoCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateVideoCoverResponse) GoString() string {
	return s.String()
}

func (s *GenerateVideoCoverResponse) SetHeaders(v map[string]*string) *GenerateVideoCoverResponse {
	s.Headers = v
	return s
}

func (s *GenerateVideoCoverResponse) SetBody(v *GenerateVideoCoverResponseBody) *GenerateVideoCoverResponse {
	s.Body = v
	return s
}

type UnderstandVideoContentRequest struct {
	// A short description of struct
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
	Async    *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s UnderstandVideoContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentRequest) SetVideoURL(v string) *UnderstandVideoContentRequest {
	s.VideoURL = &v
	return s
}

func (s *UnderstandVideoContentRequest) SetAsync(v bool) *UnderstandVideoContentRequest {
	s.Async = &v
	return s
}

type UnderstandVideoContentAdvanceRequest struct {
	VideoURLObject io.Reader `json:"VideoURLObject,omitempty" xml:"VideoURLObject,omitempty" require:"true"`
	Async          *bool     `json:"Async,omitempty" xml:"Async,omitempty"`
}

func (s UnderstandVideoContentAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentAdvanceRequest) SetVideoURLObject(v io.Reader) *UnderstandVideoContentAdvanceRequest {
	s.VideoURLObject = v
	return s
}

func (s *UnderstandVideoContentAdvanceRequest) SetAsync(v bool) *UnderstandVideoContentAdvanceRequest {
	s.Async = &v
	return s
}

type UnderstandVideoContentResponseBody struct {
	// Id of the request
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *UnderstandVideoContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s UnderstandVideoContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBody) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBody) SetRequestId(v string) *UnderstandVideoContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnderstandVideoContentResponseBody) SetData(v *UnderstandVideoContentResponseBodyData) *UnderstandVideoContentResponseBody {
	s.Data = v
	return s
}

type UnderstandVideoContentResponseBodyData struct {
	TagInfo   map[string]interface{}                           `json:"TagInfo,omitempty" xml:"TagInfo,omitempty"`
	VideoInfo *UnderstandVideoContentResponseBodyDataVideoInfo `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s UnderstandVideoContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyData) SetTagInfo(v map[string]interface{}) *UnderstandVideoContentResponseBodyData {
	s.TagInfo = v
	return s
}

func (s *UnderstandVideoContentResponseBodyData) SetVideoInfo(v *UnderstandVideoContentResponseBodyDataVideoInfo) *UnderstandVideoContentResponseBodyData {
	s.VideoInfo = v
	return s
}

type UnderstandVideoContentResponseBodyDataVideoInfo struct {
	Width    *int64   `json:"Width,omitempty" xml:"Width,omitempty"`
	Height   *int64   `json:"Height,omitempty" xml:"Height,omitempty"`
	Duration *int64   `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Fps      *float32 `json:"Fps,omitempty" xml:"Fps,omitempty"`
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponseBodyDataVideoInfo) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetWidth(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Width = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetHeight(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Height = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetDuration(v int64) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Duration = &v
	return s
}

func (s *UnderstandVideoContentResponseBodyDataVideoInfo) SetFps(v float32) *UnderstandVideoContentResponseBodyDataVideoInfo {
	s.Fps = &v
	return s
}

type UnderstandVideoContentResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnderstandVideoContentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnderstandVideoContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UnderstandVideoContentResponse) GoString() string {
	return s.String()
}

func (s *UnderstandVideoContentResponse) SetHeaders(v map[string]*string) *UnderstandVideoContentResponse {
	s.Headers = v
	return s
}

func (s *UnderstandVideoContentResponse) SetBody(v *UnderstandVideoContentResponseBody) *UnderstandVideoContentResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("videorecog"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DetectVideoShotWithOptions(request *DetectVideoShotRequest, runtime *util.RuntimeOptions) (_result *DetectVideoShotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetectVideoShotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetectVideoShot"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectVideoShot(request *DetectVideoShotRequest) (_result *DetectVideoShotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectVideoShotResponse{}
	_body, _err := client.DetectVideoShotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectVideoShotAdvance(request *DetectVideoShotAdvanceRequest, runtime *util.RuntimeOptions) (_result *DetectVideoShotResponse, _err error) {
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
		Product:  tea.String("videorecog"),
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
	detectVideoShotReq := &DetectVideoShotRequest{}
	openapiutil.Convert(request, detectVideoShotReq)
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
	detectVideoShotReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	detectVideoShotResp, _err := client.DetectVideoShotWithOptions(detectVideoShotReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = detectVideoShotResp
	return _result, _err
}

func (client *Client) GenerateVideoCoverWithOptions(request *GenerateVideoCoverRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GenerateVideoCoverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GenerateVideoCover"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateVideoCover(request *GenerateVideoCoverRequest) (_result *GenerateVideoCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateVideoCoverResponse{}
	_body, _err := client.GenerateVideoCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateVideoCoverAdvance(request *GenerateVideoCoverAdvanceRequest, runtime *util.RuntimeOptions) (_result *GenerateVideoCoverResponse, _err error) {
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
		Product:  tea.String("videorecog"),
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
	generateVideoCoverReq := &GenerateVideoCoverRequest{}
	openapiutil.Convert(request, generateVideoCoverReq)
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
	generateVideoCoverReq.VideoUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	generateVideoCoverResp, _err := client.GenerateVideoCoverWithOptions(generateVideoCoverReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generateVideoCoverResp
	return _result, _err
}

func (client *Client) UnderstandVideoContentWithOptions(request *UnderstandVideoContentRequest, runtime *util.RuntimeOptions) (_result *UnderstandVideoContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnderstandVideoContentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnderstandVideoContent"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnderstandVideoContent(request *UnderstandVideoContentRequest) (_result *UnderstandVideoContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnderstandVideoContentResponse{}
	_body, _err := client.UnderstandVideoContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnderstandVideoContentAdvance(request *UnderstandVideoContentAdvanceRequest, runtime *util.RuntimeOptions) (_result *UnderstandVideoContentResponse, _err error) {
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
		Product:  tea.String("videorecog"),
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
	understandVideoContentReq := &UnderstandVideoContentRequest{}
	openapiutil.Convert(request, understandVideoContentReq)
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
		Content:     request.VideoURLObject,
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
	understandVideoContentReq.VideoURL = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	understandVideoContentResp, _err := client.UnderstandVideoContentWithOptions(understandVideoContentReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = understandVideoContentResp
	return _result, _err
}
