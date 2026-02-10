// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 导入OSS数据集
//
// @param request - AddOSSMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOSSMultimodalFineTuneDatasetResponse
func (client *Client) AddOSSMultimodalFineTuneDatasetWithContext(ctx context.Context, request *AddOSSMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *AddOSSMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("AddOSSMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOSSMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigCreateResponse
func (client *Client) ChatBIConfigCreateWithContext(ctx context.Context, request *ChatBIConfigCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigDeleteResponse
func (client *Client) ChatBIConfigDeleteWithContext(ctx context.Context, request *ChatBIConfigDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigDeleteEntryResponse
func (client *Client) ChatBIConfigDeleteEntryWithContext(ctx context.Context, request *ChatBIConfigDeleteEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigDeleteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigQueryEntriesResponse
func (client *Client) ChatBIConfigQueryEntriesWithContext(ctx context.Context, request *ChatBIConfigQueryEntriesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigQueryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigQueryTablesResponse
func (client *Client) ChatBIConfigQueryTablesWithContext(ctx context.Context, request *ChatBIConfigQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIConfigUpdateEntryResponse
func (client *Client) ChatBIConfigUpdateEntryWithContext(ctx context.Context, request *ChatBIConfigUpdateEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIConfigUpdateEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileTemplateDownloadResponse
func (client *Client) ChatBIFileTemplateDownloadWithContext(ctx context.Context, request *ChatBIFileTemplateDownloadRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileTemplateDownloadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileUploadResponse
func (client *Client) ChatBIFileUploadWithContext(ctx context.Context, request *ChatBIFileUploadRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileUploadResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIFileUploadCallbackResponse
func (client *Client) ChatBIFileUploadCallbackWithContext(ctx context.Context, request *ChatBIFileUploadCallbackRequest, runtime *dara.RuntimeOptions) (_result *ChatBIFileUploadCallbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternCreateResponse
func (client *Client) ChatBIPatternCreateWithContext(ctx context.Context, request *ChatBIPatternCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternDeleteResponse
func (client *Client) ChatBIPatternDeleteWithContext(ctx context.Context, request *ChatBIPatternDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternDeleteEntryResponse
func (client *Client) ChatBIPatternDeleteEntryWithContext(ctx context.Context, request *ChatBIPatternDeleteEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternDeleteEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexCreateResponse
func (client *Client) ChatBIPatternIndexCreateWithContext(ctx context.Context, request *ChatBIPatternIndexCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexDeleteResponse
func (client *Client) ChatBIPatternIndexDeleteWithContext(ctx context.Context, request *ChatBIPatternIndexDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternIndexQueryTablesResponse
func (client *Client) ChatBIPatternIndexQueryTablesWithContext(ctx context.Context, request *ChatBIPatternIndexQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternIndexQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternQueryEntriesResponse
func (client *Client) ChatBIPatternQueryEntriesWithContext(ctx context.Context, request *ChatBIPatternQueryEntriesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternQueryEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternQueryTablesResponse
func (client *Client) ChatBIPatternQueryTablesWithContext(ctx context.Context, request *ChatBIPatternQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPatternUpdateEntryResponse
func (client *Client) ChatBIPatternUpdateEntryWithContext(ctx context.Context, request *ChatBIPatternUpdateEntryRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPatternUpdateEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ChatBIPredictSseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIPredictSseResponse
func (client *Client) ChatBIPredictSseWithSSECtx(ctx context.Context, tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions, _yield chan *ChatBIPredictSseResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatBIPredictSseWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
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
func (client *Client) ChatBIPredictSseWithContext(ctx context.Context, tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions) (_result *ChatBIPredictSseResponse, _err error) {
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
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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

	if !dara.IsNil(request.ThinkingMode) {
		query["ThinkingMode"] = request.ThinkingMode
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexCreateResponse
func (client *Client) ChatBISchemaIndexCreateWithContext(ctx context.Context, request *ChatBISchemaIndexCreateRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexDeleteResponse
func (client *Client) ChatBISchemaIndexDeleteWithContext(ctx context.Context, request *ChatBISchemaIndexDeleteRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexDeleteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBISchemaIndexQueryTablesResponse
func (client *Client) ChatBISchemaIndexQueryTablesWithContext(ctx context.Context, request *ChatBISchemaIndexQueryTablesRequest, runtime *dara.RuntimeOptions) (_result *ChatBISchemaIndexQueryTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatBIUpdateTableValidationColumnsResponse
func (client *Client) ChatBIUpdateTableValidationColumnsWithContext(ctx context.Context, request *ChatBIUpdateTableValidationColumnsRequest, runtime *dara.RuntimeOptions) (_result *ChatBIUpdateTableValidationColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalDatasetResponse
func (client *Client) CreateMultimodalDatasetWithContext(ctx context.Context, request *CreateMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalDatasetEmbeddingResponse
func (client *Client) CreateMultimodalDatasetEmbeddingWithContext(ctx context.Context, request *CreateMultimodalDatasetEmbeddingRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalDatasetEmbeddingResponse, _err error) {
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

	if !dara.IsNil(request.ModelMode) {
		query["ModelMode"] = request.ModelMode
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建多模态微调数据集
//
// @param request - CreateMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalFineTuneDatasetResponse
func (client *Client) CreateMultimodalFineTuneDatasetWithContext(ctx context.Context, request *CreateMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("CreateMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 部署打标服务
//
// @param tmpReq - CreateMultimodalLabelStudioServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalLabelStudioServiceResponse
func (client *Client) CreateMultimodalLabelStudioServiceWithContext(ctx context.Context, tmpReq *CreateMultimodalLabelStudioServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalLabelStudioServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMultimodalLabelStudioServiceShrinkRequest{}
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultimodalLabelStudioService"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalLabelStudioServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - CreateMultimodalSearchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalSearchTaskResponse
func (client *Client) CreateMultimodalSearchTaskWithContext(ctx context.Context, tmpReq *CreateMultimodalSearchTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalSearchTaskResponse, _err error) {
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

	if !dara.IsNil(request.ModelMode) {
		query["ModelMode"] = request.ModelMode
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从检索结果中创建微调数据集
//
// @param tmpReq - CreateMultimodalSearchTaskResultFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMultimodalSearchTaskResultFineTuneDatasetResponse
func (client *Client) CreateMultimodalSearchTaskResultFineTuneDatasetWithContext(ctx context.Context, tmpReq *CreateMultimodalSearchTaskResultFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateMultimodalSearchTaskResultFineTuneDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResultIndex) {
		request.ResultIndexShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResultIndex, dara.String("ResultIndex"), dara.String("json"))
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

	if !dara.IsNil(request.ResultIndexShrink) {
		query["ResultIndex"] = request.ResultIndexShrink
	}

	if !dara.IsNil(request.ResultMode) {
		query["ResultMode"] = request.ResultMode
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMultimodalSearchTaskResultFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMultimodalSearchTaskResultFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalDatasetResponse
func (client *Client) DeleteMultimodalDatasetWithContext(ctx context.Context, request *DeleteMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalEmbeddingResponse
func (client *Client) DeleteMultimodalEmbeddingWithContext(ctx context.Context, request *DeleteMultimodalEmbeddingRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalEmbeddingResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除多模态微调数据集
//
// @param request - DeleteMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalFineTuneDatasetResponse
func (client *Client) DeleteMultimodalFineTuneDatasetWithContext(ctx context.Context, request *DeleteMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("DeleteMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型mode可选列表
//
// @param request - DeleteMultimodalLabelStudioServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMultimodalLabelStudioServiceResponse
func (client *Client) DeleteMultimodalLabelStudioServiceWithContext(ctx context.Context, request *DeleteMultimodalLabelStudioServiceRequest, runtime *dara.RuntimeOptions) (_result *DeleteMultimodalLabelStudioServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMultimodalLabelStudioService"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMultimodalLabelStudioServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 微调数据集删除导入的OSS路径
//
// @param request - DeleteOSSMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOSSMultimodalFineTuneDatasetResponse
func (client *Client) DeleteOSSMultimodalFineTuneDatasetWithContext(ctx context.Context, request *DeleteOSSMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteOSSMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("DeleteOSSMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOSSMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadMultimodalSearchTaskResultMetadataResponse
func (client *Client) DownloadMultimodalSearchTaskResultMetadataWithContext(ctx context.Context, request *DownloadMultimodalSearchTaskResultMetadataRequest, runtime *dara.RuntimeOptions) (_result *DownloadMultimodalSearchTaskResultMetadataResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得鉴权token
//
// @param request - GetUserTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserTokenResponse
func (client *Client) GetUserTokenWithContext(ctx context.Context, request *GetUserTokenRequest, runtime *dara.RuntimeOptions) (_result *GetUserTokenResponse, _err error) {
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserToken"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalDatasetResponse
func (client *Client) ListMultimodalDatasetWithContext(ctx context.Context, request *ListMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalEmbeddingModelResponse
func (client *Client) ListMultimodalEmbeddingModelWithContext(ctx context.Context, request *ListMultimodalEmbeddingModelRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalEmbeddingModelResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模型mode可选列表
//
// @param request - ListMultimodalEmbeddingModelModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalEmbeddingModelModeResponse
func (client *Client) ListMultimodalEmbeddingModelModeWithContext(ctx context.Context, request *ListMultimodalEmbeddingModelModeRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalEmbeddingModelModeResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalEmbeddingModelMode"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalEmbeddingModelModeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param request - ListMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalFineTuneDatasetResponse
func (client *Client) ListMultimodalFineTuneDatasetWithContext(ctx context.Context, request *ListMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("ListMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询打标服务信息
//
// @param request - ListMultimodalLabelStudioServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalLabelStudioServiceResponse
func (client *Client) ListMultimodalLabelStudioServiceWithContext(ctx context.Context, request *ListMultimodalLabelStudioServiceRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalLabelStudioServiceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalLabelStudioService"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalLabelStudioServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询search模型列表
//
// @param request - ListMultimodalSearchModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalSearchModelResponse
func (client *Client) ListMultimodalSearchModelWithContext(ctx context.Context, request *ListMultimodalSearchModelRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalSearchModelResponse, _err error) {
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
		Action:      dara.String("ListMultimodalSearchModel"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalSearchModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListMultimodalSearchTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalSearchTaskResponse
func (client *Client) ListMultimodalSearchTaskWithContext(ctx context.Context, tmpReq *ListMultimodalSearchTaskRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalSearchTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListMultimodalSearchTaskShrinkRequest{}
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

	if !dara.IsNil(request.InputField) {
		query["InputField"] = request.InputField
	}

	if !dara.IsNil(request.ModelMode) {
		query["ModelMode"] = request.ModelMode
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询search task结果列表
//
// @param request - ListMultimodalSearchTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMultimodalSearchTaskResultResponse
func (client *Client) ListMultimodalSearchTaskResultWithContext(ctx context.Context, request *ListMultimodalSearchTaskResultRequest, runtime *dara.RuntimeOptions) (_result *ListMultimodalSearchTaskResultResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMultimodalSearchTaskResult"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMultimodalSearchTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultimodalDatasetResponse
func (client *Client) UpdateMultimodalDatasetWithContext(ctx context.Context, request *UpdateMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultimodalDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新微调数据集信息
//
// @param request - UpdateMultimodalFineTuneDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultimodalFineTuneDatasetResponse
func (client *Client) UpdateMultimodalFineTuneDatasetWithContext(ctx context.Context, request *UpdateMultimodalFineTuneDatasetRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultimodalFineTuneDatasetResponse, _err error) {
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
		Action:      dara.String("UpdateMultimodalFineTuneDataset"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultimodalFineTuneDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 为打标服务覆盖配置白名单
//
// @param tmpReq - UpdateMultimodalLabelStudioServiceWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMultimodalLabelStudioServiceWhiteListResponse
func (client *Client) UpdateMultimodalLabelStudioServiceWhiteListWithContext(ctx context.Context, tmpReq *UpdateMultimodalLabelStudioServiceWhiteListRequest, runtime *dara.RuntimeOptions) (_result *UpdateMultimodalLabelStudioServiceWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateMultimodalLabelStudioServiceWhiteListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WhiteList) {
		request.WhiteListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhiteList, dara.String("WhiteList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.WhiteListShrink) {
		query["WhiteList"] = request.WhiteListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMultimodalLabelStudioServiceWhiteList"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMultimodalLabelStudioServiceWhiteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadOSSMultimodalDatasetResponse
func (client *Client) UploadOSSMultimodalDatasetWithContext(ctx context.Context, request *UploadOSSMultimodalDatasetRequest, runtime *dara.RuntimeOptions) (_result *UploadOSSMultimodalDatasetResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预检验鉴权token
//
// @param request - ValidateDatabaseUserTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateDatabaseUserTokenResponse
func (client *Client) ValidateDatabaseUserTokenWithContext(ctx context.Context, request *ValidateDatabaseUserTokenRequest, runtime *dara.RuntimeOptions) (_result *ValidateDatabaseUserTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBName) {
		query["DBName"] = request.DBName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateDatabaseUserToken"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateDatabaseUserTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预检验鉴权token
//
// @param request - ValidateUserTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateUserTokenResponse
func (client *Client) ValidateUserTokenWithContext(ctx context.Context, request *ValidateUserTokenRequest, runtime *dara.RuntimeOptions) (_result *ValidateUserTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateUserToken"),
		Version:     dara.String("2025-10-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateUserTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) chatBIPredictSseWithSSECtx_opYieldFunc(_yield chan *ChatBIPredictSseResponse, _yieldErr chan error, ctx context.Context, tmpReq *ChatBIPredictSseRequest, runtime *dara.RuntimeOptions) {
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
	if !dara.IsNil(request.AuthMessage) {
		query["AuthMessage"] = request.AuthMessage
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

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

	if !dara.IsNil(request.ThinkingMode) {
		query["ThinkingMode"] = request.ThinkingMode
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
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
