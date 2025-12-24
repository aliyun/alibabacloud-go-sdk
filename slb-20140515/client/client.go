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
		"cn-qingdao":                  dara.String("slb.aliyuncs.com"),
		"cn-beijing":                  dara.String("slb.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("slb.aliyuncs.com"),
		"cn-shanghai":                 dara.String("slb.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("slb.aliyuncs.com"),
		"cn-hongkong":                 dara.String("slb.aliyuncs.com"),
		"ap-southeast-1":              dara.String("slb.aliyuncs.com"),
		"us-east-1":                   dara.String("slb.aliyuncs.com"),
		"us-west-1":                   dara.String("slb.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("slb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("slb.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("slb.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("slb.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("slb.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("slb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("slb.aliyuncs.com"),
		"cn-edge-1":                   dara.String("slb.aliyuncs.com"),
		"cn-fujian":                   dara.String("slb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("slb.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("slb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("slb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("slb-api.cn-qingdao-nebula.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("slb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("slb.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("slb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("slb.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("slb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("slb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("slb.aliyuncs.com"),
		"cn-wuhan":                    dara.String("slb.aliyuncs.com"),
		"cn-yushanfang":               dara.String("slb.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("slb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("slb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("slb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("slb.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("slb.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("slb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("slb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds IP entries to an access control list (ACL).
//
// Description:
//
// Each network ACL can contain one or more IP addresses or CIDR blocks. Take note of the following limits on network ACLs:
//
//   - The number of IP entries that can be added to a network ACL with each Alibaba Cloud account at a time: 50
//
//   - The maximum number of IP entries that each network ACL can contain: 300
//
// @param request - AddAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAccessControlListEntryResponse
func (client *Client) AddAccessControlListEntryWithOptions(request *AddAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *AddAccessControlListEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAccessControlListEntry"),
		Version:     dara.String("2014-05-15"),
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
// Adds IP entries to an access control list (ACL).
//
// Description:
//
// Each network ACL can contain one or more IP addresses or CIDR blocks. Take note of the following limits on network ACLs:
//
//   - The number of IP entries that can be added to a network ACL with each Alibaba Cloud account at a time: 50
//
//   - The maximum number of IP entries that each network ACL can contain: 300
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
// Adds backend servers.
//
// Description:
//
// >  If multiple identical Elastic Compute Service (ECS) instances are specified in a request, only the first ECS instance is added. The other ECS instances are ignored. If the backend server that you add is the same as one of the existing backend servers that are already associated with the listener, an error message is returned.
//
// @param request - AddBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBackendServersResponse
func (client *Client) AddBackendServersWithOptions(request *AddBackendServersRequest, runtime *dara.RuntimeOptions) (_result *AddBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds backend servers.
//
// Description:
//
// >  If multiple identical Elastic Compute Service (ECS) instances are specified in a request, only the first ECS instance is added. The other ECS instances are ignored. If the backend server that you add is the same as one of the existing backend servers that are already associated with the listener, an error message is returned.
//
// @param request - AddBackendServersRequest
//
// @return AddBackendServersResponse
func (client *Client) AddBackendServers(request *AddBackendServersRequest) (_result *AddBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBackendServersResponse{}
	_body, _err := client.AddBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds IP addresses to the whitelist of a listener.
//
// @param request - AddListenerWhiteListItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddListenerWhiteListItemResponse
func (client *Client) AddListenerWhiteListItemWithOptions(request *AddListenerWhiteListItemRequest, runtime *dara.RuntimeOptions) (_result *AddListenerWhiteListItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceItems) {
		query["SourceItems"] = request.SourceItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddListenerWhiteListItem"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddListenerWhiteListItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds IP addresses to the whitelist of a listener.
//
// @param request - AddListenerWhiteListItemRequest
//
// @return AddListenerWhiteListItemResponse
func (client *Client) AddListenerWhiteListItem(request *AddListenerWhiteListItemRequest) (_result *AddListenerWhiteListItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddListenerWhiteListItemResponse{}
	_body, _err := client.AddListenerWhiteListItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to an SLB instance.
//
// Description:
//
// # Limits
//
// Before you call this API, note the following limits:
//
//   - You can add up to 10 tags to each SLB instance.
//
//   - You can add up to five pairs of tags at a time.
//
//   - All the tags and keys added to an SLB instance must be unique.
//
//   - If you add a tag of which the key is the same as that of an existing tag, but the value is different, the new tag overwrites the existing one.
//
// @param request - AddTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTagsResponse
func (client *Client) AddTagsWithOptions(request *AddTagsRequest, runtime *dara.RuntimeOptions) (_result *AddTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTags"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to an SLB instance.
//
// Description:
//
// # Limits
//
// Before you call this API, note the following limits:
//
//   - You can add up to 10 tags to each SLB instance.
//
//   - You can add up to five pairs of tags at a time.
//
//   - All the tags and keys added to an SLB instance must be unique.
//
//   - If you add a tag of which the key is the same as that of an existing tag, but the value is different, the new tag overwrites the existing one.
//
// @param request - AddTagsRequest
//
// @return AddTagsResponse
func (client *Client) AddTags(request *AddTagsRequest) (_result *AddTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTagsResponse{}
	_body, _err := client.AddTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds backend servers to a vServer group.
//
// @param request - AddVServerGroupBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVServerGroupBackendServersResponse
func (client *Client) AddVServerGroupBackendServersWithOptions(request *AddVServerGroupBackendServersRequest, runtime *dara.RuntimeOptions) (_result *AddVServerGroupBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVServerGroupBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds backend servers to a vServer group.
//
// @param request - AddVServerGroupBackendServersRequest
//
// @return AddVServerGroupBackendServersResponse
func (client *Client) AddVServerGroupBackendServers(request *AddVServerGroupBackendServersRequest) (_result *AddVServerGroupBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddVServerGroupBackendServersResponse{}
	_body, _err := client.AddVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL).
//
// Description:
//
// You can create multiple ACLs. Each ACL can contain one or more IP addresses or CIDR blocks. Before you create an ACL, take note of the following limits:
//
//   - An account can have a maximum of 50 ACLs in each region.
//
//   - You can add a maximum of 50 IP addresses or CIDR blocks at a time within an account.
//
//   - Each ACL can contain a maximum of 300 IP addresses or CIDR blocks.
//
// @param request - CreateAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessControlListResponse
func (client *Client) CreateAccessControlListWithOptions(request *CreateAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessControlListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessControlList"),
		Version:     dara.String("2014-05-15"),
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
// Creates an access control list (ACL).
//
// Description:
//
// You can create multiple ACLs. Each ACL can contain one or more IP addresses or CIDR blocks. Before you create an ACL, take note of the following limits:
//
//   - An account can have a maximum of 50 ACLs in each region.
//
//   - You can add a maximum of 50 IP addresses or CIDR blocks at a time within an account.
//
//   - Each ACL can contain a maximum of 300 IP addresses or CIDR blocks.
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
// Adds an additional domain name.
//
// @param request - CreateDomainExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainExtensionResponse
func (client *Client) CreateDomainExtensionWithOptions(request *CreateDomainExtensionRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomainExtension"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an additional domain name.
//
// @param request - CreateDomainExtensionRequest
//
// @return CreateDomainExtensionResponse
func (client *Client) CreateDomainExtension(request *CreateDomainExtensionRequest) (_result *CreateDomainExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainExtensionResponse{}
	_body, _err := client.CreateDomainExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Classic Load Balancer (CLB) instance.
//
// Description:
//
//	  Before you create a CLB instance, call the [DescribeAvailableResource](~~DescribeAvailableResource~~) operation to query the resources available for purchase in the region where you want to create the CLB instance.
//
//		- After a CLB instance is created, you are charged for using the CLB instance.
//
//		- The pay-as-you-go billing method supports the pay-by-specification and pay-by-LCU metering methods.
//
// @param request - CreateLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancerWithOptions(request *CreateLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteProtection) {
		query["DeleteProtection"] = request.DeleteProtection
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.MasterZoneId) {
		query["MasterZoneId"] = request.MasterZoneId
	}

	if !dara.IsNil(request.ModificationProtectionReason) {
		query["ModificationProtectionReason"] = request.ModificationProtectionReason
	}

	if !dara.IsNil(request.ModificationProtectionStatus) {
		query["ModificationProtectionStatus"] = request.ModificationProtectionStatus
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SlaveZoneId) {
		query["SlaveZoneId"] = request.SlaveZoneId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancer"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Classic Load Balancer (CLB) instance.
//
// Description:
//
//	  Before you create a CLB instance, call the [DescribeAvailableResource](~~DescribeAvailableResource~~) operation to query the resources available for purchase in the region where you want to create the CLB instance.
//
//		- After a CLB instance is created, you are charged for using the CLB instance.
//
//		- The pay-as-you-go billing method supports the pay-by-specification and pay-by-LCU metering methods.
//
// @param request - CreateLoadBalancerRequest
//
// @return CreateLoadBalancerResponse
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (_result *CreateLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerResponse{}
	_body, _err := client.CreateLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an HTTP listener for a Classic Load Balancer (CLB) instance.
//
// Description:
//
// A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](~~StartLoadBalancerListener~~) operation to start the listener. After the listener is started, the listener can forward traffic to backend servers.
//
// ## Prerequisites
//
// A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](~~StartLoadBalancerListener~~).
//
// @param request - CreateLoadBalancerHTTPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerHTTPListenerResponse
func (client *Client) CreateLoadBalancerHTTPListenerWithOptions(request *CreateLoadBalancerHTTPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerHTTPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ForwardPort) {
		query["ForwardPort"] = request.ForwardPort
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerForward) {
		query["ListenerForward"] = request.ListenerForward
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !dara.IsNil(request.XForwardedFor_ClientSrcPort) {
		query["XForwardedFor_ClientSrcPort"] = request.XForwardedFor_ClientSrcPort
	}

	if !dara.IsNil(request.XForwardedFor_SLBID) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !dara.IsNil(request.XForwardedFor_SLBIP) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !dara.IsNil(request.XForwardedFor_SLBPORT) {
		query["XForwardedFor_SLBPORT"] = request.XForwardedFor_SLBPORT
	}

	if !dara.IsNil(request.XForwardedFor_proto) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerHTTPListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerHTTPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTP listener for a Classic Load Balancer (CLB) instance.
//
// Description:
//
// A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](~~StartLoadBalancerListener~~) operation to start the listener. After the listener is started, the listener can forward traffic to backend servers.
//
// ## Prerequisites
//
// A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](~~StartLoadBalancerListener~~).
//
// @param request - CreateLoadBalancerHTTPListenerRequest
//
// @return CreateLoadBalancerHTTPListenerResponse
func (client *Client) CreateLoadBalancerHTTPListener(request *CreateLoadBalancerHTTPListenerRequest) (_result *CreateLoadBalancerHTTPListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerHTTPListenerResponse{}
	_body, _err := client.CreateLoadBalancerHTTPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an HTTPS listener.
//
// Description:
//
// A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/27597.html) operation to start the listener. After the listener is started, the listener can forward traffic to backend servers.
//
// ## Prerequisites
//
// A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer-2).
//
// @param request - CreateLoadBalancerHTTPSListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerHTTPSListenerResponse
func (client *Client) CreateLoadBalancerHTTPSListenerWithOptions(request *CreateLoadBalancerHTTPSListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerHTTPSListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.CACertificateId) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableHttp2) {
		query["EnableHttp2"] = request.EnableHttp2
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.TLSCipherPolicy) {
		query["TLSCipherPolicy"] = request.TLSCipherPolicy
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !dara.IsNil(request.XForwardedFor_ClientSrcPort) {
		query["XForwardedFor_ClientSrcPort"] = request.XForwardedFor_ClientSrcPort
	}

	if !dara.IsNil(request.XForwardedFor_SLBID) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !dara.IsNil(request.XForwardedFor_SLBIP) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !dara.IsNil(request.XForwardedFor_SLBPORT) {
		query["XForwardedFor_SLBPORT"] = request.XForwardedFor_SLBPORT
	}

	if !dara.IsNil(request.XForwardedFor_proto) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerHTTPSListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerHTTPSListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTPS listener.
//
// Description:
//
// A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/27597.html) operation to start the listener. After the listener is started, the listener can forward traffic to backend servers.
//
// ## Prerequisites
//
// A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://www.alibabacloud.com/help/en/server-load-balancer/latest/createloadbalancer-2).
//
// @param request - CreateLoadBalancerHTTPSListenerRequest
//
// @return CreateLoadBalancerHTTPSListenerResponse
func (client *Client) CreateLoadBalancerHTTPSListener(request *CreateLoadBalancerHTTPSListenerRequest) (_result *CreateLoadBalancerHTTPSListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerHTTPSListenerResponse{}
	_body, _err := client.CreateLoadBalancerHTTPSListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a TCP listener.
//
// Description:
//
// >  Newly created listeners are in the **stopped*	- state. After a listener is created, call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/2401757.html) operation to enable the listener to forward network traffic.
//
// @param request - CreateLoadBalancerTCPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerTCPListenerResponse
func (client *Client) CreateLoadBalancerTCPListenerWithOptions(request *CreateLoadBalancerTCPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerTCPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ConnectionDrain) {
		query["ConnectionDrain"] = request.ConnectionDrain
	}

	if !dara.IsNil(request.ConnectionDrainTimeout) {
		query["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckSwitch) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersistenceTimeout) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !dara.IsNil(request.ProxyProtocolV2Enabled) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["healthCheckInterval"] = request.HealthCheckInterval
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerTCPListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerTCPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a TCP listener.
//
// Description:
//
// >  Newly created listeners are in the **stopped*	- state. After a listener is created, call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/2401757.html) operation to enable the listener to forward network traffic.
//
// @param request - CreateLoadBalancerTCPListenerRequest
//
// @return CreateLoadBalancerTCPListenerResponse
func (client *Client) CreateLoadBalancerTCPListener(request *CreateLoadBalancerTCPListenerRequest) (_result *CreateLoadBalancerTCPListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerTCPListenerResponse{}
	_body, _err := client.CreateLoadBalancerTCPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a UDP listener.
//
// Description:
//
// UDP listeners of Classic Load Balancer (CLB) instances in a classic network cannot pass client IP addresses to backend servers.
//
// >  A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/27597.html) operation to enable the listener to forward traffic to backend servers.
//
// @param request - CreateLoadBalancerUDPListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadBalancerUDPListenerResponse
func (client *Client) CreateLoadBalancerUDPListenerWithOptions(request *CreateLoadBalancerUDPListenerRequest, runtime *dara.RuntimeOptions) (_result *CreateLoadBalancerUDPListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckSwitch) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProxyProtocolV2Enabled) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.HealthCheckExp) {
		query["healthCheckExp"] = request.HealthCheckExp
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["healthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckReq) {
		query["healthCheckReq"] = request.HealthCheckReq
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerUDPListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLoadBalancerUDPListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a UDP listener.
//
// Description:
//
// UDP listeners of Classic Load Balancer (CLB) instances in a classic network cannot pass client IP addresses to backend servers.
//
// >  A newly created listener is in the **stopped*	- state. After a listener is created, you can call the [StartLoadBalancerListener](https://help.aliyun.com/document_detail/27597.html) operation to enable the listener to forward traffic to backend servers.
//
// @param request - CreateLoadBalancerUDPListenerRequest
//
// @return CreateLoadBalancerUDPListenerResponse
func (client *Client) CreateLoadBalancerUDPListener(request *CreateLoadBalancerUDPListenerRequest) (_result *CreateLoadBalancerUDPListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLoadBalancerUDPListenerResponse{}
	_body, _err := client.CreateLoadBalancerUDPListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a primary/secondary server group. Each primary/secondary server group consists of two backend servers. One backend server functions as the primary server, and the other backend server functions as the secondary backend server.
//
// @param request - CreateMasterSlaveServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMasterSlaveServerGroupResponse
func (client *Client) CreateMasterSlaveServerGroupWithOptions(request *CreateMasterSlaveServerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMasterSlaveServerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveBackendServers) {
		query["MasterSlaveBackendServers"] = request.MasterSlaveBackendServers
	}

	if !dara.IsNil(request.MasterSlaveServerGroupName) {
		query["MasterSlaveServerGroupName"] = request.MasterSlaveServerGroupName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMasterSlaveServerGroup"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMasterSlaveServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a primary/secondary server group. Each primary/secondary server group consists of two backend servers. One backend server functions as the primary server, and the other backend server functions as the secondary backend server.
//
// @param request - CreateMasterSlaveServerGroupRequest
//
// @return CreateMasterSlaveServerGroupResponse
func (client *Client) CreateMasterSlaveServerGroup(request *CreateMasterSlaveServerGroupRequest) (_result *CreateMasterSlaveServerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMasterSlaveServerGroupResponse{}
	_body, _err := client.CreateMasterSlaveServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates forwarding rules for an HTTP or HTTPS listener.
//
// @param request - CreateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRulesResponse
func (client *Client) CreateRulesWithOptions(request *CreateRulesRequest, runtime *dara.RuntimeOptions) (_result *CreateRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleList) {
		query["RuleList"] = request.RuleList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRules"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates forwarding rules for an HTTP or HTTPS listener.
//
// @param request - CreateRulesRequest
//
// @return CreateRulesResponse
func (client *Client) CreateRules(request *CreateRulesRequest) (_result *CreateRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRulesResponse{}
	_body, _err := client.CreateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Transport Layer Security (TLS) policy.
//
// @param request - CreateTLSCipherPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTLSCipherPolicyResponse
func (client *Client) CreateTLSCipherPolicyWithOptions(request *CreateTLSCipherPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTLSCipherPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphers) {
		query["Ciphers"] = request.Ciphers
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TLSVersions) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTLSCipherPolicy"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTLSCipherPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Transport Layer Security (TLS) policy.
//
// @param request - CreateTLSCipherPolicyRequest
//
// @return CreateTLSCipherPolicyResponse
func (client *Client) CreateTLSCipherPolicy(request *CreateTLSCipherPolicyRequest) (_result *CreateTLSCipherPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTLSCipherPolicyResponse{}
	_body, _err := client.CreateTLSCipherPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a vServer group and adds backend servers to the vServer group.
//
// @param request - CreateVServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVServerGroupResponse
func (client *Client) CreateVServerGroupWithOptions(request *CreateVServerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateVServerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VServerGroupName) {
		query["VServerGroupName"] = request.VServerGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVServerGroup"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vServer group and adds backend servers to the vServer group.
//
// @param request - CreateVServerGroupRequest
//
// @return CreateVServerGroupResponse
func (client *Client) CreateVServerGroup(request *CreateVServerGroupRequest) (_result *CreateVServerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVServerGroupResponse{}
	_body, _err := client.CreateVServerGroupWithOptions(request, runtime)
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
// You can delete an ACL only if it is not associated with a listener.
//
// @param request - DeleteAccessControlListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessControlListResponse
func (client *Client) DeleteAccessControlListWithOptions(request *DeleteAccessControlListRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessControlListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessControlList"),
		Version:     dara.String("2014-05-15"),
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
// Deletes an access control list (ACL).
//
// Description:
//
// You can delete an ACL only if it is not associated with a listener.
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
// Deletes the access log of a Classic Load Balancer (CLB) instance.
//
// @param request - DeleteAccessLogsDownloadAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessLogsDownloadAttributeResponse
func (client *Client) DeleteAccessLogsDownloadAttributeWithOptions(request *DeleteAccessLogsDownloadAttributeRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessLogsDownloadAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LogsDownloadAttributes) {
		query["LogsDownloadAttributes"] = request.LogsDownloadAttributes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessLogsDownloadAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessLogsDownloadAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the access log of a Classic Load Balancer (CLB) instance.
//
// @param request - DeleteAccessLogsDownloadAttributeRequest
//
// @return DeleteAccessLogsDownloadAttributeResponse
func (client *Client) DeleteAccessLogsDownloadAttribute(request *DeleteAccessLogsDownloadAttributeRequest) (_result *DeleteAccessLogsDownloadAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessLogsDownloadAttributeResponse{}
	_body, _err := client.DeleteAccessLogsDownloadAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a certificate authority (CA) certificate.
//
// Description:
//
// You cannot delete a CA certificate that is in use.
//
// @param request - DeleteCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCACertificateResponse
func (client *Client) DeleteCACertificateWithOptions(request *DeleteCACertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCACertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CACertificateId) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCACertificate"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a certificate authority (CA) certificate.
//
// Description:
//
// You cannot delete a CA certificate that is in use.
//
// @param request - DeleteCACertificateRequest
//
// @return DeleteCACertificateResponse
func (client *Client) DeleteCACertificate(request *DeleteCACertificateRequest) (_result *DeleteCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCACertificateResponse{}
	_body, _err := client.DeleteCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an additional domain name.
//
// @param request - DeleteDomainExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainExtensionResponse
func (client *Client) DeleteDomainExtensionWithOptions(request *DeleteDomainExtensionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainExtensionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainExtensionId) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainExtension"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an additional domain name.
//
// @param request - DeleteDomainExtensionRequest
//
// @return DeleteDomainExtensionResponse
func (client *Client) DeleteDomainExtension(request *DeleteDomainExtensionRequest) (_result *DeleteDomainExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainExtensionResponse{}
	_body, _err := client.DeleteDomainExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a pay-as-you-go Server Load Balancer (SLB) instance.
//
// Description:
//
// > The listeners and tags of the SLB instance are deleted along with the SLB instance.
//
// @param request - DeleteLoadBalancerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancerWithOptions(request *DeleteLoadBalancerRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoadBalancerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancer"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a pay-as-you-go Server Load Balancer (SLB) instance.
//
// Description:
//
// > The listeners and tags of the SLB instance are deleted along with the SLB instance.
//
// @param request - DeleteLoadBalancerRequest
//
// @return DeleteLoadBalancerResponse
func (client *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (_result *DeleteLoadBalancerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLoadBalancerResponse{}
	_body, _err := client.DeleteLoadBalancerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a listener.
//
// Description:
//
// >  You can delete only listeners that are in the **stopped*	- or **running*	- state.
//
// @param request - DeleteLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLoadBalancerListenerResponse
func (client *Client) DeleteLoadBalancerListenerWithOptions(request *DeleteLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *DeleteLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancerListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a listener.
//
// Description:
//
// >  You can delete only listeners that are in the **stopped*	- or **running*	- state.
//
// @param request - DeleteLoadBalancerListenerRequest
//
// @return DeleteLoadBalancerListenerResponse
func (client *Client) DeleteLoadBalancerListener(request *DeleteLoadBalancerListenerRequest) (_result *DeleteLoadBalancerListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLoadBalancerListenerResponse{}
	_body, _err := client.DeleteLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a primary/secondary server group.
//
// @param request - DeleteMasterSlaveServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMasterSlaveServerGroupResponse
func (client *Client) DeleteMasterSlaveServerGroupWithOptions(request *DeleteMasterSlaveServerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMasterSlaveServerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMasterSlaveServerGroup"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMasterSlaveServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a primary/secondary server group.
//
// @param request - DeleteMasterSlaveServerGroupRequest
//
// @return DeleteMasterSlaveServerGroupResponse
func (client *Client) DeleteMasterSlaveServerGroup(request *DeleteMasterSlaveServerGroupRequest) (_result *DeleteMasterSlaveServerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMasterSlaveServerGroupResponse{}
	_body, _err := client.DeleteMasterSlaveServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes forwarding rules.
//
// Description:
//
// You must specify at least one forwarding rule that you want to delete. You can specify at most 10 forwarding rules in each call.
//
// @param request - DeleteRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRulesResponse
func (client *Client) DeleteRulesWithOptions(request *DeleteRulesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleIds) {
		query["RuleIds"] = request.RuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRules"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes forwarding rules.
//
// Description:
//
// You must specify at least one forwarding rule that you want to delete. You can specify at most 10 forwarding rules in each call.
//
// @param request - DeleteRulesRequest
//
// @return DeleteRulesResponse
func (client *Client) DeleteRules(request *DeleteRulesRequest) (_result *DeleteRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRulesResponse{}
	_body, _err := client.DeleteRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a server certificate.
//
// Description:
//
// >  You cannot delete server certificates that are in use.
//
// @param request - DeleteServerCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServerCertificateResponse
func (client *Client) DeleteServerCertificateWithOptions(request *DeleteServerCertificateRequest, runtime *dara.RuntimeOptions) (_result *DeleteServerCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServerCertificate"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a server certificate.
//
// Description:
//
// >  You cannot delete server certificates that are in use.
//
// @param request - DeleteServerCertificateRequest
//
// @return DeleteServerCertificateResponse
func (client *Client) DeleteServerCertificate(request *DeleteServerCertificateRequest) (_result *DeleteServerCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteServerCertificateResponse{}
	_body, _err := client.DeleteServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a TLS policy.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Slb\\&api=DeleteTLSCipherPolicy\\&type=RPC\\&version=2014-05-15)
//
// @param request - DeleteTLSCipherPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTLSCipherPolicyResponse
func (client *Client) DeleteTLSCipherPolicyWithOptions(request *DeleteTLSCipherPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteTLSCipherPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TLSCipherPolicyId) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTLSCipherPolicy"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTLSCipherPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a TLS policy.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=Slb\\&api=DeleteTLSCipherPolicy\\&type=RPC\\&version=2014-05-15)
//
// @param request - DeleteTLSCipherPolicyRequest
//
// @return DeleteTLSCipherPolicyResponse
func (client *Client) DeleteTLSCipherPolicy(request *DeleteTLSCipherPolicyRequest) (_result *DeleteTLSCipherPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTLSCipherPolicyResponse{}
	_body, _err := client.DeleteTLSCipherPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a VServer group.
//
// @param request - DeleteVServerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVServerGroupResponse
func (client *Client) DeleteVServerGroupWithOptions(request *DeleteVServerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteVServerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVServerGroup"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVServerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VServer group.
//
// @param request - DeleteVServerGroupRequest
//
// @return DeleteVServerGroupResponse
func (client *Client) DeleteVServerGroup(request *DeleteVServerGroupRequest) (_result *DeleteVServerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVServerGroupResponse{}
	_body, _err := client.DeleteVServerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of an access control list (ACL).
//
// @param request - DescribeAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListAttributeResponse
func (client *Client) DescribeAccessControlListAttributeWithOptions(request *DescribeAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntryComment) {
		query["AclEntryComment"] = request.AclEntryComment
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlListAttribute"),
		Version:     dara.String("2014-05-15"),
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
// Queries the configuration of an access control list (ACL).
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
// Queries access control lists (ACLs).
//
// @param request - DescribeAccessControlListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessControlListsResponse
func (client *Client) DescribeAccessControlListsWithOptions(request *DescribeAccessControlListsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessControlListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessControlLists"),
		Version:     dara.String("2014-05-15"),
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
// Queries access control lists (ACLs).
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
// Queries the access log of a Classic Load Balancer (CLB) instance.
//
// @param request - DescribeAccessLogsDownloadAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessLogsDownloadAttributeResponse
func (client *Client) DescribeAccessLogsDownloadAttributeWithOptions(request *DescribeAccessLogsDownloadAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessLogsDownloadAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LogType) {
		query["LogType"] = request.LogType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessLogsDownloadAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessLogsDownloadAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access log of a Classic Load Balancer (CLB) instance.
//
// @param request - DescribeAccessLogsDownloadAttributeRequest
//
// @return DescribeAccessLogsDownloadAttributeResponse
func (client *Client) DescribeAccessLogsDownloadAttribute(request *DescribeAccessLogsDownloadAttributeRequest) (_result *DescribeAccessLogsDownloadAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessLogsDownloadAttributeResponse{}
	_body, _err := client.DescribeAccessLogsDownloadAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available resources and resources that are available for purchase in the zones of a region.
//
// Description:
//
// > Only the available resources and zones are returned.
//
// @param request - DescribeAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResourceWithOptions(request *DescribeAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAvailableResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableResource"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available resources and resources that are available for purchase in the zones of a region.
//
// Description:
//
// > Only the available resources and zones are returned.
//
// @param request - DescribeAvailableResourceRequest
//
// @return DescribeAvailableResourceResponse
func (client *Client) DescribeAvailableResource(request *DescribeAvailableResourceRequest) (_result *DescribeAvailableResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAvailableResourceResponse{}
	_body, _err := client.DescribeAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries certificate authority (CA) certificates.
//
// Description:
//
// > To ensure data confidentiality, only the certificate fingerprint and name are returned. The certificate content is not returned.
//
// @param request - DescribeCACertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCACertificatesResponse
func (client *Client) DescribeCACertificatesWithOptions(request *DescribeCACertificatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCACertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CACertificateId) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCACertificates"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCACertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries certificate authority (CA) certificates.
//
// Description:
//
// > To ensure data confidentiality, only the certificate fingerprint and name are returned. The certificate content is not returned.
//
// @param request - DescribeCACertificatesRequest
//
// @return DescribeCACertificatesResponse
func (client *Client) DescribeCACertificates(request *DescribeCACertificatesRequest) (_result *DescribeCACertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCACertificatesResponse{}
	_body, _err := client.DescribeCACertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an additional domain name.
//
// @param request - DescribeDomainExtensionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainExtensionAttributeResponse
func (client *Client) DescribeDomainExtensionAttributeWithOptions(request *DescribeDomainExtensionAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainExtensionAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainExtensionId) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainExtensionAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainExtensionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an additional domain name.
//
// @param request - DescribeDomainExtensionAttributeRequest
//
// @return DescribeDomainExtensionAttributeResponse
func (client *Client) DescribeDomainExtensionAttribute(request *DescribeDomainExtensionAttributeRequest) (_result *DescribeDomainExtensionAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainExtensionAttributeResponse{}
	_body, _err := client.DescribeDomainExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries additional domain names.
//
// @param request - DescribeDomainExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDomainExtensionsResponse
func (client *Client) DescribeDomainExtensionsWithOptions(request *DescribeDomainExtensionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDomainExtensionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainExtensionId) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDomainExtensions"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDomainExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries additional domain names.
//
// @param request - DescribeDomainExtensionsRequest
//
// @return DescribeDomainExtensionsResponse
func (client *Client) DescribeDomainExtensions(request *DescribeDomainExtensionsRequest) (_result *DescribeDomainExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDomainExtensionsResponse{}
	_body, _err := client.DescribeDomainExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health status of backend servers.
//
// @param request - DescribeHealthStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthStatusResponse
func (client *Client) DescribeHealthStatusWithOptions(request *DescribeHealthStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthStatus"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health status of backend servers.
//
// @param request - DescribeHealthStatusRequest
//
// @return DescribeHealthStatusResponse
func (client *Client) DescribeHealthStatus(request *DescribeHealthStatusRequest) (_result *DescribeHealthStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHealthStatusResponse{}
	_body, _err := client.DescribeHealthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of fine-grained monitoring in a region.
//
// @param request - DescribeHighDefinationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHighDefinationMonitorResponse
func (client *Client) DescribeHighDefinationMonitorWithOptions(request *DescribeHighDefinationMonitorRequest, runtime *dara.RuntimeOptions) (_result *DescribeHighDefinationMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHighDefinationMonitor"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHighDefinationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of fine-grained monitoring in a region.
//
// @param request - DescribeHighDefinationMonitorRequest
//
// @return DescribeHighDefinationMonitorResponse
func (client *Client) DescribeHighDefinationMonitor(request *DescribeHighDefinationMonitorRequest) (_result *DescribeHighDefinationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeHighDefinationMonitorResponse{}
	_body, _err := client.DescribeHighDefinationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelist configurations of a listener.
//
// @param request - DescribeListenerAccessControlAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeListenerAccessControlAttributeResponse
func (client *Client) DescribeListenerAccessControlAttributeWithOptions(request *DescribeListenerAccessControlAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeListenerAccessControlAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeListenerAccessControlAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeListenerAccessControlAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelist configurations of a listener.
//
// @param request - DescribeListenerAccessControlAttributeRequest
//
// @return DescribeListenerAccessControlAttributeResponse
func (client *Client) DescribeListenerAccessControlAttribute(request *DescribeListenerAccessControlAttributeRequest) (_result *DescribeListenerAccessControlAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeListenerAccessControlAttributeResponse{}
	_body, _err := client.DescribeListenerAccessControlAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detail of a Classic Load Balancer (CLB) instance.
//
// Description:
//
// >  If backend servers are deployed in a vServer group, you can call the [DescribeVServerGroupAttribute](https://help.aliyun.com/document_detail/35224.html) operation to query the backend servers.
//
// @param request - DescribeLoadBalancerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerAttributeResponse
func (client *Client) DescribeLoadBalancerAttributeWithOptions(request *DescribeLoadBalancerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detail of a Classic Load Balancer (CLB) instance.
//
// Description:
//
// >  If backend servers are deployed in a vServer group, you can call the [DescribeVServerGroupAttribute](https://help.aliyun.com/document_detail/35224.html) operation to query the backend servers.
//
// @param request - DescribeLoadBalancerAttributeRequest
//
// @return DescribeLoadBalancerAttributeResponse
func (client *Client) DescribeLoadBalancerAttribute(request *DescribeLoadBalancerAttributeRequest) (_result *DescribeLoadBalancerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an HTTP listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTP listener is created. For more information about how to create an HTTP listener, see [CreateLoadBalancerHTTPListener](https://help.aliyun.com/document_detail/27592.html).
//
// @param request - DescribeLoadBalancerHTTPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerHTTPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPListenerAttributeWithOptions(request *DescribeLoadBalancerHTTPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerHTTPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an HTTP listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTP listener is created. For more information about how to create an HTTP listener, see [CreateLoadBalancerHTTPListener](https://help.aliyun.com/document_detail/27592.html).
//
// @param request - DescribeLoadBalancerHTTPListenerAttributeRequest
//
// @return DescribeLoadBalancerHTTPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPListenerAttribute(request *DescribeLoadBalancerHTTPListenerAttributeRequest) (_result *DescribeLoadBalancerHTTPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerHTTPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an HTTPS listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTPS listener is created. For more information about how to create an HTTPS listener, see [CreateLoadBalancerHTTPSListener](https://help.aliyun.com/document_detail/27593.html).
//
// @param request - DescribeLoadBalancerHTTPSListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPSListenerAttributeWithOptions(request *DescribeLoadBalancerHTTPSListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerHTTPSListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an HTTPS listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTPS listener is created. For more information about how to create an HTTPS listener, see [CreateLoadBalancerHTTPSListener](https://help.aliyun.com/document_detail/27593.html).
//
// @param request - DescribeLoadBalancerHTTPSListenerAttributeRequest
//
// @return DescribeLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) DescribeLoadBalancerHTTPSListenerAttribute(request *DescribeLoadBalancerHTTPSListenerAttributeRequest) (_result *DescribeLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerHTTPSListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the listeners of a Classic Load Balancer (CLB) instance.
//
// Description:
//
//	  A CLB instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/2401685.html).
//
//		- One or more listeners are added to the CLB instance. For more information, see the following topics:
//
//	    	- [CreateLoadBalancerUDPListener](~~CreateLoadBalancerUDPListener~~)
//
//	    	- [CreateLoadBalancerTCPListener](~~CreateLoadBalancerTCPListener~~)
//
//	    	- [CreateLoadBalancerHTTPListener](~~CreateLoadBalancerHTTPListener~~)
//
//	    	- [CreateLoadBalancerHTTPSListener](~~CreateLoadBalancerHTTPSListener~~)
//
// @param request - DescribeLoadBalancerListenersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerListenersResponse
func (client *Client) DescribeLoadBalancerListenersWithOptions(request *DescribeLoadBalancerListenersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerListenersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerListeners"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerListenersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the listeners of a Classic Load Balancer (CLB) instance.
//
// Description:
//
//	  A CLB instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/2401685.html).
//
//		- One or more listeners are added to the CLB instance. For more information, see the following topics:
//
//	    	- [CreateLoadBalancerUDPListener](~~CreateLoadBalancerUDPListener~~)
//
//	    	- [CreateLoadBalancerTCPListener](~~CreateLoadBalancerTCPListener~~)
//
//	    	- [CreateLoadBalancerHTTPListener](~~CreateLoadBalancerHTTPListener~~)
//
//	    	- [CreateLoadBalancerHTTPSListener](~~CreateLoadBalancerHTTPSListener~~)
//
// @param request - DescribeLoadBalancerListenersRequest
//
// @return DescribeLoadBalancerListenersResponse
func (client *Client) DescribeLoadBalancerListeners(request *DescribeLoadBalancerListenersRequest) (_result *DescribeLoadBalancerListenersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerListenersResponse{}
	_body, _err := client.DescribeLoadBalancerListenersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a TCP listener of Classic Load Balancer (CLB).
//
// @param request - DescribeLoadBalancerTCPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerTCPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerTCPListenerAttributeWithOptions(request *DescribeLoadBalancerTCPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerTCPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a TCP listener of Classic Load Balancer (CLB).
//
// @param request - DescribeLoadBalancerTCPListenerAttributeRequest
//
// @return DescribeLoadBalancerTCPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerTCPListenerAttribute(request *DescribeLoadBalancerTCPListenerAttributeRequest) (_result *DescribeLoadBalancerTCPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerTCPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a UDP listener.
//
// @param request - DescribeLoadBalancerUDPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancerUDPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerUDPListenerAttributeWithOptions(request *DescribeLoadBalancerUDPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancerUDPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerUDPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a UDP listener.
//
// @param request - DescribeLoadBalancerUDPListenerAttributeRequest
//
// @return DescribeLoadBalancerUDPListenerAttributeResponse
func (client *Client) DescribeLoadBalancerUDPListenerAttribute(request *DescribeLoadBalancerUDPListenerAttributeRequest) (_result *DescribeLoadBalancerUDPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.DescribeLoadBalancerUDPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Classic Load Balancer (CLB) instances.
//
// @param request - DescribeLoadBalancersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLoadBalancersResponse
func (client *Client) DescribeLoadBalancersWithOptions(request *DescribeLoadBalancersRequest, runtime *dara.RuntimeOptions) (_result *DescribeLoadBalancersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressIPVersion) {
		query["AddressIPVersion"] = request.AddressIPVersion
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.LoadBalancerStatus) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !dara.IsNil(request.MasterZoneId) {
		query["MasterZoneId"] = request.MasterZoneId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerId) {
		query["ServerId"] = request.ServerId
	}

	if !dara.IsNil(request.ServerIntranetAddress) {
		query["ServerIntranetAddress"] = request.ServerIntranetAddress
	}

	if !dara.IsNil(request.SlaveZoneId) {
		query["SlaveZoneId"] = request.SlaveZoneId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLoadBalancersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Classic Load Balancer (CLB) instances.
//
// @param request - DescribeLoadBalancersRequest
//
// @return DescribeLoadBalancersResponse
func (client *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (_result *DescribeLoadBalancersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeLoadBalancersResponse{}
	_body, _err := client.DescribeLoadBalancersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a primary/secondary server group.
//
// @param request - DescribeMasterSlaveServerGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMasterSlaveServerGroupAttributeResponse
func (client *Client) DescribeMasterSlaveServerGroupAttributeWithOptions(request *DescribeMasterSlaveServerGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeMasterSlaveServerGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMasterSlaveServerGroupAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMasterSlaveServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a primary/secondary server group.
//
// @param request - DescribeMasterSlaveServerGroupAttributeRequest
//
// @return DescribeMasterSlaveServerGroupAttributeResponse
func (client *Client) DescribeMasterSlaveServerGroupAttribute(request *DescribeMasterSlaveServerGroupAttributeRequest) (_result *DescribeMasterSlaveServerGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMasterSlaveServerGroupAttributeResponse{}
	_body, _err := client.DescribeMasterSlaveServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries primary/secondary server groups.
//
// @param request - DescribeMasterSlaveServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMasterSlaveServerGroupsResponse
func (client *Client) DescribeMasterSlaveServerGroupsWithOptions(request *DescribeMasterSlaveServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeMasterSlaveServerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IncludeListener) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMasterSlaveServerGroups"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMasterSlaveServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries primary/secondary server groups.
//
// @param request - DescribeMasterSlaveServerGroupsRequest
//
// @return DescribeMasterSlaveServerGroupsResponse
func (client *Client) DescribeMasterSlaveServerGroups(request *DescribeMasterSlaveServerGroupsRequest) (_result *DescribeMasterSlaveServerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMasterSlaveServerGroupsResponse{}
	_body, _err := client.DescribeMasterSlaveServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions.
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
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2014-05-15"),
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
// Queries available regions.
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
// Queries the configurations of a specified forwarding rule.
//
// @param request - DescribeRuleAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRuleAttributeResponse
func (client *Client) DescribeRuleAttributeWithOptions(request *DescribeRuleAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeRuleAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRuleAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRuleAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a specified forwarding rule.
//
// @param request - DescribeRuleAttributeRequest
//
// @return DescribeRuleAttributeResponse
func (client *Client) DescribeRuleAttribute(request *DescribeRuleAttributeRequest) (_result *DescribeRuleAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRuleAttributeResponse{}
	_body, _err := client.DescribeRuleAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the forwarding rules that are configured for a specified listener.
//
// @param request - DescribeRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRules"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the forwarding rules that are configured for a specified listener.
//
// @param request - DescribeRulesRequest
//
// @return DescribeRulesResponse
func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the server certificates in a region.
//
// Description:
//
// >  For security reasons, only the fingerprints and names of server certificates are returned. The content of server certificates and private keys is not returned.
//
// @param request - DescribeServerCertificatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServerCertificatesResponse
func (client *Client) DescribeServerCertificatesWithOptions(request *DescribeServerCertificatesRequest, runtime *dara.RuntimeOptions) (_result *DescribeServerCertificatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServerCertificates"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServerCertificatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the server certificates in a region.
//
// Description:
//
// >  For security reasons, only the fingerprints and names of server certificates are returned. The content of server certificates and private keys is not returned.
//
// @param request - DescribeServerCertificatesRequest
//
// @return DescribeServerCertificatesResponse
func (client *Client) DescribeServerCertificates(request *DescribeServerCertificatesRequest) (_result *DescribeServerCertificatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeServerCertificatesResponse{}
	_body, _err := client.DescribeServerCertificatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// Description:
//
// Take note of the following items:
//
//   - The system queries tags based on the instance ID, tag key, and tag value specified by you.
//
//   - If the logical relationship among the specified conditions is AND, only tags that match all the specified conditions are returned.
//
//   - If the Tagkey parameter is specified and but Tagvalue parameter is not specified, all tags that contain the specified tag key are returned.
//
//   - If you specify the Tagvalue parameter in a request, you must also specify the Tagkey parameter in the request.
//
//   - If you specify both the Tagkey and Tagvalue parameters, only tags that contain the specified keys and values are returned.
//
// @param request - DescribeTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTagsResponse
func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DistinctKey) {
		query["DistinctKey"] = request.DistinctKey
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTags"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// Description:
//
// Take note of the following items:
//
//   - The system queries tags based on the instance ID, tag key, and tag value specified by you.
//
//   - If the logical relationship among the specified conditions is AND, only tags that match all the specified conditions are returned.
//
//   - If the Tagkey parameter is specified and but Tagvalue parameter is not specified, all tags that contain the specified tag key are returned.
//
//   - If you specify the Tagvalue parameter in a request, you must also specify the Tagkey parameter in the request.
//
//   - If you specify both the Tagkey and Tagvalue parameters, only tags that contain the specified keys and values are returned.
//
// @param request - DescribeTagsRequest
//
// @return DescribeTagsResponse
func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a vServer group.
//
// @param request - DescribeVServerGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVServerGroupAttributeResponse
func (client *Client) DescribeVServerGroupAttributeWithOptions(request *DescribeVServerGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeVServerGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVServerGroupAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a vServer group.
//
// @param request - DescribeVServerGroupAttributeRequest
//
// @return DescribeVServerGroupAttributeResponse
func (client *Client) DescribeVServerGroupAttribute(request *DescribeVServerGroupAttributeRequest) (_result *DescribeVServerGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVServerGroupAttributeResponse{}
	_body, _err := client.DescribeVServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries vServer groups.
//
// @param request - DescribeVServerGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVServerGroupsResponse
func (client *Client) DescribeVServerGroupsWithOptions(request *DescribeVServerGroupsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVServerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.IncludeListener) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !dara.IsNil(request.IncludeRule) {
		query["IncludeRule"] = request.IncludeRule
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVServerGroups"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVServerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries vServer groups.
//
// @param request - DescribeVServerGroupsRequest
//
// @return DescribeVServerGroupsResponse
func (client *Client) DescribeVServerGroups(request *DescribeVServerGroupsRequest) (_result *DescribeVServerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVServerGroupsResponse{}
	_body, _err := client.DescribeVServerGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones in a region.
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2014-05-15"),
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
// Queries the zones in a region.
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
// Enables fine-grained monitoring for the current region.
//
// @param request - EnableHighDefinationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHighDefinationMonitorResponse
func (client *Client) EnableHighDefinationMonitorWithOptions(request *EnableHighDefinationMonitorRequest, runtime *dara.RuntimeOptions) (_result *EnableHighDefinationMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogProject) {
		query["LogProject"] = request.LogProject
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableHighDefinationMonitor"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableHighDefinationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables fine-grained monitoring for the current region.
//
// @param request - EnableHighDefinationMonitorRequest
//
// @return EnableHighDefinationMonitorResponse
func (client *Client) EnableHighDefinationMonitor(request *EnableHighDefinationMonitorRequest) (_result *EnableHighDefinationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableHighDefinationMonitorResponse{}
	_body, _err := client.EnableHighDefinationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Transport Layer Security (TLS) policies.
//
// @param request - ListTLSCipherPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTLSCipherPoliciesResponse
func (client *Client) ListTLSCipherPoliciesWithOptions(request *ListTLSCipherPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListTLSCipherPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeListener) {
		query["IncludeListener"] = request.IncludeListener
	}

	if !dara.IsNil(request.MaxItems) {
		query["MaxItems"] = request.MaxItems
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TLSCipherPolicyId) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTLSCipherPolicies"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTLSCipherPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Transport Layer Security (TLS) policies.
//
// @param request - ListTLSCipherPoliciesRequest
//
// @return ListTLSCipherPoliciesResponse
func (client *Client) ListTLSCipherPolicies(request *ListTLSCipherPoliciesRequest) (_result *ListTLSCipherPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTLSCipherPoliciesResponse{}
	_body, _err := client.ListTLSCipherPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources.
//
// Description:
//
//	  Set **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object to be queried.
//
//		- **Tag.N*	- is a resource tag that consists of a key-value pair. If you set only **Tag.N.Key**, all tag values that are associated with the specified tag key are returned. If you set only **Tag.N.Value**, an error message is returned.
//
//		- If you set **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//		- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2014-05-15"),
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
// Queries the tags that are added to resources.
//
// Description:
//
//	  Set **ResourceId.N*	- or **Tag.N*	- that consists of **Tag.N.Key*	- and **Tag.N.Value*	- in the request to specify the object to be queried.
//
//		- **Tag.N*	- is a resource tag that consists of a key-value pair. If you set only **Tag.N.Key**, all tag values that are associated with the specified tag key are returned. If you set only **Tag.N.Value**, an error message is returned.
//
//		- If you set **Tag.N*	- and **ResourceId.N*	- to filter tags, **ResourceId.N*	- must match all specified key-value pairs.
//
//		- If you specify multiple key-value pairs, resources that contain these key-value pairs are returned.
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
// Modifies the configuration of fine-grained monitoring in a specified region.
//
// @param request - ModifyHighDefinationMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHighDefinationMonitorResponse
func (client *Client) ModifyHighDefinationMonitorWithOptions(request *ModifyHighDefinationMonitorRequest, runtime *dara.RuntimeOptions) (_result *ModifyHighDefinationMonitorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogProject) {
		query["LogProject"] = request.LogProject
	}

	if !dara.IsNil(request.LogStore) {
		query["LogStore"] = request.LogStore
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHighDefinationMonitor"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHighDefinationMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of fine-grained monitoring in a specified region.
//
// @param request - ModifyHighDefinationMonitorRequest
//
// @return ModifyHighDefinationMonitorResponse
func (client *Client) ModifyHighDefinationMonitor(request *ModifyHighDefinationMonitorRequest) (_result *ModifyHighDefinationMonitorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyHighDefinationMonitorResponse{}
	_body, _err := client.ModifyHighDefinationMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the metering method of a pay-as-you-go Classic Load Balancer (CLB) instance.
//
// Description:
//
// > 	- For pay-as-you-go CLB instances, you can only change the metering method from pay-by-specification to pay-by-LCU. You cannot change the metering method from pay-by-LCU to pay-by-specification.
//
// >	- This operation can change the metering method of only one instance at a time.
//
// @param request - ModifyLoadBalancerInstanceChargeTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLoadBalancerInstanceChargeTypeResponse
func (client *Client) ModifyLoadBalancerInstanceChargeTypeWithOptions(request *ModifyLoadBalancerInstanceChargeTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyLoadBalancerInstanceChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLoadBalancerInstanceChargeType"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLoadBalancerInstanceChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the metering method of a pay-as-you-go Classic Load Balancer (CLB) instance.
//
// Description:
//
// > 	- For pay-as-you-go CLB instances, you can only change the metering method from pay-by-specification to pay-by-LCU. You cannot change the metering method from pay-by-LCU to pay-by-specification.
//
// >	- This operation can change the metering method of only one instance at a time.
//
// @param request - ModifyLoadBalancerInstanceChargeTypeRequest
//
// @return ModifyLoadBalancerInstanceChargeTypeResponse
func (client *Client) ModifyLoadBalancerInstanceChargeType(request *ModifyLoadBalancerInstanceChargeTypeRequest) (_result *ModifyLoadBalancerInstanceChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLoadBalancerInstanceChargeTypeResponse{}
	_body, _err := client.ModifyLoadBalancerInstanceChargeTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the specification of a Classic Load Balancer (CLB) instance.
//
// @param request - ModifyLoadBalancerInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLoadBalancerInstanceSpecResponse
func (client *Client) ModifyLoadBalancerInstanceSpecWithOptions(request *ModifyLoadBalancerInstanceSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyLoadBalancerInstanceSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerSpec) {
		query["LoadBalancerSpec"] = request.LoadBalancerSpec
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLoadBalancerInstanceSpec"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLoadBalancerInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the specification of a Classic Load Balancer (CLB) instance.
//
// @param request - ModifyLoadBalancerInstanceSpecRequest
//
// @return ModifyLoadBalancerInstanceSpecResponse
func (client *Client) ModifyLoadBalancerInstanceSpec(request *ModifyLoadBalancerInstanceSpecRequest) (_result *ModifyLoadBalancerInstanceSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLoadBalancerInstanceSpecResponse{}
	_body, _err := client.ModifyLoadBalancerInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the metering method of an Internet-facing Classic Load Balancer (CLB) instance.
//
// Description:
//
// ## Description
//
//   - If you modify only the maximum bandwidth of a pay-by-bandwidth CLB instance, the new bandwidth immediately takes effect.
//
//   - If you modify the metering method (for example, switch from pay-by-bandwidth to pay-by-data-transfer), the new metering method and the other changes specified in the operation take effect at 00:00:00 the next day.
//
// @param request - ModifyLoadBalancerInternetSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLoadBalancerInternetSpecResponse
func (client *Client) ModifyLoadBalancerInternetSpecWithOptions(request *ModifyLoadBalancerInternetSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyLoadBalancerInternetSpecResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.InternetChargeType) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLoadBalancerInternetSpec"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLoadBalancerInternetSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the metering method of an Internet-facing Classic Load Balancer (CLB) instance.
//
// Description:
//
// ## Description
//
//   - If you modify only the maximum bandwidth of a pay-by-bandwidth CLB instance, the new bandwidth immediately takes effect.
//
//   - If you modify the metering method (for example, switch from pay-by-bandwidth to pay-by-data-transfer), the new metering method and the other changes specified in the operation take effect at 00:00:00 the next day.
//
// @param request - ModifyLoadBalancerInternetSpecRequest
//
// @return ModifyLoadBalancerInternetSpecResponse
func (client *Client) ModifyLoadBalancerInternetSpec(request *ModifyLoadBalancerInternetSpecRequest) (_result *ModifyLoadBalancerInternetSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLoadBalancerInternetSpecResponse{}
	_body, _err := client.ModifyLoadBalancerInternetSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyLoadBalancerPayType is deprecated
//
// Summary:
//
// Changes the billing method of a Classic Load Balancer (CLB) instance from pay-as-you-go to subscription.
//
// @param request - ModifyLoadBalancerPayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLoadBalancerPayTypeResponse
func (client *Client) ModifyLoadBalancerPayTypeWithOptions(request *ModifyLoadBalancerPayTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyLoadBalancerPayTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLoadBalancerPayType"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLoadBalancerPayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyLoadBalancerPayType is deprecated
//
// Summary:
//
// Changes the billing method of a Classic Load Balancer (CLB) instance from pay-as-you-go to subscription.
//
// @param request - ModifyLoadBalancerPayTypeRequest
//
// @return ModifyLoadBalancerPayTypeResponse
// Deprecated
func (client *Client) ModifyLoadBalancerPayType(request *ModifyLoadBalancerPayTypeRequest) (_result *ModifyLoadBalancerPayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLoadBalancerPayTypeResponse{}
	_body, _err := client.ModifyLoadBalancerPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces backend servers in a vServer group.
//
// Description:
//
// This operation can be used only to replace backend servers in a vServer group. To modify the attributes of backend servers, such as weights, call the [SetVServerGroupAttribute](https://help.aliyun.com/document_detail/35217.html) operation.
//
// @param request - ModifyVServerGroupBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVServerGroupBackendServersResponse
func (client *Client) ModifyVServerGroupBackendServersWithOptions(request *ModifyVServerGroupBackendServersRequest, runtime *dara.RuntimeOptions) (_result *ModifyVServerGroupBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewBackendServers) {
		query["NewBackendServers"] = request.NewBackendServers
	}

	if !dara.IsNil(request.OldBackendServers) {
		query["OldBackendServers"] = request.OldBackendServers
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVServerGroupBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces backend servers in a vServer group.
//
// Description:
//
// This operation can be used only to replace backend servers in a vServer group. To modify the attributes of backend servers, such as weights, call the [SetVServerGroupAttribute](https://help.aliyun.com/document_detail/35217.html) operation.
//
// @param request - ModifyVServerGroupBackendServersRequest
//
// @return ModifyVServerGroupBackendServersResponse
func (client *Client) ModifyVServerGroupBackendServers(request *ModifyVServerGroupBackendServersRequest) (_result *ModifyVServerGroupBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVServerGroupBackendServersResponse{}
	_body, _err := client.ModifyVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a resource to another resource group.
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
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2014-05-15"),
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
// Moves a resource to another resource group.
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
// Removes IP entries from the network access control list (ACL) of a Classic Load Balancer (CLB) instance.
//
// @param request - RemoveAccessControlListEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAccessControlListEntryResponse
func (client *Client) RemoveAccessControlListEntryWithOptions(request *RemoveAccessControlListEntryRequest, runtime *dara.RuntimeOptions) (_result *RemoveAccessControlListEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclEntrys) {
		query["AclEntrys"] = request.AclEntrys
	}

	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAccessControlListEntry"),
		Version:     dara.String("2014-05-15"),
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
// Removes IP entries from the network access control list (ACL) of a Classic Load Balancer (CLB) instance.
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
// Removes backend servers.
//
// Description:
//
// >  If the backend servers that you want to remove are not in the server list of the Classic Load Balancer (CLB) instance, the request fails. However, the system does not report an error.
//
// @param request - RemoveBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveBackendServersResponse
func (client *Client) RemoveBackendServersWithOptions(request *RemoveBackendServersRequest, runtime *dara.RuntimeOptions) (_result *RemoveBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes backend servers.
//
// Description:
//
// >  If the backend servers that you want to remove are not in the server list of the Classic Load Balancer (CLB) instance, the request fails. However, the system does not report an error.
//
// @param request - RemoveBackendServersRequest
//
// @return RemoveBackendServersResponse
func (client *Client) RemoveBackendServers(request *RemoveBackendServersRequest) (_result *RemoveBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveBackendServersResponse{}
	_body, _err := client.RemoveBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes IP addresses or CIDR blocks from the whitelist of a listener.
//
// @param request - RemoveListenerWhiteListItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveListenerWhiteListItemResponse
func (client *Client) RemoveListenerWhiteListItemWithOptions(request *RemoveListenerWhiteListItemRequest, runtime *dara.RuntimeOptions) (_result *RemoveListenerWhiteListItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SourceItems) {
		query["SourceItems"] = request.SourceItems
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveListenerWhiteListItem"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveListenerWhiteListItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes IP addresses or CIDR blocks from the whitelist of a listener.
//
// @param request - RemoveListenerWhiteListItemRequest
//
// @return RemoveListenerWhiteListItemResponse
func (client *Client) RemoveListenerWhiteListItem(request *RemoveListenerWhiteListItemRequest) (_result *RemoveListenerWhiteListItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveListenerWhiteListItemResponse{}
	_body, _err := client.RemoveListenerWhiteListItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from a Server Load Balancer (SLB) instance.
//
// @param request - RemoveTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTagsResponse
func (client *Client) RemoveTagsWithOptions(request *RemoveTagsRequest, runtime *dara.RuntimeOptions) (_result *RemoveTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTags"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from a Server Load Balancer (SLB) instance.
//
// @param request - RemoveTagsRequest
//
// @return RemoveTagsResponse
func (client *Client) RemoveTags(request *RemoveTagsRequest) (_result *RemoveTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTagsResponse{}
	_body, _err := client.RemoveTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes backend servers from a vServer group.
//
// Description:
//
// >  If the backend servers specified by the **BackendServers*	- parameter do not exist in the vServer group, the backend servers are ignored. No error message is returned.
//
// @param request - RemoveVServerGroupBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVServerGroupBackendServersResponse
func (client *Client) RemoveVServerGroupBackendServersWithOptions(request *RemoveVServerGroupBackendServersRequest, runtime *dara.RuntimeOptions) (_result *RemoveVServerGroupBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVServerGroupBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVServerGroupBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes backend servers from a vServer group.
//
// Description:
//
// >  If the backend servers specified by the **BackendServers*	- parameter do not exist in the vServer group, the backend servers are ignored. No error message is returned.
//
// @param request - RemoveVServerGroupBackendServersRequest
//
// @return RemoveVServerGroupBackendServersResponse
func (client *Client) RemoveVServerGroupBackendServers(request *RemoveVServerGroupBackendServersRequest) (_result *RemoveVServerGroupBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveVServerGroupBackendServersResponse{}
	_body, _err := client.RemoveVServerGroupBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames an access control list (ACL).
//
// @param request - SetAccessControlListAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccessControlListAttributeResponse
func (client *Client) SetAccessControlListAttributeWithOptions(request *SetAccessControlListAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetAccessControlListAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclName) {
		query["AclName"] = request.AclName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAccessControlListAttribute"),
		Version:     dara.String("2014-05-15"),
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
// Renames an access control list (ACL).
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
// Adds an access log forwarding rule to a Classic Load Balancer (CLB) instance.
//
// @param request - SetAccessLogsDownloadAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAccessLogsDownloadAttributeResponse
func (client *Client) SetAccessLogsDownloadAttributeWithOptions(request *SetAccessLogsDownloadAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetAccessLogsDownloadAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LogsDownloadAttributes) {
		query["LogsDownloadAttributes"] = request.LogsDownloadAttributes
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAccessLogsDownloadAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAccessLogsDownloadAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an access log forwarding rule to a Classic Load Balancer (CLB) instance.
//
// @param request - SetAccessLogsDownloadAttributeRequest
//
// @return SetAccessLogsDownloadAttributeResponse
func (client *Client) SetAccessLogsDownloadAttribute(request *SetAccessLogsDownloadAttributeRequest) (_result *SetAccessLogsDownloadAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAccessLogsDownloadAttributeResponse{}
	_body, _err := client.SetAccessLogsDownloadAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies weights for backend servers.
//
// @param request - SetBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetBackendServersResponse
func (client *Client) SetBackendServersWithOptions(request *SetBackendServersRequest, runtime *dara.RuntimeOptions) (_result *SetBackendServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetBackendServers"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies weights for backend servers.
//
// @param request - SetBackendServersRequest
//
// @return SetBackendServersResponse
func (client *Client) SetBackendServers(request *SetBackendServersRequest) (_result *SetBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetBackendServersResponse{}
	_body, _err := client.SetBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Specifies a name for a CA certificate.
//
// @param request - SetCACertificateNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCACertificateNameResponse
func (client *Client) SetCACertificateNameWithOptions(request *SetCACertificateNameRequest, runtime *dara.RuntimeOptions) (_result *SetCACertificateNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CACertificateId) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !dara.IsNil(request.CACertificateName) {
		query["CACertificateName"] = request.CACertificateName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetCACertificateName"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetCACertificateNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Specifies a name for a CA certificate.
//
// @param request - SetCACertificateNameRequest
//
// @return SetCACertificateNameResponse
func (client *Client) SetCACertificateName(request *SetCACertificateNameRequest) (_result *SetCACertificateNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCACertificateNameResponse{}
	_body, _err := client.SetCACertificateNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the certificate of an additional domain name.
//
// Description:
//
// >  You cannot replace an additional certificate for a listener that is added to a shared-resource Server Load Balancer (SLB) instance.
//
// @param request - SetDomainExtensionAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDomainExtensionAttributeResponse
func (client *Client) SetDomainExtensionAttributeWithOptions(request *SetDomainExtensionAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetDomainExtensionAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainExtensionId) {
		query["DomainExtensionId"] = request.DomainExtensionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDomainExtensionAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDomainExtensionAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the certificate of an additional domain name.
//
// Description:
//
// >  You cannot replace an additional certificate for a listener that is added to a shared-resource Server Load Balancer (SLB) instance.
//
// @param request - SetDomainExtensionAttributeRequest
//
// @return SetDomainExtensionAttributeResponse
func (client *Client) SetDomainExtensionAttribute(request *SetDomainExtensionAttributeRequest) (_result *SetDomainExtensionAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDomainExtensionAttributeResponse{}
	_body, _err := client.SetDomainExtensionAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the whitelist of a specified listener.
//
// @param request - SetListenerAccessControlStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetListenerAccessControlStatusResponse
func (client *Client) SetListenerAccessControlStatusWithOptions(request *SetListenerAccessControlStatusRequest, runtime *dara.RuntimeOptions) (_result *SetListenerAccessControlStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessControlStatus) {
		query["AccessControlStatus"] = request.AccessControlStatus
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetListenerAccessControlStatus"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetListenerAccessControlStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the whitelist of a specified listener.
//
// @param request - SetListenerAccessControlStatusRequest
//
// @return SetListenerAccessControlStatusResponse
func (client *Client) SetListenerAccessControlStatus(request *SetListenerAccessControlStatusRequest) (_result *SetListenerAccessControlStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetListenerAccessControlStatusResponse{}
	_body, _err := client.SetListenerAccessControlStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables deletion protection for an SLB instance.
//
// @param request - SetLoadBalancerDeleteProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerDeleteProtectionResponse
func (client *Client) SetLoadBalancerDeleteProtectionWithOptions(request *SetLoadBalancerDeleteProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerDeleteProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteProtection) {
		query["DeleteProtection"] = request.DeleteProtection
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerDeleteProtection"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerDeleteProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables deletion protection for an SLB instance.
//
// @param request - SetLoadBalancerDeleteProtectionRequest
//
// @return SetLoadBalancerDeleteProtectionResponse
func (client *Client) SetLoadBalancerDeleteProtection(request *SetLoadBalancerDeleteProtectionRequest) (_result *SetLoadBalancerDeleteProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerDeleteProtectionResponse{}
	_body, _err := client.SetLoadBalancerDeleteProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTP listener.
//
// Description:
//
// ### Prerequisites
//
//   - A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//   - An HTTP listener is created. For more information about how to create an HTTP listener, see [CreateLoadBalancerHTTPListener](https://help.aliyun.com/document_detail/27592.html).
//
// @param request - SetLoadBalancerHTTPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerHTTPListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPListenerAttributeWithOptions(request *SetLoadBalancerHTTPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerHTTPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroup) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !dara.IsNil(request.XForwardedFor_ClientSrcPort) {
		query["XForwardedFor_ClientSrcPort"] = request.XForwardedFor_ClientSrcPort
	}

	if !dara.IsNil(request.XForwardedFor_SLBID) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !dara.IsNil(request.XForwardedFor_SLBIP) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !dara.IsNil(request.XForwardedFor_SLBPORT) {
		query["XForwardedFor_SLBPORT"] = request.XForwardedFor_SLBPORT
	}

	if !dara.IsNil(request.XForwardedFor_proto) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTP listener.
//
// Description:
//
// ### Prerequisites
//
//   - A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//   - An HTTP listener is created. For more information about how to create an HTTP listener, see [CreateLoadBalancerHTTPListener](https://help.aliyun.com/document_detail/27592.html).
//
// @param request - SetLoadBalancerHTTPListenerAttributeRequest
//
// @return SetLoadBalancerHTTPListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPListenerAttribute(request *SetLoadBalancerHTTPListenerAttributeRequest) (_result *SetLoadBalancerHTTPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerHTTPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerHTTPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTPS listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTPS listener is created. For more information about how to create an HTTPS listener, see [CreateLoadBalancerHTTPSListener](https://help.aliyun.com/document_detail/27593.html).
//
// @param request - SetLoadBalancerHTTPSListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPSListenerAttributeWithOptions(request *SetLoadBalancerHTTPSListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.CACertificateId) {
		query["CACertificateId"] = request.CACertificateId
	}

	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableHttp2) {
		query["EnableHttp2"] = request.EnableHttp2
	}

	if !dara.IsNil(request.Gzip) {
		query["Gzip"] = request.Gzip
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckMethod) {
		query["HealthCheckMethod"] = request.HealthCheckMethod
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.IdleTimeout) {
		query["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RequestTimeout) {
		query["RequestTimeout"] = request.RequestTimeout
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.TLSCipherPolicy) {
		query["TLSCipherPolicy"] = request.TLSCipherPolicy
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroup) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	if !dara.IsNil(request.XForwardedFor_ClientSrcPort) {
		query["XForwardedFor_ClientSrcPort"] = request.XForwardedFor_ClientSrcPort
	}

	if !dara.IsNil(request.XForwardedFor_SLBID) {
		query["XForwardedFor_SLBID"] = request.XForwardedFor_SLBID
	}

	if !dara.IsNil(request.XForwardedFor_SLBIP) {
		query["XForwardedFor_SLBIP"] = request.XForwardedFor_SLBIP
	}

	if !dara.IsNil(request.XForwardedFor_SLBPORT) {
		query["XForwardedFor_SLBPORT"] = request.XForwardedFor_SLBPORT
	}

	if !dara.IsNil(request.XForwardedFor_proto) {
		query["XForwardedFor_proto"] = request.XForwardedFor_proto
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerHTTPSListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an HTTPS listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- An HTTPS listener is created. For more information about how to create an HTTPS listener, see [CreateLoadBalancerHTTPSListener](https://help.aliyun.com/document_detail/27593.html).
//
// @param request - SetLoadBalancerHTTPSListenerAttributeRequest
//
// @return SetLoadBalancerHTTPSListenerAttributeResponse
func (client *Client) SetLoadBalancerHTTPSListenerAttribute(request *SetLoadBalancerHTTPSListenerAttributeRequest) (_result *SetLoadBalancerHTTPSListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerHTTPSListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerHTTPSListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the configuration read-only mode for a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerModificationProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerModificationProtectionResponse
func (client *Client) SetLoadBalancerModificationProtectionWithOptions(request *SetLoadBalancerModificationProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerModificationProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.ModificationProtectionReason) {
		query["ModificationProtectionReason"] = request.ModificationProtectionReason
	}

	if !dara.IsNil(request.ModificationProtectionStatus) {
		query["ModificationProtectionStatus"] = request.ModificationProtectionStatus
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerModificationProtection"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerModificationProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of the configuration read-only mode for a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerModificationProtectionRequest
//
// @return SetLoadBalancerModificationProtectionResponse
func (client *Client) SetLoadBalancerModificationProtection(request *SetLoadBalancerModificationProtectionRequest) (_result *SetLoadBalancerModificationProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerModificationProtectionResponse{}
	_body, _err := client.SetLoadBalancerModificationProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerNameResponse
func (client *Client) SetLoadBalancerNameWithOptions(request *SetLoadBalancerNameRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerName) {
		query["LoadBalancerName"] = request.LoadBalancerName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerName"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the name of a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerNameRequest
//
// @return SetLoadBalancerNameResponse
func (client *Client) SetLoadBalancerName(request *SetLoadBalancerNameRequest) (_result *SetLoadBalancerNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerNameResponse{}
	_body, _err := client.SetLoadBalancerNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the state of a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerStatusResponse
func (client *Client) SetLoadBalancerStatusWithOptions(request *SetLoadBalancerStatusRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.LoadBalancerStatus) {
		query["LoadBalancerStatus"] = request.LoadBalancerStatus
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerStatus"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the state of a Classic Load Balancer (CLB) instance.
//
// @param request - SetLoadBalancerStatusRequest
//
// @return SetLoadBalancerStatusResponse
func (client *Client) SetLoadBalancerStatus(request *SetLoadBalancerStatusRequest) (_result *SetLoadBalancerStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerStatusResponse{}
	_body, _err := client.SetLoadBalancerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a TCP listener of Classic Load Balancer (CLB).
//
// Description:
//
//	  A CLB instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/2401685.html).
//
//		- A TCP listener is created. For more information, see [CreateLoadBalancerTCPListener](~~CreateLoadBalancerTCPListener~~).
//
// @param request - SetLoadBalancerTCPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerTCPListenerAttributeResponse
func (client *Client) SetLoadBalancerTCPListenerAttributeWithOptions(request *SetLoadBalancerTCPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerTCPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ConnectionDrain) {
		query["ConnectionDrain"] = request.ConnectionDrain
	}

	if !dara.IsNil(request.ConnectionDrainTimeout) {
		query["ConnectionDrainTimeout"] = request.ConnectionDrainTimeout
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckSwitch) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !dara.IsNil(request.HealthCheckType) {
		query["HealthCheckType"] = request.HealthCheckType
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveServerGroup) {
		query["MasterSlaveServerGroup"] = request.MasterSlaveServerGroup
	}

	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PersistenceTimeout) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !dara.IsNil(request.ProxyProtocolV2Enabled) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.SynProxy) {
		query["SynProxy"] = request.SynProxy
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroup) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a TCP listener of Classic Load Balancer (CLB).
//
// Description:
//
//	  A CLB instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/2401685.html).
//
//		- A TCP listener is created. For more information, see [CreateLoadBalancerTCPListener](~~CreateLoadBalancerTCPListener~~).
//
// @param request - SetLoadBalancerTCPListenerAttributeRequest
//
// @return SetLoadBalancerTCPListenerAttributeResponse
func (client *Client) SetLoadBalancerTCPListenerAttribute(request *SetLoadBalancerTCPListenerAttributeRequest) (_result *SetLoadBalancerTCPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerTCPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerTCPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a UDP listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- A UDP listener is created. For more information, see [CreateLoadBalancerUDPListener](https://help.aliyun.com/document_detail/27595.html).
//
// @param request - SetLoadBalancerUDPListenerAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerUDPListenerAttributeResponse
func (client *Client) SetLoadBalancerUDPListenerAttributeWithOptions(request *SetLoadBalancerUDPListenerAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerUDPListenerAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AclStatus) {
		query["AclStatus"] = request.AclStatus
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckConnectTimeout) {
		query["HealthCheckConnectTimeout"] = request.HealthCheckConnectTimeout
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckSwitch) {
		query["HealthCheckSwitch"] = request.HealthCheckSwitch
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveServerGroup) {
		query["MasterSlaveServerGroup"] = request.MasterSlaveServerGroup
	}

	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProxyProtocolV2Enabled) {
		query["ProxyProtocolV2Enabled"] = request.ProxyProtocolV2Enabled
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroup) {
		query["VServerGroup"] = request.VServerGroup
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.HealthCheckExp) {
		query["healthCheckExp"] = request.HealthCheckExp
	}

	if !dara.IsNil(request.HealthCheckReq) {
		query["healthCheckReq"] = request.HealthCheckReq
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerUDPListenerAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a UDP listener.
//
// Description:
//
//	  A Classic Load Balancer (CLB) instance is created. For more information, see [CreateLoadBalancer](https://help.aliyun.com/document_detail/27577.html).
//
//		- A UDP listener is created. For more information, see [CreateLoadBalancerUDPListener](https://help.aliyun.com/document_detail/27595.html).
//
// @param request - SetLoadBalancerUDPListenerAttributeRequest
//
// @return SetLoadBalancerUDPListenerAttributeResponse
func (client *Client) SetLoadBalancerUDPListenerAttribute(request *SetLoadBalancerUDPListenerAttributeRequest) (_result *SetLoadBalancerUDPListenerAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerUDPListenerAttributeResponse{}
	_body, _err := client.SetLoadBalancerUDPListenerAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule that is associated with a vServer group.
//
// @param request - SetRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRuleResponse
func (client *Client) SetRuleWithOptions(request *SetRuleRequest, runtime *dara.RuntimeOptions) (_result *SetRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckConnectPort) {
		query["HealthCheckConnectPort"] = request.HealthCheckConnectPort
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
	}

	if !dara.IsNil(request.HealthCheckInterval) {
		query["HealthCheckInterval"] = request.HealthCheckInterval
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthCheckURI) {
		query["HealthCheckURI"] = request.HealthCheckURI
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.ListenerSync) {
		query["ListenerSync"] = request.ListenerSync
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRule"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule that is associated with a vServer group.
//
// @param request - SetRuleRequest
//
// @return SetRuleResponse
func (client *Client) SetRule(request *SetRuleRequest) (_result *SetRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetRuleResponse{}
	_body, _err := client.SetRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a name for a server certificate.
//
// @param request - SetServerCertificateNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetServerCertificateNameResponse
func (client *Client) SetServerCertificateNameWithOptions(request *SetServerCertificateNameRequest, runtime *dara.RuntimeOptions) (_result *SetServerCertificateNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificateId) {
		query["ServerCertificateId"] = request.ServerCertificateId
	}

	if !dara.IsNil(request.ServerCertificateName) {
		query["ServerCertificateName"] = request.ServerCertificateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetServerCertificateName"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetServerCertificateNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a name for a server certificate.
//
// @param request - SetServerCertificateNameRequest
//
// @return SetServerCertificateNameResponse
func (client *Client) SetServerCertificateName(request *SetServerCertificateNameRequest) (_result *SetServerCertificateNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetServerCertificateNameResponse{}
	_body, _err := client.SetServerCertificateNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a Transport Layer Security (TLS) policy.
//
// @param request - SetTLSCipherPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTLSCipherPolicyAttributeResponse
func (client *Client) SetTLSCipherPolicyAttributeWithOptions(request *SetTLSCipherPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetTLSCipherPolicyAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Ciphers) {
		query["Ciphers"] = request.Ciphers
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TLSCipherPolicyId) {
		query["TLSCipherPolicyId"] = request.TLSCipherPolicyId
	}

	if !dara.IsNil(request.TLSVersions) {
		query["TLSVersions"] = request.TLSVersions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTLSCipherPolicyAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTLSCipherPolicyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a Transport Layer Security (TLS) policy.
//
// @param request - SetTLSCipherPolicyAttributeRequest
//
// @return SetTLSCipherPolicyAttributeResponse
func (client *Client) SetTLSCipherPolicyAttribute(request *SetTLSCipherPolicyAttributeRequest) (_result *SetTLSCipherPolicyAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetTLSCipherPolicyAttributeResponse{}
	_body, _err := client.SetTLSCipherPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a vServer group.
//
// Description:
//
// This operation can be used to modify the weights of backend servers and names of vServer groups.
//
//   - If you want to modify backend servers in a specified vServer group, call the [ModifyVServerGroupBackendServers](https://help.aliyun.com/document_detail/35220.html) operation.
//
//   - If you want to add backend servers to a specified vServer group, call the [AddVServerGroupBackendServers](https://help.aliyun.com/document_detail/35218.html) operation.
//
// @param request - SetVServerGroupAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetVServerGroupAttributeResponse
func (client *Client) SetVServerGroupAttributeWithOptions(request *SetVServerGroupAttributeRequest, runtime *dara.RuntimeOptions) (_result *SetVServerGroupAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServers) {
		query["BackendServers"] = request.BackendServers
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.VServerGroupName) {
		query["VServerGroupName"] = request.VServerGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetVServerGroupAttribute"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetVServerGroupAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a vServer group.
//
// Description:
//
// This operation can be used to modify the weights of backend servers and names of vServer groups.
//
//   - If you want to modify backend servers in a specified vServer group, call the [ModifyVServerGroupBackendServers](https://help.aliyun.com/document_detail/35220.html) operation.
//
//   - If you want to add backend servers to a specified vServer group, call the [AddVServerGroupBackendServers](https://help.aliyun.com/document_detail/35218.html) operation.
//
// @param request - SetVServerGroupAttributeRequest
//
// @return SetVServerGroupAttributeResponse
func (client *Client) SetVServerGroupAttribute(request *SetVServerGroupAttributeRequest) (_result *SetVServerGroupAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetVServerGroupAttributeResponse{}
	_body, _err := client.SetVServerGroupAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a listener.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - You can call the operation only when the listener is in the Stopped state.
//
//   - If the operation is successful, the listener switches to the Starting state.
//
//   - You cannot perform this operation when the Classic Load Balancer (CLB) instance to which the listener belongs is in the Locked state.
//
// @param request - StartLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartLoadBalancerListenerResponse
func (client *Client) StartLoadBalancerListenerWithOptions(request *StartLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *StartLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartLoadBalancerListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a listener.
//
// Description:
//
// When you call this operation, take note of the following items:
//
//   - You can call the operation only when the listener is in the Stopped state.
//
//   - If the operation is successful, the listener switches to the Starting state.
//
//   - You cannot perform this operation when the Classic Load Balancer (CLB) instance to which the listener belongs is in the Locked state.
//
// @param request - StartLoadBalancerListenerRequest
//
// @return StartLoadBalancerListenerResponse
func (client *Client) StartLoadBalancerListener(request *StartLoadBalancerListenerRequest) (_result *StartLoadBalancerListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartLoadBalancerListenerResponse{}
	_body, _err := client.StartLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a listener.
//
// Description:
//
// Before you make this API call, note the following:
//
//   - After the API call is successfully made, the listener enters the stopped state.
//
//   - If the Classic Load Balancer (CLB) instance to which the listener to be stopped belongs is in the locked state, this API call cannot be made.
//
// >  If you stop the listener, your services will be disrupted. Exercise caution when you perform this action.
//
// @param request - StopLoadBalancerListenerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopLoadBalancerListenerResponse
func (client *Client) StopLoadBalancerListenerWithOptions(request *StopLoadBalancerListenerRequest, runtime *dara.RuntimeOptions) (_result *StopLoadBalancerListenerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopLoadBalancerListener"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopLoadBalancerListenerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a listener.
//
// Description:
//
// Before you make this API call, note the following:
//
//   - After the API call is successfully made, the listener enters the stopped state.
//
//   - If the Classic Load Balancer (CLB) instance to which the listener to be stopped belongs is in the locked state, this API call cannot be made.
//
// >  If you stop the listener, your services will be disrupted. Exercise caution when you perform this action.
//
// @param request - StopLoadBalancerListenerRequest
//
// @return StopLoadBalancerListenerResponse
func (client *Client) StopLoadBalancerListener(request *StopLoadBalancerListenerRequest) (_result *StopLoadBalancerListenerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopLoadBalancerListenerResponse{}
	_body, _err := client.StopLoadBalancerListenerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags and adds the tags to resources.
//
// Description:
//
// >  You can add at most 20 tags to each instance. Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number is reached, an error message is returned.
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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2014-05-15"),
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
// Creates tags and adds the tags to resources.
//
// Description:
//
// >  You can add at most 20 tags to each instance. Before you add tags to a resource, Alibaba Cloud checks the number of existing tags of the resource. If the maximum number is reached, an error message is returned.
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
// Removes tags from a resource.
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

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
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
		Version:     dara.String("2014-05-15"),
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
// Removes tags from a resource.
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
// Uploads a CA certificate.
//
// Description:
//
// You can upload only one CA certificate at a time. After a CA certificate is uploaded, the certificate ID, name, and fingerprint are returned.
//
// @param request - UploadCACertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCACertificateResponse
func (client *Client) UploadCACertificateWithOptions(request *UploadCACertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadCACertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CACertificate) {
		query["CACertificate"] = request.CACertificate
	}

	if !dara.IsNil(request.CACertificateName) {
		query["CACertificateName"] = request.CACertificateName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadCACertificate"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadCACertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a CA certificate.
//
// Description:
//
// You can upload only one CA certificate at a time. After a CA certificate is uploaded, the certificate ID, name, and fingerprint are returned.
//
// @param request - UploadCACertificateRequest
//
// @return UploadCACertificateResponse
func (client *Client) UploadCACertificate(request *UploadCACertificateRequest) (_result *UploadCACertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadCACertificateResponse{}
	_body, _err := client.UploadCACertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a server certificate.
//
// Description:
//
//	  You can upload only one server certificate and its private key in each call.
//
//		- After a server certificate and its private key are uploaded, the fingerprints of all server certificates that belong to your Alibaba Cloud account are returned.
//
// @param request - UploadServerCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadServerCertificateResponse
func (client *Client) UploadServerCertificateWithOptions(request *UploadServerCertificateRequest, runtime *dara.RuntimeOptions) (_result *UploadServerCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliCloudCertificateId) {
		query["AliCloudCertificateId"] = request.AliCloudCertificateId
	}

	if !dara.IsNil(request.AliCloudCertificateName) {
		query["AliCloudCertificateName"] = request.AliCloudCertificateName
	}

	if !dara.IsNil(request.AliCloudCertificateRegionId) {
		query["AliCloudCertificateRegionId"] = request.AliCloudCertificateRegionId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrivateKey) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ServerCertificate) {
		query["ServerCertificate"] = request.ServerCertificate
	}

	if !dara.IsNil(request.ServerCertificateName) {
		query["ServerCertificateName"] = request.ServerCertificateName
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadServerCertificate"),
		Version:     dara.String("2014-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadServerCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a server certificate.
//
// Description:
//
//	  You can upload only one server certificate and its private key in each call.
//
//		- After a server certificate and its private key are uploaded, the fingerprints of all server certificates that belong to your Alibaba Cloud account are returned.
//
// @param request - UploadServerCertificateRequest
//
// @return UploadServerCertificateResponse
func (client *Client) UploadServerCertificate(request *UploadServerCertificateRequest) (_result *UploadServerCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadServerCertificateResponse{}
	_body, _err := client.UploadServerCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
