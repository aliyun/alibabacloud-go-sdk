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

type AddImageRequest struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ImageUrl  *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s AddImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageRequest) GoString() string {
	return s.String()
}

func (s *AddImageRequest) SetDbName(v string) *AddImageRequest {
	s.DbName = &v
	return s
}

func (s *AddImageRequest) SetImageUrl(v string) *AddImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *AddImageRequest) SetExtraData(v string) *AddImageRequest {
	s.ExtraData = &v
	return s
}

func (s *AddImageRequest) SetEntityId(v string) *AddImageRequest {
	s.EntityId = &v
	return s
}

type AddImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ExtraData      *string   `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	EntityId       *string   `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s AddImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s AddImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddImageAdvanceRequest) SetImageUrlObject(v io.Reader) *AddImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *AddImageAdvanceRequest) SetDbName(v string) *AddImageAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *AddImageAdvanceRequest) SetExtraData(v string) *AddImageAdvanceRequest {
	s.ExtraData = &v
	return s
}

func (s *AddImageAdvanceRequest) SetEntityId(v string) *AddImageAdvanceRequest {
	s.EntityId = &v
	return s
}

type AddImageResponseBody struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AddImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s AddImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddImageResponseBody) SetRequestId(v string) *AddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddImageResponseBody) SetData(v *AddImageResponseBodyData) *AddImageResponseBody {
	s.Data = v
	return s
}

type AddImageResponseBodyData struct {
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
}

func (s AddImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddImageResponseBodyData) SetDataId(v string) *AddImageResponseBodyData {
	s.DataId = &v
	return s
}

type AddImageResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddImageResponse) GoString() string {
	return s.String()
}

func (s *AddImageResponse) SetHeaders(v map[string]*string) *AddImageResponse {
	s.Headers = v
	return s
}

func (s *AddImageResponse) SetBody(v *AddImageResponseBody) *AddImageResponse {
	s.Body = v
	return s
}

type CreateImageDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateImageDbRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageDbRequest) GoString() string {
	return s.String()
}

func (s *CreateImageDbRequest) SetName(v string) *CreateImageDbRequest {
	s.Name = &v
	return s
}

type CreateImageDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageDbResponseBody) SetRequestId(v string) *CreateImageDbResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageDbResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateImageDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateImageDbResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageDbResponse) GoString() string {
	return s.String()
}

func (s *CreateImageDbResponse) SetHeaders(v map[string]*string) *CreateImageDbResponse {
	s.Headers = v
	return s
}

func (s *CreateImageDbResponse) SetBody(v *CreateImageDbResponseBody) *CreateImageDbResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetDbName(v string) *DeleteImageRequest {
	s.DbName = &v
	return s
}

func (s *DeleteImageRequest) SetEntityId(v string) *DeleteImageRequest {
	s.EntityId = &v
	return s
}

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteImageDbRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteImageDbRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageDbRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageDbRequest) SetName(v string) *DeleteImageDbRequest {
	s.Name = &v
	return s
}

type DeleteImageDbResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageDbResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageDbResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageDbResponseBody) SetRequestId(v string) *DeleteImageDbResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageDbResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageDbResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageDbResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageDbResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageDbResponse) SetHeaders(v map[string]*string) *DeleteImageDbResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageDbResponse) SetBody(v *DeleteImageDbResponseBody) *DeleteImageDbResponse {
	s.Body = v
	return s
}

type ListImageDbsResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListImageDbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListImageDbsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImageDbsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageDbsResponseBody) SetRequestId(v string) *ListImageDbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageDbsResponseBody) SetData(v *ListImageDbsResponseBodyData) *ListImageDbsResponseBody {
	s.Data = v
	return s
}

type ListImageDbsResponseBodyData struct {
	DbList []*ListImageDbsResponseBodyDataDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
}

func (s ListImageDbsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListImageDbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImageDbsResponseBodyData) SetDbList(v []*ListImageDbsResponseBodyDataDbList) *ListImageDbsResponseBodyData {
	s.DbList = v
	return s
}

type ListImageDbsResponseBodyDataDbList struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListImageDbsResponseBodyDataDbList) String() string {
	return tea.Prettify(s)
}

func (s ListImageDbsResponseBodyDataDbList) GoString() string {
	return s.String()
}

func (s *ListImageDbsResponseBodyDataDbList) SetName(v string) *ListImageDbsResponseBodyDataDbList {
	s.Name = &v
	return s
}

type ListImageDbsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImageDbsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImageDbsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImageDbsResponse) GoString() string {
	return s.String()
}

func (s *ListImageDbsResponse) SetHeaders(v map[string]*string) *ListImageDbsResponse {
	s.Headers = v
	return s
}

func (s *ListImageDbsResponse) SetBody(v *ListImageDbsResponseBody) *ListImageDbsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	DbName         *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Limit          *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	Order          *string `json:"Order,omitempty" xml:"Order,omitempty"`
	EntityIdPrefix *string `json:"EntityIdPrefix,omitempty" xml:"EntityIdPrefix,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetDbName(v string) *ListImagesRequest {
	s.DbName = &v
	return s
}

func (s *ListImagesRequest) SetToken(v string) *ListImagesRequest {
	s.Token = &v
	return s
}

func (s *ListImagesRequest) SetOffset(v int32) *ListImagesRequest {
	s.Offset = &v
	return s
}

func (s *ListImagesRequest) SetLimit(v int32) *ListImagesRequest {
	s.Limit = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetEntityIdPrefix(v string) *ListImagesRequest {
	s.EntityIdPrefix = &v
	return s
}

type ListImagesResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ListImagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetData(v *ListImagesResponseBodyData) *ListImagesResponseBody {
	s.Data = v
	return s
}

type ListImagesResponseBodyData struct {
	ImageList  []*ListImagesResponseBodyDataImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	Token      *string                                `json:"Token,omitempty" xml:"Token,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyData) SetImageList(v []*ListImagesResponseBodyDataImageList) *ListImagesResponseBodyData {
	s.ImageList = v
	return s
}

func (s *ListImagesResponseBodyData) SetToken(v string) *ListImagesResponseBodyData {
	s.Token = &v
	return s
}

func (s *ListImagesResponseBodyData) SetTotalCount(v int32) *ListImagesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListImagesResponseBodyDataImageList struct {
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	DataId    *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s ListImagesResponseBodyDataImageList) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyDataImageList) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyDataImageList) SetEntityId(v string) *ListImagesResponseBodyDataImageList {
	s.EntityId = &v
	return s
}

func (s *ListImagesResponseBodyDataImageList) SetCreatedAt(v int64) *ListImagesResponseBodyDataImageList {
	s.CreatedAt = &v
	return s
}

func (s *ListImagesResponseBodyDataImageList) SetUpdatedAt(v int64) *ListImagesResponseBodyDataImageList {
	s.UpdatedAt = &v
	return s
}

func (s *ListImagesResponseBodyDataImageList) SetDataId(v string) *ListImagesResponseBodyDataImageList {
	s.DataId = &v
	return s
}

func (s *ListImagesResponseBodyDataImageList) SetExtraData(v string) *ListImagesResponseBodyDataImageList {
	s.ExtraData = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type SearchImageRequest struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Limit    *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s SearchImageRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageRequest) GoString() string {
	return s.String()
}

func (s *SearchImageRequest) SetDbName(v string) *SearchImageRequest {
	s.DbName = &v
	return s
}

func (s *SearchImageRequest) SetImageUrl(v string) *SearchImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *SearchImageRequest) SetLimit(v int32) *SearchImageRequest {
	s.Limit = &v
	return s
}

type SearchImageAdvanceRequest struct {
	ImageUrlObject io.Reader `json:"ImageUrlObject,omitempty" xml:"ImageUrlObject,omitempty" require:"true"`
	DbName         *string   `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Limit          *int32    `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s SearchImageAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SearchImageAdvanceRequest) SetImageUrlObject(v io.Reader) *SearchImageAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *SearchImageAdvanceRequest) SetDbName(v string) *SearchImageAdvanceRequest {
	s.DbName = &v
	return s
}

func (s *SearchImageAdvanceRequest) SetLimit(v int32) *SearchImageAdvanceRequest {
	s.Limit = &v
	return s
}

type SearchImageResponseBody struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *SearchImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s SearchImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBody) SetRequestId(v string) *SearchImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchImageResponseBody) SetData(v *SearchImageResponseBodyData) *SearchImageResponseBody {
	s.Data = v
	return s
}

type SearchImageResponseBodyData struct {
	MatchList []*SearchImageResponseBodyDataMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s SearchImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyData) SetMatchList(v []*SearchImageResponseBodyDataMatchList) *SearchImageResponseBodyData {
	s.MatchList = v
	return s
}

type SearchImageResponseBodyDataMatchList struct {
	ImageUrl  *string  `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	EntityId  *string  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Score     *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	DataId    *string  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	ExtraData *string  `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
}

func (s SearchImageResponseBodyDataMatchList) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponseBodyDataMatchList) GoString() string {
	return s.String()
}

func (s *SearchImageResponseBodyDataMatchList) SetImageUrl(v string) *SearchImageResponseBodyDataMatchList {
	s.ImageUrl = &v
	return s
}

func (s *SearchImageResponseBodyDataMatchList) SetEntityId(v string) *SearchImageResponseBodyDataMatchList {
	s.EntityId = &v
	return s
}

func (s *SearchImageResponseBodyDataMatchList) SetScore(v float32) *SearchImageResponseBodyDataMatchList {
	s.Score = &v
	return s
}

func (s *SearchImageResponseBodyDataMatchList) SetDataId(v string) *SearchImageResponseBodyDataMatchList {
	s.DataId = &v
	return s
}

func (s *SearchImageResponseBodyDataMatchList) SetExtraData(v string) *SearchImageResponseBodyDataMatchList {
	s.ExtraData = &v
	return s
}

type SearchImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchImageResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchImageResponse) GoString() string {
	return s.String()
}

func (s *SearchImageResponse) SetHeaders(v map[string]*string) *SearchImageResponse {
	s.Headers = v
	return s
}

func (s *SearchImageResponse) SetBody(v *SearchImageResponseBody) *SearchImageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("imgsearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddImageWithOptions(request *AddImageRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddImage"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddImage(request *AddImageRequest) (_result *AddImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddImageResponse{}
	_body, _err := client.AddImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddImageAdvance(request *AddImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *AddImageResponse, _err error) {
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
		Product:  tea.String("imgsearch"),
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
	addImageReq := &AddImageRequest{}
	openapiutil.Convert(request, addImageReq)
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
		Content:     request.ImageUrlObject,
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
	addImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	addImageResp, _err := client.AddImageWithOptions(addImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = addImageResp
	return _result, _err
}

func (client *Client) CreateImageDbWithOptions(request *CreateImageDbRequest, runtime *util.RuntimeOptions) (_result *CreateImageDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateImageDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateImageDb"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImageDb(request *CreateImageDbRequest) (_result *CreateImageDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageDbResponse{}
	_body, _err := client.CreateImageDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageDbWithOptions(request *DeleteImageDbRequest, runtime *util.RuntimeOptions) (_result *DeleteImageDbResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteImageDbResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImageDb"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImageDb(request *DeleteImageDbRequest) (_result *DeleteImageDbResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageDbResponse{}
	_body, _err := client.DeleteImageDbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImageDbsWithOptions(runtime *util.RuntimeOptions) (_result *ListImageDbsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListImageDbsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListImageDbs"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImageDbs() (_result *ListImageDbsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImageDbsResponse{}
	_body, _err := client.ListImageDbsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListImages"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageWithOptions(request *SearchImageRequest, runtime *util.RuntimeOptions) (_result *SearchImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchImage"), tea.String("2020-03-20"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchImage(request *SearchImageRequest) (_result *SearchImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchImageResponse{}
	_body, _err := client.SearchImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchImageAdvance(request *SearchImageAdvanceRequest, runtime *util.RuntimeOptions) (_result *SearchImageResponse, _err error) {
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
		Product:  tea.String("imgsearch"),
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
	searchImageReq := &SearchImageRequest{}
	openapiutil.Convert(request, searchImageReq)
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
		Content:     request.ImageUrlObject,
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
	searchImageReq.ImageUrl = tea.String("http://" + tea.StringValue(authResponse.Bucket) + "." + tea.StringValue(authResponse.Endpoint) + "/" + tea.StringValue(authResponse.ObjectKey))
	searchImageResp, _err := client.SearchImageWithOptions(searchImageReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = searchImageResp
	return _result, _err
}
