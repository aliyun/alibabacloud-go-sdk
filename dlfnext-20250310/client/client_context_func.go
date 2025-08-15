// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 更新数据目录
//
// @param request - AlterCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterCatalogResponse
func (client *Client) AlterCatalogWithContext(ctx context.Context, catalog *string, request *AlterCatalogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterCatalogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Removals) {
		body["removals"] = request.Removals
	}

	if !dara.IsNil(request.Updates) {
		body["updates"] = request.Updates
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterCatalog"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs/" + dara.PercentEncode(dara.StringValue(catalog))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AlterCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据库
//
// @param request - AlterDatabaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterDatabaseResponse
func (client *Client) AlterDatabaseWithContext(ctx context.Context, catalogId *string, database *string, request *AlterDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Removals) {
		body["removals"] = request.Removals
	}

	if !dara.IsNil(request.Updates) {
		body["updates"] = request.Updates
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterDatabase"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AlterDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新接收者
//
// @param request - AlterReceiverRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterReceiverResponse
func (client *Client) AlterReceiverWithContext(ctx context.Context, receiver *string, request *AlterReceiverRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterReceiverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.ReceiverName) {
		body["receiverName"] = request.ReceiverName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterReceiver"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/receivers/" + dara.PercentEncode(dara.StringValue(receiver))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AlterReceiverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新共享
//
// @param request - AlterShareRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterShareResponse
func (client *Client) AlterShareWithContext(ctx context.Context, share *string, request *AlterShareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.ShareName) {
		body["shareName"] = request.ShareName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterShare"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AlterShareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新共享中的接收者
//
// @param request - AlterShareReceiversRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterShareReceiversResponse
func (client *Client) AlterShareReceiversWithContext(ctx context.Context, share *string, request *AlterShareReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareReceiversResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddedReceivers) {
		body["addedReceivers"] = request.AddedReceivers
	}

	if !dara.IsNil(request.RemovedReceivers) {
		body["removedReceivers"] = request.RemovedReceivers
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterShareReceivers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share)) + "/receivers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AlterShareReceiversResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改共享资源
//
// @param request - AlterShareResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterShareResourcesResponse
func (client *Client) AlterShareResourcesWithContext(ctx context.Context, share *string, request *AlterShareResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CatalogId) {
		body["catalogId"] = request.CatalogId
	}

	if !dara.IsNil(request.ShareResourceList) {
		body["shareResourceList"] = request.ShareResourceList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterShareResources"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share)) + "/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AlterShareResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更改Table
//
// @param request - AlterTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterTableResponse
func (client *Client) AlterTableWithContext(ctx context.Context, catalogId *string, database *string, table *string, request *AlterTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Changes) {
		body["changes"] = request.Changes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &AlterTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量授权
//
// @param request - BatchGrantPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGrantPermissionsResponse
func (client *Client) BatchGrantPermissionsWithContext(ctx context.Context, catalogId *string, request *BatchGrantPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGrantPermissionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Permissions) {
		body["permissions"] = request.Permissions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGrantPermissions"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/permissions/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/batchgrant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量取消授权
//
// @param request - BatchRevokePermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRevokePermissionsResponse
func (client *Client) BatchRevokePermissionsWithContext(ctx context.Context, catalogId *string, request *BatchRevokePermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchRevokePermissionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Permissions) {
		body["permissions"] = request.Permissions
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRevokePermissions"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/permissions/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/batchrevoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据目录
//
// @param request - CreateCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCatalogResponse
func (client *Client) CreateCatalogWithContext(ctx context.Context, request *CreateCatalogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCatalogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsShared) {
		body["isShared"] = request.IsShared
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.ShareId) {
		body["shareId"] = request.ShareId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCatalog"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据库
//
// @param request - CreateDatabaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseResponse
func (client *Client) CreateDatabaseWithContext(ctx context.Context, catalogId *string, request *CreateDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabase"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建接收者
//
// @param request - CreateReceiverRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReceiverResponse
func (client *Client) CreateReceiverWithContext(ctx context.Context, request *CreateReceiverRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateReceiverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.ReceiverName) {
		body["receiverName"] = request.ReceiverName
	}

	if !dara.IsNil(request.ReceiverTenantId) {
		body["receiverTenantId"] = request.ReceiverTenantId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReceiver"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/receivers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateReceiverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建角色
//
// @param request - CreateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRoleResponse
func (client *Client) CreateRoleWithContext(ctx context.Context, request *CreateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.RoleName) {
		body["roleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRole"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建共享
//
// @param request - CreateShareRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShareResponse
func (client *Client) CreateShareWithContext(ctx context.Context, request *CreateShareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.ShareName) {
		body["shareName"] = request.ShareName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateShare"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateShareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建表
//
// @param request - CreateTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableResponse
func (client *Client) CreateTableWithContext(ctx context.Context, catalogId *string, database *string, request *CreateTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Identifier) {
		body["identifier"] = request.Identifier
	}

	if !dara.IsNil(request.Schema) {
		body["schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &CreateTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除角色
//
// @param request - DeleteRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithContext(ctx context.Context, request *DeleteRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RolePrincipal) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRole"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 DLF 开通地域
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/service/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据湖Catalog
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropCatalogResponse
func (client *Client) DropCatalogWithContext(ctx context.Context, catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropCatalogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropCatalog"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs/" + dara.PercentEncode(dara.StringValue(catalog))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DropCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据库
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropDatabaseResponse
func (client *Client) DropDatabaseWithContext(ctx context.Context, catalogId *string, database *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropDatabaseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropDatabase"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DropDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除接收者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropReceiverResponse
func (client *Client) DropReceiverWithContext(ctx context.Context, receiver *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropReceiverResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropReceiver"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/receivers/" + dara.PercentEncode(dara.StringValue(receiver))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DropReceiverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除共享
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropShareResponse
func (client *Client) DropShareWithContext(ctx context.Context, share *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropShareResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropShare"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DropShareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropTableResponse
func (client *Client) DropTableWithContext(ctx context.Context, catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &DropTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据湖Catalog
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogResponse
func (client *Client) GetCatalogWithContext(ctx context.Context, catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalog"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs/" + dara.PercentEncode(dara.StringValue(catalog))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据湖Catalog
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogByIdResponse
func (client *Client) GetCatalogByIdWithContext(ctx context.Context, id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogByIdResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogById"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs/id/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogByIdResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetCatalogSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogSummaryResponse
func (client *Client) GetCatalogSummaryWithContext(ctx context.Context, catalogId *string, request *GetCatalogSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogSummary"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/storage-summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetCatalogSummaryTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogSummaryTrendResponse
func (client *Client) GetCatalogSummaryTrendWithContext(ctx context.Context, catalogId *string, request *GetCatalogSummaryTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogSummaryTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["endDate"] = request.EndDate
	}

	if !dara.IsNil(request.StartDate) {
		query["startDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogSummaryTrend"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/storage-summary/trend"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogSummaryTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据湖Catalog的临时访问凭证
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCatalogTokenResponse
func (client *Client) GetCatalogTokenWithContext(ctx context.Context, catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogTokenResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCatalogToken"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs/" + dara.PercentEncode(dara.StringValue(catalog)) + "/token"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCatalogTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据库
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithContext(ctx context.Context, catalogId *string, database *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetDatabaseSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseSummaryResponse
func (client *Client) GetDatabaseSummaryWithContext(ctx context.Context, catalogId *string, database *string, request *GetDatabaseSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatabaseSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseSummary"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/storage-summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看iceberg数据库
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIcebergNamespaceResponse
func (client *Client) GetIcebergNamespaceWithContext(ctx context.Context, catalogId *string, namespace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIcebergNamespaceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIcebergNamespace"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/iceberg/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/namespaces/" + dara.PercentEncode(dara.StringValue(namespace))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIcebergNamespaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIcebergTableResponse
func (client *Client) GetIcebergTableWithContext(ctx context.Context, catalogId *string, namespace *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIcebergTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIcebergTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/iceberg/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/tables/" + dara.PercentEncode(dara.StringValue(table))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIcebergTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接收者
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReceiverResponse
func (client *Client) GetReceiverWithContext(ctx context.Context, receiver *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReceiverResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReceiver"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/receivers/" + dara.PercentEncode(dara.StringValue(receiver))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReceiverResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 DLF 当前地域开通状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionStatusResponse
func (client *Client) GetRegionStatusWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegionStatus"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/service/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色
//
// @param request - GetRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRoleResponse
func (client *Client) GetRoleWithContext(ctx context.Context, request *GetRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RolePrincipal) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRole"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取共享
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareResponse
func (client *Client) GetShareWithContext(ctx context.Context, share *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetShareResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetShare"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetShareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableResponse
func (client *Client) GetTableWithContext(ctx context.Context, catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表快照
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableSnapshotResponse
func (client *Client) GetTableSnapshotWithContext(ctx context.Context, catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableSnapshotResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableSnapshot"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/snapshot"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableSnapshotResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - GetTableSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableSummaryResponse
func (client *Client) GetTableSummaryWithContext(ctx context.Context, catalogId *string, database *string, table *string, request *GetTableSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["date"] = request.Date
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableSummary"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/storage-summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithContext(ctx context.Context, request *GetUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserPrincipal) {
		query["userPrincipal"] = request.UserPrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/users"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量授予角色权限给用户
//
// @param request - GrantRoleToUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsersWithContext(ctx context.Context, request *GrantRoleToUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RolePrincipal) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !dara.IsNil(request.UserPrincipals) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantRoleToUsers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles/grant"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据目录列表
//
// @param request - ListCatalogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCatalogsResponse
func (client *Client) ListCatalogsWithContext(ctx context.Context, request *ListCatalogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogNamePattern) {
		query["catalogNamePattern"] = request.CatalogNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCatalogs"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/catalogs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据库列表
//
// @param request - ListDatabaseDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseDetailsResponse
func (client *Client) ListDatabaseDetailsWithContext(ctx context.Context, catalogId *string, request *ListDatabaseDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabaseDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseNamePattern) {
		query["databaseNamePattern"] = request.DatabaseNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseDetails"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/database-details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看数据库列表
//
// @param request - ListDatabasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithContext(ctx context.Context, catalogId *string, request *ListDatabasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseNamePattern) {
		query["databaseNamePattern"] = request.DatabaseNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看iceberg数据库列表
//
// @param request - ListIcebergNamespaceDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIcebergNamespaceDetailsResponse
func (client *Client) ListIcebergNamespaceDetailsWithContext(ctx context.Context, catalogId *string, request *ListIcebergNamespaceDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergNamespaceDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NamespaceNamePattern) {
		query["namespaceNamePattern"] = request.NamespaceNamePattern
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIcebergNamespaceDetails"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/iceberg/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/namespace-details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIcebergNamespaceDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看iceberg表快照列表
//
// @param request - ListIcebergSnapshotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIcebergSnapshotsResponse
func (client *Client) ListIcebergSnapshotsWithContext(ctx context.Context, catalogId *string, namespace *string, table *string, request *ListIcebergSnapshotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergSnapshotsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIcebergSnapshots"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/iceberg/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/snapshots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIcebergSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看iceberg表详情列表
//
// @param request - ListIcebergTableDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIcebergTableDetailsResponse
func (client *Client) ListIcebergTableDetailsWithContext(ctx context.Context, catalogId *string, namespace *string, request *ListIcebergTableDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergTableDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["tableNamePattern"] = request.TableNamePattern
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIcebergTableDetails"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/iceberg/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/namespaces/" + dara.PercentEncode(dara.StringValue(namespace)) + "/table-details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIcebergTableDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表
//
// @param request - ListPartitionSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPartitionSummariesResponse
func (client *Client) ListPartitionSummariesWithContext(ctx context.Context, catalogId *string, database *string, table *string, request *ListPartitionSummariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPartitionSummariesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.PartitionNamePattern) {
		query["partitionNamePattern"] = request.PartitionNamePattern
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPartitionSummaries"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/partitions/storage-summary"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPartitionSummariesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取指定资源或指定Principal的权限信息
//
// @param request - ListPermissionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionsResponse
func (client *Client) ListPermissionsWithContext(ctx context.Context, catalogId *string, request *ListPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Database) {
		query["database"] = request.Database
	}

	if !dara.IsNil(request.Function) {
		query["function"] = request.Function
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.Principal) {
		query["principal"] = request.Principal
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Table) {
		query["table"] = request.Table
	}

	if !dara.IsNil(request.View) {
		query["view"] = request.View
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissions"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/permissions/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取提供的共享列表
//
// @param request - ListProvidedSharesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProvidedSharesResponse
func (client *Client) ListProvidedSharesWithContext(ctx context.Context, request *ListProvidedSharesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProvidedSharesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProvidedShares"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/list/provided"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProvidedSharesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接收的共享列表
//
// @param request - ListReceivedSharesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReceivedSharesResponse
func (client *Client) ListReceivedSharesWithContext(ctx context.Context, request *ListReceivedSharesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReceivedSharesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReceivedShares"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/list/received"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReceivedSharesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取接收者列表
//
// @param request - ListReceiversRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReceiversResponse
func (client *Client) ListReceiversWithContext(ctx context.Context, request *ListReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReceiversResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.ReceiverName) {
		query["receiverName"] = request.ReceiverName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReceivers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/receivers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReceiversResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色用户列表
//
// @param request - ListRoleUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoleUsersResponse
func (client *Client) ListRoleUsersWithContext(ctx context.Context, request *ListRoleUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRoleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.RolePrincipal) {
		query["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoleUsers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles/users/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoleUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取角色列表
//
// @param request - ListRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRolesResponse
func (client *Client) ListRolesWithContext(ctx context.Context, request *ListRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.RoleName) {
		query["roleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRoles"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取共享中的接收者列表
//
// @param request - ListShareReceiversRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShareReceiversResponse
func (client *Client) ListShareReceiversWithContext(ctx context.Context, share *string, request *ListShareReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShareReceiversResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShareReceivers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share)) + "/receivers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShareReceiversResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取共享资源列表
//
// @param request - ListShareResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShareResourcesResponse
func (client *Client) ListShareResourcesWithContext(ctx context.Context, share *string, request *ListShareResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShareResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShareResources"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/share/shares/" + dara.PercentEncode(dara.StringValue(share)) + "/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShareResourcesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表快照列表
//
// @param request - ListSnapshotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithContext(ctx context.Context, catalogId *string, database *string, table *string, request *ListSnapshotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshots"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/snapshots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表详情列表
//
// @param request - ListTableDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableDetailsResponse
func (client *Client) ListTableDetailsWithContext(ctx context.Context, catalogId *string, database *string, request *ListTableDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTableDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["tableNamePattern"] = request.TableNamePattern
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableDetails"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/table-details"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableDetailsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看表详情列表
//
// @param request - ListTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithContext(ctx context.Context, catalogId *string, database *string, request *ListTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["tableNamePattern"] = request.TableNamePattern
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户角色列表
//
// @param request - ListUserRolesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserRolesResponse
func (client *Client) ListUserRolesWithContext(ctx context.Context, request *ListUserRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.UserPrincipal) {
		query["userPrincipal"] = request.UserPrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserRoles"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/users/roles/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserRolesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户列表
//
// @param request - ListUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithContext(ctx context.Context, request *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.PageToken) {
		query["pageToken"] = request.PageToken
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.UserName) {
		query["userName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/users/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 刷新用户同步
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshUserSyncResponse
func (client *Client) RefreshUserSyncWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RefreshUserSyncResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshUserSync"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/usersync"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RefreshUserSyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量取消授予角色权限给用户
//
// @param request - RevokeRoleFromUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeRoleFromUsersResponse
func (client *Client) RevokeRoleFromUsersWithContext(ctx context.Context, request *RevokeRoleFromUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeRoleFromUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RolePrincipal) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !dara.IsNil(request.UserPrincipals) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeRoleFromUsers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles/revoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回滚表
//
// @param request - RollbackTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackTableResponse
func (client *Client) RollbackTableWithContext(ctx context.Context, catalogId *string, database *string, table *string, request *RollbackTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Instant) {
		body["instant"] = request.Instant
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackTable"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/" + dara.PercentEncode(dara.StringValue(catalogId)) + "/databases/" + dara.PercentEncode(dara.StringValue(database)) + "/tables/" + dara.PercentEncode(dara.StringValue(table)) + "/rollback"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &RollbackTableResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 订阅当前地域的 DLF
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeResponse
func (client *Client) SubscribeWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubscribeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("Subscribe"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/service/subscribe"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &SubscribeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新角色
//
// @param request - UpdateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithContext(ctx context.Context, request *UpdateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.RolePrincipal) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRole"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新角色用户
//
// @param request - UpdateRoleUsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleUsersResponse
func (client *Client) UpdateRoleUsersWithContext(ctx context.Context, request *UpdateRoleUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RolePrincipal) {
		body["rolePrincipal"] = request.RolePrincipal
	}

	if !dara.IsNil(request.UserPrincipals) {
		body["userPrincipals"] = request.UserPrincipals
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRoleUsers"),
		Version:     dara.String("2025-03-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/dlf/v1/auth/roles/users"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("none"),
	}
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
