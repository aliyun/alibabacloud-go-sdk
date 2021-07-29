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

type GeneralRecognitionRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicContent   *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
}

func (s GeneralRecognitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionRequest) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionRequest) SetInstanceName(v string) *GeneralRecognitionRequest {
	s.InstanceName = &v
	return s
}

func (s *GeneralRecognitionRequest) SetPicContent(v string) *GeneralRecognitionRequest {
	s.PicContent = &v
	return s
}

type GeneralRecognitionAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GeneralRecognitionAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionAdvanceRequest) SetPicContentObject(v io.Reader) *GeneralRecognitionAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *GeneralRecognitionAdvanceRequest) SetInstanceName(v string) *GeneralRecognitionAdvanceRequest {
	s.InstanceName = &v
	return s
}

type GeneralRecognitionResponseBody struct {
	Code      *int32                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *GeneralRecognitionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GeneralRecognitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionResponseBody) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionResponseBody) SetCode(v int32) *GeneralRecognitionResponseBody {
	s.Code = &v
	return s
}

func (s *GeneralRecognitionResponseBody) SetMessage(v string) *GeneralRecognitionResponseBody {
	s.Message = &v
	return s
}

func (s *GeneralRecognitionResponseBody) SetData(v *GeneralRecognitionResponseBodyData) *GeneralRecognitionResponseBody {
	s.Data = v
	return s
}

func (s *GeneralRecognitionResponseBody) SetSuccess(v bool) *GeneralRecognitionResponseBody {
	s.Success = &v
	return s
}

func (s *GeneralRecognitionResponseBody) SetRequestId(v string) *GeneralRecognitionResponseBody {
	s.RequestId = &v
	return s
}

type GeneralRecognitionResponseBodyData struct {
	Result  []*GeneralRecognitionResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	Regions []*string                                   `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s GeneralRecognitionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionResponseBodyData) SetResult(v []*GeneralRecognitionResponseBodyDataResult) *GeneralRecognitionResponseBodyData {
	s.Result = v
	return s
}

func (s *GeneralRecognitionResponseBodyData) SetRegions(v []*string) *GeneralRecognitionResponseBodyData {
	s.Regions = v
	return s
}

type GeneralRecognitionResponseBodyDataResult struct {
	Tag   *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Score *string `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GeneralRecognitionResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionResponseBodyDataResult) SetTag(v string) *GeneralRecognitionResponseBodyDataResult {
	s.Tag = &v
	return s
}

func (s *GeneralRecognitionResponseBodyDataResult) SetScore(v string) *GeneralRecognitionResponseBodyDataResult {
	s.Score = &v
	return s
}

type GeneralRecognitionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GeneralRecognitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GeneralRecognitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GeneralRecognitionResponse) GoString() string {
	return s.String()
}

func (s *GeneralRecognitionResponse) SetHeaders(v map[string]*string) *GeneralRecognitionResponse {
	s.Headers = v
	return s
}

func (s *GeneralRecognitionResponse) SetBody(v *GeneralRecognitionResponseBody) *GeneralRecognitionResponse {
	s.Body = v
	return s
}

type ImageDuplicationRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicUrlList   *string `json:"PicUrlList,omitempty" xml:"PicUrlList,omitempty"`
	PicNumList   *string `json:"PicNumList,omitempty" xml:"PicNumList,omitempty"`
	ImageHeight  *int64  `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageWidth   *int64  `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
}

func (s ImageDuplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageDuplicationRequest) GoString() string {
	return s.String()
}

func (s *ImageDuplicationRequest) SetInstanceName(v string) *ImageDuplicationRequest {
	s.InstanceName = &v
	return s
}

func (s *ImageDuplicationRequest) SetPicUrlList(v string) *ImageDuplicationRequest {
	s.PicUrlList = &v
	return s
}

func (s *ImageDuplicationRequest) SetPicNumList(v string) *ImageDuplicationRequest {
	s.PicNumList = &v
	return s
}

func (s *ImageDuplicationRequest) SetImageHeight(v int64) *ImageDuplicationRequest {
	s.ImageHeight = &v
	return s
}

func (s *ImageDuplicationRequest) SetImageWidth(v int64) *ImageDuplicationRequest {
	s.ImageWidth = &v
	return s
}

type ImageDuplicationResponseBody struct {
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageDuplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageDuplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageDuplicationResponseBody) SetCode(v int32) *ImageDuplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageDuplicationResponseBody) SetMessage(v string) *ImageDuplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ImageDuplicationResponseBody) SetData(v []*string) *ImageDuplicationResponseBody {
	s.Data = v
	return s
}

func (s *ImageDuplicationResponseBody) SetSuccess(v bool) *ImageDuplicationResponseBody {
	s.Success = &v
	return s
}

func (s *ImageDuplicationResponseBody) SetRequestId(v string) *ImageDuplicationResponseBody {
	s.RequestId = &v
	return s
}

type ImageDuplicationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImageDuplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImageDuplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageDuplicationResponse) GoString() string {
	return s.String()
}

func (s *ImageDuplicationResponse) SetHeaders(v map[string]*string) *ImageDuplicationResponse {
	s.Headers = v
	return s
}

func (s *ImageDuplicationResponse) SetBody(v *ImageDuplicationResponseBody) *ImageDuplicationResponse {
	s.Body = v
	return s
}

type ImageSegmentationRequest struct {
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicContent      *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	ReturnPicFormat *string `json:"ReturnPicFormat,omitempty" xml:"ReturnPicFormat,omitempty"`
	ReturnPicType   *string `json:"ReturnPicType,omitempty" xml:"ReturnPicType,omitempty"`
}

func (s ImageSegmentationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageSegmentationRequest) GoString() string {
	return s.String()
}

func (s *ImageSegmentationRequest) SetInstanceName(v string) *ImageSegmentationRequest {
	s.InstanceName = &v
	return s
}

func (s *ImageSegmentationRequest) SetPicContent(v string) *ImageSegmentationRequest {
	s.PicContent = &v
	return s
}

func (s *ImageSegmentationRequest) SetReturnPicFormat(v string) *ImageSegmentationRequest {
	s.ReturnPicFormat = &v
	return s
}

func (s *ImageSegmentationRequest) SetReturnPicType(v string) *ImageSegmentationRequest {
	s.ReturnPicType = &v
	return s
}

type ImageSegmentationAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ReturnPicFormat  *string   `json:"ReturnPicFormat,omitempty" xml:"ReturnPicFormat,omitempty"`
	ReturnPicType    *string   `json:"ReturnPicType,omitempty" xml:"ReturnPicType,omitempty"`
}

func (s ImageSegmentationAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageSegmentationAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ImageSegmentationAdvanceRequest) SetPicContentObject(v io.Reader) *ImageSegmentationAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *ImageSegmentationAdvanceRequest) SetInstanceName(v string) *ImageSegmentationAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ImageSegmentationAdvanceRequest) SetReturnPicFormat(v string) *ImageSegmentationAdvanceRequest {
	s.ReturnPicFormat = &v
	return s
}

func (s *ImageSegmentationAdvanceRequest) SetReturnPicType(v string) *ImageSegmentationAdvanceRequest {
	s.ReturnPicType = &v
	return s
}

type ImageSegmentationResponseBody struct {
	Code      *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *ImageSegmentationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageSegmentationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageSegmentationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageSegmentationResponseBody) SetCode(v int32) *ImageSegmentationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageSegmentationResponseBody) SetMessage(v string) *ImageSegmentationResponseBody {
	s.Message = &v
	return s
}

func (s *ImageSegmentationResponseBody) SetData(v *ImageSegmentationResponseBodyData) *ImageSegmentationResponseBody {
	s.Data = v
	return s
}

func (s *ImageSegmentationResponseBody) SetSuccess(v bool) *ImageSegmentationResponseBody {
	s.Success = &v
	return s
}

func (s *ImageSegmentationResponseBody) SetRequestId(v string) *ImageSegmentationResponseBody {
	s.RequestId = &v
	return s
}

type ImageSegmentationResponseBodyData struct {
	PicUrl *string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
}

func (s ImageSegmentationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImageSegmentationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageSegmentationResponseBodyData) SetPicUrl(v string) *ImageSegmentationResponseBodyData {
	s.PicUrl = &v
	return s
}

type ImageSegmentationResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImageSegmentationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImageSegmentationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageSegmentationResponse) GoString() string {
	return s.String()
}

func (s *ImageSegmentationResponse) SetHeaders(v map[string]*string) *ImageSegmentationResponse {
	s.Headers = v
	return s
}

func (s *ImageSegmentationResponse) SetBody(v *ImageSegmentationResponseBody) *ImageSegmentationResponse {
	s.Body = v
	return s
}

type CommodityTitleRequest struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PicContent   *string `json:"PicContent,omitempty" xml:"PicContent,omitempty"`
	TitleIndex   *int64  `json:"TitleIndex,omitempty" xml:"TitleIndex,omitempty"`
}

func (s CommodityTitleRequest) String() string {
	return tea.Prettify(s)
}

func (s CommodityTitleRequest) GoString() string {
	return s.String()
}

func (s *CommodityTitleRequest) SetInstanceName(v string) *CommodityTitleRequest {
	s.InstanceName = &v
	return s
}

func (s *CommodityTitleRequest) SetPicContent(v string) *CommodityTitleRequest {
	s.PicContent = &v
	return s
}

func (s *CommodityTitleRequest) SetTitleIndex(v int64) *CommodityTitleRequest {
	s.TitleIndex = &v
	return s
}

type CommodityTitleAdvanceRequest struct {
	PicContentObject io.Reader `json:"PicContentObject,omitempty" xml:"PicContentObject,omitempty" require:"true"`
	InstanceName     *string   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	TitleIndex       *int64    `json:"TitleIndex,omitempty" xml:"TitleIndex,omitempty"`
}

func (s CommodityTitleAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CommodityTitleAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CommodityTitleAdvanceRequest) SetPicContentObject(v io.Reader) *CommodityTitleAdvanceRequest {
	s.PicContentObject = v
	return s
}

func (s *CommodityTitleAdvanceRequest) SetInstanceName(v string) *CommodityTitleAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CommodityTitleAdvanceRequest) SetTitleIndex(v int64) *CommodityTitleAdvanceRequest {
	s.TitleIndex = &v
	return s
}

type CommodityTitleResponseBody struct {
	Code      *int32                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *CommodityTitleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommodityTitleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommodityTitleResponseBody) GoString() string {
	return s.String()
}

func (s *CommodityTitleResponseBody) SetCode(v int32) *CommodityTitleResponseBody {
	s.Code = &v
	return s
}

func (s *CommodityTitleResponseBody) SetMessage(v string) *CommodityTitleResponseBody {
	s.Message = &v
	return s
}

func (s *CommodityTitleResponseBody) SetData(v *CommodityTitleResponseBodyData) *CommodityTitleResponseBody {
	s.Data = v
	return s
}

func (s *CommodityTitleResponseBody) SetSuccess(v bool) *CommodityTitleResponseBody {
	s.Success = &v
	return s
}

func (s *CommodityTitleResponseBody) SetRequestId(v string) *CommodityTitleResponseBody {
	s.RequestId = &v
	return s
}

type CommodityTitleResponseBodyData struct {
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CommodityTitleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CommodityTitleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CommodityTitleResponseBodyData) SetTitle(v string) *CommodityTitleResponseBodyData {
	s.Title = &v
	return s
}

type CommodityTitleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CommodityTitleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommodityTitleResponse) String() string {
	return tea.Prettify(s)
}

func (s CommodityTitleResponse) GoString() string {
	return s.String()
}

func (s *CommodityTitleResponse) SetHeaders(v map[string]*string) *CommodityTitleResponse {
	s.Headers = v
	return s
}

func (s *CommodityTitleResponse) SetBody(v *CommodityTitleResponseBody) *CommodityTitleResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imagesearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GeneralRecognitionWithOptions(request *GeneralRecognitionRequest, runtime *util.RuntimeOptions) (_result *GeneralRecognitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GeneralRecognitionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GeneralRecognition"), tea.String("2021-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GeneralRecognition(request *GeneralRecognitionRequest) (_result *GeneralRecognitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GeneralRecognitionResponse{}
	_body, _err := client.GeneralRecognitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GeneralRecognitionAdvance(request *GeneralRecognitionAdvanceRequest, runtime *util.RuntimeOptions) (_result *GeneralRecognitionResponse, _err error) {
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

	authConfig := &rpc.Config{
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
		Product:  tea.String("ImageSearch"),
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
	generalRecognitionReq := &GeneralRecognitionRequest{}
	openapiutil.Convert(request, generalRecognitionReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
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
			Content:     request.PicContentObject,
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
		generalRecognitionReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	generalRecognitionResp, _err := client.GeneralRecognitionWithOptions(generalRecognitionReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = generalRecognitionResp
	return _result, _err
}

func (client *Client) ImageDuplicationWithOptions(request *ImageDuplicationRequest, runtime *util.RuntimeOptions) (_result *ImageDuplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImageDuplicationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImageDuplication"), tea.String("2021-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageDuplication(request *ImageDuplicationRequest) (_result *ImageDuplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageDuplicationResponse{}
	_body, _err := client.ImageDuplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageSegmentationWithOptions(request *ImageSegmentationRequest, runtime *util.RuntimeOptions) (_result *ImageSegmentationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ImageSegmentationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ImageSegmentation"), tea.String("2021-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageSegmentation(request *ImageSegmentationRequest) (_result *ImageSegmentationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImageSegmentationResponse{}
	_body, _err := client.ImageSegmentationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageSegmentationAdvance(request *ImageSegmentationAdvanceRequest, runtime *util.RuntimeOptions) (_result *ImageSegmentationResponse, _err error) {
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

	authConfig := &rpc.Config{
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
		Product:  tea.String("ImageSearch"),
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
	imageSegmentationReq := &ImageSegmentationRequest{}
	openapiutil.Convert(request, imageSegmentationReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
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
			Content:     request.PicContentObject,
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
		imageSegmentationReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	imageSegmentationResp, _err := client.ImageSegmentationWithOptions(imageSegmentationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = imageSegmentationResp
	return _result, _err
}

func (client *Client) CommodityTitleWithOptions(request *CommodityTitleRequest, runtime *util.RuntimeOptions) (_result *CommodityTitleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CommodityTitleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CommodityTitle"), tea.String("2021-01-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommodityTitle(request *CommodityTitleRequest) (_result *CommodityTitleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommodityTitleResponse{}
	_body, _err := client.CommodityTitleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CommodityTitleAdvance(request *CommodityTitleAdvanceRequest, runtime *util.RuntimeOptions) (_result *CommodityTitleResponse, _err error) {
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

	authConfig := &rpc.Config{
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
		Product:  tea.String("ImageSearch"),
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
	commodityTitleReq := &CommodityTitleRequest{}
	openapiutil.Convert(request, commodityTitleReq)
	if !tea.BoolValue(util.IsUnset(request.PicContentObject)) {
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
			Content:     request.PicContentObject,
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
		commodityTitleReq.PicContent = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	}

	commodityTitleResp, _err := client.CommodityTitleWithOptions(commodityTitleReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = commodityTitleResp
	return _result, _err
}
