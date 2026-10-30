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
		"cn-qingdao":                  dara.String("ddosbgp.aliyuncs.com"),
		"cn-beijing":                  dara.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote":                dara.String("ddosbgp.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai":                 dara.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("ddosbgp.aliyuncs.com"),
		"cn-heyuan":                   dara.String("ddosbgp.aliyuncs.com"),
		"cn-guangzhou":                dara.String("ddosbgp.aliyuncs.com"),
		"cn-chengdu":                  dara.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-1":              dara.String("ddosbgp.aliyuncs.com"),
		"ap-southeast-2":              dara.String("ddosbgp.aliyuncs.com"),
		"ap-southeast-3":              dara.String("ddosbgp.aliyuncs.com"),
		"ap-southeast-5":              dara.String("ddosbgp.aliyuncs.com"),
		"eu-west-1":                   dara.String("ddosbgp.aliyuncs.com"),
		"eu-central-1":                dara.String("ddosbgp.aliyuncs.com"),
		"ap-south-1":                  dara.String("ddosbgp.aliyuncs.com"),
		"me-east-1":                   dara.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("ddosbgp.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("ddosbgp.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("ddosbgp.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("ddosbgp.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("ddosbgp.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("ddosbgp.aliyuncs.com"),
		"cn-edge-1":                   dara.String("ddosbgp.aliyuncs.com"),
		"cn-fujian":                   dara.String("ddosbgp.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("ddosbgp.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("ddosbgp.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("ddosbgp.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("ddosbgp.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("ddosbgp.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("ddosbgp.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("ddosbgp.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("ddosbgp.aliyuncs.com"),
		"cn-wuhan":                    dara.String("ddosbgp.aliyuncs.com"),
		"cn-yushanfang":               dara.String("ddosbgp.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("ddosbgp.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("ddosbgp.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("ddosbgp.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("ddosbgp.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("ddosbgp.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("ddosbgp.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ddosbgp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds protected IP addresses to an Anti-DDoS Origin instance.
//
// @param request - AddIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpResponse
func (client *Client) AddIpWithOptions(request *AddIpRequest, runtime *dara.RuntimeOptions) (_result *AddIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddIp"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds protected IP addresses to an Anti-DDoS Origin instance.
//
// @param request - AddIpRequest
//
// @return AddIpResponse
func (client *Client) AddIp(request *AddIpRequest) (_result *AddIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddIpResponse{}
	_body, _err := client.AddIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds members to a resource directory.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to add members to the resource directory.
//
// @param tmpReq - AddRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRdMemberListResponse
func (client *Client) AddRdMemberListWithOptions(tmpReq *AddRdMemberListRequest, runtime *dara.RuntimeOptions) (_result *AddRdMemberListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddRdMemberListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MemberList) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, dara.String("MemberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberListShrink) {
		query["MemberList"] = request.MemberListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRdMemberList"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds members to a resource directory.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to add members to the resource directory.
//
// @param request - AddRdMemberListRequest
//
// @return AddRdMemberListResponse
func (client *Client) AddRdMemberList(request *AddRdMemberListRequest) (_result *AddRdMemberListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRdMemberListResponse{}
	_body, _err := client.AddRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an asset with an Anti-DDoS Origin instance of a paid edition.
//
// @param tmpReq - AttachAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachAssetGroupToInstanceResponse
func (client *Client) AttachAssetGroupToInstanceWithOptions(tmpReq *AttachAssetGroupToInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetGroupList) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, dara.String("AssetGroupList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetGroupListShrink) {
		query["AssetGroupList"] = request.AssetGroupListShrink
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
		Action:      dara.String("AttachAssetGroupToInstance"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an asset with an Anti-DDoS Origin instance of a paid edition.
//
// @param request - AttachAssetGroupToInstanceRequest
//
// @return AttachAssetGroupToInstanceResponse
func (client *Client) AttachAssetGroupToInstance(request *AttachAssetGroupToInstanceRequest) (_result *AttachAssetGroupToInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachAssetGroupToInstanceResponse{}
	_body, _err := client.AttachAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates protection objects with a mitigation policy.
//
// Description:
//
// A mitigation policy that is associated with protection objects cannot be deleted.
//
// @param tmpReq - AttachToPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachToPolicyResponse
func (client *Client) AttachToPolicyWithOptions(tmpReq *AttachToPolicyRequest, runtime *dara.RuntimeOptions) (_result *AttachToPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachToPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpPortProtocolList) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, dara.String("IpPortProtocolList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IpPortProtocolListShrink) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachToPolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachToPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates protection objects with a mitigation policy.
//
// Description:
//
// A mitigation policy that is associated with protection objects cannot be deleted.
//
// @param request - AttachToPolicyRequest
//
// @return AttachToPolicyResponse
func (client *Client) AttachToPolicy(request *AttachToPolicyRequest) (_result *AttachToPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachToPolicyResponse{}
	_body, _err := client.AttachToPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to access Simple Log Service.
//
// @param request - CheckAccessLogAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAccessLogAuthResponse
func (client *Client) CheckAccessLogAuthWithOptions(request *CheckAccessLogAuthRequest, runtime *dara.RuntimeOptions) (_result *CheckAccessLogAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("CheckAccessLogAuth"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to access Simple Log Service.
//
// @param request - CheckAccessLogAuthRequest
//
// @return CheckAccessLogAuthResponse
func (client *Client) CheckAccessLogAuth(request *CheckAccessLogAuthRequest) (_result *CheckAccessLogAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAccessLogAuthResponse{}
	_body, _err := client.CheckAccessLogAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// Description:
//
// You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckGrantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckGrantResponse
func (client *Client) CheckGrantWithOptions(request *CheckGrantRequest, runtime *dara.RuntimeOptions) (_result *CheckGrantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckGrant"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckGrantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// Description:
//
// You can call the CheckGrant operation to query whether Anti-DDoS Origin is authorized to obtain information about the assets within the current Alibaba Cloud account.
//
// ### Limits
//
// You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - CheckGrantRequest
//
// @return CheckGrantResponse
func (client *Client) CheckGrant(request *CheckGrantRequest) (_result *CheckGrantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckGrantResponse{}
	_body, _err := client.CheckGrantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mitigation policy.
//
// Description:
//
// A mitigation policy that is associated with protected objects cannot be deleted.
//
// @param request - CreatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mitigation policy.
//
// Description:
//
// A mitigation policy that is associated with protected objects cannot be deleted.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering status for a single protected IP address.
//
// Description:
//
// This operation deactivates blackhole filtering status for a single IP address protected by an Anti-DDoS Origin instance.
//
// Before you invoke this operation, you can invoke [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) to query the protection status of all IP assets protected by Anti-DDoS Origin instances, including whether the IP addresses are in blackhole filtering status.
//
// ### QPS limit
//
// The QPS limit on this operation is 10 requests per second per account. If the number of requests per second exceeds the limit, throttling is triggered, which may affect your business. Invoke the API as appropriate.
//
// @param request - DeleteBlackholeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBlackholeResponse
func (client *Client) DeleteBlackholeWithOptions(request *DeleteBlackholeRequest, runtime *dara.RuntimeOptions) (_result *DeleteBlackholeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
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
		Action:      dara.String("DeleteBlackhole"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deactivates blackhole filtering status for a single protected IP address.
//
// Description:
//
// This operation deactivates blackhole filtering status for a single IP address protected by an Anti-DDoS Origin instance.
//
// Before you invoke this operation, you can invoke [DescribePackIpList](https://help.aliyun.com/document_detail/118701.html) to query the protection status of all IP assets protected by Anti-DDoS Origin instances, including whether the IP addresses are in blackhole filtering status.
//
// ### QPS limit
//
// The QPS limit on this operation is 10 requests per second per account. If the number of requests per second exceeds the limit, throttling is triggered, which may affect your business. Invoke the API as appropriate.
//
// @param request - DeleteBlackholeRequest
//
// @return DeleteBlackholeResponse
func (client *Client) DeleteBlackhole(request *DeleteBlackholeRequest) (_result *DeleteBlackholeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBlackholeResponse{}
	_body, _err := client.DeleteBlackholeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes protected IP addresses from Anti-DDoS Origin and cancels protection.
//
// @param request - DeleteIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpResponse
func (client *Client) DeleteIpWithOptions(request *DeleteIpRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIp"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes protected IP addresses from Anti-DDoS Origin and cancels protection.
//
// @param request - DeleteIpRequest
//
// @return DeleteIpResponse
func (client *Client) DeleteIp(request *DeleteIpRequest) (_result *DeleteIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpResponse{}
	_body, _err := client.DeleteIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a mitigation policy.
//
// Description:
//
// You cannot delete a mitigation policy that is associated with a protected object.
//
// @param request - DeletePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(request *DeletePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a mitigation policy.
//
// Description:
//
// You cannot delete a mitigation policy that is associated with a protected object.
//
// @param request - DeletePolicyRequest
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(request *DeletePolicyRequest) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes members.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to delete members.
//
// @param tmpReq - DeleteRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRdMemberListResponse
func (client *Client) DeleteRdMemberListWithOptions(tmpReq *DeleteRdMemberListRequest, runtime *dara.RuntimeOptions) (_result *DeleteRdMemberListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteRdMemberListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MemberList) {
		request.MemberListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberList, dara.String("MemberList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberListShrink) {
		query["MemberList"] = request.MemberListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRdMemberList"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes members.
//
// Description:
//
// Only a delegated administrator account or the management account of a resource directory can be used to delete members.
//
// @param request - DeleteRdMemberListRequest
//
// @return DeleteRdMemberListResponse
func (client *Client) DeleteRdMemberList(request *DeleteRdMemberListRequest) (_result *DeleteRdMemberListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRdMemberListResponse{}
	_body, _err := client.DeleteRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetGroupResponse
func (client *Client) DescribeAssetGroupWithOptions(request *DescribeAssetGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAssetGroup"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupRequest
//
// @return DescribeAssetGroupResponse
func (client *Client) DescribeAssetGroup(request *DescribeAssetGroupRequest) (_result *DescribeAssetGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAssetGroupResponse{}
	_body, _err := client.DescribeAssetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAssetGroupToInstanceResponse
func (client *Client) DescribeAssetGroupToInstanceWithOptions(request *DescribeAssetGroupToInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAssetGroupToInstance"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an asset and an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DescribeAssetGroupToInstanceRequest
//
// @return DescribeAssetGroupToInstanceResponse
func (client *Client) DescribeAssetGroupToInstance(request *DescribeAssetGroupToInstanceRequest) (_result *DescribeAssetGroupToInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAssetGroupToInstanceResponse{}
	_body, _err := client.DescribeAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of DDoS attack events that occurred on a single Anti-DDoS Origin instance.
//
// Description:
//
// You can invoke this operation to perform a paged query for the details of DDoS attack events that occurred on a single Anti-DDoS Origin instance. The paging result includes the start and end time of the attack, the attacked IP address, and the event status.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param request - DescribeDdosEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosEventResponse
func (client *Client) DescribeDdosEventWithOptions(request *DescribeDdosEventRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDdosEvent"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of DDoS attack events that occurred on a single Anti-DDoS Origin instance.
//
// Description:
//
// You can invoke this operation to perform a paged query for the details of DDoS attack events that occurred on a single Anti-DDoS Origin instance. The paging result includes the start and end time of the attack, the attacked IP address, and the event status.
//
// ### QPS limits
//
// The QPS limit per user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Make API calls at a reasonable rate.
//
// @param request - DescribeDdosEventRequest
//
// @return DescribeDdosEventResponse
func (client *Client) DescribeDdosEvent(request *DescribeDdosEventRequest) (_result *DescribeDdosEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosEventResponse{}
	_body, _err := client.DescribeDdosEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pay-as-you-go billing information of an Anti-DDoS Origin instance.
//
// Description:
//
// This operation is used to perform a paged query of the billing details of all Anti-DDoS Origin instances owned by the current Alibaba Cloud account. The billing details include instance IDs, validity periods, and statuses. Paging is supported for this query.
//
// @param request - DescribeDdosOriginInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDdosOriginInstanceBillResponse
func (client *Client) DescribeDdosOriginInstanceBillWithOptions(request *DescribeDdosOriginInstanceBillRequest, runtime *dara.RuntimeOptions) (_result *DescribeDdosOriginInstanceBillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IsShowList) {
		query["IsShowList"] = request.IsShowList
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
		Action:      dara.String("DescribeDdosOriginInstanceBill"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDdosOriginInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pay-as-you-go billing information of an Anti-DDoS Origin instance.
//
// Description:
//
// This operation is used to perform a paged query of the billing details of all Anti-DDoS Origin instances owned by the current Alibaba Cloud account. The billing details include instance IDs, validity periods, and statuses. Paging is supported for this query.
//
// @param request - DescribeDdosOriginInstanceBillRequest
//
// @return DescribeDdosOriginInstanceBillResponse
func (client *Client) DescribeDdosOriginInstanceBill(request *DescribeDdosOriginInstanceBillRequest) (_result *DescribeDdosOriginInstanceBillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDdosOriginInstanceBillResponse{}
	_body, _err := client.DescribeDdosOriginInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are assigned public IP addresses in abnormal states and the number of Anti-DDoS Origin instances that are about to expire in a specified region.
//
// Description:
//
// This operation queries the number of abnormal IP addresses (such as those in the Black Hole Activated state) among protected assets that are assigned public IP addresses (such as a cloud server ECS instance, a Server Load Balancer (SLB) instance, and an elastic IP address (EIP)) and the number of Anti-DDoS Origin instances that are about to expire (with a remaining validity period of less than 7 calendar days) in a specified region.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you invoke the API and exceed the limit, your requests are throttled, which may affect your services. Invoke the API at a reasonable rate.
//
// @param request - DescribeExcpetionCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExcpetionCountResponse
func (client *Client) DescribeExcpetionCountWithOptions(request *DescribeExcpetionCountRequest, runtime *dara.RuntimeOptions) (_result *DescribeExcpetionCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeExcpetionCount"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of assets that are assigned public IP addresses in abnormal states and the number of Anti-DDoS Origin instances that are about to expire in a specified region.
//
// Description:
//
// This operation queries the number of abnormal IP addresses (such as those in the Black Hole Activated state) among protected assets that are assigned public IP addresses (such as a cloud server ECS instance, a Server Load Balancer (SLB) instance, and an elastic IP address (EIP)) and the number of Anti-DDoS Origin instances that are about to expire (with a remaining validity period of less than 7 calendar days) in a specified region.
//
// ### QPS limit
//
// The queries per second (QPS) limit for a single user is 10 calls per second. If you invoke the API and exceed the limit, your requests are throttled, which may affect your services. Invoke the API at a reasonable rate.
//
// @param request - DescribeExcpetionCountRequest
//
// @return DescribeExcpetionCountResponse
func (client *Client) DescribeExcpetionCount(request *DescribeExcpetionCountRequest) (_result *DescribeExcpetionCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExcpetionCountResponse{}
	_body, _err := client.DescribeExcpetionCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of Anti-DDoS Origin instances.
//
// Description:
//
// This operation is used to query the details of all Anti-DDoS Origin instances owned by the current Alibaba Cloud account by paging, such as instance IDs, validity periods, and statuses.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - DescribeInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceListResponse
func (client *Client) DescribeInstanceListWithOptions(request *DescribeInstanceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.InstanceTypeList) {
		query["InstanceTypeList"] = request.InstanceTypeList
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.IpVersion) {
		query["IpVersion"] = request.IpVersion
	}

	if !dara.IsNil(request.Orderby) {
		query["Orderby"] = request.Orderby
	}

	if !dara.IsNil(request.Orderdire) {
		query["Orderdire"] = request.Orderdire
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceList"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of Anti-DDoS Origin instances.
//
// Description:
//
// This operation is used to query the details of all Anti-DDoS Origin instances owned by the current Alibaba Cloud account by paging, such as instance IDs, validity periods, and statuses.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API call is throttled, which may affect your business. Invoke this operation as needed.
//
// @param request - DescribeInstanceListRequest
//
// @return DescribeInstanceListResponse
func (client *Client) DescribeInstanceList(request *DescribeInstanceListRequest) (_result *DescribeInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceListResponse{}
	_body, _err := client.DescribeInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the specifications of a specified Anti-DDoS Origin instance.
//
// @param request - DescribeInstanceSpecsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecsWithOptions(request *DescribeInstanceSpecsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceSpecsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIdList) {
		query["InstanceIdList"] = request.InstanceIdList
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
		Action:      dara.String("DescribeInstanceSpecs"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the specifications of a specified Anti-DDoS Origin instance.
//
// @param request - DescribeInstanceSpecsRequest
//
// @return DescribeInstanceSpecsResponse
func (client *Client) DescribeInstanceSpecs(request *DescribeInstanceSpecsRequest) (_result *DescribeInstanceSpecsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceSpecsResponse{}
	_body, _err := client.DescribeInstanceSpecsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network-layer interception records of Anti-DDoS Origin instances.
//
// Description:
//
// This operation is used to query the details of Layer 3 and Layer 4 packet interception records for all Anti-DDoS Origin instances owned by the current Alibaba Cloud account by using paging.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API invoke is throttled, which may affect your business. Invoke this operation at an appropriate frequency.
//
// @param request - DescribeNetworkLayerInterceptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetworkLayerInterceptsResponse
func (client *Client) DescribeNetworkLayerInterceptsWithOptions(request *DescribeNetworkLayerInterceptsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetworkLayerInterceptsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationIp) {
		query["DestinationIp"] = request.DestinationIp
	}

	if !dara.IsNil(request.DestinationPort) {
		query["DestinationPort"] = request.DestinationPort
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InterceptModule) {
		query["InterceptModule"] = request.InterceptModule
	}

	if !dara.IsNil(request.NetworkProtocol) {
		query["NetworkProtocol"] = request.NetworkProtocol
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProtocolNumber) {
		query["ProtocolNumber"] = request.ProtocolNumber
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetworkLayerIntercepts"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNetworkLayerInterceptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network-layer interception records of Anti-DDoS Origin instances.
//
// Description:
//
// This operation is used to query the details of Layer 3 and Layer 4 packet interception records for all Anti-DDoS Origin instances owned by the current Alibaba Cloud account by using paging.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 10 calls per second. If the limit is exceeded, the API invoke is throttled, which may affect your business. Invoke this operation at an appropriate frequency.
//
// @param request - DescribeNetworkLayerInterceptsRequest
//
// @return DescribeNetworkLayerInterceptsResponse
func (client *Client) DescribeNetworkLayerIntercepts(request *DescribeNetworkLayerInterceptsRequest) (_result *DescribeNetworkLayerInterceptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNetworkLayerInterceptsResponse{}
	_body, _err := client.DescribeNetworkLayerInterceptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Anti-DDoS Origin.
//
// Description:
//
// This API operation is used to invoke a paged query with paging for the Operation logs of Anti-DDoS Origin.
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Invoke this API operation at a reasonable frequency.
//
// @param request - DescribeOpEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntitiesWithOptions(request *DescribeOpEntitiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeOpEntitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OpAction) {
		query["OpAction"] = request.OpAction
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderDir) {
		query["OrderDir"] = request.OrderDir
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeOpEntities"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the operation logs of Anti-DDoS Origin.
//
// Description:
//
// This API operation is used to invoke a paged query with paging for the Operation logs of Anti-DDoS Origin.
//
// ### QPS limits
//
// The QPS limit for a single user is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Invoke this API operation at a reasonable frequency.
//
// @param request - DescribeOpEntitiesRequest
//
// @return DescribeOpEntitiesResponse
func (client *Client) DescribeOpEntities(request *DescribeOpEntitiesRequest) (_result *DescribeOpEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeOpEntitiesResponse{}
	_body, _err := client.DescribeOpEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of protected IP addresses for a single Anti-DDoS Origin instance.
//
// Description:
//
// This operation performs a paged query to retrieve the list of protected IP addresses for a single Anti-DDoS Origin instance. The information includes the IP address, the cloud asset type to which the IP address belongs, and the current status of the IP address, such as the blackhole filtering status.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency.
//
// @param request - DescribePackIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackIpListResponse
func (client *Client) DescribePackIpListWithOptions(request *DescribePackIpListRequest, runtime *dara.RuntimeOptions) (_result *DescribePackIpListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.MemberUid) {
		query["MemberUid"] = request.MemberUid
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductName) {
		query["ProductName"] = request.ProductName
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
		Action:      dara.String("DescribePackIpList"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of protected IP addresses for a single Anti-DDoS Origin instance.
//
// Description:
//
// This operation performs a paged query to retrieve the list of protected IP addresses for a single Anti-DDoS Origin instance. The information includes the IP address, the cloud asset type to which the IP address belongs, and the current status of the IP address, such as the blackhole filtering status.
//
// ### QPS limit
//
// The queries per second (QPS) limit per user is 10. If the limit is exceeded, API calls are throttled, which may affect your services. Call the operation at a reasonable frequency.
//
// @param request - DescribePackIpListRequest
//
// @return DescribePackIpListResponse
func (client *Client) DescribePackIpList(request *DescribePackIpListRequest) (_result *DescribePackIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackIpListResponse{}
	_body, _err := client.DescribePackIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries members in a resource directory.
//
// @param request - DescribeRdMemberListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdMemberListResponse
func (client *Client) DescribeRdMemberListWithOptions(request *DescribeRdMemberListRequest, runtime *dara.RuntimeOptions) (_result *DescribeRdMemberListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceDirectoryId) {
		query["ResourceDirectoryId"] = request.ResourceDirectoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdMemberList"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdMemberListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries members in a resource directory.
//
// @param request - DescribeRdMemberListRequest
//
// @return DescribeRdMemberListResponse
func (client *Client) DescribeRdMemberList(request *DescribeRdMemberListRequest) (_result *DescribeRdMemberListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdMemberListResponse{}
	_body, _err := client.DescribeRdMemberListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the multi-account management feature.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRdStatusResponse
func (client *Client) DescribeRdStatusWithOptions(runtime *dara.RuntimeOptions) (_result *DescribeRdStatusResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRdStatus"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRdStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the multi-account management feature.
//
// @return DescribeRdStatusResponse
func (client *Client) DescribeRdStatus() (_result *DescribeRdStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRdStatusResponse{}
	_body, _err := client.DescribeRdStatusWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all regions supported by Anti-DDoS Origin Enterprise.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2018-07-20"),
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
// Queries all regions supported by Anti-DDoS Origin Enterprise.
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
// Queries the traffic statistics of a specified Anti-DDoS Origin instance within a specified time period.
//
// Description:
//
// Queries the traffic statistics of a specified Anti-DDoS Origin instance within a specified time range.
//
// > When calling this operation, you must set the **InstanceId*	- parameter to specify the Anti-DDoS Origin instance to query.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 1 call per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeTrafficRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTrafficResponse
func (client *Client) DescribeTrafficWithOptions(request *DescribeTrafficRequest, runtime *dara.RuntimeOptions) (_result *DescribeTrafficResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FlowType) {
		query["FlowType"] = request.FlowType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.Ipnet) {
		query["Ipnet"] = request.Ipnet
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTraffic"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the traffic statistics of a specified Anti-DDoS Origin instance within a specified time period.
//
// Description:
//
// Queries the traffic statistics of a specified Anti-DDoS Origin instance within a specified time range.
//
// > When calling this operation, you must set the **InstanceId*	- parameter to specify the Anti-DDoS Origin instance to query.
//
// ### QPS limit
//
// The single-user QPS limit for this operation is 1 call per second. If the limit is exceeded, the API call is throttled, which may affect your business. Call this operation as appropriate.
//
// @param request - DescribeTrafficRequest
//
// @return DescribeTrafficResponse
func (client *Client) DescribeTraffic(request *DescribeTrafficRequest) (_result *DescribeTrafficResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTrafficResponse{}
	_body, _err := client.DescribeTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates a mitigation policy from a protected object.
//
// @param tmpReq - DetachFromPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachFromPolicyResponse
func (client *Client) DetachFromPolicyWithOptions(tmpReq *DetachFromPolicyRequest, runtime *dara.RuntimeOptions) (_result *DetachFromPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachFromPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpPortProtocolList) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, dara.String("IpPortProtocolList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IpPortProtocolListShrink) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachFromPolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachFromPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a mitigation policy from a protected object.
//
// @param request - DetachFromPolicyRequest
//
// @return DetachFromPolicyResponse
func (client *Client) DetachFromPolicy(request *DetachFromPolicyRequest) (_result *DetachFromPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachFromPolicyResponse{}
	_body, _err := client.DetachFromPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates an asset from an Anti-DDoS Origin instance of a paid edition.
//
// @param tmpReq - DettachAssetGroupToInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DettachAssetGroupToInstanceResponse
func (client *Client) DettachAssetGroupToInstanceWithOptions(tmpReq *DettachAssetGroupToInstanceRequest, runtime *dara.RuntimeOptions) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DettachAssetGroupToInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AssetGroupList) {
		request.AssetGroupListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AssetGroupList, dara.String("AssetGroupList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetGroupListShrink) {
		query["AssetGroupList"] = request.AssetGroupListShrink
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
		Action:      dara.String("DettachAssetGroupToInstance"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates an asset from an Anti-DDoS Origin instance of a paid edition.
//
// @param request - DettachAssetGroupToInstanceRequest
//
// @return DettachAssetGroupToInstanceResponse
func (client *Client) DettachAssetGroupToInstance(request *DettachAssetGroupToInstanceRequest) (_result *DettachAssetGroupToInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DettachAssetGroupToInstanceResponse{}
	_body, _err := client.DettachAssetGroupToInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the activation status of Simple Log Service for the current Alibaba Cloud account.
//
// @param request - GetSlsOpenStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSlsOpenStatusResponse
func (client *Client) GetSlsOpenStatusWithOptions(request *GetSlsOpenStatusRequest, runtime *dara.RuntimeOptions) (_result *GetSlsOpenStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetSlsOpenStatus"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the activation status of Simple Log Service for the current Alibaba Cloud account.
//
// @param request - GetSlsOpenStatusRequest
//
// @return GetSlsOpenStatusResponse
func (client *Client) GetSlsOpenStatus(request *GetSlsOpenStatusRequest) (_result *GetSlsOpenStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSlsOpenStatusResponse{}
	_body, _err := client.GetSlsOpenStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Anti-DDoS Origin instances that have log analysis enabled.
//
// @param request - ListOpenedAccessLogInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpenedAccessLogInstancesResponse
func (client *Client) ListOpenedAccessLogInstancesWithOptions(request *ListOpenedAccessLogInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("ListOpenedAccessLogInstances"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Anti-DDoS Origin instances that have log analysis enabled.
//
// @param request - ListOpenedAccessLogInstancesRequest
//
// @return ListOpenedAccessLogInstancesResponse
func (client *Client) ListOpenedAccessLogInstances(request *ListOpenedAccessLogInstancesRequest) (_result *ListOpenedAccessLogInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOpenedAccessLogInstancesResponse{}
	_body, _err := client.ListOpenedAccessLogInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries mitigation policies.
//
// @param request - ListPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyResponse
func (client *Client) ListPolicyWithOptions(request *ListPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries mitigation policies.
//
// @param request - ListPolicyRequest
//
// @return ListPolicyResponse
func (client *Client) ListPolicy(request *ListPolicyRequest) (_result *ListPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicyResponse{}
	_body, _err := client.ListPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the associations of mitigation policies.
//
// @param tmpReq - ListPolicyAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyAttachmentResponse
func (client *Client) ListPolicyAttachmentWithOptions(tmpReq *ListPolicyAttachmentRequest, runtime *dara.RuntimeOptions) (_result *ListPolicyAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListPolicyAttachmentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpPortProtocolList) {
		request.IpPortProtocolListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpPortProtocolList, dara.String("IpPortProtocolList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IpPortProtocolListShrink) {
		query["IpPortProtocolList"] = request.IpPortProtocolListShrink
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		query["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyAttachment"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the associations of mitigation policies.
//
// @param request - ListPolicyAttachmentRequest
//
// @return ListPolicyAttachmentResponse
func (client *Client) ListPolicyAttachment(request *ListPolicyAttachmentRequest) (_result *ListPolicyAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicyAttachmentResponse{}
	_body, _err := client.ListPolicyAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all tags.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2018-07-20"),
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
// Queries all tags.
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
// Queries the mapping between resources (Anti-DDoS Origin instances) and tags.
//
// Description:
//
// This operation queries the tags bound to Anti-DDoS Origin instances in batches.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2018-07-20"),
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
// Queries the mapping between resources (Anti-DDoS Origin instances) and tags.
//
// Description:
//
// This operation queries the tags bound to Anti-DDoS Origin instances in batches.
//
// ### QPS limit
//
// The QPS limit per user for this operation is 10 calls per second. If the limit is exceeded, API calls are throttled, which may affect your services. Call this operation at a reasonable rate.
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
// Modifies a mitigation policy.
//
// Description:
//
// Modifies a mitigation policy.
//
// @param tmpReq - ModifyPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicyWithOptions(tmpReq *ModifyPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicy"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a mitigation policy.
//
// Description:
//
// Modifies a mitigation policy.
//
// @param request - ModifyPolicyRequest
//
// @return ModifyPolicyResponse
func (client *Client) ModifyPolicy(request *ModifyPolicyRequest) (_result *ModifyPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyResponse{}
	_body, _err := client.ModifyPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the content of a mitigation policy.
//
// Description:
//
// Ensure that all parameters are specified when you call this operation. If a parameter is not specified, the corresponding configuration is deleted.
//
// @param tmpReq - ModifyPolicyContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPolicyContentResponse
func (client *Client) ModifyPolicyContentWithOptions(tmpReq *ModifyPolicyContentRequest, runtime *dara.RuntimeOptions) (_result *ModifyPolicyContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyPolicyContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Content) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, dara.String("Content"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ContentShrink) {
		query["Content"] = request.ContentShrink
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PortVersion) {
		query["PortVersion"] = request.PortVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPolicyContent"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPolicyContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the content of a mitigation policy.
//
// Description:
//
// Ensure that all parameters are specified when you call this operation. If a parameter is not specified, the corresponding configuration is deleted.
//
// @param request - ModifyPolicyContentRequest
//
// @return ModifyPolicyContentResponse
func (client *Client) ModifyPolicyContent(request *ModifyPolicyContentRequest) (_result *ModifyPolicyContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPolicyContentResponse{}
	_body, _err := client.ModifyPolicyContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a remark for a single Anti-DDoS Origin instance.
//
// Description:
//
// Sets a remark for a single Anti-DDoS Origin instance.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this API appropriately.
//
// @param request - ModifyRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRemarkResponse
func (client *Client) ModifyRemarkWithOptions(request *ModifyRemarkRequest, runtime *dara.RuntimeOptions) (_result *ModifyRemarkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRemark"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a remark for a single Anti-DDoS Origin instance.
//
// Description:
//
// Sets a remark for a single Anti-DDoS Origin instance.
//
// ### QPS limit
//
// The single-user QPS limit for this API is 10 calls per second. If this limit is exceeded, the API calls are throttled, which may affect your business. Call this API appropriately.
//
// @param request - ModifyRemarkRequest
//
// @return ModifyRemarkResponse
func (client *Client) ModifyRemark(request *ModifyRemarkRequest) (_result *ModifyRemarkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyRemarkResponse{}
	_body, _err := client.ModifyRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Migrates protected IP addresses from an Anti-DDoS Origin instance.
//
// @param request - MoveIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveIpResponse
func (client *Client) MoveIpWithOptions(request *MoveIpRequest, runtime *dara.RuntimeOptions) (_result *MoveIpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
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

	if !dara.IsNil(request.SourceInstanceId) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveIp"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates protected IP addresses from an Anti-DDoS Origin instance.
//
// @param request - MoveIpRequest
//
// @return MoveIpResponse
func (client *Client) MoveIp(request *MoveIpRequest) (_result *MoveIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveIpResponse{}
	_body, _err := client.MoveIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a cloud resource belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a cloud resource belongs.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go Anti-DDoS Origin instance.
//
// @param request - ReleaseDdosOriginInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseDdosOriginInstanceResponse
func (client *Client) ReleaseDdosOriginInstanceWithOptions(request *ReleaseDdosOriginInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseDdosOriginInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseDdosOriginInstance"),
		Version:     dara.String("2018-07-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseDdosOriginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go Anti-DDoS Origin instance.
//
// @param request - ReleaseDdosOriginInstanceRequest
//
// @return ReleaseDdosOriginInstanceResponse
func (client *Client) ReleaseDdosOriginInstance(request *ReleaseDdosOriginInstanceRequest) (_result *ReleaseDdosOriginInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseDdosOriginInstanceResponse{}
	_body, _err := client.ReleaseDdosOriginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add tags to Anti-DDoS Origin instances.
//
// Description:
//
// You can call the TagResources operation to add tags to instances.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2018-07-20"),
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
// Add tags to Anti-DDoS Origin instances.
//
// Description:
//
// You can call the TagResources operation to add tags to instances.
//
// ### Limit
//
// You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
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
// Removes tags from Anti-DDoS Origin instances.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     dara.String("2018-07-20"),
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
// Removes tags from Anti-DDoS Origin instances.
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
