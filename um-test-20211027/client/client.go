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

type GetOssUploadParamRequest struct {
	// App版本号
	AppVersion *string `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	// 数据源id（appKey)
	DataSourceId *string `json:"dataSourceId,omitempty" xml:"dataSourceId,omitempty"`
	// 文件名称，后缀只允许为txt,so,sym,zip,gz
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// 文件类型(1 mapping文件；2 so文件；3 dSYM文件压缩包)
	FileType *int32 `json:"fileType,omitempty" xml:"fileType,omitempty"`
}

func (s GetOssUploadParamRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadParamRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadParamRequest) SetAppVersion(v string) *GetOssUploadParamRequest {
	s.AppVersion = &v
	return s
}

func (s *GetOssUploadParamRequest) SetDataSourceId(v string) *GetOssUploadParamRequest {
	s.DataSourceId = &v
	return s
}

func (s *GetOssUploadParamRequest) SetFileName(v string) *GetOssUploadParamRequest {
	s.FileName = &v
	return s
}

func (s *GetOssUploadParamRequest) SetFileType(v int32) *GetOssUploadParamRequest {
	s.FileType = &v
	return s
}

type GetOssUploadParamResponseBody struct {
	// 请求唯一ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// code
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *GetOssUploadParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// 异常描述
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetOssUploadParamResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadParamResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUploadParamResponseBody) SetRequestId(v string) *GetOssUploadParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUploadParamResponseBody) SetCode(v int64) *GetOssUploadParamResponseBody {
	s.Code = &v
	return s
}

func (s *GetOssUploadParamResponseBody) SetData(v *GetOssUploadParamResponseBodyData) *GetOssUploadParamResponseBody {
	s.Data = v
	return s
}

func (s *GetOssUploadParamResponseBody) SetMsg(v string) *GetOssUploadParamResponseBody {
	s.Msg = &v
	return s
}

func (s *GetOssUploadParamResponseBody) SetSuccess(v bool) *GetOssUploadParamResponseBody {
	s.Success = &v
	return s
}

func (s *GetOssUploadParamResponseBody) SetTraceId(v string) *GetOssUploadParamResponseBody {
	s.TraceId = &v
	return s
}

type GetOssUploadParamResponseBodyData struct {
	// 文件上传表单必要参数
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// 文件上传表单必要参数
	Callback *string `json:"callback,omitempty" xml:"callback,omitempty"`
	// 文件上传表单必要参数
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// 文件上传表单必要参数
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// 文件上传表单必要参数
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// 文件上传地址
	UploadAddress *string `json:"uploadAddress,omitempty" xml:"uploadAddress,omitempty"`
}

func (s GetOssUploadParamResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOssUploadParamResponseBodyData) SetAccessKeyId(v string) *GetOssUploadParamResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetOssUploadParamResponseBodyData) SetCallback(v string) *GetOssUploadParamResponseBodyData {
	s.Callback = &v
	return s
}

func (s *GetOssUploadParamResponseBodyData) SetKey(v string) *GetOssUploadParamResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetOssUploadParamResponseBodyData) SetPolicy(v string) *GetOssUploadParamResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetOssUploadParamResponseBodyData) SetSignature(v string) *GetOssUploadParamResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetOssUploadParamResponseBodyData) SetUploadAddress(v string) *GetOssUploadParamResponseBodyData {
	s.UploadAddress = &v
	return s
}

type GetOssUploadParamResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetOssUploadParamResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOssUploadParamResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOssUploadParamResponse) GoString() string {
	return s.String()
}

func (s *GetOssUploadParamResponse) SetHeaders(v map[string]*string) *GetOssUploadParamResponse {
	s.Headers = v
	return s
}

func (s *GetOssUploadParamResponse) SetBody(v *GetOssUploadParamResponseBody) *GetOssUploadParamResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("um-test"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetOssUploadParamWithOptions(request *GetOssUploadParamRequest, runtime *util.RuntimeOptions) (_result *GetOssUploadParamResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOssUploadParam"),
		Version:     tea.String("2021-10-27"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOssUploadParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOssUploadParam(request *GetOssUploadParamRequest) (_result *GetOssUploadParamResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOssUploadParamResponse{}
	_body, _err := client.GetOssUploadParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
