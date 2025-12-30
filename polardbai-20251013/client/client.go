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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("polardbai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建chatbi配置表
//
// @param request - ChatBIConfigCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigCreateResponse
func (client *Client) ChatBIConfigCreateWithOptions(request *ChatBIConfigCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigCreate"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi配置表
//
// @param request - ChatBIConfigCreateRequest
//
// @return ChatBIConfigCreateResponse
func (client *Client) ChatBIConfigCreate(request *ChatBIConfigCreateRequest) (_result *ChatBIConfigCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigCreateResponse{}
	_body, _err := client.ChatBIConfigCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除chatbi配置表
//
// @param request - ChatBIConfigDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigDeleteResponse
func (client *Client) ChatBIConfigDeleteWithOptions(request *ChatBIConfigDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigDelete"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除chatbi配置表
//
// @param request - ChatBIConfigDeleteRequest
//
// @return ChatBIConfigDeleteResponse
func (client *Client) ChatBIConfigDelete(request *ChatBIConfigDeleteRequest) (_result *ChatBIConfigDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigDeleteResponse{}
	_body, _err := client.ChatBIConfigDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIConfigDeleteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigDeleteEntryResponse
func (client *Client) ChatBIConfigDeleteEntryWithOptions(request *ChatBIConfigDeleteEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigDeleteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigDeleteEntry"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigDeleteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIConfigDeleteEntryRequest
//
// @return ChatBIConfigDeleteEntryResponse
func (client *Client) ChatBIConfigDeleteEntry(request *ChatBIConfigDeleteEntryRequest) (_result *ChatBIConfigDeleteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigDeleteEntryResponse{}
	_body, _err := client.ChatBIConfigDeleteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi配置表内容
//
// @param request - ChatBIConfigQueryEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigQueryEntriesResponse
func (client *Client) ChatBIConfigQueryEntriesWithOptions(request *ChatBIConfigQueryEntriesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigQueryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigQueryEntries"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigQueryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi配置表内容
//
// @param request - ChatBIConfigQueryEntriesRequest
//
// @return ChatBIConfigQueryEntriesResponse
func (client *Client) ChatBIConfigQueryEntries(request *ChatBIConfigQueryEntriesRequest) (_result *ChatBIConfigQueryEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigQueryEntriesResponse{}
	_body, _err := client.ChatBIConfigQueryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户chatbi配置表名
//
// @param request - ChatBIConfigQueryTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigQueryTablesResponse
func (client *Client) ChatBIConfigQueryTablesWithOptions(request *ChatBIConfigQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigQueryTables"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigQueryTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户chatbi配置表名
//
// @param request - ChatBIConfigQueryTablesRequest
//
// @return ChatBIConfigQueryTablesResponse
func (client *Client) ChatBIConfigQueryTables(request *ChatBIConfigQueryTablesRequest) (_result *ChatBIConfigQueryTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigQueryTablesResponse{}
	_body, _err := client.ChatBIConfigQueryTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改chatbi配置表内容
//
// @param request - ChatBIConfigUpdateEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigUpdateEntryResponse
func (client *Client) ChatBIConfigUpdateEntryWithOptions(request *ChatBIConfigUpdateEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigUpdateEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FormulaFunction) {
		query["FormulaFunction"] = request.FormulaFunction
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IsFunctional) {
		query["IsFunctional"] = request.IsFunctional
	}

	if !dara.IsNil(request.QueryFunction) {
		query["QueryFunction"] = request.QueryFunction
	}

	if !dara.IsNil(request.SqlCondition) {
		query["SqlCondition"] = request.SqlCondition
	}

	if !dara.IsNil(request.SqlFunction) {
		query["SqlFunction"] = request.SqlFunction
	}

	if !dara.IsNil(request.TextCondition) {
		query["TextCondition"] = request.TextCondition
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIConfigUpdateEntry"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIConfigUpdateEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改chatbi配置表内容
//
// @param request - ChatBIConfigUpdateEntryRequest
//
// @return ChatBIConfigUpdateEntryResponse
func (client *Client) ChatBIConfigUpdateEntry(request *ChatBIConfigUpdateEntryRequest) (_result *ChatBIConfigUpdateEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIConfigUpdateEntryResponse{}
	_body, _err := client.ChatBIConfigUpdateEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileTemplateDownloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileTemplateDownloadResponse
func (client *Client) ChatBIFileTemplateDownloadWithOptions(request *ChatBIFileTemplateDownloadRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileTemplateDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIFileTemplateDownload"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIFileTemplateDownloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileTemplateDownloadRequest
//
// @return ChatBIFileTemplateDownloadResponse
func (client *Client) ChatBIFileTemplateDownload(request *ChatBIFileTemplateDownloadRequest) (_result *ChatBIFileTemplateDownloadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIFileTemplateDownloadResponse{}
	_body, _err := client.ChatBIFileTemplateDownloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileUploadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileUploadResponse
func (client *Client) ChatBIFileUploadWithOptions(request *ChatBIFileUploadRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIFileUpload"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIFileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileUploadRequest
//
// @return ChatBIFileUploadResponse
func (client *Client) ChatBIFileUpload(request *ChatBIFileUploadRequest) (_result *ChatBIFileUploadResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIFileUploadResponse{}
	_body, _err := client.ChatBIFileUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileUploadCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileUploadCallbackResponse
func (client *Client) ChatBIFileUploadCallbackWithOptions(request *ChatBIFileUploadCallbackRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileUploadCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CharacterSetName) {
		query["CharacterSetName"] = request.CharacterSetName
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIFileUploadCallback"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIFileUploadCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIFileUploadCallbackRequest
//
// @return ChatBIFileUploadCallbackResponse
func (client *Client) ChatBIFileUploadCallback(request *ChatBIFileUploadCallbackRequest) (_result *ChatBIFileUploadCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIFileUploadCallbackResponse{}
	_body, _err := client.ChatBIFileUploadCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIPatternCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternCreateResponse
func (client *Client) ChatBIPatternCreateWithOptions(request *ChatBIPatternCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableNameSuffix) {
		query["TableNameSuffix"] = request.TableNameSuffix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternCreate"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi问题模板表
//
// @param request - ChatBIPatternCreateRequest
//
// @return ChatBIPatternCreateResponse
func (client *Client) ChatBIPatternCreate(request *ChatBIPatternCreateRequest) (_result *ChatBIPatternCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternCreateResponse{}
	_body, _err := client.ChatBIPatternCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除chatbi问题模板表
//
// @param request - ChatBIPatternDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternDeleteResponse
func (client *Client) ChatBIPatternDeleteWithOptions(request *ChatBIPatternDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternDelete"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除chatbi问题模板表
//
// @param request - ChatBIPatternDeleteRequest
//
// @return ChatBIPatternDeleteResponse
func (client *Client) ChatBIPatternDelete(request *ChatBIPatternDeleteRequest) (_result *ChatBIPatternDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternDeleteResponse{}
	_body, _err := client.ChatBIPatternDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIPatternDeleteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternDeleteEntryResponse
func (client *Client) ChatBIPatternDeleteEntryWithOptions(request *ChatBIPatternDeleteEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternDeleteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternDeleteEntry"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternDeleteEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIPatternDeleteEntryRequest
//
// @return ChatBIPatternDeleteEntryResponse
func (client *Client) ChatBIPatternDeleteEntry(request *ChatBIPatternDeleteEntryRequest) (_result *ChatBIPatternDeleteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternDeleteEntryResponse{}
	_body, _err := client.ChatBIPatternDeleteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi检索索引表
//
// @param request - ChatBIPatternIndexCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexCreateResponse
func (client *Client) ChatBIPatternIndexCreateWithOptions(request *ChatBIPatternIndexCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PatternTableName) {
		query["PatternTableName"] = request.PatternTableName
	}

	if !dara.IsNil(request.TableNameSuffix) {
		query["TableNameSuffix"] = request.TableNameSuffix
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternIndexCreate"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternIndexCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi检索索引表
//
// @param request - ChatBIPatternIndexCreateRequest
//
// @return ChatBIPatternIndexCreateResponse
func (client *Client) ChatBIPatternIndexCreate(request *ChatBIPatternIndexCreateRequest) (_result *ChatBIPatternIndexCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternIndexCreateResponse{}
	_body, _err := client.ChatBIPatternIndexCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除检索索引表
//
// @param request - ChatBIPatternIndexDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexDeleteResponse
func (client *Client) ChatBIPatternIndexDeleteWithOptions(request *ChatBIPatternIndexDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternIndexDelete"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternIndexDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除检索索引表
//
// @param request - ChatBIPatternIndexDeleteRequest
//
// @return ChatBIPatternIndexDeleteResponse
func (client *Client) ChatBIPatternIndexDelete(request *ChatBIPatternIndexDeleteRequest) (_result *ChatBIPatternIndexDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternIndexDeleteResponse{}
	_body, _err := client.ChatBIPatternIndexDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询检索索引表名
//
// @param request - ChatBIPatternIndexQueryTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexQueryTablesResponse
func (client *Client) ChatBIPatternIndexQueryTablesWithOptions(request *ChatBIPatternIndexQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternIndexQueryTables"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternIndexQueryTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检索索引表名
//
// @param request - ChatBIPatternIndexQueryTablesRequest
//
// @return ChatBIPatternIndexQueryTablesResponse
func (client *Client) ChatBIPatternIndexQueryTables(request *ChatBIPatternIndexQueryTablesRequest) (_result *ChatBIPatternIndexQueryTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternIndexQueryTablesResponse{}
	_body, _err := client.ChatBIPatternIndexQueryTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIPatternQueryEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternQueryEntriesResponse
func (client *Client) ChatBIPatternQueryEntriesWithOptions(request *ChatBIPatternQueryEntriesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternQueryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternQueryEntries"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternQueryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIPatternQueryEntriesRequest
//
// @return ChatBIPatternQueryEntriesResponse
func (client *Client) ChatBIPatternQueryEntries(request *ChatBIPatternQueryEntriesRequest) (_result *ChatBIPatternQueryEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternQueryEntriesResponse{}
	_body, _err := client.ChatBIPatternQueryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表名
//
// @param request - ChatBIPatternQueryTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternQueryTablesResponse
func (client *Client) ChatBIPatternQueryTablesWithOptions(request *ChatBIPatternQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternQueryTables"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternQueryTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表名
//
// @param request - ChatBIPatternQueryTablesRequest
//
// @return ChatBIPatternQueryTablesResponse
func (client *Client) ChatBIPatternQueryTables(request *ChatBIPatternQueryTablesRequest) (_result *ChatBIPatternQueryTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternQueryTablesResponse{}
	_body, _err := client.ChatBIPatternQueryTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改chatbi问题模板表内容
//
// @param request - ChatBIPatternUpdateEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternUpdateEntryResponse
func (client *Client) ChatBIPatternUpdateEntryWithOptions(request *ChatBIPatternUpdateEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternUpdateEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PatternDescription) {
		query["PatternDescription"] = request.PatternDescription
	}

	if !dara.IsNil(request.PatternParams) {
		query["PatternParams"] = request.PatternParams
	}

	if !dara.IsNil(request.PatternQuestion) {
		query["PatternQuestion"] = request.PatternQuestion
	}

	if !dara.IsNil(request.PatternSql) {
		query["PatternSql"] = request.PatternSql
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPatternUpdateEntry"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPatternUpdateEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改chatbi问题模板表内容
//
// @param request - ChatBIPatternUpdateEntryRequest
//
// @return ChatBIPatternUpdateEntryResponse
func (client *Client) ChatBIPatternUpdateEntry(request *ChatBIPatternUpdateEntryRequest) (_result *ChatBIPatternUpdateEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPatternUpdateEntryResponse{}
	_body, _err := client.ChatBIPatternUpdateEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 流式返回
//
// @param tmpReq - ChatBIPredictSseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPredictSseResponse
func (client *Client) ChatBIPredictSseWithSSE(tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions, _yield chan *ChatBIPredictSseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatBIPredictSseWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// 流式返回
//
// @param tmpReq - ChatBIPredictSseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPredictSseResponse
func (client *Client) ChatBIPredictSseWithOptions(tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPredictSseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChatBIPredictSseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.GenerateChart) {
		query["GenerateChart"] = request.GenerateChart
	}

	if !dara.IsNil(request.GenerateSummary) {
		query["GenerateSummary"] = request.GenerateSummary
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.PatternIndexTableName) {
		query["PatternIndexTableName"] = request.PatternIndexTableName
	}

	if !dara.IsNil(request.Question) {
		query["Question"] = request.Question
	}

	if !dara.IsNil(request.SchemaIndexTableName) {
		query["SchemaIndexTableName"] = request.SchemaIndexTableName
	}

	if !dara.IsNil(request.SelectData) {
		query["SelectData"] = request.SelectData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPredictSse"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIPredictSseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式返回
//
// @param request - ChatBIPredictSseRequest
//
// @return ChatBIPredictSseResponse
func (client *Client) ChatBIPredictSse(request *ChatBIPredictSseRequest) (_result *ChatBIPredictSseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIPredictSseResponse{}
	_body, _err := client.ChatBIPredictSseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建chatbi检索索引表
//
// @param request - ChatBISchemaIndexCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexCreateResponse
func (client *Client) ChatBISchemaIndexCreateWithOptions(request *ChatBISchemaIndexCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnsExcluded) {
		query["ColumnsExcluded"] = request.ColumnsExcluded
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableNameSuffix) {
		query["TableNameSuffix"] = request.TableNameSuffix
	}

	if !dara.IsNil(request.TablesIncluded) {
		query["TablesIncluded"] = request.TablesIncluded
	}

	if !dara.IsNil(request.ToSample) {
		query["ToSample"] = request.ToSample
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBISchemaIndexCreate"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBISchemaIndexCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建chatbi检索索引表
//
// @param request - ChatBISchemaIndexCreateRequest
//
// @return ChatBISchemaIndexCreateResponse
func (client *Client) ChatBISchemaIndexCreate(request *ChatBISchemaIndexCreateRequest) (_result *ChatBISchemaIndexCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBISchemaIndexCreateResponse{}
	_body, _err := client.ChatBISchemaIndexCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除检索索引表
//
// @param request - ChatBISchemaIndexDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexDeleteResponse
func (client *Client) ChatBISchemaIndexDeleteWithOptions(request *ChatBISchemaIndexDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBISchemaIndexDelete"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBISchemaIndexDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除检索索引表
//
// @param request - ChatBISchemaIndexDeleteRequest
//
// @return ChatBISchemaIndexDeleteResponse
func (client *Client) ChatBISchemaIndexDelete(request *ChatBISchemaIndexDeleteRequest) (_result *ChatBISchemaIndexDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBISchemaIndexDeleteResponse{}
	_body, _err := client.ChatBISchemaIndexDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询检索索引表名
//
// @param request - ChatBISchemaIndexQueryTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexQueryTablesResponse
func (client *Client) ChatBISchemaIndexQueryTablesWithOptions(request *ChatBISchemaIndexQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBISchemaIndexQueryTables"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBISchemaIndexQueryTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询检索索引表名
//
// @param request - ChatBISchemaIndexQueryTablesRequest
//
// @return ChatBISchemaIndexQueryTablesResponse
func (client *Client) ChatBISchemaIndexQueryTables(request *ChatBISchemaIndexQueryTablesRequest) (_result *ChatBISchemaIndexQueryTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBISchemaIndexQueryTablesResponse{}
	_body, _err := client.ChatBISchemaIndexQueryTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIUpdateTableValidationColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIUpdateTableValidationColumnsResponse
func (client *Client) ChatBIUpdateTableValidationColumnsWithOptions(request *ChatBIUpdateTableValidationColumnsRequest, runtime *dara.RuntimeOptions) (_result *ChatBIUpdateTableValidationColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIUpdateTableValidationColumns"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatBIUpdateTableValidationColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询chatbi问题模板表内容
//
// @param request - ChatBIUpdateTableValidationColumnsRequest
//
// @return ChatBIUpdateTableValidationColumnsResponse
func (client *Client) ChatBIUpdateTableValidationColumns(request *ChatBIUpdateTableValidationColumnsRequest) (_result *ChatBIUpdateTableValidationColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatBIUpdateTableValidationColumnsResponse{}
	_body, _err := client.ChatBIUpdateTableValidationColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询多模态数据集列表
//
// @param request - CreateMultimodalDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalDatasetResponse
func (client *Client) CreateMultimodalDatasetWithOptions(request *CreateMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetDescription) {
		query["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultimodalDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询多模态数据集列表
//
// @param request - CreateMultimodalDatasetRequest
//
// @return CreateMultimodalDatasetResponse
func (client *Client) CreateMultimodalDataset(request *CreateMultimodalDatasetRequest) (_result *CreateMultimodalDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultimodalDatasetResponse{}
	_body, _err := client.CreateMultimodalDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Embedding
//
// @param request - CreateMultimodalDatasetEmbeddingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalDatasetEmbeddingResponse
func (client *Client) CreateMultimodalDatasetEmbeddingWithOptions(request *CreateMultimodalDatasetEmbeddingRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalDatasetEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultimodalDatasetEmbedding"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalDatasetEmbeddingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Embedding
//
// @param request - CreateMultimodalDatasetEmbeddingRequest
//
// @return CreateMultimodalDatasetEmbeddingResponse
func (client *Client) CreateMultimodalDatasetEmbedding(request *CreateMultimodalDatasetEmbeddingRequest) (_result *CreateMultimodalDatasetEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultimodalDatasetEmbeddingResponse{}
	_body, _err := client.CreateMultimodalDatasetEmbeddingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建SearchTask
//
// @param tmpReq - CreateMultimodalSearchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalSearchTaskResponse
func (client *Client) CreateMultimodalSearchTaskWithOptions(tmpReq *CreateMultimodalSearchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalSearchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMultimodalSearchTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatasetIds) {
		request.DatasetIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatasetIds, dara.String("DatasetIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetIdsShrink) {
		query["DatasetIds"] = request.DatasetIdsShrink
	}

	if !dara.IsNil(request.EmbeddingModel) {
		query["EmbeddingModel"] = request.EmbeddingModel
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.SearchModel) {
		query["SearchModel"] = request.SearchModel
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultimodalSearchTask"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalSearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建SearchTask
//
// @param request - CreateMultimodalSearchTaskRequest
//
// @return CreateMultimodalSearchTaskResponse
func (client *Client) CreateMultimodalSearchTask(request *CreateMultimodalSearchTaskRequest) (_result *CreateMultimodalSearchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMultimodalSearchTaskResponse{}
	_body, _err := client.CreateMultimodalSearchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除多模态数据集
//
// @param request - DeleteMultimodalDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalDatasetResponse
func (client *Client) DeleteMultimodalDatasetWithOptions(request *DeleteMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultimodalDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultimodalDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除多模态数据集
//
// @param request - DeleteMultimodalDatasetRequest
//
// @return DeleteMultimodalDatasetResponse
func (client *Client) DeleteMultimodalDataset(request *DeleteMultimodalDatasetRequest) (_result *DeleteMultimodalDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMultimodalDatasetResponse{}
	_body, _err := client.DeleteMultimodalDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除多模态数据集embedding
//
// @param request - DeleteMultimodalEmbeddingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalEmbeddingResponse
func (client *Client) DeleteMultimodalEmbeddingWithOptions(request *DeleteMultimodalEmbeddingRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalEmbeddingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.Embedding) {
		query["Embedding"] = request.Embedding
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultimodalEmbedding"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultimodalEmbeddingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除多模态数据集embedding
//
// @param request - DeleteMultimodalEmbeddingRequest
//
// @return DeleteMultimodalEmbeddingResponse
func (client *Client) DeleteMultimodalEmbedding(request *DeleteMultimodalEmbeddingRequest) (_result *DeleteMultimodalEmbeddingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMultimodalEmbeddingResponse{}
	_body, _err := client.DeleteMultimodalEmbeddingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下载Search任务的结果元数据
//
// @param request - DownloadMultimodalSearchTaskResultMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadMultimodalSearchTaskResultMetadataResponse
func (client *Client) DownloadMultimodalSearchTaskResultMetadataWithOptions(request *DownloadMultimodalSearchTaskResultMetadataRequest, runtime *dara.RuntimeOptions) (_result *DownloadMultimodalSearchTaskResultMetadataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadMultimodalSearchTaskResultMetadata"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadMultimodalSearchTaskResultMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下载Search任务的结果元数据
//
// @param request - DownloadMultimodalSearchTaskResultMetadataRequest
//
// @return DownloadMultimodalSearchTaskResultMetadataResponse
func (client *Client) DownloadMultimodalSearchTaskResultMetadata(request *DownloadMultimodalSearchTaskResultMetadataRequest) (_result *DownloadMultimodalSearchTaskResultMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadMultimodalSearchTaskResultMetadataResponse{}
	_body, _err := client.DownloadMultimodalSearchTaskResultMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询多模态数据集列表
//
// @param request - ListMultimodalDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalDatasetResponse
func (client *Client) ListMultimodalDatasetWithOptions(request *ListMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询多模态数据集列表
//
// @param request - ListMultimodalDatasetRequest
//
// @return ListMultimodalDatasetResponse
func (client *Client) ListMultimodalDataset(request *ListMultimodalDatasetRequest) (_result *ListMultimodalDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultimodalDatasetResponse{}
	_body, _err := client.ListMultimodalDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模型列表
//
// @param request - ListMultimodalEmbeddingModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalEmbeddingModelResponse
func (client *Client) ListMultimodalEmbeddingModelWithOptions(request *ListMultimodalEmbeddingModelRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalEmbeddingModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalEmbeddingModel"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalEmbeddingModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型列表
//
// @param request - ListMultimodalEmbeddingModelRequest
//
// @return ListMultimodalEmbeddingModelResponse
func (client *Client) ListMultimodalEmbeddingModel(request *ListMultimodalEmbeddingModelRequest) (_result *ListMultimodalEmbeddingModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultimodalEmbeddingModelResponse{}
	_body, _err := client.ListMultimodalEmbeddingModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询search task列表
//
// @param request - ListMultimodalSearchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalSearchTaskResponse
func (client *Client) ListMultimodalSearchTaskWithOptions(request *ListMultimodalSearchTaskRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalSearchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalSearchTask"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalSearchTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询search task列表
//
// @param request - ListMultimodalSearchTaskRequest
//
// @return ListMultimodalSearchTaskResponse
func (client *Client) ListMultimodalSearchTask(request *ListMultimodalSearchTaskRequest) (_result *ListMultimodalSearchTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMultimodalSearchTaskResponse{}
	_body, _err := client.ListMultimodalSearchTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新数据集信息
//
// @param request - UpdateMultimodalDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultimodalDatasetResponse
func (client *Client) UpdateMultimodalDatasetWithOptions(request *UpdateMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultimodalDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetDescription) {
		query["DatasetDescription"] = request.DatasetDescription
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultimodalDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultimodalDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集信息
//
// @param request - UpdateMultimodalDatasetRequest
//
// @return UpdateMultimodalDatasetResponse
func (client *Client) UpdateMultimodalDataset(request *UpdateMultimodalDatasetRequest) (_result *UpdateMultimodalDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMultimodalDatasetResponse{}
	_body, _err := client.UpdateMultimodalDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入OSS数据集
//
// @param request - UploadOSSMultimodalDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadOSSMultimodalDatasetResponse
func (client *Client) UploadOSSMultimodalDatasetWithOptions(request *UploadOSSMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *UploadOSSMultimodalDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.OssUrl) {
		query["OssUrl"] = request.OssUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadOSSMultimodalDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadOSSMultimodalDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入OSS数据集
//
// @param request - UploadOSSMultimodalDatasetRequest
//
// @return UploadOSSMultimodalDatasetResponse
func (client *Client) UploadOSSMultimodalDataset(request *UploadOSSMultimodalDatasetRequest) (_result *UploadOSSMultimodalDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadOSSMultimodalDatasetResponse{}
	_body, _err := client.UploadOSSMultimodalDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) chatBIPredictSseWithSSE_opYieldFunc(_yield chan *ChatBIPredictSseResponse, _yieldErr chan error, tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &ChatBIPredictSseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.GenerateChart) {
		query["GenerateChart"] = request.GenerateChart
	}

	if !dara.IsNil(request.GenerateSummary) {
		query["GenerateSummary"] = request.GenerateSummary
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.PatternIndexTableName) {
		query["PatternIndexTableName"] = request.PatternIndexTableName
	}

	if !dara.IsNil(request.Question) {
		query["Question"] = request.Question
	}

	if !dara.IsNil(request.SchemaIndexTableName) {
		query["SchemaIndexTableName"] = request.SchemaIndexTableName
	}

	if !dara.IsNil(request.SelectData) {
		query["SelectData"] = request.SelectData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatBIPredictSse"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
