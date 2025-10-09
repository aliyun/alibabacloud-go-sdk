// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("emas-appmonitor"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取错误事件详情
//
// @param request - GetErrorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorResponse
func (client *Client) GetErrorWithOptions(request *GetErrorRequest, runtime *dara.RuntimeOptions) (_result *GetErrorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.BizModule) {
		body["BizModule"] = request.BizModule
	}

	if !dara.IsNil(request.ClientTime) {
		body["ClientTime"] = request.ClientTime
	}

	if !dara.IsNil(request.Did) {
		body["Did"] = request.Did
	}

	if !dara.IsNil(request.DigestHash) {
		body["DigestHash"] = request.DigestHash
	}

	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetError"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取错误事件详情
//
// @param request - GetErrorRequest
//
// @return GetErrorResponse
func (client *Client) GetError(request *GetErrorRequest) (_result *GetErrorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetErrorResponse{}
	_body, _err := client.GetErrorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某一聚合错误下所有的错误事件列表
//
// @param tmpReq - GetErrorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorsResponse
func (client *Client) GetErrorsWithOptions(tmpReq *GetErrorsRequest, runtime *dara.RuntimeOptions) (_result *GetErrorsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetErrorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.BizModule) {
		body["BizModule"] = request.BizModule
	}

	if !dara.IsNil(request.DigestHash) {
		body["DigestHash"] = request.DigestHash
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TimeRange) {
		bodyFlat["TimeRange"] = request.TimeRange
	}

	if !dara.IsNil(request.Utdid) {
		body["Utdid"] = request.Utdid
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErrors"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErrorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某一聚合错误下所有的错误事件列表
//
// @param request - GetErrorsRequest
//
// @return GetErrorsResponse
func (client *Client) GetErrors(request *GetErrorsRequest) (_result *GetErrorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetErrorsResponse{}
	_body, _err := client.GetErrorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取聚合错误详情
//
// @param tmpReq - GetIssueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIssueResponse
func (client *Client) GetIssueWithOptions(tmpReq *GetIssueRequest, runtime *dara.RuntimeOptions) (_result *GetIssueResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetIssueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.BizModule) {
		body["BizModule"] = request.BizModule
	}

	if !dara.IsNil(request.DigestHash) {
		body["DigestHash"] = request.DigestHash
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TimeRange) {
		bodyFlat["TimeRange"] = request.TimeRange
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIssue"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIssueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取聚合错误详情
//
// @param request - GetIssueRequest
//
// @return GetIssueResponse
func (client *Client) GetIssue(request *GetIssueRequest) (_result *GetIssueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIssueResponse{}
	_body, _err := client.GetIssueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取聚合错误列表
//
// @param tmpReq - GetIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIssuesResponse
func (client *Client) GetIssuesWithOptions(tmpReq *GetIssuesRequest, runtime *dara.RuntimeOptions) (_result *GetIssuesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetIssuesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.BizModule) {
		body["BizModule"] = request.BizModule
	}

	if !dara.IsNil(request.FilterShrink) {
		body["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		body["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TimeRange) {
		bodyFlat["TimeRange"] = request.TimeRange
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIssues"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIssuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取聚合错误列表
//
// @param request - GetIssuesRequest
//
// @return GetIssuesResponse
func (client *Client) GetIssues(request *GetIssuesRequest) (_result *GetIssuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIssuesResponse{}
	_body, _err := client.GetIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取符号表文件列表
//
// @param request - GetSymbolicFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSymbolicFilesResponse
func (client *Client) GetSymbolicFilesWithOptions(request *GetSymbolicFilesRequest, runtime *dara.RuntimeOptions) (_result *GetSymbolicFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExportStatus) {
		body["ExportStatus"] = request.ExportStatus
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.PageIndex) {
		body["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSymbolicFiles"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSymbolicFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取符号表文件列表
//
// @param request - GetSymbolicFilesRequest
//
// @return GetSymbolicFilesResponse
func (client *Client) GetSymbolicFiles(request *GetSymbolicFilesRequest) (_result *GetSymbolicFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSymbolicFilesResponse{}
	_body, _err := client.GetSymbolicFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RequestUploadToken
//
// @param request - RequestUploadTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequestUploadTokenResponse
func (client *Client) RequestUploadTokenWithOptions(request *RequestUploadTokenRequest, runtime *dara.RuntimeOptions) (_result *RequestUploadTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RequestUploadToken"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RequestUploadTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RequestUploadToken
//
// @param request - RequestUploadTokenRequest
//
// @return RequestUploadTokenResponse
func (client *Client) RequestUploadToken(request *RequestUploadTokenRequest) (_result *RequestUploadTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RequestUploadTokenResponse{}
	_body, _err := client.RequestUploadTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SubmitSymbolic
//
// @param request - SubmitSymbolicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSymbolicResponse
func (client *Client) SubmitSymbolicWithOptions(request *SubmitSymbolicRequest, runtime *dara.RuntimeOptions) (_result *SubmitSymbolicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppVersion) {
		body["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSymbolic"),
		Version:     dara.String("2019-06-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSymbolicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SubmitSymbolic
//
// @param request - SubmitSymbolicRequest
//
// @return SubmitSymbolicResponse
func (client *Client) SubmitSymbolic(request *SubmitSymbolicRequest) (_result *SubmitSymbolicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSymbolicResponse{}
	_body, _err := client.SubmitSymbolicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
