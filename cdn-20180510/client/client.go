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
	client.EndpointRule = dara.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":     dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1": dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2": dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3": dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":   dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":      dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"me-east-1":      dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-east-1":      dara.String("cdn.ap-southeast-1.aliyuncs.com"),
		"us-west-1":      dara.String("cdn.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cdn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds a domain name to accelerate.
//
// Description:
//
//	  You must activate Alibaba Cloud CDN before you can add a domain name to it. For more information, see [Activate Alibaba Cloud CDN](https://help.aliyun.com/document_detail/27272.html).
//
//		- The domain name that you want to add has a valid Internet Content Provider (ICP) number.
//
//		- You can add only one domain name to Alibaba Cloud CDN in each call. Each Alibaba Cloud account can add a maximum of 50 domain names to Alibaba Cloud CDN.
//
//		- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. The review will be completed by the end of the next business day after you submit the application.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - AddCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCdnDomainResponse
func (client *Client) AddCdnDomainWithOptions(request *AddCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *AddCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdnType) {
		query["CdnType"] = request.CdnType
	}

	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to accelerate.
//
// Description:
//
//	  You must activate Alibaba Cloud CDN before you can add a domain name to it. For more information, see [Activate Alibaba Cloud CDN](https://help.aliyun.com/document_detail/27272.html).
//
//		- The domain name that you want to add has a valid Internet Content Provider (ICP) number.
//
//		- You can add only one domain name to Alibaba Cloud CDN in each call. Each Alibaba Cloud account can add a maximum of 50 domain names to Alibaba Cloud CDN.
//
//		- If the content of the origin server is not stored on Alibaba Cloud, the content must be reviewed. The review will be completed by the end of the next business day after you submit the application.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - AddCdnDomainRequest
//
// @return AddCdnDomainResponse
func (client *Client) AddCdnDomain(request *AddCdnDomainRequest) (_result *AddCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCdnDomainResponse{}
	_body, _err := client.AddCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a Function Compute trigger.
//
// @param request - AddFCTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFCTriggerResponse
func (client *Client) AddFCTriggerWithOptions(request *AddFCTriggerRequest, runtime *dara.RuntimeOptions) (_result *AddFCTriggerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TriggerARN) {
		query["TriggerARN"] = request.TriggerARN
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EventMetaName) {
		body["EventMetaName"] = request.EventMetaName
	}

	if !dara.IsNil(request.EventMetaVersion) {
		body["EventMetaVersion"] = request.EventMetaVersion
	}

	if !dara.IsNil(request.FunctionARN) {
		body["FunctionARN"] = request.FunctionARN
	}

	if !dara.IsNil(request.Notes) {
		body["Notes"] = request.Notes
	}

	if !dara.IsNil(request.RoleARN) {
		body["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.SourceARN) {
		body["SourceARN"] = request.SourceARN
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFCTrigger"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFCTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a Function Compute trigger.
//
// @param request - AddFCTriggerRequest
//
// @return AddFCTriggerResponse
func (client *Client) AddFCTrigger(request *AddFCTriggerRequest) (_result *AddFCTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFCTriggerResponse{}
	_body, _err := client.AddFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more domain names to Alibaba Cloud CDN. You can add a maximum of 50 domain names at a time.
//
// Description:
//
//	  You must activate Alibaba Cloud CDN before you can add a domain name to it. For more information, see [Activate Alibaba Cloud CDN](https://help.aliyun.com/document_detail/27272.html).
//
//		- If the acceleration region is Chinese Mainland Only or Global, you must apply for an ICP filing for the domain name.
//
//		- You can specify multiple domain names and separate them with commas (,). You can specify at most 50 domain names in each call.
//
//		- For more information, see [Add a domain name](https://help.aliyun.com/document_detail/122181.html).
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchAddCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddCdnDomainResponse
func (client *Client) BatchAddCdnDomainWithOptions(request *BatchAddCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchAddCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdnType) {
		query["CdnType"] = request.CdnType
	}

	if !dara.IsNil(request.CheckUrl) {
		query["CheckUrl"] = request.CheckUrl
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more domain names to Alibaba Cloud CDN. You can add a maximum of 50 domain names at a time.
//
// Description:
//
//	  You must activate Alibaba Cloud CDN before you can add a domain name to it. For more information, see [Activate Alibaba Cloud CDN](https://help.aliyun.com/document_detail/27272.html).
//
//		- If the acceleration region is Chinese Mainland Only or Global, you must apply for an ICP filing for the domain name.
//
//		- You can specify multiple domain names and separate them with commas (,). You can specify at most 50 domain names in each call.
//
//		- For more information, see [Add a domain name](https://help.aliyun.com/document_detail/122181.html).
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchAddCdnDomainRequest
//
// @return BatchAddCdnDomainResponse
func (client *Client) BatchAddCdnDomain(request *BatchAddCdnDomainRequest) (_result *BatchAddCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchAddCdnDomainResponse{}
	_body, _err := client.BatchAddCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes configurations of multiple accelerated domain names at a time.
//
// Description:
//
//	  You can specify up to 50 domain names in each request.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchDeleteCdnDomainConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteCdnDomainConfigResponse
func (client *Client) BatchDeleteCdnDomainConfigWithOptions(request *BatchDeleteCdnDomainConfigRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteCdnDomainConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteCdnDomainConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteCdnDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes configurations of multiple accelerated domain names at a time.
//
// Description:
//
//	  You can specify up to 50 domain names in each request.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - BatchDeleteCdnDomainConfigRequest
//
// @return BatchDeleteCdnDomainConfigResponse
func (client *Client) BatchDeleteCdnDomainConfig(request *BatchDeleteCdnDomainConfigRequest) (_result *BatchDeleteCdnDomainConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteCdnDomainConfigResponse{}
	_body, _err := client.BatchDeleteCdnDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether one or more IP addresses are assigned to Alibaba Cloud CDN.
//
// Description:
//
// >The maximum number of times that each user can call this operation per second is 20.
//
// @param request - BatchDescribeCdnIpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDescribeCdnIpInfoResponse
func (client *Client) BatchDescribeCdnIpInfoWithOptions(request *BatchDescribeCdnIpInfoRequest, runtime *dara.RuntimeOptions) (_result *BatchDescribeCdnIpInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpAddrList) {
		query["IpAddrList"] = request.IpAddrList
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDescribeCdnIpInfo"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDescribeCdnIpInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether one or more IP addresses are assigned to Alibaba Cloud CDN.
//
// Description:
//
// >The maximum number of times that each user can call this operation per second is 20.
//
// @param request - BatchDescribeCdnIpInfoRequest
//
// @return BatchDescribeCdnIpInfoResponse
func (client *Client) BatchDescribeCdnIpInfo(request *BatchDescribeCdnIpInfoRequest) (_result *BatchDescribeCdnIpInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDescribeCdnIpInfoResponse{}
	_body, _err := client.BatchDescribeCdnIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures multiple accelerated domain names at a time.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- You can specify multiple domain names and must separate them with commas (,). You can specify up to 50 domain names in each call.
//
//		- If the BatchSetCdnDomainConfig operation is successful, a unique configuration ID (ConfigId) is generated. You can use configuration IDs to update or delete configurations. For more information, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/388994.html).
//
// @param request - BatchSetCdnDomainConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetCdnDomainConfigResponse
func (client *Client) BatchSetCdnDomainConfigWithOptions(request *BatchSetCdnDomainConfigRequest, runtime *dara.RuntimeOptions) (_result *BatchSetCdnDomainConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetCdnDomainConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetCdnDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures multiple accelerated domain names at a time.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- You can specify multiple domain names and must separate them with commas (,). You can specify up to 50 domain names in each call.
//
//		- If the BatchSetCdnDomainConfig operation is successful, a unique configuration ID (ConfigId) is generated. You can use configuration IDs to update or delete configurations. For more information, see [Usage notes on ConfigId](https://help.aliyun.com/document_detail/388994.html).
//
// @param request - BatchSetCdnDomainConfigRequest
//
// @return BatchSetCdnDomainConfigResponse
func (client *Client) BatchSetCdnDomainConfig(request *BatchSetCdnDomainConfigRequest) (_result *BatchSetCdnDomainConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetCdnDomainConfigResponse{}
	_body, _err := client.BatchSetCdnDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量配置多个域名的灰度动态功能
//
// @param request - BatchSetGrayDomainFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetGrayDomainFunctionResponse
func (client *Client) BatchSetGrayDomainFunctionWithOptions(request *BatchSetGrayDomainFunctionRequest, runtime *dara.RuntimeOptions) (_result *BatchSetGrayDomainFunctionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSetGrayDomainFunction"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSetGrayDomainFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量配置多个域名的灰度动态功能
//
// @param request - BatchSetGrayDomainFunctionRequest
//
// @return BatchSetGrayDomainFunctionResponse
func (client *Client) BatchSetGrayDomainFunction(request *BatchSetGrayDomainFunctionRequest) (_result *BatchSetGrayDomainFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSetGrayDomainFunctionResponse{}
	_body, _err := client.BatchSetGrayDomainFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables one or more domain names at a time. After a domain name is enabled, the value of the DomainStatus parameter is changed to Online.
//
// Description:
//
//	  If a domain name specified in the request is in an invalid state or your account has an overdue payment, the domain name cannot be enabled.
//
//		- You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request.
//
// @param request - BatchStartCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStartCdnDomainResponse
func (client *Client) BatchStartCdnDomainWithOptions(request *BatchStartCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStartCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStartCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStartCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables one or more domain names at a time. After a domain name is enabled, the value of the DomainStatus parameter is changed to Online.
//
// Description:
//
//	  If a domain name specified in the request is in an invalid state or your account has an overdue payment, the domain name cannot be enabled.
//
//		- You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request.
//
// @param request - BatchStartCdnDomainRequest
//
// @return BatchStartCdnDomainResponse
func (client *Client) BatchStartCdnDomain(request *BatchStartCdnDomainRequest) (_result *BatchStartCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStartCdnDomainResponse{}
	_body, _err := client.BatchStartCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables one or more accelerated domain names at a time. After an accelerated domain name is disabled, the value of the DomainStatus parameter is changed to Offline.
//
// Description:
//
//	  After an accelerated domain name is disabled, Alibaba Cloud CDN retains its information and reroutes all the requests that are destined for the accelerated domain name to the origin.
//
//		- If you need to temporarily disable CDN acceleration for a domain name, we recommend that you call the StopDomain operation.
//
//		- You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request.
//
// @param request - BatchStopCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchStopCdnDomainResponse
func (client *Client) BatchStopCdnDomainWithOptions(request *BatchStopCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchStopCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchStopCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchStopCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables one or more accelerated domain names at a time. After an accelerated domain name is disabled, the value of the DomainStatus parameter is changed to Offline.
//
// Description:
//
//	  After an accelerated domain name is disabled, Alibaba Cloud CDN retains its information and reroutes all the requests that are destined for the accelerated domain name to the origin.
//
//		- If you need to temporarily disable CDN acceleration for a domain name, we recommend that you call the StopDomain operation.
//
//		- You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request.
//
// @param request - BatchStopCdnDomainRequest
//
// @return BatchStopCdnDomainResponse
func (client *Client) BatchStopCdnDomain(request *BatchStopCdnDomainRequest) (_result *BatchStopCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchStopCdnDomainResponse{}
	_body, _err := client.BatchStopCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of multiple accelerated domain names at a time.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - BatchUpdateCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateCdnDomainResponse
func (client *Client) BatchUpdateCdnDomainWithOptions(request *BatchUpdateCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of multiple accelerated domain names at a time.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- You can specify up to 50 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - BatchUpdateCdnDomainRequest
//
// @return BatchUpdateCdnDomainResponse
func (client *Client) BatchUpdateCdnDomain(request *BatchUpdateCdnDomainRequest) (_result *BatchUpdateCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateCdnDomainResponse{}
	_body, _err := client.BatchUpdateCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers the dynamic routing feature of Dynamic Content Delivery Network (DCDN) for an Alibaba Cloud CDN-accelerated domain name. After the registration is successful, the routing center generates the dynamic routing information and send it to DCDN points of presence (POPs). This is a prerequisite for you to transfer a domain name from Alibaba Cloud CDN to DCDN.
//
// @param request - CdnMigrateRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CdnMigrateRegisterResponse
func (client *Client) CdnMigrateRegisterWithOptions(request *CdnMigrateRegisterRequest, runtime *dara.RuntimeOptions) (_result *CdnMigrateRegisterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CdnMigrateRegister"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CdnMigrateRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers the dynamic routing feature of Dynamic Content Delivery Network (DCDN) for an Alibaba Cloud CDN-accelerated domain name. After the registration is successful, the routing center generates the dynamic routing information and send it to DCDN points of presence (POPs). This is a prerequisite for you to transfer a domain name from Alibaba Cloud CDN to DCDN.
//
// @param request - CdnMigrateRegisterRequest
//
// @return CdnMigrateRegisterResponse
func (client *Client) CdnMigrateRegister(request *CdnMigrateRegisterRequest) (_result *CdnMigrateRegisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CdnMigrateRegisterResponse{}
	_body, _err := client.CdnMigrateRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfer a domain name from Alibaba Cloud CDN to DCDN.
//
// @param request - ChangeCdnDomainToDcdnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeCdnDomainToDcdnResponse
func (client *Client) ChangeCdnDomainToDcdnWithOptions(request *ChangeCdnDomainToDcdnRequest, runtime *dara.RuntimeOptions) (_result *ChangeCdnDomainToDcdnResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Operation) {
		query["Operation"] = request.Operation
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeCdnDomainToDcdn"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeCdnDomainToDcdnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfer a domain name from Alibaba Cloud CDN to DCDN.
//
// @param request - ChangeCdnDomainToDcdnRequest
//
// @return ChangeCdnDomainToDcdnResponse
func (client *Client) ChangeCdnDomainToDcdn(request *ChangeCdnDomainToDcdnRequest) (_result *ChangeCdnDomainToDcdnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeCdnDomainToDcdnResponse{}
	_body, _err := client.ChangeCdnDomainToDcdnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a domain name exists.
//
// @param request - CheckCdnDomainExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCdnDomainExistResponse
func (client *Client) CheckCdnDomainExistWithOptions(request *CheckCdnDomainExistRequest, runtime *dara.RuntimeOptions) (_result *CheckCdnDomainExistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCdnDomainExist"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCdnDomainExistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a domain name exists.
//
// @param request - CheckCdnDomainExistRequest
//
// @return CheckCdnDomainExistResponse
func (client *Client) CheckCdnDomainExist(request *CheckCdnDomainExistRequest) (_result *CheckCdnDomainExistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCdnDomainExistResponse{}
	_body, _err := client.CheckCdnDomainExistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether an ICP filing is obtained for the domain name.
//
// @param request - CheckCdnDomainICPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCdnDomainICPResponse
func (client *Client) CheckCdnDomainICPWithOptions(request *CheckCdnDomainICPRequest, runtime *dara.RuntimeOptions) (_result *CheckCdnDomainICPResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckCdnDomainICP"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCdnDomainICPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether an ICP filing is obtained for the domain name.
//
// @param request - CheckCdnDomainICPRequest
//
// @return CheckCdnDomainICPResponse
func (client *Client) CheckCdnDomainICP(request *CheckCdnDomainICPRequest) (_result *CheckCdnDomainICPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCdnDomainICPResponse{}
	_body, _err := client.CheckCdnDomainICPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR).
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CreateCdnCertificateSigningRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCdnCertificateSigningRequestResponse
func (client *Client) CreateCdnCertificateSigningRequestWithOptions(request *CreateCdnCertificateSigningRequestRequest, runtime *dara.RuntimeOptions) (_result *CreateCdnCertificateSigningRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.CommonName) {
		query["CommonName"] = request.CommonName
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Organization) {
		query["Organization"] = request.Organization
	}

	if !dara.IsNil(request.OrganizationUnit) {
		query["OrganizationUnit"] = request.OrganizationUnit
	}

	if !dara.IsNil(request.SANs) {
		query["SANs"] = request.SANs
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCdnCertificateSigningRequest"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCdnCertificateSigningRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a certificate signing request (CSR).
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - CreateCdnCertificateSigningRequestRequest
//
// @return CreateCdnCertificateSigningRequestResponse
func (client *Client) CreateCdnCertificateSigningRequest(request *CreateCdnCertificateSigningRequestRequest) (_result *CreateCdnCertificateSigningRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCdnCertificateSigningRequestResponse{}
	_body, _err := client.CreateCdnCertificateSigningRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tracking task. After you create a tracking task, the system sends operations reports to you by email on a regular basis.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - CreateCdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCdnDeliverTaskResponse
func (client *Client) CreateCdnDeliverTaskWithOptions(request *CreateCdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCdnDeliverTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tracking task. After you create a tracking task, the system sends operations reports to you by email on a regular basis.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - CreateCdnDeliverTaskRequest
//
// @return CreateCdnDeliverTaskResponse
func (client *Client) CreateCdnDeliverTask(request *CreateCdnDeliverTaskRequest) (_result *CreateCdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCdnDeliverTaskResponse{}
	_body, _err := client.CreateCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom operations report.
//
// Description:
//
//	  This operation allows you to create a custom operations report for a specific domain name. You can view the statistics about the domain name in the report.
//
//		- You can call this operation up to three times per second per account.
//
// @param request - CreateCdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCdnSubTaskResponse
func (client *Client) CreateCdnSubTaskWithOptions(request *CreateCdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCdnSubTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom operations report.
//
// Description:
//
//	  This operation allows you to create a custom operations report for a specific domain name. You can view the statistics about the domain name in the report.
//
//		- You can call this operation up to three times per second per account.
//
// @param request - CreateCdnSubTaskRequest
//
// @return CreateCdnSubTaskResponse
func (client *Client) CreateCdnSubTask(request *CreateCdnSubTaskRequest) (_result *CreateCdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCdnSubTaskResponse{}
	_body, _err := client.CreateCdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables real-time log delivery for specific accelerated domain names.
//
// Description:
//
// >  You can call this API operation up to 100 times per second per account.
//
// @param request - CreateRealTimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRealTimeLogDeliveryResponse
func (client *Client) CreateRealTimeLogDeliveryWithOptions(request *CreateRealTimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *CreateRealTimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRealTimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRealTimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables real-time log delivery for specific accelerated domain names.
//
// Description:
//
// >  You can call this API operation up to 100 times per second per account.
//
// @param request - CreateRealTimeLogDeliveryRequest
//
// @return CreateRealTimeLogDeliveryResponse
func (client *Client) CreateRealTimeLogDelivery(request *CreateRealTimeLogDeliveryRequest) (_result *CreateRealTimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRealTimeLogDeliveryResponse{}
	_body, _err := client.CreateRealTimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to export resource usage details to an Excel file.
//
// Description:
//
//	  You can create a task to query data in the last year. The maximum time range that can be queried is one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - CreateUsageDetailDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsageDetailDataExportTaskResponse
func (client *Client) CreateUsageDetailDataExportTaskWithOptions(request *CreateUsageDetailDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateUsageDetailDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUsageDetailDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUsageDetailDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to export resource usage details to an Excel file.
//
// Description:
//
//	  You can create a task to query data in the last year. The maximum time range that can be queried is one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - CreateUsageDetailDataExportTaskRequest
//
// @return CreateUsageDetailDataExportTaskResponse
func (client *Client) CreateUsageDetailDataExportTask(request *CreateUsageDetailDataExportTaskRequest) (_result *CreateUsageDetailDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUsageDetailDataExportTaskResponse{}
	_body, _err := client.CreateUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to export your resource usage history to a PDF file.
//
// Description:
//
//	  You can create a task to query data in the last year. The maximum time range that can be queried is one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - CreateUserUsageDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserUsageDataExportTaskResponse
func (client *Client) CreateUserUsageDataExportTaskWithOptions(request *CreateUserUsageDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateUserUsageDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserUsageDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserUsageDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to export your resource usage history to a PDF file.
//
// Description:
//
//	  You can create a task to query data in the last year. The maximum time range that can be queried is one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - CreateUserUsageDataExportTaskRequest
//
// @return CreateUserUsageDataExportTaskResponse
func (client *Client) CreateUserUsageDataExportTask(request *CreateUserUsageDataExportTaskRequest) (_result *CreateUserUsageDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserUsageDataExportTaskResponse{}
	_body, _err := client.CreateUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes tracking tasks by task ID.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DeleteCdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCdnDeliverTaskResponse
func (client *Client) DeleteCdnDeliverTaskWithOptions(request *DeleteCdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteCdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCdnDeliverTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes tracking tasks by task ID.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DeleteCdnDeliverTaskRequest
//
// @return DeleteCdnDeliverTaskResponse
func (client *Client) DeleteCdnDeliverTask(request *DeleteCdnDeliverTaskRequest) (_result *DeleteCdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCdnDeliverTaskResponse{}
	_body, _err := client.DeleteCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an accelerated domain name from Alibaba Cloud CDN.
//
// Description:
//
//	  We recommend that you add an A record for the domain name in the system of your DNS service provider before you remove the domain name from Alibaba Cloud CDN. Otherwise, the domain name may become inaccessible. Proceed with caution.
//
//		- After you successfully call the DeleteCdnDomain operation, all records of the removed domain name are deleted. If you need to only disable the domain name, we recommend that you call the StopCdnDomain operation.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DeleteCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCdnDomainResponse
func (client *Client) DeleteCdnDomainWithOptions(request *DeleteCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an accelerated domain name from Alibaba Cloud CDN.
//
// Description:
//
//	  We recommend that you add an A record for the domain name in the system of your DNS service provider before you remove the domain name from Alibaba Cloud CDN. Otherwise, the domain name may become inaccessible. Proceed with caution.
//
//		- After you successfully call the DeleteCdnDomain operation, all records of the removed domain name are deleted. If you need to only disable the domain name, we recommend that you call the StopCdnDomain operation.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DeleteCdnDomainRequest
//
// @return DeleteCdnDomainResponse
func (client *Client) DeleteCdnDomain(request *DeleteCdnDomainRequest) (_result *DeleteCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCdnDomainResponse{}
	_body, _err := client.DeleteCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// >  You can call this API operation up to three times per second per account.
//
// @param request - DeleteCdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCdnSubTaskResponse
func (client *Client) DeleteCdnSubTaskWithOptions(runtime *dara.RuntimeOptions) (_result *DeleteCdnSubTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCdnSubTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// Description:
//
// >  You can call this API operation up to three times per second per account.
//
// @return DeleteCdnSubTaskResponse
func (client *Client) DeleteCdnSubTask() (_result *DeleteCdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCdnSubTaskResponse{}
	_body, _err := client.DeleteCdnSubTaskWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制实时日志删除接口
//
// @param request - DeleteCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainSampleRateResponse
func (client *Client) DeleteCustomDomainSampleRateWithOptions(request *DeleteCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomDomainSampleRate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制实时日志删除接口
//
// @param request - DeleteCustomDomainSampleRateRequest
//
// @return DeleteCustomDomainSampleRateResponse
func (client *Client) DeleteCustomDomainSampleRate(request *DeleteCustomDomainSampleRateRequest) (_result *DeleteCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomDomainSampleRateResponse{}
	_body, _err := client.DeleteCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified Function Compute trigger.
//
// @param request - DeleteFCTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFCTriggerResponse
func (client *Client) DeleteFCTriggerWithOptions(request *DeleteFCTriggerRequest, runtime *dara.RuntimeOptions) (_result *DeleteFCTriggerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TriggerARN) {
		query["TriggerARN"] = request.TriggerARN
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFCTrigger"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFCTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified Function Compute trigger.
//
// @param request - DeleteFCTriggerRequest
//
// @return DeleteFCTriggerResponse
func (client *Client) DeleteFCTrigger(request *DeleteFCTriggerRequest) (_result *DeleteFCTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFCTriggerResponse{}
	_body, _err := client.DeleteFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the Logstore that is used by a specified configuration record of real-time log delivery.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRealTimeLogLogstoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRealTimeLogLogstoreResponse
func (client *Client) DeleteRealTimeLogLogstoreWithOptions(request *DeleteRealTimeLogLogstoreRequest, runtime *dara.RuntimeOptions) (_result *DeleteRealTimeLogLogstoreResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRealTimeLogLogstore"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRealTimeLogLogstoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the Logstore that is used by a specified configuration record of real-time log delivery.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRealTimeLogLogstoreRequest
//
// @return DeleteRealTimeLogLogstoreResponse
func (client *Client) DeleteRealTimeLogLogstore(request *DeleteRealTimeLogLogstoreRequest) (_result *DeleteRealTimeLogLogstoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRealTimeLogLogstoreResponse{}
	_body, _err := client.DeleteRealTimeLogLogstoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the configurations of real-time log delivery for specific accelerated domain names.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRealtimeLogDeliveryResponse
func (client *Client) DeleteRealtimeLogDeliveryWithOptions(request *DeleteRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DeleteRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the configurations of real-time log delivery for specific accelerated domain names.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRealtimeLogDeliveryRequest
//
// @return DeleteRealtimeLogDeliveryResponse
func (client *Client) DeleteRealtimeLogDelivery(request *DeleteRealtimeLogDeliveryRequest) (_result *DeleteRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRealtimeLogDeliveryResponse{}
	_body, _err := client.DeleteRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes specified configurations of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DeleteSpecificConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSpecificConfigResponse
func (client *Client) DeleteSpecificConfigWithOptions(request *DeleteSpecificConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSpecificConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSpecificConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSpecificConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes specified configurations of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DeleteSpecificConfigRequest
//
// @return DeleteSpecificConfigResponse
func (client *Client) DeleteSpecificConfig(request *DeleteSpecificConfigRequest) (_result *DeleteSpecificConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSpecificConfigResponse{}
	_body, _err := client.DeleteSpecificConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified configuration of the staging environment.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DeleteSpecificStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSpecificStagingConfigResponse
func (client *Client) DeleteSpecificStagingConfigWithOptions(request *DeleteSpecificStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteSpecificStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSpecificStagingConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSpecificStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified configuration of the staging environment.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DeleteSpecificStagingConfigRequest
//
// @return DeleteSpecificStagingConfigResponse
func (client *Client) DeleteSpecificStagingConfig(request *DeleteSpecificStagingConfigRequest) (_result *DeleteSpecificStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSpecificStagingConfigResponse{}
	_body, _err := client.DeleteSpecificStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task that was used to export usage details.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteUsageDetailDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUsageDetailDataExportTaskResponse
func (client *Client) DeleteUsageDetailDataExportTaskWithOptions(request *DeleteUsageDetailDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteUsageDetailDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUsageDetailDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUsageDetailDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task that was used to export usage details.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteUsageDetailDataExportTaskRequest
//
// @return DeleteUsageDetailDataExportTaskResponse
func (client *Client) DeleteUsageDetailDataExportTask(request *DeleteUsageDetailDataExportTaskRequest) (_result *DeleteUsageDetailDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUsageDetailDataExportTaskResponse{}
	_body, _err := client.DeleteUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task that was used to export usage history.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteUserUsageDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserUsageDataExportTaskResponse
func (client *Client) DeleteUserUsageDataExportTaskWithOptions(request *DeleteUserUsageDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserUsageDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserUsageDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserUsageDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task that was used to export usage history.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DeleteUserUsageDataExportTaskRequest
//
// @return DeleteUserUsageDataExportTaskResponse
func (client *Client) DeleteUserUsageDataExportTask(request *DeleteUserUsageDataExportTaskRequest) (_result *DeleteUserUsageDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserUsageDataExportTaskResponse{}
	_body, _err := client.DeleteUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries countries and regions that can be added to the blacklist.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeBlockedRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBlockedRegionsResponse
func (client *Client) DescribeBlockedRegionsWithOptions(request *DescribeBlockedRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeBlockedRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBlockedRegions"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBlockedRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries countries and regions that can be added to the blacklist.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeBlockedRegionsRequest
//
// @return DescribeBlockedRegionsResponse
func (client *Client) DescribeBlockedRegions(request *DescribeBlockedRegionsRequest) (_result *DescribeBlockedRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBlockedRegionsResponse{}
	_body, _err := client.DescribeBlockedRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an SSL certificate.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeCdnCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnCertificateDetailResponse
func (client *Client) DescribeCdnCertificateDetailWithOptions(request *DescribeCdnCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnCertificateDetail"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an SSL certificate.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeCdnCertificateDetailRequest
//
// @return DescribeCdnCertificateDetailResponse
func (client *Client) DescribeCdnCertificateDetail(request *DescribeCdnCertificateDetailRequest) (_result *DescribeCdnCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnCertificateDetailResponse{}
	_body, _err := client.DescribeCdnCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries certificate details by certificate ID.
//
// @param request - DescribeCdnCertificateDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnCertificateDetailByIdResponse
func (client *Client) DescribeCdnCertificateDetailByIdWithOptions(request *DescribeCdnCertificateDetailByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnCertificateDetailByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnCertificateDetailById"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnCertificateDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries certificate details by certificate ID.
//
// @param request - DescribeCdnCertificateDetailByIdRequest
//
// @return DescribeCdnCertificateDetailByIdResponse
func (client *Client) DescribeCdnCertificateDetailById(request *DescribeCdnCertificateDetailByIdRequest) (_result *DescribeCdnCertificateDetailByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnCertificateDetailByIdResponse{}
	_body, _err := client.DescribeCdnCertificateDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeCdnCertificateList is deprecated, please use Cdn::2018-05-10::DescribeCdnSSLCertificateList instead.
//
// Summary:
//
// Queries the certificates of accelerated domain names.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnCertificateListResponse
// Deprecated
func (client *Client) DescribeCdnCertificateListWithOptions(request *DescribeCdnCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnCertificateList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeCdnCertificateList is deprecated, please use Cdn::2018-05-10::DescribeCdnSSLCertificateList instead.
//
// Summary:
//
// Queries the certificates of accelerated domain names.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnCertificateListRequest
//
// @return DescribeCdnCertificateListResponse
// Deprecated
func (client *Client) DescribeCdnCertificateList(request *DescribeCdnCertificateListRequest) (_result *DescribeCdnCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnCertificateListResponse{}
	_body, _err := client.DescribeCdnCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Internet service provider (ISP), region, and country that are required for advanced conditions.
//
// @param request - DescribeCdnConditionIPBInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnConditionIPBInfoResponse
func (client *Client) DescribeCdnConditionIPBInfoWithOptions(request *DescribeCdnConditionIPBInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnConditionIPBInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		query["DataId"] = request.DataId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnConditionIPBInfo"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnConditionIPBInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Internet service provider (ISP), region, and country that are required for advanced conditions.
//
// @param request - DescribeCdnConditionIPBInfoRequest
//
// @return DescribeCdnConditionIPBInfoResponse
func (client *Client) DescribeCdnConditionIPBInfo(request *DescribeCdnConditionIPBInfoRequest) (_result *DescribeCdnConditionIPBInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnConditionIPBInfoResponse{}
	_body, _err := client.DescribeCdnConditionIPBInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that are deleted from your account.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeCdnDeletedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDeletedDomainsResponse
func (client *Client) DescribeCdnDeletedDomainsWithOptions(request *DescribeCdnDeletedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDeletedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeCdnDeletedDomains"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDeletedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are deleted from your account.
//
// Description:
//
// > You can call this operation up to 10 times per second per account.
//
// @param request - DescribeCdnDeletedDomainsRequest
//
// @return DescribeCdnDeletedDomainsResponse
func (client *Client) DescribeCdnDeletedDomains(request *DescribeCdnDeletedDomainsRequest) (_result *DescribeCdnDeletedDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDeletedDomainsResponse{}
	_body, _err := client.DescribeCdnDeletedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more tracking tasks of operations reports.
//
// Description:
//
// > You can call this operation up to 3 times per second per account.
//
// @param request - DescribeCdnDeliverListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDeliverListResponse
func (client *Client) DescribeCdnDeliverListWithOptions(request *DescribeCdnDeliverListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDeliverListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliverId) {
		query["DeliverId"] = request.DeliverId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDeliverList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDeliverListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more tracking tasks of operations reports.
//
// Description:
//
// > You can call this operation up to 3 times per second per account.
//
// @param request - DescribeCdnDeliverListRequest
//
// @return DescribeCdnDeliverListResponse
func (client *Client) DescribeCdnDeliverList(request *DescribeCdnDeliverListRequest) (_result *DescribeCdnDeliverListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDeliverListResponse{}
	_body, _err := client.DescribeCdnDeliverListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 天翼定制化小时日志下载接口
//
// @param request - DescribeCdnDomainAtoaLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainAtoaLogsResponse
func (client *Client) DescribeCdnDomainAtoaLogsWithOptions(request *DescribeCdnDomainAtoaLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainAtoaLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainAtoaLogs"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainAtoaLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 天翼定制化小时日志下载接口
//
// @param request - DescribeCdnDomainAtoaLogsRequest
//
// @return DescribeCdnDomainAtoaLogsResponse
func (client *Client) DescribeCdnDomainAtoaLogs(request *DescribeCdnDomainAtoaLogsRequest) (_result *DescribeCdnDomainAtoaLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainAtoaLogsResponse{}
	_body, _err := client.DescribeCdnDomainAtoaLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by SSL certificate.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnDomainByCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainByCertificateResponse
func (client *Client) DescribeCdnDomainByCertificateWithOptions(request *DescribeCdnDomainByCertificateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainByCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Exact) {
		query["Exact"] = request.Exact
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SSLStatus) {
		query["SSLStatus"] = request.SSLStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainByCertificate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainByCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by SSL certificate.
//
// Description:
//
// >  You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnDomainByCertificateRequest
//
// @return DescribeCdnDomainByCertificateResponse
func (client *Client) DescribeCdnDomainByCertificate(request *DescribeCdnDomainByCertificateRequest) (_result *DescribeCdnDomainByCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainByCertificateResponse{}
	_body, _err := client.DescribeCdnDomainByCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name. You can query the configurations of one or more features at the same time.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnDomainConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainConfigsResponse
func (client *Client) DescribeCdnDomainConfigsWithOptions(request *DescribeCdnDomainConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainConfigs"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an accelerated domain name. You can query the configurations of one or more features at the same time.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnDomainConfigsRequest
//
// @return DescribeCdnDomainConfigsResponse
func (client *Client) DescribeCdnDomainConfigs(request *DescribeCdnDomainConfigsRequest) (_result *DescribeCdnDomainConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainConfigsResponse{}
	_body, _err := client.DescribeCdnDomainConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainDetailResponse
func (client *Client) DescribeCdnDomainDetailWithOptions(request *DescribeCdnDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainDetail"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnDomainDetailRequest
//
// @return DescribeCdnDomainDetailResponse
func (client *Client) DescribeCdnDomainDetail(request *DescribeCdnDomainDetailRequest) (_result *DescribeCdnDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainDetailResponse{}
	_body, _err := client.DescribeCdnDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the address where you can download the log data of a specific domain name.
//
// Description:
//
//	  If you do not set **StartTime*	- or **EndTime**, the request returns the data collected in the last 24 hours. If you set both **StartTime*	- and **EndTime**, the request returns the data collected within the specified time range.
//
//		- The log data is collected every hour.
//
//		- You can call this operation up to 100 times per second per account.
//
//		- You can query only logs in the last month. The start time and the current time cannot exceed 31 days.
//
// @param request - DescribeCdnDomainLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainLogsResponse
func (client *Client) DescribeCdnDomainLogsWithOptions(request *DescribeCdnDomainLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainLogs"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the address where you can download the log data of a specific domain name.
//
// Description:
//
//	  If you do not set **StartTime*	- or **EndTime**, the request returns the data collected in the last 24 hours. If you set both **StartTime*	- and **EndTime**, the request returns the data collected within the specified time range.
//
//		- The log data is collected every hour.
//
//		- You can call this operation up to 100 times per second per account.
//
//		- You can query only logs in the last month. The start time and the current time cannot exceed 31 days.
//
// @param request - DescribeCdnDomainLogsRequest
//
// @return DescribeCdnDomainLogsResponse
func (client *Client) DescribeCdnDomainLogs(request *DescribeCdnDomainLogsRequest) (_result *DescribeCdnDomainLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainLogsResponse{}
	_body, _err := client.DescribeCdnDomainLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询离线日志下载地址
//
// @param request - DescribeCdnDomainLogsExTtlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainLogsExTtlResponse
func (client *Client) DescribeCdnDomainLogsExTtlWithOptions(request *DescribeCdnDomainLogsExTtlRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainLogsExTtlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainLogsExTtl"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainLogsExTtlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询离线日志下载地址
//
// @param request - DescribeCdnDomainLogsExTtlRequest
//
// @return DescribeCdnDomainLogsExTtlResponse
func (client *Client) DescribeCdnDomainLogsExTtl(request *DescribeCdnDomainLogsExTtlRequest) (_result *DescribeCdnDomainLogsExTtlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainLogsExTtlResponse{}
	_body, _err := client.DescribeCdnDomainLogsExTtlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of features in the staging environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnDomainStagingConfigResponse
func (client *Client) DescribeCdnDomainStagingConfigWithOptions(request *DescribeCdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		query["FunctionNames"] = request.FunctionNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnDomainStagingConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of features in the staging environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnDomainStagingConfigRequest
//
// @return DescribeCdnDomainStagingConfigResponse
func (client *Client) DescribeCdnDomainStagingConfig(request *DescribeCdnDomainStagingConfigRequest) (_result *DescribeCdnDomainStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnDomainStagingConfigResponse{}
	_body, _err := client.DescribeCdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeCdnFullDomainsBlockIPConfig operation to query the configurations of full blocking.
//
// Description:
//
// >
//
//   - To use this operation,[submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - If you specify IP addresses or CIDR blocks, IP addresses that are effective and the corresponding expiration time are returned. If you do not specify IP addresses or CIDR blocks, all effective IP addresses and the corresponding expiration time are returned.
//
//   - The results are written to OSS and returned as OSS URLs. The content in OSS objects is in the format of IP address-Corresponding expiration time. The expiration time is in the YYYY-MM-DD hh:mm:ss format.
//
//   - You can share OSS URLs with others. The shared URLs are valid for three days.
//
// @param request - DescribeCdnFullDomainsBlockIPConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnFullDomainsBlockIPConfigResponse
func (client *Client) DescribeCdnFullDomainsBlockIPConfigWithOptions(request *DescribeCdnFullDomainsBlockIPConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnFullDomainsBlockIPConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnFullDomainsBlockIPConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnFullDomainsBlockIPConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeCdnFullDomainsBlockIPConfig operation to query the configurations of full blocking.
//
// Description:
//
// >
//
//   - To use this operation,[submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - If you specify IP addresses or CIDR blocks, IP addresses that are effective and the corresponding expiration time are returned. If you do not specify IP addresses or CIDR blocks, all effective IP addresses and the corresponding expiration time are returned.
//
//   - The results are written to OSS and returned as OSS URLs. The content in OSS objects is in the format of IP address-Corresponding expiration time. The expiration time is in the YYYY-MM-DD hh:mm:ss format.
//
//   - You can share OSS URLs with others. The shared URLs are valid for three days.
//
// @param request - DescribeCdnFullDomainsBlockIPConfigRequest
//
// @return DescribeCdnFullDomainsBlockIPConfigResponse
func (client *Client) DescribeCdnFullDomainsBlockIPConfig(request *DescribeCdnFullDomainsBlockIPConfigRequest) (_result *DescribeCdnFullDomainsBlockIPConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnFullDomainsBlockIPConfigResponse{}
	_body, _err := client.DescribeCdnFullDomainsBlockIPConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the blocking history.
//
// Description:
//
// >
//
//   - To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - For a specified IP addresses and time range, the time when the IP address was delivered to the edge and the corresponding result are returned.
//
//   - If a specified IP address or CIDR block has multiple blocking records in a specified time range, the records are sorted by delivery time in descending order.
//
//   - The maximum time range to query is 90 days.
//
//   - If no blocking record exists or delivery fails for the given IP address and time range, the delivery time is empty.
//
// @param request - DescribeCdnFullDomainsBlockIPHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnFullDomainsBlockIPHistoryResponse
func (client *Client) DescribeCdnFullDomainsBlockIPHistoryWithOptions(request *DescribeCdnFullDomainsBlockIPHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnFullDomainsBlockIPHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnFullDomainsBlockIPHistory"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnFullDomainsBlockIPHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the blocking history.
//
// Description:
//
// >
//
//   - To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - For a specified IP addresses and time range, the time when the IP address was delivered to the edge and the corresponding result are returned.
//
//   - If a specified IP address or CIDR block has multiple blocking records in a specified time range, the records are sorted by delivery time in descending order.
//
//   - The maximum time range to query is 90 days.
//
//   - If no blocking record exists or delivery fails for the given IP address and time range, the delivery time is empty.
//
// @param request - DescribeCdnFullDomainsBlockIPHistoryRequest
//
// @return DescribeCdnFullDomainsBlockIPHistoryResponse
func (client *Client) DescribeCdnFullDomainsBlockIPHistory(request *DescribeCdnFullDomainsBlockIPHistoryRequest) (_result *DescribeCdnFullDomainsBlockIPHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnFullDomainsBlockIPHistoryResponse{}
	_body, _err := client.DescribeCdnFullDomainsBlockIPHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about SSL certificates that belong to your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnHttpsDomainListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnHttpsDomainListResponse
func (client *Client) DescribeCdnHttpsDomainListWithOptions(request *DescribeCdnHttpsDomainListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnHttpsDomainListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
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
		Action:      dara.String("DescribeCdnHttpsDomainList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnHttpsDomainListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about SSL certificates that belong to your Alibaba Cloud account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnHttpsDomainListRequest
//
// @return DescribeCdnHttpsDomainListResponse
func (client *Client) DescribeCdnHttpsDomainList(request *DescribeCdnHttpsDomainListRequest) (_result *DescribeCdnHttpsDomainListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnHttpsDomainListResponse{}
	_body, _err := client.DescribeCdnHttpsDomainListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the registration status of the dynamic routing feature of Dynamic Content Delivery Network (DCDN) for a domain name that is added to Alibaba Cloud CDN.
//
// Description:
//
// >  If a domain name is not transferred from Alibaba Cloud CDN to DCDN after it is registered in the routing center of DCDN, the registration information is retained for only one day.
//
// @param request - DescribeCdnMigrateRegisterStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnMigrateRegisterStatusResponse
func (client *Client) DescribeCdnMigrateRegisterStatusWithOptions(request *DescribeCdnMigrateRegisterStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnMigrateRegisterStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnMigrateRegisterStatus"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnMigrateRegisterStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the registration status of the dynamic routing feature of Dynamic Content Delivery Network (DCDN) for a domain name that is added to Alibaba Cloud CDN.
//
// Description:
//
// >  If a domain name is not transferred from Alibaba Cloud CDN to DCDN after it is registered in the routing center of DCDN, the registration information is retained for only one day.
//
// @param request - DescribeCdnMigrateRegisterStatusRequest
//
// @return DescribeCdnMigrateRegisterStatusResponse
func (client *Client) DescribeCdnMigrateRegisterStatus(request *DescribeCdnMigrateRegisterStatusRequest) (_result *DescribeCdnMigrateRegisterStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnMigrateRegisterStatusResponse{}
	_body, _err := client.DescribeCdnMigrateRegisterStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the code of a commodity by account UID.
//
// @param request - DescribeCdnOrderCommodityCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnOrderCommodityCodeResponse
func (client *Client) DescribeCdnOrderCommodityCodeWithOptions(request *DescribeCdnOrderCommodityCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnOrderCommodityCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityCode) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnOrderCommodityCode"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnOrderCommodityCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the code of a commodity by account UID.
//
// @param request - DescribeCdnOrderCommodityCodeRequest
//
// @return DescribeCdnOrderCommodityCodeResponse
func (client *Client) DescribeCdnOrderCommodityCode(request *DescribeCdnOrderCommodityCodeRequest) (_result *DescribeCdnOrderCommodityCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnOrderCommodityCodeResponse{}
	_body, _err := client.DescribeCdnOrderCommodityCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Internet service providers (ISPs) and regions that are supported by Alibaba Cloud CDN.
//
// Description:
//
//	  The lists of ISPs and regions that are supported by Alibaba Cloud CDN are updated and published on the Alibaba Cloud International site.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnRegionAndIspRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnRegionAndIspResponse
func (client *Client) DescribeCdnRegionAndIspWithOptions(request *DescribeCdnRegionAndIspRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnRegionAndIspResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnRegionAndIsp"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnRegionAndIspResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Internet service providers (ISPs) and regions that are supported by Alibaba Cloud CDN.
//
// Description:
//
//	  The lists of ISPs and regions that are supported by Alibaba Cloud CDN are updated and published on the Alibaba Cloud International site.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnRegionAndIspRequest
//
// @return DescribeCdnRegionAndIspResponse
func (client *Client) DescribeCdnRegionAndIsp(request *DescribeCdnRegionAndIspRequest) (_result *DescribeCdnRegionAndIspResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnRegionAndIspResponse{}
	_body, _err := client.DescribeCdnRegionAndIspWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the content of an operations report.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeCdnReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnReportResponse
func (client *Client) DescribeCdnReportWithOptions(request *DescribeCdnReportRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnReportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HttpCode) {
		query["HttpCode"] = request.HttpCode
	}

	if !dara.IsNil(request.IsOverseas) {
		query["IsOverseas"] = request.IsOverseas
	}

	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnReport"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the content of an operations report.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - DescribeCdnReportRequest
//
// @return DescribeCdnReportResponse
func (client *Client) DescribeCdnReport(request *DescribeCdnReportRequest) (_result *DescribeCdnReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnReportResponse{}
	_body, _err := client.DescribeCdnReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries operations reports.
//
// Description:
//
//	  This operation queries the metadata of all operations reports. The statistics in the reports are not returned.
//
//		- You can call this operation up to three times per second per account.
//
// @param request - DescribeCdnReportListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnReportListResponse
func (client *Client) DescribeCdnReportListWithOptions(request *DescribeCdnReportListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnReportListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnReportList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnReportListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries operations reports.
//
// Description:
//
//	  This operation queries the metadata of all operations reports. The statistics in the reports are not returned.
//
//		- You can call this operation up to three times per second per account.
//
// @param request - DescribeCdnReportListRequest
//
// @return DescribeCdnReportListResponse
func (client *Client) DescribeCdnReportList(request *DescribeCdnReportListRequest) (_result *DescribeCdnReportListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnReportListResponse{}
	_body, _err := client.DescribeCdnReportListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a ShangMi (SM) certificate.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeCdnSMCertificateDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnSMCertificateDetailResponse
func (client *Client) DescribeCdnSMCertificateDetailWithOptions(request *DescribeCdnSMCertificateDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnSMCertificateDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnSMCertificateDetail"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnSMCertificateDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a ShangMi (SM) certificate.
//
// Description:
//
// > You can call this operation up to 20 times per second per account.
//
// @param request - DescribeCdnSMCertificateDetailRequest
//
// @return DescribeCdnSMCertificateDetailResponse
func (client *Client) DescribeCdnSMCertificateDetail(request *DescribeCdnSMCertificateDetailRequest) (_result *DescribeCdnSMCertificateDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnSMCertificateDetailResponse{}
	_body, _err := client.DescribeCdnSMCertificateDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ShangMi (SM) certificates of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnSMCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnSMCertificateListResponse
func (client *Client) DescribeCdnSMCertificateListWithOptions(request *DescribeCdnSMCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnSMCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnSMCertificateList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnSMCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ShangMi (SM) certificates of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnSMCertificateListRequest
//
// @return DescribeCdnSMCertificateListResponse
func (client *Client) DescribeCdnSMCertificateList(request *DescribeCdnSMCertificateListRequest) (_result *DescribeCdnSMCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnSMCertificateListResponse{}
	_body, _err := client.DescribeCdnSMCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate list by domain name.
//
// @param request - DescribeCdnSSLCertificateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnSSLCertificateListResponse
func (client *Client) DescribeCdnSSLCertificateListWithOptions(request *DescribeCdnSSLCertificateListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnSSLCertificateListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKeyword) {
		query["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnSSLCertificateList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnSSLCertificateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate list by domain name.
//
// @param request - DescribeCdnSSLCertificateListRequest
//
// @return DescribeCdnSSLCertificateListResponse
func (client *Client) DescribeCdnSSLCertificateList(request *DescribeCdnSSLCertificateListRequest) (_result *DescribeCdnSSLCertificateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnSSLCertificateListResponse{}
	_body, _err := client.DescribeCdnSSLCertificateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about security features of Alibaba Cloud CDN.
//
// @param request - DescribeCdnSecFuncInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnSecFuncInfoResponse
func (client *Client) DescribeCdnSecFuncInfoWithOptions(request *DescribeCdnSecFuncInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnSecFuncInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.SecFuncType) {
		query["SecFuncType"] = request.SecFuncType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnSecFuncInfo"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnSecFuncInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about security features of Alibaba Cloud CDN.
//
// @param request - DescribeCdnSecFuncInfoRequest
//
// @return DescribeCdnSecFuncInfoResponse
func (client *Client) DescribeCdnSecFuncInfo(request *DescribeCdnSecFuncInfoRequest) (_result *DescribeCdnSecFuncInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnSecFuncInfoResponse{}
	_body, _err := client.DescribeCdnSecFuncInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of your Alibaba Cloud CDN service. The information includes the service activation time, the current service status, the current metering method, and the metering method for the next cycle.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnServiceResponse
func (client *Client) DescribeCdnServiceWithOptions(request *DescribeCdnServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnService"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of your Alibaba Cloud CDN service. The information includes the service activation time, the current service status, the current metering method, and the metering method for the next cycle.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnServiceRequest
//
// @return DescribeCdnServiceResponse
func (client *Client) DescribeCdnService(request *DescribeCdnServiceRequest) (_result *DescribeCdnServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnServiceResponse{}
	_body, _err := client.DescribeCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tracking tasks that you have created.
//
// Description:
//
//	  By default, this operation queries all custom operations reports. However, only one operations report can be displayed. Therefore, only one operations report is returned.
//
//		- You can call this operation up to three times per second per account.
//
// @param request - DescribeCdnSubListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnSubListResponse
func (client *Client) DescribeCdnSubListWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeCdnSubListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnSubList"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnSubListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tracking tasks that you have created.
//
// Description:
//
//	  By default, this operation queries all custom operations reports. However, only one operations report can be displayed. Therefore, only one operations report is returned.
//
//		- You can call this operation up to three times per second per account.
//
// @return DescribeCdnSubListResponse
func (client *Client) DescribeCdnSubList() (_result *DescribeCdnSubListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnSubListResponse{}
	_body, _err := client.DescribeCdnSubListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types of domain names.
//
// @param request - DescribeCdnTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnTypesResponse
func (client *Client) DescribeCdnTypesWithOptions(request *DescribeCdnTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnTypes"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of domain names.
//
// @param request - DescribeCdnTypesRequest
//
// @return DescribeCdnTypesResponse
func (client *Client) DescribeCdnTypes(request *DescribeCdnTypesRequest) (_result *DescribeCdnTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnTypesResponse{}
	_body, _err := client.DescribeCdnTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the billing history under your Alibaba Cloud account.
//
// Description:
//
//	  You can query billing history up to the last one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnUserBillHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserBillHistoryResponse
func (client *Client) DescribeCdnUserBillHistoryWithOptions(request *DescribeCdnUserBillHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserBillHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserBillHistory"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserBillHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the billing history under your Alibaba Cloud account.
//
// Description:
//
//	  You can query billing history up to the last one month.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnUserBillHistoryRequest
//
// @return DescribeCdnUserBillHistoryResponse
func (client *Client) DescribeCdnUserBillHistory(request *DescribeCdnUserBillHistoryRequest) (_result *DescribeCdnUserBillHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserBillHistoryResponse{}
	_body, _err := client.DescribeCdnUserBillHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Estimates resource usage of the current month.
//
// Description:
//
// You can call this operation to estimate resource usage of the current month based on the metering method that is specified on the first day of the current month. You can call this operation to estimate resource usage only of the current month within your Alibaba Cloud account. The time range used for the estimation starts at 00:00 on the first day of the current month and ends 2 hours earlier than the current time.
//
//   - Pay by monthly 95th percentile: The top 5% values between the start time and end time are excluded. The estimated value is the highest value among the remaining values.
//
//   - Pay by average daily peak bandwidth per month: Estimated value = Sum of daily peak bandwidth values/Number of days. The current day is excluded.
//
//   - Pay by 4th peak bandwidth per month: The estimated value is the 4th peak bandwidth value between the start time and end time. If the time range is less than four days, the estimated value is 0.
//
//   - Pay by average daily 95th percentile bandwidth per month: Estimated value = Sum of daily 95th percentile bandwidth values/Number of days. The current day is excluded.
//
//   - Pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00: The top 5% values between the start time and end time are excluded. The estimated value is the highest value among the remaining values.
//
// > You can call this operation only once per second per account.
//
// @param request - DescribeCdnUserBillPredictionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserBillPredictionResponse
func (client *Client) DescribeCdnUserBillPredictionWithOptions(request *DescribeCdnUserBillPredictionRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserBillPredictionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserBillPrediction"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserBillPredictionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Estimates resource usage of the current month.
//
// Description:
//
// You can call this operation to estimate resource usage of the current month based on the metering method that is specified on the first day of the current month. You can call this operation to estimate resource usage only of the current month within your Alibaba Cloud account. The time range used for the estimation starts at 00:00 on the first day of the current month and ends 2 hours earlier than the current time.
//
//   - Pay by monthly 95th percentile: The top 5% values between the start time and end time are excluded. The estimated value is the highest value among the remaining values.
//
//   - Pay by average daily peak bandwidth per month: Estimated value = Sum of daily peak bandwidth values/Number of days. The current day is excluded.
//
//   - Pay by 4th peak bandwidth per month: The estimated value is the 4th peak bandwidth value between the start time and end time. If the time range is less than four days, the estimated value is 0.
//
//   - Pay by average daily 95th percentile bandwidth per month: Estimated value = Sum of daily 95th percentile bandwidth values/Number of days. The current day is excluded.
//
//   - Pay by 95th percentile bandwidth with 50% off from 00:00 to 08:00: The top 5% values between the start time and end time are excluded. The estimated value is the highest value among the remaining values.
//
// > You can call this operation only once per second per account.
//
// @param request - DescribeCdnUserBillPredictionRequest
//
// @return DescribeCdnUserBillPredictionResponse
func (client *Client) DescribeCdnUserBillPrediction(request *DescribeCdnUserBillPredictionRequest) (_result *DescribeCdnUserBillPredictionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserBillPredictionResponse{}
	_body, _err := client.DescribeCdnUserBillPredictionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the metering methods of an account. The maximum time range to query is one month.
//
// Description:
//
//	You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnUserBillTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserBillTypeResponse
func (client *Client) DescribeCdnUserBillTypeWithOptions(request *DescribeCdnUserBillTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserBillTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserBillType"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserBillTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the metering methods of an account. The maximum time range to query is one month.
//
// Description:
//
//	You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCdnUserBillTypeRequest
//
// @return DescribeCdnUserBillTypeResponse
func (client *Client) DescribeCdnUserBillType(request *DescribeCdnUserBillTypeRequest) (_result *DescribeCdnUserBillTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserBillTypeResponse{}
	_body, _err := client.DescribeCdnUserBillTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries configurations of security features.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserConfigsResponse
func (client *Client) DescribeCdnUserConfigsWithOptions(request *DescribeCdnUserConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserConfigs"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configurations of security features.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserConfigsRequest
//
// @return DescribeCdnUserConfigsResponse
func (client *Client) DescribeCdnUserConfigs(request *DescribeCdnUserConfigsRequest) (_result *DescribeCdnUserConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserConfigsResponse{}
	_body, _err := client.DescribeCdnUserConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names that have specified features configured and the status of the domain names.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 100.
//
// @param request - DescribeCdnUserDomainsByFuncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserDomainsByFuncResponse
func (client *Client) DescribeCdnUserDomainsByFuncWithOptions(request *DescribeCdnUserDomainsByFuncRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserDomainsByFuncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FuncId) {
		query["FuncId"] = request.FuncId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserDomainsByFunc"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserDomainsByFuncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names that have specified features configured and the status of the domain names.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 100.
//
// @param request - DescribeCdnUserDomainsByFuncRequest
//
// @return DescribeCdnUserDomainsByFuncResponse
func (client *Client) DescribeCdnUserDomainsByFunc(request *DescribeCdnUserDomainsByFuncRequest) (_result *DescribeCdnUserDomainsByFuncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserDomainsByFuncResponse{}
	_body, _err := client.DescribeCdnUserDomainsByFuncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quotas and usage of Alibaba Cloud CDN resources.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserQuotaResponse
func (client *Client) DescribeCdnUserQuotaWithOptions(request *DescribeCdnUserQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserQuota"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quotas and usage of Alibaba Cloud CDN resources.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserQuotaRequest
//
// @return DescribeCdnUserQuotaResponse
func (client *Client) DescribeCdnUserQuota(request *DescribeCdnUserQuotaRequest) (_result *DescribeCdnUserQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserQuotaResponse{}
	_body, _err := client.DescribeCdnUserQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource plans that you have purchased for Alibaba Cloud CDN.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserResourcePackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnUserResourcePackageResponse
func (client *Client) DescribeCdnUserResourcePackageWithOptions(request *DescribeCdnUserResourcePackageRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnUserResourcePackageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnUserResourcePackage"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnUserResourcePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource plans that you have purchased for Alibaba Cloud CDN.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeCdnUserResourcePackageRequest
//
// @return DescribeCdnUserResourcePackageResponse
func (client *Client) DescribeCdnUserResourcePackage(request *DescribeCdnUserResourcePackageRequest) (_result *DescribeCdnUserResourcePackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnUserResourcePackageResponse{}
	_body, _err := client.DescribeCdnUserResourcePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names that use Web Application Firewall (WAF).
//
// Description:
//
// > You can call this operation up to 150 times per second per account.
//
// @param request - DescribeCdnWafDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCdnWafDomainResponse
func (client *Client) DescribeCdnWafDomainWithOptions(request *DescribeCdnWafDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeCdnWafDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCdnWafDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCdnWafDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names that use Web Application Firewall (WAF).
//
// Description:
//
// > You can call this operation up to 150 times per second per account.
//
// @param request - DescribeCdnWafDomainRequest
//
// @return DescribeCdnWafDomainResponse
func (client *Client) DescribeCdnWafDomain(request *DescribeCdnWafDomainRequest) (_result *DescribeCdnWafDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCdnWafDomainResponse{}
	_body, _err := client.DescribeCdnWafDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a specific certificate by certificate ID.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- If a certificate is associated with a domain name but the certificate is not enabled, the result of this operation shows that the certificate does not exist.
//
// @param request - DescribeCertificateInfoByIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertificateInfoByIDResponse
func (client *Client) DescribeCertificateInfoByIDWithOptions(request *DescribeCertificateInfoByIDRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertificateInfoByIDResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCertificateInfoByID"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertificateInfoByIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a specific certificate by certificate ID.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- If a certificate is associated with a domain name but the certificate is not enabled, the result of this operation shows that the certificate does not exist.
//
// @param request - DescribeCertificateInfoByIDRequest
//
// @return DescribeCertificateInfoByIDResponse
func (client *Client) DescribeCertificateInfoByID(request *DescribeCertificateInfoByIDRequest) (_result *DescribeCertificateInfoByIDResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertificateInfoByIDResponse{}
	_body, _err := client.DescribeCertificateInfoByIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制查询域名采样率
//
// @param request - DescribeCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomDomainSampleRateResponse
func (client *Client) DescribeCustomDomainSampleRateWithOptions(request *DescribeCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomDomainSampleRate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制查询域名采样率
//
// @param request - DescribeCustomDomainSampleRateRequest
//
// @return DescribeCustomDomainSampleRateResponse
func (client *Client) DescribeCustomDomainSampleRate(request *DescribeCustomDomainSampleRateRequest) (_result *DescribeCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomDomainSampleRateResponse{}
	_body, _err := client.DescribeCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a custom logging configuration.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCustomLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomLogConfigResponse
func (client *Client) DescribeCustomLogConfigWithOptions(request *DescribeCustomLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomLogConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a custom logging configuration.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeCustomLogConfigRequest
//
// @return DescribeCustomLogConfigResponse
func (client *Client) DescribeCustomLogConfig(request *DescribeCustomLogConfigRequest) (_result *DescribeCustomLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomLogConfigResponse{}
	_body, _err := client.DescribeCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the average response time of one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 100 times per second per account.
//
// >	- You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - DescribeDomainAverageResponseTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainAverageResponseTimeResponse
func (client *Client) DescribeDomainAverageResponseTimeWithOptions(request *DescribeDomainAverageResponseTimeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainAverageResponseTimeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["DomainType"] = request.DomainType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeMerge) {
		query["TimeMerge"] = request.TimeMerge
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainAverageResponseTime"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainAverageResponseTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average response time of one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 100 times per second per account.
//
// >	- You can specify up to 500 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - DescribeDomainAverageResponseTimeRequest
//
// @return DescribeDomainAverageResponseTimeResponse
func (client *Client) DescribeDomainAverageResponseTime(request *DescribeDomainAverageResponseTimeRequest) (_result *DescribeDomainAverageResponseTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainAverageResponseTimeResponse{}
	_body, _err := client.DescribeDomainAverageResponseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth monitoring data for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 150 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainBpsDataResponse
func (client *Client) DescribeDomainBpsDataWithOptions(request *DescribeDomainBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainBpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth monitoring data for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 150 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainBpsDataRequest
//
// @return DescribeDomainBpsDataResponse
func (client *Client) DescribeDomainBpsData(request *DescribeDomainBpsDataRequest) (_result *DescribeDomainBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainBpsDataResponse{}
	_body, _err := client.DescribeDomainBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth data by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainBpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainBpsDataByLayerResponse
func (client *Client) DescribeDomainBpsDataByLayerWithOptions(request *DescribeDomainBpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainBpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainBpsDataByLayer"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainBpsDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth data by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainBpsDataByLayerRequest
//
// @return DescribeDomainBpsDataByLayerResponse
func (client *Client) DescribeDomainBpsDataByLayer(request *DescribeDomainBpsDataByLayerRequest) (_result *DescribeDomainBpsDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainBpsDataByLayerResponse{}
	_body, _err := client.DescribeDomainBpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data at a specified time for an accelerated domain.
//
// Description:
//
//	  The bandwidth is measured in bit/s.
//
//		- You can specify only one accelerated domain name in each request.
//
//		- The data is collected every 5 minutes.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDomainBpsDataByTimeStampRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainBpsDataByTimeStampResponse
func (client *Client) DescribeDomainBpsDataByTimeStampWithOptions(request *DescribeDomainBpsDataByTimeStampRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.IspNames) {
		query["IspNames"] = request.IspNames
	}

	if !dara.IsNil(request.LocationNames) {
		query["LocationNames"] = request.LocationNames
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainBpsDataByTimeStamp"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data at a specified time for an accelerated domain.
//
// Description:
//
//	  The bandwidth is measured in bit/s.
//
//		- You can specify only one accelerated domain name in each request.
//
//		- The data is collected every 5 minutes.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDomainBpsDataByTimeStampRequest
//
// @return DescribeDomainBpsDataByTimeStampResponse
func (client *Client) DescribeDomainBpsDataByTimeStamp(request *DescribeDomainBpsDataByTimeStampRequest) (_result *DescribeDomainBpsDataByTimeStampResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainBpsDataByTimeStampResponse{}
	_body, _err := client.DescribeDomainBpsDataByTimeStampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries log entries of rate limiting.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both parameters empty.
//
//		- You can specify up to 20 domain names in reach request. If you specify multiple domain names, separate them with commas (,).
//
//		- You can query data collected over the last 30 days.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainCcActivityLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainCcActivityLogResponse
func (client *Client) DescribeDomainCcActivityLogWithOptions(request *DescribeDomainCcActivityLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainCcActivityLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TriggerObject) {
		query["TriggerObject"] = request.TriggerObject
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainCcActivityLog"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainCcActivityLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries log entries of rate limiting.
//
// Description:
//
//	  If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range. You must set both parameters or leave both parameters empty.
//
//		- You can specify up to 20 domain names in reach request. If you specify multiple domain names, separate them with commas (,).
//
//		- You can query data collected over the last 30 days.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainCcActivityLogRequest
//
// @return DescribeDomainCcActivityLogResponse
func (client *Client) DescribeDomainCcActivityLog(request *DescribeDomainCcActivityLogRequest) (_result *DescribeDomainCcActivityLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainCcActivityLogResponse{}
	_body, _err := client.DescribeDomainCcActivityLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificate information of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainCertificateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainCertificateInfoResponse
func (client *Client) DescribeDomainCertificateInfoWithOptions(request *DescribeDomainCertificateInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainCertificateInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainCertificateInfo"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainCertificateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificate information of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainCertificateInfoRequest
//
// @return DescribeDomainCertificateInfoResponse
func (client *Client) DescribeDomainCertificateInfo(request *DescribeDomainCertificateInfoRequest) (_result *DescribeDomainCertificateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainCertificateInfoResponse{}
	_body, _err := client.DescribeDomainCertificateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects the CNAME for an accelerated domain name. You can check the resolution result to determine whether the CNAME is configured.
//
// @param request - DescribeDomainCnameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainCnameResponse
func (client *Client) DescribeDomainCnameWithOptions(request *DescribeDomainCnameRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainCnameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainCname"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainCnameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects the CNAME for an accelerated domain name. You can check the resolution result to determine whether the CNAME is configured.
//
// @param request - DescribeDomainCnameRequest
//
// @return DescribeDomainCnameResponse
func (client *Client) DescribeDomainCname(request *DescribeDomainCnameRequest) (_result *DescribeDomainCnameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainCnameResponse{}
	_body, _err := client.DescribeDomainCnameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom log configuration of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainCustomLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainCustomLogConfigResponse
func (client *Client) DescribeDomainCustomLogConfigWithOptions(request *DescribeDomainCustomLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainCustomLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainCustomLogConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainCustomLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom log configuration of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainCustomLogConfigRequest
//
// @return DescribeDomainCustomLogConfigResponse
func (client *Client) DescribeDomainCustomLogConfig(request *DescribeDomainCustomLogConfigRequest) (_result *DescribeDomainCustomLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainCustomLogConfigResponse{}
	_body, _err := client.DescribeDomainCustomLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The domain name that you want to query. You can specify multiple domain names and separate them with commas (,). You can specify at most 30 domain names in each call.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDomainDetailDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainDetailDataByLayerResponse
func (client *Client) DescribeDomainDetailDataByLayerWithOptions(request *DescribeDomainDetailDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainDetailDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainDetailDataByLayer"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainDetailDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The domain name that you want to query. You can specify multiple domain names and separate them with commas (,). You can specify at most 30 domain names in each call.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
// @param request - DescribeDomainDetailDataByLayerRequest
//
// @return DescribeDomainDetailDataByLayerResponse
func (client *Client) DescribeDomainDetailDataByLayer(request *DescribeDomainDetailDataByLayerRequest) (_result *DescribeDomainDetailDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainDetailDataByLayerResponse{}
	_body, _err := client.DescribeDomainDetailDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries byte hit ratios that are measured in percentage.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainHitRateDataResponse
func (client *Client) DescribeDomainHitRateDataWithOptions(request *DescribeDomainHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainHitRateData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries byte hit ratios that are measured in percentage.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHitRateDataRequest
//
// @return DescribeDomainHitRateDataResponse
func (client *Client) DescribeDomainHitRateData(request *DescribeDomainHitRateDataRequest) (_result *DescribeDomainHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainHitRateDataResponse{}
	_body, _err := client.DescribeDomainHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from an accelerated domain name. The data is collected every 5 minutes.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainHttpCodeDataResponse
func (client *Client) DescribeDomainHttpCodeDataWithOptions(request *DescribeDomainHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainHttpCodeData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from an accelerated domain name. The data is collected every 5 minutes.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHttpCodeDataRequest
//
// @return DescribeDomainHttpCodeDataResponse
func (client *Client) DescribeDomainHttpCodeData(request *DescribeDomainHttpCodeDataRequest) (_result *DescribeDomainHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries HTTP status codes by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// ### Time granularity
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHttpCodeDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainHttpCodeDataByLayerResponse
func (client *Client) DescribeDomainHttpCodeDataByLayerWithOptions(request *DescribeDomainHttpCodeDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainHttpCodeDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainHttpCodeDataByLayer"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries HTTP status codes by protocol.
//
// Description:
//
// You can call this operation up to 20 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// ### Time granularity
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainHttpCodeDataByLayerRequest
//
// @return DescribeDomainHttpCodeDataByLayerResponse
func (client *Client) DescribeDomainHttpCodeDataByLayer(request *DescribeDomainHttpCodeDataByLayerRequest) (_result *DescribeDomainHttpCodeDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainHttpCodeDataByLayerResponse{}
	_body, _err := client.DescribeDomainHttpCodeDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proportions of data usage of different Internet service providers (ISPs). Data is collected every day. You can query data collected in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set StartTime or EndTime, the request returns the data collected in the last 24 hours. If you set both StartTime and EndTime, the request returns the data collected within the specified time range.
//
// >	- This operation queries proportions of data usage of different ISPs for only a specific accelerated domain name, or for all accelerated domain names in your Alibaba Cloud account.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainISPDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainISPDataResponse
func (client *Client) DescribeDomainISPDataWithOptions(request *DescribeDomainISPDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainISPDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainISPData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of data usage of different Internet service providers (ISPs). Data is collected every day. You can query data collected in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set StartTime or EndTime, the request returns the data collected in the last 24 hours. If you set both StartTime and EndTime, the request returns the data collected within the specified time range.
//
// >	- This operation queries proportions of data usage of different ISPs for only a specific accelerated domain name, or for all accelerated domain names in your Alibaba Cloud account.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainISPDataRequest
//
// @return DescribeDomainISPDataResponse
func (client *Client) DescribeDomainISPData(request *DescribeDomainISPDataRequest) (_result *DescribeDomainISPDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainISPDataResponse{}
	_body, _err := client.DescribeDomainISPDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the 95th percentile bandwidth data of a domain name.
//
// Description:
//
// *You can use one of the following methods to query data:**
//
//   - If you specify the StartTime and EndTime parameters and the time range that is specified by these parameters is less than or equal to 24 hours, the 95th percentile bandwidth data on the day of the start time is returned. If the time range that is specified by these parameters is more than 24 hours, the 95th percentile bandwidth data in the month of the start time is returned.
//
//   - If you specify the TimePoint and Cycle parameters, the 95th percentile bandwidth data of the cycle is returned.
//
//   - If you specify the StartTime, EndTime, and Cycle parameters, the 95th percentile bandwidth data of the cycle is returned.
//
// If you do not use one of the methods, the 95th percentile bandwidth data of the previous 24 hours is returned by default.
//
//   - Maximum time range to query: 90 days
//
//   - Minimum data granularity to query: 1 day
//
//   - Historical data available: 90 days
//
// - You can call this operation up to 100 times per second per account.
//
// - The unit of the bandwidth data returned is bit/s.
//
// @param request - DescribeDomainMax95BpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainMax95BpsDataResponse
func (client *Client) DescribeDomainMax95BpsDataWithOptions(request *DescribeDomainMax95BpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainMax95BpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cycle) {
		query["Cycle"] = request.Cycle
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimePoint) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainMax95BpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainMax95BpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the 95th percentile bandwidth data of a domain name.
//
// Description:
//
// *You can use one of the following methods to query data:**
//
//   - If you specify the StartTime and EndTime parameters and the time range that is specified by these parameters is less than or equal to 24 hours, the 95th percentile bandwidth data on the day of the start time is returned. If the time range that is specified by these parameters is more than 24 hours, the 95th percentile bandwidth data in the month of the start time is returned.
//
//   - If you specify the TimePoint and Cycle parameters, the 95th percentile bandwidth data of the cycle is returned.
//
//   - If you specify the StartTime, EndTime, and Cycle parameters, the 95th percentile bandwidth data of the cycle is returned.
//
// If you do not use one of the methods, the 95th percentile bandwidth data of the previous 24 hours is returned by default.
//
//   - Maximum time range to query: 90 days
//
//   - Minimum data granularity to query: 1 day
//
//   - Historical data available: 90 days
//
// - You can call this operation up to 100 times per second per account.
//
// - The unit of the bandwidth data returned is bit/s.
//
// @param request - DescribeDomainMax95BpsDataRequest
//
// @return DescribeDomainMax95BpsDataResponse
func (client *Client) DescribeDomainMax95BpsData(request *DescribeDomainMax95BpsDataRequest) (_result *DescribeDomainMax95BpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainMax95BpsDataResponse{}
	_body, _err := client.DescribeDomainMax95BpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic data and the number of requests for multiple accelerated domain names at a time.
//
// Description:
//
//	  If you do not set StartTime or EndTime, data collected within the last 10 minutes is queried.
//
//		- The maximum interval between StartTime and EndTime is 1 hour.
//
//		- You can query data within the last 90 days.
//
//		- You can query the traffic data and the number of requests for accelerated domain names that are deleted.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainMultiUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainMultiUsageDataResponse
func (client *Client) DescribeDomainMultiUsageDataWithOptions(request *DescribeDomainMultiUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainMultiUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainMultiUsageData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainMultiUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic data and the number of requests for multiple accelerated domain names at a time.
//
// Description:
//
//	  If you do not set StartTime or EndTime, data collected within the last 10 minutes is queried.
//
//		- The maximum interval between StartTime and EndTime is 1 hour.
//
//		- You can query data within the last 90 days.
//
//		- You can query the traffic data and the number of requests for accelerated domain names that are deleted.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainMultiUsageDataRequest
//
// @return DescribeDomainMultiUsageDataResponse
func (client *Client) DescribeDomainMultiUsageData(request *DescribeDomainMultiUsageDataRequest) (_result *DescribeDomainMultiUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainMultiUsageDataResponse{}
	_body, _err := client.DescribeDomainMultiUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries monitoring data including the amount of network traffic and the number of visits by directory.
//
// Description:
//
//	  This operation is available only to users that are on the whitelist. If the daily peak bandwidth value of your workloads reaches 10 Gbit/s, you can [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex) to apply to be included in the whitelist.
//
//		- You can call this API operation up to 6,000 times per second per account.
//
//		- Data collection by directory is available only to specified domain names within your Alibaba Cloud account. It cannot be enabled for all domain names within your Alibaba Cloud account.
//
//		- The average size of the files that belong to the domain name must be larger than 1 MB.
//
//		- The number of directories specified for a single domain name cannot exceed 100. If the number of directories exceeds 100, the data accuracy reduces.
//
//		- If you do not set StartTime or EndTime, data collected within the last 24 hours is queried. If you set both StartTime and EndTime, data within the specified time range is queried.
//
//		- You can query data collected within the last 30 days.
//
// @param request - DescribeDomainPathDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainPathDataResponse
func (client *Client) DescribeDomainPathDataWithOptions(request *DescribeDomainPathDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainPathDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainPathData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainPathDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries monitoring data including the amount of network traffic and the number of visits by directory.
//
// Description:
//
//	  This operation is available only to users that are on the whitelist. If the daily peak bandwidth value of your workloads reaches 10 Gbit/s, you can [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex) to apply to be included in the whitelist.
//
//		- You can call this API operation up to 6,000 times per second per account.
//
//		- Data collection by directory is available only to specified domain names within your Alibaba Cloud account. It cannot be enabled for all domain names within your Alibaba Cloud account.
//
//		- The average size of the files that belong to the domain name must be larger than 1 MB.
//
//		- The number of directories specified for a single domain name cannot exceed 100. If the number of directories exceeds 100, the data accuracy reduces.
//
//		- If you do not set StartTime or EndTime, data collected within the last 24 hours is queried. If you set both StartTime and EndTime, data within the specified time range is queried.
//
//		- You can query data collected within the last 30 days.
//
// @param request - DescribeDomainPathDataRequest
//
// @return DescribeDomainPathDataResponse
func (client *Client) DescribeDomainPathData(request *DescribeDomainPathDataRequest) (_result *DescribeDomainPathDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainPathDataResponse{}
	_body, _err := client.DescribeDomainPathDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the page view (PV) data of an accelerated domain name. The data is collected at an interval of 1 hour. You can query data in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainPvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainPvDataResponse
func (client *Client) DescribeDomainPvDataWithOptions(request *DescribeDomainPvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainPvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainPvData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainPvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the page view (PV) data of an accelerated domain name. The data is collected at an interval of 1 hour. You can query data in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 50 times per second per account.
//
// @param request - DescribeDomainPvDataRequest
//
// @return DescribeDomainPvDataResponse
func (client *Client) DescribeDomainPvData(request *DescribeDomainPvDataRequest) (_result *DescribeDomainPvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainPvDataResponse{}
	_body, _err := client.DescribeDomainPvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) for an accelerated domain name. The data is collected every 5 minutes. You can query data collected within the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainQpsDataResponse
func (client *Client) DescribeDomainQpsDataWithOptions(request *DescribeDomainQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainQpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) for an accelerated domain name. The data is collected every 5 minutes. You can query data collected within the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainQpsDataRequest
//
// @return DescribeDomainQpsDataResponse
func (client *Client) DescribeDomainQpsData(request *DescribeDomainQpsDataRequest) (_result *DescribeDomainQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainQpsDataResponse{}
	_body, _err := client.DescribeDomainQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) at a specific layer for one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainQpsDataByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainQpsDataByLayerResponse
func (client *Client) DescribeDomainQpsDataByLayerWithOptions(request *DescribeDomainQpsDataByLayerRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainQpsDataByLayerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.Layer) {
		query["Layer"] = request.Layer
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainQpsDataByLayer"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainQpsDataByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) at a specific layer for one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// You can call this operation up to 20 times per second per user.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainQpsDataByLayerRequest
//
// @return DescribeDomainQpsDataByLayerResponse
func (client *Client) DescribeDomainQpsDataByLayer(request *DescribeDomainQpsDataByLayerRequest) (_result *DescribeDomainQpsDataByLayerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainQpsDataByLayerResponse{}
	_body, _err := client.DescribeDomainQpsDataByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data about one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeBpsDataResponse
func (client *Client) DescribeDomainRealTimeBpsDataWithOptions(request *DescribeDomainRealTimeBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeBpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth data about one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity*	- The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeBpsDataRequest
//
// @return DescribeDomainRealTimeBpsDataResponse
func (client *Client) DescribeDomainRealTimeBpsData(request *DescribeDomainRealTimeBpsDataRequest) (_result *DescribeDomainRealTimeBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeBpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the byte hit ratios of accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeByteHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeDomainRealTimeByteHitRateDataWithOptions(request *DescribeDomainRealTimeByteHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeByteHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeByteHitRateData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the byte hit ratios of accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the byte hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeByteHitRateDataRequest
//
// @return DescribeDomainRealTimeByteHitRateDataResponse
func (client *Client) DescribeDomainRealTimeByteHitRateData(request *DescribeDomainRealTimeByteHitRateDataRequest) (_result *DescribeDomainRealTimeByteHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeByteHitRateDataResponse{}
	_body, _err := client.DescribeDomainRealTimeByteHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time monitoring data for a domain name.
//
// Description:
//
//	  You can query data in the last seven days. Data is collected every minute.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainRealTimeDetailDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeDetailDataResponse
func (client *Client) DescribeDomainRealTimeDetailDataWithOptions(request *DescribeDomainRealTimeDetailDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeDetailDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeDetailData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeDetailDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time monitoring data for a domain name.
//
// Description:
//
//	  You can query data in the last seven days. Data is collected every minute.
//
//		- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainRealTimeDetailDataRequest
//
// @return DescribeDomainRealTimeDetailDataResponse
func (client *Client) DescribeDomainRealTimeDetailData(request *DescribeDomainRealTimeDetailDataRequest) (_result *DescribeDomainRealTimeDetailDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeDetailDataResponse{}
	_body, _err := client.DescribeDomainRealTimeDetailDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from an accelerated domain name.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeDomainRealTimeHttpCodeDataWithOptions(request *DescribeDomainRealTimeHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeHttpCodeData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned from an accelerated domain name.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeHttpCodeDataRequest
//
// @return DescribeDomainRealTimeHttpCodeDataResponse
func (client *Client) DescribeDomainRealTimeHttpCodeData(request *DescribeDomainRealTimeHttpCodeDataRequest) (_result *DescribeDomainRealTimeHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainRealTimeHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeQpsDataResponse
func (client *Client) DescribeDomainRealTimeQpsDataWithOptions(request *DescribeDomainRealTimeQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeQpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeQpsDataRequest
//
// @return DescribeDomainRealTimeQpsDataResponse
func (client *Client) DescribeDomainRealTimeQpsData(request *DescribeDomainRealTimeQpsDataRequest) (_result *DescribeDomainRealTimeQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeQpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request hit ratios for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
//   - By default, requests in the Go programming language use the POST request method. You must manually change the request method to GET by declaring: request.Method="GET".
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the request hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeDomainRealTimeReqHitRateDataWithOptions(request *DescribeDomainRealTimeReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeReqHitRateData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit ratios for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
//   - By default, requests in the Go programming language use the POST request method. You must manually change the request method to GET by declaring: request.Method="GET".
//
//   - The network traffic destined for different domain names may be redirected to the same origin server. Therefore, the request hit ratios may be inaccurate. The accuracy of query results is based on the actual configurations.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeReqHitRateDataRequest
//
// @return DescribeDomainRealTimeReqHitRateDataResponse
func (client *Client) DescribeDomainRealTimeReqHitRateData(request *DescribeDomainRealTimeReqHitRateDataRequest) (_result *DescribeDomainRealTimeReqHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeReqHitRateDataResponse{}
	_body, _err := client.DescribeDomainRealTimeReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries origin bandwidth data for accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeSrcBpsDataResponse
func (client *Client) DescribeDomainRealTimeSrcBpsDataWithOptions(request *DescribeDomainRealTimeSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeSrcBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeSrcBpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin bandwidth data for accelerated domain names.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcBpsDataRequest
//
// @return DescribeDomainRealTimeSrcBpsDataResponse
func (client *Client) DescribeDomainRealTimeSrcBpsData(request *DescribeDomainRealTimeSrcBpsDataRequest) (_result *DescribeDomainRealTimeSrcBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcBpsDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned during back-to-origin routing.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeSrcHttpCodeDataResponse
func (client *Client) DescribeDomainRealTimeSrcHttpCodeDataWithOptions(request *DescribeDomainRealTimeSrcHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeSrcHttpCodeData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number and proportions of HTTP status codes returned during back-to-origin routing.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you set both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcHttpCodeDataRequest
//
// @return DescribeDomainRealTimeSrcHttpCodeDataResponse
func (client *Client) DescribeDomainRealTimeSrcHttpCodeData(request *DescribeDomainRealTimeSrcHttpCodeDataRequest) (_result *DescribeDomainRealTimeSrcHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of back-to-origin traffic for one or more specified accelerated domains. The data is collected every minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour by default. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeSrcTrafficDataResponse
func (client *Client) DescribeDomainRealTimeSrcTrafficDataWithOptions(request *DescribeDomainRealTimeSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeSrcTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeSrcTrafficData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of back-to-origin traffic for one or more specified accelerated domains. The data is collected every minute.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour by default. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeSrcTrafficDataRequest
//
// @return DescribeDomainRealTimeSrcTrafficDataResponse
func (client *Client) DescribeDomainRealTimeSrcTrafficData(request *DescribeDomainRealTimeSrcTrafficDataRequest) (_result *DescribeDomainRealTimeSrcTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeSrcTrafficDataResponse{}
	_body, _err := client.DescribeDomainRealTimeSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of network traffic for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 50 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealTimeTrafficDataResponse
func (client *Client) DescribeDomainRealTimeTrafficDataWithOptions(request *DescribeDomainRealTimeTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealTimeTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealTimeTrafficData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealTimeTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of network traffic for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 50 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last hour. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity varies with the time range specified by the StartTime and EndTime parameters. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |1 minute|1 hour|7 days|5 minutes|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// @param request - DescribeDomainRealTimeTrafficDataRequest
//
// @return DescribeDomainRealTimeTrafficDataResponse
func (client *Client) DescribeDomainRealTimeTrafficData(request *DescribeDomainRealTimeTrafficDataRequest) (_result *DescribeDomainRealTimeTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealTimeTrafficDataResponse{}
	_body, _err := client.DescribeDomainRealTimeTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time log delivery information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRealtimeLogDeliveryResponse
func (client *Client) DescribeDomainRealtimeLogDeliveryWithOptions(request *DescribeDomainRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time log delivery information about an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainRealtimeLogDeliveryRequest
//
// @return DescribeDomainRealtimeLogDeliveryResponse
func (client *Client) DescribeDomainRealtimeLogDelivery(request *DescribeDomainRealtimeLogDeliveryRequest) (_result *DescribeDomainRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRealtimeLogDeliveryResponse{}
	_body, _err := client.DescribeDomainRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the geographic distribution of users for a domain name. The data is collected at an interval of one day. You can query the data in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you not use this operation because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not specify the **StartTime*	- or **EndTime*	- parameter, data collected within the last **24*	- hours is queried. If you specify both the **StartTime*	- and **EndTime*	- parameters, data collected within the specified time range is queried.
//
// >	- There is delay in data collection. If you want to query data collected within the last day, we recommend that you query the data on the next day.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainRegionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainRegionDataResponse
func (client *Client) DescribeDomainRegionDataWithOptions(request *DescribeDomainRegionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainRegionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainRegionData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the geographic distribution of users for a domain name. The data is collected at an interval of one day. You can query the data in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you not use this operation because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not specify the **StartTime*	- or **EndTime*	- parameter, data collected within the last **24*	- hours is queried. If you specify both the **StartTime*	- and **EndTime*	- parameters, data collected within the specified time range is queried.
//
// >	- There is delay in data collection. If you want to query data collected within the last day, we recommend that you query the data on the next day.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainRegionDataRequest
//
// @return DescribeDomainRegionDataResponse
func (client *Client) DescribeDomainRegionData(request *DescribeDomainRegionDataRequest) (_result *DescribeDomainRegionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainRegionDataResponse{}
	_body, _err := client.DescribeDomainRegionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request hit ratio in percentage.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainReqHitRateDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainReqHitRateDataResponse
func (client *Client) DescribeDomainReqHitRateDataWithOptions(request *DescribeDomainReqHitRateDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainReqHitRateData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request hit ratio in percentage.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainReqHitRateDataRequest
//
// @return DescribeDomainReqHitRateDataResponse
func (client *Client) DescribeDomainReqHitRateData(request *DescribeDomainReqHitRateDataRequest) (_result *DescribeDomainReqHitRateDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainReqHitRateDataResponse{}
	_body, _err := client.DescribeDomainReqHitRateDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries bandwidth monitoring data of requests that are redirected to origin servers for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcBpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSrcBpsDataResponse
func (client *Client) DescribeDomainSrcBpsDataWithOptions(request *DescribeDomainSrcBpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSrcBpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries bandwidth monitoring data of requests that are redirected to origin servers for one or more accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not specify the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you specify both the StartTime and EndTime parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the time range to query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcBpsDataRequest
//
// @return DescribeDomainSrcBpsDataResponse
func (client *Client) DescribeDomainSrcBpsData(request *DescribeDomainSrcBpsDataRequest) (_result *DescribeDomainSrcBpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSrcBpsDataResponse{}
	_body, _err := client.DescribeDomainSrcBpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proportions of HTTP status codes that are returned during back-to-origin routing.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcHttpCodeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSrcHttpCodeDataResponse
func (client *Client) DescribeDomainSrcHttpCodeDataWithOptions(request *DescribeDomainSrcHttpCodeDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSrcHttpCodeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSrcHttpCodeData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSrcHttpCodeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proportions of HTTP status codes that are returned during back-to-origin routing.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter varies with the maximum time range per query. The following table describes the time period within which historical data is available and the data delay.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcHttpCodeDataRequest
//
// @return DescribeDomainSrcHttpCodeDataResponse
func (client *Client) DescribeDomainSrcHttpCodeData(request *DescribeDomainSrcHttpCodeDataRequest) (_result *DescribeDomainSrcHttpCodeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSrcHttpCodeDataResponse{}
	_body, _err := client.DescribeDomainSrcHttpCodeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) that are sent to the origin server. You can query data collected in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// ### Time granularity
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSrcQpsDataResponse
func (client *Client) DescribeDomainSrcQpsDataWithOptions(request *DescribeDomainSrcQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSrcQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSrcQpsData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSrcQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of queries per second (QPS) that are sent to the origin server. You can query data collected in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// ### Time granularity
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcQpsDataRequest
//
// @return DescribeDomainSrcQpsDataResponse
func (client *Client) DescribeDomainSrcQpsData(request *DescribeDomainSrcQpsDataRequest) (_result *DescribeDomainSrcQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSrcQpsDataResponse{}
	_body, _err := client.DescribeDomainSrcQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries frequently requested origin URLs of one or more accelerated domain names.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- The data is collected at an interval of 5 minutes.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainSrcTopUrlVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSrcTopUrlVisitResponse
func (client *Client) DescribeDomainSrcTopUrlVisitWithOptions(request *DescribeDomainSrcTopUrlVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSrcTopUrlVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSrcTopUrlVisit"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSrcTopUrlVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries frequently requested origin URLs of one or more accelerated domain names.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- The data is collected at an interval of 5 minutes.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainSrcTopUrlVisitRequest
//
// @return DescribeDomainSrcTopUrlVisitResponse
func (client *Client) DescribeDomainSrcTopUrlVisit(request *DescribeDomainSrcTopUrlVisitRequest) (_result *DescribeDomainSrcTopUrlVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSrcTopUrlVisitResponse{}
	_body, _err := client.DescribeDomainSrcTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries origin traffic for one or more specified accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainSrcTrafficDataResponse
func (client *Client) DescribeDomainSrcTrafficDataWithOptions(request *DescribeDomainSrcTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainSrcTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainSrcTrafficData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainSrcTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin traffic for one or more specified accelerated domain names.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainSrcTrafficDataRequest
//
// @return DescribeDomainSrcTrafficDataResponse
func (client *Client) DescribeDomainSrcTrafficData(request *DescribeDomainSrcTrafficDataRequest) (_result *DescribeDomainSrcTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainSrcTrafficDataResponse{}
	_body, _err := client.DescribeDomainSrcTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries client IP addresses that are ranked by the number of requests or the amount of network traffic within a specific time range for one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- Data is collected every hour.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopClientIpVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopClientIpVisitResponse
func (client *Client) DescribeDomainTopClientIpVisitWithOptions(request *DescribeDomainTopClientIpVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopClientIpVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopClientIpVisit"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopClientIpVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries client IP addresses that are ranked by the number of requests or the amount of network traffic within a specific time range for one or more accelerated domain names. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature to for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- Data is collected every hour.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopClientIpVisitRequest
//
// @return DescribeDomainTopClientIpVisitResponse
func (client *Client) DescribeDomainTopClientIpVisit(request *DescribeDomainTopClientIpVisitRequest) (_result *DescribeDomainTopClientIpVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopClientIpVisitResponse{}
	_body, _err := client.DescribeDomainTopClientIpVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries frequently requested web pages of one or more accelerated domain names on a specified day and sorts the web pages. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature or [ship real-time logs in Log Service](https://help.aliyun.com/document_detail/440145.html) to analyze data.
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - Data is collected at an interval of five minutes.
//
//   - You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopReferVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopReferVisitResponse
func (client *Client) DescribeDomainTopReferVisitWithOptions(request *DescribeDomainTopReferVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopReferVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopReferVisit"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopReferVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries frequently requested web pages of one or more accelerated domain names on a specified day and sorts the web pages. You can query data collected within the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature or [ship real-time logs in Log Service](https://help.aliyun.com/document_detail/440145.html) to analyze data.
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - Data is collected at an interval of five minutes.
//
//   - You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopReferVisitRequest
//
// @return DescribeDomainTopReferVisitResponse
func (client *Client) DescribeDomainTopReferVisit(request *DescribeDomainTopReferVisitRequest) (_result *DescribeDomainTopReferVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopReferVisitResponse{}
	_body, _err := client.DescribeDomainTopReferVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries top 100 frequently requested URLs of an accelerated domain name within a specified time range.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can query data collected in the last 90 days.
//
// >	- You can specify only one domain name in each call.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopUrlVisitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTopUrlVisitResponse
func (client *Client) DescribeDomainTopUrlVisitWithOptions(request *DescribeDomainTopUrlVisitRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTopUrlVisitResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTopUrlVisit"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTopUrlVisitResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries top 100 frequently requested URLs of an accelerated domain name within a specified time range.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can query data collected in the last 90 days.
//
// >	- You can specify only one domain name in each call.
//
// >	- You can call this operation up to 10 times per second per account.
//
// @param request - DescribeDomainTopUrlVisitRequest
//
// @return DescribeDomainTopUrlVisitResponse
func (client *Client) DescribeDomainTopUrlVisit(request *DescribeDomainTopUrlVisitRequest) (_result *DescribeDomainTopUrlVisitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTopUrlVisitResponse{}
	_body, _err := client.DescribeDomainTopUrlVisitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries network traffic for one or more accelerated domain names. You can query data that is collected in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366|04:00 on the next day|
//
// @param request - DescribeDomainTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainTrafficDataResponse
func (client *Client) DescribeDomainTrafficDataWithOptions(request *DescribeDomainTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IspNameEn) {
		query["IspNameEn"] = request.IspNameEn
	}

	if !dara.IsNil(request.LocationNameEn) {
		query["LocationNameEn"] = request.LocationNameEn
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainTrafficData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network traffic for one or more accelerated domain names. You can query data that is collected in the last 90 days.
//
// Description:
//
// You can call this operation up to 100 times per second per account.
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
// **Time granularity**
//
// The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|366 days|366|04:00 on the next day|
//
// @param request - DescribeDomainTrafficDataRequest
//
// @return DescribeDomainTrafficDataResponse
func (client *Client) DescribeDomainTrafficData(request *DescribeDomainTrafficDataRequest) (_result *DescribeDomainTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainTrafficDataResponse{}
	_body, _err := client.DescribeDomainTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource usage data of specific domain names in a specified billable region.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|90 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainUsageDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainUsageDataResponse
func (client *Client) DescribeDomainUsageDataWithOptions(request *DescribeDomainUsageDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainUsageDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.DataProtocol) {
		query["DataProtocol"] = request.DataProtocol
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Field) {
		query["Field"] = request.Field
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainUsageData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainUsageDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource usage data of specific domain names in a specified billable region.
//
// Description:
//
// You can call this operation up to 10 times per second per account.
//
//   - The time granularity supported by the Interval parameter, the maximum time period within which historical data is available, and the data delay vary with the maximum time range per query, as described in the following table.
//
// |Time granularity|Maximum time range per query|Historical data available|Data delay|
//
// |---|---|---|---|
//
// |5 minutes|3 days|93 days|15 minutes|
//
// |1 hour|31 days|186 days|4 hours|
//
// |1 day|90 days|366 days|04:00 on the next day|
//
// @param request - DescribeDomainUsageDataRequest
//
// @return DescribeDomainUsageDataResponse
func (client *Client) DescribeDomainUsageData(request *DescribeDomainUsageDataRequest) (_result *DescribeDomainUsageDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainUsageDataResponse{}
	_body, _err := client.DescribeDomainUsageDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the unique visitor (UV) data of an accelerated domain name. Data is collected every hour. You can query data collected in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - You can specify only one accelerated domain name or all accelerated domain names in your Alibaba Cloud account.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainUvDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainUvDataResponse
func (client *Client) DescribeDomainUvDataWithOptions(request *DescribeDomainUvDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainUvDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainUvData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the unique visitor (UV) data of an accelerated domain name. Data is collected every hour. You can query data collected in the last 90 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// >
//
//   - If you do not set the StartTime or EndTime parameter, the request returns the data collected in the last 24 hours. If you set both these parameters, the request returns the data collected within the specified time range.
//
//   - You can specify only one accelerated domain name or all accelerated domain names in your Alibaba Cloud account.
//
//   - You can call this operation up to 100 times per second per account.
//
// @param request - DescribeDomainUvDataRequest
//
// @return DescribeDomainUvDataResponse
func (client *Client) DescribeDomainUvData(request *DescribeDomainUvDataRequest) (_result *DescribeDomainUvDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainUvDataResponse{}
	_body, _err := client.DescribeDomainUvDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the verification content of a domain name, including the host record and record value.
//
// Description:
//
// You can call this operation to query the verification content of an accelerated domain name based on whether the global resource plan is enabled.
//
// @param request - DescribeDomainVerifyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainVerifyDataResponse
func (client *Client) DescribeDomainVerifyDataWithOptions(request *DescribeDomainVerifyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainVerifyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GlobalResourcePlan) {
		query["GlobalResourcePlan"] = request.GlobalResourcePlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainVerifyData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainVerifyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the verification content of a domain name, including the host record and record value.
//
// Description:
//
// You can call this operation to query the verification content of an accelerated domain name based on whether the global resource plan is enabled.
//
// @param request - DescribeDomainVerifyDataRequest
//
// @return DescribeDomainVerifyDataResponse
func (client *Client) DescribeDomainVerifyData(request *DescribeDomainVerifyDataRequest) (_result *DescribeDomainVerifyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainVerifyDataResponse{}
	_body, _err := client.DescribeDomainVerifyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by origin server.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDomainsBySourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsBySourceResponse
func (client *Client) DescribeDomainsBySourceWithOptions(request *DescribeDomainsBySourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainsBySourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainsBySource"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries accelerated domain names by origin server.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeDomainsBySourceRequest
//
// @return DescribeDomainsBySourceResponse
func (client *Client) DescribeDomainsBySource(request *DescribeDomainsBySourceRequest) (_result *DescribeDomainsBySourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainsBySourceResponse{}
	_body, _err := client.DescribeDomainsBySourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of an accelerated domain name. Data is collected every day. You can query data collected within the last 90 days.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- If you do not set StartTime or EndTime, data within the last 24 hours is queried. If you set both StartTime and EndTime, data within the specified time range is queried.
//
//		- You can query the monitoring data of a specific accelerated domain name or all accelerated domain names that belong to your Alibaba Cloud account.
//
// @param request - DescribeDomainsUsageByDayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsUsageByDayResponse
func (client *Client) DescribeDomainsUsageByDayWithOptions(request *DescribeDomainsUsageByDayRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainsUsageByDayResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainsUsageByDay"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring data of an accelerated domain name. Data is collected every day. You can query data collected within the last 90 days.
//
// Description:
//
//	  You can call this operation up to 10 times per second per account.
//
//		- If you do not set StartTime or EndTime, data within the last 24 hours is queried. If you set both StartTime and EndTime, data within the specified time range is queried.
//
//		- You can query the monitoring data of a specific accelerated domain name or all accelerated domain names that belong to your Alibaba Cloud account.
//
// @param request - DescribeDomainsUsageByDayRequest
//
// @return DescribeDomainsUsageByDayResponse
func (client *Client) DescribeDomainsUsageByDay(request *DescribeDomainsUsageByDayRequest) (_result *DescribeDomainsUsageByDayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainsUsageByDayResponse{}
	_body, _err := client.DescribeDomainsUsageByDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution errors of a script in EdgeScript (ES).
//
// Description:
//
//	You can call this operation up to 30 times per second per account.
//
// @param request - DescribeEsExceptionDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEsExceptionDataResponse
func (client *Client) DescribeEsExceptionDataWithOptions(request *DescribeEsExceptionDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEsExceptionDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEsExceptionData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEsExceptionDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution errors of a script in EdgeScript (ES).
//
// Description:
//
//	You can call this operation up to 30 times per second per account.
//
// @param request - DescribeEsExceptionDataRequest
//
// @return DescribeEsExceptionDataResponse
func (client *Client) DescribeEsExceptionData(request *DescribeEsExceptionDataRequest) (_result *DescribeEsExceptionDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEsExceptionDataResponse{}
	_body, _err := client.DescribeEsExceptionDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status of scripts in EdgeScript (ES).
//
// Description:
//
//	You can call this operation up to 30 times per second per account.
//
// @param request - DescribeEsExecuteDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEsExecuteDataResponse
func (client *Client) DescribeEsExecuteDataWithOptions(request *DescribeEsExecuteDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeEsExecuteDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeEsExecuteData"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEsExecuteDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of scripts in EdgeScript (ES).
//
// Description:
//
//	You can call this operation up to 30 times per second per account.
//
// @param request - DescribeEsExecuteDataRequest
//
// @return DescribeEsExecuteDataResponse
func (client *Client) DescribeEsExecuteData(request *DescribeEsExecuteDataRequest) (_result *DescribeEsExecuteDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeEsExecuteDataResponse{}
	_body, _err := client.DescribeEsExecuteDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specified Function Compute trigger.
//
// @param request - DescribeFCTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFCTriggerResponse
func (client *Client) DescribeFCTriggerWithOptions(request *DescribeFCTriggerRequest, runtime *dara.RuntimeOptions) (_result *DescribeFCTriggerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFCTrigger"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFCTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specified Function Compute trigger.
//
// @param request - DescribeFCTriggerRequest
//
// @return DescribeFCTriggerResponse
func (client *Client) DescribeFCTrigger(request *DescribeFCTriggerRequest) (_result *DescribeFCTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFCTriggerResponse{}
	_body, _err := client.DescribeFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a specified IP address is the IP address of a CDN point of presence (POP).
//
// @param request - DescribeIpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpInfoResponse
func (client *Client) DescribeIpInfoWithOptions(request *DescribeIpInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpInfo"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a specified IP address is the IP address of a CDN point of presence (POP).
//
// @param request - DescribeIpInfoRequest
//
// @return DescribeIpInfoResponse
func (client *Client) DescribeIpInfo(request *DescribeIpInfoRequest) (_result *DescribeIpInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpInfoResponse{}
	_body, _err := client.DescribeIpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of IP addresses of points of presence (POPs). The status of an IP address of a POP indicates whether content delivery acceleration is supported by the POP.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeIpStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpStatusResponse
func (client *Client) DescribeIpStatusWithOptions(request *DescribeIpStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpStatus"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of IP addresses of points of presence (POPs). The status of an IP address of a POP indicates whether content delivery acceleration is supported by the POP.
//
// Description:
//
// > You can call this operation up to 50 times per second per account.
//
// @param request - DescribeIpStatusRequest
//
// @return DescribeIpStatusResponse
func (client *Client) DescribeIpStatus(request *DescribeIpStatusRequest) (_result *DescribeIpStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpStatusResponse{}
	_body, _err := client.DescribeIpStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual IP addresses (VIPs) of L2 CDN points of presence (POPs) for a specific domain name.
//
// Description:
//
//	  This operation is available only to users whose daily peak bandwidth value is higher than 1 Gbit/s. If you meet this requirement, you can [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex) to apply for permissions to use this operation.
//
//		- You can call this operation up to 40 times per second per account.
//
// @param request - DescribeL2VipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeL2VipsByDomainResponse
func (client *Client) DescribeL2VipsByDomainWithOptions(request *DescribeL2VipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeL2VipsByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeL2VipsByDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeL2VipsByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual IP addresses (VIPs) of L2 CDN points of presence (POPs) for a specific domain name.
//
// Description:
//
//	  This operation is available only to users whose daily peak bandwidth value is higher than 1 Gbit/s. If you meet this requirement, you can [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex) to apply for permissions to use this operation.
//
//		- You can call this operation up to 40 times per second per account.
//
// @param request - DescribeL2VipsByDomainRequest
//
// @return DescribeL2VipsByDomainResponse
func (client *Client) DescribeL2VipsByDomain(request *DescribeL2VipsByDomainRequest) (_result *DescribeL2VipsByDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeL2VipsByDomainResponse{}
	_body, _err := client.DescribeL2VipsByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the prefetch details of a task, including the prefetch progress of all resources in the task. Only users who are included in the whitelist can use this operation. You can contact your business manager to apply for the whitelist.
//
// Description:
//
//	  You can query data within the last 3 days.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribePreloadDetailByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePreloadDetailByIdResponse
func (client *Client) DescribePreloadDetailByIdWithOptions(request *DescribePreloadDetailByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribePreloadDetailByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePreloadDetailById"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePreloadDetailByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prefetch details of a task, including the prefetch progress of all resources in the task. Only users who are included in the whitelist can use this operation. You can contact your business manager to apply for the whitelist.
//
// Description:
//
//	  You can query data within the last 3 days.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribePreloadDetailByIdRequest
//
// @return DescribePreloadDetailByIdResponse
func (client *Client) DescribePreloadDetailById(request *DescribePreloadDetailByIdRequest) (_result *DescribePreloadDetailByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePreloadDetailByIdResponse{}
	_body, _err := client.DescribePreloadDetailByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bandwidth values by Internet service provider (ISP) and region.
//
// Description:
//
//	  The data is collected every 5 minutes.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - DescribeRangeDataByLocateAndIspServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRangeDataByLocateAndIspServiceResponse
func (client *Client) DescribeRangeDataByLocateAndIspServiceWithOptions(request *DescribeRangeDataByLocateAndIspServiceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRangeDataByLocateAndIspServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainNames) {
		query["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IspNames) {
		query["IspNames"] = request.IspNames
	}

	if !dara.IsNil(request.LocationNames) {
		query["LocationNames"] = request.LocationNames
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRangeDataByLocateAndIspService"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRangeDataByLocateAndIspServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bandwidth values by Internet service provider (ISP) and region.
//
// Description:
//
//	  The data is collected every 5 minutes.
//
//		- You can call this operation up to 20 times per second per account.
//
// @param request - DescribeRangeDataByLocateAndIspServiceRequest
//
// @return DescribeRangeDataByLocateAndIspServiceResponse
func (client *Client) DescribeRangeDataByLocateAndIspService(request *DescribeRangeDataByLocateAndIspServiceRequest) (_result *DescribeRangeDataByLocateAndIspServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRangeDataByLocateAndIspServiceResponse{}
	_body, _err := client.DescribeRangeDataByLocateAndIspServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of real-time log deliveries.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRealtimeDeliveryAccRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRealtimeDeliveryAccResponse
func (client *Client) DescribeRealtimeDeliveryAccWithOptions(request *DescribeRealtimeDeliveryAccRequest, runtime *dara.RuntimeOptions) (_result *DescribeRealtimeDeliveryAccResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRealtimeDeliveryAcc"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRealtimeDeliveryAccResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of real-time log deliveries.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeRealtimeDeliveryAccRequest
//
// @return DescribeRealtimeDeliveryAccResponse
func (client *Client) DescribeRealtimeDeliveryAcc(request *DescribeRealtimeDeliveryAccRequest) (_result *DescribeRealtimeDeliveryAccResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRealtimeDeliveryAccResponse{}
	_body, _err := client.DescribeRealtimeDeliveryAccWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeRefreshQuota
//
// Description:
//
// Queries the maximum and remaining numbers of URLs and directories that can be refreshed, the maximum and remaining numbers of times that you can prefetch content, and the maximum and remaining numbers of URLs and directories that can be blocked on the current day.
//
// @param request - DescribeRefreshQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRefreshQuotaResponse
func (client *Client) DescribeRefreshQuotaWithOptions(request *DescribeRefreshQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeRefreshQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRefreshQuota"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeRefreshQuota
//
// Description:
//
// Queries the maximum and remaining numbers of URLs and directories that can be refreshed, the maximum and remaining numbers of times that you can prefetch content, and the maximum and remaining numbers of URLs and directories that can be blocked on the current day.
//
// @param request - DescribeRefreshQuotaRequest
//
// @return DescribeRefreshQuotaResponse
func (client *Client) DescribeRefreshQuota(request *DescribeRefreshQuotaRequest) (_result *DescribeRefreshQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRefreshQuotaResponse{}
	_body, _err := client.DescribeRefreshQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statuses of refresh or prefetch tasks by task ID.
//
// Description:
//
//	  You can query data in the last three days.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeRefreshTaskByIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRefreshTaskByIdResponse
func (client *Client) DescribeRefreshTaskByIdWithOptions(request *DescribeRefreshTaskByIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeRefreshTaskByIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRefreshTaskById"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRefreshTaskByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statuses of refresh or prefetch tasks by task ID.
//
// Description:
//
//	  You can query data in the last three days.
//
//		- You can call this operation up to 30 times per second per account.
//
// @param request - DescribeRefreshTaskByIdRequest
//
// @return DescribeRefreshTaskByIdResponse
func (client *Client) DescribeRefreshTaskById(request *DescribeRefreshTaskByIdRequest) (_result *DescribeRefreshTaskByIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRefreshTaskByIdResponse{}
	_body, _err := client.DescribeRefreshTaskByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of refresh or prefetch tasks that belong to an accelerated domain name.
//
// Description:
//
//	  You can query the status of tasks by task ID or URL.
//
//		- You can set both the **TaskId*	- and **ObjectPath*	- parameters. If you do not set the **TaskId*	- or **ObjectPath*	- parameter, data entries on the first page (20 entries) collected in the last 3 days are returned.
//
//		- You can query data collected in the last 3 days.
//
//		- If auto CDN cache update is enabled in the Object Storage Service (OSS) console, you cannot call the DescribeRefreshTasks operation to query automatic refresh tasks in OSS.
//
//		- You can call this operation up to 10 times per second per account. If you want to query tasks at a higher frequency, call the [DescribeRefreshTaskById](https://help.aliyun.com/document_detail/187709.html) operation. This operation allows you to query tasks by task ID.
//
// @param request - DescribeRefreshTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRefreshTasksResponse
func (client *Client) DescribeRefreshTasksWithOptions(request *DescribeRefreshTasksRequest, runtime *dara.RuntimeOptions) (_result *DescribeRefreshTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRefreshTasks"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRefreshTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of refresh or prefetch tasks that belong to an accelerated domain name.
//
// Description:
//
//	  You can query the status of tasks by task ID or URL.
//
//		- You can set both the **TaskId*	- and **ObjectPath*	- parameters. If you do not set the **TaskId*	- or **ObjectPath*	- parameter, data entries on the first page (20 entries) collected in the last 3 days are returned.
//
//		- You can query data collected in the last 3 days.
//
//		- If auto CDN cache update is enabled in the Object Storage Service (OSS) console, you cannot call the DescribeRefreshTasks operation to query automatic refresh tasks in OSS.
//
//		- You can call this operation up to 10 times per second per account. If you want to query tasks at a higher frequency, call the [DescribeRefreshTaskById](https://help.aliyun.com/document_detail/187709.html) operation. This operation allows you to query tasks by task ID.
//
// @param request - DescribeRefreshTasksRequest
//
// @return DescribeRefreshTasksResponse
func (client *Client) DescribeRefreshTasks(request *DescribeRefreshTasksRequest) (_result *DescribeRefreshTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRefreshTasksResponse{}
	_body, _err := client.DescribeRefreshTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries node IP addresses in the staging environment.
//
// Description:
//
// >The maximum number of times that each user can call this operation per second is 30.
//
// @param request - DescribeStagingIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeStagingIpResponse
func (client *Client) DescribeStagingIpWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeStagingIpResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeStagingIp"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeStagingIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries node IP addresses in the staging environment.
//
// Description:
//
// >The maximum number of times that each user can call this operation per second is 30.
//
// @return DescribeStagingIpResponse
func (client *Client) DescribeStagingIp() (_result *DescribeStagingIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeStagingIpResponse{}
	_body, _err := client.DescribeStagingIpWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tags that are added to specified resources.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 10.
//
// @param request - DescribeTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagResourcesResponse
func (client *Client) DescribeTagResourcesWithOptions(request *DescribeTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTagResources"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags that are added to specified resources.
//
// Description:
//
// >  The maximum number of times that each user can call this operation per second is 10.
//
// @param request - DescribeTagResourcesRequest
//
// @return DescribeTagResourcesResponse
func (client *Client) DescribeTagResources(request *DescribeTagResourcesRequest) (_result *DescribeTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagResourcesResponse{}
	_body, _err := client.DescribeTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top N domain names ranked by network traffic. You can query data collected in the last 30 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the current month. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeTopDomainsByFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTopDomainsByFlowResponse
func (client *Client) DescribeTopDomainsByFlowWithOptions(request *DescribeTopDomainsByFlowRequest, runtime *dara.RuntimeOptions) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTopDomainsByFlow"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top N domain names ranked by network traffic. You can query data collected in the last 30 days.
//
// Description:
//
// The statistical analysis feature of Alibaba Cloud CDN is no longer available. The API operations related to the statistical analysis feature are no longer maintained. We recommend that you do not use the API operations because data may be missing or inaccurate. You can use the [operations report](https://help.aliyun.com/document_detail/279577.html) feature for data analysis.
//
// > 	- If you do not set the **StartTime*	- or **EndTime*	- parameter, the request returns the data collected in the current month. If you set both these parameters, the request returns the data collected within the specified time range.
//
// >	- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeTopDomainsByFlowRequest
//
// @return DescribeTopDomainsByFlowResponse
func (client *Client) DescribeTopDomainsByFlow(request *DescribeTopDomainsByFlowRequest) (_result *DescribeTopDomainsByFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTopDomainsByFlowResponse{}
	_body, _err := client.DescribeTopDomainsByFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a user.
//
// @param request - DescribeUserCdnStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserCdnStatusResponse
func (client *Client) DescribeUserCdnStatusWithOptions(request *DescribeUserCdnStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserCdnStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserCdnStatus"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserCdnStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a user.
//
// @param request - DescribeUserCdnStatusRequest
//
// @return DescribeUserCdnStatusResponse
func (client *Client) DescribeUserCdnStatus(request *DescribeUserCdnStatusRequest) (_result *DescribeUserCdnStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserCdnStatusResponse{}
	_body, _err := client.DescribeUserCdnStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of domain names whose SSL certificates are about to expire or have already expired.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserCertificateExpireCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserCertificateExpireCountResponse
func (client *Client) DescribeUserCertificateExpireCountWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeUserCertificateExpireCountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserCertificateExpireCount"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserCertificateExpireCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of domain names whose SSL certificates are about to expire or have already expired.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeUserCertificateExpireCountResponse
func (client *Client) DescribeUserCertificateExpireCount() (_result *DescribeUserCertificateExpireCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserCertificateExpireCountResponse{}
	_body, _err := client.DescribeUserCertificateExpireCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries configurations of security features.
//
// @param request - DescribeUserConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserConfigsResponse
func (client *Client) DescribeUserConfigsWithOptions(request *DescribeUserConfigsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserConfigs"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configurations of security features.
//
// @param request - DescribeUserConfigsRequest
//
// @return DescribeUserConfigsResponse
func (client *Client) DescribeUserConfigs(request *DescribeUserConfigsRequest) (_result *DescribeUserConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserConfigsResponse{}
	_body, _err := client.DescribeUserConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all accelerated domain names in your Alibaba Cloud account and the status of the accelerated domain names. You can filter domain names by name or status. Fuzzy match is supported.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can specify up to 50 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - DescribeUserDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserDomainsResponse
func (client *Client) DescribeUserDomainsWithOptions(request *DescribeUserDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CdnType) {
		query["CdnType"] = request.CdnType
	}

	if !dara.IsNil(request.ChangeEndTime) {
		query["ChangeEndTime"] = request.ChangeEndTime
	}

	if !dara.IsNil(request.ChangeStartTime) {
		query["ChangeStartTime"] = request.ChangeStartTime
	}

	if !dara.IsNil(request.CheckDomainShow) {
		query["CheckDomainShow"] = request.CheckDomainShow
	}

	if !dara.IsNil(request.Coverage) {
		query["Coverage"] = request.Coverage
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainSearchType) {
		query["DomainSearchType"] = request.DomainSearchType
	}

	if !dara.IsNil(request.DomainStatus) {
		query["DomainStatus"] = request.DomainStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserDomains"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all accelerated domain names in your Alibaba Cloud account and the status of the accelerated domain names. You can filter domain names by name or status. Fuzzy match is supported.
//
// Description:
//
//	  You can call this operation up to 100 times per second per account.
//
//		- You can specify up to 50 domain names in each request. Separate multiple domain names with commas (,).
//
// @param request - DescribeUserDomainsRequest
//
// @return DescribeUserDomainsResponse
func (client *Client) DescribeUserDomains(request *DescribeUserDomainsRequest) (_result *DescribeUserDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserDomainsResponse{}
	_body, _err := client.DescribeUserDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user tags.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserTagsResponse
func (client *Client) DescribeUserTagsWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeUserTagsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserTags"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user tags.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return DescribeUserTagsResponse
func (client *Client) DescribeUserTags() (_result *DescribeUserTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserTagsResponse{}
	_body, _err := client.DescribeUserTagsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries usage export tasks that were created in the last three months.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserUsageDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserUsageDataExportTaskResponse
func (client *Client) DescribeUserUsageDataExportTaskWithOptions(request *DescribeUserUsageDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserUsageDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeUserUsageDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserUsageDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries usage export tasks that were created in the last three months.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserUsageDataExportTaskRequest
//
// @return DescribeUserUsageDataExportTaskResponse
func (client *Client) DescribeUserUsageDataExportTask(request *DescribeUserUsageDataExportTaskRequest) (_result *DescribeUserUsageDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserUsageDataExportTaskResponse{}
	_body, _err := client.DescribeUserUsageDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tasks that were used to export resource usage details of one or more accelerated domain names that belong to your Alibaba Cloud account. Resource usage information is collected every five minutes.
//
// Description:
//
//	  This operation has been available since July 20, 2018. You can query information about resource usage collected within the last three months.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserUsageDetailDataExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserUsageDetailDataExportTaskResponse
func (client *Client) DescribeUserUsageDetailDataExportTaskWithOptions(request *DescribeUserUsageDetailDataExportTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserUsageDetailDataExportTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeUserUsageDetailDataExportTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserUsageDetailDataExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tasks that were used to export resource usage details of one or more accelerated domain names that belong to your Alibaba Cloud account. Resource usage information is collected every five minutes.
//
// Description:
//
//	  This operation has been available since July 20, 2018. You can query information about resource usage collected within the last three months.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - DescribeUserUsageDetailDataExportTaskRequest
//
// @return DescribeUserUsageDetailDataExportTaskResponse
func (client *Client) DescribeUserUsageDetailDataExportTask(request *DescribeUserUsageDetailDataExportTaskRequest) (_result *DescribeUserUsageDetailDataExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserUsageDetailDataExportTaskResponse{}
	_body, _err := client.DescribeUserUsageDetailDataExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virtual IP addresses (VIPs) of CDN points of presence (POPs) by domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserVipsByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserVipsByDomainResponse
func (client *Client) DescribeUserVipsByDomainWithOptions(request *DescribeUserVipsByDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserVipsByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserVipsByDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserVipsByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virtual IP addresses (VIPs) of CDN points of presence (POPs) by domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - DescribeUserVipsByDomainRequest
//
// @return DescribeUserVipsByDomainResponse
func (client *Client) DescribeUserVipsByDomain(request *DescribeUserVipsByDomainRequest) (_result *DescribeUserVipsByDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserVipsByDomainResponse{}
	_body, _err := client.DescribeUserVipsByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content of an accelerated domain name.
//
// @param request - DescribeVerifyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVerifyContentResponse
func (client *Client) DescribeVerifyContentWithOptions(request *DescribeVerifyContentRequest, runtime *dara.RuntimeOptions) (_result *DescribeVerifyContentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVerifyContent"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVerifyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ownership verification content of an accelerated domain name.
//
// @param request - DescribeVerifyContentRequest
//
// @return DescribeVerifyContentResponse
func (client *Client) DescribeVerifyContent(request *DescribeVerifyContentRequest) (_result *DescribeVerifyContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVerifyContentResponse{}
	_body, _err := client.DescribeVerifyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables real-time log delivery for specific accelerated domain names.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DisableRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableRealtimeLogDeliveryResponse
func (client *Client) DisableRealtimeLogDeliveryWithOptions(request *DisableRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *DisableRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables real-time log delivery for specific accelerated domain names.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - DisableRealtimeLogDeliveryRequest
//
// @return DisableRealtimeLogDeliveryResponse
func (client *Client) DisableRealtimeLogDelivery(request *DisableRealtimeLogDeliveryRequest) (_result *DisableRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableRealtimeLogDeliveryResponse{}
	_body, _err := client.DisableRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables real-time log delivery for an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - EnableRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableRealtimeLogDeliveryResponse
func (client *Client) EnableRealtimeLogDeliveryWithOptions(request *EnableRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *EnableRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables real-time log delivery for an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - EnableRealtimeLogDeliveryRequest
//
// @return EnableRealtimeLogDeliveryResponse
func (client *Client) EnableRealtimeLogDelivery(request *EnableRealtimeLogDeliveryRequest) (_result *EnableRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableRealtimeLogDeliveryResponse{}
	_body, _err := client.EnableRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 按域名、functionName查询灰度配置信息，返回的信息中包含当前的灰度状态、灰度进度
//
// @param request - GetGrayDomainFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGrayDomainFunctionResponse
func (client *Client) GetGrayDomainFunctionWithOptions(request *GetGrayDomainFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetGrayDomainFunctionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.FunctionNames) {
		body["FunctionNames"] = request.FunctionNames
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGrayDomainFunction"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGrayDomainFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 按域名、functionName查询灰度配置信息，返回的信息中包含当前的灰度状态、灰度进度
//
// @param request - GetGrayDomainFunctionRequest
//
// @return GetGrayDomainFunctionResponse
func (client *Client) GetGrayDomainFunction(request *GetGrayDomainFunctionRequest) (_result *GetGrayDomainFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetGrayDomainFunctionResponse{}
	_body, _err := client.GetGrayDomainFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain names by log configuration ID.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListDomainsByLogConfigIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsByLogConfigIdResponse
func (client *Client) ListDomainsByLogConfigIdWithOptions(request *ListDomainsByLogConfigIdRequest, runtime *dara.RuntimeOptions) (_result *ListDomainsByLogConfigIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomainsByLogConfigId"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsByLogConfigIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain names by log configuration ID.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListDomainsByLogConfigIdRequest
//
// @return ListDomainsByLogConfigIdResponse
func (client *Client) ListDomainsByLogConfigId(request *ListDomainsByLogConfigIdRequest) (_result *ListDomainsByLogConfigIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainsByLogConfigIdResponse{}
	_body, _err := client.ListDomainsByLogConfigIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Function Compute trigger that is set for an Alibaba Cloud CDN event.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListFCTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFCTriggerResponse
func (client *Client) ListFCTriggerWithOptions(request *ListFCTriggerRequest, runtime *dara.RuntimeOptions) (_result *ListFCTriggerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFCTrigger"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFCTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Function Compute trigger that is set for an Alibaba Cloud CDN event.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListFCTriggerRequest
//
// @return ListFCTriggerResponse
func (client *Client) ListFCTrigger(request *ListFCTriggerRequest) (_result *ListFCTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFCTriggerResponse{}
	_body, _err := client.ListFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all real-time log delivery tasks within your Alibaba Cloud account.
//
// @param request - ListRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRealtimeLogDeliveryResponse
func (client *Client) ListRealtimeLogDeliveryWithOptions(runtime *dara.RuntimeOptions) (_result *ListRealtimeLogDeliveryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all real-time log delivery tasks within your Alibaba Cloud account.
//
// @return ListRealtimeLogDeliveryResponse
func (client *Client) ListRealtimeLogDelivery() (_result *ListRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRealtimeLogDeliveryResponse{}
	_body, _err := client.ListRealtimeLogDeliveryWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all domain names that are associated with a specific real-time log delivery configuration record.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListRealtimeLogDeliveryDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRealtimeLogDeliveryDomainsResponse
func (client *Client) ListRealtimeLogDeliveryDomainsWithOptions(request *ListRealtimeLogDeliveryDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListRealtimeLogDeliveryDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRealtimeLogDeliveryDomains"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all domain names that are associated with a specific real-time log delivery configuration record.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListRealtimeLogDeliveryDomainsRequest
//
// @return ListRealtimeLogDeliveryDomainsResponse
func (client *Client) ListRealtimeLogDeliveryDomains(request *ListRealtimeLogDeliveryDomainsRequest) (_result *ListRealtimeLogDeliveryDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRealtimeLogDeliveryDomainsResponse{}
	_body, _err := client.ListRealtimeLogDeliveryDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the real-time log delivery feature in a specified region.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListRealtimeLogDeliveryInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRealtimeLogDeliveryInfosResponse
func (client *Client) ListRealtimeLogDeliveryInfosWithOptions(runtime *dara.RuntimeOptions) (_result *ListRealtimeLogDeliveryInfosResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListRealtimeLogDeliveryInfos"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the real-time log delivery feature in a specified region.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return ListRealtimeLogDeliveryInfosResponse
func (client *Client) ListRealtimeLogDeliveryInfos() (_result *ListRealtimeLogDeliveryInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRealtimeLogDeliveryInfosResponse{}
	_body, _err := client.ListRealtimeLogDeliveryInfosWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a resource.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TagOwnerBid) {
		query["TagOwnerBid"] = request.TagOwnerBid
	}

	if !dara.IsNil(request.TagOwnerUid) {
		query["TagOwnerUid"] = request.TagOwnerUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a resource.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all custom log configurations in your account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ListUserCustomLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserCustomLogConfigResponse
func (client *Client) ListUserCustomLogConfigWithOptions(runtime *dara.RuntimeOptions) (_result *ListUserCustomLogConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserCustomLogConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserCustomLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all custom log configurations in your account.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @return ListUserCustomLogConfigResponse
func (client *Client) ListUserCustomLogConfig() (_result *ListUserCustomLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserCustomLogConfigResponse{}
	_body, _err := client.ListUserCustomLogConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - ModifyCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdnDomainResponse
func (client *Client) ModifyCdnDomainWithOptions(request *ModifyCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sources) {
		query["Sources"] = request.Sources
	}

	if !dara.IsNil(request.TopLevelDomain) {
		query["TopLevelDomain"] = request.TopLevelDomain
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - ModifyCdnDomainRequest
//
// @return ModifyCdnDomainResponse
func (client *Client) ModifyCdnDomain(request *ModifyCdnDomainRequest) (_result *ModifyCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdnDomainResponse{}
	_body, _err := client.ModifyCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Transfer domain names from an Alibaba Cloud account to the current account.
//
// Description:
//
// This operation is used in the following scenario:
//
//   - You have multiple Alibaba Cloud accounts and want to transfer domain names from Account A to Account B.
//
//   - You are prompted that a domain name has been added when you add the domain name to Alibaba Cloud CDN. You do not know which account does the domain name belong to, and you want to transfer the domain name to your current account.
//
// @param request - ModifyCdnDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdnDomainOwnerResponse
func (client *Client) ModifyCdnDomainOwnerWithOptions(request *ModifyCdnDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdnDomainOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdnDomainOwner"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdnDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Transfer domain names from an Alibaba Cloud account to the current account.
//
// Description:
//
// This operation is used in the following scenario:
//
//   - You have multiple Alibaba Cloud accounts and want to transfer domain names from Account A to Account B.
//
//   - You are prompted that a domain name has been added when you add the domain name to Alibaba Cloud CDN. You do not know which account does the domain name belong to, and you want to transfer the domain name to your current account.
//
// @param request - ModifyCdnDomainOwnerRequest
//
// @return ModifyCdnDomainOwnerResponse
func (client *Client) ModifyCdnDomainOwner(request *ModifyCdnDomainOwnerRequest) (_result *ModifyCdnDomainOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdnDomainOwnerResponse{}
	_body, _err := client.ModifyCdnDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the acceleration region for an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyCdnDomainSchdmByPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdnDomainSchdmByPropertyResponse
func (client *Client) ModifyCdnDomainSchdmByPropertyWithOptions(request *ModifyCdnDomainSchdmByPropertyRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdnDomainSchdmByPropertyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Property) {
		query["Property"] = request.Property
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdnDomainSchdmByProperty"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the acceleration region for an accelerated domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyCdnDomainSchdmByPropertyRequest
//
// @return ModifyCdnDomainSchdmByPropertyResponse
func (client *Client) ModifyCdnDomainSchdmByProperty(request *ModifyCdnDomainSchdmByPropertyRequest) (_result *ModifyCdnDomainSchdmByPropertyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdnDomainSchdmByPropertyResponse{}
	_body, _err := client.ModifyCdnDomainSchdmByPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the metering method of Alibaba Cloud CDN.
//
// @param request - ModifyCdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCdnServiceResponse
func (client *Client) ModifyCdnServiceWithOptions(request *ModifyCdnServiceRequest, runtime *dara.RuntimeOptions) (_result *ModifyCdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCdnService"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the metering method of Alibaba Cloud CDN.
//
// @param request - ModifyCdnServiceRequest
//
// @return ModifyCdnServiceResponse
func (client *Client) ModifyCdnService(request *ModifyCdnServiceRequest) (_result *ModifyCdnServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCdnServiceResponse{}
	_body, _err := client.ModifyCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # A客户定制新增修改域名采样率接口
//
// @param request - ModifyCustomDomainSampleRateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCustomDomainSampleRateResponse
func (client *Client) ModifyCustomDomainSampleRateWithOptions(request *ModifyCustomDomainSampleRateRequest, runtime *dara.RuntimeOptions) (_result *ModifyCustomDomainSampleRateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseConfigID) {
		body["BaseConfigID"] = request.BaseConfigID
	}

	if !dara.IsNil(request.DomainNames) {
		body["DomainNames"] = request.DomainNames
	}

	if !dara.IsNil(request.SampleRate) {
		body["SampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SinkID) {
		body["SinkID"] = request.SinkID
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCustomDomainSampleRate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCustomDomainSampleRateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # A客户定制新增修改域名采样率接口
//
// @param request - ModifyCustomDomainSampleRateRequest
//
// @return ModifyCustomDomainSampleRateResponse
func (client *Client) ModifyCustomDomainSampleRate(request *ModifyCustomDomainSampleRateRequest) (_result *ModifyCustomDomainSampleRateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCustomDomainSampleRateResponse{}
	_body, _err := client.ModifyCustomDomainSampleRateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of real-time log delivery for a specific domain name. Each domain name supports only one Logstore.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyRealtimeLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRealtimeLogDeliveryResponse
func (client *Client) ModifyRealtimeLogDeliveryWithOptions(request *ModifyRealtimeLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *ModifyRealtimeLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRealtimeLogDelivery"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRealtimeLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of real-time log delivery for a specific domain name. Each domain name supports only one Logstore.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - ModifyRealtimeLogDeliveryRequest
//
// @return ModifyRealtimeLogDeliveryResponse
func (client *Client) ModifyRealtimeLogDelivery(request *ModifyRealtimeLogDeliveryRequest) (_result *ModifyRealtimeLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRealtimeLogDeliveryResponse{}
	_body, _err := client.ModifyRealtimeLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Activates Alibaba Cloud CDN. You must activate Alibaba Cloud CDN before you can manage domain names in Alibaba Cloud CDN.
//
// Description:
//
//	  Alibaba Cloud CDN can be activated only once per Alibaba Cloud account. The Alibaba Cloud account must complete real-name verification to activate Alibaba Cloud CDN.
//
//		- You can call this operation up to five times per second per user.
//
// @param request - OpenCdnServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenCdnServiceResponse
func (client *Client) OpenCdnServiceWithOptions(request *OpenCdnServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenCdnServiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenCdnService"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates Alibaba Cloud CDN. You must activate Alibaba Cloud CDN before you can manage domain names in Alibaba Cloud CDN.
//
// Description:
//
//	  Alibaba Cloud CDN can be activated only once per Alibaba Cloud account. The Alibaba Cloud account must complete real-name verification to activate Alibaba Cloud CDN.
//
//		- You can call this operation up to five times per second per user.
//
// @param request - OpenCdnServiceRequest
//
// @return OpenCdnServiceResponse
func (client *Client) OpenCdnService(request *OpenCdnServiceRequest) (_result *OpenCdnServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenCdnServiceResponse{}
	_body, _err := client.OpenCdnServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布灰度配置到线上，支持多种模式，如全网发布、指定方式(灰度发布)，回滚
//
// @param request - PublishGrayDomainConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishGrayDomainConfigResponse
func (client *Client) PublishGrayDomainConfigWithOptions(request *PublishGrayDomainConfigRequest, runtime *dara.RuntimeOptions) (_result *PublishGrayDomainConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomCountryId) {
		body["CustomCountryId"] = request.CustomCountryId
	}

	if !dara.IsNil(request.CustomPercent) {
		body["CustomPercent"] = request.CustomPercent
	}

	if !dara.IsNil(request.CustomProvinceId) {
		body["CustomProvinceId"] = request.CustomProvinceId
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.PublishMode) {
		body["PublishMode"] = request.PublishMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishGrayDomainConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishGrayDomainConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布灰度配置到线上，支持多种模式，如全网发布、指定方式(灰度发布)，回滚
//
// @param request - PublishGrayDomainConfigRequest
//
// @return PublishGrayDomainConfigResponse
func (client *Client) PublishGrayDomainConfig(request *PublishGrayDomainConfigRequest) (_result *PublishGrayDomainConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishGrayDomainConfigResponse{}
	_body, _err := client.PublishGrayDomainConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes the configurations of the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - PublishStagingConfigToProductionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishStagingConfigToProductionResponse
func (client *Client) PublishStagingConfigToProductionWithOptions(request *PublishStagingConfigToProductionRequest, runtime *dara.RuntimeOptions) (_result *PublishStagingConfigToProductionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishStagingConfigToProduction"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishStagingConfigToProductionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes the configurations of the staging environment to the production environment.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - PublishStagingConfigToProductionRequest
//
// @return PublishStagingConfigToProductionResponse
func (client *Client) PublishStagingConfigToProduction(request *PublishStagingConfigToProductionRequest) (_result *PublishStagingConfigToProductionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishStagingConfigToProductionResponse{}
	_body, _err := client.PublishStagingConfigToProductionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Prefetches content from origin servers to points of presence (POPs). This reduces loads on origin servers because users can directly hit cache upon their first visits.
//
// Description:
//
//	  Alibaba Cloud CDN supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshObjectCaches](https://help.aliyun.com/document_detail/91164.html) operation to refresh content and call the [PushObjectCache](https://help.aliyun.com/document_detail/91161.html) operation to prefetch content.
//
//		- By default, each Alibaba Cloud account can submit up to 1,000 URLs per day. If the daily peak bandwidth value of your workloads exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to increase your daily quota. Alibaba Cloud reviews your application and then increases the quota accordingly.
//
//		- You can specify at most 100 URLs in each prefetch request.
//
//		- For each Alibaba Cloud account, the prefetch queue can contain up to 50,000 URLs. Content is prefetched based on the time when the URLs are submitted. The URL that is submitted the earliest has the highest priority. If the number of URLs in the queue reaches 50,000, you cannot submit more URLs until the number drops below 50,000.
//
//		- You can call this operation up to 50 times per second per account.
//
//		- For more information about how to automate refresh or prefetch tasks, see [Run scripts to refresh and prefetch content](https://help.aliyun.com/document_detail/151829.html).
//
// ## Precautions
//
//   - After a prefetch task is submitted and completed, the POPs immediately start to retrieve resources from the origin server. Therefore, a large number of refresh tasks cause a large number of concurrent download tasks. This increases the number of requests that are redirected to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - The time required for a prefetch task to complete is proportional to the size of the prefetched file. In actual practice, most prefetch tasks require 5 to 30 minutes to complete. A task with a smaller average file size requires less time.
//
//   - To allow RAM users to perform this operation, you must first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/260300.html).
//
// @param request - PushObjectCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushObjectCacheResponse
func (client *Client) PushObjectCacheWithOptions(request *PushObjectCacheRequest, runtime *dara.RuntimeOptions) (_result *PushObjectCacheResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.ConsistencyHash) {
		query["ConsistencyHash"] = request.ConsistencyHash
	}

	if !dara.IsNil(request.L2Preload) {
		query["L2Preload"] = request.L2Preload
	}

	if !dara.IsNil(request.ObjectPath) {
		query["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryHashkey) {
		query["QueryHashkey"] = request.QueryHashkey
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WithHeader) {
		query["WithHeader"] = request.WithHeader
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushObjectCache"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prefetches content from origin servers to points of presence (POPs). This reduces loads on origin servers because users can directly hit cache upon their first visits.
//
// Description:
//
//	  Alibaba Cloud CDN supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshObjectCaches](https://help.aliyun.com/document_detail/91164.html) operation to refresh content and call the [PushObjectCache](https://help.aliyun.com/document_detail/91161.html) operation to prefetch content.
//
//		- By default, each Alibaba Cloud account can submit up to 1,000 URLs per day. If the daily peak bandwidth value of your workloads exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to increase your daily quota. Alibaba Cloud reviews your application and then increases the quota accordingly.
//
//		- You can specify at most 100 URLs in each prefetch request.
//
//		- For each Alibaba Cloud account, the prefetch queue can contain up to 50,000 URLs. Content is prefetched based on the time when the URLs are submitted. The URL that is submitted the earliest has the highest priority. If the number of URLs in the queue reaches 50,000, you cannot submit more URLs until the number drops below 50,000.
//
//		- You can call this operation up to 50 times per second per account.
//
//		- For more information about how to automate refresh or prefetch tasks, see [Run scripts to refresh and prefetch content](https://help.aliyun.com/document_detail/151829.html).
//
// ## Precautions
//
//   - After a prefetch task is submitted and completed, the POPs immediately start to retrieve resources from the origin server. Therefore, a large number of refresh tasks cause a large number of concurrent download tasks. This increases the number of requests that are redirected to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - The time required for a prefetch task to complete is proportional to the size of the prefetched file. In actual practice, most prefetch tasks require 5 to 30 minutes to complete. A task with a smaller average file size requires less time.
//
//   - To allow RAM users to perform this operation, you must first grant them the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/260300.html).
//
// @param request - PushObjectCacheRequest
//
// @return PushObjectCacheResponse
func (client *Client) PushObjectCache(request *PushObjectCacheRequest) (_result *PushObjectCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PushObjectCacheResponse{}
	_body, _err := client.PushObjectCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refreshes the cache based on cache tags that you configured.
//
// @param request - RefreshObjectCacheByCacheTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshObjectCacheByCacheTagResponse
func (client *Client) RefreshObjectCacheByCacheTagWithOptions(request *RefreshObjectCacheByCacheTagRequest, runtime *dara.RuntimeOptions) (_result *RefreshObjectCacheByCacheTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheTag) {
		query["CacheTag"] = request.CacheTag
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshObjectCacheByCacheTag"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshObjectCacheByCacheTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes the cache based on cache tags that you configured.
//
// @param request - RefreshObjectCacheByCacheTagRequest
//
// @return RefreshObjectCacheByCacheTagResponse
func (client *Client) RefreshObjectCacheByCacheTag(request *RefreshObjectCacheByCacheTagRequest) (_result *RefreshObjectCacheByCacheTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshObjectCacheByCacheTagResponse{}
	_body, _err := client.RefreshObjectCacheByCacheTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Refreshes files on Alibaba Cloud CDN points of presence (POPs). After files are refreshed, the original file content immediately becomes invalid. If clients request the original file content, Alibaba Cloud CDN forwards the requests to the origin server. Then, Alibaba Cloud CDN caches the latest content to the POPs and returns the content to the clients. Alibaba Cloud CDN allows you to refresh content from multiple URLs at the same time.
//
// Description:
//
//	  Alibaba Cloud CDN supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshObjectCaches](https://help.aliyun.com/document_detail/91164.html) operation to refresh content and call the [PushObjectCache](https://help.aliyun.com/document_detail/91161.html) operation to prefetch content.
//
//		- You can call the RefreshObjectCaches operation up to 50 times per second per account.
//
//		- For more information about how to automatically refresh or prefetch tasks, see [Run scripts to refresh and prefetch content](https://help.aliyun.com/document_detail/151829.html).
//
// ## Precautions
//
//   - After a refresh task is submitted and completed, specific resources are removed from POPs. When a POP receives a request for the removed resources, the POP forwards the request to the origin server to retrieve the resources. The retrieved resources are returned to the client and cached on the POP. Multiple refresh tasks may cause a large number of resources to be removed from the POPs. This increases the number of requests that are forwarded to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - A refresh task takes effect 5 to 6 minutes after being submitted. This means that if the resource you want to refresh has a TTL of less than five minutes, you wait for it to expire instead of manually running a refresh task.
//
//   - If you want to use RAM users to refresh or prefetch resources, you must obtain the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/260300.html).
//
// ### Refresh quota
//
//   - By default, each Alibaba Cloud account can refresh content from up to 10,000 URLs and 100 directories per day. The directories include subdirectories. If the daily peak bandwidth value exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to request a quota increase. Alibaba Cloud CDN evaluates your application based on your workloads.
//
//   - By default, each Alibaba Cloud account can submit up to 20 refresh rules that contain regular expressions per day. If the daily peak bandwidth of your Alibaba Cloud account exceeds 10 Gbit/s, you can [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex) to request a quota increase.
//
//   - You can specify up to 1,000 URL refresh rules, 100 directory refresh rules, or 1 refresh rule that contains regular expressions in each call.
//
//   - You can refresh up to 1,000 URLs per minute for each domain name.
//
// @param request - RefreshObjectCachesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshObjectCachesResponse
func (client *Client) RefreshObjectCachesWithOptions(request *RefreshObjectCachesRequest, runtime *dara.RuntimeOptions) (_result *RefreshObjectCachesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.ObjectPath) {
		body["ObjectPath"] = request.ObjectPath
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshObjectCaches"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Refreshes files on Alibaba Cloud CDN points of presence (POPs). After files are refreshed, the original file content immediately becomes invalid. If clients request the original file content, Alibaba Cloud CDN forwards the requests to the origin server. Then, Alibaba Cloud CDN caches the latest content to the POPs and returns the content to the clients. Alibaba Cloud CDN allows you to refresh content from multiple URLs at the same time.
//
// Description:
//
//	  Alibaba Cloud CDN supports POST requests in which parameters are sent as a form.
//
//		- You can call the [RefreshObjectCaches](https://help.aliyun.com/document_detail/91164.html) operation to refresh content and call the [PushObjectCache](https://help.aliyun.com/document_detail/91161.html) operation to prefetch content.
//
//		- You can call the RefreshObjectCaches operation up to 50 times per second per account.
//
//		- For more information about how to automatically refresh or prefetch tasks, see [Run scripts to refresh and prefetch content](https://help.aliyun.com/document_detail/151829.html).
//
// ## Precautions
//
//   - After a refresh task is submitted and completed, specific resources are removed from POPs. When a POP receives a request for the removed resources, the POP forwards the request to the origin server to retrieve the resources. The retrieved resources are returned to the client and cached on the POP. Multiple refresh tasks may cause a large number of resources to be removed from the POPs. This increases the number of requests that are forwarded to the origin server. The back-to-origin routing process consumes more bandwidth resources and the origin server may be overwhelmed.
//
//   - A refresh task takes effect 5 to 6 minutes after being submitted. This means that if the resource you want to refresh has a TTL of less than five minutes, you wait for it to expire instead of manually running a refresh task.
//
//   - If you want to use RAM users to refresh or prefetch resources, you must obtain the required permissions. For more information, see [Authorize a RAM user to prefetch and refresh resources](https://help.aliyun.com/document_detail/260300.html).
//
// ### Refresh quota
//
//   - By default, each Alibaba Cloud account can refresh content from up to 10,000 URLs and 100 directories per day. The directories include subdirectories. If the daily peak bandwidth value exceeds 200 Mbit/s, you can [submit a ticket](https://account.alibabacloud.com/login/login.htm?oauth_callback=https%3A//ticket-intl.console.aliyun.com/%23/ticket/createIndex) to request a quota increase. Alibaba Cloud CDN evaluates your application based on your workloads.
//
//   - By default, each Alibaba Cloud account can submit up to 20 refresh rules that contain regular expressions per day. If the daily peak bandwidth of your Alibaba Cloud account exceeds 10 Gbit/s, you can [submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex) to request a quota increase.
//
//   - You can specify up to 1,000 URL refresh rules, 100 directory refresh rules, or 1 refresh rule that contains regular expressions in each call.
//
//   - You can refresh up to 1,000 URLs per minute for each domain name.
//
// @param request - RefreshObjectCachesRequest
//
// @return RefreshObjectCachesResponse
func (client *Client) RefreshObjectCaches(request *RefreshObjectCachesRequest) (_result *RefreshObjectCachesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshObjectCachesResponse{}
	_body, _err := client.RefreshObjectCachesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back configurations in the staging environment. After you call this operation, all configurations in the staging environment are cleared.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - RollbackStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackStagingConfigResponse
func (client *Client) RollbackStagingConfigWithOptions(request *RollbackStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *RollbackStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackStagingConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back configurations in the staging environment. After you call this operation, all configurations in the staging environment are cleared.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - RollbackStagingConfigRequest
//
// @return RollbackStagingConfigResponse
func (client *Client) RollbackStagingConfig(request *RollbackStagingConfigRequest) (_result *RollbackStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RollbackStagingConfigResponse{}
	_body, _err := client.RollbackStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to configure an SSL certificate for a specific domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - SetCdnDomainCSRCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCdnDomainCSRCertificateResponse
func (client *Client) SetCdnDomainCSRCertificateWithOptions(request *SetCdnDomainCSRCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetCdnDomainCSRCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ServerCertificate) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCdnDomainCSRCertificate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCdnDomainCSRCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to configure an SSL certificate for a specific domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - SetCdnDomainCSRCertificateRequest
//
// @return SetCdnDomainCSRCertificateResponse
func (client *Client) SetCdnDomainCSRCertificate(request *SetCdnDomainCSRCertificateRequest) (_result *SetCdnDomainCSRCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCdnDomainCSRCertificateResponse{}
	_body, _err := client.SetCdnDomainCSRCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a ShangMi (SM) certificate for a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetCdnDomainSMCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCdnDomainSMCertificateResponse
func (client *Client) SetCdnDomainSMCertificateWithOptions(request *SetCdnDomainSMCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetCdnDomainSMCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCdnDomainSMCertificate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCdnDomainSMCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a ShangMi (SM) certificate for a domain name.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetCdnDomainSMCertificateRequest
//
// @return SetCdnDomainSMCertificateResponse
func (client *Client) SetCdnDomainSMCertificate(request *SetCdnDomainSMCertificateRequest) (_result *SetCdnDomainSMCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCdnDomainSMCertificateResponse{}
	_body, _err := client.SetCdnDomainSMCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the certificate for a domain name and updates the certificate information.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- Method: POST.
//
// @param request - SetCdnDomainSSLCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCdnDomainSSLCertificateResponse
func (client *Client) SetCdnDomainSSLCertificateWithOptions(request *SetCdnDomainSSLCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetCdnDomainSSLCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.CertRegion) {
		query["CertRegion"] = request.CertRegion
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SSLPri) {
		query["SSLPri"] = request.SSLPri
	}

	if !dara.IsNil(request.SSLProtocol) {
		query["SSLProtocol"] = request.SSLProtocol
	}

	if !dara.IsNil(request.SSLPub) {
		query["SSLPub"] = request.SSLPub
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCdnDomainSSLCertificate"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCdnDomainSSLCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the certificate for a domain name and updates the certificate information.
//
// Description:
//
//	  You can call this operation up to 30 times per second per account.
//
//		- Method: POST.
//
// @param request - SetCdnDomainSSLCertificateRequest
//
// @return SetCdnDomainSSLCertificateResponse
func (client *Client) SetCdnDomainSSLCertificate(request *SetCdnDomainSSLCertificateRequest) (_result *SetCdnDomainSSLCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCdnDomainSSLCertificateResponse{}
	_body, _err := client.SetCdnDomainSSLCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a domain name to be accelerated in the staging environment.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - SetCdnDomainStagingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCdnDomainStagingConfigResponse
func (client *Client) SetCdnDomainStagingConfigWithOptions(request *SetCdnDomainStagingConfigRequest, runtime *dara.RuntimeOptions) (_result *SetCdnDomainStagingConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Functions) {
		query["Functions"] = request.Functions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCdnDomainStagingConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCdnDomainStagingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a domain name to be accelerated in the staging environment.
//
// Description:
//
// >  You can call this operation up to 30 times per second per account.
//
// @param request - SetCdnDomainStagingConfigRequest
//
// @return SetCdnDomainStagingConfigResponse
func (client *Client) SetCdnDomainStagingConfig(request *SetCdnDomainStagingConfigRequest) (_result *SetCdnDomainStagingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCdnDomainStagingConfigResponse{}
	_body, _err := client.SetCdnDomainStagingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Blocks or unblocks IP addresses from accessing domain names.
//
// Description:
//
// >
//
//   - To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - This operation is suitable for blocking or unblocking a maximum of 1,000 IP addresses or CIDR blocks at a time.
//
// @param request - SetCdnFullDomainsBlockIPRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCdnFullDomainsBlockIPResponse
func (client *Client) SetCdnFullDomainsBlockIPWithOptions(request *SetCdnFullDomainsBlockIPRequest, runtime *dara.RuntimeOptions) (_result *SetCdnFullDomainsBlockIPResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockInterval) {
		body["BlockInterval"] = request.BlockInterval
	}

	if !dara.IsNil(request.IPList) {
		body["IPList"] = request.IPList
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.UpdateType) {
		body["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCdnFullDomainsBlockIP"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCdnFullDomainsBlockIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Blocks or unblocks IP addresses from accessing domain names.
//
// Description:
//
// >
//
//   - To use this operation, [submit a ticket](https://workorder-intl.console.aliyun.com/?spm=5176.2020520001.aliyun_topbar.18.dbd44bd3e4f845#/ticket/createIndex).
//
//   - This operation is suitable for blocking or unblocking a maximum of 1,000 IP addresses or CIDR blocks at a time.
//
// @param request - SetCdnFullDomainsBlockIPRequest
//
// @return SetCdnFullDomainsBlockIPResponse
func (client *Client) SetCdnFullDomainsBlockIP(request *SetCdnFullDomainsBlockIPRequest) (_result *SetCdnFullDomainsBlockIPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCdnFullDomainsBlockIPResponse{}
	_body, _err := client.SetCdnFullDomainsBlockIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a custom origin header.
//
// @param request - SetReqHeaderConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetReqHeaderConfigResponse
func (client *Client) SetReqHeaderConfigWithOptions(request *SetReqHeaderConfigRequest, runtime *dara.RuntimeOptions) (_result *SetReqHeaderConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetReqHeaderConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetReqHeaderConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a custom origin header.
//
// @param request - SetReqHeaderConfigRequest
//
// @return SetReqHeaderConfigResponse
func (client *Client) SetReqHeaderConfig(request *SetReqHeaderConfigRequest) (_result *SetReqHeaderConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetReqHeaderConfigResponse{}
	_body, _err := client.SetReqHeaderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the virtual waiting room feature for an accelerated domain name. This operation is available only for accelerated domain names of the Dynamic CDN workload type.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetWaitingRoomConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWaitingRoomConfigResponse
func (client *Client) SetWaitingRoomConfigWithOptions(request *SetWaitingRoomConfigRequest, runtime *dara.RuntimeOptions) (_result *SetWaitingRoomConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowPct) {
		query["AllowPct"] = request.AllowPct
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GapTime) {
		query["GapTime"] = request.GapTime
	}

	if !dara.IsNil(request.MaxTimeWait) {
		query["MaxTimeWait"] = request.MaxTimeWait
	}

	if !dara.IsNil(request.WaitUri) {
		query["WaitUri"] = request.WaitUri
	}

	if !dara.IsNil(request.WaitUrl) {
		query["WaitUrl"] = request.WaitUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWaitingRoomConfig"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWaitingRoomConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the virtual waiting room feature for an accelerated domain name. This operation is available only for accelerated domain names of the Dynamic CDN workload type.
//
// Description:
//
// > You can call this operation up to 30 times per second per account.
//
// @param request - SetWaitingRoomConfigRequest
//
// @return SetWaitingRoomConfigResponse
func (client *Client) SetWaitingRoomConfig(request *SetWaitingRoomConfigRequest) (_result *SetWaitingRoomConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetWaitingRoomConfigResponse{}
	_body, _err := client.SetWaitingRoomConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a disabled domain name. After the domain name is enabled, the value of the DomainStatus parameter is changed to Online.
//
// Description:
//
//	  If the domain name is in an invalid state or you have an overdue payment in your account, the domain name cannot be enabled.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - StartCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCdnDomainResponse
func (client *Client) StartCdnDomainWithOptions(request *StartCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StartCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a disabled domain name. After the domain name is enabled, the value of the DomainStatus parameter is changed to Online.
//
// Description:
//
//	  If the domain name is in an invalid state or you have an overdue payment in your account, the domain name cannot be enabled.
//
//		- You can call this operation up to 100 times per second per account.
//
// @param request - StartCdnDomainRequest
//
// @return StartCdnDomainResponse
func (client *Client) StartCdnDomain(request *StartCdnDomainRequest) (_result *StartCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartCdnDomainResponse{}
	_body, _err := client.StartCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an accelerated domain name. After the domain name is disabled, the value of the DomainStatus parameter is changed to Offline.
//
// Description:
//
//	  After an accelerated domain is disabled, Alibaba Cloud CDN retains its information and routes all the requests that are destined for the accelerated domain to the origin server.
//
//		- You can call this operation up to 40 times per second per account.
//
// @param request - StopCdnDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCdnDomainResponse
func (client *Client) StopCdnDomainWithOptions(request *StopCdnDomainRequest, runtime *dara.RuntimeOptions) (_result *StopCdnDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCdnDomain"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCdnDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an accelerated domain name. After the domain name is disabled, the value of the DomainStatus parameter is changed to Offline.
//
// Description:
//
//	  After an accelerated domain is disabled, Alibaba Cloud CDN retains its information and routes all the requests that are destined for the accelerated domain to the origin server.
//
//		- You can call this operation up to 40 times per second per account.
//
// @param request - StopCdnDomainRequest
//
// @return StopCdnDomainResponse
func (client *Client) StopCdnDomain(request *StopCdnDomainRequest) (_result *StopCdnDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopCdnDomainResponse{}
	_body, _err := client.StopCdnDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more tags to specific resources.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more tags to specific resources.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from specified resources.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from specified resources.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tracking task.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateCdnDeliverTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCdnDeliverTaskResponse
func (client *Client) UpdateCdnDeliverTaskWithOptions(request *UpdateCdnDeliverTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateCdnDeliverTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Deliver) {
		body["Deliver"] = request.Deliver
	}

	if !dara.IsNil(request.DeliverId) {
		body["DeliverId"] = request.DeliverId
	}

	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Reports) {
		body["Reports"] = request.Reports
	}

	if !dara.IsNil(request.Schedule) {
		body["Schedule"] = request.Schedule
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCdnDeliverTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCdnDeliverTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tracking task.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateCdnDeliverTaskRequest
//
// @return UpdateCdnDeliverTaskResponse
func (client *Client) UpdateCdnDeliverTask(request *UpdateCdnDeliverTaskRequest) (_result *UpdateCdnDeliverTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCdnDeliverTaskResponse{}
	_body, _err := client.UpdateCdnDeliverTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates one or more operations reports.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateCdnSubTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCdnSubTaskResponse
func (client *Client) UpdateCdnSubTaskWithOptions(request *UpdateCdnSubTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateCdnSubTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		body["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReportIds) {
		body["ReportIds"] = request.ReportIds
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCdnSubTask"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCdnSubTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates one or more operations reports.
//
// Description:
//
// > You can call this operation up to three times per second per account.
//
// @param request - UpdateCdnSubTaskRequest
//
// @return UpdateCdnSubTaskResponse
func (client *Client) UpdateCdnSubTask(request *UpdateCdnSubTaskRequest) (_result *UpdateCdnSubTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCdnSubTaskResponse{}
	_body, _err := client.UpdateCdnSubTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified Function Compute trigger.
//
// @param request - UpdateFCTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFCTriggerResponse
func (client *Client) UpdateFCTriggerWithOptions(request *UpdateFCTriggerRequest, runtime *dara.RuntimeOptions) (_result *UpdateFCTriggerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TriggerARN) {
		query["TriggerARN"] = request.TriggerARN
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionARN) {
		body["FunctionARN"] = request.FunctionARN
	}

	if !dara.IsNil(request.Notes) {
		body["Notes"] = request.Notes
	}

	if !dara.IsNil(request.RoleARN) {
		body["RoleARN"] = request.RoleARN
	}

	if !dara.IsNil(request.SourceARN) {
		body["SourceARN"] = request.SourceARN
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFCTrigger"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFCTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified Function Compute trigger.
//
// @param request - UpdateFCTriggerRequest
//
// @return UpdateFCTriggerResponse
func (client *Client) UpdateFCTrigger(request *UpdateFCTriggerRequest) (_result *UpdateFCTriggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFCTriggerResponse{}
	_body, _err := client.UpdateFCTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a specified domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - VerifyDomainOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyDomainOwnerResponse
func (client *Client) VerifyDomainOwnerWithOptions(request *VerifyDomainOwnerRequest, runtime *dara.RuntimeOptions) (_result *VerifyDomainOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyDomainOwner"),
		Version:     dara.String("2018-05-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the ownership of a specified domain name.
//
// Description:
//
// > You can call this operation up to 100 times per second per account.
//
// @param request - VerifyDomainOwnerRequest
//
// @return VerifyDomainOwnerResponse
func (client *Client) VerifyDomainOwner(request *VerifyDomainOwnerRequest) (_result *VerifyDomainOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyDomainOwnerResponse{}
	_body, _err := client.VerifyDomainOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
