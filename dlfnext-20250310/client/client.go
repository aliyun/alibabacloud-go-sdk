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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dlfnext"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 更新数据目录
//
// @param request - AlterCatalogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterCatalogResponse
func (client *Client) AlterCatalogWithOptions(catalog *string, request *AlterCatalogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterCatalogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据目录
//
// @param request - AlterCatalogRequest
//
// @return AlterCatalogResponse
func (client *Client) AlterCatalog(catalog *string, request *AlterCatalogRequest) (_result *AlterCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterCatalogResponse{}
	_body, _err := client.AlterCatalogWithOptions(catalog, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterDatabaseWithOptions(catalogId *string, database *string, request *AlterDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterDatabaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterDatabaseResponse
func (client *Client) AlterDatabase(catalogId *string, database *string, request *AlterDatabaseRequest) (_result *AlterDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterDatabaseResponse{}
	_body, _err := client.AlterDatabaseWithOptions(catalogId, database, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterReceiverWithOptions(receiver *string, request *AlterReceiverRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterReceiverResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterReceiverResponse
func (client *Client) AlterReceiver(receiver *string, request *AlterReceiverRequest) (_result *AlterReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterReceiverResponse{}
	_body, _err := client.AlterReceiverWithOptions(receiver, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterShareWithOptions(share *string, request *AlterShareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.EnableWrite) {
		body["enableWrite"] = request.EnableWrite
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterShareResponse
func (client *Client) AlterShare(share *string, request *AlterShareRequest) (_result *AlterShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterShareResponse{}
	_body, _err := client.AlterShareWithOptions(share, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterShareReceiversWithOptions(share *string, request *AlterShareReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareReceiversResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterShareReceiversResponse
func (client *Client) AlterShareReceivers(share *string, request *AlterShareReceiversRequest) (_result *AlterShareReceiversResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterShareReceiversResponse{}
	_body, _err := client.AlterShareReceiversWithOptions(share, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterShareResourcesWithOptions(share *string, request *AlterShareResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterShareResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterShareResourcesResponse
func (client *Client) AlterShareResources(share *string, request *AlterShareResourcesRequest) (_result *AlterShareResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterShareResourcesResponse{}
	_body, _err := client.AlterShareResourcesWithOptions(share, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AlterTableWithOptions(catalogId *string, database *string, table *string, request *AlterTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AlterTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AlterTableResponse
func (client *Client) AlterTable(catalogId *string, database *string, table *string, request *AlterTableRequest) (_result *AlterTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AlterTableResponse{}
	_body, _err := client.AlterTableWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchGrantPermissionsWithOptions(catalogId *string, request *BatchGrantPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGrantPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchGrantPermissionsResponse
func (client *Client) BatchGrantPermissions(catalogId *string, request *BatchGrantPermissionsRequest) (_result *BatchGrantPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.BatchGrantPermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchRevokePermissionsWithOptions(catalogId *string, request *BatchRevokePermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchRevokePermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchRevokePermissionsResponse
func (client *Client) BatchRevokePermissions(catalogId *string, request *BatchRevokePermissionsRequest) (_result *BatchRevokePermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.BatchRevokePermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCatalogWithOptions(request *CreateCatalogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCatalogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCatalogResponse
func (client *Client) CreateCatalog(request *CreateCatalogRequest) (_result *CreateCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCatalogResponse{}
	_body, _err := client.CreateCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDatabaseWithOptions(catalogId *string, request *CreateDatabaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDatabaseResponse
func (client *Client) CreateDatabase(catalogId *string, request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateReceiverWithOptions(request *CreateReceiverRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateReceiverResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateReceiverResponse
func (client *Client) CreateReceiver(request *CreateReceiverRequest) (_result *CreateReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateReceiverResponse{}
	_body, _err := client.CreateReceiverWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRoleResponse
func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateShareWithOptions(request *CreateShareRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["comment"] = request.Comment
	}

	if !dara.IsNil(request.EnableWrite) {
		body["enableWrite"] = request.EnableWrite
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateShareResponse
func (client *Client) CreateShare(request *CreateShareRequest) (_result *CreateShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateShareResponse{}
	_body, _err := client.CreateShareWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTableWithOptions(catalogId *string, database *string, request *CreateTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTableResponse
func (client *Client) CreateTable(catalogId *string, database *string, request *CreateTableRequest) (_result *CreateTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTableResponse{}
	_body, _err := client.CreateTableWithOptions(catalogId, database, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRoleResponse
func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DropCatalogWithOptions(catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropCatalogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DropCatalogResponse
func (client *Client) DropCatalog(catalog *string) (_result *DropCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropCatalogResponse{}
	_body, _err := client.DropCatalogWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DropDatabaseWithOptions(catalogId *string, database *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropDatabaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DropDatabaseResponse
func (client *Client) DropDatabase(catalogId *string, database *string) (_result *DropDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropDatabaseResponse{}
	_body, _err := client.DropDatabaseWithOptions(catalogId, database, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DropReceiverWithOptions(receiver *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropReceiverResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DropReceiverResponse
func (client *Client) DropReceiver(receiver *string) (_result *DropReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropReceiverResponse{}
	_body, _err := client.DropReceiverWithOptions(receiver, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DropShareWithOptions(share *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropShareResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DropShareResponse
func (client *Client) DropShare(share *string) (_result *DropShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropShareResponse{}
	_body, _err := client.DropShareWithOptions(share, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DropTableWithOptions(catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DropTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DropTableResponse
func (client *Client) DropTable(catalogId *string, database *string, table *string) (_result *DropTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DropTableResponse{}
	_body, _err := client.DropTableWithOptions(catalogId, database, table, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCatalogWithOptions(catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCatalogResponse
func (client *Client) GetCatalog(catalog *string) (_result *GetCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogResponse{}
	_body, _err := client.GetCatalogWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCatalogByIdWithOptions(id *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogByIdResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCatalogByIdResponse
func (client *Client) GetCatalogById(id *string) (_result *GetCatalogByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogByIdResponse{}
	_body, _err := client.GetCatalogByIdWithOptions(id, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCatalogSummaryWithOptions(catalogId *string, request *GetCatalogSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCatalogSummaryResponse
func (client *Client) GetCatalogSummary(catalogId *string, request *GetCatalogSummaryRequest) (_result *GetCatalogSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogSummaryResponse{}
	_body, _err := client.GetCatalogSummaryWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCatalogSummaryTrendWithOptions(catalogId *string, request *GetCatalogSummaryTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogSummaryTrendResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCatalogSummaryTrendResponse
func (client *Client) GetCatalogSummaryTrend(catalogId *string, request *GetCatalogSummaryTrendRequest) (_result *GetCatalogSummaryTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogSummaryTrendResponse{}
	_body, _err := client.GetCatalogSummaryTrendWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCatalogTokenWithOptions(catalog *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCatalogTokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCatalogTokenResponse
func (client *Client) GetCatalogToken(catalog *string) (_result *GetCatalogTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogTokenResponse{}
	_body, _err := client.GetCatalogTokenWithOptions(catalog, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatabaseWithOptions(catalogId *string, database *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatabaseResponse
func (client *Client) GetDatabase(catalogId *string, database *string) (_result *GetDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(catalogId, database, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatabaseSummaryWithOptions(catalogId *string, database *string, request *GetDatabaseSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatabaseSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatabaseSummaryResponse
func (client *Client) GetDatabaseSummary(catalogId *string, database *string, request *GetDatabaseSummaryRequest) (_result *GetDatabaseSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatabaseSummaryResponse{}
	_body, _err := client.GetDatabaseSummaryWithOptions(catalogId, database, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIcebergNamespaceWithOptions(catalogId *string, namespace *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIcebergNamespaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIcebergNamespaceResponse
func (client *Client) GetIcebergNamespace(catalogId *string, namespace *string) (_result *GetIcebergNamespaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIcebergNamespaceResponse{}
	_body, _err := client.GetIcebergNamespaceWithOptions(catalogId, namespace, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIcebergTableWithOptions(catalogId *string, namespace *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIcebergTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIcebergTableResponse
func (client *Client) GetIcebergTable(catalogId *string, namespace *string, table *string) (_result *GetIcebergTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIcebergTableResponse{}
	_body, _err := client.GetIcebergTableWithOptions(catalogId, namespace, table, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetReceiverWithOptions(receiver *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReceiverResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetReceiverResponse
func (client *Client) GetReceiver(receiver *string) (_result *GetReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReceiverResponse{}
	_body, _err := client.GetReceiverWithOptions(receiver, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRegionStatusWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRegionStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRegionStatusResponse
func (client *Client) GetRegionStatus() (_result *GetRegionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionStatusResponse{}
	_body, _err := client.GetRegionStatusWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetRoleWithOptions(request *GetRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetRoleResponse
func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetShareWithOptions(share *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetShareResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetShareResponse
func (client *Client) GetShare(share *string) (_result *GetShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShareResponse{}
	_body, _err := client.GetShareWithOptions(share, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTableWithOptions(catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTableResponse
func (client *Client) GetTable(catalogId *string, database *string, table *string) (_result *GetTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(catalogId, database, table, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTableSnapshotWithOptions(catalogId *string, database *string, table *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableSnapshotResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTableSnapshotResponse
func (client *Client) GetTableSnapshot(catalogId *string, database *string, table *string) (_result *GetTableSnapshotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableSnapshotResponse{}
	_body, _err := client.GetTableSnapshotWithOptions(catalogId, database, table, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTableSummaryWithOptions(catalogId *string, database *string, table *string, request *GetTableSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTableSummaryResponse
func (client *Client) GetTableSummary(catalogId *string, database *string, table *string, request *GetTableSummaryRequest) (_result *GetTableSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableSummaryResponse{}
	_body, _err := client.GetTableSummaryWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GrantRoleToUsersWithOptions(request *GrantRoleToUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GrantRoleToUsersResponse
func (client *Client) GrantRoleToUsers(request *GrantRoleToUsersRequest) (_result *GrantRoleToUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.GrantRoleToUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListCatalogsWithOptions(request *ListCatalogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListCatalogsResponse
func (client *Client) ListCatalogs(request *ListCatalogsRequest) (_result *ListCatalogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCatalogsResponse{}
	_body, _err := client.ListCatalogsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatabaseDetailsWithOptions(catalogId *string, request *ListDatabaseDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabaseDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatabaseDetailsResponse
func (client *Client) ListDatabaseDetails(catalogId *string, request *ListDatabaseDetailsRequest) (_result *ListDatabaseDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatabaseDetailsResponse{}
	_body, _err := client.ListDatabaseDetailsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatabasesWithOptions(catalogId *string, request *ListDatabasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatabasesResponse
func (client *Client) ListDatabases(catalogId *string, request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIcebergNamespaceDetailsWithOptions(catalogId *string, request *ListIcebergNamespaceDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergNamespaceDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIcebergNamespaceDetailsResponse
func (client *Client) ListIcebergNamespaceDetails(catalogId *string, request *ListIcebergNamespaceDetailsRequest) (_result *ListIcebergNamespaceDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIcebergNamespaceDetailsResponse{}
	_body, _err := client.ListIcebergNamespaceDetailsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIcebergSnapshotsWithOptions(catalogId *string, namespace *string, table *string, request *ListIcebergSnapshotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergSnapshotsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIcebergSnapshotsResponse
func (client *Client) ListIcebergSnapshots(catalogId *string, namespace *string, table *string, request *ListIcebergSnapshotsRequest) (_result *ListIcebergSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIcebergSnapshotsResponse{}
	_body, _err := client.ListIcebergSnapshotsWithOptions(catalogId, namespace, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListIcebergTableDetailsWithOptions(catalogId *string, namespace *string, request *ListIcebergTableDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIcebergTableDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListIcebergTableDetailsResponse
func (client *Client) ListIcebergTableDetails(catalogId *string, namespace *string, request *ListIcebergTableDetailsRequest) (_result *ListIcebergTableDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIcebergTableDetailsResponse{}
	_body, _err := client.ListIcebergTableDetailsWithOptions(catalogId, namespace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPartitionSummariesWithOptions(catalogId *string, database *string, table *string, request *ListPartitionSummariesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPartitionSummariesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPartitionSummariesResponse
func (client *Client) ListPartitionSummaries(catalogId *string, database *string, table *string, request *ListPartitionSummariesRequest) (_result *ListPartitionSummariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionSummariesResponse{}
	_body, _err := client.ListPartitionSummariesWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPermissionsWithOptions(catalogId *string, request *ListPermissionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPermissionsResponse
func (client *Client) ListPermissions(catalogId *string, request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(catalogId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListProvidedSharesWithOptions(request *ListProvidedSharesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListProvidedSharesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListProvidedSharesResponse
func (client *Client) ListProvidedShares(request *ListProvidedSharesRequest) (_result *ListProvidedSharesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProvidedSharesResponse{}
	_body, _err := client.ListProvidedSharesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListReceivedSharesWithOptions(request *ListReceivedSharesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReceivedSharesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListReceivedSharesResponse
func (client *Client) ListReceivedShares(request *ListReceivedSharesRequest) (_result *ListReceivedSharesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListReceivedSharesResponse{}
	_body, _err := client.ListReceivedSharesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListReceiversWithOptions(request *ListReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListReceiversResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListReceiversResponse
func (client *Client) ListReceivers(request *ListReceiversRequest) (_result *ListReceiversResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListReceiversResponse{}
	_body, _err := client.ListReceiversWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRoleUsersWithOptions(request *ListRoleUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRoleUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRoleUsersResponse
func (client *Client) ListRoleUsers(request *ListRoleUsersRequest) (_result *ListRoleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRoleUsersResponse{}
	_body, _err := client.ListRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListRolesWithOptions(request *ListRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRolesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListRolesResponse
func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListShareReceiversWithOptions(share *string, request *ListShareReceiversRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShareReceiversResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListShareReceiversResponse
func (client *Client) ListShareReceivers(share *string, request *ListShareReceiversRequest) (_result *ListShareReceiversResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShareReceiversResponse{}
	_body, _err := client.ListShareReceiversWithOptions(share, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListShareResourcesWithOptions(share *string, request *ListShareResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListShareResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListShareResourcesResponse
func (client *Client) ListShareResources(share *string, request *ListShareResourcesRequest) (_result *ListShareResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShareResourcesResponse{}
	_body, _err := client.ListShareResourcesWithOptions(share, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListSnapshotsWithOptions(catalogId *string, database *string, table *string, request *ListSnapshotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListSnapshotsResponse
func (client *Client) ListSnapshots(catalogId *string, database *string, table *string, request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTableDetailsWithOptions(catalogId *string, database *string, request *ListTableDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTableDetailsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTableDetailsResponse
func (client *Client) ListTableDetails(catalogId *string, database *string, request *ListTableDetailsRequest) (_result *ListTableDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableDetailsResponse{}
	_body, _err := client.ListTableDetailsWithOptions(catalogId, database, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTablesWithOptions(catalogId *string, database *string, request *ListTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTablesResponse
func (client *Client) ListTables(catalogId *string, database *string, request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(catalogId, database, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListUserRolesWithOptions(request *ListUserRolesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserRolesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListUserRolesResponse
func (client *Client) ListUserRoles(request *ListUserRolesRequest) (_result *ListUserRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserRolesResponse{}
	_body, _err := client.ListUserRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RefreshUserSyncWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RefreshUserSyncResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RefreshUserSyncResponse
func (client *Client) RefreshUserSync() (_result *RefreshUserSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshUserSyncResponse{}
	_body, _err := client.RefreshUserSyncWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RevokeRoleFromUsersWithOptions(request *RevokeRoleFromUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RevokeRoleFromUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RevokeRoleFromUsersResponse
func (client *Client) RevokeRoleFromUsers(request *RevokeRoleFromUsersRequest) (_result *RevokeRoleFromUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.RevokeRoleFromUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RollbackTableWithOptions(catalogId *string, database *string, table *string, request *RollbackTableRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RollbackTableResponse
func (client *Client) RollbackTable(catalogId *string, database *string, table *string, request *RollbackTableRequest) (_result *RollbackTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollbackTableResponse{}
	_body, _err := client.RollbackTableWithOptions(catalogId, database, table, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubscribeWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubscribeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubscribeResponse
func (client *Client) Subscribe() (_result *SubscribeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubscribeResponse{}
	_body, _err := client.SubscribeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRoleResponse
func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateRoleUsersWithOptions(request *UpdateRoleUsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleUsersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateRoleUsersResponse
func (client *Client) UpdateRoleUsers(request *UpdateRoleUsersRequest) (_result *UpdateRoleUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.UpdateRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
