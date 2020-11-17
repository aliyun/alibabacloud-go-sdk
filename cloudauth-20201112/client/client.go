// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
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

type LivenessDetectResponse struct {
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	ResultObject *LivenessDetectResponseResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" require:"true" type:"Struct"`
}

func (s LivenessDetectResponse) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponse) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponse) SetRequestId(v string) *LivenessDetectResponse {
	s.RequestId = &v
	return s
}

func (s *LivenessDetectResponse) SetCode(v string) *LivenessDetectResponse {
	s.Code = &v
	return s
}

func (s *LivenessDetectResponse) SetMessage(v string) *LivenessDetectResponse {
	s.Message = &v
	return s
}

func (s *LivenessDetectResponse) SetResultObject(v *LivenessDetectResponseResultObject) *LivenessDetectResponse {
	s.ResultObject = v
	return s
}

type LivenessDetectResponseResultObject struct {
	Passed   *string  `json:"Passed,omitempty" xml:"Passed,omitempty" require:"true"`
	Score    *float32 `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	FrameUrl *string  `json:"FrameUrl,omitempty" xml:"FrameUrl,omitempty" require:"true"`
}

func (s LivenessDetectResponseResultObject) String() string {
	return tea.Prettify(s)
}

func (s LivenessDetectResponseResultObject) GoString() string {
	return s.String()
}

func (s *LivenessDetectResponseResultObject) SetPassed(v string) *LivenessDetectResponseResultObject {
	s.Passed = &v
	return s
}

func (s *LivenessDetectResponseResultObject) SetScore(v float32) *LivenessDetectResponseResultObject {
	s.Score = &v
	return s
}

func (s *LivenessDetectResponseResultObject) SetFrameUrl(v string) *LivenessDetectResponseResultObject {
	s.FrameUrl = &v
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

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) LivenessDetect(request *LivenessDetectRequest, runtime *util.RuntimeOptions) (_result *LivenessDetectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &LivenessDetectResponse{}
	_body, _err := client.DoRequest(tea.String("LivenessDetect"), tea.String("HTTPS"), tea.String("POST"), tea.String("2020-11-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	rpcutil.Convert(runtime, ossRuntime)
	livenessDetectreq := &LivenessDetectRequest{}
	rpcutil.Convert(request, livenessDetectreq)
	authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
	if _err != nil {
		return _result, _err
	}

	ossConfig.AccessKeyId = authResponse.AccessKeyId
	ossConfig.Endpoint = rpcutil.GetEndpoint(authResponse.Endpoint, authResponse.UseAccelerate, client.EndpointType)
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
	livenessDetectreq.MediaFile = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	livenessDetectResp, _err := client.LivenessDetect(livenessDetectreq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = livenessDetectResp
	return _result, _err
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
