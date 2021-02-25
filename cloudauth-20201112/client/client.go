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

type LivenessDetectRequest struct {
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	MediaCategory *string `json:"MediaCategory,omitempty" xml:"MediaCategory,omitempty"`
	MediaUrl      *string `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
	MediaFile     *string `json:"MediaFile,omitempty" xml:"MediaFile,omitempty"`
}

func (s LivenessDetectRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectRequest) GoString() string {
	return s.String()
}

func (s *LivenessDetectRequest) SetBizType(v string) *LivenessDetectRequest {
	s.BizType = &v
	return s
}

func (s *LivenessDetectRequest) SetBizId(v string) *LivenessDetectRequest {
	s.BizId = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaCategory(v string) *LivenessDetectRequest {
	s.MediaCategory = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaUrl(v string) *LivenessDetectRequest {
	s.MediaUrl = &v
	return s
}

func (s *LivenessDetectRequest) SetMediaFile(v string) *LivenessDetectRequest {
	s.MediaFile = &v
	return s
}

type LivenessDetectAdvanceRequest struct {
	MediaFileObject io.Reader `json:"MediaFileObject,omitempty" xml:"MediaFileObject,omitempty" require:"true"`
	BizType         *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	BizId           *string   `json:"BizId,omitempty" xml:"BizId,omitempty"`
	MediaCategory   *string   `json:"MediaCategory,omitempty" xml:"MediaCategory,omitempty"`
	MediaUrl        *string   `json:"MediaUrl,omitempty" xml:"MediaUrl,omitempty"`
}

func (s LivenessDetectAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectAdvanceRequest) GoString() string {
	return s.String()
}

func (s *LivenessDetectAdvanceRequest) SetMediaFileObject(v io.Reader) *LivenessDetectAdvanceRequest {
	s.MediaFileObject = v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetBizType(v string) *LivenessDetectAdvanceRequest {
	s.BizType = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetBizId(v string) *LivenessDetectAdvanceRequest {
	s.BizId = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetMediaCategory(v string) *LivenessDetectAdvanceRequest {
	s.MediaCategory = &v
	return s
}

func (s *LivenessDetectAdvanceRequest) SetMediaUrl(v string) *LivenessDetectAdvanceRequest {
	s.MediaUrl = &v
	return s
}

type LivenessDetectResponseBody struct {
	ResultObject *LivenessDetectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s LivenessDetectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponseBody) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponseBody) SetResultObject(v *LivenessDetectResponseBodyResultObject) *LivenessDetectResponseBody {
	s.ResultObject = v
	return s
}

func (s *LivenessDetectResponseBody) SetMessage(v string) *LivenessDetectResponseBody {
	s.Message = &v
	return s
}

func (s *LivenessDetectResponseBody) SetRequestId(v string) *LivenessDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *LivenessDetectResponseBody) SetCode(v string) *LivenessDetectResponseBody {
	s.Code = &v
	return s
}

type LivenessDetectResponseBodyResultObject struct {
	Score    *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	FrameUrl *string  `json:"FrameUrl,omitempty" xml:"FrameUrl,omitempty"`
	Passed   *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
}

func (s LivenessDetectResponseBodyResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponseBodyResultObject) SetScore(v float32) *LivenessDetectResponseBodyResultObject {
	s.Score = &v
	return s
}

func (s *LivenessDetectResponseBodyResultObject) SetFrameUrl(v string) *LivenessDetectResponseBodyResultObject {
	s.FrameUrl = &v
	return s
}

func (s *LivenessDetectResponseBodyResultObject) SetPassed(v string) *LivenessDetectResponseBodyResultObject {
	s.Passed = &v
	return s
}

type LivenessDetectResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LivenessDetectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LivenessDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponse) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponse) SetHeaders(v map[string]*string) *LivenessDetectResponse {
	s.Headers = v
	return s
}

func (s *LivenessDetectResponse) SetBody(v *LivenessDetectResponseBody) *LivenessDetectResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) LivenessDetectWithOptions(request *LivenessDetectRequest, runtime *util.RuntimeOptions) (_result *LivenessDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LivenessDetectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LivenessDetect"), tea.String("2020-11-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LivenessDetect(request *LivenessDetectRequest) (_result *LivenessDetectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LivenessDetectResponse{}
	_body, _err := client.LivenessDetectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LivenessDetectAdvance(request *LivenessDetectAdvanceRequest, runtime *util.RuntimeOptions) (_result *LivenessDetectResponse, _err error) {
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
		Product:  tea.String("Cloudauth"),
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
	livenessDetectReq := &LivenessDetectRequest{}
	openapiutil.Convert(request, livenessDetectReq)
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
		Content:     request.MediaFileObject,
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
	livenessDetectReq.MediaFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	livenessDetectResp, _err := client.LivenessDetectWithOptions(livenessDetectReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = livenessDetectResp
	return _result, _err
}
