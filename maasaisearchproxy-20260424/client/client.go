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
	client.Endpoint, _err = client.GetEndpoint(dara.String("maasaisearchproxy"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 提交单条记录导入任务（通过AccessKey认证）
//
// @param request - OpenDatasetImportDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDatasetImportDataResponse
func (client *Client) OpenDatasetImportDataWithOptions(request *OpenDatasetImportDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenDatasetImportDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDatasetImportData"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/upsert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDatasetImportDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交单条记录导入任务（通过AccessKey认证）
//
// @param request - OpenDatasetImportDataRequest
//
// @return OpenDatasetImportDataResponse
func (client *Client) OpenDatasetImportData(request *OpenDatasetImportDataRequest) (_result *OpenDatasetImportDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenDatasetImportDataResponse{}
	_body, _err := client.OpenDatasetImportDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取数据集资源 OSS 访问地址（通过AccessKey认证）
//
// @param request - OpenDatasetResourceUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenDatasetResourceUrlResponse
func (client *Client) OpenDatasetResourceUrlWithOptions(request *OpenDatasetResourceUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OpenDatasetResourceUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.PrimaryKey) {
		body["primaryKey"] = request.PrimaryKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenDatasetResourceUrl"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenDatasetResourceUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据集资源 OSS 访问地址（通过AccessKey认证）
//
// @param request - OpenDatasetResourceUrlRequest
//
// @return OpenDatasetResourceUrlResponse
func (client *Client) OpenDatasetResourceUrl(request *OpenDatasetResourceUrlRequest) (_result *OpenDatasetResourceUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenDatasetResourceUrlResponse{}
	_body, _err := client.OpenDatasetResourceUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 联网搜API
//
// @param request - WebSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebSearchResponse
func (client *Client) WebSearchWithOptions(request *WebSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *WebSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExcludeDomain) {
		body["excludeDomain"] = request.ExcludeDomain
	}

	if !dara.IsNil(request.IncludeDomain) {
		body["includeDomain"] = request.IncludeDomain
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Region) {
		body["region"] = request.Region
	}

	if !dara.IsNil(request.SearchType) {
		body["searchType"] = request.SearchType
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("WebSearch"),
		Version:     dara.String("2026-04-24"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/web-search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &WebSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 联网搜API
//
// @param request - WebSearchRequest
//
// @return WebSearchResponse
func (client *Client) WebSearch(request *WebSearchRequest) (_result *WebSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebSearchResponse{}
	_body, _err := client.WebSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
