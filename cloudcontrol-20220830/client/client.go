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
		"ap-northeast-1": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-6": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-7": dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"cn-beijing":     dara.String("cloudcontrol.aliyuncs.com"),
		"cn-chengdu":     dara.String("cloudcontrol.aliyuncs.com"),
		"cn-fuzhou":      dara.String("cloudcontrol.aliyuncs.com"),
		"cn-guangzhou":   dara.String("cloudcontrol.aliyuncs.com"),
		"cn-hangzhou":    dara.String("cloudcontrol.aliyuncs.com"),
		"cn-heyuan":      dara.String("cloudcontrol.aliyuncs.com"),
		"cn-hongkong":    dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"cn-huhehaote":   dara.String("cloudcontrol.aliyuncs.com"),
		"cn-nanjing":     dara.String("cloudcontrol.aliyuncs.com"),
		"cn-qingdao":     dara.String("cloudcontrol.aliyuncs.com"),
		"cn-shanghai":    dara.String("cloudcontrol.aliyuncs.com"),
		"cn-shenzhen":    dara.String("cloudcontrol.aliyuncs.com"),
		"cn-wulanchabu":  dara.String("cloudcontrol.aliyuncs.com"),
		"cn-zhangjiakou": dara.String("cloudcontrol.aliyuncs.com"),
		"us-west-1":      dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"us-east-1":      dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":      dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"me-east-1":      dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"me-central-1":   dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     dara.String("cloudcontrol.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudcontrol"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Calls this operation to cancel a specified asynchronous task.
//
// Description:
//
// Only tasks that are in the Pending or Running state can be canceled.
//
// You can call the CancelTask operation to cancel a Cloud Control API task, but the tasks that have been started in the downstream Alibaba Cloud services cannot be canceled.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelTaskResponse
func (client *Client) CancelTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelTask"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks/" + dara.PercentEncode(dara.StringValue(taskId)) + "/operation/cancel"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to cancel a specified asynchronous task.
//
// Description:
//
// Only tasks that are in the Pending or Running state can be canceled.
//
// You can call the CancelTask operation to cancel a Cloud Control API task, but the tasks that have been started in the downstream Alibaba Cloud services cannot be canceled.
//
// @return CancelTaskResponse
func (client *Client) CancelTask(taskId *string) (_result *CancelTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls this operation to create resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param request - CreateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(requestPath *string, request *CreateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to create resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(requestPath *string, request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view resource documentation and try Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - DeleteResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(requestPath *string, tmpReq *DeleteResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view resource documentation and try Cloud Control API.
//
// @param requestPath - the whole path of resource string
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(requestPath *string, request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries pricing based on an OpenAPI triplet and input parameters.
//
// @param request - GetApiPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiPriceResponse
func (client *Client) GetApiPriceWithOptions(request *GetApiPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiPrice"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/price/quote"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries pricing based on an OpenAPI triplet and input parameters.
//
// @param request - GetApiPriceRequest
//
// @return GetApiPriceResponse
func (client *Client) GetApiPrice(request *GetApiPriceRequest) (_result *GetApiPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApiPriceResponse{}
	_body, _err := client.GetApiPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// An RFQ interface through which users can query resource prices.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPriceResponse
func (client *Client) GetPriceWithOptions(requestPath *string, tmpReq *GetPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetPriceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceAttributes) {
		request.ResourceAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceAttributes, dara.String("resourceAttributes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceAttributesShrink) {
		query["resourceAttributes"] = request.ResourceAttributesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrice"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An RFQ interface through which users can query resource prices.
//
// @param requestPath - the whole path of resource string
//
// @param request - GetPriceRequest
//
// @return GetPriceResponse
func (client *Client) GetPrice(requestPath *string, request *GetPriceRequest) (_result *GetPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPriceResponse{}
	_body, _err := client.GetPriceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves resource metadata.
//
// @param requestPath - the whole path of resource string
//
// @param headers - GetResourceTypeHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceTypeWithOptions(requestPath *string, headers *GetResourceTypeHeaders, runtime *dara.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceType"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves resource metadata.
//
// @param requestPath - the whole path of resource string
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceType(requestPath *string) (_result *GetResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GetResourceTypeHeaders{}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(requestPath, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query resources.
//
// Description:
//
// You can go to the [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the resource documentation and test the Cloud Control API.
//
// This API provides Get and List operations for resources that you can invoke using different request URIs.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - GetResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcesResponse
func (client *Client) GetResourcesWithOptions(requestPath *string, tmpReq *GetResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResources"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query resources.
//
// Description:
//
// You can go to the [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the resource documentation and test the Cloud Control API.
//
// This API provides Get and List operations for resources that you can invoke using different request URIs.
//
// @param requestPath - the whole path of resource string
//
// @param request - GetResourcesRequest
//
// @return GetResourcesResponse
func (client *Client) GetResources(requestPath *string, request *GetResourcesRequest) (_result *GetResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourcesResponse{}
	_body, _err := client.GetResourcesWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls this operation to query a specified asynchronous task.
//
// Description:
//
// GET /api/v1/tasks/{taskId}.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tasks/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to query a specified asynchronous task.
//
// Description:
//
// GET /api/v1/tasks/{taskId}.
//
// @return GetTaskResponse
func (client *Client) GetTask(taskId *string) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves pricing mapping catalogs in batches by Terraform resource type for cost estimation during the RunIaC plan phase.
//
// Description:
//
// Retrieves the mappings between schema properties in the Terraform alicloud provider and OpenAPI parameters.
//
// @param request - GetTerraformPricingMappingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTerraformPricingMappingsResponse
func (client *Client) GetTerraformPricingMappingsWithOptions(request *GetTerraformPricingMappingsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTerraformPricingMappingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTerraformPricingMappings"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/price/terraform-mappings"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTerraformPricingMappingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves pricing mapping catalogs in batches by Terraform resource type for cost estimation during the RunIaC plan phase.
//
// Description:
//
// Retrieves the mappings between schema properties in the Terraform alicloud provider and OpenAPI parameters.
//
// @param request - GetTerraformPricingMappingsRequest
//
// @return GetTerraformPricingMappingsResponse
func (client *Client) GetTerraformPricingMappings(request *GetTerraformPricingMappingsRequest) (_result *GetTerraformPricingMappingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTerraformPricingMappingsResponse{}
	_body, _err := client.GetTerraformPricingMappingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the valid values of resource attributes, such as RegionID and ZoneId.
//
// @param requestPath - the whole path of resource string
//
// @param tmpReq - ListDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(requestPath *string, tmpReq *ListDataSourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataSourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttributeName) {
		query["attributeName"] = request.AttributeName
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the valid values of resource attributes, such as RegionID and ZoneId.
//
// @param requestPath - the whole path of resource string
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(requestPath *string, request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls this operation to list the supported services.
//
// Description:
//
// GET /api/v1/providers/{provider}/products.
//
// @param request - ListProductsRequest
//
// @param headers - ListProductsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(provider *string, request *ListProductsRequest, headers *ListProductsHeaders, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/providers/" + dara.PercentEncode(dara.StringValue(provider)) + "/products"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to list the supported services.
//
// Description:
//
// GET /api/v1/providers/{provider}/products.
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(provider *string, request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListProductsHeaders{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(provider, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls this operation to list the resource types of a service.
//
// Description:
//
// GET /api/v1/providers/{provider}/products/{product}/resourceTypes.
//
// @param tmpReq - ListResourceTypesRequest
//
// @param headers - ListResourceTypesHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithOptions(provider *string, product *string, tmpReq *ListResourceTypesRequest, headers *ListResourceTypesHeaders, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListResourceTypesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceTypes) {
		request.ResourceTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypes, dara.String("resourceTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceTypesShrink) {
		query["resourceTypes"] = request.ResourceTypesShrink
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XAcsAcceptLanguage) {
		realHeaders["x-acs-accept-language"] = dara.String(dara.ToString(dara.StringValue(headers.XAcsAcceptLanguage)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypes"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/providers/" + dara.PercentEncode(dara.StringValue(provider)) + "/products/" + dara.PercentEncode(dara.StringValue(product)) + "/resourceTypes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to list the resource types of a service.
//
// Description:
//
// GET /api/v1/providers/{provider}/products/{product}/resourceTypes.
//
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(provider *string, product *string, request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &ListResourceTypesHeaders{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(provider, product, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the OpenAPI triplets that currently support price inquiry.
//
// @param request - ListSupportedPricingApisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSupportedPricingApisResponse
func (client *Client) ListSupportedPricingApisWithOptions(request *ListSupportedPricingApisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSupportedPricingApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSupportedPricingApis"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/price/supported-apis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSupportedPricingApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the OpenAPI triplets that currently support price inquiry.
//
// @param request - ListSupportedPricingApisRequest
//
// @return ListSupportedPricingApisResponse
func (client *Client) ListSupportedPricingApis(request *ListSupportedPricingApisRequest) (_result *ListSupportedPricingApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSupportedPricingApisResponse{}
	_body, _err := client.ListSupportedPricingApisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls this operation to update resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// If resources fail to be updated at any time, the Cloud Control API does not roll the resource back to the original status.
//
// The resource APIs cannot be rolled back. If the API operation is partially failed to be called, you can call the GetResource operation to view the latest status of the resource. If necessary, you can call the UpdateResource or DeleteResource operation to manually compensate for the failure.
//
// @param requestPath - the whole path of resource string
//
// @param request - UpdateResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(requestPath *string, request *UpdateResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateResource"),
		Version:     dara.String("2022-08-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String(dara.StringValue(requestPath)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls this operation to update resources.
//
// Description:
//
// You can go to [OpenAPI Explorer](https://next.api.aliyun.com/cloudcontrol) to view the documentation and try out Cloud Control API.
//
// If resources fail to be updated at any time, the Cloud Control API does not roll the resource back to the original status.
//
// The resource APIs cannot be rolled back. If the API operation is partially failed to be called, you can call the GetResource operation to view the latest status of the resource. If necessary, you can call the UpdateResource or DeleteResource operation to manually compensate for the failure.
//
// @param requestPath - the whole path of resource string
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(requestPath *string, request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(requestPath, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
