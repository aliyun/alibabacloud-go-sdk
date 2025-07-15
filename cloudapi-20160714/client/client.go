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
		"cn-qingdao":                  dara.String("apigateway.cn-qingdao.aliyuncs.com"),
		"cn-beijing":                  dara.String("apigateway.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("apigateway.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                dara.String("apigateway.cn-huhehaote.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("apigateway.cn-wulanchabu.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("apigateway.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":                 dara.String("apigateway.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("apigateway.cn-shenzhen.aliyuncs.com"),
		"cn-heyuan":                   dara.String("apigateway.cn-heyuan.aliyuncs.com"),
		"cn-guangzhou":                dara.String("apigateway.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":                  dara.String("apigateway.cn-chengdu.aliyuncs.com"),
		"cn-hongkong":                 dara.String("apigateway.cn-hongkong.aliyuncs.com"),
		"ap-northeast-1":              dara.String("apigateway.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1":              dara.String("apigateway.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              dara.String("apigateway.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              dara.String("apigateway.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              dara.String("apigateway.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-6":              dara.String("apigateway.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-7":              dara.String("apigateway.ap-southeast-7.aliyuncs.com"),
		"us-east-1":                   dara.String("apigateway.us-east-1.aliyuncs.com"),
		"us-west-1":                   dara.String("apigateway.us-west-1.aliyuncs.com"),
		"eu-west-1":                   dara.String("apigateway.eu-west-1.aliyuncs.com"),
		"eu-central-1":                dara.String("apigateway.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("apigateway.ap-south-1.aliyuncs.com"),
		"me-east-1":                   dara.String("apigateway.me-east-1.aliyuncs.com"),
		"me-central-1":                dara.String("apigateway.me-central-1.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("apigateway.cn-hangzhou-finance.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("apigateway.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("apigateway.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("apigateway.cn-north-2-gov-1.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("apigateway.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("apigateway.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("apigateway.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("apigateway.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("apigateway.aliyuncs.com"),
		"cn-edge-1":                   dara.String("apigateway.aliyuncs.com"),
		"cn-fujian":                   dara.String("apigateway.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("apigateway.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("apigateway.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("apigateway.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("apigateway.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("apigateway.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("apigateway.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("apigateway.cn-shanghai-inner.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("apigateway.aliyuncs.com"),
		"cn-wuhan":                    dara.String("apigateway.aliyuncs.com"),
		"cn-yushanfang":               dara.String("apigateway.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("apigateway.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("apigateway.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("apigateway.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("apigateway.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("apigateway.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("apigateway.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cloudapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Unpublishes a specified API from a specified runtime environment.
//
// Description:
//
//	  This operation is intended for API providers and is the opposite of DeployApi.
//
//		- An API can be unpublished from a specified runtime environment in under 5 seconds.
//
//		- An unpublished API cannot be called in the specified runtime environment.
//
// @param request - AbolishApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishApiResponse
func (client *Client) AbolishApiWithOptions(request *AbolishApiRequest, runtime *dara.RuntimeOptions) (_result *AbolishApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbolishApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbolishApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes a specified API from a specified runtime environment.
//
// Description:
//
//	  This operation is intended for API providers and is the opposite of DeployApi.
//
//		- An API can be unpublished from a specified runtime environment in under 5 seconds.
//
//		- An unpublished API cannot be called in the specified runtime environment.
//
// @param request - AbolishApiRequest
//
// @return AbolishApiResponse
func (client *Client) AbolishApi(request *AbolishApiRequest) (_result *AbolishApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbolishApiResponse{}
	_body, _err := client.AbolishApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Adds an IP address entry to the access control polocy of an instance.
//
// @param request - AddAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAccessControlListEntryResponse
func (client *Client) AddAccessControlListEntryWithOptions(request *AddAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *AddAccessControlListEntryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAccessControlListEntry"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAccessControlListEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Adds an IP address entry to the access control polocy of an instance.
//
// @param request - AddAccessControlListEntryRequest
//
// @return AddAccessControlListEntryResponse
func (client *Client) AddAccessControlListEntry(request *AddAccessControlListEntryRequest) (_result *AddAccessControlListEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAccessControlListEntryResponse{}
	_body, _err := client.AddAccessControlListEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a policy to an existing ACL.
//
// Description:
//
// When you call this operation, note that:
//
//   - This operation is intended for API providers.
//
//   - An added policy immediately takes effect on all APIs that are bound to the access control list (ACL).
//
//   - A maximum of 100 policies can be added to an ACL.
//
// @param request - AddIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpControlPolicyItemResponse
func (client *Client) AddIpControlPolicyItemWithOptions(request *AddIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *AddIpControlPolicyItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CidrIp) {
		query["CidrIp"] = request.CidrIp
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a policy to an existing ACL.
//
// Description:
//
// When you call this operation, note that:
//
//   - This operation is intended for API providers.
//
//   - An added policy immediately takes effect on all APIs that are bound to the access control list (ACL).
//
//   - A maximum of 100 policies can be added to an ACL.
//
// @param request - AddIpControlPolicyItemRequest
//
// @return AddIpControlPolicyItemResponse
func (client *Client) AddIpControlPolicyItem(request *AddIpControlPolicyItemRequest) (_result *AddIpControlPolicyItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddIpControlPolicyItemResponse{}
	_body, _err := client.AddIpControlPolicyItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a custom special policy to a specified throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the input SpecialKey already exists, the previous configuration is overwritten. Use caution when calling this operation.
//
//		- Special throttling policies must be added to an existing throttling policy, and can take effect on all the APIs to which the throttling policy is bound.
//
// @param request - AddTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrafficSpecialControlResponse
func (client *Client) AddTrafficSpecialControlWithOptions(request *AddTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *AddTrafficSpecialControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecialKey) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !dara.IsNil(request.SpecialType) {
		query["SpecialType"] = request.SpecialType
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficValue) {
		query["TrafficValue"] = request.TrafficValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom special policy to a specified throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the input SpecialKey already exists, the previous configuration is overwritten. Use caution when calling this operation.
//
//		- Special throttling policies must be added to an existing throttling policy, and can take effect on all the APIs to which the throttling policy is bound.
//
// @param request - AddTrafficSpecialControlRequest
//
// @return AddTrafficSpecialControlResponse
func (client *Client) AddTrafficSpecialControl(request *AddTrafficSpecialControlRequest) (_result *AddTrafficSpecialControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTrafficSpecialControlResponse{}
	_body, _err := client.AddTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an internal domain name resolution with a dedicated instance.
//
// Description:
//
// An internal domain name resolution can be associated only with a dedicated instance, not with a shared instance or shared instance cluster.
//
// @param tmpReq - AssociateInstanceWithPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateInstanceWithPrivateDNSResponse
func (client *Client) AssociateInstanceWithPrivateDNSWithOptions(tmpReq *AssociateInstanceWithPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *AssociateInstanceWithPrivateDNSResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AssociateInstanceWithPrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntranetDomains) {
		request.IntranetDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntranetDomains, dara.String("IntranetDomains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomainsShrink) {
		body["IntranetDomains"] = request.IntranetDomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateInstanceWithPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an internal domain name resolution with a dedicated instance.
//
// Description:
//
// An internal domain name resolution can be associated only with a dedicated instance, not with a shared instance or shared instance cluster.
//
// @param request - AssociateInstanceWithPrivateDNSRequest
//
// @return AssociateInstanceWithPrivateDNSResponse
func (client *Client) AssociateInstanceWithPrivateDNS(request *AssociateInstanceWithPrivateDNSRequest) (_result *AssociateInstanceWithPrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.AssociateInstanceWithPrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches APIs to an API product. If the API product does not exist, the system automatically creates the API product.
//
// @param request - AttachApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApiProductResponse
func (client *Client) AttachApiProductWithOptions(request *AttachApiProductRequest, runtime *dara.RuntimeOptions) (_result *AttachApiProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.Apis) {
		query["Apis"] = request.Apis
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachApiProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches APIs to an API product. If the API product does not exist, the system automatically creates the API product.
//
// @param request - AttachApiProductRequest
//
// @return AttachApiProductResponse
func (client *Client) AttachApiProduct(request *AttachApiProductRequest) (_result *AttachApiProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachApiProductResponse{}
	_body, _err := client.AttachApiProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attach plugin to API group.
//
// @param request - AttachGroupPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachGroupPluginResponse
func (client *Client) AttachGroupPluginWithOptions(request *AttachGroupPluginRequest, runtime *dara.RuntimeOptions) (_result *AttachGroupPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachGroupPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachGroupPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach plugin to API group.
//
// @param request - AttachGroupPluginRequest
//
// @return AttachGroupPluginResponse
func (client *Client) AttachGroupPlugin(request *AttachGroupPluginRequest) (_result *AttachGroupPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachGroupPluginResponse{}
	_body, _err := client.AttachGroupPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a plug-in to an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can only bind plug-ins to published APIs.
//
//		- The plug-in takes effect immediately after it is bound to an API.
//
//		- If you bind a different plug-in to an API, this plug-in takes effect immediately.
//
// @param request - AttachPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPluginResponse
func (client *Client) AttachPluginWithOptions(request *AttachPluginRequest, runtime *dara.RuntimeOptions) (_result *AttachPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a plug-in to an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can only bind plug-ins to published APIs.
//
//		- The plug-in takes effect immediately after it is bound to an API.
//
//		- If you bind a different plug-in to an API, this plug-in takes effect immediately.
//
// @param request - AttachPluginRequest
//
// @return AttachPluginResponse
func (client *Client) AttachPlugin(request *AttachPluginRequest) (_result *AttachPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachPluginResponse{}
	_body, _err := client.AttachPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unpublishes multiple published APIs at a time.
//
// @param request - BatchAbolishApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAbolishApisResponse
func (client *Client) BatchAbolishApisWithOptions(request *BatchAbolishApisRequest, runtime *dara.RuntimeOptions) (_result *BatchAbolishApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAbolishApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAbolishApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes multiple published APIs at a time.
//
// @param request - BatchAbolishApisRequest
//
// @return BatchAbolishApisResponse
func (client *Client) BatchAbolishApis(request *BatchAbolishApisRequest) (_result *BatchAbolishApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchAbolishApisResponse{}
	_body, _err := client.BatchAbolishApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes multiple APIs at a time.
//
// @param request - BatchDeployApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeployApisResponse
func (client *Client) BatchDeployApisWithOptions(request *BatchDeployApisRequest, runtime *dara.RuntimeOptions) (_result *BatchDeployApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["Api"] = request.Api
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeployApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeployApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes multiple APIs at a time.
//
// @param request - BatchDeployApisRequest
//
// @return BatchDeployApisResponse
func (client *Client) BatchDeployApis(request *BatchDeployApisRequest) (_result *BatchDeployApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeployApisResponse{}
	_body, _err := client.BatchDeployApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Creates an Access Control List (ACL). Each user is allowed to create five ACLs in each region.
//
// @param request - CreateAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessControlListResponse
func (client *Client) CreateAccessControlListWithOptions(request *CreateAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessControlListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessControlList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Creates an Access Control List (ACL). Each user is allowed to create five ACLs in each region.
//
// @param request - CreateAccessControlListRequest
//
// @return CreateAccessControlListResponse
func (client *Client) CreateAccessControlList(request *CreateAccessControlListRequest) (_result *CreateAccessControlListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessControlListResponse{}
	_body, _err := client.CreateAccessControlListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of an API must be unique within an API group.
//
//		- A request path must be unique within an API group.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiResponse
func (client *Client) CreateApiWithOptions(request *CreateApiRequest, runtime *dara.RuntimeOptions) (_result *CreateApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendEnable) {
		query["BackendEnable"] = request.BackendEnable
	}

	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpenIdConnectConfig) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !dara.IsNil(request.RequestConfig) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !dara.IsNil(request.ResultBodyModel) {
		query["ResultBodyModel"] = request.ResultBodyModel
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceConfig) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WebSocketApiType) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConstantParameters) {
		body["ConstantParameters"] = request.ConstantParameters
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		body["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		body["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.RequestParameters) {
		body["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.ResultDescriptions) {
		body["ResultDescriptions"] = request.ResultDescriptions
	}

	if !dara.IsNil(request.ResultSample) {
		body["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		body["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.SystemParameters) {
		body["SystemParameters"] = request.SystemParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of an API must be unique within an API group.
//
//		- A request path must be unique within an API group.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateApiRequest
//
// @return CreateApiResponse
func (client *Client) CreateApi(request *CreateApiRequest) (_result *CreateApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiResponse{}
	_body, _err := client.CreateApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建API分组
//
// @param request - CreateApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiGroupResponse
func (client *Client) CreateApiGroupWithOptions(request *CreateApiGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建API分组
//
// @param request - CreateApiGroupRequest
//
// @return CreateApiGroupResponse
func (client *Client) CreateApiGroup(request *CreateApiGroupRequest) (_result *CreateApiGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiGroupResponse{}
	_body, _err := client.CreateApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a variable to an environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - CreateApiStageVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiStageVariableResponse
func (client *Client) CreateApiStageVariableWithOptions(request *CreateApiStageVariableRequest, runtime *dara.RuntimeOptions) (_result *CreateApiStageVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageId) {
		query["StageId"] = request.StageId
	}

	if !dara.IsNil(request.StageRouteModel) {
		query["StageRouteModel"] = request.StageRouteModel
	}

	if !dara.IsNil(request.SupportRoute) {
		query["SupportRoute"] = request.SupportRoute
	}

	if !dara.IsNil(request.VariableName) {
		query["VariableName"] = request.VariableName
	}

	if !dara.IsNil(request.VariableValue) {
		query["VariableValue"] = request.VariableValue
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiStageVariable"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiStageVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a variable to an environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - CreateApiStageVariableRequest
//
// @return CreateApiStageVariableResponse
func (client *Client) CreateApiStageVariable(request *CreateApiStageVariableRequest) (_result *CreateApiStageVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiStageVariableResponse{}
	_body, _err := client.CreateApiStageVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application for calling APIs in API Gateway.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- Each application has a key-value pair which is used for identity verification when you call an API.
//
//		- An application must be authorized to call an API.
//
//		- Each application has only one key-value pair, which can be reset if the pair is leaked.
//
//		- A maximum of 1,000 applications can be created for each Alibaba Cloud account.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *dara.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppSecret) {
		query["AppSecret"] = request.AppSecret
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application for calling APIs in API Gateway.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- Each application has a key-value pair which is used for identity verification when you call an API.
//
//		- An application must be authorized to call an API.
//
//		- Each application has only one key-value pair, which can be reset if the pair is leaked.
//
//		- A maximum of 1,000 applications can be created for each Alibaba Cloud account.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AppCode to an application.
//
// @param request - CreateAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppCodeResponse
func (client *Client) CreateAppCodeWithOptions(request *CreateAppCodeRequest, runtime *dara.RuntimeOptions) (_result *CreateAppCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AppCode to an application.
//
// @param request - CreateAppCodeRequest
//
// @return CreateAppCodeResponse
func (client *Client) CreateAppCode(request *CreateAppCodeRequest) (_result *CreateAppCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppCodeResponse{}
	_body, _err := client.CreateAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an AppKey and AppSecret pair to an application.
//
// @param request - CreateAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppKeyResponse
func (client *Client) CreateAppKeyWithOptions(request *CreateAppKeyRequest, runtime *dara.RuntimeOptions) (_result *CreateAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppSecret) {
		query["AppSecret"] = request.AppSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppKey"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an AppKey and AppSecret pair to an application.
//
// @param request - CreateAppKeyRequest
//
// @return CreateAppKeyResponse
func (client *Client) CreateAppKey(request *CreateAppKeyRequest) (_result *CreateAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppKeyResponse{}
	_body, _err := client.CreateAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backend service in API Gateway.
//
// @param request - CreateBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackendResponse
func (client *Client) CreateBackendWithOptions(request *CreateBackendRequest, runtime *dara.RuntimeOptions) (_result *CreateBackendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.CreateEventBridgeServiceLinkedRole) {
		query["CreateEventBridgeServiceLinkedRole"] = request.CreateEventBridgeServiceLinkedRole
	}

	if !dara.IsNil(request.CreateSlr) {
		query["CreateSlr"] = request.CreateSlr
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backend service in API Gateway.
//
// @param request - CreateBackendRequest
//
// @return CreateBackendResponse
func (client *Client) CreateBackend(request *CreateBackendRequest) (_result *CreateBackendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackendResponse{}
	_body, _err := client.CreateBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建后端服务在环境上的配置
//
// @param request - CreateBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBackendModelResponse
func (client *Client) CreateBackendModelWithOptions(request *CreateBackendModelRequest, runtime *dara.RuntimeOptions) (_result *CreateBackendModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelData) {
		query["BackendModelData"] = request.BackendModelData
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBackendModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建后端服务在环境上的配置
//
// @param request - CreateBackendModelRequest
//
// @return CreateBackendModelResponse
func (client *Client) CreateBackendModel(request *CreateBackendModelRequest) (_result *CreateBackendModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBackendModelResponse{}
	_body, _err := client.CreateBackendModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom dataset.
//
// @param request - CreateDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithOptions(request *CreateDatasetRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.DatasetType) {
		query["DatasetType"] = request.DatasetType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom dataset.
//
// @param request - CreateDatasetRequest
//
// @return CreateDatasetResponse
func (client *Client) CreateDataset(request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义数据集条目
//
// @param request - CreateDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetItemResponse
func (client *Client) CreateDatasetItemWithOptions(request *CreateDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *CreateDatasetItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
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
		Action:      dara.String("CreateDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义数据集条目
//
// @param request - CreateDatasetItemRequest
//
// @return CreateDatasetItemResponse
func (client *Client) CreateDatasetItem(request *CreateDatasetItemRequest) (_result *CreateDatasetItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatasetItemResponse{}
	_body, _err := client.CreateDatasetItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an API Gateway instance.
//
// @param request - CreateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.InstanceCidr) {
		query["InstanceCidr"] = request.InstanceCidr
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.UserVpcId) {
		query["UserVpcId"] = request.UserVpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneVSwitchSecurityGroup) {
		query["ZoneVSwitchSecurityGroup"] = request.ZoneVSwitchSecurityGroup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API Gateway instance.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建内网域名
//
// @param request - CreateIntranetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIntranetDomainResponse
func (client *Client) CreateIntranetDomainWithOptions(request *CreateIntranetDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateIntranetDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIntranetDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIntranetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内网域名
//
// @param request - CreateIntranetDomainRequest
//
// @return CreateIntranetDomainResponse
func (client *Client) CreateIntranetDomain(request *CreateIntranetDomainRequest) (_result *CreateIntranetDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIntranetDomainResponse{}
	_body, _err := client.CreateIntranetDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a region.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An ACL must be bound to an API to take effect. After an ACL is bound to an API, the ACL takes effect on the API immediately.
//
//		- You can add policies to an ACL when you create the ACL.
//
//		- If an ACL does not have any policy, the ACL is ineffective.
//
// @param request - CreateIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIpControlResponse
func (client *Client) CreateIpControlWithOptions(request *CreateIpControlRequest, runtime *dara.RuntimeOptions) (_result *CreateIpControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.IpControlPolicys) {
		query["IpControlPolicys"] = request.IpControlPolicys
	}

	if !dara.IsNil(request.IpControlType) {
		query["IpControlType"] = request.IpControlType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a region.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An ACL must be bound to an API to take effect. After an ACL is bound to an API, the ACL takes effect on the API immediately.
//
//		- You can add policies to an ACL when you create the ACL.
//
//		- If an ACL does not have any policy, the ACL is ineffective.
//
// @param request - CreateIpControlRequest
//
// @return CreateIpControlResponse
func (client *Client) CreateIpControl(request *CreateIpControlRequest) (_result *CreateIpControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIpControlResponse{}
	_body, _err := client.CreateIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Simple Log Service configuration for an API.
//
// @param request - CreateLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogConfigResponse
func (client *Client) CreateLogConfigWithOptions(request *CreateLogConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSlr) {
		query["CreateSlr"] = request.CreateSlr
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlsLogStore) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Simple Log Service configuration for an API.
//
// @param request - CreateLogConfigRequest
//
// @return CreateLogConfigResponse
func (client *Client) CreateLogConfig(request *CreateLogConfigRequest) (_result *CreateLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogConfigResponse{}
	_body, _err := client.CreateLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model for an API group.
//
// Description:
//
//	  For more information about the model definition, see [JSON Schema Draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04?spm=a2c4g.11186623.2.10.2e977ff7p4BpQd).
//
//		- JSON Schema supports only element attributes of the Object type.
//
// @param request - CreateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelResponse
func (client *Client) CreateModelWithOptions(request *CreateModelRequest, runtime *dara.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model for an API group.
//
// Description:
//
//	  For more information about the model definition, see [JSON Schema Draft 4](https://tools.ietf.org/html/draft-zyp-json-schema-04?spm=a2c4g.11186623.2.10.2e977ff7p4BpQd).
//
//		- JSON Schema supports only element attributes of the Object type.
//
// @param request - CreateModelRequest
//
// @return CreateModelResponse
func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables CloudMonitor alerting for a specified API group.
//
// @param request - CreateMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroupWithOptions(request *CreateMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auth) {
		query["Auth"] = request.Auth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RawMonitorGroupId) {
		query["RawMonitorGroupId"] = request.RawMonitorGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMonitorGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables CloudMonitor alerting for a specified API group.
//
// @param request - CreateMonitorGroupRequest
//
// @return CreateMonitorGroupResponse
func (client *Client) CreateMonitorGroup(request *CreateMonitorGroupRequest) (_result *CreateMonitorGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CreateMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The number of plug-ins of the same type that each user can create is limited. Different limits apply to different plug-in types.
//
//		- The plug-in definitions for advanced features are restricted.
//
//		- Plug-ins must be bound to APIs to take effect. After a plug-in is bound, it takes effect on that API immediately.
//
// @param request - CreatePluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginResponse
func (client *Client) CreatePluginWithOptions(request *CreatePluginRequest, runtime *dara.RuntimeOptions) (_result *CreatePluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PluginData) {
		query["PluginData"] = request.PluginData
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The number of plug-ins of the same type that each user can create is limited. Different limits apply to different plug-in types.
//
//		- The plug-in definitions for advanced features are restricted.
//
//		- Plug-ins must be bound to APIs to take effect. After a plug-in is bound, it takes effect on that API immediately.
//
// @param request - CreatePluginRequest
//
// @return CreatePluginResponse
func (client *Client) CreatePlugin(request *CreatePluginRequest) (_result *CreatePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePluginResponse{}
	_body, _err := client.CreatePluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an internal domain name resolution and adds a resolution record.
//
// Description:
//
// An internal domain name resolution of the virtual private cloud (VPC) type can be bound only to traditional dedicated instances. An internal domain name resolution of the A type can be bound only to VPC integration dedicated instances.
//
// @param tmpReq - CreatePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateDNSResponse
func (client *Client) CreatePrivateDNSWithOptions(tmpReq *CreatePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateDNSResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Records) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, dara.String("Records"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RecordsShrink) {
		body["Records"] = request.RecordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal domain name resolution and adds a resolution record.
//
// Description:
//
// An internal domain name resolution of the virtual private cloud (VPC) type can be bound only to traditional dedicated instances. An internal domain name resolution of the A type can be bound only to VPC integration dedicated instances.
//
// @param request - CreatePrivateDNSRequest
//
// @return CreatePrivateDNSResponse
func (client *Client) CreatePrivateDNS(request *CreatePrivateDNSRequest) (_result *CreatePrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePrivateDNSResponse{}
	_body, _err := client.CreatePrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The API operation only creates a key policy. You must call the binding operation to bind the key to an API.
//
//		- After the key is bound to the API, requests sent from API Gateway to the backend service contain signature strings. You can specify whether your backend service verifies these signature strings.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSignatureResponse
func (client *Client) CreateSignatureWithOptions(request *CreateSignatureRequest, runtime *dara.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureKey) {
		query["SignatureKey"] = request.SignatureKey
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	if !dara.IsNil(request.SignatureSecret) {
		query["SignatureSecret"] = request.SignatureSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The API operation only creates a key policy. You must call the binding operation to bind the key to an API.
//
//		- After the key is bound to the API, requests sent from API Gateway to the backend service contain signature strings. You can specify whether your backend service verifies these signature strings.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateSignatureRequest
//
// @return CreateSignatureResponse
func (client *Client) CreateSignature(request *CreateSignatureRequest) (_result *CreateSignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CreateSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Throttling policies must be bound to APIs to take effect. After a policy is bound to an API, it goes into effect on that API immediately.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficControlResponse
func (client *Client) CreateTrafficControlWithOptions(request *CreateTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *CreateTrafficControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDefault) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !dara.IsNil(request.AppDefault) {
		query["AppDefault"] = request.AppDefault
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !dara.IsNil(request.TrafficControlUnit) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !dara.IsNil(request.UserDefault) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Throttling policies must be bound to APIs to take effect. After a policy is bound to an API, it goes into effect on that API immediately.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - CreateTrafficControlRequest
//
// @return CreateTrafficControlResponse
func (client *Client) CreateTrafficControl(request *CreateTrafficControlRequest) (_result *CreateTrafficControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTrafficControlResponse{}
	_body, _err := client.CreateTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Deletes an access control policy.
//
// @param request - DeleteAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessControlListResponse
func (client *Client) DeleteAccessControlListWithOptions(request *DeleteAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessControlListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessControlList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessControlListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated instances. Deletes an access control policy.
//
// @param request - DeleteAccessControlListRequest
//
// @return DeleteAccessControlListResponse
func (client *Client) DeleteAccessControlList(request *DeleteAccessControlListRequest) (_result *DeleteAccessControlListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessControlListResponse{}
	_body, _err := client.DeleteAccessControlListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes all custom special policies of a specified throttling policy.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DeleteAllTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAllTrafficSpecialControlResponse
func (client *Client) DeleteAllTrafficSpecialControlWithOptions(request *DeleteAllTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteAllTrafficSpecialControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAllTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAllTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes all custom special policies of a specified throttling policy.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DeleteAllTrafficSpecialControlRequest
//
// @return DeleteAllTrafficSpecialControlResponse
func (client *Client) DeleteAllTrafficSpecialControl(request *DeleteAllTrafficSpecialControlRequest) (_result *DeleteAllTrafficSpecialControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAllTrafficSpecialControlResponse{}
	_body, _err := client.DeleteAllTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the definition of a specified API.
//
// Description:
//
//	  This operation is intended for API providers and cannot be undone after it is complete.
//
//		- An API that is running in the runtime environment must be unpublished before you can delete the API.****
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiResponse
func (client *Client) DeleteApiWithOptions(request *DeleteApiRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the definition of a specified API.
//
// Description:
//
//	  This operation is intended for API providers and cannot be undone after it is complete.
//
//		- An API that is running in the runtime environment must be unpublished before you can delete the API.****
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiRequest
//
// @return DeleteApiResponse
func (client *Client) DeleteApi(request *DeleteApiRequest) (_result *DeleteApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiResponse{}
	_body, _err := client.DeleteApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An API group that contains APIs cannot be deleted. To delete the API group, you must first delete its APIs.
//
//		- After an API group is deleted, the second-level domain name bound to the API group is automatically invalidated.
//
//		- If the specified API group does not exist, a success response is returned.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiGroupResponse
func (client *Client) DeleteApiGroupWithOptions(request *DeleteApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- An API group that contains APIs cannot be deleted. To delete the API group, you must first delete its APIs.
//
//		- After an API group is deleted, the second-level domain name bound to the API group is automatically invalidated.
//
//		- If the specified API group does not exist, a success response is returned.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteApiGroupRequest
//
// @return DeleteApiGroupResponse
func (client *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (_result *DeleteApiGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiGroupResponse{}
	_body, _err := client.DeleteApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API product. Deleting an API product causes the association between APIs and the deleted API product to be deleted as well. Exercise caution when you delete an API product. If any API in the API product is associated with an application, the API product fails to be deleted.
//
// @param request - DeleteApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiProductResponse
func (client *Client) DeleteApiProductWithOptions(request *DeleteApiProductRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API product. Deleting an API product causes the association between APIs and the deleted API product to be deleted as well. Exercise caution when you delete an API product. If any API in the API product is associated with an application, the API product fails to be deleted.
//
// @param request - DeleteApiProductRequest
//
// @return DeleteApiProductResponse
func (client *Client) DeleteApiProduct(request *DeleteApiProductRequest) (_result *DeleteApiProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiProductResponse{}
	_body, _err := client.DeleteApiProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified variable in a specified environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DeleteApiStageVariableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiStageVariableResponse
func (client *Client) DeleteApiStageVariableWithOptions(request *DeleteApiStageVariableRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiStageVariableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageId) {
		query["StageId"] = request.StageId
	}

	if !dara.IsNil(request.VariableName) {
		query["VariableName"] = request.VariableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiStageVariable"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiStageVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified variable in a specified environment.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DeleteApiStageVariableRequest
//
// @return DeleteApiStageVariableResponse
func (client *Client) DeleteApiStageVariable(request *DeleteApiStageVariableRequest) (_result *DeleteApiStageVariableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiStageVariableResponse{}
	_body, _err := client.DeleteApiStageVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- After an application is deleted, the application and its API authorization cannot be restored.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- After an application is deleted, the application and its API authorization cannot be restored.
//
//		- You can call this operation up to 50 times per second per account.
//
// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the AppCode of an application.
//
// @param request - DeleteAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppCodeResponse
func (client *Client) DeleteAppCodeWithOptions(request *DeleteAppCodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the AppCode of an application.
//
// @param request - DeleteAppCodeRequest
//
// @return DeleteAppCodeResponse
func (client *Client) DeleteAppCode(request *DeleteAppCodeRequest) (_result *DeleteAppCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppCodeResponse{}
	_body, _err := client.DeleteAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the AppKey and AppSecret of an application.
//
// @param request - DeleteAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppKeyResponse
func (client *Client) DeleteAppKeyWithOptions(request *DeleteAppKeyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppKey"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the AppKey and AppSecret of an application.
//
// @param request - DeleteAppKeyRequest
//
// @return DeleteAppKeyResponse
func (client *Client) DeleteAppKey(request *DeleteAppKeyRequest) (_result *DeleteAppKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppKeyResponse{}
	_body, _err := client.DeleteAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backend service.
//
// @param request - DeleteBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackendResponse
func (client *Client) DeleteBackendWithOptions(request *DeleteBackendRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backend service.
//
// @param request - DeleteBackendRequest
//
// @return DeleteBackendResponse
func (client *Client) DeleteBackend(request *DeleteBackendRequest) (_result *DeleteBackendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackendResponse{}
	_body, _err := client.DeleteBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the definition of a backend service in an environment. After the definition is deleted, the API that uses the backend service and is published to this environment will be unpublished.
//
// @param request - DeleteBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackendModelResponse
func (client *Client) DeleteBackendModelWithOptions(request *DeleteBackendModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteBackendModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelId) {
		query["BackendModelId"] = request.BackendModelId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackendModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the definition of a backend service in an environment. After the definition is deleted, the API that uses the backend service and is published to this environment will be unpublished.
//
// @param request - DeleteBackendModelRequest
//
// @return DeleteBackendModelResponse
func (client *Client) DeleteBackendModel(request *DeleteBackendModelRequest) (_result *DeleteBackendModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBackendModelResponse{}
	_body, _err := client.DeleteBackendModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义数据集
//
// @param request - DeleteDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithOptions(request *DeleteDatasetRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义数据集
//
// @param request - DeleteDatasetRequest
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data entry from a custom dataset.
//
// @param request - DeleteDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetItemResponse
func (client *Client) DeleteDatasetItemWithOptions(request *DeleteDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *DeleteDatasetItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data entry from a custom dataset.
//
// @param request - DeleteDatasetItemRequest
//
// @return DeleteDatasetItemResponse
func (client *Client) DeleteDatasetItem(request *DeleteDatasetItemRequest) (_result *DeleteDatasetItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDatasetItemResponse{}
	_body, _err := client.DeleteDatasetItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a custom domain name from an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the specified domain name does not exist, a successful response will still appear.
//
//		- Unbinding a domain name from an API group will affect access to the APIs in the group. Exercise caution when using this operation.
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
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2016-07-14"),
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
// Unbinds a custom domain name from an API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the specified domain name does not exist, a successful response will still appear.
//
//		- Unbinding a domain name from an API group will affect access to the APIs in the group. Exercise caution when using this operation.
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
// Deletes the SSL certificate of a specified domain name. This operation is intended for API providers. If the SSL certificate does not exist, a success response is still returned. If the specified API group does not exist, the InvalidGroupId.NotFound error is returned. Access over HTTPS is not supported after the SSL certificate is deleted. Exercise caution when using this API operation.
//
// @param request - DeleteDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainCertificateResponse
func (client *Client) DeleteDomainCertificateWithOptions(request *DeleteDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertificateId) {
		query["CertificateId"] = request.CertificateId
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainCertificate"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the SSL certificate of a specified domain name. This operation is intended for API providers. If the SSL certificate does not exist, a success response is still returned. If the specified API group does not exist, the InvalidGroupId.NotFound error is returned. Access over HTTPS is not supported after the SSL certificate is deleted. Exercise caution when using this API operation.
//
// @param request - DeleteDomainCertificateRequest
//
// @return DeleteDomainCertificateResponse
func (client *Client) DeleteDomainCertificate(request *DeleteDomainCertificateRequest) (_result *DeleteDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainCertificateResponse{}
	_body, _err := client.DeleteDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API Gateway instance.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API Gateway instance.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the ACL is bound to an API, you must unbind the ACL from the API before you can delete the ACL. Otherwise, an error is returned.
//
//		- If you call this operation on an ACL that does not exist, a success message is returned.
//
// @param request - DeleteIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpControlResponse
func (client *Client) DeleteIpControlWithOptions(request *DeleteIpControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- If the ACL is bound to an API, you must unbind the ACL from the API before you can delete the ACL. Otherwise, an error is returned.
//
//		- If you call this operation on an ACL that does not exist, a success message is returned.
//
// @param request - DeleteIpControlRequest
//
// @return DeleteIpControlResponse
func (client *Client) DeleteIpControl(request *DeleteIpControlRequest) (_result *DeleteIpControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpControlResponse{}
	_body, _err := client.DeleteIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the specified log configuration.
//
// @param request - DeleteLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogConfigResponse
func (client *Client) DeleteLogConfigWithOptions(request *DeleteLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the specified log configuration.
//
// @param request - DeleteLogConfigRequest
//
// @return DeleteLogConfigResponse
func (client *Client) DeleteLogConfig(request *DeleteLogConfigRequest) (_result *DeleteLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogConfigResponse{}
	_body, _err := client.DeleteLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @param request - DeleteModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model.
//
// @param request - DeleteModelRequest
//
// @return DeleteModelResponse
func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a CloudMonitor application group corresponding to an API group.
//
// @param request - DeleteMonitorGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroupWithOptions(request *DeleteMonitorGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMonitorGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.RawMonitorGroupId) {
		query["RawMonitorGroupId"] = request.RawMonitorGroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMonitorGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a CloudMonitor application group corresponding to an API group.
//
// @param request - DeleteMonitorGroupRequest
//
// @return DeleteMonitorGroupResponse
func (client *Client) DeleteMonitorGroup(request *DeleteMonitorGroupRequest) (_result *DeleteMonitorGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMonitorGroupResponse{}
	_body, _err := client.DeleteMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must first unbind the plug-in from the API. Otherwise, an error is reported when you delete the plug-in.
//
// @param request - DeletePluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginResponse
func (client *Client) DeletePluginWithOptions(request *DeletePluginRequest, runtime *dara.RuntimeOptions) (_result *DeletePluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must first unbind the plug-in from the API. Otherwise, an error is reported when you delete the plug-in.
//
// @param request - DeletePluginRequest
//
// @return DeletePluginResponse
func (client *Client) DeletePlugin(request *DeletePluginRequest) (_result *DeletePluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePluginResponse{}
	_body, _err := client.DeletePluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an internal domain name resolution.
//
// @param request - DeletePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateDNSResponse
func (client *Client) DeletePrivateDNSWithOptions(request *DeletePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateDNSResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal domain name resolution.
//
// @param request - DeletePrivateDNSRequest
//
// @return DeletePrivateDNSResponse
func (client *Client) DeletePrivateDNS(request *DeletePrivateDNSRequest) (_result *DeletePrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateDNSResponse{}
	_body, _err := client.DeletePrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation deletes an existing backend signature key.
//
//		- You cannot delete a key that is bound to an API. To delete the key, you must unbind it first.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteSignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSignatureResponse
func (client *Client) DeleteSignatureWithOptions(request *DeleteSignatureRequest, runtime *dara.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation deletes an existing backend signature key.
//
//		- You cannot delete a key that is bound to an API. To delete the key, you must unbind it first.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteSignatureRequest
//
// @return DeleteSignatureResponse
func (client *Client) DeleteSignature(request *DeleteSignatureRequest) (_result *DeleteSignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DeleteSignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom throttling policy and the special throttling rules in the policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the throttling policy you want to delete is bound to APIs, you need to unbind the policy first. Otherwise, an error is reported when you delete the policy.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficControlResponse
func (client *Client) DeleteTrafficControlWithOptions(request *DeleteTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom throttling policy and the special throttling rules in the policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- If the throttling policy you want to delete is bound to APIs, you need to unbind the policy first. Otherwise, an error is reported when you delete the policy.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeleteTrafficControlRequest
//
// @return DeleteTrafficControlResponse
func (client *Client) DeleteTrafficControl(request *DeleteTrafficControlRequest) (_result *DeleteTrafficControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTrafficControlResponse{}
	_body, _err := client.DeleteTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom special throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can obtain the input parameters required in this operation by calling other APIs.
//
// @param request - DeleteTrafficSpecialControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTrafficSpecialControlResponse
func (client *Client) DeleteTrafficSpecialControlWithOptions(request *DeleteTrafficSpecialControlRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficSpecialControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecialKey) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !dara.IsNil(request.SpecialType) {
		query["SpecialType"] = request.SpecialType
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTrafficSpecialControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom special throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can obtain the input parameters required in this operation by calling other APIs.
//
// @param request - DeleteTrafficSpecialControlRequest
//
// @return DeleteTrafficSpecialControlResponse
func (client *Client) DeleteTrafficSpecialControl(request *DeleteTrafficSpecialControlRequest) (_result *DeleteTrafficSpecialControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTrafficSpecialControlResponse{}
	_body, _err := client.DeleteTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes an API to an environment.
//
// Description:
//
//	  This operation is intended for API providers. Only the API that you have defined and published to a runtime environment can be called.
//
//		- An API is published to a cluster in under 5 seconds.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeployApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApiResponse
func (client *Client) DeployApiWithOptions(request *DeployApiRequest, runtime *dara.RuntimeOptions) (_result *DeployApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an API to an environment.
//
// Description:
//
//	  This operation is intended for API providers. Only the API that you have defined and published to a runtime environment can be called.
//
//		- An API is published to a cluster in under 5 seconds.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - DeployApiRequest
//
// @return DeployApiResponse
func (client *Client) DeployApi(request *DeployApiRequest) (_result *DeployApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployApiResponse{}
	_body, _err := client.DeployApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询批量下线API任务
//
// @param request - DescribeAbolishApiTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAbolishApiTaskResponse
func (client *Client) DescribeAbolishApiTaskWithOptions(request *DescribeAbolishApiTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeAbolishApiTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAbolishApiTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAbolishApiTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量下线API任务
//
// @param request - DescribeAbolishApiTaskRequest
//
// @return DescribeAbolishApiTaskResponse
func (client *Client) DescribeAbolishApiTask(request *DescribeAbolishApiTaskRequest) (_result *DescribeAbolishApiTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAbolishApiTaskResponse{}
	_body, _err := client.DescribeAbolishApiTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries the details of an access control policy.
//
// @param request - DescribeAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListAttributeResponse
func (client *Client) DescribeAccessControlListAttributeWithOptions(request *DescribeAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlListAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessControlListAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries the details of an access control policy.
//
// @param request - DescribeAccessControlListAttributeRequest
//
// @return DescribeAccessControlListAttributeResponse
func (client *Client) DescribeAccessControlListAttribute(request *DescribeAccessControlListAttributeRequest) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessControlListAttributeResponse{}
	_body, _err := client.DescribeAccessControlListAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries access control policies.
//
// @param request - DescribeAccessControlListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListsResponse
func (client *Client) DescribeAccessControlListsWithOptions(request *DescribeAccessControlListsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlLists"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessControlListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control for dedicated instances. Queries access control policies.
//
// @param request - DescribeAccessControlListsRequest
//
// @return DescribeAccessControlListsResponse
func (client *Client) DescribeAccessControlLists(request *DescribeAccessControlListsRequest) (_result *DescribeAccessControlListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessControlListsResponse{}
	_body, _err := client.DescribeAccessControlListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the definition of an API.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiResponse
func (client *Client) DescribeApiWithOptions(request *DescribeApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the definition of an API.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiRequest
//
// @return DescribeApiResponse
func (client *Client) DescribeApi(request *DescribeApiRequest) (_result *DescribeApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiResponse{}
	_body, _err := client.DescribeApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the documentation of an API.
//
// Description:
//
//	  For API callers, the specified API must be a public or authorized private API that has been published to a runtime environment.
//
//		- When you call this operation as an API caller, the service information, parameter definitions, and other details of the API you specify are returned.
//
//		- When you call this operation as an API provider, the definition of the specified API running in the specified runtime environment is returned. The returned definition takes effect in the runtime environment, and may be different from the definition of the API you modify.
//
//		- Before you call this operation as an API provider, ensure that the API to be queried is a public one or that your application has been authorized to call the API, because authentication on API callers is required.
//
// @param request - DescribeApiDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiDocResponse
func (client *Client) DescribeApiDocWithOptions(request *DescribeApiDocRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiDocResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiDoc"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the documentation of an API.
//
// Description:
//
//	  For API callers, the specified API must be a public or authorized private API that has been published to a runtime environment.
//
//		- When you call this operation as an API caller, the service information, parameter definitions, and other details of the API you specify are returned.
//
//		- When you call this operation as an API provider, the definition of the specified API running in the specified runtime environment is returned. The returned definition takes effect in the runtime environment, and may be different from the definition of the API you modify.
//
//		- Before you call this operation as an API provider, ensure that the API to be queried is a public one or that your application has been authorized to call the API, because authentication on API callers is required.
//
// @param request - DescribeApiDocRequest
//
// @return DescribeApiDocResponse
func (client *Client) DescribeApiDoc(request *DescribeApiDocRequest) (_result *DescribeApiDocResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiDocResponse{}
	_body, _err := client.DescribeApiDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query details about an API group, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupResponse
func (client *Client) DescribeApiGroupWithOptions(request *DescribeApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query details about an API group, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupRequest
//
// @return DescribeApiGroupResponse
func (client *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (_result *DescribeApiGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiGroupResponse{}
	_body, _err := client.DescribeApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the VPC whitelist that is allowed to access an API group.
//
// @param request - DescribeApiGroupVpcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupVpcWhitelistResponse
func (client *Client) DescribeApiGroupVpcWhitelistWithOptions(request *DescribeApiGroupVpcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupVpcWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroupVpcWhitelist"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the VPC whitelist that is allowed to access an API group.
//
// @param request - DescribeApiGroupVpcWhitelistRequest
//
// @return DescribeApiGroupVpcWhitelistResponse
func (client *Client) DescribeApiGroupVpcWhitelist(request *DescribeApiGroupVpcWhitelistRequest) (_result *DescribeApiGroupVpcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiGroupVpcWhitelistResponse{}
	_body, _err := client.DescribeApiGroupVpcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries existing API groups and their basic information.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiGroupsResponse
func (client *Client) DescribeApiGroupsWithOptions(request *DescribeApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries existing API groups and their basic information.
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - DescribeApiGroupsRequest
//
// @return DescribeApiGroupsResponse
func (client *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (_result *DescribeApiGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.DescribeApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical versions of a specified API.
//
// Description:
//
//	  This operation is intended for API providers. Only APIs that have been published have historical version records.
//
//		- This operation allows you to obtain the historical versions of an API. This operation is always called by other operations.
//
// @param request - DescribeApiHistoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiHistoriesResponse
func (client *Client) DescribeApiHistoriesWithOptions(request *DescribeApiHistoriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiHistoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiHistories"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical versions of a specified API.
//
// Description:
//
//	  This operation is intended for API providers. Only APIs that have been published have historical version records.
//
//		- This operation allows you to obtain the historical versions of an API. This operation is always called by other operations.
//
// @param request - DescribeApiHistoriesRequest
//
// @return DescribeApiHistoriesResponse
func (client *Client) DescribeApiHistories(request *DescribeApiHistoriesRequest) (_result *DescribeApiHistoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiHistoriesResponse{}
	_body, _err := client.DescribeApiHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified historical version of a specified API definition.
//
// Description:
//
// Queries the details of a specified historical version of a specified API definition.
//
//   - This API is intended for API providers.
//
//   - API Gateway records the time and definition of an API every time the API is published. You can use the version number obtained from other operations to query definition details at a certain publication.
//
// @param request - DescribeApiHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiHistoryResponse
func (client *Client) DescribeApiHistoryWithOptions(request *DescribeApiHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryVersion) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiHistory"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified historical version of a specified API definition.
//
// Description:
//
// Queries the details of a specified historical version of a specified API definition.
//
//   - This API is intended for API providers.
//
//   - API Gateway records the time and definition of an API every time the API is published. You can use the version number obtained from other operations to query definition details at a certain publication.
//
// @param request - DescribeApiHistoryRequest
//
// @return DescribeApiHistoryResponse
func (client *Client) DescribeApiHistory(request *DescribeApiHistoryRequest) (_result *DescribeApiHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiHistoryResponse{}
	_body, _err := client.DescribeApiHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) that are bound to all the APIs in an API group in a specified environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- If an optional parameter is not specified, all results are returned on separate pages.
//
// ·
//
// @param request - DescribeApiIpControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiIpControlsResponse
func (client *Client) DescribeApiIpControlsWithOptions(request *DescribeApiIpControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiIpControlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiIpControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiIpControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) that are bound to all the APIs in an API group in a specified environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- If an optional parameter is not specified, all results are returned on separate pages.
//
// ·
//
// @param request - DescribeApiIpControlsRequest
//
// @return DescribeApiIpControlsResponse
func (client *Client) DescribeApiIpControls(request *DescribeApiIpControlsRequest) (_result *DescribeApiIpControlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiIpControlsResponse{}
	_body, _err := client.DescribeApiIpControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the response time statistics of an API.
//
// Description:
//
// You can call this operation to query the latency metrics in milliseconds for a specified API.
//
//   - This API is intended for API providers.
//
//   - Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiLatencyDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiLatencyDataResponse
func (client *Client) DescribeApiLatencyDataWithOptions(request *DescribeApiLatencyDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiLatencyDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiLatencyData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiLatencyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the response time statistics of an API.
//
// Description:
//
// You can call this operation to query the latency metrics in milliseconds for a specified API.
//
//   - This API is intended for API providers.
//
//   - Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiLatencyDataRequest
//
// @return DescribeApiLatencyDataResponse
func (client *Client) DescribeApiLatencyData(request *DescribeApiLatencyDataRequest) (_result *DescribeApiLatencyDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiLatencyDataResponse{}
	_body, _err := client.DescribeApiLatencyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud Marketplace attributes of an API.
//
// @param request - DescribeApiMarketAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiMarketAttributesResponse
func (client *Client) DescribeApiMarketAttributesWithOptions(request *DescribeApiMarketAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiMarketAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiMarketAttributes"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiMarketAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud Marketplace attributes of an API.
//
// @param request - DescribeApiMarketAttributesRequest
//
// @return DescribeApiMarketAttributesResponse
func (client *Client) DescribeApiMarketAttributes(request *DescribeApiMarketAttributesRequest) (_result *DescribeApiMarketAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiMarketAttributesResponse{}
	_body, _err := client.DescribeApiMarketAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attached APIs of an API product.
//
// @param request - DescribeApiProductApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiProductApisResponse
func (client *Client) DescribeApiProductApisWithOptions(request *DescribeApiProductApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiProductApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiProductApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiProductApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attached APIs of an API product.
//
// @param request - DescribeApiProductApisRequest
//
// @return DescribeApiProductApisResponse
func (client *Client) DescribeApiProductApis(request *DescribeApiProductApisRequest) (_result *DescribeApiProductApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiProductApisResponse{}
	_body, _err := client.DescribeApiProductApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries API products by application.
//
// @param request - DescribeApiProductsByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiProductsByAppResponse
func (client *Client) DescribeApiProductsByAppWithOptions(request *DescribeApiProductsByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiProductsByAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiProductsByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiProductsByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API products by application.
//
// @param request - DescribeApiProductsByAppRequest
//
// @return DescribeApiProductsByAppResponse
func (client *Client) DescribeApiProductsByApp(request *DescribeApiProductsByAppRequest) (_result *DescribeApiProductsByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiProductsByAppResponse{}
	_body, _err := client.DescribeApiProductsByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the QPS statistics of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiQpsDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiQpsDataResponse
func (client *Client) DescribeApiQpsDataWithOptions(request *DescribeApiQpsDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiQpsDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiQpsData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the QPS statistics of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiQpsDataRequest
//
// @return DescribeApiQpsDataResponse
func (client *Client) DescribeApiQpsData(request *DescribeApiQpsDataRequest) (_result *DescribeApiQpsDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiQpsDataResponse{}
	_body, _err := client.DescribeApiQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to the APIs of a specified API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiSignaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiSignaturesResponse
func (client *Client) DescribeApiSignaturesWithOptions(request *DescribeApiSignaturesRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiSignaturesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiSignatures"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to the APIs of a specified API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiSignaturesRequest
//
// @return DescribeApiSignaturesResponse
func (client *Client) DescribeApiSignatures(request *DescribeApiSignaturesRequest) (_result *DescribeApiSignaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiSignaturesResponse{}
	_body, _err := client.DescribeApiSignaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the throttling policies bound to all members of an API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiTrafficControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiTrafficControlsResponse
func (client *Client) DescribeApiTrafficControlsWithOptions(request *DescribeApiTrafficControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiTrafficControlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiTrafficControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiTrafficControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling policies bound to all members of an API group in a specified environment.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The ApiIds parameter is optional. If this parameter is not specified, all results in the specified environment of an API group are returned.
//
// @param request - DescribeApiTrafficControlsRequest
//
// @return DescribeApiTrafficControlsResponse
func (client *Client) DescribeApiTrafficControls(request *DescribeApiTrafficControlsRequest) (_result *DescribeApiTrafficControlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiTrafficControlsResponse{}
	_body, _err := client.DescribeApiTrafficControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the traffic of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiTrafficDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiTrafficDataResponse
func (client *Client) DescribeApiTrafficDataWithOptions(request *DescribeApiTrafficDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeApiTrafficDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiTrafficData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the traffic of an API.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Only statistics for API calls made in the release environment are collected by default.
//
// @param request - DescribeApiTrafficDataRequest
//
// @return DescribeApiTrafficDataResponse
func (client *Client) DescribeApiTrafficData(request *DescribeApiTrafficDataRequest) (_result *DescribeApiTrafficDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApiTrafficDataResponse{}
	_body, _err := client.DescribeApiTrafficDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of APIs that are being defined.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation returns a list of all APIs that are being defined. The basic information about these APIs is also returned in the list.
//
//		- This operation returns all APIs that are being edited, regardless of their environments. The returned definitions may be different from the definitions in the environments.
//
// @param request - DescribeApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisResponse
func (client *Client) DescribeApisWithOptions(request *DescribeApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiMethod) {
		query["ApiMethod"] = request.ApiMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiPath) {
		query["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.CatalogId) {
		query["CatalogId"] = request.CatalogId
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnDeployed) {
		query["UnDeployed"] = request.UnDeployed
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of APIs that are being defined.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation returns a list of all APIs that are being defined. The basic information about these APIs is also returned in the list.
//
//		- This operation returns all APIs that are being edited, regardless of their environments. The returned definitions may be different from the definitions in the environments.
//
// @param request - DescribeApisRequest
//
// @return DescribeApisResponse
func (client *Client) DescribeApis(request *DescribeApisRequest) (_result *DescribeApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisResponse{}
	_body, _err := client.DescribeApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs with which an application is associated.
//
// @param request - DescribeApisByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByAppResponse
func (client *Client) DescribeApisByAppWithOptions(request *DescribeApisByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiUid) {
		query["ApiUid"] = request.ApiUid
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs with which an application is associated.
//
// @param request - DescribeApisByAppRequest
//
// @return DescribeApisByAppResponse
func (client *Client) DescribeApisByApp(request *DescribeApisByAppRequest) (_result *DescribeApisByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisByAppResponse{}
	_body, _err := client.DescribeApisByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries APIs in the draft or published state that are created by using a specified backend service.
//
// @param request - DescribeApisByBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByBackendResponse
func (client *Client) DescribeApisByBackendWithOptions(request *DescribeApisByBackendRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByBackendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs in the draft or published state that are created by using a specified backend service.
//
// @param request - DescribeApisByBackendRequest
//
// @return DescribeApisByBackendResponse
func (client *Client) DescribeApisByBackend(request *DescribeApisByBackendRequest) (_result *DescribeApisByBackendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisByBackendResponse{}
	_body, _err := client.DescribeApisByBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs that are bound to an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByIpControlResponse
func (client *Client) DescribeApisByIpControlWithOptions(request *DescribeApisByIpControlRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByIpControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that are bound to an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByIpControlRequest
//
// @return DescribeApisByIpControlResponse
func (client *Client) DescribeApisByIpControl(request *DescribeApisByIpControlRequest) (_result *DescribeApisByIpControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisByIpControlResponse{}
	_body, _err := client.DescribeApisByIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified backend signature key is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The results are returned on separate pages. You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisBySignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisBySignatureResponse
func (client *Client) DescribeApisBySignatureWithOptions(request *DescribeApisBySignatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisBySignatureResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisBySignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisBySignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified backend signature key is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The results are returned on separate pages. You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisBySignatureRequest
//
// @return DescribeApisBySignatureResponse
func (client *Client) DescribeApisBySignature(request *DescribeApisBySignatureRequest) (_result *DescribeApisBySignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisBySignatureResponse{}
	_body, _err := client.DescribeApisBySignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified throttling policy is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByTrafficControlResponse
func (client *Client) DescribeApisByTrafficControlWithOptions(request *DescribeApisByTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByTrafficControlResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified throttling policy is bound.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- You can specify PageNumber to obtain the result on the specified page.
//
// @param request - DescribeApisByTrafficControlRequest
//
// @return DescribeApisByTrafficControlResponse
func (client *Client) DescribeApisByTrafficControl(request *DescribeApisByTrafficControlRequest) (_result *DescribeApisByTrafficControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisByTrafficControlResponse{}
	_body, _err := client.DescribeApisByTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs that are associated with a virtual private cloud (VPC) access authorization in a region.
//
// @param request - DescribeApisByVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisByVpcAccessResponse
func (client *Client) DescribeApisByVpcAccessWithOptions(request *DescribeApisByVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisByVpcAccessResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcName) {
		query["VpcName"] = request.VpcName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisByVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisByVpcAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that are associated with a virtual private cloud (VPC) access authorization in a region.
//
// @param request - DescribeApisByVpcAccessRequest
//
// @return DescribeApisByVpcAccessResponse
func (client *Client) DescribeApisByVpcAccess(request *DescribeApisByVpcAccessRequest) (_result *DescribeApisByVpcAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisByVpcAccessResponse{}
	_body, _err := client.DescribeApisByVpcAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries APIs by application. The environment information is also returned.
//
// @param request - DescribeApisWithStageNameIntegratedByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApisWithStageNameIntegratedByAppResponse
func (client *Client) DescribeApisWithStageNameIntegratedByAppWithOptions(request *DescribeApisWithStageNameIntegratedByAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeApisWithStageNameIntegratedByAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiUid) {
		query["ApiUid"] = request.ApiUid
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApisWithStageNameIntegratedByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApisWithStageNameIntegratedByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs by application. The environment information is also returned.
//
// @param request - DescribeApisWithStageNameIntegratedByAppRequest
//
// @return DescribeApisWithStageNameIntegratedByAppResponse
func (client *Client) DescribeApisWithStageNameIntegratedByApp(request *DescribeApisWithStageNameIntegratedByAppRequest) (_result *DescribeApisWithStageNameIntegratedByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeApisWithStageNameIntegratedByAppResponse{}
	_body, _err := client.DescribeApisWithStageNameIntegratedByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the apps that can be authorized.
//
// @param request - DescribeAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppResponse
func (client *Client) DescribeAppWithOptions(request *DescribeAppRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the apps that can be authorized.
//
// @param request - DescribeAppRequest
//
// @return DescribeAppResponse
func (client *Client) DescribeApp(request *DescribeAppRequest) (_result *DescribeAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppResponse{}
	_body, _err := client.DescribeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries apps and their basic information.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppId is optional.
//
// @param request - DescribeAppAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppAttributesResponse
func (client *Client) DescribeAppAttributesWithOptions(request *DescribeAppAttributesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppAttributesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppAttributes"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries apps and their basic information.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppId is optional.
//
// @param request - DescribeAppAttributesRequest
//
// @return DescribeAppAttributesResponse
func (client *Client) DescribeAppAttributes(request *DescribeAppAttributesRequest) (_result *DescribeAppAttributesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppAttributesResponse{}
	_body, _err := client.DescribeAppAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the key-related information of an application.
//
// @param request - DescribeAppSecuritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppSecuritiesResponse
func (client *Client) DescribeAppSecuritiesWithOptions(request *DescribeAppSecuritiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppSecuritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppSecurities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppSecuritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the key-related information of an application.
//
// @param request - DescribeAppSecuritiesRequest
//
// @return DescribeAppSecuritiesResponse
func (client *Client) DescribeAppSecurities(request *DescribeAppSecuritiesRequest) (_result *DescribeAppSecuritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppSecuritiesResponse{}
	_body, _err := client.DescribeAppSecuritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This key is used for authentication when an API call is made.
//
// Description:
//
//	This operation is intended for API callers.
//
// @param request - DescribeAppSecurityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppSecurityResponse
func (client *Client) DescribeAppSecurityWithOptions(request *DescribeAppSecurityRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppSecurityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppSecurity"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppSecurityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This key is used for authentication when an API call is made.
//
// Description:
//
//	This operation is intended for API callers.
//
// @param request - DescribeAppSecurityRequest
//
// @return DescribeAppSecurityResponse
func (client *Client) DescribeAppSecurity(request *DescribeAppSecurityRequest) (_result *DescribeAppSecurityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppSecurityResponse{}
	_body, _err := client.DescribeAppSecurityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the apps of a user. App information is returned only to the app owner.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- API providers can use the app IDs or their Alibaba Cloud accounts to query app information.
//
//		- Each provider can call this operation for a maximum of 200 times every day in a region.
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppOwner) {
		query["AppOwner"] = request.AppOwner
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the apps of a user. App information is returned only to the app owner.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- API providers can use the app IDs or their Alibaba Cloud accounts to query app information.
//
//		- Each provider can call this operation for a maximum of 200 times every day in a region.
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries authorized applications by API product.
//
// @param request - DescribeAppsByApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsByApiProductResponse
func (client *Client) DescribeAppsByApiProductWithOptions(request *DescribeAppsByApiProductRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppsByApiProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppsByApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppsByApiProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authorized applications by API product.
//
// @param request - DescribeAppsByApiProductRequest
//
// @return DescribeAppsByApiProductResponse
func (client *Client) DescribeAppsByApiProduct(request *DescribeAppsByApiProductRequest) (_result *DescribeAppsByApiProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppsByApiProductResponse{}
	_body, _err := client.DescribeAppsByApiProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorized APIs of a specified APP.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The specified application can call all APIs included in the responses.
//
// @param request - DescribeAuthorizedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthorizedApisResponse
func (client *Client) DescribeAuthorizedApisWithOptions(request *DescribeAuthorizedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthorizedApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthorizedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthorizedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorized APIs of a specified APP.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The specified application can call all APIs included in the responses.
//
// @param request - DescribeAuthorizedApisRequest
//
// @return DescribeAuthorizedApisResponse
func (client *Client) DescribeAuthorizedApis(request *DescribeAuthorizedApisRequest) (_result *DescribeAuthorizedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthorizedApisResponse{}
	_body, _err := client.DescribeAuthorizedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the current apps.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- All applications included in the responses have access to the specified API.
//
// @param request - DescribeAuthorizedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAuthorizedAppsResponse
func (client *Client) DescribeAuthorizedAppsWithOptions(request *DescribeAuthorizedAppsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAuthorizedAppsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppOwnerId) {
		query["AppOwnerId"] = request.AppOwnerId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAuthorizedApps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAuthorizedAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the current apps.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- All applications included in the responses have access to the specified API.
//
// @param request - DescribeAuthorizedAppsRequest
//
// @return DescribeAuthorizedAppsResponse
func (client *Client) DescribeAuthorizedApps(request *DescribeAuthorizedAppsRequest) (_result *DescribeAuthorizedAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAuthorizedAppsResponse{}
	_body, _err := client.DescribeAuthorizedAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a backend service and its URL configured for each environment.
//
// @param request - DescribeBackendInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackendInfoResponse
func (client *Client) DescribeBackendInfoWithOptions(request *DescribeBackendInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackendInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackendInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackendInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a backend service and its URL configured for each environment.
//
// @param request - DescribeBackendInfoRequest
//
// @return DescribeBackendInfoResponse
func (client *Client) DescribeBackendInfo(request *DescribeBackendInfoRequest) (_result *DescribeBackendInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackendInfoResponse{}
	_body, _err := client.DescribeBackendInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backend services. You can filter backend services by backend service name and backend service type.
//
// @param request - DescribeBackendListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackendListResponse
func (client *Client) DescribeBackendListWithOptions(request *DescribeBackendListRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackendListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackendList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackendListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backend services. You can filter backend services by backend service name and backend service type.
//
// @param request - DescribeBackendListRequest
//
// @return DescribeBackendListResponse
func (client *Client) DescribeBackendList(request *DescribeBackendListRequest) (_result *DescribeBackendListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackendListResponse{}
	_body, _err := client.DescribeBackendListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a single dataset.
//
// @param request - DescribeDatasetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetInfoResponse
func (client *Client) DescribeDatasetInfoWithOptions(request *DescribeDatasetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a single dataset.
//
// @param request - DescribeDatasetInfoRequest
//
// @return DescribeDatasetInfoResponse
func (client *Client) DescribeDatasetInfo(request *DescribeDatasetInfoRequest) (_result *DescribeDatasetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatasetInfoResponse{}
	_body, _err := client.DescribeDatasetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a data entry in a custom dataset.
//
// @param request - DescribeDatasetItemInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetItemInfoResponse
func (client *Client) DescribeDatasetItemInfoWithOptions(request *DescribeDatasetItemInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetItemInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
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
		Action:      dara.String("DescribeDatasetItemInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetItemInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a data entry in a custom dataset.
//
// @param request - DescribeDatasetItemInfoRequest
//
// @return DescribeDatasetItemInfoResponse
func (client *Client) DescribeDatasetItemInfo(request *DescribeDatasetItemInfoRequest) (_result *DescribeDatasetItemInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatasetItemInfoResponse{}
	_body, _err := client.DescribeDatasetItemInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data entries of a custom dataset.
//
// @param request - DescribeDatasetItemListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetItemListResponse
func (client *Client) DescribeDatasetItemListWithOptions(request *DescribeDatasetItemListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetItemListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemIds) {
		query["DatasetItemIds"] = request.DatasetItemIds
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetItemList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetItemListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data entries of a custom dataset.
//
// @param request - DescribeDatasetItemListRequest
//
// @return DescribeDatasetItemListResponse
func (client *Client) DescribeDatasetItemList(request *DescribeDatasetItemListRequest) (_result *DescribeDatasetItemListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatasetItemListResponse{}
	_body, _err := client.DescribeDatasetItemListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom datasets.
//
// @param request - DescribeDatasetListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatasetListResponse
func (client *Client) DescribeDatasetListWithOptions(request *DescribeDatasetListRequest, runtime *dara.RuntimeOptions) (_result *DescribeDatasetListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetIds) {
		query["DatasetIds"] = request.DatasetIds
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDatasetList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDatasetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom datasets.
//
// @param request - DescribeDatasetListRequest
//
// @return DescribeDatasetListResponse
func (client *Client) DescribeDatasetList(request *DescribeDatasetListRequest) (_result *DescribeDatasetListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDatasetListResponse{}
	_body, _err := client.DescribeDatasetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress of an asynchronous API publishing task.
//
// @param request - DescribeDeployApiTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployApiTaskResponse
func (client *Client) DescribeDeployApiTaskWithOptions(request *DescribeDeployApiTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployApiTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployApiTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployApiTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of an asynchronous API publishing task.
//
// @param request - DescribeDeployApiTaskRequest
//
// @return DescribeDeployApiTaskResponse
func (client *Client) DescribeDeployApiTask(request *DescribeDeployApiTaskRequest) (_result *DescribeDeployApiTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeployApiTaskResponse{}
	_body, _err := client.DescribeDeployApiTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the definition of an API that takes effect in an environment. The definition may differ from the definition being edited.
//
// @param request - DescribeDeployedApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployedApiResponse
func (client *Client) DescribeDeployedApiWithOptions(request *DescribeDeployedApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployedApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployedApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployedApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the definition of an API that takes effect in an environment. The definition may differ from the definition being edited.
//
// @param request - DescribeDeployedApiRequest
//
// @return DescribeDeployedApiResponse
func (client *Client) DescribeDeployedApi(request *DescribeDeployedApiRequest) (_result *DescribeDeployedApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeployedApiResponse{}
	_body, _err := client.DescribeDeployedApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs that have been published to a specified environment.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeDeployedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeployedApisResponse
func (client *Client) DescribeDeployedApisWithOptions(request *DescribeDeployedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeployedApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiMethod) {
		query["ApiMethod"] = request.ApiMethod
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiPath) {
		query["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.EnableTagAuth) {
		query["EnableTagAuth"] = request.EnableTagAuth
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeployedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeployedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that have been published to a specified environment.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeDeployedApisRequest
//
// @return DescribeDeployedApisResponse
func (client *Client) DescribeDeployedApis(request *DescribeDeployedApisRequest) (_result *DescribeDeployedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeployedApisResponse{}
	_body, _err := client.DescribeDeployedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about a bound custom domain name, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// @param request - DescribeDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainResponse
func (client *Client) DescribeDomainWithOptions(request *DescribeDomainRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about a bound custom domain name, including the automatically assigned second-level domain name, custom domain name, and SSL certificate.
//
// @param request - DescribeDomainRequest
//
// @return DescribeDomainResponse
func (client *Client) DescribeDomain(request *DescribeDomainRequest) (_result *DescribeDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainResponse{}
	_body, _err := client.DescribeDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the average latency of an API group in an environment.
//
// @param request - DescribeGroupLatencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupLatencyResponse
func (client *Client) DescribeGroupLatencyWithOptions(request *DescribeGroupLatencyRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupLatencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupLatency"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupLatencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average latency of an API group in an environment.
//
// @param request - DescribeGroupLatencyRequest
//
// @return DescribeGroupLatencyResponse
func (client *Client) DescribeGroupLatency(request *DescribeGroupLatencyRequest) (_result *DescribeGroupLatencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupLatencyResponse{}
	_body, _err := client.DescribeGroupLatencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of requests directed to an API group within a period of time.
//
// @param request - DescribeGroupQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupQpsResponse
func (client *Client) DescribeGroupQpsWithOptions(request *DescribeGroupQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupQpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupQps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of requests directed to an API group within a period of time.
//
// @param request - DescribeGroupQpsRequest
//
// @return DescribeGroupQpsResponse
func (client *Client) DescribeGroupQps(request *DescribeGroupQpsRequest) (_result *DescribeGroupQpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupQpsResponse{}
	_body, _err := client.DescribeGroupQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the traffic of an API group.
//
// @param request - DescribeGroupTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupTrafficResponse
func (client *Client) DescribeGroupTrafficWithOptions(request *DescribeGroupTrafficRequest, runtime *dara.RuntimeOptions) (_result *DescribeGroupTrafficResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGroupTraffic"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGroupTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic of an API group.
//
// @param request - DescribeGroupTrafficRequest
//
// @return DescribeGroupTrafficResponse
func (client *Client) DescribeGroupTraffic(request *DescribeGroupTrafficRequest) (_result *DescribeGroupTrafficResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGroupTrafficResponse{}
	_body, _err := client.DescribeGroupTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical versions of an API.
//
// @param request - DescribeHistoryApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHistoryApisResponse
func (client *Client) DescribeHistoryApisWithOptions(request *DescribeHistoryApisRequest, runtime *dara.RuntimeOptions) (_result *DescribeHistoryApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHistoryApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHistoryApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical versions of an API.
//
// @param request - DescribeHistoryApisRequest
//
// @return DescribeHistoryApisResponse
func (client *Client) DescribeHistoryApis(request *DescribeHistoryApisRequest) (_result *DescribeHistoryApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHistoryApisResponse{}
	_body, _err := client.DescribeHistoryApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an OAS API import task.
//
// @param request - DescribeImportOASTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImportOASTaskResponse
func (client *Client) DescribeImportOASTaskWithOptions(request *DescribeImportOASTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeImportOASTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImportOASTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImportOASTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an OAS API import task.
//
// @param request - DescribeImportOASTaskRequest
//
// @return DescribeImportOASTaskResponse
func (client *Client) DescribeImportOASTask(request *DescribeImportOASTaskRequest) (_result *DescribeImportOASTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImportOASTaskResponse{}
	_body, _err := client.DescribeImportOASTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a dedicated instance cluster.
//
// @param request - DescribeInstanceClusterInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceClusterInfoResponse
func (client *Client) DescribeInstanceClusterInfoWithOptions(request *DescribeInstanceClusterInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceClusterInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceClusterInfo"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceClusterInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a dedicated instance cluster.
//
// @param request - DescribeInstanceClusterInfoRequest
//
// @return DescribeInstanceClusterInfoResponse
func (client *Client) DescribeInstanceClusterInfo(request *DescribeInstanceClusterInfoRequest) (_result *DescribeInstanceClusterInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceClusterInfoResponse{}
	_body, _err := client.DescribeInstanceClusterInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries dedicated instance clusters.
//
// @param request - DescribeInstanceClusterListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceClusterListResponse
func (client *Client) DescribeInstanceClusterListWithOptions(request *DescribeInstanceClusterListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceClusterListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceClusterId) {
		query["InstanceClusterId"] = request.InstanceClusterId
	}

	if !dara.IsNil(request.InstanceClusterName) {
		query["InstanceClusterName"] = request.InstanceClusterName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceClusterList"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceClusterListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries dedicated instance clusters.
//
// @param request - DescribeInstanceClusterListRequest
//
// @return DescribeInstanceClusterListResponse
func (client *Client) DescribeInstanceClusterList(request *DescribeInstanceClusterListRequest) (_result *DescribeInstanceClusterListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceClusterListResponse{}
	_body, _err := client.DescribeInstanceClusterListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of lost connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceDropConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDropConnectionsResponse
func (client *Client) DescribeInstanceDropConnectionsWithOptions(request *DescribeInstanceDropConnectionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDropConnectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDropConnections"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDropConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of lost connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceDropConnectionsRequest
//
// @return DescribeInstanceDropConnectionsResponse
func (client *Client) DescribeInstanceDropConnections(request *DescribeInstanceDropConnectionsRequest) (_result *DescribeInstanceDropConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDropConnectionsResponse{}
	_body, _err := client.DescribeInstanceDropConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of dropped packets within a period of time.
//
// @param request - DescribeInstanceDropPacketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDropPacketResponse
func (client *Client) DescribeInstanceDropPacketWithOptions(request *DescribeInstanceDropPacketRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDropPacketResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDropPacket"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDropPacketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of dropped packets within a period of time.
//
// @param request - DescribeInstanceDropPacketRequest
//
// @return DescribeInstanceDropPacketResponse
func (client *Client) DescribeInstanceDropPacket(request *DescribeInstanceDropPacketRequest) (_result *DescribeInstanceDropPacketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceDropPacketResponse{}
	_body, _err := client.DescribeInstanceDropPacketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceHttpCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceHttpCodeResponse
func (client *Client) DescribeInstanceHttpCodeWithOptions(request *DescribeInstanceHttpCodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceHttpCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceHttpCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceHttpCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of HTTP status codes of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceHttpCodeRequest
//
// @return DescribeInstanceHttpCodeResponse
func (client *Client) DescribeInstanceHttpCode(request *DescribeInstanceHttpCodeRequest) (_result *DescribeInstanceHttpCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceHttpCodeResponse{}
	_body, _err := client.DescribeInstanceHttpCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the average latency of a dedicated instance over a period of time.
//
// @param request - DescribeInstanceLatencyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceLatencyResponse
func (client *Client) DescribeInstanceLatencyWithOptions(request *DescribeInstanceLatencyRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceLatencyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceLatency"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceLatencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the average latency of a dedicated instance over a period of time.
//
// @param request - DescribeInstanceLatencyRequest
//
// @return DescribeInstanceLatencyResponse
func (client *Client) DescribeInstanceLatency(request *DescribeInstanceLatencyRequest) (_result *DescribeInstanceLatencyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceLatencyResponse{}
	_body, _err := client.DescribeInstanceLatencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of new connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceNewConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceNewConnectionsResponse
func (client *Client) DescribeInstanceNewConnectionsWithOptions(request *DescribeInstanceNewConnectionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceNewConnectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceNewConnections"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceNewConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of new connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceNewConnectionsRequest
//
// @return DescribeInstanceNewConnectionsResponse
func (client *Client) DescribeInstanceNewConnections(request *DescribeInstanceNewConnectionsRequest) (_result *DescribeInstanceNewConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceNewConnectionsResponse{}
	_body, _err := client.DescribeInstanceNewConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the numbers of inbound and outbound packets of a dedicated instance within a period of time.
//
// @param request - DescribeInstancePacketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancePacketsResponse
func (client *Client) DescribeInstancePacketsWithOptions(request *DescribeInstancePacketsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancePacketsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstancePackets"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancePacketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the numbers of inbound and outbound packets of a dedicated instance within a period of time.
//
// @param request - DescribeInstancePacketsRequest
//
// @return DescribeInstancePacketsResponse
func (client *Client) DescribeInstancePackets(request *DescribeInstancePacketsRequest) (_result *DescribeInstancePacketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancePacketsResponse{}
	_body, _err := client.DescribeInstancePacketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceQpsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceQpsResponse
func (client *Client) DescribeInstanceQpsWithOptions(request *DescribeInstanceQpsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceQpsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceQps"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of requests to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceQpsRequest
//
// @return DescribeInstanceQpsResponse
func (client *Client) DescribeInstanceQps(request *DescribeInstanceQpsRequest) (_result *DescribeInstanceQpsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceQpsResponse{}
	_body, _err := client.DescribeInstanceQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of concurrent connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceSlbConnectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSlbConnectResponse
func (client *Client) DescribeInstanceSlbConnectWithOptions(request *DescribeInstanceSlbConnectRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSlbConnectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SbcName) {
		query["SbcName"] = request.SbcName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceSlbConnect"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSlbConnectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of concurrent connections to a dedicated instance within a period of time.
//
// @param request - DescribeInstanceSlbConnectRequest
//
// @return DescribeInstanceSlbConnectResponse
func (client *Client) DescribeInstanceSlbConnect(request *DescribeInstanceSlbConnectRequest) (_result *DescribeInstanceSlbConnectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSlbConnectResponse{}
	_body, _err := client.DescribeInstanceSlbConnectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request traffic and response traffic of a dedicated instance within a period of time.
//
// @param request - DescribeInstanceTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceTrafficResponse
func (client *Client) DescribeInstanceTrafficWithOptions(request *DescribeInstanceTrafficRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceTrafficResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceTraffic"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request traffic and response traffic of a dedicated instance within a period of time.
//
// @param request - DescribeInstanceTrafficRequest
//
// @return DescribeInstanceTrafficResponse
func (client *Client) DescribeInstanceTraffic(request *DescribeInstanceTrafficRequest) (_result *DescribeInstanceTrafficResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceTrafficResponse{}
	_body, _err := client.DescribeInstanceTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of instances in a region. The instances include shared instances and dedicated instances.
//
// @param request - DescribeInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableTagAuthorization) {
		query["EnableTagAuthorization"] = request.EnableTagAuthorization
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of instances in a region. The instances include shared instances and dedicated instances.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rule entries of an IP address-based traffic control policy.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can filter the query results by policy ID.
//
// @param request - DescribeIpControlPolicyItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpControlPolicyItemsResponse
func (client *Client) DescribeIpControlPolicyItemsWithOptions(request *DescribeIpControlPolicyItemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpControlPolicyItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyItemId) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpControlPolicyItems"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpControlPolicyItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rule entries of an IP address-based traffic control policy.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You can filter the query results by policy ID.
//
// @param request - DescribeIpControlPolicyItemsRequest
//
// @return DescribeIpControlPolicyItemsResponse
func (client *Client) DescribeIpControlPolicyItems(request *DescribeIpControlPolicyItemsRequest) (_result *DescribeIpControlPolicyItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpControlPolicyItemsResponse{}
	_body, _err := client.DescribeIpControlPolicyItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom access control lists (ACLs) on separate pages.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation is used to query the ACLs in a region. Region is a system parameter.
//
//		- You can filter the query results by ACL ID, name, or type.
//
//		- This operation cannot be used to query specific policies. If you want to query specific policies, call the [DescribeIpControlPolicyItems](~~DescribeIpControlPolicyItems~~) operation.
//
// @param request - DescribeIpControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeIpControlsResponse
func (client *Client) DescribeIpControlsWithOptions(request *DescribeIpControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeIpControlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.IpControlType) {
		query["IpControlType"] = request.IpControlType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeIpControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeIpControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom access control lists (ACLs) on separate pages.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation is used to query the ACLs in a region. Region is a system parameter.
//
//		- You can filter the query results by ACL ID, name, or type.
//
//		- This operation cannot be used to query specific policies. If you want to query specific policies, call the [DescribeIpControlPolicyItems](~~DescribeIpControlPolicyItems~~) operation.
//
// @param request - DescribeIpControlsRequest
//
// @return DescribeIpControlsResponse
func (client *Client) DescribeIpControls(request *DescribeIpControlsRequest) (_result *DescribeIpControlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeIpControlsResponse{}
	_body, _err := client.DescribeIpControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询日志配置
//
// @param request - DescribeLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLogConfigResponse
func (client *Client) DescribeLogConfigWithOptions(request *DescribeLogConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询日志配置
//
// @param request - DescribeLogConfigRequest
//
// @return DescribeLogConfigResponse
func (client *Client) DescribeLogConfig(request *DescribeLogConfigRequest) (_result *DescribeLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLogConfigResponse{}
	_body, _err := client.DescribeLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of remaining ordered relationships for a purchaser.
//
// @param request - DescribeMarketRemainsQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMarketRemainsQuotaResponse
func (client *Client) DescribeMarketRemainsQuotaWithOptions(request *DescribeMarketRemainsQuotaRequest, runtime *dara.RuntimeOptions) (_result *DescribeMarketRemainsQuotaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMarketRemainsQuota"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMarketRemainsQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of remaining ordered relationships for a purchaser.
//
// @param request - DescribeMarketRemainsQuotaRequest
//
// @return DescribeMarketRemainsQuotaResponse
func (client *Client) DescribeMarketRemainsQuota(request *DescribeMarketRemainsQuotaRequest) (_result *DescribeMarketRemainsQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMarketRemainsQuotaResponse{}
	_body, _err := client.DescribeMarketRemainsQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the created models of an API group.
//
// Description:
//
//	Fuzzy queries are supported.
//
// @param request - DescribeModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelsResponse
func (client *Client) DescribeModelsWithOptions(request *DescribeModelsRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModels"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the created models of an API group.
//
// Description:
//
//	Fuzzy queries are supported.
//
// @param request - DescribeModelsRequest
//
// @return DescribeModelsResponse
func (client *Client) DescribeModels(request *DescribeModelsRequest) (_result *DescribeModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeModelsResponse{}
	_body, _err := client.DescribeModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified plug-in is bound.
//
// @param request - DescribePluginApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginApisResponse
func (client *Client) DescribePluginApisWithOptions(request *DescribePluginApisRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs to which a specified plug-in is bound.
//
// @param request - DescribePluginApisRequest
//
// @return DescribePluginApisResponse
func (client *Client) DescribePluginApis(request *DescribePluginApisRequest) (_result *DescribePluginApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginApisResponse{}
	_body, _err := client.DescribePluginApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of groups bound to a plugin based on the plugin ID
//
// @param request - DescribePluginGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginGroupsResponse
func (client *Client) DescribePluginGroupsWithOptions(request *DescribePluginGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of groups bound to a plugin based on the plugin ID
//
// @param request - DescribePluginGroupsRequest
//
// @return DescribePluginGroupsResponse
func (client *Client) DescribePluginGroups(request *DescribePluginGroupsRequest) (_result *DescribePluginGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginGroupsResponse{}
	_body, _err := client.DescribePluginGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询插件列表
//
// @param request - DescribePluginSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginSchemasResponse
func (client *Client) DescribePluginSchemasWithOptions(request *DescribePluginSchemasRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginSchemasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginSchemas"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询插件列表
//
// @param request - DescribePluginSchemasRequest
//
// @return DescribePluginSchemasResponse
func (client *Client) DescribePluginSchemas(request *DescribePluginSchemasRequest) (_result *DescribePluginSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginSchemasResponse{}
	_body, _err := client.DescribePluginSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询插件模板
//
// @param request - DescribePluginTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginTemplatesResponse
func (client *Client) DescribePluginTemplatesWithOptions(request *DescribePluginTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginTemplates"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询插件模板
//
// @param request - DescribePluginTemplatesRequest
//
// @return DescribePluginTemplatesResponse
func (client *Client) DescribePluginTemplates(request *DescribePluginTemplatesRequest) (_result *DescribePluginTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginTemplatesResponse{}
	_body, _err := client.DescribePluginTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries API Gateway plug-ins and the details of the plug-ins.
//
// Description:
//
//	  This operation supports pagination.
//
//		- This operation allows you to query plug-ins by business type.
//
//		- This operation allows you to query plug-ins by ID.
//
//		- This operation allows you to query plug-ins by name.
//
// @param request - DescribePluginsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsResponse
func (client *Client) DescribePluginsWithOptions(request *DescribePluginsRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsResponse, _err error) {
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

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlugins"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries API Gateway plug-ins and the details of the plug-ins.
//
// Description:
//
//	  This operation supports pagination.
//
//		- This operation allows you to query plug-ins by business type.
//
//		- This operation allows you to query plug-ins by ID.
//
//		- This operation allows you to query plug-ins by name.
//
// @param request - DescribePluginsRequest
//
// @return DescribePluginsResponse
func (client *Client) DescribePlugins(request *DescribePluginsRequest) (_result *DescribePluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginsResponse{}
	_body, _err := client.DescribePluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the plug-ins that are bound to a running API in an environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation supports pagination.
//
// @param request - DescribePluginsByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsByApiResponse
func (client *Client) DescribePluginsByApiWithOptions(request *DescribePluginsByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsByApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginsByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the plug-ins that are bound to a running API in an environment.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- This operation supports pagination.
//
// @param request - DescribePluginsByApiRequest
//
// @return DescribePluginsByApiResponse
func (client *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (_result *DescribePluginsByApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginsByApiResponse{}
	_body, _err := client.DescribePluginsByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Plugins Bound to API Group
//
// @param request - DescribePluginsByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePluginsByGroupResponse
func (client *Client) DescribePluginsByGroupWithOptions(request *DescribePluginsByGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePluginsByGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePluginsByGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePluginsByGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Plugins Bound to API Group
//
// @param request - DescribePluginsByGroupRequest
//
// @return DescribePluginsByGroupResponse
func (client *Client) DescribePluginsByGroup(request *DescribePluginsByGroupRequest) (_result *DescribePluginsByGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePluginsByGroupResponse{}
	_body, _err := client.DescribePluginsByGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about an API group purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApiGroupResponse
func (client *Client) DescribePurchasedApiGroupWithOptions(request *DescribePurchasedApiGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an API group purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupRequest
//
// @return DescribePurchasedApiGroupResponse
func (client *Client) DescribePurchasedApiGroup(request *DescribePurchasedApiGroupRequest) (_result *DescribePurchasedApiGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurchasedApiGroupResponse{}
	_body, _err := client.DescribePurchasedApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the API groups purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApiGroupsResponse
func (client *Client) DescribePurchasedApiGroupsWithOptions(request *DescribePurchasedApiGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApiGroupsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApiGroups"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the API groups purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApiGroupsRequest
//
// @return DescribePurchasedApiGroupsResponse
func (client *Client) DescribePurchasedApiGroups(request *DescribePurchasedApiGroupsRequest) (_result *DescribePurchasedApiGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurchasedApiGroupsResponse{}
	_body, _err := client.DescribePurchasedApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries APIs that are purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePurchasedApisResponse
func (client *Client) DescribePurchasedApisWithOptions(request *DescribePurchasedApisRequest, runtime *dara.RuntimeOptions) (_result *DescribePurchasedApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePurchasedApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePurchasedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries APIs that are purchased from Alibaba Cloud Marketplace.
//
// @param request - DescribePurchasedApisRequest
//
// @return DescribePurchasedApisResponse
func (client *Client) DescribePurchasedApis(request *DescribePurchasedApisRequest) (_result *DescribePurchasedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePurchasedApisResponse{}
	_body, _err := client.DescribePurchasedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Alibaba Cloud regions that are supported by API Gateway.
//
// Description:
//
// This operation queries regions in which API Gateway is available.
//
//   - This operation is intended for API providers and callers.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Queries the Alibaba Cloud regions that are supported by API Gateway.
//
// Description:
//
// This operation queries regions in which API Gateway is available.
//
//   - This operation is intended for API providers and callers.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries backend signature keys.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This operation is used to query the backend signature keys in a Region. Region is a system parameter.
//
// @param request - DescribeSignaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSignaturesResponse
func (client *Client) DescribeSignaturesWithOptions(request *DescribeSignaturesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSignaturesResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSignatures"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries backend signature keys.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This operation is used to query the backend signature keys in a Region. Region is a system parameter.
//
// @param request - DescribeSignaturesRequest
//
// @return DescribeSignaturesResponse
func (client *Client) DescribeSignatures(request *DescribeSignaturesRequest) (_result *DescribeSignaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSignaturesResponse{}
	_body, _err := client.DescribeSignaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to a specified API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeSignaturesByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSignaturesByApiResponse
func (client *Client) DescribeSignaturesByApiWithOptions(request *DescribeSignaturesByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeSignaturesByApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSignaturesByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSignaturesByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the backend signature keys that are bound to a specified API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeSignaturesByApiRequest
//
// @return DescribeSignaturesByApiResponse
func (client *Client) DescribeSignaturesByApi(request *DescribeSignaturesByApiRequest) (_result *DescribeSignaturesByApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSignaturesByApiResponse{}
	_body, _err := client.DescribeSignaturesByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of API Gateway resources in a region.
//
// @param request - DescribeSummaryDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSummaryDataResponse
func (client *Client) DescribeSummaryDataWithOptions(request *DescribeSummaryDataRequest, runtime *dara.RuntimeOptions) (_result *DescribeSummaryDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSummaryData"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSummaryDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of API Gateway resources in a region.
//
// @param request - DescribeSummaryDataRequest
//
// @return DescribeSummaryDataResponse
func (client *Client) DescribeSummaryData(request *DescribeSummaryDataRequest) (_result *DescribeSummaryDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSummaryDataResponse{}
	_body, _err := client.DescribeSummaryDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the common parameters supported by the system.
//
// Description:
//
//	  This API is intended for API callers.
//
//		- The response of this API contains the system parameters that are optional in API definitions.
//
// @param request - DescribeSystemParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemParametersResponse
func (client *Client) DescribeSystemParametersWithOptions(request *DescribeSystemParametersRequest, runtime *dara.RuntimeOptions) (_result *DescribeSystemParametersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSystemParameters"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the common parameters supported by the system.
//
// Description:
//
//	  This API is intended for API callers.
//
//		- The response of this API contains the system parameters that are optional in API definitions.
//
// @param request - DescribeSystemParametersRequest
//
// @return DescribeSystemParametersResponse
func (client *Client) DescribeSystemParameters(request *DescribeSystemParametersRequest) (_result *DescribeSystemParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSystemParametersResponse{}
	_body, _err := client.DescribeSystemParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom throttling policies and their details. Conditional queries are supported.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API can be used to query all existing throttling policies (including special throttling policies) and their details.
//
//		- You can specify query conditions. For example, you can query the throttling policies bound to a specified API or in a specified environment.
//
// @param request - DescribeTrafficControlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficControlsResponse
func (client *Client) DescribeTrafficControlsWithOptions(request *DescribeTrafficControlsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficControlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrafficControls"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom throttling policies and their details. Conditional queries are supported.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API can be used to query all existing throttling policies (including special throttling policies) and their details.
//
//		- You can specify query conditions. For example, you can query the throttling policies bound to a specified API or in a specified environment.
//
// @param request - DescribeTrafficControlsRequest
//
// @return DescribeTrafficControlsResponse
func (client *Client) DescribeTrafficControls(request *DescribeTrafficControlsRequest) (_result *DescribeTrafficControlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrafficControlsResponse{}
	_body, _err := client.DescribeTrafficControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the throttling policy that is bound to a specific API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeTrafficControlsByApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficControlsByApiResponse
func (client *Client) DescribeTrafficControlsByApiWithOptions(request *DescribeTrafficControlsByApiRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficControlsByApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTrafficControlsByApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficControlsByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the throttling policy that is bound to a specific API.
//
// Description:
//
//	This API is intended for API providers.
//
// @param request - DescribeTrafficControlsByApiRequest
//
// @return DescribeTrafficControlsByApiResponse
func (client *Client) DescribeTrafficControlsByApi(request *DescribeTrafficControlsByApiRequest) (_result *DescribeTrafficControlsByApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrafficControlsByApiResponse{}
	_body, _err := client.DescribeTrafficControlsByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询批量更新API后端元定结果
//
// @param request - DescribeUpdateBackendTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpdateBackendTaskResponse
func (client *Client) DescribeUpdateBackendTaskWithOptions(request *DescribeUpdateBackendTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpdateBackendTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpdateBackendTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpdateBackendTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批量更新API后端元定结果
//
// @param request - DescribeUpdateBackendTaskRequest
//
// @return DescribeUpdateBackendTaskResponse
func (client *Client) DescribeUpdateBackendTask(request *DescribeUpdateBackendTaskRequest) (_result *DescribeUpdateBackendTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUpdateBackendTaskResponse{}
	_body, _err := client.DescribeUpdateBackendTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询更新VPC授权的任务
//
// @param request - DescribeUpdateVpcInfoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUpdateVpcInfoTaskResponse
func (client *Client) DescribeUpdateVpcInfoTaskWithOptions(request *DescribeUpdateVpcInfoTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeUpdateVpcInfoTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationUid) {
		query["OperationUid"] = request.OperationUid
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUpdateVpcInfoTask"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUpdateVpcInfoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询更新VPC授权的任务
//
// @param request - DescribeUpdateVpcInfoTaskRequest
//
// @return DescribeUpdateVpcInfoTaskResponse
func (client *Client) DescribeUpdateVpcInfoTask(request *DescribeUpdateVpcInfoTaskRequest) (_result *DescribeUpdateVpcInfoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeUpdateVpcInfoTaskResponse{}
	_body, _err := client.DescribeUpdateVpcInfoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries VPC access authorizations.
//
// @param request - DescribeVpcAccessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVpcAccessesResponse
func (client *Client) DescribeVpcAccessesWithOptions(request *DescribeVpcAccessesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVpcAccessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccurateQuery) {
		query["AccurateQuery"] = request.AccurateQuery
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcAccessId) {
		query["VpcAccessId"] = request.VpcAccessId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVpcAccesses"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVpcAccessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries VPC access authorizations.
//
// @param request - DescribeVpcAccessesRequest
//
// @return DescribeVpcAccessesResponse
func (client *Client) DescribeVpcAccesses(request *DescribeVpcAccessesRequest) (_result *DescribeVpcAccessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVpcAccessesResponse{}
	_body, _err := client.DescribeVpcAccessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries zones in a region.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries zones in a region.
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches APIs from an API product.
//
// @param request - DetachApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachApiProductResponse
func (client *Client) DetachApiProductWithOptions(request *DetachApiProductRequest, runtime *dara.RuntimeOptions) (_result *DetachApiProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.Apis) {
		query["Apis"] = request.Apis
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachApiProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches APIs from an API product.
//
// @param request - DetachApiProductRequest
//
// @return DetachApiProductResponse
func (client *Client) DetachApiProduct(request *DetachApiProductRequest) (_result *DetachApiProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachApiProductResponse{}
	_body, _err := client.DetachApiProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Unbind group plugin
//
// @param request - DetachGroupPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachGroupPluginResponse
func (client *Client) DetachGroupPluginWithOptions(request *DetachGroupPluginRequest, runtime *dara.RuntimeOptions) (_result *DetachGroupPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachGroupPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachGroupPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Unbind group plugin
//
// @param request - DetachGroupPluginRequest
//
// @return DetachGroupPluginResponse
func (client *Client) DetachGroupPlugin(request *DetachGroupPluginRequest) (_result *DetachGroupPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachGroupPluginResponse{}
	_body, _err := client.DetachGroupPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑插件
//
// @param request - DetachPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPluginResponse
func (client *Client) DetachPluginWithOptions(request *DetachPluginRequest, runtime *dara.RuntimeOptions) (_result *DetachPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑插件
//
// @param request - DetachPluginRequest
//
// @return DetachPluginResponse
func (client *Client) DetachPlugin(request *DetachPluginRequest) (_result *DetachPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachPluginResponse{}
	_body, _err := client.DetachPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Disables access control on an instance.
//
// @param request - DisableInstanceAccessControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableInstanceAccessControlResponse
func (client *Client) DisableInstanceAccessControlWithOptions(request *DisableInstanceAccessControlRequest, runtime *dara.RuntimeOptions) (_result *DisableInstanceAccessControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableInstanceAccessControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableInstanceAccessControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Disables access control on an instance.
//
// @param request - DisableInstanceAccessControlRequest
//
// @return DisableInstanceAccessControlResponse
func (client *Client) DisableInstanceAccessControl(request *DisableInstanceAccessControlRequest) (_result *DisableInstanceAccessControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableInstanceAccessControlResponse{}
	_body, _err := client.DisableInstanceAccessControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates an internal domain name resolution from a dedicated instance.
//
// @param tmpReq - DissociateInstanceWithPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateInstanceWithPrivateDNSResponse
func (client *Client) DissociateInstanceWithPrivateDNSWithOptions(tmpReq *DissociateInstanceWithPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *DissociateInstanceWithPrivateDNSResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DissociateInstanceWithPrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntranetDomains) {
		request.IntranetDomainsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntranetDomains, dara.String("IntranetDomains"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomainsShrink) {
		body["IntranetDomains"] = request.IntranetDomainsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateInstanceWithPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an internal domain name resolution from a dedicated instance.
//
// @param request - DissociateInstanceWithPrivateDNSRequest
//
// @return DissociateInstanceWithPrivateDNSResponse
func (client *Client) DissociateInstanceWithPrivateDNS(request *DissociateInstanceWithPrivateDNSRequest) (_result *DissociateInstanceWithPrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateInstanceWithPrivateDNSResponse{}
	_body, _err := client.DissociateInstanceWithPrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the syntax before Swagger-compliant data is imported.
//
// @param tmpReq - DryRunSwaggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DryRunSwaggerResponse
func (client *Client) DryRunSwaggerWithOptions(tmpReq *DryRunSwaggerRequest, runtime *dara.RuntimeOptions) (_result *DryRunSwaggerResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DryRunSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalCondition) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, dara.String("GlobalCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.GlobalConditionShrink) {
		query["GlobalCondition"] = request.GlobalConditionShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DryRunSwagger"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DryRunSwaggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the syntax before Swagger-compliant data is imported.
//
// @param request - DryRunSwaggerRequest
//
// @return DryRunSwaggerResponse
func (client *Client) DryRunSwagger(request *DryRunSwaggerRequest) (_result *DryRunSwaggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DryRunSwaggerResponse{}
	_body, _err := client.DryRunSwaggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Specifies the access control policy of an instance.
//
// @param request - EnableInstanceAccessControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInstanceAccessControlResponse
func (client *Client) EnableInstanceAccessControlWithOptions(request *EnableInstanceAccessControlRequest, runtime *dara.RuntimeOptions) (_result *EnableInstanceAccessControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableInstanceAccessControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInstanceAccessControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This feature provides instance-level access control capabilities for dedicated API Gateway instances. Specifies the access control policy of an instance.
//
// @param request - EnableInstanceAccessControlRequest
//
// @return EnableInstanceAccessControlResponse
func (client *Client) EnableInstanceAccessControl(request *EnableInstanceAccessControlRequest) (_result *EnableInstanceAccessControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableInstanceAccessControlResponse{}
	_body, _err := client.EnableInstanceAccessControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出OAS
//
// @param tmpReq - ExportOASRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportOASResponse
func (client *Client) ExportOASWithOptions(tmpReq *ExportOASRequest, runtime *dara.RuntimeOptions) (_result *ExportOASResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExportOASShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiIdList) {
		request.ApiIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiIdList, dara.String("ApiIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIdListShrink) {
		query["ApiIdList"] = request.ApiIdListShrink
	}

	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OasVersion) {
		query["OasVersion"] = request.OasVersion
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.WithXExtensions) {
		query["WithXExtensions"] = request.WithXExtensions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportOAS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportOASResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出OAS
//
// @param request - ExportOASRequest
//
// @return ExportOASResponse
func (client *Client) ExportOAS(request *ExportOASRequest) (_result *ExportOASResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportOASResponse{}
	_body, _err := client.ExportOASWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports APIs based on the OAS standard.
//
// @param request - ImportOASRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportOASResponse
func (client *Client) ImportOASWithOptions(request *ImportOASRequest, runtime *dara.RuntimeOptions) (_result *ImportOASResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IgnoreWarning) {
		query["IgnoreWarning"] = request.IgnoreWarning
	}

	if !dara.IsNil(request.OASVersion) {
		query["OASVersion"] = request.OASVersion
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.RequestMode) {
		query["RequestMode"] = request.RequestMode
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SkipDryRun) {
		query["SkipDryRun"] = request.SkipDryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportOAS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportOASResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports APIs based on the OAS standard.
//
// @param request - ImportOASRequest
//
// @return ImportOASResponse
func (client *Client) ImportOAS(request *ImportOASRequest) (_result *ImportOASResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportOASResponse{}
	_body, _err := client.ImportOASWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an API by importing Swagger-compliant data.
//
// Description:
//
//	  Alibaba Cloud supports extensions based on Swagger 2.0.
//
//		- Alibaba Cloud supports Swagger configuration files in JSON and YAML formats.
//
// @param tmpReq - ImportSwaggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportSwaggerResponse
func (client *Client) ImportSwaggerWithOptions(tmpReq *ImportSwaggerRequest, runtime *dara.RuntimeOptions) (_result *ImportSwaggerResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ImportSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GlobalCondition) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, dara.String("GlobalCondition"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DataFormat) {
		query["DataFormat"] = request.DataFormat
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GlobalConditionShrink) {
		query["GlobalCondition"] = request.GlobalConditionShrink
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportSwagger"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportSwaggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API by importing Swagger-compliant data.
//
// Description:
//
//	  Alibaba Cloud supports extensions based on Swagger 2.0.
//
//		- Alibaba Cloud supports Swagger configuration files in JSON and YAML formats.
//
// @param request - ImportSwaggerRequest
//
// @return ImportSwaggerResponse
func (client *Client) ImportSwagger(request *ImportSwaggerRequest) (_result *ImportSwaggerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportSwaggerResponse{}
	_body, _err := client.ImportSwaggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries internal domain name resolutions by domain name or resolution type.
//
// @param request - ListPrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateDNSResponse
func (client *Client) ListPrivateDNSWithOptions(request *ListPrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateDNSResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries internal domain name resolutions by domain name or resolution type.
//
// @param request - ListPrivateDNSRequest
//
// @return ListPrivateDNSResponse
func (client *Client) ListPrivateDNS(request *ListPrivateDNSRequest) (_result *ListPrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateDNSResponse{}
	_body, _err := client.ListPrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the visible resource tags.
//
// Description:
//
//	  The Tag.N.Key and Tag.N.Value parameters constitute a key-value pair.
//
//		- ResourceId.N must meet all the key-value pairs that are entered. If you enter multiple key-value pairs, resources that contain the specified key-value pairs are returned.
//
//		- This operation is used to query resource tags based on conditions. If no relationship matches the conditions, an empty list is returned.
//
//		- You can query both user tags and visible system tags.
//
//		- In addition to the required parameters, you can also specify ResourceId.N to query the visible resource tags of a specified resource in a region.
//
//		- You can also specify Tag.N.Key to query the visible keys of a specified key in a region.
//
//		- At least one of ResourceId.N, Tag.N.Key, and Tag.N.Value exists.
//
//		- You can query tags of the same type or different types in a single operation.
//
//		- You can query all your user types and visible system tags.
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2016-07-14"),
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
// Queries the visible resource tags.
//
// Description:
//
//	  The Tag.N.Key and Tag.N.Value parameters constitute a key-value pair.
//
//		- ResourceId.N must meet all the key-value pairs that are entered. If you enter multiple key-value pairs, resources that contain the specified key-value pairs are returned.
//
//		- This operation is used to query resource tags based on conditions. If no relationship matches the conditions, an empty list is returned.
//
//		- You can query both user tags and visible system tags.
//
//		- In addition to the required parameters, you can also specify ResourceId.N to query the visible resource tags of a specified resource in a region.
//
//		- You can also specify Tag.N.Key to query the visible keys of a specified key in a region.
//
//		- At least one of ResourceId.N, Tag.N.Key, and Tag.N.Value exists.
//
//		- You can query tags of the same type or different types in a single operation.
//
//		- You can query all your user types and visible system tags.
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
// Modifies the definition of an API.
//
// Description:
//
// *This operation is intended for API providers.**
//
//   - This API operation requires a full update. Updates of partial parameters are not supported.
//
//   - When you modify an API name, make sure that the name of each API within the same group is unique.
//
//   - When you modify the request path, make sure that each request path within the same group is unique.
//
//   - The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiResponse
func (client *Client) ModifyApiWithOptions(request *ModifyApiRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendEnable) {
		query["BackendEnable"] = request.BackendEnable
	}

	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.OpenIdConnectConfig) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !dara.IsNil(request.RequestConfig) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !dara.IsNil(request.ResultBodyModel) {
		query["ResultBodyModel"] = request.ResultBodyModel
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceConfig) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.WebSocketApiType) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConstantParameters) {
		body["ConstantParameters"] = request.ConstantParameters
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		body["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		body["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.RequestParameters) {
		body["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.ResultDescriptions) {
		body["ResultDescriptions"] = request.ResultDescriptions
	}

	if !dara.IsNil(request.ResultSample) {
		body["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ServiceParameters) {
		body["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		body["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.SystemParameters) {
		body["SystemParameters"] = request.SystemParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the definition of an API.
//
// Description:
//
// *This operation is intended for API providers.**
//
//   - This API operation requires a full update. Updates of partial parameters are not supported.
//
//   - When you modify an API name, make sure that the name of each API within the same group is unique.
//
//   - When you modify the request path, make sure that each request path within the same group is unique.
//
//   - The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiRequest
//
// @return ModifyApiResponse
func (client *Client) ModifyApi(request *ModifyApiRequest) (_result *ModifyApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiResponse{}
	_body, _err := client.ModifyApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the draft definition of an API. This operation is different from the ModifyApi operation. This operation does not require all information about the API. You need to only specify the parameters that you want to modify. For example, if you want to change the authentication method of the API from Anonymous to APP, you specify APP as the value of AuthType and do not need to configure other parameters.
//
// @param request - ModifyApiConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiConfigurationResponse
func (client *Client) ModifyApiConfigurationWithOptions(request *ModifyApiConfigurationRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowSignatureMethod) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
	}

	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.AppCodeAuthType) {
		query["AppCodeAuthType"] = request.AppCodeAuthType
	}

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BodyFormat) {
		query["BodyFormat"] = request.BodyFormat
	}

	if !dara.IsNil(request.BodyModel) {
		query["BodyModel"] = request.BodyModel
	}

	if !dara.IsNil(request.ContentTypeCategory) {
		query["ContentTypeCategory"] = request.ContentTypeCategory
	}

	if !dara.IsNil(request.ContentTypeValue) {
		query["ContentTypeValue"] = request.ContentTypeValue
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableInternet) {
		query["DisableInternet"] = request.DisableInternet
	}

	if !dara.IsNil(request.ErrorCodeSamples) {
		query["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !dara.IsNil(request.FailResultSample) {
		query["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.ForceNonceCheck) {
		query["ForceNonceCheck"] = request.ForceNonceCheck
	}

	if !dara.IsNil(request.FunctionComputeConfig) {
		query["FunctionComputeConfig"] = request.FunctionComputeConfig
	}

	if !dara.IsNil(request.HttpConfig) {
		query["HttpConfig"] = request.HttpConfig
	}

	if !dara.IsNil(request.MockConfig) {
		query["MockConfig"] = request.MockConfig
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.OssConfig) {
		query["OssConfig"] = request.OssConfig
	}

	if !dara.IsNil(request.PostBodyDescription) {
		query["PostBodyDescription"] = request.PostBodyDescription
	}

	if !dara.IsNil(request.RequestHttpMethod) {
		query["RequestHttpMethod"] = request.RequestHttpMethod
	}

	if !dara.IsNil(request.RequestMode) {
		query["RequestMode"] = request.RequestMode
	}

	if !dara.IsNil(request.RequestParameters) {
		query["RequestParameters"] = request.RequestParameters
	}

	if !dara.IsNil(request.RequestPath) {
		query["RequestPath"] = request.RequestPath
	}

	if !dara.IsNil(request.RequestProtocol) {
		query["RequestProtocol"] = request.RequestProtocol
	}

	if !dara.IsNil(request.ResultSample) {
		query["ResultSample"] = request.ResultSample
	}

	if !dara.IsNil(request.ResultType) {
		query["ResultType"] = request.ResultType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	if !dara.IsNil(request.ServiceParametersMap) {
		query["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !dara.IsNil(request.ServiceProtocol) {
		query["ServiceProtocol"] = request.ServiceProtocol
	}

	if !dara.IsNil(request.ServiceTimeout) {
		query["ServiceTimeout"] = request.ServiceTimeout
	}

	if !dara.IsNil(request.UseBackendService) {
		query["UseBackendService"] = request.UseBackendService
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	if !dara.IsNil(request.VpcConfig) {
		query["VpcConfig"] = request.VpcConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiConfiguration"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the draft definition of an API. This operation is different from the ModifyApi operation. This operation does not require all information about the API. You need to only specify the parameters that you want to modify. For example, if you want to change the authentication method of the API from Anonymous to APP, you specify APP as the value of AuthType and do not need to configure other parameters.
//
// @param request - ModifyApiConfigurationRequest
//
// @return ModifyApiConfigurationResponse
func (client *Client) ModifyApiConfiguration(request *ModifyApiConfigurationRequest) (_result *ModifyApiConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiConfigurationResponse{}
	_body, _err := client.ModifyApiConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or basepath of an existing API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupResponse
func (client *Client) ModifyApiGroupWithOptions(request *ModifyApiGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BasePath) {
		query["BasePath"] = request.BasePath
	}

	if !dara.IsNil(request.CompatibleFlags) {
		query["CompatibleFlags"] = request.CompatibleFlags
	}

	if !dara.IsNil(request.CustomAppCodeConfig) {
		query["CustomAppCodeConfig"] = request.CustomAppCodeConfig
	}

	if !dara.IsNil(request.CustomTraceConfig) {
		query["CustomTraceConfig"] = request.CustomTraceConfig
	}

	if !dara.IsNil(request.CustomerConfigs) {
		query["CustomerConfigs"] = request.CustomerConfigs
	}

	if !dara.IsNil(request.DefaultDomain) {
		query["DefaultDomain"] = request.DefaultDomain
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FilterAppCodeForBackend) {
		query["FilterAppCodeForBackend"] = request.FilterAppCodeForBackend
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PassthroughHeaders) {
		query["PassthroughHeaders"] = request.PassthroughHeaders
	}

	if !dara.IsNil(request.RpcPattern) {
		query["RpcPattern"] = request.RpcPattern
	}

	if !dara.IsNil(request.RpsLimitForServerless) {
		query["RpsLimitForServerless"] = request.RpsLimitForServerless
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SupportSSE) {
		query["SupportSSE"] = request.SupportSSE
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UserLogConfig) {
		query["UserLogConfig"] = request.UserLogConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, or basepath of an existing API group.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyApiGroupRequest
//
// @return ModifyApiGroupResponse
func (client *Client) ModifyApiGroup(request *ModifyApiGroupRequest) (_result *ModifyApiGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiGroupResponse{}
	_body, _err := client.ModifyApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变更分组实例
//
// @param request - ModifyApiGroupInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupInstanceResponse
func (client *Client) ModifyApiGroupInstanceWithOptions(request *ModifyApiGroupInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetInstanceId) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupInstance"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变更分组实例
//
// @param request - ModifyApiGroupInstanceRequest
//
// @return ModifyApiGroupInstanceResponse
func (client *Client) ModifyApiGroupInstance(request *ModifyApiGroupInstanceRequest) (_result *ModifyApiGroupInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiGroupInstanceResponse{}
	_body, _err := client.ModifyApiGroupInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the network policy of an API group.
//
// @param request - ModifyApiGroupNetworkPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupNetworkPolicyResponse
func (client *Client) ModifyApiGroupNetworkPolicyWithOptions(request *ModifyApiGroupNetworkPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupNetworkPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.InnerDomainEnable) {
		query["InnerDomainEnable"] = request.InnerDomainEnable
	}

	if !dara.IsNil(request.InternetEnable) {
		query["InternetEnable"] = request.InternetEnable
	}

	if !dara.IsNil(request.InternetIPV6Enable) {
		query["InternetIPV6Enable"] = request.InternetIPV6Enable
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIntranetEnable) {
		query["VpcIntranetEnable"] = request.VpcIntranetEnable
	}

	if !dara.IsNil(request.VpcSlbIntranetEnable) {
		query["VpcSlbIntranetEnable"] = request.VpcSlbIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupNetworkPolicy"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupNetworkPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network policy of an API group.
//
// @param request - ModifyApiGroupNetworkPolicyRequest
//
// @return ModifyApiGroupNetworkPolicyResponse
func (client *Client) ModifyApiGroupNetworkPolicy(request *ModifyApiGroupNetworkPolicyRequest) (_result *ModifyApiGroupNetworkPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiGroupNetworkPolicyResponse{}
	_body, _err := client.ModifyApiGroupNetworkPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the VPC whitelist of an API group.
//
// @param request - ModifyApiGroupVpcWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApiGroupVpcWhitelistResponse
func (client *Client) ModifyApiGroupVpcWhitelistWithOptions(request *ModifyApiGroupVpcWhitelistRequest, runtime *dara.RuntimeOptions) (_result *ModifyApiGroupVpcWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIds) {
		query["VpcIds"] = request.VpcIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApiGroupVpcWhitelist"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the VPC whitelist of an API group.
//
// @param request - ModifyApiGroupVpcWhitelistRequest
//
// @return ModifyApiGroupVpcWhitelistResponse
func (client *Client) ModifyApiGroupVpcWhitelist(request *ModifyApiGroupVpcWhitelistRequest) (_result *ModifyApiGroupVpcWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApiGroupVpcWhitelistResponse{}
	_body, _err := client.ModifyApiGroupVpcWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a specified app.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppName or Description can be modified. If these parameters are not specified, no modifications are made and the operation will directly return a successful response.********
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *dara.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Disabled) {
		query["Disabled"] = request.Disabled
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a specified app.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- AppName or Description can be modified. If these parameters are not specified, no modifications are made and the operation will directly return a successful response.********
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改后端服务
//
// @param request - ModifyBackendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackendResponse
func (client *Client) ModifyBackendWithOptions(request *ModifyBackendRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendName) {
		query["BackendName"] = request.BackendName
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackend"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改后端服务
//
// @param request - ModifyBackendRequest
//
// @return ModifyBackendResponse
func (client *Client) ModifyBackend(request *ModifyBackendRequest) (_result *ModifyBackendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackendResponse{}
	_body, _err := client.ModifyBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改后端服务在环境上的定义
//
// @param request - ModifyBackendModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBackendModelResponse
func (client *Client) ModifyBackendModelWithOptions(request *ModifyBackendModelRequest, runtime *dara.RuntimeOptions) (_result *ModifyBackendModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendId) {
		query["BackendId"] = request.BackendId
	}

	if !dara.IsNil(request.BackendModelData) {
		query["BackendModelData"] = request.BackendModelData
	}

	if !dara.IsNil(request.BackendModelId) {
		query["BackendModelId"] = request.BackendModelId
	}

	if !dara.IsNil(request.BackendType) {
		query["BackendType"] = request.BackendType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBackendModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBackendModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改后端服务在环境上的定义
//
// @param request - ModifyBackendModelRequest
//
// @return ModifyBackendModelResponse
func (client *Client) ModifyBackendModel(request *ModifyBackendModelRequest) (_result *ModifyBackendModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBackendModelResponse{}
	_body, _err := client.ModifyBackendModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a custom dataset.
//
// @param request - ModifyDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatasetResponse
func (client *Client) ModifyDatasetWithOptions(request *ModifyDatasetRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatasetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetName) {
		query["DatasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataset"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a custom dataset.
//
// @param request - ModifyDatasetRequest
//
// @return ModifyDatasetResponse
func (client *Client) ModifyDataset(request *ModifyDatasetRequest) (_result *ModifyDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatasetResponse{}
	_body, _err := client.ModifyDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the expiration time and description of a data entry in a custom dataset.
//
// @param request - ModifyDatasetItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatasetItemResponse
func (client *Client) ModifyDatasetItemWithOptions(request *ModifyDatasetItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyDatasetItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		query["DatasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.DatasetItemId) {
		query["DatasetItemId"] = request.DatasetItemId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExpiredTime) {
		query["ExpiredTime"] = request.ExpiredTime
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDatasetItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDatasetItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the expiration time and description of a data entry in a custom dataset.
//
// @param request - ModifyDatasetItemRequest
//
// @return ModifyDatasetItemResponse
func (client *Client) ModifyDatasetItem(request *ModifyDatasetItemRequest) (_result *ModifyDatasetItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDatasetItemResponse{}
	_body, _err := client.ModifyDatasetItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the properties of an API Gateway instance.
//
// @param tmpReq - ModifyInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttributeWithOptions(tmpReq *ModifyInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyInstanceAttributeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ToConnectVpcIpBlock) {
		request.ToConnectVpcIpBlockShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToConnectVpcIpBlock, dara.String("ToConnectVpcIpBlock"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteVpcIpBlock) {
		query["DeleteVpcIpBlock"] = request.DeleteVpcIpBlock
	}

	if !dara.IsNil(request.EgressIpv6Enable) {
		query["EgressIpv6Enable"] = request.EgressIpv6Enable
	}

	if !dara.IsNil(request.HttpsPolicy) {
		query["HttpsPolicy"] = request.HttpsPolicy
	}

	if !dara.IsNil(request.IPV6Enabled) {
		query["IPV6Enabled"] = request.IPV6Enabled
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.IntranetSegments) {
		query["IntranetSegments"] = request.IntranetSegments
	}

	if !dara.IsNil(request.MaintainEndTime) {
		query["MaintainEndTime"] = request.MaintainEndTime
	}

	if !dara.IsNil(request.MaintainStartTime) {
		query["MaintainStartTime"] = request.MaintainStartTime
	}

	if !dara.IsNil(request.ToConnectVpcIpBlockShrink) {
		query["ToConnectVpcIpBlock"] = request.ToConnectVpcIpBlockShrink
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcSlbIntranetEnable) {
		query["VpcSlbIntranetEnable"] = request.VpcSlbIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of an API Gateway instance.
//
// @param request - ModifyInstanceAttributeRequest
//
// @return ModifyInstanceAttributeResponse
func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the configurations of an API Gateway instance.
//
// @param request - ModifyInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpecWithOptions(request *ModifyInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.ModifyAction) {
		query["ModifyAction"] = request.ModifyAction
	}

	if !dara.IsNil(request.SkipWaitSwitch) {
		query["SkipWaitSwitch"] = request.SkipWaitSwitch
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceSpec"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades or downgrades the configurations of an API Gateway instance.
//
// @param request - ModifyInstanceSpecRequest
//
// @return ModifyInstanceSpecResponse
func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (_result *ModifyInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.ModifyInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify instance client VPC config.
//
// @param request - ModifyInstanceVpcAttributeForConsoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVpcAttributeForConsoleResponse
func (client *Client) ModifyInstanceVpcAttributeForConsoleWithOptions(request *ModifyInstanceVpcAttributeForConsoleRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceVpcAttributeForConsoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteVpcAccess) {
		query["DeleteVpcAccess"] = request.DeleteVpcAccess
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		query["VpcOwnerId"] = request.VpcOwnerId
	}

	if !dara.IsNil(request.VswitchId) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceVpcAttributeForConsole"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceVpcAttributeForConsoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify instance client VPC config.
//
// @param request - ModifyInstanceVpcAttributeForConsoleRequest
//
// @return ModifyInstanceVpcAttributeForConsoleResponse
func (client *Client) ModifyInstanceVpcAttributeForConsole(request *ModifyInstanceVpcAttributeForConsoleRequest) (_result *ModifyInstanceVpcAttributeForConsoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceVpcAttributeForConsoleResponse{}
	_body, _err := client.ModifyInstanceVpcAttributeForConsoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the VPC domain name policy of an API group.
//
// @param request - ModifyIntranetDomainPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIntranetDomainPolicyResponse
func (client *Client) ModifyIntranetDomainPolicyWithOptions(request *ModifyIntranetDomainPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyIntranetDomainPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcIntranetEnable) {
		query["VpcIntranetEnable"] = request.VpcIntranetEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIntranetDomainPolicy"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIntranetDomainPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the VPC domain name policy of an API group.
//
// @param request - ModifyIntranetDomainPolicyRequest
//
// @return ModifyIntranetDomainPolicyResponse
func (client *Client) ModifyIntranetDomainPolicy(request *ModifyIntranetDomainPolicyRequest) (_result *ModifyIntranetDomainPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIntranetDomainPolicyResponse{}
	_body, _err := client.ModifyIntranetDomainPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation allows you to modify only the name and description of an ACL. You cannot modify the type of the ACL.
//
// @param request - ModifyIpControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpControlResponse
func (client *Client) ModifyIpControlWithOptions(request *ModifyIpControlRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.IpControlName) {
		query["IpControlName"] = request.IpControlName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- This operation allows you to modify only the name and description of an ACL. You cannot modify the type of the ACL.
//
// @param request - ModifyIpControlRequest
//
// @return ModifyIpControlResponse
func (client *Client) ModifyIpControl(request *ModifyIpControlRequest) (_result *ModifyIpControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIpControlResponse{}
	_body, _err := client.ModifyIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a policy in an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The modification immediately takes effect on all the APIs that are bound to the policy.
//
//		- This operation causes a full modification of the content of a policy.
//
// @param request - ModifyIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyIpControlPolicyItemResponse
func (client *Client) ModifyIpControlPolicyItemWithOptions(request *ModifyIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *ModifyIpControlPolicyItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CidrIp) {
		query["CidrIp"] = request.CidrIp
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PolicyItemId) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a policy in an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The modification immediately takes effect on all the APIs that are bound to the policy.
//
//		- This operation causes a full modification of the content of a policy.
//
// @param request - ModifyIpControlPolicyItemRequest
//
// @return ModifyIpControlPolicyItemResponse
func (client *Client) ModifyIpControlPolicyItem(request *ModifyIpControlPolicyItemRequest) (_result *ModifyIpControlPolicyItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyIpControlPolicyItemResponse{}
	_body, _err := client.ModifyIpControlPolicyItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改日志配置
//
// @param request - ModifyLogConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLogConfigResponse
func (client *Client) ModifyLogConfigWithOptions(request *ModifyLogConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyLogConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SlsLogStore) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !dara.IsNil(request.SlsProject) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLogConfig"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改日志配置
//
// @param request - ModifyLogConfigRequest
//
// @return ModifyLogConfigResponse
func (client *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (_result *ModifyLogConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLogConfigResponse{}
	_body, _err := client.ModifyLogConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the model of an API group.
//
// @param request - ModifyModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyModelResponse
func (client *Client) ModifyModelWithOptions(request *ModifyModelRequest, runtime *dara.RuntimeOptions) (_result *ModifyModelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ModelName) {
		query["ModelName"] = request.ModelName
	}

	if !dara.IsNil(request.NewModelName) {
		query["NewModelName"] = request.NewModelName
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyModel"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the model of an API group.
//
// @param request - ModifyModelRequest
//
// @return ModifyModelResponse
func (client *Client) ModifyModel(request *ModifyModelRequest) (_result *ModifyModelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyModelResponse{}
	_body, _err := client.ModifyModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of the plug-in must be unique.
//
// @param request - ModifyPluginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPluginResponse
func (client *Client) ModifyPluginWithOptions(request *ModifyPluginRequest, runtime *dara.RuntimeOptions) (_result *ModifyPluginResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.PluginData) {
		query["PluginData"] = request.PluginData
	}

	if !dara.IsNil(request.PluginId) {
		query["PluginId"] = request.PluginId
	}

	if !dara.IsNil(request.PluginName) {
		query["PluginName"] = request.PluginName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPlugin"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a plug-in.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The name of the plug-in must be unique.
//
// @param request - ModifyPluginRequest
//
// @return ModifyPluginResponse
func (client *Client) ModifyPlugin(request *ModifyPluginRequest) (_result *ModifyPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPluginResponse{}
	_body, _err := client.ModifyPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation modifies the name, Key value, and Secret value of an existing signature key.
//
//		- Note that the modification takes effect immediately. If the key has been bound to an API, you must adjust the backend signature verification based on the new key accordingly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifySignatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySignatureResponse
func (client *Client) ModifySignatureWithOptions(request *ModifySignatureRequest, runtime *dara.RuntimeOptions) (_result *ModifySignatureResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.SignatureKey) {
		query["SignatureKey"] = request.SignatureKey
	}

	if !dara.IsNil(request.SignatureName) {
		query["SignatureName"] = request.SignatureName
	}

	if !dara.IsNil(request.SignatureSecret) {
		query["SignatureSecret"] = request.SignatureSecret
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySignature"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a backend signature key.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API operation modifies the name, Key value, and Secret value of an existing signature key.
//
//		- Note that the modification takes effect immediately. If the key has been bound to an API, you must adjust the backend signature verification based on the new key accordingly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifySignatureRequest
//
// @return ModifySignatureResponse
func (client *Client) ModifySignature(request *ModifySignatureRequest) (_result *ModifySignatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySignatureResponse{}
	_body, _err := client.ModifySignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the settings of a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The modifications take effect on the bound APIs instantly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyTrafficControlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrafficControlResponse
func (client *Client) ModifyTrafficControlWithOptions(request *ModifyTrafficControlRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrafficControlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiDefault) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !dara.IsNil(request.AppDefault) {
		query["AppDefault"] = request.AppDefault
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !dara.IsNil(request.TrafficControlName) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !dara.IsNil(request.TrafficControlUnit) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !dara.IsNil(request.UserDefault) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTrafficControl"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of a custom throttling policy.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The modifications take effect on the bound APIs instantly.
//
//		- The QPS limit on this operation is 50 per user.
//
// @param request - ModifyTrafficControlRequest
//
// @return ModifyTrafficControlResponse
func (client *Client) ModifyTrafficControl(request *ModifyTrafficControlRequest) (_result *ModifyTrafficControlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTrafficControlResponse{}
	_body, _err := client.ModifyTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a virtual private cloud (VPC) authorization and updates the metadata of the API associated with the VPC authorization.
//
// @param request - ModifyVpcAccessAndUpdateApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVpcAccessAndUpdateApisResponse
func (client *Client) ModifyVpcAccessAndUpdateApisWithOptions(request *ModifyVpcAccessAndUpdateApisRequest, runtime *dara.RuntimeOptions) (_result *ModifyVpcAccessAndUpdateApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Refresh) {
		query["Refresh"] = request.Refresh
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcTargetHostName) {
		query["VpcTargetHostName"] = request.VpcTargetHostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVpcAccessAndUpdateApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVpcAccessAndUpdateApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a virtual private cloud (VPC) authorization and updates the metadata of the API associated with the VPC authorization.
//
// @param request - ModifyVpcAccessAndUpdateApisRequest
//
// @return ModifyVpcAccessAndUpdateApisResponse
func (client *Client) ModifyVpcAccessAndUpdateApis(request *ModifyVpcAccessAndUpdateApisRequest) (_result *ModifyVpcAccessAndUpdateApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVpcAccessAndUpdateApisResponse{}
	_body, _err := client.ModifyVpcAccessAndUpdateApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通API网关服务
//
// @param request - OpenApiGatewayServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenApiGatewayServiceResponse
func (client *Client) OpenApiGatewayServiceWithOptions(runtime *dara.RuntimeOptions) (_result *OpenApiGatewayServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("OpenApiGatewayService"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenApiGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通API网关服务
//
// @return OpenApiGatewayServiceResponse
func (client *Client) OpenApiGatewayService() (_result *OpenApiGatewayServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenApiGatewayServiceResponse{}
	_body, _err := client.OpenApiGatewayServiceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the request logs of a user.
//
// @param request - QueryRequestLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRequestLogsResponse
func (client *Client) QueryRequestLogsWithOptions(request *QueryRequestLogsRequest, runtime *dara.RuntimeOptions) (_result *QueryRequestLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RequestLogId) {
		query["RequestLogId"] = request.RequestLogId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRequestLogs"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRequestLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the request logs of a user.
//
// @param request - QueryRequestLogsRequest
//
// @return QueryRequestLogsResponse
func (client *Client) QueryRequestLogs(request *QueryRequestLogsRequest) (_result *QueryRequestLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRequestLogsResponse{}
	_body, _err := client.QueryRequestLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reactivates a custom domain name whose validity status is Abnormal.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must solve the problem that is mentioned in the domain name exception prompt before you can reactivate the domain name.
//
//		- A typical reason why a custom domain name becomes abnormal is that the domain name does not have an ICP filing or the domain name is included in a blacklist by the administration. When a custom domain name is abnormal, users cannot use it to call APIs.
//
//		- You can call this operation to reactivate the domain name to resume normal access.
//
// @param request - ReactivateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReactivateDomainResponse
func (client *Client) ReactivateDomainWithOptions(request *ReactivateDomainRequest, runtime *dara.RuntimeOptions) (_result *ReactivateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReactivateDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReactivateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reactivates a custom domain name whose validity status is Abnormal.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- You must solve the problem that is mentioned in the domain name exception prompt before you can reactivate the domain name.
//
//		- A typical reason why a custom domain name becomes abnormal is that the domain name does not have an ICP filing or the domain name is included in a blacklist by the administration. When a custom domain name is abnormal, users cannot use it to call APIs.
//
//		- You can call this operation to reactivate the domain name to resume normal access.
//
// @param request - ReactivateDomainRequest
//
// @return ReactivateDomainResponse
func (client *Client) ReactivateDomain(request *ReactivateDomainRequest) (_result *ReactivateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReactivateDomainResponse{}
	_body, _err := client.ReactivateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除访问控制策略中IP条目
//
// @param request - RemoveAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAccessControlListEntryResponse
func (client *Client) RemoveAccessControlListEntryWithOptions(request *RemoveAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *RemoveAccessControlListEntryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAccessControlListEntry"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAccessControlListEntryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除访问控制策略中IP条目
//
// @param request - RemoveAccessControlListEntryRequest
//
// @return RemoveAccessControlListEntryResponse
func (client *Client) RemoveAccessControlListEntry(request *RemoveAccessControlListEntryRequest) (_result *RemoveAccessControlListEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveAccessControlListEntryResponse{}
	_body, _err := client.RemoveAccessControlListEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on API products from an application.
//
// @param tmpReq - RemoveApiProductsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApiProductsAuthoritiesResponse
func (client *Client) RemoveApiProductsAuthoritiesWithOptions(tmpReq *RemoveApiProductsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveApiProductsAuthoritiesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveApiProductsAuthoritiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiProductIds) {
		request.ApiProductIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiProductIds, dara.String("ApiProductIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductIdsShrink) {
		query["ApiProductIds"] = request.ApiProductIdsShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApiProductsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApiProductsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on API products from an application.
//
// @param request - RemoveApiProductsAuthoritiesRequest
//
// @return RemoveApiProductsAuthoritiesResponse
func (client *Client) RemoveApiProductsAuthorities(request *RemoveApiProductsAuthoritiesRequest) (_result *RemoveApiProductsAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveApiProductsAuthoritiesResponse{}
	_body, _err := client.RemoveApiProductsAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on multiple APIs from a specified application.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveApisAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveApisAuthoritiesResponse
func (client *Client) RemoveApisAuthoritiesWithOptions(request *RemoveApisAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveApisAuthoritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveApisAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveApisAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on multiple APIs from a specified application.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveApisAuthoritiesRequest
//
// @return RemoveApisAuthoritiesResponse
func (client *Client) RemoveApisAuthorities(request *RemoveApisAuthoritiesRequest) (_result *RemoveApisAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveApisAuthoritiesResponse{}
	_body, _err := client.RemoveApisAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on a specified API from multiple applications. In this case, multiple applications map to a single API.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveAppsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppsAuthoritiesResponse
func (client *Client) RemoveAppsAuthoritiesWithOptions(request *RemoveAppsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *RemoveAppsAuthoritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAppsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAppsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on a specified API from multiple applications. In this case, multiple applications map to a single API.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- Before you revoke access permissions, check by whom the permissions were granted. API providers can only revoke permissions granted by a Provider, and API callers can only revoke permissions granted by a Consumer.
//
// @param request - RemoveAppsAuthoritiesRequest
//
// @return RemoveAppsAuthoritiesResponse
func (client *Client) RemoveAppsAuthorities(request *RemoveAppsAuthoritiesRequest) (_result *RemoveAppsAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveAppsAuthoritiesResponse{}
	_body, _err := client.RemoveAppsAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds an API from an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The unbinding takes effect immediately. After the API is unbound from the ACL, the corresponding environment does not have any IP address access control in place for the API.
//
// @param request - RemoveIpControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpControlApisResponse
func (client *Client) RemoveIpControlApisWithOptions(request *RemoveIpControlApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpControlApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveIpControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds an API from an access control list (ACL).
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- The unbinding takes effect immediately. After the API is unbound from the ACL, the corresponding environment does not have any IP address access control in place for the API.
//
// @param request - RemoveIpControlApisRequest
//
// @return RemoveIpControlApisResponse
func (client *Client) RemoveIpControlApis(request *RemoveIpControlApisRequest) (_result *RemoveIpControlApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveIpControlApisResponse{}
	_body, _err := client.RemoveIpControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more policies from an access control list (ACL).
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - RemoveIpControlPolicyItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveIpControlPolicyItemResponse
func (client *Client) RemoveIpControlPolicyItemWithOptions(request *RemoveIpControlPolicyItemRequest, runtime *dara.RuntimeOptions) (_result *RemoveIpControlPolicyItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.PolicyItemIds) {
		query["PolicyItemIds"] = request.PolicyItemIds
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveIpControlPolicyItem"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more policies from an access control list (ACL).
//
// Description:
//
//	This operation is intended for API providers.
//
// @param request - RemoveIpControlPolicyItemRequest
//
// @return RemoveIpControlPolicyItemResponse
func (client *Client) RemoveIpControlPolicyItem(request *RemoveIpControlPolicyItemRequest) (_result *RemoveIpControlPolicyItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveIpControlPolicyItemResponse{}
	_body, _err := client.RemoveIpControlPolicyItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a backend signature key from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The operation takes effect immediately. The request sent from API Gateway to the backend service does not contain the signature string. The corresponding verification step can be removed from the backend.
//
// @param request - RemoveSignatureApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSignatureApisResponse
func (client *Client) RemoveSignatureApisWithOptions(request *RemoveSignatureApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveSignatureApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveSignatureApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveSignatureApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a backend signature key from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- The operation takes effect immediately. The request sent from API Gateway to the backend service does not contain the signature string. The corresponding verification step can be removed from the backend.
//
// @param request - RemoveSignatureApisRequest
//
// @return RemoveSignatureApisResponse
func (client *Client) RemoveSignatureApis(request *RemoveSignatureApisRequest) (_result *RemoveSignatureApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveSignatureApisResponse{}
	_body, _err := client.RemoveSignatureApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a specified throttling policy from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to unbind a specified throttling policy from up to 100 APIs at a time.
//
// @param request - RemoveTrafficControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTrafficControlApisResponse
func (client *Client) RemoveTrafficControlApisWithOptions(request *RemoveTrafficControlApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveTrafficControlApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTrafficControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTrafficControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a specified throttling policy from APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to unbind a specified throttling policy from up to 100 APIs at a time.
//
// @param request - RemoveTrafficControlApisRequest
//
// @return RemoveTrafficControlApisResponse
func (client *Client) RemoveTrafficControlApis(request *RemoveTrafficControlApisRequest) (_result *RemoveTrafficControlApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTrafficControlApisResponse{}
	_body, _err := client.RemoveTrafficControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a VPC authorization without unpublishing the associated APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Revokes the permissions of API Gateway to access your VPC instance.
//
// >  Deleting an authorization affects the associated API. Before you delete the authorization, make sure that it is not used by the API.
//
// @param request - RemoveVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVpcAccessResponse
func (client *Client) RemoveVpcAccessWithOptions(request *RemoveVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *RemoveVpcAccessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVpcAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPC authorization without unpublishing the associated APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- Revokes the permissions of API Gateway to access your VPC instance.
//
// >  Deleting an authorization affects the associated API. Before you delete the authorization, make sure that it is not used by the API.
//
// @param request - RemoveVpcAccessRequest
//
// @return RemoveVpcAccessResponse
func (client *Client) RemoveVpcAccess(request *RemoveVpcAccessRequest) (_result *RemoveVpcAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveVpcAccessResponse{}
	_body, _err := client.RemoveVpcAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除VPC授权并下线关联API
//
// @param request - RemoveVpcAccessAndAbolishApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVpcAccessAndAbolishApisResponse
func (client *Client) RemoveVpcAccessAndAbolishApisWithOptions(request *RemoveVpcAccessAndAbolishApisRequest, runtime *dara.RuntimeOptions) (_result *RemoveVpcAccessAndAbolishApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedBatchWork) {
		query["NeedBatchWork"] = request.NeedBatchWork
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVpcAccessAndAbolishApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVpcAccessAndAbolishApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除VPC授权并下线关联API
//
// @param request - RemoveVpcAccessAndAbolishApisRequest
//
// @return RemoveVpcAccessAndAbolishApisResponse
func (client *Client) RemoveVpcAccessAndAbolishApis(request *RemoveVpcAccessAndAbolishApisRequest) (_result *RemoveVpcAccessAndAbolishApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveVpcAccessAndAbolishApisResponse{}
	_body, _err := client.RemoveVpcAccessAndAbolishApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the AppCode of an application. You can call this operation only once per minute.
//
// @param request - ResetAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAppCodeResponse
func (client *Client) ResetAppCodeWithOptions(request *ResetAppCodeRequest, runtime *dara.RuntimeOptions) (_result *ResetAppCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppCode) {
		query["AppCode"] = request.AppCode
	}

	if !dara.IsNil(request.NewAppCode) {
		query["NewAppCode"] = request.NewAppCode
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the AppCode of an application. You can call this operation only once per minute.
//
// @param request - ResetAppCodeRequest
//
// @return ResetAppCodeResponse
func (client *Client) ResetAppCode(request *ResetAppCodeRequest) (_result *ResetAppCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAppCodeResponse{}
	_body, _err := client.ResetAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the key of an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A new secret is automatically generated after you have called this operation. This secret cannot be customized.
//
//		- The results returned by this operation do not contain the application secret. You can obtain the secret by calling DescribeAppSecurity.
//
// @param request - ResetAppSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAppSecretResponse
func (client *Client) ResetAppSecretWithOptions(request *ResetAppSecretRequest, runtime *dara.RuntimeOptions) (_result *ResetAppSecretResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		query["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.NewAppKey) {
		query["NewAppKey"] = request.NewAppKey
	}

	if !dara.IsNil(request.NewAppSecret) {
		query["NewAppSecret"] = request.NewAppSecret
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAppSecret"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the key of an application.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A new secret is automatically generated after you have called this operation. This secret cannot be customized.
//
//		- The results returned by this operation do not contain the application secret. You can obtain the secret by calling DescribeAppSecurity.
//
// @param request - ResetAppSecretRequest
//
// @return ResetAppSecretResponse
func (client *Client) ResetAppSecret(request *ResetAppSecretRequest) (_result *ResetAppSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAppSecretResponse{}
	_body, _err := client.ResetAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据APP生成SDK
//
// @param request - SdkGenerateByAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByAppResponse
func (client *Client) SdkGenerateByAppWithOptions(request *SdkGenerateByAppRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByAppResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByApp"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据APP生成SDK
//
// @param request - SdkGenerateByAppRequest
//
// @return SdkGenerateByAppResponse
func (client *Client) SdkGenerateByApp(request *SdkGenerateByAppRequest) (_result *SdkGenerateByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdkGenerateByAppResponse{}
	_body, _err := client.SdkGenerateByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成与App关联的API的SDK
//
// @param request - SdkGenerateByAppForRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByAppForRegionResponse
func (client *Client) SdkGenerateByAppForRegionWithOptions(request *SdkGenerateByAppForRegionRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByAppForRegionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByAppForRegion"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByAppForRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成与App关联的API的SDK
//
// @param request - SdkGenerateByAppForRegionRequest
//
// @return SdkGenerateByAppForRegionResponse
func (client *Client) SdkGenerateByAppForRegion(request *SdkGenerateByAppForRegionRequest) (_result *SdkGenerateByAppForRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdkGenerateByAppForRegionResponse{}
	_body, _err := client.SdkGenerateByAppForRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据分组生成SDK
//
// @param request - SdkGenerateByGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SdkGenerateByGroupResponse
func (client *Client) SdkGenerateByGroupWithOptions(request *SdkGenerateByGroupRequest, runtime *dara.RuntimeOptions) (_result *SdkGenerateByGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SdkGenerateByGroup"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SdkGenerateByGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据分组生成SDK
//
// @param request - SdkGenerateByGroupRequest
//
// @return SdkGenerateByGroupResponse
func (client *Client) SdkGenerateByGroup(request *SdkGenerateByGroupRequest) (_result *SdkGenerateByGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SdkGenerateByGroupResponse{}
	_body, _err := client.SdkGenerateByGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改访问控制策略的名称
//
// @param request - SetAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccessControlListAttributeResponse
func (client *Client) SetAccessControlListAttributeWithOptions(request *SetAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetAccessControlListAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAccessControlListAttribute"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAccessControlListAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改访问控制策略的名称
//
// @param request - SetAccessControlListAttributeRequest
//
// @return SetAccessControlListAttributeResponse
func (client *Client) SetAccessControlListAttribute(request *SetAccessControlListAttributeRequest) (_result *SetAccessControlListAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAccessControlListAttributeResponse{}
	_body, _err := client.SetAccessControlListAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on API products to an application.
//
// @param tmpReq - SetApiProductsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApiProductsAuthoritiesResponse
func (client *Client) SetApiProductsAuthoritiesWithOptions(tmpReq *SetApiProductsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetApiProductsAuthoritiesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetApiProductsAuthoritiesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiProductIds) {
		request.ApiProductIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiProductIds, dara.String("ApiProductIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductIdsShrink) {
		query["ApiProductIds"] = request.ApiProductIdsShrink
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApiProductsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApiProductsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on API products to an application.
//
// @param request - SetApiProductsAuthoritiesRequest
//
// @return SetApiProductsAuthoritiesResponse
func (client *Client) SetApiProductsAuthorities(request *SetApiProductsAuthoritiesRequest) (_result *SetApiProductsAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApiProductsAuthoritiesResponse{}
	_body, _err := client.SetApiProductsAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a specified application to call multiple APIs.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetApisAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetApisAuthoritiesResponse
func (client *Client) SetApisAuthoritiesWithOptions(request *SetApisAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetApisAuthoritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetApisAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetApisAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a specified application to call multiple APIs.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetApisAuthoritiesRequest
//
// @return SetApisAuthoritiesResponse
func (client *Client) SetApisAuthorities(request *SetApisAuthoritiesRequest) (_result *SetApisAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetApisAuthoritiesResponse{}
	_body, _err := client.SetApisAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes multiple applications to call APIs in an API product.
//
// @param request - SetAppsAuthToApiProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppsAuthToApiProductResponse
func (client *Client) SetAppsAuthToApiProductWithOptions(request *SetAppsAuthToApiProductRequest, runtime *dara.RuntimeOptions) (_result *SetAppsAuthToApiProductResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiProductId) {
		query["ApiProductId"] = request.ApiProductId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppsAuthToApiProduct"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppsAuthToApiProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes multiple applications to call APIs in an API product.
//
// @param request - SetAppsAuthToApiProductRequest
//
// @return SetAppsAuthToApiProductResponse
func (client *Client) SetAppsAuthToApiProduct(request *SetAppsAuthToApiProductRequest) (_result *SetAppsAuthToApiProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAppsAuthToApiProductResponse{}
	_body, _err := client.SetAppsAuthToApiProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants access permissions on a specified API to multiple applications.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetAppsAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAppsAuthoritiesResponse
func (client *Client) SetAppsAuthoritiesWithOptions(request *SetAppsAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *SetAppsAuthoritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AuthValidTime) {
		query["AuthValidTime"] = request.AuthValidTime
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAppsAuthorities"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAppsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants access permissions on a specified API to multiple applications.
//
// Description:
//
//	  This operation is intended for API providers and callers.
//
//		- API providers can authorize all applications to call their APIs.
//
//		- API callers can authorize their own applications to call the APIs that they have purchased.
//
// @param request - SetAppsAuthoritiesRequest
//
// @return SetAppsAuthoritiesResponse
func (client *Client) SetAppsAuthorities(request *SetAppsAuthoritiesRequest) (_result *SetAppsAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAppsAuthoritiesResponse{}
	_body, _err := client.SetAppsAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a custom domain name to a specified API group.
//
// @param request - SetDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainResponse
func (client *Client) SetDomainWithOptions(request *SetDomainRequest, runtime *dara.RuntimeOptions) (_result *SetDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindStageName) {
		query["BindStageName"] = request.BindStageName
	}

	if !dara.IsNil(request.CustomDomainType) {
		query["CustomDomainType"] = request.CustomDomainType
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IsForce) {
		query["IsForce"] = request.IsForce
	}

	if !dara.IsNil(request.IsHttpRedirectToHttps) {
		query["IsHttpRedirectToHttps"] = request.IsHttpRedirectToHttps
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomain"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a custom domain name to a specified API group.
//
// @param request - SetDomainRequest
//
// @return SetDomainResponse
func (client *Client) SetDomain(request *SetDomainRequest) (_result *SetDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainResponse{}
	_body, _err := client.SetDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an SSL certificate for a specified custom domain name.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The SSL certificate must match the custom domain name.
//
//		- After the SSL certificate is bound, HTTPS-based API services become available.
//
// @param request - SetDomainCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainCertificateResponse
func (client *Client) SetDomainCertificateWithOptions(request *SetDomainCertificateRequest, runtime *dara.RuntimeOptions) (_result *SetDomainCertificateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaCertificateBody) {
		query["CaCertificateBody"] = request.CaCertificateBody
	}

	if !dara.IsNil(request.CertificateBody) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !dara.IsNil(request.CertificateName) {
		query["CertificateName"] = request.CertificateName
	}

	if !dara.IsNil(request.CertificatePrivateKey) {
		query["CertificatePrivateKey"] = request.CertificatePrivateKey
	}

	if !dara.IsNil(request.ClientCertSDnPassThrough) {
		query["ClientCertSDnPassThrough"] = request.ClientCertSDnPassThrough
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SslOcspCacheEnable) {
		query["SslOcspCacheEnable"] = request.SslOcspCacheEnable
	}

	if !dara.IsNil(request.SslOcspEnable) {
		query["SslOcspEnable"] = request.SslOcspEnable
	}

	if !dara.IsNil(request.SslVerifyDepth) {
		query["SslVerifyDepth"] = request.SslVerifyDepth
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainCertificate"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an SSL certificate for a specified custom domain name.
//
// Description:
//
//	  This operation is intended for API providers.
//
//		- The SSL certificate must match the custom domain name.
//
//		- After the SSL certificate is bound, HTTPS-based API services become available.
//
// @param request - SetDomainCertificateRequest
//
// @return SetDomainCertificateResponse
func (client *Client) SetDomainCertificate(request *SetDomainCertificateRequest) (_result *SetDomainCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainCertificateResponse{}
	_body, _err := client.SetDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables WebSocket for a custom domain name.
//
// @param request - SetDomainWebSocketStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainWebSocketStatusResponse
func (client *Client) SetDomainWebSocketStatusWithOptions(request *SetDomainWebSocketStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDomainWebSocketStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionValue) {
		query["ActionValue"] = request.ActionValue
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WSSEnable) {
		query["WSSEnable"] = request.WSSEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainWebSocketStatus"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainWebSocketStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables WebSocket for a custom domain name.
//
// @param request - SetDomainWebSocketStatusRequest
//
// @return SetDomainWebSocketStatusResponse
func (client *Client) SetDomainWebSocketStatus(request *SetDomainWebSocketStatusRequest) (_result *SetDomainWebSocketStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainWebSocketStatusResponse{}
	_body, _err := client.SetDomainWebSocketStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置分组授权AppCode
//
// @param request - SetGroupAuthAppCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetGroupAuthAppCodeResponse
func (client *Client) SetGroupAuthAppCodeWithOptions(request *SetGroupAuthAppCodeRequest, runtime *dara.RuntimeOptions) (_result *SetGroupAuthAppCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthAppCode) {
		query["AuthAppCode"] = request.AuthAppCode
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetGroupAuthAppCode"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetGroupAuthAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置分组授权AppCode
//
// @param request - SetGroupAuthAppCodeRequest
//
// @return SetGroupAuthAppCodeResponse
func (client *Client) SetGroupAuthAppCode(request *SetGroupAuthAppCodeRequest) (_result *SetGroupAuthAppCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetGroupAuthAppCodeResponse{}
	_body, _err := client.SetGroupAuthAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a binding relationship between specified access control lists (ACLs) and APIs.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A maximum of 100 APIs can be bound at a time.
//
// @param request - SetIpControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetIpControlApisResponse
func (client *Client) SetIpControlApisWithOptions(request *SetIpControlApisRequest, runtime *dara.RuntimeOptions) (_result *SetIpControlApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.IpControlId) {
		query["IpControlId"] = request.IpControlId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetIpControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetIpControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a binding relationship between specified access control lists (ACLs) and APIs.
//
// Description:
//
//	  This operation is intended for API callers.
//
//		- A maximum of 100 APIs can be bound at a time.
//
// @param request - SetIpControlApisRequest
//
// @return SetIpControlApisResponse
func (client *Client) SetIpControlApis(request *SetIpControlApisRequest) (_result *SetIpControlApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetIpControlApisResponse{}
	_body, _err := client.SetIpControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a signature key to APIs.
//
// @param request - SetSignatureApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSignatureApisResponse
func (client *Client) SetSignatureApisWithOptions(request *SetSignatureApisRequest, runtime *dara.RuntimeOptions) (_result *SetSignatureApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SignatureId) {
		query["SignatureId"] = request.SignatureId
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSignatureApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSignatureApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a signature key to APIs.
//
// @param request - SetSignatureApisRequest
//
// @return SetSignatureApisResponse
func (client *Client) SetSignatureApis(request *SetSignatureApisRequest) (_result *SetSignatureApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSignatureApisResponse{}
	_body, _err := client.SetSignatureApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a throttling policy to APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to bind a specific throttling policy to up to 100 APIs at a time.
//
// @param request - SetTrafficControlApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTrafficControlApisResponse
func (client *Client) SetTrafficControlApisWithOptions(request *SetTrafficControlApisRequest, runtime *dara.RuntimeOptions) (_result *SetTrafficControlApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiIds) {
		query["ApiIds"] = request.ApiIds
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	if !dara.IsNil(request.TrafficControlId) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTrafficControlApis"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTrafficControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a throttling policy to APIs.
//
// Description:
//
//	  This API is intended for API providers.
//
//		- This API allows you to bind a specific throttling policy to up to 100 APIs at a time.
//
// @param request - SetTrafficControlApisRequest
//
// @return SetTrafficControlApisResponse
func (client *Client) SetTrafficControlApis(request *SetTrafficControlApisRequest) (_result *SetTrafficControlApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetTrafficControlApisResponse{}
	_body, _err := client.SetTrafficControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) access authorization and enables reverse access.
//
// Description:
//
// This operation is intended for API providers.
//
//   - This operation is used to authorize API Gateway to access your VPC instance.
//
// @param request - SetVpcAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVpcAccessResponse
func (client *Client) SetVpcAccessWithOptions(request *SetVpcAccessRequest, runtime *dara.RuntimeOptions) (_result *SetVpcAccessResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcTargetHostName) {
		query["VpcTargetHostName"] = request.VpcTargetHostName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVpcAccess"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVpcAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a virtual private cloud (VPC) access authorization and enables reverse access.
//
// Description:
//
// This operation is intended for API providers.
//
//   - This operation is used to authorize API Gateway to access your VPC instance.
//
// @param request - SetVpcAccessRequest
//
// @return SetVpcAccessResponse
func (client *Client) SetVpcAccess(request *SetVpcAccessRequest) (_result *SetVpcAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetVpcAccessResponse{}
	_body, _err := client.SetVpcAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a wildcard domain name template for a bound custom domain name.
//
// @param request - SetWildcardDomainPatternsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWildcardDomainPatternsResponse
func (client *Client) SetWildcardDomainPatternsWithOptions(request *SetWildcardDomainPatternsRequest, runtime *dara.RuntimeOptions) (_result *SetWildcardDomainPatternsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WildcardDomainPatterns) {
		query["WildcardDomainPatterns"] = request.WildcardDomainPatterns
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWildcardDomainPatterns"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWildcardDomainPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a wildcard domain name template for a bound custom domain name.
//
// @param request - SetWildcardDomainPatternsRequest
//
// @return SetWildcardDomainPatternsResponse
func (client *Client) SetWildcardDomainPatterns(request *SetWildcardDomainPatternsRequest) (_result *SetWildcardDomainPatternsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetWildcardDomainPatternsResponse{}
	_body, _err := client.SetWildcardDomainPatternsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches the definition of an API in a specified runtime environment to a historical version.
//
// @param request - SwitchApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchApiResponse
func (client *Client) SwitchApiWithOptions(request *SwitchApiRequest, runtime *dara.RuntimeOptions) (_result *SwitchApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryVersion) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StageName) {
		query["StageName"] = request.StageName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchApi"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the definition of an API in a specified runtime environment to a historical version.
//
// @param request - SwitchApiRequest
//
// @return SwitchApiResponse
func (client *Client) SwitchApi(request *SwitchApiRequest) (_result *SwitchApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchApiResponse{}
	_body, _err := client.SwitchApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tag-resource relationship.
//
// Description:
//
//	  All tags (key-value pairs) are applied to all resources of a specified ResourceId, with each resource specified as ResourceId.N.
//
//		- Tag.N is a resource tag consisting of a key-value pair: Tag.N.Key and Tag.N.Value.
//
//		- If you call this operation to tag multiple resources simultaneously, either all or none of the resources will be tagged.
//
//		- If you specify Tag.1.Value in addition to required parameters, you must also specify Tag.1.Key. Otherwise, an InvalidParameter.TagKey error is reported. A tag that has a value must have the corresponding key, but the key can be an empty string.
//
//		- If a tag with the same key has been bound to a resource, the new tag will overwrite the existing one.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2016-07-14"),
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
// Creates a tag-resource relationship.
//
// Description:
//
//	  All tags (key-value pairs) are applied to all resources of a specified ResourceId, with each resource specified as ResourceId.N.
//
//		- Tag.N is a resource tag consisting of a key-value pair: Tag.N.Key and Tag.N.Value.
//
//		- If you call this operation to tag multiple resources simultaneously, either all or none of the resources will be tagged.
//
//		- If you specify Tag.1.Value in addition to required parameters, you must also specify Tag.1.Key. Otherwise, an InvalidParameter.TagKey error is reported. A tag that has a value must have the corresponding key, but the key can be an empty string.
//
//		- If a tag with the same key has been bound to a resource, the new tag will overwrite the existing one.
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
// Removes tags from resources.
//
// Description:
//
//	  If you call this operation to untag multiple resources simultaneously, either all or none of the resources will be untagged.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to true, all tags bound to the specified resources will be deleted. If a resource does not have any tags, the request is not processed but a success is returned.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to false, the request is not processed but a success is returned.
//
//		- When tag keys are specified, the All parameter is invalid.
//
//		- When multiple resources and key-value pairs are specified, the specified tags bound to the resources are deleted.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2016-07-14"),
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
// Removes tags from resources.
//
// Description:
//
//	  If you call this operation to untag multiple resources simultaneously, either all or none of the resources will be untagged.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to true, all tags bound to the specified resources will be deleted. If a resource does not have any tags, the request is not processed but a success is returned.
//
//		- If you specify resource IDs without specifying tag keys and set the All parameter to false, the request is not processed but a success is returned.
//
//		- When tag keys are specified, the All parameter is invalid.
//
//		- When multiple resources and key-value pairs are specified, the specified tags bound to the resources are deleted.
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
// Modifies an internal domain name resolution.
//
// @param tmpReq - UpdatePrivateDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateDNSResponse
func (client *Client) UpdatePrivateDNSWithOptions(tmpReq *UpdatePrivateDNSRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateDNSResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePrivateDNSShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Records) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, dara.String("Records"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IntranetDomain) {
		query["IntranetDomain"] = request.IntranetDomain
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.RecordsShrink) {
		body["Records"] = request.RecordsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateDNS"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateDNSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an internal domain name resolution.
//
// @param request - UpdatePrivateDNSRequest
//
// @return UpdatePrivateDNSResponse
func (client *Client) UpdatePrivateDNS(request *UpdatePrivateDNSRequest) (_result *UpdatePrivateDNSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePrivateDNSResponse{}
	_body, _err := client.UpdatePrivateDNSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the network connectivity between an API Gateway instance and a port on an Elastic Compute Service (ECS) or Server Load Balance (SLB) instance in a virtual private cloud (VPC) access authorization.
//
// @param request - ValidateVpcConnectivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateVpcConnectivityResponse
func (client *Client) ValidateVpcConnectivityWithOptions(request *ValidateVpcConnectivityRequest, runtime *dara.RuntimeOptions) (_result *ValidateVpcConnectivityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.VpcAccessId) {
		query["VpcAccessId"] = request.VpcAccessId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateVpcConnectivity"),
		Version:     dara.String("2016-07-14"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateVpcConnectivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the network connectivity between an API Gateway instance and a port on an Elastic Compute Service (ECS) or Server Load Balance (SLB) instance in a virtual private cloud (VPC) access authorization.
//
// @param request - ValidateVpcConnectivityRequest
//
// @return ValidateVpcConnectivityResponse
func (client *Client) ValidateVpcConnectivity(request *ValidateVpcConnectivityRequest) (_result *ValidateVpcConnectivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateVpcConnectivityResponse{}
	_body, _err := client.ValidateVpcConnectivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
