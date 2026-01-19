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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          dara.String("acc.aliyuncs.com"),
		"ap-south-1":                  dara.String("acc.aliyuncs.com"),
		"ap-southeast-2":              dara.String("acc.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("acc.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("acc.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("acc.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("acc.aliyuncs.com"),
		"cn-edge-1":                   dara.String("acc.aliyuncs.com"),
		"cn-fujian":                   dara.String("acc.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("acc.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("acc.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("acc.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("acc.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("acc.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("acc.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("acc.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("acc.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("acc.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("acc.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("acc.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("acc.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("acc.aliyuncs.com"),
		"cn-wuhan":                    dara.String("acc.aliyuncs.com"),
		"cn-yushanfang":               dara.String("acc.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("acc.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("acc.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("acc.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("acc.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("acc.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("acc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("acc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建镜像缓存
//
// @param tmpReq - CreateImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageCacheResponse
func (client *Client) CreateImageCacheWithOptions(tmpReq *CreateImageCacheRequest, runtime *dara.RuntimeOptions) (_result *CreateImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateImageCacheShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetworkConfig) {
		request.NetworkConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkConfig, dara.String("NetworkConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcrRegistryInfos) {
		query["AcrRegistryInfos"] = request.AcrRegistryInfos
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ImageCacheName) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !dara.IsNil(request.ImageRegistryCredentials) {
		query["ImageRegistryCredentials"] = request.ImageRegistryCredentials
	}

	if !dara.IsNil(request.Images) {
		query["Images"] = request.Images
	}

	if !dara.IsNil(request.NetworkConfigShrink) {
		query["NetworkConfig"] = request.NetworkConfigShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageCache"),
		Version:     dara.String("2024-04-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建镜像缓存
//
// @param request - CreateImageCacheRequest
//
// @return CreateImageCacheResponse
func (client *Client) CreateImageCache(request *CreateImageCacheRequest) (_result *CreateImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImageCacheResponse{}
	_body, _err := client.CreateImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除镜像缓存
//
// @param request - DeleteImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteImageCacheResponse
func (client *Client) DeleteImageCacheWithOptions(request *DeleteImageCacheRequest, runtime *dara.RuntimeOptions) (_result *DeleteImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.ImageCacheId) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteImageCache"),
		Version:     dara.String("2024-04-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除镜像缓存
//
// @param request - DeleteImageCacheRequest
//
// @return DeleteImageCacheResponse
func (client *Client) DeleteImageCache(request *DeleteImageCacheRequest) (_result *DeleteImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteImageCacheResponse{}
	_body, _err := client.DeleteImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询镜像缓存
//
// @param request - GetImageCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageCacheResponse
func (client *Client) GetImageCacheWithOptions(request *GetImageCacheRequest, runtime *dara.RuntimeOptions) (_result *GetImageCacheResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageCacheId) {
		query["ImageCacheId"] = request.ImageCacheId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageCache"),
		Version:     dara.String("2024-04-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询镜像缓存
//
// @param request - GetImageCacheRequest
//
// @return GetImageCacheResponse
func (client *Client) GetImageCache(request *GetImageCacheRequest) (_result *GetImageCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetImageCacheResponse{}
	_body, _err := client.GetImageCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询镜像缓存
//
// @param request - ListImageCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImageCachesResponse
func (client *Client) ListImageCachesWithOptions(request *ListImageCachesRequest, runtime *dara.RuntimeOptions) (_result *ListImageCachesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageCacheName) {
		query["ImageCacheName"] = request.ImageCacheName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImageCaches"),
		Version:     dara.String("2024-04-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListImageCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询镜像缓存
//
// @param request - ListImageCachesRequest
//
// @return ListImageCachesResponse
func (client *Client) ListImageCaches(request *ListImageCachesRequest) (_result *ListImageCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListImageCachesResponse{}
	_body, _err := client.ListImageCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
