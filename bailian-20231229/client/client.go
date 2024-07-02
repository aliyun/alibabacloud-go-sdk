// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) SetCategoryId(v string) *AddFileRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileRequest) SetLeaseId(v string) *AddFileRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileRequest) SetParser(v string) *AddFileRequest {
	s.Parser = &v
	return s
}

type AddFileResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) SetCode(v string) *AddFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddFileResponseBody) SetData(v *AddFileResponseBodyData) *AddFileResponseBody {
	s.Data = v
	return s
}

func (s *AddFileResponseBody) SetMessage(v string) *AddFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddFileResponseBody) SetRequestId(v string) *AddFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFileResponseBody) SetStatus(v string) *AddFileResponseBody {
	s.Status = &v
	return s
}

func (s *AddFileResponseBody) SetSuccess(v string) *AddFileResponseBody {
	s.Success = &v
	return s
}

type AddFileResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFileResponseBodyData) SetFileId(v string) *AddFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBodyData) SetParser(v string) *AddFileResponseBodyData {
	s.Parser = &v
	return s
}

type AddFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetStatusCode(v int32) *AddFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

type ApplyFileUploadLeaseRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19657c391f6c70bcea63c154d8606bb3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SizeInBytes *string `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
}

func (s ApplyFileUploadLeaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseRequest) SetFileName(v string) *ApplyFileUploadLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetMd5(v string) *ApplyFileUploadLeaseRequest {
	s.Md5 = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest {
	s.SizeInBytes = &v
	return s
}

type ApplyFileUploadLeaseResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ApplyFileUploadLeaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBody) SetCode(v string) *ApplyFileUploadLeaseResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetData(v *ApplyFileUploadLeaseResponseBodyData) *ApplyFileUploadLeaseResponseBody {
	s.Data = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetMessage(v string) *ApplyFileUploadLeaseResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetRequestId(v string) *ApplyFileUploadLeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetStatus(v string) *ApplyFileUploadLeaseResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetSuccess(v bool) *ApplyFileUploadLeaseResponseBody {
	s.Success = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyData struct {
	// example:
	//
	// 1e6a159107384782be5e45ac4759b247.1719325231035
	FileUploadLeaseId *string                                    `json:"FileUploadLeaseId,omitempty" xml:"FileUploadLeaseId,omitempty"`
	Param             *ApplyFileUploadLeaseResponseBodyDataParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetFileUploadLeaseId(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.FileUploadLeaseId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetParam(v *ApplyFileUploadLeaseResponseBodyDataParam) *ApplyFileUploadLeaseResponseBodyData {
	s.Param = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetType(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.Type = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyDataParam struct {
	// example:
	//
	// "X-bailian-extra": "MTAwNTQyNjQ5NTE2OTE3OA==",
	//
	//         "Content-Type": "application/pdf"
	Headers interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// PUT
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// https://bailian-datahub-data-origin-prod.oss-cn-hangzhou.aliyuncs.com/1005426495169178/10024405/68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847.pdf?Expires=1716699536&OSSAccessKeyId=TestID&Signature=HfwPUZo4pR6DatSDym0zFKVh9Wg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetHeaders(v interface{}) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetMethod(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Method = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetUrl(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Url = &v
	return s
}

type ApplyFileUploadLeaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFileUploadLeaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFileUploadLeaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponse) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponse) SetHeaders(v map[string]*string) *ApplyFileUploadLeaseResponse {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetStatusCode(v int32) *ApplyFileUploadLeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetBody(v *ApplyFileUploadLeaseResponseBody) *ApplyFileUploadLeaseResponse {
	s.Body = v
	return s
}

type DescribeFileResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBody) SetCode(v string) *DescribeFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileResponseBody) SetData(v *DescribeFileResponseBodyData) *DescribeFileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileResponseBody) SetMessage(v string) *DescribeFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFileResponseBody) SetRequestId(v string) *DescribeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileResponseBody) SetStatus(v string) *DescribeFileResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFileResponseBody) SetSuccess(v bool) *DescribeFileResponseBody {
	s.Success = &v
	return s
}

type DescribeFileResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2024-05-26 12:45:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// test.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBodyData) SetCategoryId(v string) *DescribeFileResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetCreateTime(v string) *DescribeFileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileId(v string) *DescribeFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileName(v string) *DescribeFileResponseBodyData {
	s.FileName = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileType(v string) *DescribeFileResponseBodyData {
	s.FileType = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParser(v string) *DescribeFileResponseBodyData {
	s.Parser = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetSizeInBytes(v int64) *DescribeFileResponseBodyData {
	s.SizeInBytes = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetStatus(v string) *DescribeFileResponseBodyData {
	s.Status = &v
	return s
}

type DescribeFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileResponse) SetHeaders(v map[string]*string) *DescribeFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileResponse) SetStatusCode(v int32) *DescribeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileResponse) SetBody(v *DescribeFileResponseBody) *DescribeFileResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFileResponse
func (client *Client) AddFileWithOptions(WorkspaceId *string, request *AddFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LeaseId)) {
		body["LeaseId"] = request.LeaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Parser)) {
		body["Parser"] = request.Parser
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @return AddFileResponse
func (client *Client) AddFile(WorkspaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLeaseWithOptions(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyFileUploadLeaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		body["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInBytes)) {
		body["SizeInBytes"] = request.SizeInBytes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyFileUploadLease"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/category/" + tea.StringValue(openapiutil.GetEncodeParam(CategoryId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLease(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest) (_result *ApplyFileUploadLeaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.ApplyFileUploadLeaseWithOptions(CategoryId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileResponse
func (client *Client) DescribeFileWithOptions(WorkspaceId *string, FileId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file/" + tea.StringValue(openapiutil.GetEncodeParam(FileId)) + "/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @return DescribeFileResponse
func (client *Client) DescribeFile(WorkspaceId *string, FileId *string) (_result *DescribeFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(WorkspaceId, FileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
