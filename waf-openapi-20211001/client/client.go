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
	client.EndpointMap = map[string]*string{
		"cn-qingdao":            dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":            dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-chengdu":            dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-huhehaote":          dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou":           dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen":           dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-heyuan":             dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":        dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("wafopenapi.ap-southeast-1.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("wafopenapi.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("waf-openapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Changes the resource group to which a protected object belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a protected object belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Clears an IP address blacklist for major event protection.
//
// @param request - ClearMajorProtectionBlackIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearMajorProtectionBlackIpResponse
func (client *Client) ClearMajorProtectionBlackIpWithOptions(request *ClearMajorProtectionBlackIpRequest, runtime *dara.RuntimeOptions) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearMajorProtectionBlackIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears an IP address blacklist for major event protection.
//
// @param request - ClearMajorProtectionBlackIpRequest
//
// @return ClearMajorProtectionBlackIpResponse
func (client *Client) ClearMajorProtectionBlackIp(request *ClearMajorProtectionBlackIpRequest) (_result *ClearMajorProtectionBlackIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearMajorProtectionBlackIpResponse{}
	_body, _err := client.ClearMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new protection template from the copy.
//
// @param request - CopyDefenseTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyDefenseTemplateResponse
func (client *Client) CopyDefenseTemplateWithOptions(request *CopyDefenseTemplateRequest, runtime *dara.RuntimeOptions) (_result *CopyDefenseTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyDefenseTemplate"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new protection template from the copy.
//
// @param request - CopyDefenseTemplateRequest
//
// @return CopyDefenseTemplateResponse
func (client *Client) CopyDefenseTemplate(request *CopyDefenseTemplateRequest) (_result *CopyDefenseTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyDefenseTemplateResponse{}
	_body, _err := client.CopyDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data export task in the API security module.
//
// @param request - CreateApiExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiExportResponse
func (client *Client) CreateApiExportWithOptions(request *CreateApiExportRequest, runtime *dara.RuntimeOptions) (_result *CreateApiExportResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiExport"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data export task in the API security module.
//
// @param request - CreateApiExportRequest
//
// @return CreateApiExportResponse
func (client *Client) CreateApiExport(request *CreateApiExportRequest) (_result *CreateApiExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiExportResponse{}
	_body, _err := client.CreateApiExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a certificate that uses an internationally accepted algorithm for a domain name added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param request - CreateCertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCertsResponse
func (client *Client) CreateCertsWithOptions(request *CreateCertsRequest, runtime *dara.RuntimeOptions) (_result *CreateCertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertContent) {
		query["CertContent"] = request.CertContent
	}

	if !dara.IsNil(request.CertKey) {
		query["CertKey"] = request.CertKey
	}

	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCerts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a certificate that uses an internationally accepted algorithm for a domain name added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param request - CreateCertsRequest
//
// @return CreateCertsResponse
func (client *Client) CreateCerts(request *CreateCertsRequest) (_result *CreateCertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCertsResponse{}
	_body, _err := client.CreateCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a service to Web Application Firewall (WAF). This operation is supported for only the Elastic Compute Service (ECS) and Classic Load Balancer (CLB) services.
//
// @param tmpReq - CreateCloudResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudResourceResponse
func (client *Client) CreateCloudResourceWithOptions(tmpReq *CreateCloudResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateCloudResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Listen) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, dara.String("Listen"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Redirect) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, dara.String("Redirect"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ListenShrink) {
		query["Listen"] = request.ListenShrink
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.RedirectShrink) {
		query["Redirect"] = request.RedirectShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a service to Web Application Firewall (WAF). This operation is supported for only the Elastic Compute Service (ECS) and Classic Load Balancer (CLB) services.
//
// @param request - CreateCloudResourceRequest
//
// @return CreateCloudResourceResponse
func (client *Client) CreateCloudResource(request *CreateCloudResourceRequest) (_result *CreateCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCloudResourceResponse{}
	_body, _err := client.CreateCloudResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建防护对象
//
// @param tmpReq - CreateDefenseResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDefenseResourceResponse
func (client *Client) CreateDefenseResourceWithOptions(tmpReq *CreateDefenseResourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDefenseResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDefenseResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomHeaders) {
		request.CustomHeadersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomHeaders, dara.String("CustomHeaders"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomHeadersShrink) {
		query["CustomHeaders"] = request.CustomHeadersShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Detail) {
		query["Detail"] = request.Detail
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Pattern) {
		query["Pattern"] = request.Pattern
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceOrigin) {
		query["ResourceOrigin"] = request.ResourceOrigin
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.XffStatus) {
		query["XffStatus"] = request.XffStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDefenseResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDefenseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建防护对象
//
// @param request - CreateDefenseResourceRequest
//
// @return CreateDefenseResourceResponse
func (client *Client) CreateDefenseResource(request *CreateDefenseResourceRequest) (_result *CreateDefenseResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDefenseResourceResponse{}
	_body, _err := client.CreateDefenseResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a protected object group.
//
// @param request - CreateDefenseResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDefenseResourceGroupResponse
func (client *Client) CreateDefenseResourceGroupWithOptions(request *CreateDefenseResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDefenseResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddList) {
		query["AddList"] = request.AddList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDefenseResourceGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protected object group.
//
// @param request - CreateDefenseResourceGroupRequest
//
// @return CreateDefenseResourceGroupResponse
func (client *Client) CreateDefenseResourceGroup(request *CreateDefenseResourceGroupRequest) (_result *CreateDefenseResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDefenseResourceGroupResponse{}
	_body, _err := client.CreateDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a protection rule.
//
// @param request - CreateDefenseRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDefenseRuleResponse
func (client *Client) CreateDefenseRuleWithOptions(request *CreateDefenseRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDefenseRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDefenseRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protection rule.
//
// @param request - CreateDefenseRuleRequest
//
// @return CreateDefenseRuleResponse
func (client *Client) CreateDefenseRule(request *CreateDefenseRuleRequest) (_result *CreateDefenseRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDefenseRuleResponse{}
	_body, _err := client.CreateDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a protection rule template.
//
// @param request - CreateDefenseTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDefenseTemplateResponse
func (client *Client) CreateDefenseTemplateWithOptions(request *CreateDefenseTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateDefenseTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateOrigin) {
		query["TemplateOrigin"] = request.TemplateOrigin
	}

	if !dara.IsNil(request.TemplateStatus) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.UnbindResourceGroups) {
		query["UnbindResourceGroups"] = request.UnbindResourceGroups
	}

	if !dara.IsNil(request.UnbindResources) {
		query["UnbindResources"] = request.UnbindResources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDefenseTemplate"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a protection rule template.
//
// @param request - CreateDefenseTemplateRequest
//
// @return CreateDefenseTemplateResponse
func (client *Client) CreateDefenseTemplate(request *CreateDefenseTemplateRequest) (_result *CreateDefenseTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDefenseTemplateResponse{}
	_body, _err := client.CreateDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a domain name to Web Application Firewall (WAF).
//
// @param tmpReq - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(tmpReq *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Listen) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, dara.String("Listen"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Redirect) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, dara.String("Redirect"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ListenShrink) {
		query["Listen"] = request.ListenShrink
	}

	if !dara.IsNil(request.RedirectShrink) {
		query["Redirect"] = request.RedirectShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a domain name to Web Application Firewall (WAF).
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a hybrid cloud cluster.
//
// @param request - CreateHybridCloudClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridCloudClusterResponse
func (client *Client) CreateHybridCloudClusterWithOptions(request *CreateHybridCloudClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridCloudClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.AccessRegion) {
		query["AccessRegion"] = request.AccessRegion
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.HttpPorts) {
		query["HttpPorts"] = request.HttpPorts
	}

	if !dara.IsNil(request.HttpsPorts) {
		query["HttpsPorts"] = request.HttpsPorts
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogFieldsNotReturned) {
		query["LogFieldsNotReturned"] = request.LogFieldsNotReturned
	}

	if !dara.IsNil(request.ProtectionServerCount) {
		query["ProtectionServerCount"] = request.ProtectionServerCount
	}

	if !dara.IsNil(request.ProxyStatus) {
		query["ProxyStatus"] = request.ProxyStatus
	}

	if !dara.IsNil(request.ProxyType) {
		query["ProxyType"] = request.ProxyType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridCloudCluster"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridCloudClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a hybrid cloud cluster.
//
// @param request - CreateHybridCloudClusterRequest
//
// @return CreateHybridCloudClusterResponse
func (client *Client) CreateHybridCloudCluster(request *CreateHybridCloudClusterRequest) (_result *CreateHybridCloudClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridCloudClusterResponse{}
	_body, _err := client.CreateHybridCloudClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增集群规则信息
//
// @param request - CreateHybridCloudClusterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridCloudClusterRuleResponse
func (client *Client) CreateHybridCloudClusterRuleWithOptions(request *CreateHybridCloudClusterRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridCloudClusterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridCloudClusterRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridCloudClusterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增集群规则信息
//
// @param request - CreateHybridCloudClusterRuleRequest
//
// @return CreateHybridCloudClusterRuleResponse
func (client *Client) CreateHybridCloudClusterRule(request *CreateHybridCloudClusterRuleRequest) (_result *CreateHybridCloudClusterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridCloudClusterRuleResponse{}
	_body, _err := client.CreateHybridCloudClusterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a node group for a hybrid cloud cluster.
//
// @param request - CreateHybridCloudGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHybridCloudGroupResponse
func (client *Client) CreateHybridCloudGroupWithOptions(request *CreateHybridCloudGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateHybridCloudGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackSourceMark) {
		query["BackSourceMark"] = request.BackSourceMark
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LoadBalanceIp) {
		query["LoadBalanceIp"] = request.LoadBalanceIp
	}

	if !dara.IsNil(request.LocationCode) {
		query["LocationCode"] = request.LocationCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHybridCloudGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHybridCloudGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node group for a hybrid cloud cluster.
//
// @param request - CreateHybridCloudGroupRequest
//
// @return CreateHybridCloudGroupResponse
func (client *Client) CreateHybridCloudGroup(request *CreateHybridCloudGroupRequest) (_result *CreateHybridCloudGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHybridCloudGroupResponse{}
	_body, _err := client.CreateHybridCloudGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IP address blacklist for major event protection.
//
// Description:
//
// This operation is available only on the China site (aliyun.com).
//
// @param request - CreateMajorProtectionBlackIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMajorProtectionBlackIpResponse
func (client *Client) CreateMajorProtectionBlackIpWithOptions(request *CreateMajorProtectionBlackIpRequest, runtime *dara.RuntimeOptions) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IpList) {
		query["IpList"] = request.IpList
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMajorProtectionBlackIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP address blacklist for major event protection.
//
// Description:
//
// This operation is available only on the China site (aliyun.com).
//
// @param request - CreateMajorProtectionBlackIpRequest
//
// @return CreateMajorProtectionBlackIpResponse
func (client *Client) CreateMajorProtectionBlackIp(request *CreateMajorProtectionBlackIpRequest) (_result *CreateMajorProtectionBlackIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMajorProtectionBlackIpResponse{}
	_body, _err := client.CreateMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to use the multi-account management feature of Web Application Firewall (WAF).
//
// @param request - CreateMemberAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemberAccountsResponse
func (client *Client) CreateMemberAccountsWithOptions(request *CreateMemberAccountsRequest, runtime *dara.RuntimeOptions) (_result *CreateMemberAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberAccountIds) {
		query["MemberAccountIds"] = request.MemberAccountIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemberAccounts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemberAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to use the multi-account management feature of Web Application Firewall (WAF).
//
// @param request - CreateMemberAccountsRequest
//
// @return CreateMemberAccountsResponse
func (client *Client) CreateMemberAccounts(request *CreateMemberAccountsRequest) (_result *CreateMemberAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMemberAccountsResponse{}
	_body, _err := client.CreateMemberAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go Web Application Firewall (WAF) 3.0 instance.
//
// @param request - CreatePostpaidInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePostpaidInstanceResponse
func (client *Client) CreatePostpaidInstanceWithOptions(request *CreatePostpaidInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreatePostpaidInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePostpaidInstance"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePostpaidInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go Web Application Firewall (WAF) 3.0 instance.
//
// @param request - CreatePostpaidInstanceRequest
//
// @return CreatePostpaidInstanceResponse
func (client *Client) CreatePostpaidInstance(request *CreatePostpaidInstanceRequest) (_result *CreatePostpaidInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePostpaidInstanceResponse{}
	_body, _err := client.CreatePostpaidInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a ShangMi (SM) certificate for a domain name that is added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param request - CreateSM2CertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSM2CertResponse
func (client *Client) CreateSM2CertWithOptions(request *CreateSM2CertRequest, runtime *dara.RuntimeOptions) (_result *CreateSM2CertResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertName) {
		query["CertName"] = request.CertName
	}

	if !dara.IsNil(request.EncryptCertificate) {
		query["EncryptCertificate"] = request.EncryptCertificate
	}

	if !dara.IsNil(request.EncryptPrivateKey) {
		query["EncryptPrivateKey"] = request.EncryptPrivateKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SignCertificate) {
		query["SignCertificate"] = request.SignCertificate
	}

	if !dara.IsNil(request.SignPrivateKey) {
		query["SignPrivateKey"] = request.SignPrivateKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSM2Cert"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSM2CertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a ShangMi (SM) certificate for a domain name that is added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param request - CreateSM2CertRequest
//
// @return CreateSM2CertResponse
func (client *Client) CreateSM2Cert(request *CreateSM2CertRequest) (_result *CreateSM2CertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSM2CertResponse{}
	_body, _err := client.CreateSM2CertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple risks detected by the API security module at a time.
//
// @param request - DeleteApisecAbnormalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApisecAbnormalsResponse
func (client *Client) DeleteApisecAbnormalsWithOptions(request *DeleteApisecAbnormalsRequest, runtime *dara.RuntimeOptions) (_result *DeleteApisecAbnormalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbnormalIds) {
		query["AbnormalIds"] = request.AbnormalIds
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApisecAbnormals"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApisecAbnormalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple risks detected by the API security module at a time.
//
// @param request - DeleteApisecAbnormalsRequest
//
// @return DeleteApisecAbnormalsResponse
func (client *Client) DeleteApisecAbnormals(request *DeleteApisecAbnormalsRequest) (_result *DeleteApisecAbnormalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApisecAbnormalsResponse{}
	_body, _err := client.DeleteApisecAbnormalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple security events detected by the API security module at a time.
//
// @param request - DeleteApisecEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApisecEventsResponse
func (client *Client) DeleteApisecEventsWithOptions(request *DeleteApisecEventsRequest, runtime *dara.RuntimeOptions) (_result *DeleteApisecEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EventIds) {
		query["EventIds"] = request.EventIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApisecEvents"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApisecEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple security events detected by the API security module at a time.
//
// @param request - DeleteApisecEventsRequest
//
// @return DeleteApisecEventsResponse
func (client *Client) DeleteApisecEvents(request *DeleteApisecEventsRequest) (_result *DeleteApisecEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApisecEventsResponse{}
	_body, _err := client.DeleteApisecEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a service from Web Application Firewall (WAF). This operation is supported for only the Elastic Compute Service (ECS) and Classic Load Balancer (CLB) services.
//
// @param request - DeleteCloudResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudResourceResponse
func (client *Client) DeleteCloudResourceWithOptions(request *DeleteCloudResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCloudResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a service from Web Application Firewall (WAF). This operation is supported for only the Elastic Compute Service (ECS) and Classic Load Balancer (CLB) services.
//
// @param request - DeleteCloudResourceRequest
//
// @return DeleteCloudResourceResponse
func (client *Client) DeleteCloudResource(request *DeleteCloudResourceRequest) (_result *DeleteCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCloudResourceResponse{}
	_body, _err := client.DeleteCloudResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除单个防护对象
//
// @param request - DeleteDefenseResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDefenseResourceResponse
func (client *Client) DeleteDefenseResourceWithOptions(request *DeleteDefenseResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDefenseResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDefenseResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDefenseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除单个防护对象
//
// @param request - DeleteDefenseResourceRequest
//
// @return DeleteDefenseResourceResponse
func (client *Client) DeleteDefenseResource(request *DeleteDefenseResourceRequest) (_result *DeleteDefenseResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDefenseResourceResponse{}
	_body, _err := client.DeleteDefenseResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a protected object group.
//
// @param request - DeleteDefenseResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDefenseResourceGroupResponse
func (client *Client) DeleteDefenseResourceGroupWithOptions(request *DeleteDefenseResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDefenseResourceGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protected object group.
//
// @param request - DeleteDefenseResourceGroupRequest
//
// @return DeleteDefenseResourceGroupResponse
func (client *Client) DeleteDefenseResourceGroup(request *DeleteDefenseResourceGroupRequest) (_result *DeleteDefenseResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDefenseResourceGroupResponse{}
	_body, _err := client.DeleteDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a protection rule.
//
// @param request - DeleteDefenseRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDefenseRuleResponse
func (client *Client) DeleteDefenseRuleWithOptions(request *DeleteDefenseRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDefenseRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDefenseRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protection rule.
//
// @param request - DeleteDefenseRuleRequest
//
// @return DeleteDefenseRuleResponse
func (client *Client) DeleteDefenseRule(request *DeleteDefenseRuleRequest) (_result *DeleteDefenseRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDefenseRuleResponse{}
	_body, _err := client.DeleteDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新防护规则封禁Ip
//
// @param request - DeleteDefenseRuleBlockIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDefenseRuleBlockIpResponse
func (client *Client) DeleteDefenseRuleBlockIpWithOptions(request *DeleteDefenseRuleBlockIpRequest, runtime *dara.RuntimeOptions) (_result *DeleteDefenseRuleBlockIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDefenseRuleBlockIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDefenseRuleBlockIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新防护规则封禁Ip
//
// @param request - DeleteDefenseRuleBlockIpRequest
//
// @return DeleteDefenseRuleBlockIpResponse
func (client *Client) DeleteDefenseRuleBlockIp(request *DeleteDefenseRuleBlockIpRequest) (_result *DeleteDefenseRuleBlockIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDefenseRuleBlockIpResponse{}
	_body, _err := client.DeleteDefenseRuleBlockIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a protection rule template.
//
// @param request - DeleteDefenseTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDefenseTemplateResponse
func (client *Client) DeleteDefenseTemplateWithOptions(request *DeleteDefenseTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDefenseTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDefenseTemplate"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a protection rule template.
//
// @param request - DeleteDefenseTemplateRequest
//
// @return DeleteDefenseTemplateResponse
func (client *Client) DeleteDefenseTemplate(request *DeleteDefenseTemplateRequest) (_result *DeleteDefenseTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDefenseTemplateResponse{}
	_body, _err := client.DeleteDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a domain name that is added to Web Application Firewall (WAF).
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name that is added to Web Application Firewall (WAF).
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除集群规则信息
//
// @param request - DeleteHybridCloudClusterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHybridCloudClusterRuleResponse
func (client *Client) DeleteHybridCloudClusterRuleWithOptions(request *DeleteHybridCloudClusterRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteHybridCloudClusterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterRuleResourceId) {
		query["ClusterRuleResourceId"] = request.ClusterRuleResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHybridCloudClusterRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHybridCloudClusterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除集群规则信息
//
// @param request - DeleteHybridCloudClusterRuleRequest
//
// @return DeleteHybridCloudClusterRuleResponse
func (client *Client) DeleteHybridCloudClusterRule(request *DeleteHybridCloudClusterRuleRequest) (_result *DeleteHybridCloudClusterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHybridCloudClusterRuleResponse{}
	_body, _err := client.DeleteHybridCloudClusterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IP address blacklist for major event protection.
//
// @param request - DeleteMajorProtectionBlackIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMajorProtectionBlackIpResponse
func (client *Client) DeleteMajorProtectionBlackIpWithOptions(request *DeleteMajorProtectionBlackIpRequest, runtime *dara.RuntimeOptions) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IpList) {
		query["IpList"] = request.IpList
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMajorProtectionBlackIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IP address blacklist for major event protection.
//
// @param request - DeleteMajorProtectionBlackIpRequest
//
// @return DeleteMajorProtectionBlackIpResponse
func (client *Client) DeleteMajorProtectionBlackIp(request *DeleteMajorProtectionBlackIpRequest) (_result *DeleteMajorProtectionBlackIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMajorProtectionBlackIpResponse{}
	_body, _err := client.DeleteMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes the members that are added for multi-account management in Web Application Firewall (WAF).
//
// @param request - DeleteMemberAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemberAccountResponse
func (client *Client) DeleteMemberAccountWithOptions(request *DeleteMemberAccountRequest, runtime *dara.RuntimeOptions) (_result *DeleteMemberAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberAccountId) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemberAccount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemberAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes the members that are added for multi-account management in Web Application Firewall (WAF).
//
// @param request - DeleteMemberAccountRequest
//
// @return DeleteMemberAccountResponse
func (client *Client) DeleteMemberAccount(request *DeleteMemberAccountRequest) (_result *DeleteMemberAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMemberAccountResponse{}
	_body, _err := client.DeleteMemberAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询异常的云产品接入资源
//
// @param request - DescribeAbnormalCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAbnormalCloudResourcesResponse
func (client *Client) DescribeAbnormalCloudResourcesWithOptions(request *DescribeAbnormalCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAbnormalCloudResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAbnormalCloudResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAbnormalCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询异常的云产品接入资源
//
// @param request - DescribeAbnormalCloudResourcesRequest
//
// @return DescribeAbnormalCloudResourcesResponse
func (client *Client) DescribeAbnormalCloudResources(request *DescribeAbnormalCloudResourcesRequest) (_result *DescribeAbnormalCloudResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAbnormalCloudResourcesResponse{}
	_body, _err := client.DescribeAbnormalCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether an Alibaba Cloud account is the delegated administrator account of a Web Application Firewall (WAF) instance.
//
// @param request - DescribeAccountDelegatedStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountDelegatedStatusResponse
func (client *Client) DescribeAccountDelegatedStatusWithOptions(request *DescribeAccountDelegatedStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccountDelegatedStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccountDelegatedStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccountDelegatedStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether an Alibaba Cloud account is the delegated administrator account of a Web Application Firewall (WAF) instance.
//
// @param request - DescribeAccountDelegatedStatusRequest
//
// @return DescribeAccountDelegatedStatusResponse
func (client *Client) DescribeAccountDelegatedStatus(request *DescribeAccountDelegatedStatusRequest) (_result *DescribeAccountDelegatedStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccountDelegatedStatusResponse{}
	_body, _err := client.DescribeAccountDelegatedStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of data export tasks in the API security module.
//
// @param request - DescribeApiExportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiExportsResponse
func (client *Client) DescribeApiExportsWithOptions(request *DescribeApiExportsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiExportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiExports"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiExportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of data export tasks in the API security module.
//
// @param request - DescribeApiExportsRequest
//
// @return DescribeApiExportsResponse
func (client *Client) DescribeApiExports(request *DescribeApiExportsRequest) (_result *DescribeApiExportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiExportsResponse{}
	_body, _err := client.DescribeApiExportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which risks are detected by the API security module.
//
// @param request - DescribeApisecAbnormalDomainStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecAbnormalDomainStatisticResponse
func (client *Client) DescribeApisecAbnormalDomainStatisticWithOptions(request *DescribeApisecAbnormalDomainStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecAbnormalDomainStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecAbnormalDomainStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecAbnormalDomainStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which risks are detected by the API security module.
//
// @param request - DescribeApisecAbnormalDomainStatisticRequest
//
// @return DescribeApisecAbnormalDomainStatisticResponse
func (client *Client) DescribeApisecAbnormalDomainStatistic(request *DescribeApisecAbnormalDomainStatisticRequest) (_result *DescribeApisecAbnormalDomainStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecAbnormalDomainStatisticResponse{}
	_body, _err := client.DescribeApisecAbnormalDomainStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of API security risks.
//
// @param request - DescribeApisecAbnormalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecAbnormalsResponse
func (client *Client) DescribeApisecAbnormalsWithOptions(request *DescribeApisecAbnormalsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecAbnormalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbnormalId) {
		query["AbnormalId"] = request.AbnormalId
	}

	if !dara.IsNil(request.AbnormalLevel) {
		query["AbnormalLevel"] = request.AbnormalLevel
	}

	if !dara.IsNil(request.AbnormalTag) {
		query["AbnormalTag"] = request.AbnormalTag
	}

	if !dara.IsNil(request.ApiFormat) {
		query["ApiFormat"] = request.ApiFormat
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiTag) {
		query["ApiTag"] = request.ApiTag
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.OrderKey) {
		query["OrderKey"] = request.OrderKey
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecAbnormals"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecAbnormalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of API security risks.
//
// @param request - DescribeApisecAbnormalsRequest
//
// @return DescribeApisecAbnormalsResponse
func (client *Client) DescribeApisecAbnormals(request *DescribeApisecAbnormalsRequest) (_result *DescribeApisecAbnormalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecAbnormalsResponse{}
	_body, _err := client.DescribeApisecAbnormalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries API assets in the API security module.
//
// @param request - DescribeApisecApiResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecApiResourcesResponse
func (client *Client) DescribeApisecApiResourcesWithOptions(request *DescribeApisecApiResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecApiResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiFormat) {
		query["ApiFormat"] = request.ApiFormat
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiMethod) {
		query["ApiMethod"] = request.ApiMethod
	}

	if !dara.IsNil(request.ApiStatus) {
		query["ApiStatus"] = request.ApiStatus
	}

	if !dara.IsNil(request.ApiTag) {
		query["ApiTag"] = request.ApiTag
	}

	if !dara.IsNil(request.ApiType) {
		query["ApiType"] = request.ApiType
	}

	if !dara.IsNil(request.AuthFlag) {
		query["AuthFlag"] = request.AuthFlag
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Follow) {
		query["Follow"] = request.Follow
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.OrderKey) {
		query["OrderKey"] = request.OrderKey
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestSensitiveType) {
		query["RequestSensitiveType"] = request.RequestSensitiveType
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SensitiveLevel) {
		query["SensitiveLevel"] = request.SensitiveLevel
	}

	if !dara.IsNil(request.SensitiveType) {
		query["SensitiveType"] = request.SensitiveType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecApiResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecApiResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API assets in the API security module.
//
// @param request - DescribeApisecApiResourcesRequest
//
// @return DescribeApisecApiResourcesResponse
func (client *Client) DescribeApisecApiResources(request *DescribeApisecApiResourcesRequest) (_result *DescribeApisecApiResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecApiResourcesResponse{}
	_body, _err := client.DescribeApisecApiResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the asset trends in the API security module.
//
// @param request - DescribeApisecAssetTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecAssetTrendResponse
func (client *Client) DescribeApisecAssetTrendWithOptions(request *DescribeApisecAssetTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecAssetTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecAssetTrend"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecAssetTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the asset trends in the API security module.
//
// @param request - DescribeApisecAssetTrendRequest
//
// @return DescribeApisecAssetTrendResponse
func (client *Client) DescribeApisecAssetTrend(request *DescribeApisecAssetTrendRequest) (_result *DescribeApisecAssetTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecAssetTrendResponse{}
	_body, _err := client.DescribeApisecAssetTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which security events are detected by the API security module.
//
// @param request - DescribeApisecEventDomainStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecEventDomainStatisticResponse
func (client *Client) DescribeApisecEventDomainStatisticWithOptions(request *DescribeApisecEventDomainStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecEventDomainStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecEventDomainStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecEventDomainStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which security events are detected by the API security module.
//
// @param request - DescribeApisecEventDomainStatisticRequest
//
// @return DescribeApisecEventDomainStatisticResponse
func (client *Client) DescribeApisecEventDomainStatistic(request *DescribeApisecEventDomainStatisticRequest) (_result *DescribeApisecEventDomainStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecEventDomainStatisticResponse{}
	_body, _err := client.DescribeApisecEventDomainStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries API security events.
//
// @param request - DescribeApisecEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecEventsResponse
func (client *Client) DescribeApisecEventsWithOptions(request *DescribeApisecEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiFormat) {
		query["ApiFormat"] = request.ApiFormat
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiTag) {
		query["ApiTag"] = request.ApiTag
	}

	if !dara.IsNil(request.AttackIp) {
		query["AttackIp"] = request.AttackIp
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.EventLevel) {
		query["EventLevel"] = request.EventLevel
	}

	if !dara.IsNil(request.EventTag) {
		query["EventTag"] = request.EventTag
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.OrderKey) {
		query["OrderKey"] = request.OrderKey
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecEvents"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API security events.
//
// @param request - DescribeApisecEventsRequest
//
// @return DescribeApisecEventsResponse
func (client *Client) DescribeApisecEvents(request *DescribeApisecEventsRequest) (_result *DescribeApisecEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecEventsResponse{}
	_body, _err := client.DescribeApisecEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of API security log subscription.
//
// @param request - DescribeApisecLogDeliveriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecLogDeliveriesResponse
func (client *Client) DescribeApisecLogDeliveriesWithOptions(request *DescribeApisecLogDeliveriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecLogDeliveriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecLogDeliveries"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecLogDeliveriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of API security log subscription.
//
// @param request - DescribeApisecLogDeliveriesRequest
//
// @return DescribeApisecLogDeliveriesResponse
func (client *Client) DescribeApisecLogDeliveries(request *DescribeApisecLogDeliveriesRequest) (_result *DescribeApisecLogDeliveriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecLogDeliveriesResponse{}
	_body, _err := client.DescribeApisecLogDeliveriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of domain names detected in the API security module.
//
// @param request - DescribeApisecMatchedHostsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecMatchedHostsResponse
func (client *Client) DescribeApisecMatchedHostsWithOptions(request *DescribeApisecMatchedHostsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecMatchedHostsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecMatchedHosts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecMatchedHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of domain names detected in the API security module.
//
// @param request - DescribeApisecMatchedHostsRequest
//
// @return DescribeApisecMatchedHostsResponse
func (client *Client) DescribeApisecMatchedHosts(request *DescribeApisecMatchedHostsRequest) (_result *DescribeApisecMatchedHostsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecMatchedHostsResponse{}
	_body, _err := client.DescribeApisecMatchedHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of protected object groups to which API security policies are applied.
//
// @param request - DescribeApisecProtectionGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecProtectionGroupsResponse
func (client *Client) DescribeApisecProtectionGroupsWithOptions(request *DescribeApisecProtectionGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecProtectionGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApisecStatus) {
		query["ApisecStatus"] = request.ApisecStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecProtectionGroups"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecProtectionGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of protected object groups to which API security policies are applied.
//
// @param request - DescribeApisecProtectionGroupsRequest
//
// @return DescribeApisecProtectionGroupsResponse
func (client *Client) DescribeApisecProtectionGroups(request *DescribeApisecProtectionGroupsRequest) (_result *DescribeApisecProtectionGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecProtectionGroupsResponse{}
	_body, _err := client.DescribeApisecProtectionGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of protected objects to which API security policies are applied.
//
// @param request - DescribeApisecProtectionResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecProtectionResourcesResponse
func (client *Client) DescribeApisecProtectionResourcesWithOptions(request *DescribeApisecProtectionResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecProtectionResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApisecStatus) {
		query["ApisecStatus"] = request.ApisecStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecProtectionResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecProtectionResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of protected objects to which API security policies are applied.
//
// @param request - DescribeApisecProtectionResourcesRequest
//
// @return DescribeApisecProtectionResourcesResponse
func (client *Client) DescribeApisecProtectionResources(request *DescribeApisecProtectionResourcesRequest) (_result *DescribeApisecProtectionResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecProtectionResourcesResponse{}
	_body, _err := client.DescribeApisecProtectionResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the policies configured in the API security module.
//
// @param request - DescribeApisecRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecRulesResponse
func (client *Client) DescribeApisecRulesWithOptions(request *DescribeApisecRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecRules"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the policies configured in the API security module.
//
// @param request - DescribeApisecRulesRequest
//
// @return DescribeApisecRulesResponse
func (client *Client) DescribeApisecRules(request *DescribeApisecRulesRequest) (_result *DescribeApisecRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecRulesResponse{}
	_body, _err := client.DescribeApisecRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which sensitive data is detected by the API security module.
//
// @param request - DescribeApisecSensitiveDomainStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecSensitiveDomainStatisticResponse
func (client *Client) DescribeApisecSensitiveDomainStatisticWithOptions(request *DescribeApisecSensitiveDomainStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecSensitiveDomainStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
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
		Action:      dara.String("DescribeApisecSensitiveDomainStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecSensitiveDomainStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on domain names on which sensitive data is detected by the API security module.
//
// @param request - DescribeApisecSensitiveDomainStatisticRequest
//
// @return DescribeApisecSensitiveDomainStatisticResponse
func (client *Client) DescribeApisecSensitiveDomainStatistic(request *DescribeApisecSensitiveDomainStatisticRequest) (_result *DescribeApisecSensitiveDomainStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecSensitiveDomainStatisticResponse{}
	_body, _err := client.DescribeApisecSensitiveDomainStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Logstores whose names start with apisec- in Simple Log Service.
//
// @param request - DescribeApisecSlsLogStoresRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecSlsLogStoresResponse
func (client *Client) DescribeApisecSlsLogStoresWithOptions(request *DescribeApisecSlsLogStoresRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecSlsLogStoresResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogRegionId) {
		query["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecSlsLogStores"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecSlsLogStoresResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Logstores whose names start with apisec- in Simple Log Service.
//
// @param request - DescribeApisecSlsLogStoresRequest
//
// @return DescribeApisecSlsLogStoresResponse
func (client *Client) DescribeApisecSlsLogStores(request *DescribeApisecSlsLogStoresRequest) (_result *DescribeApisecSlsLogStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecSlsLogStoresResponse{}
	_body, _err := client.DescribeApisecSlsLogStoresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the projects whose names start with apisec- in Simple Log Service.
//
// @param request - DescribeApisecSlsProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecSlsProjectsResponse
func (client *Client) DescribeApisecSlsProjectsWithOptions(request *DescribeApisecSlsProjectsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecSlsProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogRegionId) {
		query["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecSlsProjects"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecSlsProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the projects whose names start with apisec- in Simple Log Service.
//
// @param request - DescribeApisecSlsProjectsRequest
//
// @return DescribeApisecSlsProjectsResponse
func (client *Client) DescribeApisecSlsProjects(request *DescribeApisecSlsProjectsRequest) (_result *DescribeApisecSlsProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecSlsProjectsResponse{}
	_body, _err := client.DescribeApisecSlsProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of API security-related risks and events.
//
// @param request - DescribeApisecStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecStatisticsResponse
func (client *Client) DescribeApisecStatisticsWithOptions(request *DescribeApisecStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserStatusList) {
		query["UserStatusList"] = request.UserStatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecStatistics"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of API security-related risks and events.
//
// @param request - DescribeApisecStatisticsRequest
//
// @return DescribeApisecStatisticsResponse
func (client *Client) DescribeApisecStatistics(request *DescribeApisecStatisticsRequest) (_result *DescribeApisecStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecStatisticsResponse{}
	_body, _err := client.DescribeApisecStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protection suggestions for APIs.
//
// @param request - DescribeApisecSuggestionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecSuggestionsResponse
func (client *Client) DescribeApisecSuggestionsWithOptions(request *DescribeApisecSuggestionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecSuggestionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecSuggestions"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecSuggestionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protection suggestions for APIs.
//
// @param request - DescribeApisecSuggestionsRequest
//
// @return DescribeApisecSuggestionsResponse
func (client *Client) DescribeApisecSuggestions(request *DescribeApisecSuggestionsRequest) (_result *DescribeApisecSuggestionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecSuggestionsResponse{}
	_body, _err := client.DescribeApisecSuggestionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user operation records in the API security module.
//
// @param request - DescribeApisecUserOperationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisecUserOperationsResponse
func (client *Client) DescribeApisecUserOperationsWithOptions(request *DescribeApisecUserOperationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisecUserOperationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisecUserOperations"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisecUserOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user operation records in the API security module.
//
// @param request - DescribeApisecUserOperationsRequest
//
// @return DescribeApisecUserOperationsResponse
func (client *Client) DescribeApisecUserOperations(request *DescribeApisecUserOperationsRequest) (_result *DescribeApisecUserOperationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisecUserOperationsResponse{}
	_body, _err := client.DescribeApisecUserOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询基础防护系统规则集
//
// @param request - DescribeBaseSystemRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBaseSystemRulesResponse
func (client *Client) DescribeBaseSystemRulesWithOptions(request *DescribeBaseSystemRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeBaseSystemRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DetectType) {
		query["DetectType"] = request.DetectType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.RuleAction) {
		query["RuleAction"] = request.RuleAction
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBaseSystemRules"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBaseSystemRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询基础防护系统规则集
//
// @param request - DescribeBaseSystemRulesRequest
//
// @return DescribeBaseSystemRulesResponse
func (client *Client) DescribeBaseSystemRules(request *DescribeBaseSystemRulesRequest) (_result *DescribeBaseSystemRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBaseSystemRulesResponse{}
	_body, _err := client.DescribeBaseSystemRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate, such as the certificate name, expiration time, issuance time, and associated domain name.
//
// @param request - DescribeCertDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertDetailResponse
func (client *Client) DescribeCertDetailWithOptions(request *DescribeCertDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertIdentifier) {
		query["CertIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCertDetail"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a certificate, such as the certificate name, expiration time, issuance time, and associated domain name.
//
// @param request - DescribeCertDetailRequest
//
// @return DescribeCertDetailResponse
func (client *Client) DescribeCertDetail(request *DescribeCertDetailRequest) (_result *DescribeCertDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertDetailResponse{}
	_body, _err := client.DescribeCertDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificates issued for your domain names that are added to Web Application Firewall (WAF).
//
// @param request - DescribeCertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCertsResponse
func (client *Client) DescribeCertsWithOptions(request *DescribeCertsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCerts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates issued for your domain names that are added to Web Application Firewall (WAF).
//
// @param request - DescribeCertsRequest
//
// @return DescribeCertsResponse
func (client *Client) DescribeCerts(request *DescribeCertsRequest) (_result *DescribeCertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCertsResponse{}
	_body, _err := client.DescribeCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a port of the cloud service that is added to Web Application Firewall (WAF). This operation is supported for only Elastic Compute Service (ECS) and Classic Load Balancer (CLB).
//
// @param request - DescribeCloudResourceAccessPortDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourceAccessPortDetailsResponse
func (client *Client) DescribeCloudResourceAccessPortDetailsWithOptions(request *DescribeCloudResourceAccessPortDetailsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudResourceAccessPortDetailsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudResourceAccessPortDetails"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudResourceAccessPortDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a port of the cloud service that is added to Web Application Firewall (WAF). This operation is supported for only Elastic Compute Service (ECS) and Classic Load Balancer (CLB).
//
// @param request - DescribeCloudResourceAccessPortDetailsRequest
//
// @return DescribeCloudResourceAccessPortDetailsResponse
func (client *Client) DescribeCloudResourceAccessPortDetails(request *DescribeCloudResourceAccessPortDetailsRequest) (_result *DescribeCloudResourceAccessPortDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudResourceAccessPortDetailsResponse{}
	_body, _err := client.DescribeCloudResourceAccessPortDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ports of the cloud service that is added to Web Application Firewall (WAF). This operation is supported for only Elastic Compute Service (ECS) and Classic Load Balancer (CLB).
//
// @param request - DescribeCloudResourceAccessedPortsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourceAccessedPortsResponse
func (client *Client) DescribeCloudResourceAccessedPortsWithOptions(request *DescribeCloudResourceAccessedPortsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudResourceAccessedPortsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudResourceAccessedPorts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudResourceAccessedPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports of the cloud service that is added to Web Application Firewall (WAF). This operation is supported for only Elastic Compute Service (ECS) and Classic Load Balancer (CLB).
//
// @param request - DescribeCloudResourceAccessedPortsRequest
//
// @return DescribeCloudResourceAccessedPortsResponse
func (client *Client) DescribeCloudResourceAccessedPorts(request *DescribeCloudResourceAccessedPortsRequest) (_result *DescribeCloudResourceAccessedPortsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudResourceAccessedPortsResponse{}
	_body, _err := client.DescribeCloudResourceAccessedPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries cloud service resources that are added to Web Application Firewall (WAF).
//
// @param request - DescribeCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudResourcesResponse
func (client *Client) DescribeCloudResourcesWithOptions(request *DescribeCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceDomain) {
		query["ResourceDomain"] = request.ResourceDomain
	}

	if !dara.IsNil(request.ResourceFunction) {
		query["ResourceFunction"] = request.ResourceFunction
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceInstanceName) {
		query["ResourceInstanceName"] = request.ResourceInstanceName
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceRouteName) {
		query["ResourceRouteName"] = request.ResourceRouteName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries cloud service resources that are added to Web Application Firewall (WAF).
//
// @param request - DescribeCloudResourcesRequest
//
// @return DescribeCloudResourcesResponse
func (client *Client) DescribeCloudResources(request *DescribeCloudResourcesRequest) (_result *DescribeCloudResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCloudResourcesResponse{}
	_body, _err := client.DescribeCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the total number of domain names that are added to Web Application Firewall (WAF) in CNAME record mode and hybrid cloud reverse proxy mode.
//
// @param request - DescribeCnameCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCnameCountResponse
func (client *Client) DescribeCnameCountWithOptions(request *DescribeCnameCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeCnameCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCnameCount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCnameCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the total number of domain names that are added to Web Application Firewall (WAF) in CNAME record mode and hybrid cloud reverse proxy mode.
//
// @param request - DescribeCnameCountRequest
//
// @return DescribeCnameCountResponse
func (client *Client) DescribeCnameCount(request *DescribeCnameCountRequest) (_result *DescribeCnameCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCnameCountResponse{}
	_body, _err := client.DescribeCnameCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询日志服务支持的所有字段
//
// @param tmpReq - DescribeCommonLogFieldsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommonLogFieldsResponse
func (client *Client) DescribeCommonLogFieldsWithOptions(tmpReq *DescribeCommonLogFieldsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCommonLogFieldsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCommonLogFieldsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LogKeyList) {
		request.LogKeyListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LogKeyList, dara.String("LogKeyList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.IsRequired) {
		query["IsRequired"] = request.IsRequired
	}

	if !dara.IsNil(request.LogKeyListShrink) {
		query["LogKeyList"] = request.LogKeyListShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCommonLogFields"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCommonLogFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日志服务支持的所有字段
//
// @param request - DescribeCommonLogFieldsRequest
//
// @return DescribeCommonLogFieldsResponse
func (client *Client) DescribeCommonLogFields(request *DescribeCommonLogFieldsRequest) (_result *DescribeCommonLogFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCommonLogFieldsResponse{}
	_body, _err := client.DescribeCommonLogFieldsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义正则规则编译结果
//
// @param request - DescribeCustomBaseRuleCompileResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomBaseRuleCompileResultResponse
func (client *Client) DescribeCustomBaseRuleCompileResultWithOptions(request *DescribeCustomBaseRuleCompileResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomBaseRuleCompileResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomBaseRuleCompileResult"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomBaseRuleCompileResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义正则规则编译结果
//
// @param request - DescribeCustomBaseRuleCompileResultRequest
//
// @return DescribeCustomBaseRuleCompileResultResponse
func (client *Client) DescribeCustomBaseRuleCompileResult(request *DescribeCustomBaseRuleCompileResultRequest) (_result *DescribeCustomBaseRuleCompileResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCustomBaseRuleCompileResultResponse{}
	_body, _err := client.DescribeCustomBaseRuleCompileResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether DDoS attacks occur on specific domain names protected by a Web Application Firewall (WAF) instance.
//
// @param request - DescribeDDoSStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDDoSStatusResponse
func (client *Client) DescribeDDoSStatusWithOptions(request *DescribeDDoSStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeDDoSStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDDoSStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDDoSStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether DDoS attacks occur on specific domain names protected by a Web Application Firewall (WAF) instance.
//
// @param request - DescribeDDoSStatusRequest
//
// @return DescribeDDoSStatusResponse
func (client *Client) DescribeDDoSStatus(request *DescribeDDoSStatusRequest) (_result *DescribeDDoSStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDDoSStatusResponse{}
	_body, _err := client.DescribeDDoSStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default SSL and Transport Layer Security (TLS) settings.
//
// @param request - DescribeDefaultHttpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefaultHttpsResponse
func (client *Client) DescribeDefaultHttpsWithOptions(request *DescribeDefaultHttpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefaultHttpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefaultHttps"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefaultHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default SSL and Transport Layer Security (TLS) settings.
//
// @param request - DescribeDefaultHttpsRequest
//
// @return DescribeDefaultHttpsResponse
func (client *Client) DescribeDefaultHttps(request *DescribeDefaultHttpsRequest) (_result *DescribeDefaultHttpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefaultHttpsResponse{}
	_body, _err := client.DescribeDefaultHttpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询可以被防护组绑定的防护对象列表
//
// @param request - DescribeDefenseGroupValidResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseGroupValidResourcesResponse
func (client *Client) DescribeDefenseGroupValidResourcesWithOptions(request *DescribeDefenseGroupValidResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseGroupValidResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseGroupValidResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseGroupValidResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询可以被防护组绑定的防护对象列表
//
// @param request - DescribeDefenseGroupValidResourcesRequest
//
// @return DescribeDefenseGroupValidResourcesResponse
func (client *Client) DescribeDefenseGroupValidResources(request *DescribeDefenseGroupValidResourcesRequest) (_result *DescribeDefenseGroupValidResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseGroupValidResourcesResponse{}
	_body, _err := client.DescribeDefenseGroupValidResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a protected object.
//
// @param request - DescribeDefenseResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceResponse
func (client *Client) DescribeDefenseResourceWithOptions(request *DescribeDefenseResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a protected object.
//
// @param request - DescribeDefenseResourceRequest
//
// @return DescribeDefenseResourceResponse
func (client *Client) DescribeDefenseResource(request *DescribeDefenseResourceRequest) (_result *DescribeDefenseResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceResponse{}
	_body, _err := client.DescribeDefenseResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a protected object group.
//
// @param request - DescribeDefenseResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceGroupResponse
func (client *Client) DescribeDefenseResourceGroupWithOptions(request *DescribeDefenseResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a protected object group.
//
// @param request - DescribeDefenseResourceGroupRequest
//
// @return DescribeDefenseResourceGroupResponse
func (client *Client) DescribeDefenseResourceGroup(request *DescribeDefenseResourceGroupRequest) (_result *DescribeDefenseResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupResponse{}
	_body, _err := client.DescribeDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of protected object groups.
//
// @param request - DescribeDefenseResourceGroupNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceGroupNamesResponse
func (client *Client) DescribeDefenseResourceGroupNamesWithOptions(request *DescribeDefenseResourceGroupNamesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceGroupNamesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupNameLike) {
		query["GroupNameLike"] = request.GroupNameLike
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceGroupNames"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceGroupNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of protected object groups.
//
// @param request - DescribeDefenseResourceGroupNamesRequest
//
// @return DescribeDefenseResourceGroupNamesResponse
func (client *Client) DescribeDefenseResourceGroupNames(request *DescribeDefenseResourceGroupNamesRequest) (_result *DescribeDefenseResourceGroupNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupNamesResponse{}
	_body, _err := client.DescribeDefenseResourceGroupNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a pagination query to retrieve the information about protected object groups.
//
// @param request - DescribeDefenseResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceGroupsResponse
func (client *Client) DescribeDefenseResourceGroupsWithOptions(request *DescribeDefenseResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupNameLike) {
		query["GroupNameLike"] = request.GroupNameLike
	}

	if !dara.IsNil(request.GroupNames) {
		query["GroupNames"] = request.GroupNames
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceGroups"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a pagination query to retrieve the information about protected object groups.
//
// @param request - DescribeDefenseResourceGroupsRequest
//
// @return DescribeDefenseResourceGroupsResponse
func (client *Client) DescribeDefenseResourceGroups(request *DescribeDefenseResourceGroupsRequest) (_result *DescribeDefenseResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceGroupsResponse{}
	_body, _err := client.DescribeDefenseResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a pagination query to retrieve the names of protected objects.
//
// @param request - DescribeDefenseResourceNamesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceNamesResponse
func (client *Client) DescribeDefenseResourceNamesWithOptions(request *DescribeDefenseResourceNamesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceNamesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceNames"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a pagination query to retrieve the names of protected objects.
//
// @param request - DescribeDefenseResourceNamesRequest
//
// @return DescribeDefenseResourceNamesResponse
func (client *Client) DescribeDefenseResourceNames(request *DescribeDefenseResourceNamesRequest) (_result *DescribeDefenseResourceNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceNamesResponse{}
	_body, _err := client.DescribeDefenseResourceNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询防护对象和所属资源的关系
//
// @param request - DescribeDefenseResourceOwnerUidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceOwnerUidResponse
func (client *Client) DescribeDefenseResourceOwnerUidWithOptions(request *DescribeDefenseResourceOwnerUidRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceOwnerUidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceNames) {
		query["ResourceNames"] = request.ResourceNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceOwnerUid"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceOwnerUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询防护对象和所属资源的关系
//
// @param request - DescribeDefenseResourceOwnerUidRequest
//
// @return DescribeDefenseResourceOwnerUidResponse
func (client *Client) DescribeDefenseResourceOwnerUid(request *DescribeDefenseResourceOwnerUidRequest) (_result *DescribeDefenseResourceOwnerUidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceOwnerUidResponse{}
	_body, _err := client.DescribeDefenseResourceOwnerUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protection templates that are associated with a protected object or protected object group.
//
// @param request - DescribeDefenseResourceTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourceTemplatesResponse
func (client *Client) DescribeDefenseResourceTemplatesWithOptions(request *DescribeDefenseResourceTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourceTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResourceTemplates"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourceTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protection templates that are associated with a protected object or protected object group.
//
// @param request - DescribeDefenseResourceTemplatesRequest
//
// @return DescribeDefenseResourceTemplatesResponse
func (client *Client) DescribeDefenseResourceTemplates(request *DescribeDefenseResourceTemplatesRequest) (_result *DescribeDefenseResourceTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourceTemplatesResponse{}
	_body, _err := client.DescribeDefenseResourceTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries protected objects by page.
//
// @param request - DescribeDefenseResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseResourcesResponse
func (client *Client) DescribeDefenseResourcesWithOptions(request *DescribeDefenseResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries protected objects by page.
//
// @param request - DescribeDefenseResourcesRequest
//
// @return DescribeDefenseResourcesResponse
func (client *Client) DescribeDefenseResources(request *DescribeDefenseResourcesRequest) (_result *DescribeDefenseResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseResourcesResponse{}
	_body, _err := client.DescribeDefenseResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a protection rule.
//
// @param request - DescribeDefenseRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseRuleResponse
func (client *Client) DescribeDefenseRuleWithOptions(request *DescribeDefenseRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a protection rule.
//
// @param request - DescribeDefenseRuleRequest
//
// @return DescribeDefenseRuleResponse
func (client *Client) DescribeDefenseRule(request *DescribeDefenseRuleRequest) (_result *DescribeDefenseRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseRuleResponse{}
	_body, _err := client.DescribeDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询防护规则的统计信息
//
// @param request - DescribeDefenseRuleStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseRuleStatisticsResponse
func (client *Client) DescribeDefenseRuleStatisticsWithOptions(request *DescribeDefenseRuleStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseRuleStatisticsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FourthKey) {
		query["FourthKey"] = request.FourthKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PrimaryKey) {
		query["PrimaryKey"] = request.PrimaryKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SecondaryKey) {
		query["SecondaryKey"] = request.SecondaryKey
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.ThirdKey) {
		query["ThirdKey"] = request.ThirdKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseRuleStatistics"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseRuleStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询防护规则的统计信息
//
// @param request - DescribeDefenseRuleStatisticsRequest
//
// @return DescribeDefenseRuleStatisticsResponse
func (client *Client) DescribeDefenseRuleStatistics(request *DescribeDefenseRuleStatisticsRequest) (_result *DescribeDefenseRuleStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseRuleStatisticsResponse{}
	_body, _err := client.DescribeDefenseRuleStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries protection rules by page.
//
// @param request - DescribeDefenseRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseRulesResponse
func (client *Client) DescribeDefenseRulesWithOptions(request *DescribeDefenseRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseRules"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries protection rules by page.
//
// @param request - DescribeDefenseRulesRequest
//
// @return DescribeDefenseRulesResponse
func (client *Client) DescribeDefenseRules(request *DescribeDefenseRulesRequest) (_result *DescribeDefenseRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseRulesResponse{}
	_body, _err := client.DescribeDefenseRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户防护场景的配置
//
// @param request - DescribeDefenseSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseSceneConfigResponse
func (client *Client) DescribeDefenseSceneConfigWithOptions(request *DescribeDefenseSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseSceneConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseSceneConfig"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseSceneConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户防护场景的配置
//
// @param request - DescribeDefenseSceneConfigRequest
//
// @return DescribeDefenseSceneConfigResponse
func (client *Client) DescribeDefenseSceneConfig(request *DescribeDefenseSceneConfigRequest) (_result *DescribeDefenseSceneConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseSceneConfigResponse{}
	_body, _err := client.DescribeDefenseSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a protection rule template.
//
// @param request - DescribeDefenseTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseTemplateResponse
func (client *Client) DescribeDefenseTemplateWithOptions(request *DescribeDefenseTemplateRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseTemplate"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a protection rule template.
//
// @param request - DescribeDefenseTemplateRequest
//
// @return DescribeDefenseTemplateResponse
func (client *Client) DescribeDefenseTemplate(request *DescribeDefenseTemplateRequest) (_result *DescribeDefenseTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseTemplateResponse{}
	_body, _err := client.DescribeDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of protected object groups for which a protection template can take effect.
//
// @param request - DescribeDefenseTemplateValidGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseTemplateValidGroupsResponse
func (client *Client) DescribeDefenseTemplateValidGroupsWithOptions(request *DescribeDefenseTemplateValidGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseTemplateValidGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseTemplateValidGroups"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseTemplateValidGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of protected object groups for which a protection template can take effect.
//
// @param request - DescribeDefenseTemplateValidGroupsRequest
//
// @return DescribeDefenseTemplateValidGroupsResponse
func (client *Client) DescribeDefenseTemplateValidGroups(request *DescribeDefenseTemplateValidGroupsRequest) (_result *DescribeDefenseTemplateValidGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseTemplateValidGroupsResponse{}
	_body, _err := client.DescribeDefenseTemplateValidGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询可以被自定义模板绑定的防护对象列表
//
// @param request - DescribeDefenseTemplateValidResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseTemplateValidResourcesResponse
func (client *Client) DescribeDefenseTemplateValidResourcesWithOptions(request *DescribeDefenseTemplateValidResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseTemplateValidResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseTemplateValidResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseTemplateValidResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询可以被自定义模板绑定的防护对象列表
//
// @param request - DescribeDefenseTemplateValidResourcesRequest
//
// @return DescribeDefenseTemplateValidResourcesResponse
func (client *Client) DescribeDefenseTemplateValidResources(request *DescribeDefenseTemplateValidResourcesRequest) (_result *DescribeDefenseTemplateValidResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseTemplateValidResourcesResponse{}
	_body, _err := client.DescribeDefenseTemplateValidResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve protection templates.
//
// @param request - DescribeDefenseTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDefenseTemplatesResponse
func (client *Client) DescribeDefenseTemplatesWithOptions(request *DescribeDefenseTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDefenseTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.DefenseSubScene) {
		query["DefenseSubScene"] = request.DefenseSubScene
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDefenseTemplates"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDefenseTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a paging query to retrieve protection templates.
//
// @param request - DescribeDefenseTemplatesRequest
//
// @return DescribeDefenseTemplatesResponse
func (client *Client) DescribeDefenseTemplates(request *DescribeDefenseTemplatesRequest) (_result *DescribeDefenseTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDefenseTemplatesResponse{}
	_body, _err := client.DescribeDefenseTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the Domain Name System (DNS) settings of a domain name are properly configured.
//
// @param request - DescribeDomainDNSRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainDNSRecordResponse
func (client *Client) DescribeDomainDNSRecordWithOptions(request *DescribeDomainDNSRecordRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainDNSRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainDNSRecord"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainDNSRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the Domain Name System (DNS) settings of a domain name are properly configured.
//
// @param request - DescribeDomainDNSRecordRequest
//
// @return DescribeDomainDNSRecordResponse
func (client *Client) DescribeDomainDNSRecord(request *DescribeDomainDNSRecordRequest) (_result *DescribeDomainDNSRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainDNSRecordResponse{}
	_body, _err := client.DescribeDomainDNSRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a domain name that is added to Web Application Firewall (WAF).
//
// @param request - DescribeDomainDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainDetailResponse
func (client *Client) DescribeDomainDetailWithOptions(request *DescribeDomainDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainDetail"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a domain name that is added to Web Application Firewall (WAF).
//
// @param request - DescribeDomainDetailRequest
//
// @return DescribeDomainDetailResponse
func (client *Client) DescribeDomainDetail(request *DescribeDomainDetailRequest) (_result *DescribeDomainDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainDetailResponse{}
	_body, _err := client.DescribeDomainDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询域名已使用的端口
//
// @param request - DescribeDomainUsedPortsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainUsedPortsResponse
func (client *Client) DescribeDomainUsedPortsWithOptions(request *DescribeDomainUsedPortsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainUsedPortsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainUsedPorts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainUsedPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询域名已使用的端口
//
// @param request - DescribeDomainUsedPortsRequest
//
// @return DescribeDomainUsedPortsResponse
func (client *Client) DescribeDomainUsedPorts(request *DescribeDomainUsedPortsRequest) (_result *DescribeDomainUsedPortsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainUsedPortsResponse{}
	_body, _err := client.DescribeDomainUsedPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that are added to Web Application Firewall (WAF).
//
// @param request - DescribeDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomainsWithOptions(request *DescribeDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Backend) {
		query["Backend"] = request.Backend
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomains"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are added to Web Application Firewall (WAF).
//
// @param request - DescribeDomainsRequest
//
// @return DescribeDomainsResponse
func (client *Client) DescribeDomains(request *DescribeDomainsRequest) (_result *DescribeDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainsResponse{}
	_body, _err := client.DescribeDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic statistics of requests that are forwarded to Web Application Firewall (WAF).
//
// @param request - DescribeFlowChartRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowChartResponse
func (client *Client) DescribeFlowChartWithOptions(request *DescribeFlowChartRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowChartResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowChart"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic statistics of requests that are forwarded to Web Application Firewall (WAF).
//
// @param request - DescribeFlowChartRequest
//
// @return DescribeFlowChartResponse
func (client *Client) DescribeFlowChart(request *DescribeFlowChartRequest) (_result *DescribeFlowChartResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowChartResponse{}
	_body, _err := client.DescribeFlowChartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 protected objects that receive requests.
//
// @param request - DescribeFlowTopResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowTopResourceResponse
func (client *Client) DescribeFlowTopResourceWithOptions(request *DescribeFlowTopResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowTopResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowTopResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 protected objects that receive requests.
//
// @param request - DescribeFlowTopResourceRequest
//
// @return DescribeFlowTopResourceResponse
func (client *Client) DescribeFlowTopResource(request *DescribeFlowTopResourceRequest) (_result *DescribeFlowTopResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowTopResourceResponse{}
	_body, _err := client.DescribeFlowTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 URLs that are used to initiate requests.
//
// @param request - DescribeFlowTopUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowTopUrlResponse
func (client *Client) DescribeFlowTopUrlWithOptions(request *DescribeFlowTopUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowTopUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowTopUrl"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 URLs that are used to initiate requests.
//
// @param request - DescribeFlowTopUrlRequest
//
// @return DescribeFlowTopUrlResponse
func (client *Client) DescribeFlowTopUrl(request *DescribeFlowTopUrlRequest) (_result *DescribeFlowTopUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowTopUrlResponse{}
	_body, _err := client.DescribeFlowTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the asset statistics provided by basic detection in the API security module.
//
// @param request - DescribeFreeUserAssetCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFreeUserAssetCountResponse
func (client *Client) DescribeFreeUserAssetCountWithOptions(request *DescribeFreeUserAssetCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeFreeUserAssetCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFreeUserAssetCount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFreeUserAssetCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the asset statistics provided by basic detection in the API security module.
//
// @param request - DescribeFreeUserAssetCountRequest
//
// @return DescribeFreeUserAssetCountResponse
func (client *Client) DescribeFreeUserAssetCount(request *DescribeFreeUserAssetCountRequest) (_result *DescribeFreeUserAssetCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFreeUserAssetCountResponse{}
	_body, _err := client.DescribeFreeUserAssetCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics of security events that are detected by using the basic detection feature of the API security module.
//
// @param request - DescribeFreeUserEventCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFreeUserEventCountResponse
func (client *Client) DescribeFreeUserEventCountWithOptions(request *DescribeFreeUserEventCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeFreeUserEventCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFreeUserEventCount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFreeUserEventCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics of security events that are detected by using the basic detection feature of the API security module.
//
// @param request - DescribeFreeUserEventCountRequest
//
// @return DescribeFreeUserEventCountResponse
func (client *Client) DescribeFreeUserEventCount(request *DescribeFreeUserEventCountRequest) (_result *DescribeFreeUserEventCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFreeUserEventCountResponse{}
	_body, _err := client.DescribeFreeUserEventCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types of security events on which basic detection is performed in the API security module.
//
// @param request - DescribeFreeUserEventTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFreeUserEventTypesResponse
func (client *Client) DescribeFreeUserEventTypesWithOptions(request *DescribeFreeUserEventTypesRequest, runtime *dara.RuntimeOptions) (_result *DescribeFreeUserEventTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFreeUserEventTypes"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFreeUserEventTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of security events on which basic detection is performed in the API security module.
//
// @param request - DescribeFreeUserEventTypesRequest
//
// @return DescribeFreeUserEventTypesResponse
func (client *Client) DescribeFreeUserEventTypes(request *DescribeFreeUserEventTypesRequest) (_result *DescribeFreeUserEventTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFreeUserEventTypesResponse{}
	_body, _err := client.DescribeFreeUserEventTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of security events on which basic detection is performed in the API security module.
//
// @param request - DescribeFreeUserEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFreeUserEventsResponse
func (client *Client) DescribeFreeUserEventsWithOptions(request *DescribeFreeUserEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFreeUserEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFreeUserEvents"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFreeUserEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of security events on which basic detection is performed in the API security module.
//
// @param request - DescribeFreeUserEventsRequest
//
// @return DescribeFreeUserEventsResponse
func (client *Client) DescribeFreeUserEvents(request *DescribeFreeUserEventsRequest) (_result *DescribeFreeUserEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFreeUserEventsResponse{}
	_body, _err := client.DescribeFreeUserEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the system status of a node in a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudBasicMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudBasicMonitorResponse
func (client *Client) DescribeHybridCloudBasicMonitorWithOptions(request *DescribeHybridCloudBasicMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudBasicMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mid) {
		query["Mid"] = request.Mid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudBasicMonitor"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudBasicMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the system status of a node in a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudBasicMonitorRequest
//
// @return DescribeHybridCloudBasicMonitorResponse
func (client *Client) DescribeHybridCloudBasicMonitor(request *DescribeHybridCloudBasicMonitorRequest) (_result *DescribeHybridCloudBasicMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudBasicMonitorResponse{}
	_body, _err := client.DescribeHybridCloudBasicMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the rule information about a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudClusterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudClusterRuleResponse
func (client *Client) DescribeHybridCloudClusterRuleWithOptions(request *DescribeHybridCloudClusterRuleRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudClusterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudClusterRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudClusterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the rule information about a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudClusterRuleRequest
//
// @return DescribeHybridCloudClusterRuleResponse
func (client *Client) DescribeHybridCloudClusterRule(request *DescribeHybridCloudClusterRuleRequest) (_result *DescribeHybridCloudClusterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudClusterRuleResponse{}
	_body, _err := client.DescribeHybridCloudClusterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集群规则列表
//
// @param request - DescribeHybridCloudClusterRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudClusterRulesResponse
func (client *Client) DescribeHybridCloudClusterRulesWithOptions(request *DescribeHybridCloudClusterRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudClusterRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleContent) {
		query["RuleContent"] = request.RuleContent
	}

	if !dara.IsNil(request.RuleMatchType) {
		query["RuleMatchType"] = request.RuleMatchType
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudClusterRules"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudClusterRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群规则列表
//
// @param request - DescribeHybridCloudClusterRulesRequest
//
// @return DescribeHybridCloudClusterRulesResponse
func (client *Client) DescribeHybridCloudClusterRules(request *DescribeHybridCloudClusterRulesRequest) (_result *DescribeHybridCloudClusterRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudClusterRulesResponse{}
	_body, _err := client.DescribeHybridCloudClusterRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集群机器列表
//
// @param request - DescribeHybridCloudClusterServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudClusterServersResponse
func (client *Client) DescribeHybridCloudClusterServersWithOptions(request *DescribeHybridCloudClusterServersRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudClusterServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudClusterServers"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudClusterServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群机器列表
//
// @param request - DescribeHybridCloudClusterServersRequest
//
// @return DescribeHybridCloudClusterServersResponse
func (client *Client) DescribeHybridCloudClusterServers(request *DescribeHybridCloudClusterServersRequest) (_result *DescribeHybridCloudClusterServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudClusterServersResponse{}
	_body, _err := client.DescribeHybridCloudClusterServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of hybrid cloud clusters.
//
// @param request - DescribeHybridCloudClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudClustersResponse
func (client *Client) DescribeHybridCloudClustersWithOptions(request *DescribeHybridCloudClustersRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudClusters"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of hybrid cloud clusters.
//
// @param request - DescribeHybridCloudClustersRequest
//
// @return DescribeHybridCloudClustersResponse
func (client *Client) DescribeHybridCloudClusters(request *DescribeHybridCloudClustersRequest) (_result *DescribeHybridCloudClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudClustersResponse{}
	_body, _err := client.DescribeHybridCloudClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the hybrid cloud node groups that are added to Web Application Firewall (WAF).
//
// @param request - DescribeHybridCloudGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudGroupsResponse
func (client *Client) DescribeHybridCloudGroupsWithOptions(request *DescribeHybridCloudGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterProxyType) {
		query["ClusterProxyType"] = request.ClusterProxyType
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.GroupType) {
		query["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudGroups"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the hybrid cloud node groups that are added to Web Application Firewall (WAF).
//
// @param request - DescribeHybridCloudGroupsRequest
//
// @return DescribeHybridCloudGroupsResponse
func (client *Client) DescribeHybridCloudGroups(request *DescribeHybridCloudGroupsRequest) (_result *DescribeHybridCloudGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudGroupsResponse{}
	_body, _err := client.DescribeHybridCloudGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of applications running on a hybrid cloud cluster node.
//
// @param request - DescribeHybridCloudProcessMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudProcessMonitorResponse
func (client *Client) DescribeHybridCloudProcessMonitorWithOptions(request *DescribeHybridCloudProcessMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudProcessMonitorResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mid) {
		query["Mid"] = request.Mid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudProcessMonitor"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudProcessMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of applications running on a hybrid cloud cluster node.
//
// @param request - DescribeHybridCloudProcessMonitorRequest
//
// @return DescribeHybridCloudProcessMonitorResponse
func (client *Client) DescribeHybridCloudProcessMonitor(request *DescribeHybridCloudProcessMonitorRequest) (_result *DescribeHybridCloudProcessMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudProcessMonitorResponse{}
	_body, _err := client.DescribeHybridCloudProcessMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询混合云域名详情
//
// @param request - DescribeHybridCloudResourceDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudResourceDetailResponse
func (client *Client) DescribeHybridCloudResourceDetailWithOptions(request *DescribeHybridCloudResourceDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudResourceDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Backend) {
		query["Backend"] = request.Backend
	}

	if !dara.IsNil(request.CnameEnabled) {
		query["CnameEnabled"] = request.CnameEnabled
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudResourceDetail"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudResourceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询混合云域名详情
//
// @param request - DescribeHybridCloudResourceDetailRequest
//
// @return DescribeHybridCloudResourceDetailResponse
func (client *Client) DescribeHybridCloudResourceDetail(request *DescribeHybridCloudResourceDetailRequest) (_result *DescribeHybridCloudResourceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudResourceDetailResponse{}
	_body, _err := client.DescribeHybridCloudResourceDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the domain names that are added to a Web Application Firewall (WAF) instance in hybrid cloud mode.
//
// @param request - DescribeHybridCloudResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudResourcesResponse
func (client *Client) DescribeHybridCloudResourcesWithOptions(request *DescribeHybridCloudResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Backend) {
		query["Backend"] = request.Backend
	}

	if !dara.IsNil(request.CnameEnabled) {
		query["CnameEnabled"] = request.CnameEnabled
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the domain names that are added to a Web Application Firewall (WAF) instance in hybrid cloud mode.
//
// @param request - DescribeHybridCloudResourcesRequest
//
// @return DescribeHybridCloudResourcesResponse
func (client *Client) DescribeHybridCloudResources(request *DescribeHybridCloudResourcesRequest) (_result *DescribeHybridCloudResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudResourcesResponse{}
	_body, _err := client.DescribeHybridCloudResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取SDK信息
//
// @param request - DescribeHybridCloudSdkServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudSdkServersResponse
func (client *Client) DescribeHybridCloudSdkServersWithOptions(request *DescribeHybridCloudSdkServersRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudSdkServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudSdkServers"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudSdkServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取SDK信息
//
// @param request - DescribeHybridCloudSdkServersRequest
//
// @return DescribeHybridCloudSdkServersResponse
func (client *Client) DescribeHybridCloudSdkServers(request *DescribeHybridCloudSdkServersRequest) (_result *DescribeHybridCloudSdkServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudSdkServersResponse{}
	_body, _err := client.DescribeHybridCloudSdkServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the regions that the hybrid cloud mode supports, such as the Internet service providers (ISPs), continents, and cities.
//
// @param request - DescribeHybridCloudServerRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudServerRegionsResponse
func (client *Client) DescribeHybridCloudServerRegionsWithOptions(request *DescribeHybridCloudServerRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudServerRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionType) {
		query["RegionType"] = request.RegionType
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudServerRegions"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudServerRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the regions that the hybrid cloud mode supports, such as the Internet service providers (ISPs), continents, and cities.
//
// @param request - DescribeHybridCloudServerRegionsRequest
//
// @return DescribeHybridCloudServerRegionsResponse
func (client *Client) DescribeHybridCloudServerRegions(request *DescribeHybridCloudServerRegionsRequest) (_result *DescribeHybridCloudServerRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudServerRegionsResponse{}
	_body, _err := client.DescribeHybridCloudServerRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询接入区域
//
// @param request - DescribeHybridCloudSupportRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudSupportRegionsResponse
func (client *Client) DescribeHybridCloudSupportRegionsWithOptions(request *DescribeHybridCloudSupportRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudSupportRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudSupportRegions"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudSupportRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询接入区域
//
// @param request - DescribeHybridCloudSupportRegionsRequest
//
// @return DescribeHybridCloudSupportRegionsResponse
func (client *Client) DescribeHybridCloudSupportRegions(request *DescribeHybridCloudSupportRegionsRequest) (_result *DescribeHybridCloudSupportRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudSupportRegionsResponse{}
	_body, _err := client.DescribeHybridCloudSupportRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries servers that are not assigned to a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudUnassignedMachinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudUnassignedMachinesResponse
func (client *Client) DescribeHybridCloudUnassignedMachinesWithOptions(request *DescribeHybridCloudUnassignedMachinesRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudUnassignedMachinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudUnassignedMachines"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudUnassignedMachinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries servers that are not assigned to a hybrid cloud cluster.
//
// @param request - DescribeHybridCloudUnassignedMachinesRequest
//
// @return DescribeHybridCloudUnassignedMachinesResponse
func (client *Client) DescribeHybridCloudUnassignedMachines(request *DescribeHybridCloudUnassignedMachinesRequest) (_result *DescribeHybridCloudUnassignedMachinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudUnassignedMachinesResponse{}
	_body, _err := client.DescribeHybridCloudUnassignedMachinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ports that are not supported by the hybrid cloud mode.
//
// @param request - DescribeHybridCloudUnsupportPortsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudUnsupportPortsResponse
func (client *Client) DescribeHybridCloudUnsupportPortsWithOptions(request *DescribeHybridCloudUnsupportPortsRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudUnsupportPortsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudUnsupportPorts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudUnsupportPortsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports that are not supported by the hybrid cloud mode.
//
// @param request - DescribeHybridCloudUnsupportPortsRequest
//
// @return DescribeHybridCloudUnsupportPortsResponse
func (client *Client) DescribeHybridCloudUnsupportPorts(request *DescribeHybridCloudUnsupportPortsRequest) (_result *DescribeHybridCloudUnsupportPortsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudUnsupportPortsResponse{}
	_body, _err := client.DescribeHybridCloudUnsupportPortsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the HTTP and HTTPS ports that you can use when you add a domain name to Web Application Firewall (WAF) in hybrid cloud mode.
//
// @param request - DescribeHybridCloudUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHybridCloudUserResponse
func (client *Client) DescribeHybridCloudUserWithOptions(request *DescribeHybridCloudUserRequest, runtime *dara.RuntimeOptions) (_result *DescribeHybridCloudUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHybridCloudUser"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHybridCloudUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the HTTP and HTTPS ports that you can use when you add a domain name to Web Application Firewall (WAF) in hybrid cloud mode.
//
// @param request - DescribeHybridCloudUserRequest
//
// @return DescribeHybridCloudUserResponse
func (client *Client) DescribeHybridCloudUser(request *DescribeHybridCloudUserRequest) (_result *DescribeHybridCloudUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHybridCloudUserResponse{}
	_body, _err := client.DescribeHybridCloudUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Web Application Firewall (WAF) instance within the current Alibaba Cloud account.
//
// @param request - DescribeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstance"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Web Application Firewall (WAF) instance within the current Alibaba Cloud account.
//
// @param request - DescribeInstanceRequest
//
// @return DescribeInstanceResponse
func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取支持的海外IP区域封禁支持的国际及地域。
//
// @param request - DescribeIpAbroadCountryInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpAbroadCountryInfosResponse
func (client *Client) DescribeIpAbroadCountryInfosWithOptions(request *DescribeIpAbroadCountryInfosRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpAbroadCountryInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbroadRegion) {
		query["AbroadRegion"] = request.AbroadRegion
	}

	if !dara.IsNil(request.Country) {
		query["Country"] = request.Country
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
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

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpAbroadCountryInfos"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpAbroadCountryInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取支持的海外IP区域封禁支持的国际及地域。
//
// @param request - DescribeIpAbroadCountryInfosRequest
//
// @return DescribeIpAbroadCountryInfosResponse
func (client *Client) DescribeIpAbroadCountryInfos(request *DescribeIpAbroadCountryInfosRequest) (_result *DescribeIpAbroadCountryInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpAbroadCountryInfosResponse{}
	_body, _err := client.DescribeIpAbroadCountryInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries IP addresses in an IP address blacklist for major event protection by page.
//
// @param request - DescribeMajorProtectionBlackIpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMajorProtectionBlackIpsResponse
func (client *Client) DescribeMajorProtectionBlackIpsWithOptions(request *DescribeMajorProtectionBlackIpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IpLike) {
		query["IpLike"] = request.IpLike
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMajorProtectionBlackIps"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries IP addresses in an IP address blacklist for major event protection by page.
//
// @param request - DescribeMajorProtectionBlackIpsRequest
//
// @return DescribeMajorProtectionBlackIpsResponse
func (client *Client) DescribeMajorProtectionBlackIps(request *DescribeMajorProtectionBlackIpsRequest) (_result *DescribeMajorProtectionBlackIpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMajorProtectionBlackIpsResponse{}
	_body, _err := client.DescribeMajorProtectionBlackIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about members.
//
// @param request - DescribeMemberAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMemberAccountsResponse
func (client *Client) DescribeMemberAccountsWithOptions(request *DescribeMemberAccountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMemberAccountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountStatus) {
		query["AccountStatus"] = request.AccountStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMemberAccounts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMemberAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about members.
//
// @param request - DescribeMemberAccountsRequest
//
// @return DescribeMemberAccountsResponse
func (client *Client) DescribeMemberAccounts(request *DescribeMemberAccountsRequest) (_result *DescribeMemberAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMemberAccountsResponse{}
	_body, _err := client.DescribeMemberAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves time-series data for all network traffic, including both malicious and legitimate requests.
//
// @param tmpReq - DescribeNetworkFlowTimeSeriesMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkFlowTimeSeriesMetricResponse
func (client *Client) DescribeNetworkFlowTimeSeriesMetricWithOptions(tmpReq *DescribeNetworkFlowTimeSeriesMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkFlowTimeSeriesMetricResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeNetworkFlowTimeSeriesMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkFlowTimeSeriesMetric"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkFlowTimeSeriesMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves time-series data for all network traffic, including both malicious and legitimate requests.
//
// @param request - DescribeNetworkFlowTimeSeriesMetricRequest
//
// @return DescribeNetworkFlowTimeSeriesMetricResponse
func (client *Client) DescribeNetworkFlowTimeSeriesMetric(request *DescribeNetworkFlowTimeSeriesMetricRequest) (_result *DescribeNetworkFlowTimeSeriesMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkFlowTimeSeriesMetricResponse{}
	_body, _err := client.DescribeNetworkFlowTimeSeriesMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves top aggregated traffic statistics, sorted by various dimensions, including malicious and legitimate requests.
//
// @param tmpReq - DescribeNetworkFlowTopNMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkFlowTopNMetricResponse
func (client *Client) DescribeNetworkFlowTopNMetricWithOptions(tmpReq *DescribeNetworkFlowTopNMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkFlowTopNMetricResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeNetworkFlowTopNMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkFlowTopNMetric"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkFlowTopNMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves top aggregated traffic statistics, sorted by various dimensions, including malicious and legitimate requests.
//
// @param request - DescribeNetworkFlowTopNMetricRequest
//
// @return DescribeNetworkFlowTopNMetricResponse
func (client *Client) DescribeNetworkFlowTopNMetric(request *DescribeNetworkFlowTopNMetricRequest) (_result *DescribeNetworkFlowTopNMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkFlowTopNMetricResponse{}
	_body, _err := client.DescribeNetworkFlowTopNMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the protection status of Web Application Firewall (WAF).
//
// @param request - DescribePauseProtectionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePauseProtectionStatusResponse
func (client *Client) DescribePauseProtectionStatusWithOptions(request *DescribePauseProtectionStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribePauseProtectionStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePauseProtectionStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePauseProtectionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the protection status of Web Application Firewall (WAF).
//
// @param request - DescribePauseProtectionStatusRequest
//
// @return DescribePauseProtectionStatusResponse
func (client *Client) DescribePauseProtectionStatus(request *DescribePauseProtectionStatusRequest) (_result *DescribePauseProtectionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePauseProtectionStatusResponse{}
	_body, _err := client.DescribePauseProtectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the queries per second (QPS) statistics of a WAF instance.
//
// @param request - DescribePeakTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePeakTrendResponse
func (client *Client) DescribePeakTrendWithOptions(request *DescribePeakTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribePeakTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePeakTrend"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the queries per second (QPS) statistics of a WAF instance.
//
// @param request - DescribePeakTrendRequest
//
// @return DescribePeakTrendResponse
func (client *Client) DescribePeakTrend(request *DescribePeakTrendRequest) (_result *DescribePeakTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePeakTrendResponse{}
	_body, _err := client.DescribePeakTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cloud service instances to be added to Web Application Firewall (WAF) in transparent proxy mode.
//
// @param request - DescribeProductInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductInstancesResponse
func (client *Client) DescribeProductInstancesWithOptions(request *DescribeProductInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeProductInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceAccessStatus) {
		query["ResourceInstanceAccessStatus"] = request.ResourceInstanceAccessStatus
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceInstanceIp) {
		query["ResourceInstanceIp"] = request.ResourceInstanceIp
	}

	if !dara.IsNil(request.ResourceInstanceName) {
		query["ResourceInstanceName"] = request.ResourceInstanceName
	}

	if !dara.IsNil(request.ResourceIp) {
		query["ResourceIp"] = request.ResourceIp
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProductInstances"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProductInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cloud service instances to be added to Web Application Firewall (WAF) in transparent proxy mode.
//
// @param request - DescribeProductInstancesRequest
//
// @return DescribeProductInstancesResponse
func (client *Client) DescribeProductInstances(request *DescribeProductInstancesRequest) (_result *DescribeProductInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeProductInstancesResponse{}
	_body, _err := client.DescribeProductInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of domain names that are added to Web Application Firewall (WAF) and penalized for failing to obtain an Internet Content Provider (ICP) filing.
//
// @param request - DescribePunishedDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePunishedDomainsResponse
func (client *Client) DescribePunishedDomainsWithOptions(request *DescribePunishedDomainsRequest, runtime *dara.RuntimeOptions) (_result *DescribePunishedDomainsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domains) {
		query["Domains"] = request.Domains
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PunishType) {
		query["PunishType"] = request.PunishType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePunishedDomains"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePunishedDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names that are added to Web Application Firewall (WAF) and penalized for failing to obtain an Internet Content Provider (ICP) filing.
//
// @param request - DescribePunishedDomainsRequest
//
// @return DescribePunishedDomainsResponse
func (client *Client) DescribePunishedDomains(request *DescribePunishedDomainsRequest) (_result *DescribePunishedDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePunishedDomainsResponse{}
	_body, _err := client.DescribePunishedDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the certificates that are used in cloud service instances. The certificates returned include the certificates within the delegated administrator account and the certificates within members to which specific instances belong. For example, the delegated administrator account has certificate 1, instance lb-xx-1 belongs to member B, and member B has certificate 2. If you specify instance lb-xx-1 in the request, certificate 1 and certificate 2 are returned.
//
// @param request - DescribeResourceInstanceCertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceInstanceCertsResponse
func (client *Client) DescribeResourceInstanceCertsWithOptions(request *DescribeResourceInstanceCertsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceInstanceCertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceInstanceCerts"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceInstanceCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the certificates that are used in cloud service instances. The certificates returned include the certificates within the delegated administrator account and the certificates within members to which specific instances belong. For example, the delegated administrator account has certificate 1, instance lb-xx-1 belongs to member B, and member B has certificate 2. If you specify instance lb-xx-1 in the request, certificate 1 and certificate 2 are returned.
//
// @param request - DescribeResourceInstanceCertsRequest
//
// @return DescribeResourceInstanceCertsResponse
func (client *Client) DescribeResourceInstanceCerts(request *DescribeResourceInstanceCertsRequest) (_result *DescribeResourceInstanceCertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceInstanceCertsResponse{}
	_body, _err := client.DescribeResourceInstanceCertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the log collection feature is enabled for a protected object.
//
// @param request - DescribeResourceLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceLogStatusResponse
func (client *Client) DescribeResourceLogStatusWithOptions(request *DescribeResourceLogStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceLogStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceLogStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the log collection feature is enabled for a protected object.
//
// @param request - DescribeResourceLogStatusRequest
//
// @return DescribeResourceLogStatusResponse
func (client *Client) DescribeResourceLogStatus(request *DescribeResourceLogStatusRequest) (_result *DescribeResourceLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceLogStatusResponse{}
	_body, _err := client.DescribeResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ports of a cloud service instance that are added to Web Application Firewall (WAF).
//
// @param request - DescribeResourcePortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcePortResponse
func (client *Client) DescribeResourcePortWithOptions(request *DescribeResourcePortRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourcePortResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourcePort"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports of a cloud service instance that are added to Web Application Firewall (WAF).
//
// @param request - DescribeResourcePortRequest
//
// @return DescribeResourcePortResponse
func (client *Client) DescribeResourcePort(request *DescribeResourcePortRequest) (_result *DescribeResourcePortResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourcePortResponse{}
	_body, _err := client.DescribeResourcePortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region IDs of the resources that are added to Web Application Firewall (WAF) by using the SDK integration mode. The resources refer to Application Load Balancer (ALB) and Microservices Engine (MSE) instances.
//
// @param request - DescribeResourceRegionIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceRegionIdResponse
func (client *Client) DescribeResourceRegionIdWithOptions(request *DescribeResourceRegionIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceRegionIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceRegionId"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceRegionIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region IDs of the resources that are added to Web Application Firewall (WAF) by using the SDK integration mode. The resources refer to Application Load Balancer (ALB) and Microservices Engine (MSE) instances.
//
// @param request - DescribeResourceRegionIdRequest
//
// @return DescribeResourceRegionIdResponse
func (client *Client) DescribeResourceRegionId(request *DescribeResourceRegionIdRequest) (_result *DescribeResourceRegionIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceRegionIdResponse{}
	_body, _err := client.DescribeResourceRegionIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the region IDs of the Classic Load Balancer (CLB) and Elastic Compute Service (ECS) instances that are added to Web Application Firewall (WAF) in cloud native mode.
//
// @param request - DescribeResourceSupportRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceSupportRegionsResponse
func (client *Client) DescribeResourceSupportRegionsWithOptions(request *DescribeResourceSupportRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeResourceSupportRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceSupportRegions"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceSupportRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the region IDs of the Classic Load Balancer (CLB) and Elastic Compute Service (ECS) instances that are added to Web Application Firewall (WAF) in cloud native mode.
//
// @param request - DescribeResourceSupportRegionsRequest
//
// @return DescribeResourceSupportRegionsResponse
func (client *Client) DescribeResourceSupportRegions(request *DescribeResourceSupportRegionsRequest) (_result *DescribeResourceSupportRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResourceSupportRegionsResponse{}
	_body, _err := client.DescribeResourceSupportRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of error codes that are returned to clients or Web Application Firewall (WAF). The error codes include 302, 405, 444, 499, and 5XX.
//
// @param request - DescribeResponseCodeTrendGraphRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResponseCodeTrendGraphResponse
func (client *Client) DescribeResponseCodeTrendGraphWithOptions(request *DescribeResponseCodeTrendGraphRequest, runtime *dara.RuntimeOptions) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResponseCodeTrendGraph"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of error codes that are returned to clients or Web Application Firewall (WAF). The error codes include 302, 405, 444, 499, and 5XX.
//
// @param request - DescribeResponseCodeTrendGraphRequest
//
// @return DescribeResponseCodeTrendGraphResponse
func (client *Client) DescribeResponseCodeTrendGraph(request *DescribeResponseCodeTrendGraphRequest) (_result *DescribeResponseCodeTrendGraphResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeResponseCodeTrendGraphResponse{}
	_body, _err := client.DescribeResponseCodeTrendGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regular expression rule groups by page.
//
// @param request - DescribeRuleGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleGroupsResponse
func (client *Client) DescribeRuleGroupsWithOptions(request *DescribeRuleGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SearchType) {
		query["SearchType"] = request.SearchType
	}

	if !dara.IsNil(request.SearchValue) {
		query["SearchValue"] = request.SearchValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleGroups"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regular expression rule groups by page.
//
// @param request - DescribeRuleGroupsRequest
//
// @return DescribeRuleGroupsResponse
func (client *Client) DescribeRuleGroups(request *DescribeRuleGroupsRequest) (_result *DescribeRuleGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleGroupsResponse{}
	_body, _err := client.DescribeRuleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 IP addresses from which attacks are initiated.
//
// @param request - DescribeRuleHitsTopClientIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopClientIpResponse
func (client *Client) DescribeRuleHitsTopClientIpWithOptions(request *DescribeRuleHitsTopClientIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopClientIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 IP addresses from which attacks are initiated.
//
// @param request - DescribeRuleHitsTopClientIpRequest
//
// @return DescribeRuleHitsTopClientIpResponse
func (client *Client) DescribeRuleHitsTopClientIp(request *DescribeRuleHitsTopClientIpRequest) (_result *DescribeRuleHitsTopClientIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopClientIpResponse{}
	_body, _err := client.DescribeRuleHitsTopClientIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 protected objects that trigger protection rules.
//
// @param request - DescribeRuleHitsTopResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopResourceResponse
func (client *Client) DescribeRuleHitsTopResourceWithOptions(request *DescribeRuleHitsTopResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 protected objects that trigger protection rules.
//
// @param request - DescribeRuleHitsTopResourceRequest
//
// @return DescribeRuleHitsTopResourceResponse
func (client *Client) DescribeRuleHitsTopResource(request *DescribeRuleHitsTopResourceRequest) (_result *DescribeRuleHitsTopResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopResourceResponse{}
	_body, _err := client.DescribeRuleHitsTopResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IDs of the top 10 protection rules that are matched by requests.
//
// @param request - DescribeRuleHitsTopRuleIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopRuleIdResponse
func (client *Client) DescribeRuleHitsTopRuleIdWithOptions(request *DescribeRuleHitsTopRuleIdRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsGroupResource) {
		query["IsGroupResource"] = request.IsGroupResource
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopRuleId"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of the top 10 protection rules that are matched by requests.
//
// @param request - DescribeRuleHitsTopRuleIdRequest
//
// @return DescribeRuleHitsTopRuleIdResponse
func (client *Client) DescribeRuleHitsTopRuleId(request *DescribeRuleHitsTopRuleIdRequest) (_result *DescribeRuleHitsTopRuleIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopRuleIdResponse{}
	_body, _err := client.DescribeRuleHitsTopRuleIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 protection modules that are matched.
//
// @param request - DescribeRuleHitsTopTuleTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopTuleTypeResponse
func (client *Client) DescribeRuleHitsTopTuleTypeWithOptions(request *DescribeRuleHitsTopTuleTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopTuleType"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 protection modules that are matched.
//
// @param request - DescribeRuleHitsTopTuleTypeRequest
//
// @return DescribeRuleHitsTopTuleTypeResponse
func (client *Client) DescribeRuleHitsTopTuleType(request *DescribeRuleHitsTopTuleTypeRequest) (_result *DescribeRuleHitsTopTuleTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopTuleTypeResponse{}
	_body, _err := client.DescribeRuleHitsTopTuleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 user agents that are used to initiate attacks.
//
// @param request - DescribeRuleHitsTopUaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopUaResponse
func (client *Client) DescribeRuleHitsTopUaWithOptions(request *DescribeRuleHitsTopUaRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopUa"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 user agents that are used to initiate attacks.
//
// @param request - DescribeRuleHitsTopUaRequest
//
// @return DescribeRuleHitsTopUaResponse
func (client *Client) DescribeRuleHitsTopUa(request *DescribeRuleHitsTopUaRequest) (_result *DescribeRuleHitsTopUaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUaResponse{}
	_body, _err := client.DescribeRuleHitsTopUaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 URLs that trigger protection rules.
//
// @param request - DescribeRuleHitsTopUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleHitsTopUrlResponse
func (client *Client) DescribeRuleHitsTopUrlWithOptions(request *DescribeRuleHitsTopUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleHitsTopUrl"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 URLs that trigger protection rules.
//
// @param request - DescribeRuleHitsTopUrlRequest
//
// @return DescribeRuleHitsTopUrlResponse
func (client *Client) DescribeRuleHitsTopUrl(request *DescribeRuleHitsTopUrlRequest) (_result *DescribeRuleHitsTopUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleHitsTopUrlResponse{}
	_body, _err := client.DescribeRuleHitsTopUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of attack traffic. Each log records the details of a request that matches protection rules.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param tmpReq - DescribeSecurityEventLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventLogsResponse
func (client *Client) DescribeSecurityEventLogsWithOptions(tmpReq *DescribeSecurityEventLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventLogsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSecurityEventLogsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventLogs"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of attack traffic. Each log records the details of a request that matches protection rules.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param request - DescribeSecurityEventLogsRequest
//
// @return DescribeSecurityEventLogsResponse
func (client *Client) DescribeSecurityEventLogs(request *DescribeSecurityEventLogsRequest) (_result *DescribeSecurityEventLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityEventLogsResponse{}
	_body, _err := client.DescribeSecurityEventLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the time series data of attack traffic. Attack requests refer to requests that match protection rules and are identified as risky.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param tmpReq - DescribeSecurityEventTimeSeriesMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventTimeSeriesMetricResponse
func (client *Client) DescribeSecurityEventTimeSeriesMetricWithOptions(tmpReq *DescribeSecurityEventTimeSeriesMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventTimeSeriesMetricResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSecurityEventTimeSeriesMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventTimeSeriesMetric"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventTimeSeriesMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time series data of attack traffic. Attack requests refer to requests that match protection rules and are identified as risky.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param request - DescribeSecurityEventTimeSeriesMetricRequest
//
// @return DescribeSecurityEventTimeSeriesMetricResponse
func (client *Client) DescribeSecurityEventTimeSeriesMetric(request *DescribeSecurityEventTimeSeriesMetricRequest) (_result *DescribeSecurityEventTimeSeriesMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityEventTimeSeriesMetricResponse{}
	_body, _err := client.DescribeSecurityEventTimeSeriesMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries top N data entries of attack traffic. The system performs statistical aggregation on attack traffic from specific dimensions and returns top N data entries.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param tmpReq - DescribeSecurityEventTopNMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventTopNMetricResponse
func (client *Client) DescribeSecurityEventTopNMetricWithOptions(tmpReq *DescribeSecurityEventTopNMetricRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventTopNMetricResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSecurityEventTopNMetricShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Metric) {
		query["Metric"] = request.Metric
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventTopNMetric"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventTopNMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries top N data entries of attack traffic. The system performs statistical aggregation on attack traffic from specific dimensions and returns top N data entries.
//
// Description:
//
// Attack traffic refers to the traffic of requests that match protection rules and are identified as risky. The following types of requests are excluded:
//
//   - Requests that match the protection rules of the whitelist module.
//
//   - Requests that match the protection rules of the bot management module. The actions of the protection rules are set to Add Tag.
//
//   - Requests that match protection rules with actions set to Dynamic Token-based Authentication, Slider CAPTCHA, Strict Slider CAPTCHA Verification, and JavaScript Validation, pass the verifications specified by the actions, and are allowed.
//
// @param request - DescribeSecurityEventTopNMetricRequest
//
// @return DescribeSecurityEventTopNMetricResponse
func (client *Client) DescribeSecurityEventTopNMetric(request *DescribeSecurityEventTopNMetricRequest) (_result *DescribeSecurityEventTopNMetricResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityEventTopNMetricResponse{}
	_body, _err := client.DescribeSecurityEventTopNMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the personal information-related APIs and domain names.
//
// @param request - DescribeSensitiveApiStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveApiStatisticResponse
func (client *Client) DescribeSensitiveApiStatisticWithOptions(request *DescribeSensitiveApiStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveApiStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
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
		Action:      dara.String("DescribeSensitiveApiStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveApiStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the personal information-related APIs and domain names.
//
// @param request - DescribeSensitiveApiStatisticRequest
//
// @return DescribeSensitiveApiStatisticResponse
func (client *Client) DescribeSensitiveApiStatistic(request *DescribeSensitiveApiStatisticRequest) (_result *DescribeSensitiveApiStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveApiStatisticResponse{}
	_body, _err := client.DescribeSensitiveApiStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the compliance check results of API security.
//
// @param request - DescribeSensitiveDetectionResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveDetectionResultResponse
func (client *Client) DescribeSensitiveDetectionResultWithOptions(request *DescribeSensitiveDetectionResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveDetectionResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveDetectionResult"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveDetectionResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the compliance check results of API security.
//
// @param request - DescribeSensitiveDetectionResultRequest
//
// @return DescribeSensitiveDetectionResultResponse
func (client *Client) DescribeSensitiveDetectionResult(request *DescribeSensitiveDetectionResultRequest) (_result *DescribeSensitiveDetectionResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveDetectionResultResponse{}
	_body, _err := client.DescribeSensitiveDetectionResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic distribution of personal information records involved in cross-border data transfer.
//
// @param request - DescribeSensitiveOutboundDistributionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveOutboundDistributionResponse
func (client *Client) DescribeSensitiveOutboundDistributionWithOptions(request *DescribeSensitiveOutboundDistributionRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveOutboundDistributionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveOutboundDistribution"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveOutboundDistributionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic distribution of personal information records involved in cross-border data transfer.
//
// @param request - DescribeSensitiveOutboundDistributionRequest
//
// @return DescribeSensitiveOutboundDistributionResponse
func (client *Client) DescribeSensitiveOutboundDistribution(request *DescribeSensitiveOutboundDistributionRequest) (_result *DescribeSensitiveOutboundDistributionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveOutboundDistributionResponse{}
	_body, _err := client.DescribeSensitiveOutboundDistributionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data types of personal information involved in cross-border data transfer.
//
// @param request - DescribeSensitiveOutboundStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveOutboundStatisticResponse
func (client *Client) DescribeSensitiveOutboundStatisticWithOptions(request *DescribeSensitiveOutboundStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveOutboundStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DetectionResult) {
		query["DetectionResult"] = request.DetectionResult
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderKey) {
		query["OrderKey"] = request.OrderKey
	}

	if !dara.IsNil(request.OrderWay) {
		query["OrderWay"] = request.OrderWay
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SensitiveCode) {
		query["SensitiveCode"] = request.SensitiveCode
	}

	if !dara.IsNil(request.SensitiveLevel) {
		query["SensitiveLevel"] = request.SensitiveLevel
	}

	if !dara.IsNil(request.SensitiveType) {
		query["SensitiveType"] = request.SensitiveType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveOutboundStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveOutboundStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data types of personal information involved in cross-border data transfer.
//
// @param request - DescribeSensitiveOutboundStatisticRequest
//
// @return DescribeSensitiveOutboundStatisticResponse
func (client *Client) DescribeSensitiveOutboundStatistic(request *DescribeSensitiveOutboundStatisticRequest) (_result *DescribeSensitiveOutboundStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveOutboundStatisticResponse{}
	_body, _err := client.DescribeSensitiveOutboundStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trends of cross-border data transfer of personal information.
//
// @param request - DescribeSensitiveOutboundTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveOutboundTrendResponse
func (client *Client) DescribeSensitiveOutboundTrendWithOptions(request *DescribeSensitiveOutboundTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveOutboundTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveOutboundTrend"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveOutboundTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trends of cross-border data transfer of personal information.
//
// @param request - DescribeSensitiveOutboundTrendRequest
//
// @return DescribeSensitiveOutboundTrendResponse
func (client *Client) DescribeSensitiveOutboundTrend(request *DescribeSensitiveOutboundTrendRequest) (_result *DescribeSensitiveOutboundTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveOutboundTrendResponse{}
	_body, _err := client.DescribeSensitiveOutboundTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access logs of sensitive data.
//
// @param request - DescribeSensitiveRequestLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveRequestLogResponse
func (client *Client) DescribeSensitiveRequestLogWithOptions(request *DescribeSensitiveRequestLogRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveRequestLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiFormat) {
		query["ApiFormat"] = request.ApiFormat
	}

	if !dara.IsNil(request.ClientIP) {
		query["ClientIP"] = request.ClientIP
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MatchedHost) {
		query["MatchedHost"] = request.MatchedHost
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SensitiveCode) {
		query["SensitiveCode"] = request.SensitiveCode
	}

	if !dara.IsNil(request.SensitiveData) {
		query["SensitiveData"] = request.SensitiveData
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveRequestLog"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveRequestLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access logs of sensitive data.
//
// @param request - DescribeSensitiveRequestLogRequest
//
// @return DescribeSensitiveRequestLogResponse
func (client *Client) DescribeSensitiveRequestLog(request *DescribeSensitiveRequestLogRequest) (_result *DescribeSensitiveRequestLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveRequestLogResponse{}
	_body, _err := client.DescribeSensitiveRequestLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tracing results of sensitive data.
//
// @param request - DescribeSensitiveRequestsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveRequestsResponse
func (client *Client) DescribeSensitiveRequestsWithOptions(request *DescribeSensitiveRequestsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveRequestsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SensitiveCode) {
		query["SensitiveCode"] = request.SensitiveCode
	}

	if !dara.IsNil(request.SensitiveData) {
		query["SensitiveData"] = request.SensitiveData
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveRequests"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveRequestsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tracing results of sensitive data.
//
// @param request - DescribeSensitiveRequestsRequest
//
// @return DescribeSensitiveRequestsResponse
func (client *Client) DescribeSensitiveRequests(request *DescribeSensitiveRequestsRequest) (_result *DescribeSensitiveRequestsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveRequestsResponse{}
	_body, _err := client.DescribeSensitiveRequestsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the sensitive data statistics of the tracing and auditing feature.
//
// @param request - DescribeSensitiveStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSensitiveStatisticResponse
func (client *Client) DescribeSensitiveStatisticWithOptions(request *DescribeSensitiveStatisticRequest, runtime *dara.RuntimeOptions) (_result *DescribeSensitiveStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisticType) {
		query["StatisticType"] = request.StatisticType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSensitiveStatistic"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSensitiveStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sensitive data statistics of the tracing and auditing feature.
//
// @param request - DescribeSensitiveStatisticRequest
//
// @return DescribeSensitiveStatisticResponse
func (client *Client) DescribeSensitiveStatistic(request *DescribeSensitiveStatisticRequest) (_result *DescribeSensitiveStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSensitiveStatisticResponse{}
	_body, _err := client.DescribeSensitiveStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Web Application Firewall (WAF) is authorized to access Logstores.
//
// @param request - DescribeSlsAuthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsAuthStatusResponse
func (client *Client) DescribeSlsAuthStatusWithOptions(request *DescribeSlsAuthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsAuthStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsAuthStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Web Application Firewall (WAF) is authorized to access Logstores.
//
// @param request - DescribeSlsAuthStatusRequest
//
// @return DescribeSlsAuthStatusResponse
func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (_result *DescribeSlsAuthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsAuthStatusResponse{}
	_body, _err := client.DescribeSlsAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a Logstore, such as the total capacity, storage duration, and used capacity.
//
// @param request - DescribeSlsLogStoreRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsLogStoreResponse
func (client *Client) DescribeSlsLogStoreWithOptions(request *DescribeSlsLogStoreRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsLogStoreResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsLogStore"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a Logstore, such as the total capacity, storage duration, and used capacity.
//
// @param request - DescribeSlsLogStoreRequest
//
// @return DescribeSlsLogStoreResponse
func (client *Client) DescribeSlsLogStore(request *DescribeSlsLogStoreRequest) (_result *DescribeSlsLogStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsLogStoreResponse{}
	_body, _err := client.DescribeSlsLogStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a Simple Log Service Logstore.
//
// @param request - DescribeSlsLogStoreStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlsLogStoreStatusResponse
func (client *Client) DescribeSlsLogStoreStatusWithOptions(request *DescribeSlsLogStoreStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeSlsLogStoreStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSlsLogStoreStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSlsLogStoreStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a Simple Log Service Logstore.
//
// @param request - DescribeSlsLogStoreStatusRequest
//
// @return DescribeSlsLogStoreStatusResponse
func (client *Client) DescribeSlsLogStoreStatus(request *DescribeSlsLogStoreStatusRequest) (_result *DescribeSlsLogStoreStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSlsLogStoreStatusResponse{}
	_body, _err := client.DescribeSlsLogStoreStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of protected resources for which a protection template takes effect.
//
// @param request - DescribeTemplateResourceCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateResourceCountResponse
func (client *Client) DescribeTemplateResourceCountWithOptions(request *DescribeTemplateResourceCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateResourceCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateResourceCount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateResourceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of protected resources for which a protection template takes effect.
//
// @param request - DescribeTemplateResourceCountRequest
//
// @return DescribeTemplateResourceCountResponse
func (client *Client) DescribeTemplateResourceCount(request *DescribeTemplateResourceCountRequest) (_result *DescribeTemplateResourceCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateResourceCountResponse{}
	_body, _err := client.DescribeTemplateResourceCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated to a protection rule template.
//
// @param request - DescribeTemplateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTemplateResourcesResponse
func (client *Client) DescribeTemplateResourcesWithOptions(request *DescribeTemplateResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeTemplateResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTemplateResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources that are associated to a protection rule template.
//
// @param request - DescribeTemplateResourcesRequest
//
// @return DescribeTemplateResourcesResponse
func (client *Client) DescribeTemplateResources(request *DescribeTemplateResourcesRequest) (_result *DescribeTemplateResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTemplateResourcesResponse{}
	_body, _err := client.DescribeTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trends of API security risks.
//
// @param request - DescribeUserAbnormalTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAbnormalTrendResponse
func (client *Client) DescribeUserAbnormalTrendWithOptions(request *DescribeUserAbnormalTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAbnormalTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAbnormalTrend"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAbnormalTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trends of API security risks.
//
// @param request - DescribeUserAbnormalTrendRequest
//
// @return DescribeUserAbnormalTrendResponse
func (client *Client) DescribeUserAbnormalTrend(request *DescribeUserAbnormalTrendRequest) (_result *DescribeUserAbnormalTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAbnormalTrendResponse{}
	_body, _err := client.DescribeUserAbnormalTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types and statistics of risks in the API security module.
//
// @param request - DescribeUserAbnormalTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAbnormalTypeResponse
func (client *Client) DescribeUserAbnormalTypeWithOptions(request *DescribeUserAbnormalTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAbnormalTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserStatusList) {
		query["UserStatusList"] = request.UserStatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAbnormalType"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAbnormalTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types and statistics of risks in the API security module.
//
// @param request - DescribeUserAbnormalTypeRequest
//
// @return DescribeUserAbnormalTypeResponse
func (client *Client) DescribeUserAbnormalType(request *DescribeUserAbnormalTypeRequest) (_result *DescribeUserAbnormalTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAbnormalTypeResponse{}
	_body, _err := client.DescribeUserAbnormalTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic statistics of an API.
//
// @param request - DescribeUserApiRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserApiRequestResponse
func (client *Client) DescribeUserApiRequestWithOptions(request *DescribeUserApiRequestRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserApiRequestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiFormat) {
		query["ApiFormat"] = request.ApiFormat
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserApiRequest"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserApiRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic statistics of an API.
//
// @param request - DescribeUserApiRequestRequest
//
// @return DescribeUserApiRequestResponse
func (client *Client) DescribeUserApiRequest(request *DescribeUserApiRequestRequest) (_result *DescribeUserApiRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserApiRequestResponse{}
	_body, _err := client.DescribeUserApiRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the user asset statistics in the API security module.
//
// @param request - DescribeUserAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAssetResponse
func (client *Client) DescribeUserAssetWithOptions(request *DescribeUserAssetRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserAssetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.Days) {
		query["Days"] = request.Days
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserAsset"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the user asset statistics in the API security module.
//
// @param request - DescribeUserAssetRequest
//
// @return DescribeUserAssetResponse
func (client *Client) DescribeUserAsset(request *DescribeUserAssetRequest) (_result *DescribeUserAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserAssetResponse{}
	_body, _err := client.DescribeUserAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trends of attacks detected by the API security module.
//
// @param request - DescribeUserEventTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEventTrendResponse
func (client *Client) DescribeUserEventTrendWithOptions(request *DescribeUserEventTrendRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserEventTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEventTrend"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserEventTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trends of attacks detected by the API security module.
//
// @param request - DescribeUserEventTrendRequest
//
// @return DescribeUserEventTrendResponse
func (client *Client) DescribeUserEventTrend(request *DescribeUserEventTrendRequest) (_result *DescribeUserEventTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserEventTrendResponse{}
	_body, _err := client.DescribeUserEventTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types and statistics of security events in the API security module.
//
// @param request - DescribeUserEventTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserEventTypeResponse
func (client *Client) DescribeUserEventTypeWithOptions(request *DescribeUserEventTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserEventTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UserStatusList) {
		query["UserStatusList"] = request.UserStatusList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserEventType"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserEventTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types and statistics of security events in the API security module.
//
// @param request - DescribeUserEventTypeRequest
//
// @return DescribeUserEventTypeResponse
func (client *Client) DescribeUserEventType(request *DescribeUserEventTypeRequest) (_result *DescribeUserEventTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserEventTypeResponse{}
	_body, _err := client.DescribeUserEventTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户日志配置
//
// @param request - DescribeUserLogFieldConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserLogFieldConfigResponse
func (client *Client) DescribeUserLogFieldConfigWithOptions(request *DescribeUserLogFieldConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserLogFieldConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryType) {
		query["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserLogFieldConfig"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserLogFieldConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户日志配置
//
// @param request - DescribeUserLogFieldConfigRequest
//
// @return DescribeUserLogFieldConfigResponse
func (client *Client) DescribeUserLogFieldConfig(request *DescribeUserLogFieldConfigRequest) (_result *DescribeUserLogFieldConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserLogFieldConfigResponse{}
	_body, _err := client.DescribeUserLogFieldConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions for log storage.
//
// @param request - DescribeUserSlsLogRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserSlsLogRegionsResponse
func (client *Client) DescribeUserSlsLogRegionsWithOptions(request *DescribeUserSlsLogRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserSlsLogRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserSlsLogRegions"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserSlsLogRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions for log storage.
//
// @param request - DescribeUserSlsLogRegionsRequest
//
// @return DescribeUserSlsLogRegionsResponse
func (client *Client) DescribeUserSlsLogRegions(request *DescribeUserSlsLogRegionsRequest) (_result *DescribeUserSlsLogRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserSlsLogRegionsResponse{}
	_body, _err := client.DescribeUserSlsLogRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status, region ID, and status modification time of Web Application Firewall (WAF) logs.
//
// @param request - DescribeUserWafLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserWafLogStatusResponse
func (client *Client) DescribeUserWafLogStatusWithOptions(request *DescribeUserWafLogStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserWafLogStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserWafLogStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserWafLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status, region ID, and status modification time of Web Application Firewall (WAF) logs.
//
// @param request - DescribeUserWafLogStatusRequest
//
// @return DescribeUserWafLogStatusResponse
func (client *Client) DescribeUserWafLogStatus(request *DescribeUserWafLogStatusRequest) (_result *DescribeUserWafLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUserWafLogStatusResponse{}
	_body, _err := client.DescribeUserWafLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 IP addresses from which requests are sent.
//
// @param request - DescribeVisitTopIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVisitTopIpResponse
func (client *Client) DescribeVisitTopIpWithOptions(request *DescribeVisitTopIpRequest, runtime *dara.RuntimeOptions) (_result *DescribeVisitTopIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVisitTopIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 IP addresses from which requests are sent.
//
// @param request - DescribeVisitTopIpRequest
//
// @return DescribeVisitTopIpResponse
func (client *Client) DescribeVisitTopIp(request *DescribeVisitTopIpRequest) (_result *DescribeVisitTopIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVisitTopIpResponse{}
	_body, _err := client.DescribeVisitTopIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the top 10 user agents that are used to initiate requests.
//
// @param request - DescribeVisitUasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVisitUasResponse
func (client *Client) DescribeVisitUasWithOptions(request *DescribeVisitUasRequest, runtime *dara.RuntimeOptions) (_result *DescribeVisitUasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		query["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.StartTimestamp) {
		query["StartTimestamp"] = request.StartTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVisitUas"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the top 10 user agents that are used to initiate requests.
//
// @param request - DescribeVisitUasRequest
//
// @return DescribeVisitUasResponse
func (client *Client) DescribeVisitUas(request *DescribeVisitUasRequest) (_result *DescribeVisitUasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVisitUasResponse{}
	_body, _err := client.DescribeVisitUasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of a Web Application Firewall (WAF) instance.
//
// @param request - DescribeWafSourceIpSegmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWafSourceIpSegmentResponse
func (client *Client) DescribeWafSourceIpSegmentWithOptions(request *DescribeWafSourceIpSegmentRequest, runtime *dara.RuntimeOptions) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWafSourceIpSegment"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the back-to-origin CIDR blocks of a Web Application Firewall (WAF) instance.
//
// @param request - DescribeWafSourceIpSegmentRequest
//
// @return DescribeWafSourceIpSegmentResponse
func (client *Client) DescribeWafSourceIpSegment(request *DescribeWafSourceIpSegmentRequest) (_result *DescribeWafSourceIpSegmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWafSourceIpSegmentResponse{}
	_body, _err := client.DescribeWafSourceIpSegmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tag keys.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to a resource.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2021-10-01"),
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
// Queries the tags that are added to a resource.
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
// Queries the tag values of a tag key.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values of a tag key.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of multiple risks detected by the API security module at a time.
//
// @param request - ModifyApisecAbnormalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecAbnormalsResponse
func (client *Client) ModifyApisecAbnormalsWithOptions(request *ModifyApisecAbnormalsRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecAbnormalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbnormalIds) {
		query["AbnormalIds"] = request.AbnormalIds
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecAbnormals"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecAbnormalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of multiple risks detected by the API security module at a time.
//
// @param request - ModifyApisecAbnormalsRequest
//
// @return ModifyApisecAbnormalsResponse
func (client *Client) ModifyApisecAbnormals(request *ModifyApisecAbnormalsRequest) (_result *ModifyApisecAbnormalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecAbnormalsResponse{}
	_body, _err := client.ModifyApisecAbnormalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the annotations of APIs in the API security module.
//
// @param request - ModifyApisecApiResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecApiResourceResponse
func (client *Client) ModifyApisecApiResourceWithOptions(request *ModifyApisecApiResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecApiResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Follow) {
		query["Follow"] = request.Follow
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecApiResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecApiResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the annotations of APIs in the API security module.
//
// @param request - ModifyApisecApiResourceRequest
//
// @return ModifyApisecApiResourceResponse
func (client *Client) ModifyApisecApiResource(request *ModifyApisecApiResourceRequest) (_result *ModifyApisecApiResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecApiResourceResponse{}
	_body, _err := client.ModifyApisecApiResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of multiple security events detected by the API security module at a time.
//
// @param request - ModifyApisecEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecEventsResponse
func (client *Client) ModifyApisecEventsWithOptions(request *ModifyApisecEventsRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecEventsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EventIds) {
		query["EventIds"] = request.EventIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.UserStatus) {
		query["UserStatus"] = request.UserStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecEvents"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of multiple security events detected by the API security module at a time.
//
// @param request - ModifyApisecEventsRequest
//
// @return ModifyApisecEventsResponse
func (client *Client) ModifyApisecEvents(request *ModifyApisecEventsRequest) (_result *ModifyApisecEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecEventsResponse{}
	_body, _err := client.ModifyApisecEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of API security log subscription.
//
// @param request - ModifyApisecLogDeliveryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecLogDeliveryResponse
func (client *Client) ModifyApisecLogDeliveryWithOptions(request *ModifyApisecLogDeliveryRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecLogDeliveryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssertKey) {
		query["AssertKey"] = request.AssertKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogRegionId) {
		query["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecLogDelivery"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecLogDeliveryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of API security log subscription.
//
// @param request - ModifyApisecLogDeliveryRequest
//
// @return ModifyApisecLogDeliveryResponse
func (client *Client) ModifyApisecLogDelivery(request *ModifyApisecLogDeliveryRequest) (_result *ModifyApisecLogDeliveryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecLogDeliveryResponse{}
	_body, _err := client.ModifyApisecLogDeliveryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of API security log subscription.
//
// @param request - ModifyApisecLogDeliveryStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecLogDeliveryStatusResponse
func (client *Client) ModifyApisecLogDeliveryStatusWithOptions(request *ModifyApisecLogDeliveryStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecLogDeliveryStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssertKey) {
		query["AssertKey"] = request.AssertKey
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecLogDeliveryStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecLogDeliveryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of API security log subscription.
//
// @param request - ModifyApisecLogDeliveryStatusRequest
//
// @return ModifyApisecLogDeliveryStatusResponse
func (client *Client) ModifyApisecLogDeliveryStatus(request *ModifyApisecLogDeliveryStatusRequest) (_result *ModifyApisecLogDeliveryStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecLogDeliveryStatusResponse{}
	_body, _err := client.ModifyApisecLogDeliveryStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of features in the API security module for protected objects or protected object groups.
//
// @param request - ModifyApisecModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecModuleStatusResponse
func (client *Client) ModifyApisecModuleStatusWithOptions(request *ModifyApisecModuleStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecModuleStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReportStatus) {
		query["ReportStatus"] = request.ReportStatus
	}

	if !dara.IsNil(request.ResourceGroups) {
		query["ResourceGroups"] = request.ResourceGroups
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	if !dara.IsNil(request.TraceStatus) {
		query["TraceStatus"] = request.TraceStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecModuleStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecModuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of features in the API security module for protected objects or protected object groups.
//
// @param request - ModifyApisecModuleStatusRequest
//
// @return ModifyApisecModuleStatusResponse
func (client *Client) ModifyApisecModuleStatus(request *ModifyApisecModuleStatusRequest) (_result *ModifyApisecModuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecModuleStatusResponse{}
	_body, _err := client.ModifyApisecModuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of the API security module for protected objects or protected object groups.
//
// @param request - ModifyApisecStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApisecStatusResponse
func (client *Client) ModifyApisecStatusWithOptions(request *ModifyApisecStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyApisecStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApisecStatus) {
		query["ApisecStatus"] = request.ApisecStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroups) {
		query["ResourceGroups"] = request.ResourceGroups
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Resources) {
		query["Resources"] = request.Resources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApisecStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApisecStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of the API security module for protected objects or protected object groups.
//
// @param request - ModifyApisecStatusRequest
//
// @return ModifyApisecStatusResponse
func (client *Client) ModifyApisecStatus(request *ModifyApisecStatusRequest) (_result *ModifyApisecStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApisecStatusResponse{}
	_body, _err := client.ModifyApisecStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a service that is added to Web Application Firewall (WAF).
//
// @param tmpReq - ModifyCloudResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudResourceResponse
func (client *Client) ModifyCloudResourceWithOptions(tmpReq *ModifyCloudResourceRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudResourceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyCloudResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Listen) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, dara.String("Listen"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Redirect) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, dara.String("Redirect"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ListenShrink) {
		query["Listen"] = request.ListenShrink
	}

	if !dara.IsNil(request.RedirectShrink) {
		query["Redirect"] = request.RedirectShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCloudResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a service that is added to Web Application Firewall (WAF).
//
// @param request - ModifyCloudResourceRequest
//
// @return ModifyCloudResourceResponse
func (client *Client) ModifyCloudResource(request *ModifyCloudResourceRequest) (_result *ModifyCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCloudResourceResponse{}
	_body, _err := client.ModifyCloudResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the default Secure Sockets Layer (SSL) and Transport Layer Security (TLS) settings.
//
// @param request - ModifyDefaultHttpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefaultHttpsResponse
func (client *Client) ModifyDefaultHttpsWithOptions(request *ModifyDefaultHttpsRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefaultHttpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertId) {
		query["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CipherSuite) {
		query["CipherSuite"] = request.CipherSuite
	}

	if !dara.IsNil(request.CustomCiphers) {
		query["CustomCiphers"] = request.CustomCiphers
	}

	if !dara.IsNil(request.EnableTLSv3) {
		query["EnableTLSv3"] = request.EnableTLSv3
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TLSVersion) {
		query["TLSVersion"] = request.TLSVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefaultHttps"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefaultHttpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the default Secure Sockets Layer (SSL) and Transport Layer Security (TLS) settings.
//
// @param request - ModifyDefaultHttpsRequest
//
// @return ModifyDefaultHttpsResponse
func (client *Client) ModifyDefaultHttps(request *ModifyDefaultHttpsRequest) (_result *ModifyDefaultHttpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefaultHttpsResponse{}
	_body, _err := client.ModifyDefaultHttpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protected object group.
//
// @param request - ModifyDefenseResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseResourceGroupResponse
func (client *Client) ModifyDefenseResourceGroupWithOptions(request *ModifyDefenseResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddList) {
		query["AddList"] = request.AddList
	}

	if !dara.IsNil(request.DeleteList) {
		query["DeleteList"] = request.DeleteList
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseResourceGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protected object group.
//
// @param request - ModifyDefenseResourceGroupRequest
//
// @return ModifyDefenseResourceGroupResponse
func (client *Client) ModifyDefenseResourceGroup(request *ModifyDefenseResourceGroupRequest) (_result *ModifyDefenseResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseResourceGroupResponse{}
	_body, _err := client.ModifyDefenseResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the cookie settings of a protected object and the method to identify the originating IP addresses of clients.
//
// @param request - ModifyDefenseResourceXffRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseResourceXffResponse
func (client *Client) ModifyDefenseResourceXffWithOptions(request *ModifyDefenseResourceXffRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseResourceXffResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcwCookieStatus) {
		query["AcwCookieStatus"] = request.AcwCookieStatus
	}

	if !dara.IsNil(request.AcwSecureStatus) {
		query["AcwSecureStatus"] = request.AcwSecureStatus
	}

	if !dara.IsNil(request.AcwV3SecureStatus) {
		query["AcwV3SecureStatus"] = request.AcwV3SecureStatus
	}

	if !dara.IsNil(request.CustomHeaders) {
		query["CustomHeaders"] = request.CustomHeaders
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResponseHeaders) {
		query["ResponseHeaders"] = request.ResponseHeaders
	}

	if !dara.IsNil(request.XffStatus) {
		query["XffStatus"] = request.XffStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseResourceXff"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseResourceXffResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the cookie settings of a protected object and the method to identify the originating IP addresses of clients.
//
// @param request - ModifyDefenseResourceXffRequest
//
// @return ModifyDefenseResourceXffResponse
func (client *Client) ModifyDefenseResourceXff(request *ModifyDefenseResourceXffRequest) (_result *ModifyDefenseResourceXffResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseResourceXffResponse{}
	_body, _err := client.ModifyDefenseResourceXffWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protection rule.
//
// @param request - ModifyDefenseRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseRuleResponse
func (client *Client) ModifyDefenseRuleWithOptions(request *ModifyDefenseRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Rules) {
		body["Rules"] = request.Rules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protection rule.
//
// @param request - ModifyDefenseRuleRequest
//
// @return ModifyDefenseRuleResponse
func (client *Client) ModifyDefenseRule(request *ModifyDefenseRuleRequest) (_result *ModifyDefenseRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseRuleResponse{}
	_body, _err := client.ModifyDefenseRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the cached page of a website that is protected based on a website tamper-proofing rule.
//
// @param request - ModifyDefenseRuleCacheRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseRuleCacheResponse
func (client *Client) ModifyDefenseRuleCacheWithOptions(request *ModifyDefenseRuleCacheRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseRuleCacheResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseRuleCache"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseRuleCacheResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the cached page of a website that is protected based on a website tamper-proofing rule.
//
// @param request - ModifyDefenseRuleCacheRequest
//
// @return ModifyDefenseRuleCacheResponse
func (client *Client) ModifyDefenseRuleCache(request *ModifyDefenseRuleCacheRequest) (_result *ModifyDefenseRuleCacheResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseRuleCacheResponse{}
	_body, _err := client.ModifyDefenseRuleCacheWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a protection rule.
//
// @param request - ModifyDefenseRuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseRuleStatusResponse
func (client *Client) ModifyDefenseRuleStatusWithOptions(request *ModifyDefenseRuleStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseRuleStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefenseType) {
		query["DefenseType"] = request.DefenseType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseRuleStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a protection rule.
//
// @param request - ModifyDefenseRuleStatusRequest
//
// @return ModifyDefenseRuleStatusResponse
func (client *Client) ModifyDefenseRuleStatus(request *ModifyDefenseRuleStatusRequest) (_result *ModifyDefenseRuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseRuleStatusResponse{}
	_body, _err := client.ModifyDefenseRuleStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户防护场景的配置
//
// @param request - ModifyDefenseSceneConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseSceneConfigResponse
func (client *Client) ModifyDefenseSceneConfigWithOptions(request *ModifyDefenseSceneConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseSceneConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.ConfigValue) {
		query["ConfigValue"] = request.ConfigValue
	}

	if !dara.IsNil(request.DefenseScene) {
		query["DefenseScene"] = request.DefenseScene
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseSceneConfig"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseSceneConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户防护场景的配置
//
// @param request - ModifyDefenseSceneConfigRequest
//
// @return ModifyDefenseSceneConfigResponse
func (client *Client) ModifyDefenseSceneConfig(request *ModifyDefenseSceneConfigRequest) (_result *ModifyDefenseSceneConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseSceneConfigResponse{}
	_body, _err := client.ModifyDefenseSceneConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protection rule template.
//
// @param request - ModifyDefenseTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseTemplateResponse
func (client *Client) ModifyDefenseTemplateWithOptions(request *ModifyDefenseTemplateRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseTemplate"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a protection rule template.
//
// @param request - ModifyDefenseTemplateRequest
//
// @return ModifyDefenseTemplateResponse
func (client *Client) ModifyDefenseTemplate(request *ModifyDefenseTemplateRequest) (_result *ModifyDefenseTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseTemplateResponse{}
	_body, _err := client.ModifyDefenseTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the status of a protection rule template.
//
// @param request - ModifyDefenseTemplateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDefenseTemplateStatusResponse
func (client *Client) ModifyDefenseTemplateStatusWithOptions(request *ModifyDefenseTemplateStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateStatus) {
		query["TemplateStatus"] = request.TemplateStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDefenseTemplateStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the status of a protection rule template.
//
// @param request - ModifyDefenseTemplateStatusRequest
//
// @return ModifyDefenseTemplateStatusResponse
func (client *Client) ModifyDefenseTemplateStatus(request *ModifyDefenseTemplateStatusRequest) (_result *ModifyDefenseTemplateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDefenseTemplateStatusResponse{}
	_body, _err := client.ModifyDefenseTemplateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a domain name that is added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param tmpReq - ModifyDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDomainResponse
func (client *Client) ModifyDomainWithOptions(tmpReq *ModifyDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyDomainResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDomainShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Listen) {
		request.ListenShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Listen, dara.String("Listen"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Redirect) {
		request.RedirectShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Redirect, dara.String("Redirect"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessType) {
		query["AccessType"] = request.AccessType
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ListenShrink) {
		query["Listen"] = request.ListenShrink
	}

	if !dara.IsNil(request.RedirectShrink) {
		query["Redirect"] = request.RedirectShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDomain"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a domain name that is added to Web Application Firewall (WAF) in CNAME record mode.
//
// @param request - ModifyDomainRequest
//
// @return ModifyDomainResponse
func (client *Client) ModifyDomain(request *ModifyDomainRequest) (_result *ModifyDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDomainResponse{}
	_body, _err := client.ModifyDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-adds a domain name that is penalized for failing to obtain an Internet Content Provider (ICP) filing to Web Application Firewall (WAF).
//
// @param request - ModifyDomainPunishStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDomainPunishStatusResponse
func (client *Client) ModifyDomainPunishStatusWithOptions(request *ModifyDomainPunishStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyDomainPunishStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDomainPunishStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDomainPunishStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-adds a domain name that is penalized for failing to obtain an Internet Content Provider (ICP) filing to Web Application Firewall (WAF).
//
// @param request - ModifyDomainPunishStatusRequest
//
// @return ModifyDomainPunishStatusResponse
func (client *Client) ModifyDomainPunishStatus(request *ModifyDomainPunishStatusRequest) (_result *ModifyDomainPunishStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDomainPunishStatusResponse{}
	_body, _err := client.ModifyDomainPunishStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies information about a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudClusterResponse
func (client *Client) ModifyHybridCloudClusterWithOptions(request *ModifyHybridCloudClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessMode) {
		query["AccessMode"] = request.AccessMode
	}

	if !dara.IsNil(request.AccessRegion) {
		query["AccessRegion"] = request.AccessRegion
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.HttpPorts) {
		query["HttpPorts"] = request.HttpPorts
	}

	if !dara.IsNil(request.HttpsPorts) {
		query["HttpsPorts"] = request.HttpsPorts
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogFieldsNotReturned) {
		query["LogFieldsNotReturned"] = request.LogFieldsNotReturned
	}

	if !dara.IsNil(request.ProtectionServerCount) {
		query["ProtectionServerCount"] = request.ProtectionServerCount
	}

	if !dara.IsNil(request.ProxyStatus) {
		query["ProxyStatus"] = request.ProxyStatus
	}

	if !dara.IsNil(request.ProxyType) {
		query["ProxyType"] = request.ProxyType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudCluster"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies information about a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudClusterRequest
//
// @return ModifyHybridCloudClusterResponse
func (client *Client) ModifyHybridCloudCluster(request *ModifyHybridCloudClusterRequest) (_result *ModifyHybridCloudClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudClusterResponse{}
	_body, _err := client.ModifyHybridCloudClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables manual bypass for a hybrid cloud cluster whose type is set to SDK Integration Mode.
//
// @param request - ModifyHybridCloudClusterBypassStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudClusterBypassStatusResponse
func (client *Client) ModifyHybridCloudClusterBypassStatusWithOptions(request *ModifyHybridCloudClusterBypassStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudClusterBypassStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterResourceId) {
		query["ClusterResourceId"] = request.ClusterResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudClusterBypassStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudClusterBypassStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables manual bypass for a hybrid cloud cluster whose type is set to SDK Integration Mode.
//
// @param request - ModifyHybridCloudClusterBypassStatusRequest
//
// @return ModifyHybridCloudClusterBypassStatusResponse
func (client *Client) ModifyHybridCloudClusterBypassStatus(request *ModifyHybridCloudClusterBypassStatusRequest) (_result *ModifyHybridCloudClusterBypassStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudClusterBypassStatusResponse{}
	_body, _err := client.ModifyHybridCloudClusterBypassStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the rule of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudClusterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudClusterRuleResponse
func (client *Client) ModifyHybridCloudClusterRuleWithOptions(request *ModifyHybridCloudClusterRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudClusterRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterRuleResourceId) {
		query["ClusterRuleResourceId"] = request.ClusterRuleResourceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleConfig) {
		query["RuleConfig"] = request.RuleConfig
	}

	if !dara.IsNil(request.RuleStatus) {
		query["RuleStatus"] = request.RuleStatus
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudClusterRule"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudClusterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the rule of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudClusterRuleRequest
//
// @return ModifyHybridCloudClusterRuleResponse
func (client *Client) ModifyHybridCloudClusterRule(request *ModifyHybridCloudClusterRuleRequest) (_result *ModifyHybridCloudClusterRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudClusterRuleResponse{}
	_body, _err := client.ModifyHybridCloudClusterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a node group in a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudGroupResponse
func (client *Client) ModifyHybridCloudGroupWithOptions(request *ModifyHybridCloudGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudGroup"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a node group in a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupRequest
//
// @return ModifyHybridCloudGroupResponse
func (client *Client) ModifyHybridCloudGroup(request *ModifyHybridCloudGroupRequest) (_result *ModifyHybridCloudGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudGroupResponse{}
	_body, _err := client.ModifyHybridCloudGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a node to a node group of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupExpansionServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudGroupExpansionServerResponse
func (client *Client) ModifyHybridCloudGroupExpansionServerWithOptions(request *ModifyHybridCloudGroupExpansionServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudGroupExpansionServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mids) {
		query["Mids"] = request.Mids
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudGroupExpansionServer"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudGroupExpansionServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a node to a node group of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupExpansionServerRequest
//
// @return ModifyHybridCloudGroupExpansionServerResponse
func (client *Client) ModifyHybridCloudGroupExpansionServer(request *ModifyHybridCloudGroupExpansionServerRequest) (_result *ModifyHybridCloudGroupExpansionServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudGroupExpansionServerResponse{}
	_body, _err := client.ModifyHybridCloudGroupExpansionServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a node from a node group of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupShrinkServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudGroupShrinkServerResponse
func (client *Client) ModifyHybridCloudGroupShrinkServerWithOptions(request *ModifyHybridCloudGroupShrinkServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudGroupShrinkServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mids) {
		query["Mids"] = request.Mids
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudGroupShrinkServer"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudGroupShrinkServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a node from a node group of a hybrid cloud cluster.
//
// @param request - ModifyHybridCloudGroupShrinkServerRequest
//
// @return ModifyHybridCloudGroupShrinkServerResponse
func (client *Client) ModifyHybridCloudGroupShrinkServer(request *ModifyHybridCloudGroupShrinkServerRequest) (_result *ModifyHybridCloudGroupShrinkServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudGroupShrinkServerResponse{}
	_body, _err := client.ModifyHybridCloudGroupShrinkServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the traffic redirection status of a hybrid cloud cluster by using an SDK.
//
// @param request - ModifyHybridCloudSdkPullinStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudSdkPullinStatusResponse
func (client *Client) ModifyHybridCloudSdkPullinStatusWithOptions(request *ModifyHybridCloudSdkPullinStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudSdkPullinStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mid) {
		query["Mid"] = request.Mid
	}

	if !dara.IsNil(request.PullinStatus) {
		query["PullinStatus"] = request.PullinStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudSdkPullinStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudSdkPullinStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the traffic redirection status of a hybrid cloud cluster by using an SDK.
//
// @param request - ModifyHybridCloudSdkPullinStatusRequest
//
// @return ModifyHybridCloudSdkPullinStatusResponse
func (client *Client) ModifyHybridCloudSdkPullinStatus(request *ModifyHybridCloudSdkPullinStatusRequest) (_result *ModifyHybridCloudSdkPullinStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudSdkPullinStatusResponse{}
	_body, _err := client.ModifyHybridCloudSdkPullinStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a hybrid cloud node.
//
// @param request - ModifyHybridCloudServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHybridCloudServerResponse
func (client *Client) ModifyHybridCloudServerWithOptions(request *ModifyHybridCloudServerRequest, runtime *dara.RuntimeOptions) (_result *ModifyHybridCloudServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Continents) {
		query["Continents"] = request.Continents
	}

	if !dara.IsNil(request.CustomName) {
		query["CustomName"] = request.CustomName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mid) {
		query["Mid"] = request.Mid
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	if !dara.IsNil(request.RegionCode) {
		query["RegionCode"] = request.RegionCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHybridCloudServer"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHybridCloudServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a hybrid cloud node.
//
// @param request - ModifyHybridCloudServerRequest
//
// @return ModifyHybridCloudServerResponse
func (client *Client) ModifyHybridCloudServer(request *ModifyHybridCloudServerRequest) (_result *ModifyHybridCloudServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHybridCloudServerResponse{}
	_body, _err := client.ModifyHybridCloudServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an IP address blacklist for major event protection.
//
// @param request - ModifyMajorProtectionBlackIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMajorProtectionBlackIpResponse
func (client *Client) ModifyMajorProtectionBlackIpWithOptions(request *ModifyMajorProtectionBlackIpRequest, runtime *dara.RuntimeOptions) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IpList) {
		query["IpList"] = request.IpList
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMajorProtectionBlackIp"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an IP address blacklist for major event protection.
//
// @param request - ModifyMajorProtectionBlackIpRequest
//
// @return ModifyMajorProtectionBlackIpResponse
func (client *Client) ModifyMajorProtectionBlackIp(request *ModifyMajorProtectionBlackIpRequest) (_result *ModifyMajorProtectionBlackIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMajorProtectionBlackIpResponse{}
	_body, _err := client.ModifyMajorProtectionBlackIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about members that are added for multi-account management.
//
// @param request - ModifyMemberAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMemberAccountResponse
func (client *Client) ModifyMemberAccountWithOptions(request *ModifyMemberAccountRequest, runtime *dara.RuntimeOptions) (_result *ModifyMemberAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberAccountId) {
		query["MemberAccountId"] = request.MemberAccountId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMemberAccount"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMemberAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about members that are added for multi-account management.
//
// @param request - ModifyMemberAccountRequest
//
// @return ModifyMemberAccountResponse
func (client *Client) ModifyMemberAccount(request *ModifyMemberAccountRequest) (_result *ModifyMemberAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMemberAccountResponse{}
	_body, _err := client.ModifyMemberAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the protection status of Web Application Firewall (WAF).
//
// @param request - ModifyPauseProtectionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPauseProtectionStatusResponse
func (client *Client) ModifyPauseProtectionStatusWithOptions(request *ModifyPauseProtectionStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyPauseProtectionStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PauseStatus) {
		query["PauseStatus"] = request.PauseStatus
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPauseProtectionStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPauseProtectionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the protection status of Web Application Firewall (WAF).
//
// @param request - ModifyPauseProtectionStatusRequest
//
// @return ModifyPauseProtectionStatusResponse
func (client *Client) ModifyPauseProtectionStatus(request *ModifyPauseProtectionStatusRequest) (_result *ModifyPauseProtectionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPauseProtectionStatusResponse{}
	_body, _err := client.ModifyPauseProtectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the log collection feature for a protected object.
//
// @param request - ModifyResourceLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyResourceLogStatusResponse
func (client *Client) ModifyResourceLogStatusWithOptions(request *ModifyResourceLogStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyResourceLogStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyResourceLogStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the log collection feature for a protected object.
//
// @param request - ModifyResourceLogStatusRequest
//
// @return ModifyResourceLogStatusResponse
func (client *Client) ModifyResourceLogStatus(request *ModifyResourceLogStatusRequest) (_result *ModifyResourceLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyResourceLogStatusResponse{}
	_body, _err := client.ModifyResourceLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates or disassociates a protected object or protected object group with or from a protection rule template.
//
// @param request - ModifyTemplateResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTemplateResourcesResponse
func (client *Client) ModifyTemplateResourcesWithOptions(request *ModifyTemplateResourcesRequest, runtime *dara.RuntimeOptions) (_result *ModifyTemplateResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindResourceGroups) {
		query["BindResourceGroups"] = request.BindResourceGroups
	}

	if !dara.IsNil(request.BindResources) {
		query["BindResources"] = request.BindResources
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UnbindResourceGroups) {
		query["UnbindResourceGroups"] = request.UnbindResourceGroups
	}

	if !dara.IsNil(request.UnbindResources) {
		query["UnbindResources"] = request.UnbindResources
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTemplateResources"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates or disassociates a protected object or protected object group with or from a protection rule template.
//
// @param request - ModifyTemplateResourcesRequest
//
// @return ModifyTemplateResourcesResponse
func (client *Client) ModifyTemplateResources(request *ModifyTemplateResourcesRequest) (_result *ModifyTemplateResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTemplateResourcesResponse{}
	_body, _err := client.ModifyTemplateResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户日志服务的默认字段配置
//
// @param request - ModifyUserLogFieldConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserLogFieldConfigResponse
func (client *Client) ModifyUserLogFieldConfigWithOptions(request *ModifyUserLogFieldConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserLogFieldConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeliveryType) {
		query["DeliveryType"] = request.DeliveryType
	}

	if !dara.IsNil(request.ExtendConfig) {
		query["ExtendConfig"] = request.ExtendConfig
	}

	if !dara.IsNil(request.FieldList) {
		query["FieldList"] = request.FieldList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogDeliveryStrategy) {
		query["LogDeliveryStrategy"] = request.LogDeliveryStrategy
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserLogFieldConfig"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserLogFieldConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户日志服务的默认字段配置
//
// @param request - ModifyUserLogFieldConfigRequest
//
// @return ModifyUserLogFieldConfigResponse
func (client *Client) ModifyUserLogFieldConfig(request *ModifyUserLogFieldConfigRequest) (_result *ModifyUserLogFieldConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserLogFieldConfigResponse{}
	_body, _err := client.ModifyUserLogFieldConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通或关闭WAF日志服务
//
// @param request - ModifyUserWafLogStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserWafLogStatusResponse
func (client *Client) ModifyUserWafLogStatusWithOptions(request *ModifyUserWafLogStatusRequest, runtime *dara.RuntimeOptions) (_result *ModifyUserWafLogStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogRegionId) {
		query["LogRegionId"] = request.LogRegionId
	}

	if !dara.IsNil(request.LogStatus) {
		query["LogStatus"] = request.LogStatus
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserWafLogStatus"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserWafLogStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通或关闭WAF日志服务
//
// @param request - ModifyUserWafLogStatusRequest
//
// @return ModifyUserWafLogStatusResponse
func (client *Client) ModifyUserWafLogStatus(request *ModifyUserWafLogStatusRequest) (_result *ModifyUserWafLogStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyUserWafLogStatusResponse{}
	_body, _err := client.ModifyUserWafLogStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新接入云产品
//
// @param request - ReCreateCloudResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReCreateCloudResourceResponse
func (client *Client) ReCreateCloudResourceWithOptions(request *ReCreateCloudResourceRequest, runtime *dara.RuntimeOptions) (_result *ReCreateCloudResourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceInstanceId) {
		query["ResourceInstanceId"] = request.ResourceInstanceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceProduct) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReCreateCloudResource"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReCreateCloudResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新接入云产品
//
// @param request - ReCreateCloudResourceRequest
//
// @return ReCreateCloudResourceResponse
func (client *Client) ReCreateCloudResource(request *ReCreateCloudResourceRequest) (_result *ReCreateCloudResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReCreateCloudResourceResponse{}
	_body, _err := client.ReCreateCloudResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a Web Application Firewall (WAF) 3.0 instance.
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a Web Application Firewall (WAF) 3.0 instance.
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes Elastic Compute Service (ECS), Classic Load Balancer (CLB), and Network Load Balancer (NLB) instances to Web Application Firewall (WAF).
//
// Description:
//
// SyncProductInstance is an asynchronous operation. You can call the [DescribeProductInstances](https://help.aliyun.com/document_detail/2743168.html) operation to query the status of the task.
//
// @param request - SyncProductInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncProductInstanceResponse
func (client *Client) SyncProductInstanceWithOptions(request *SyncProductInstanceRequest, runtime *dara.RuntimeOptions) (_result *SyncProductInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncProductInstance"),
		Version:     dara.String("2021-10-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncProductInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes Elastic Compute Service (ECS), Classic Load Balancer (CLB), and Network Load Balancer (NLB) instances to Web Application Firewall (WAF).
//
// Description:
//
// SyncProductInstance is an asynchronous operation. You can call the [DescribeProductInstances](https://help.aliyun.com/document_detail/2743168.html) operation to query the status of the task.
//
// @param request - SyncProductInstanceRequest
//
// @return SyncProductInstanceResponse
func (client *Client) SyncProductInstance(request *SyncProductInstanceRequest) (_result *SyncProductInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncProductInstanceResponse{}
	_body, _err := client.SyncProductInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to resources.
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
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2021-10-01"),
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
// Adds tags to resources.
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
// Removes tags from resources and then deletes the tags.
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Version:     dara.String("2021-10-01"),
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
// Removes tags from resources and then deletes the tags.
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
