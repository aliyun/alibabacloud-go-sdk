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
// 添加后端服务器
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBackendServers"),
		Version:     dara.String("2013-02-21"),
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
// 添加后端服务器
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
// 创建实例
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

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.IsPublicAddress) {
		query["IsPublicAddress"] = request.IsPublicAddress
	}

	if !dara.IsNil(request.LoadBalancerMode) {
		query["LoadBalancerMode"] = request.LoadBalancerMode
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
		Action:      dara.String("CreateLoadBalancer"),
		Version:     dara.String("2013-02-21"),
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
// 创建实例
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
// 创建HTTP监听
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

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerStatus) {
		query["ListenerStatus"] = request.ListenerStatus
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

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerHTTPListener"),
		Version:     dara.String("2013-02-21"),
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
// 创建HTTP监听
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
// 创建TCP监听
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
	if !dara.IsNil(request.BackendServerPort) {
		query["BackendServerPort"] = request.BackendServerPort
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.ConnectPort) {
		query["ConnectPort"] = request.ConnectPort
	}

	if !dara.IsNil(request.ConnectTimeout) {
		query["ConnectTimeout"] = request.ConnectTimeout
	}

	if !dara.IsNil(request.EstablishedTimeout) {
		query["EstablishedTimeout"] = request.EstablishedTimeout
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckDomain) {
		query["HealthCheckDomain"] = request.HealthCheckDomain
	}

	if !dara.IsNil(request.HealthCheckHttpCode) {
		query["HealthCheckHttpCode"] = request.HealthCheckHttpCode
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

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerStatus) {
		query["ListenerStatus"] = request.ListenerStatus
	}

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.MasterSlaveServerGroupId) {
		query["MasterSlaveServerGroupId"] = request.MasterSlaveServerGroupId
	}

	if !dara.IsNil(request.MaxConnection) {
		query["MaxConnection"] = request.MaxConnection
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLoadBalancerTCPListener"),
		Version:     dara.String("2013-02-21"),
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
// 创建TCP监听
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
// 删除实例
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancer"),
		Version:     dara.String("2013-02-21"),
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
// 删除实例
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
// 删除监听
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

	if !dara.IsNil(request.LoadBalancerId) {
		query["LoadBalancerId"] = request.LoadBalancerId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLoadBalancerListener"),
		Version:     dara.String("2013-02-21"),
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
// 删除监听
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
// 查询后端服务器
//
// @param request - DescribeBackendServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackendServersResponse
func (client *Client) DescribeBackendServersWithOptions(request *DescribeBackendServersRequest, runtime *dara.RuntimeOptions) (_result *DescribeBackendServersResponse, _err error) {
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackendServers"),
		Version:     dara.String("2013-02-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackendServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询后端服务器
//
// @param request - DescribeBackendServersRequest
//
// @return DescribeBackendServersResponse
func (client *Client) DescribeBackendServers(request *DescribeBackendServersRequest) (_result *DescribeBackendServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeBackendServersResponse{}
	_body, _err := client.DescribeBackendServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例信息
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerAttribute"),
		Version:     dara.String("2013-02-21"),
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
// 查询实例信息
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
// 查询HTTP监听配置
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2013-02-21"),
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
// 查询HTTP监听配置
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
// 查询TCP监听配置
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2013-02-21"),
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
// 查询TCP监听配置
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
// slb实例批量查询
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

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLoadBalancers"),
		Version:     dara.String("2013-02-21"),
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
// slb实例批量查询
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
// 查询可用地域
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
	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2013-02-21"),
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
// 查询可用地域
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
// 删除默认服务器组的后端服务器
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveBackendServers"),
		Version:     dara.String("2013-02-21"),
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
// 删除默认服务器组的后端服务器
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
// 更新http监听
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
	if !dara.IsNil(request.Cookie) {
		query["Cookie"] = request.Cookie
	}

	if !dara.IsNil(request.CookieTimeout) {
		query["CookieTimeout"] = request.CookieTimeout
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthCheckTimeout) {
		query["HealthCheckTimeout"] = request.HealthCheckTimeout
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
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

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.StickySession) {
		query["StickySession"] = request.StickySession
	}

	if !dara.IsNil(request.StickySessionType) {
		query["StickySessionType"] = request.StickySessionType
	}

	if !dara.IsNil(request.URI) {
		query["URI"] = request.URI
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	if !dara.IsNil(request.XForwardedFor) {
		query["XForwardedFor"] = request.XForwardedFor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerHTTPListenerAttribute"),
		Version:     dara.String("2013-02-21"),
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
// 更新http监听
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
// 修改监听状态
//
// @param request - SetLoadBalancerListenerStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetLoadBalancerListenerStatusResponse
func (client *Client) SetLoadBalancerListenerStatusWithOptions(request *SetLoadBalancerListenerStatusRequest, runtime *dara.RuntimeOptions) (_result *SetLoadBalancerListenerStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerStatus) {
		query["ListenerStatus"] = request.ListenerStatus
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerListenerStatus"),
		Version:     dara.String("2013-02-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetLoadBalancerListenerStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改监听状态
//
// @param request - SetLoadBalancerListenerStatusRequest
//
// @return SetLoadBalancerListenerStatusResponse
func (client *Client) SetLoadBalancerListenerStatus(request *SetLoadBalancerListenerStatusRequest) (_result *SetLoadBalancerListenerStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetLoadBalancerListenerStatusResponse{}
	_body, _err := client.SetLoadBalancerListenerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例名称
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerName"),
		Version:     dara.String("2013-02-21"),
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
// 修改实例名称
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
// 修改实例状态
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.AccessKeyId) {
		query["access_key_id"] = request.AccessKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerStatus"),
		Version:     dara.String("2013-02-21"),
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
// 修改实例状态
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
// 更新tcp监听
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
	if !dara.IsNil(request.ConnectPort) {
		query["ConnectPort"] = request.ConnectPort
	}

	if !dara.IsNil(request.ConnectTimeout) {
		query["ConnectTimeout"] = request.ConnectTimeout
	}

	if !dara.IsNil(request.HealthCheck) {
		query["HealthCheck"] = request.HealthCheck
	}

	if !dara.IsNil(request.HealthyThreshold) {
		query["HealthyThreshold"] = request.HealthyThreshold
	}

	if !dara.IsNil(request.HostId) {
		query["HostId"] = request.HostId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
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

	if !dara.IsNil(request.PersistenceTimeout) {
		query["PersistenceTimeout"] = request.PersistenceTimeout
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.UnhealthyThreshold) {
		query["UnhealthyThreshold"] = request.UnhealthyThreshold
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetLoadBalancerTCPListenerAttribute"),
		Version:     dara.String("2013-02-21"),
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
// 更新tcp监听
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
