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
		"ap-northeast-2-pop":          dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":                  dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-3":              dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-5":              dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("edas.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("edas.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("edas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("edas.aliyuncs.com"),
		"cn-chengdu":                  dara.String("edas.aliyuncs.com"),
		"cn-edge-1":                   dara.String("edas.aliyuncs.com"),
		"cn-fujian":                   dara.String("edas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("edas.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("edas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("edas.aliyuncs.com"),
		"cn-huhehaote":                dara.String("edas.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("edas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("edas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("edas.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("edas.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("edas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("edas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("edas.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("edas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("edas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("edas.aliyuncs.com"),
		"cn-wuhan":                    dara.String("edas.aliyuncs.com"),
		"cn-yushanfang":               dara.String("edas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("edas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("edas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("edas.aliyuncs.com"),
		"eu-west-1":                   dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"me-east-1":                   dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("edas.ap-northeast-1.aliyuncs.com"),
		"us-west-1":                   dara.String("edas.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("edas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Terminates a change process and rolls back the application. This operation is applicable to applications that are deployed in Container Service for Kubernetes (ACK) clusters.
//
// @param request - AbortAndRollbackChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortAndRollbackChangeOrderResponse
func (client *Client) AbortAndRollbackChangeOrderWithOptions(request *AbortAndRollbackChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortAndRollbackChangeOrder"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/change_order_abort_and_rollback"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a change process and rolls back the application. This operation is applicable to applications that are deployed in Container Service for Kubernetes (ACK) clusters.
//
// @param request - AbortAndRollbackChangeOrderRequest
//
// @return AbortAndRollbackChangeOrderResponse
func (client *Client) AbortAndRollbackChangeOrder(request *AbortAndRollbackChangeOrderRequest) (_result *AbortAndRollbackChangeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortAndRollbackChangeOrderResponse{}
	_body, _err := client.AbortAndRollbackChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a change process.
//
// @param request - AbortChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbortChangeOrderResponse
func (client *Client) AbortChangeOrderWithOptions(request *AbortChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AbortChangeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbortChangeOrder"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/change_order_abort"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a change process.
//
// @param request - AbortChangeOrderRequest
//
// @return AbortChangeOrderResponse
func (client *Client) AbortChangeOrder(request *AbortChangeOrderRequest) (_result *AbortChangeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortChangeOrderResponse{}
	_body, _err := client.AbortChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a log directory to an application. This operation is applicable to applications that are deployed in Alibaba Cloud Elastic Compute Service (ECS) clusters and hybrid cloud ECS clusters.
//
// @param request - AddLogPathRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLogPathResponse
func (client *Client) AddLogPathWithOptions(request *AddLogPathRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddLogPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Path) {
		body["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLogPath"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/log/popListLogDirs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLogPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a log directory to an application. This operation is applicable to applications that are deployed in Alibaba Cloud Elastic Compute Service (ECS) clusters and hybrid cloud ECS clusters.
//
// @param request - AddLogPathRequest
//
// @return AddLogPathResponse
func (client *Client) AddLogPath(request *AddLogPathRequest) (_result *AddLogPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddLogPathResponse{}
	_body, _err := client.AddLogPathWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a Resource Access Management (RAM) user the permissions on a specified application.
//
// @param request - AuthorizeApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeApplicationResponse
func (client *Client) AuthorizeApplicationWithOptions(request *AuthorizeApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthorizeApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/authorize_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a Resource Access Management (RAM) user the permissions on a specified application.
//
// @param request - AuthorizeApplicationRequest
//
// @return AuthorizeApplicationResponse
func (client *Client) AuthorizeApplication(request *AuthorizeApplicationRequest) (_result *AuthorizeApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthorizeApplicationResponse{}
	_body, _err := client.AuthorizeApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a Resource Access Management (RAM) user the permissions on a resource group.
//
// @param request - AuthorizeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResourceGroupResponse
func (client *Client) AuthorizeResourceGroupWithOptions(request *AuthorizeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthorizeResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupIds) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeResourceGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/authorize_res_group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a Resource Access Management (RAM) user the permissions on a resource group.
//
// @param request - AuthorizeResourceGroupRequest
//
// @return AuthorizeResourceGroupResponse
func (client *Client) AuthorizeResourceGroup(request *AuthorizeResourceGroupRequest) (_result *AuthorizeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthorizeResourceGroupResponse{}
	_body, _err := client.AuthorizeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns one or more roles to a RAM user.
//
// @param request - AuthorizeRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeRoleResponse
func (client *Client) AuthorizeRoleWithOptions(request *AuthorizeRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AuthorizeRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleIds) {
		query["RoleIds"] = request.RoleIds
	}

	if !dara.IsNil(request.TargetUserId) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthorizeRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/authorize_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthorizeRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns one or more roles to a RAM user.
//
// @param request - AuthorizeRoleRequest
//
// @return AuthorizeRoleResponse
func (client *Client) AuthorizeRole(request *AuthorizeRoleRequest) (_result *AuthorizeRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthorizeRoleResponse{}
	_body, _err := client.AuthorizeRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application that is deployed in an Elastic Compute Service (ECS) cluster.
//
// @param request - BindEcsSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEcsSlbResponse
func (client *Client) BindEcsSlbWithOptions(request *BindEcsSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindEcsSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DeployGroupId) {
		query["DeployGroupId"] = request.DeployGroupId
	}

	if !dara.IsNil(request.ListenerHealthCheckUrl) {
		query["ListenerHealthCheckUrl"] = request.ListenerHealthCheckUrl
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.ListenerProtocol) {
		query["ListenerProtocol"] = request.ListenerProtocol
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.VForwardingUrlRule) {
		query["VForwardingUrlRule"] = request.VForwardingUrlRule
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	if !dara.IsNil(request.VServerGroupName) {
		query["VServerGroupName"] = request.VServerGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindEcsSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/slb/bind_slb"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindEcsSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application that is deployed in an Elastic Compute Service (ECS) cluster.
//
// @param request - BindEcsSlbRequest
//
// @return BindEcsSlbResponse
func (client *Client) BindEcsSlb(request *BindEcsSlbRequest) (_result *BindEcsSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindEcsSlbResponse{}
	_body, _err := client.BindEcsSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - BindK8sSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindK8sSlbResponse
func (client *Client) BindK8sSlbWithOptions(request *BindK8sSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindK8sSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServicePortInfos) {
		query["ServicePortInfos"] = request.ServicePortInfos
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.SlbProtocol) {
		query["SlbProtocol"] = request.SlbProtocol
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.TargetPort) {
		query["TargetPort"] = request.TargetPort
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindK8sSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_slb_binding"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindK8sSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - BindK8sSlbRequest
//
// @return BindK8sSlbResponse
func (client *Client) BindK8sSlb(request *BindK8sSlbRequest) (_result *BindK8sSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindK8sSlbResponse{}
	_body, _err := client.BindK8sSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application in Enterprise Distributed Application Service (EDAS).
//
// @param request - BindSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSlbResponse
func (client *Client) BindSlbWithOptions(request *BindSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BindSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ListenerPort) {
		query["ListenerPort"] = request.ListenerPort
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.SlbIp) {
		query["SlbIp"] = request.SlbIp
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VServerGroupId) {
		query["VServerGroupId"] = request.VServerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/app/bind_slb_json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a Server Load Balancer (SLB) instance to an application in Enterprise Distributed Application Service (EDAS).
//
// @param request - BindSlbRequest
//
// @return BindSlbResponse
func (client *Client) BindSlb(request *BindSlbRequest) (_result *BindSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindSlbResponse{}
	_body, _err := client.BindSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the application instance group for an Elastic Compute Service (ECS) instance in an ECS cluster.
//
// @param request - ChangeDeployGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDeployGroupResponse
func (client *Client) ChangeDeployGroupWithOptions(request *ChangeDeployGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeDeployGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	if !dara.IsNil(request.ForceStatus) {
		query["ForceStatus"] = request.ForceStatus
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeDeployGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_change_group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDeployGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the application instance group for an Elastic Compute Service (ECS) instance in an ECS cluster.
//
// @param request - ChangeDeployGroupRequest
//
// @return ChangeDeployGroupResponse
func (client *Client) ChangeDeployGroup(request *ChangeDeployGroupRequest) (_result *ChangeDeployGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeDeployGroupResponse{}
	_body, _err := client.ChangeDeployGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually confirms the release of the next batch.
//
// @param request - ContinuePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinuePipelineResponse
func (client *Client) ContinuePipelineWithOptions(request *ContinuePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ContinuePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Confirm) {
		query["Confirm"] = request.Confirm
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinuePipeline"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/pipeline_batch_confirm"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinuePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually confirms the release of the next batch.
//
// @param request - ContinuePipelineRequest
//
// @return ContinuePipelineResponse
func (client *Client) ContinuePipeline(request *ContinuePipelineRequest) (_result *ContinuePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ContinuePipelineResponse{}
	_body, _err := client.ContinuePipelineWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts a Deployment into an application.
//
// @param request - ConvertK8sResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertK8sResourceResponse
func (client *Client) ConvertK8sResourceWithOptions(request *ConvertK8sResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConvertK8sResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ResourceName) {
		query["ResourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertK8sResource"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/k8s_resource_convert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertK8sResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts a Deployment into an application.
//
// @param request - ConvertK8sResourceRequest
//
// @return ConvertK8sResourceResponse
func (client *Client) ConvertK8sResource(request *ConvertK8sResourceRequest) (_result *ConvertK8sResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConvertK8sResourceResponse{}
	_body, _err := client.ConvertK8sResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an auto scaling policy for an application.
//
// @param request - CreateApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationScalingRuleResponse
func (client *Client) CreateApplicationScalingRuleWithOptions(request *CreateApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApplicationScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingBehaviour) {
		query["ScalingBehaviour"] = request.ScalingBehaviour
	}

	if !dara.IsNil(request.ScalingRuleEnable) {
		query["ScalingRuleEnable"] = request.ScalingRuleEnable
	}

	if !dara.IsNil(request.ScalingRuleMetric) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.ScalingRuleTimer) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	if !dara.IsNil(request.ScalingRuleTrigger) {
		query["ScalingRuleTrigger"] = request.ScalingRuleTrigger
	}

	if !dara.IsNil(request.ScalingRuleType) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplicationScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/application_scaling_rule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an auto scaling policy for an application.
//
// @param request - CreateApplicationScalingRuleRequest
//
// @return CreateApplicationScalingRuleResponse
func (client *Client) CreateApplicationScalingRule(request *CreateApplicationScalingRuleRequest) (_result *CreateApplicationScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApplicationScalingRuleResponse{}
	_body, _err := client.CreateApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration template.
//
// @param request - CreateConfigTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigTemplateResponse
func (client *Client) CreateConfigTemplateWithOptions(request *CreateConfigTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConfigTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfigTemplate"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/config_template"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration template.
//
// @param request - CreateConfigTemplateRequest
//
// @return CreateConfigTemplateResponse
func (client *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (_result *CreateConfigTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigTemplateResponse{}
	_body, _err := client.CreateConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a command that is used to import instances to a hybrid cloud Elastic Compute Service (ECS) cluster.
//
// Description:
//
// ## Description
//
// You must call the CreateIDCImportCommand operation first to generate a command used to import hybrid cloud ECS instances to a hybrid cloud ECS cluster. Then, run this command on the instances to import the instances to the cluster.
//
// @param request - CreateIDCImportCommandRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIDCImportCommandResponse
func (client *Client) CreateIDCImportCommandWithOptions(request *CreateIDCImportCommandRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIDCImportCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIDCImportCommand"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/create_idc_import_command"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIDCImportCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a command that is used to import instances to a hybrid cloud Elastic Compute Service (ECS) cluster.
//
// Description:
//
// ## Description
//
// You must call the CreateIDCImportCommand operation first to generate a command used to import hybrid cloud ECS instances to a hybrid cloud ECS cluster. Then, run this command on the instances to import the instances to the cluster.
//
// @param request - CreateIDCImportCommandRequest
//
// @return CreateIDCImportCommandResponse
func (client *Client) CreateIDCImportCommand(request *CreateIDCImportCommandRequest) (_result *CreateIDCImportCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIDCImportCommandResponse{}
	_body, _err := client.CreateIDCImportCommandWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Kubernetes ConfigMap.
//
// @param request - CreateK8sConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateK8sConfigMapResponse
func (client *Client) CreateK8sConfigMapWithOptions(request *CreateK8sConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateK8sConfigMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateK8sConfigMap"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_config_map"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateK8sConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Kubernetes ConfigMap.
//
// @param request - CreateK8sConfigMapRequest
//
// @return CreateK8sConfigMapResponse
func (client *Client) CreateK8sConfigMap(request *CreateK8sConfigMapRequest) (_result *CreateK8sConfigMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateK8sConfigMapResponse{}
	_body, _err := client.CreateK8sConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Ingress.
//
// @param request - CreateK8sIngressRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateK8sIngressRuleResponse
func (client *Client) CreateK8sIngressRuleWithOptions(request *CreateK8sIngressRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateK8sIngressRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IngressConf) {
		query["IngressConf"] = request.IngressConf
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateK8sIngressRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_ingress"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateK8sIngressRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Ingress.
//
// @param request - CreateK8sIngressRuleRequest
//
// @return CreateK8sIngressRuleResponse
func (client *Client) CreateK8sIngressRule(request *CreateK8sIngressRuleRequest) (_result *CreateK8sIngressRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateK8sIngressRuleResponse{}
	_body, _err := client.CreateK8sIngressRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Kubernetes Secret.
//
// @param request - CreateK8sSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateK8sSecretResponse
func (client *Client) CreateK8sSecretWithOptions(request *CreateK8sSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateK8sSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Base64Encoded) {
		body["Base64Encoded"] = request.Base64Encoded
	}

	if !dara.IsNil(request.CertId) {
		body["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertRegionId) {
		body["CertRegionId"] = request.CertRegionId
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateK8sSecret"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_secret"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateK8sSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Kubernetes Secret.
//
// @param request - CreateK8sSecretRequest
//
// @return CreateK8sSecretResponse
func (client *Client) CreateK8sSecret(request *CreateK8sSecretRequest) (_result *CreateK8sSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateK8sSecretResponse{}
	_body, _err := client.CreateK8sSecretWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application service in a Kubernetes cluster.
//
// @param request - CreateK8sServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateK8sServiceResponse
func (client *Client) CreateK8sServiceWithOptions(request *CreateK8sServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateK8sServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ExternalTrafficPolicy) {
		query["ExternalTrafficPolicy"] = request.ExternalTrafficPolicy
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ServicePorts) {
		query["ServicePorts"] = request.ServicePorts
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateK8sService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_service"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateK8sServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application service in a Kubernetes cluster.
//
// @param request - CreateK8sServiceRequest
//
// @return CreateK8sServiceResponse
func (client *Client) CreateK8sService(request *CreateK8sServiceRequest) (_result *CreateK8sServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateK8sServiceResponse{}
	_body, _err := client.CreateK8sServiceWithOptions(request, headers, runtime)
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
// @param request - DeleteApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_delete_app"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationResponse{}
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
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an auto scaling policy for an application.
//
// @param request - DeleteApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationScalingRuleResponse
func (client *Client) DeleteApplicationScalingRuleWithOptions(request *DeleteApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApplicationScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplicationScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/application_scaling_rule"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an auto scaling policy for an application.
//
// @param request - DeleteApplicationScalingRuleRequest
//
// @return DeleteApplicationScalingRuleResponse
func (client *Client) DeleteApplicationScalingRule(request *DeleteApplicationScalingRuleRequest) (_result *DeleteApplicationScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApplicationScalingRuleResponse{}
	_body, _err := client.DeleteApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Elastic Compute Service (ECS) cluster or cancels the import of a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeleteClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Elastic Compute Service (ECS) cluster or cancels the import of a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an Elastic Compute Service (ECS) instance from a cluster.
//
// @param request - DeleteClusterMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterMemberResponse
func (client *Client) DeleteClusterMemberWithOptions(request *DeleteClusterMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteClusterMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterMemberId) {
		query["ClusterMemberId"] = request.ClusterMemberId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClusterMember"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster_member"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClusterMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an Elastic Compute Service (ECS) instance from a cluster.
//
// @param request - DeleteClusterMemberRequest
//
// @return DeleteClusterMemberResponse
func (client *Client) DeleteClusterMember(request *DeleteClusterMemberRequest) (_result *DeleteClusterMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterMemberResponse{}
	_body, _err := client.DeleteClusterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a configuration template.
//
// @param request - DeleteConfigTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigTemplateResponse
func (client *Client) DeleteConfigTemplateWithOptions(request *DeleteConfigTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigTemplateResponse, _err error) {
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
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConfigTemplate"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/config_template"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a configuration template.
//
// @param request - DeleteConfigTemplateRequest
//
// @return DeleteConfigTemplateResponse
func (client *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (_result *DeleteConfigTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigTemplateResponse{}
	_body, _err := client.DeleteConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an instance group for an application.
//
// @param request - DeleteDeployGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeployGroupResponse
func (client *Client) DeleteDeployGroupWithOptions(request *DeleteDeployGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDeployGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeployGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/deploy_group"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeployGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an instance group for an application.
//
// @param request - DeleteDeployGroupRequest
//
// @return DeleteDeployGroupResponse
func (client *Client) DeleteDeployGroup(request *DeleteDeployGroupRequest) (_result *DeleteDeployGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDeployGroupResponse{}
	_body, _err := client.DeleteDeployGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Elastic Compute Unit (ECU).
//
// @param request - DeleteEcuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEcuResponse
func (client *Client) DeleteEcuWithOptions(request *DeleteEcuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEcuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EcuId) {
		query["EcuId"] = request.EcuId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEcu"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/delete_ecu"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Elastic Compute Unit (ECU).
//
// @param request - DeleteEcuRequest
//
// @return DeleteEcuResponse
func (client *Client) DeleteEcu(request *DeleteEcuRequest) (_result *DeleteEcuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEcuResponse{}
	_body, _err := client.DeleteEcuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application from a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeleteK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteK8sApplicationResponse
func (client *Client) DeleteK8sApplicationWithOptions(request *DeleteK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Force) {
		query["Force"] = request.Force
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_apps"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application from a Container Service for Kubernetes (ACK) cluster.
//
// @param request - DeleteK8sApplicationRequest
//
// @return DeleteK8sApplicationResponse
func (client *Client) DeleteK8sApplication(request *DeleteK8sApplicationRequest) (_result *DeleteK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteK8sApplicationResponse{}
	_body, _err := client.DeleteK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Kubernetes ConfigMap.
//
// @param request - DeleteK8sConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteK8sConfigMapResponse
func (client *Client) DeleteK8sConfigMapWithOptions(request *DeleteK8sConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteK8sConfigMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteK8sConfigMap"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_config_map"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteK8sConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Kubernetes ConfigMap.
//
// @param request - DeleteK8sConfigMapRequest
//
// @return DeleteK8sConfigMapResponse
func (client *Client) DeleteK8sConfigMap(request *DeleteK8sConfigMapRequest) (_result *DeleteK8sConfigMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteK8sConfigMapResponse{}
	_body, _err := client.DeleteK8sConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ingress.
//
// @param request - DeleteK8sIngressRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteK8sIngressRuleResponse
func (client *Client) DeleteK8sIngressRuleWithOptions(request *DeleteK8sIngressRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteK8sIngressRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteK8sIngressRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_ingress"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteK8sIngressRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ingress.
//
// @param request - DeleteK8sIngressRuleRequest
//
// @return DeleteK8sIngressRuleResponse
func (client *Client) DeleteK8sIngressRule(request *DeleteK8sIngressRuleRequest) (_result *DeleteK8sIngressRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteK8sIngressRuleResponse{}
	_body, _err := client.DeleteK8sIngressRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Kubernetes Secret.
//
// @param request - DeleteK8sSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteK8sSecretResponse
func (client *Client) DeleteK8sSecretWithOptions(request *DeleteK8sSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteK8sSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteK8sSecret"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_secret"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteK8sSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Kubernetes Secret.
//
// @param request - DeleteK8sSecretRequest
//
// @return DeleteK8sSecretResponse
func (client *Client) DeleteK8sSecret(request *DeleteK8sSecretRequest) (_result *DeleteK8sSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteK8sSecretResponse{}
	_body, _err := client.DeleteK8sSecretWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application service from a Kubernetes cluster.
//
// @param request - DeleteK8sServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteK8sServiceResponse
func (client *Client) DeleteK8sServiceWithOptions(request *DeleteK8sServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteK8sServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteK8sService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_service"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteK8sServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application service from a Kubernetes cluster.
//
// @param request - DeleteK8sServiceRequest
//
// @return DeleteK8sServiceResponse
func (client *Client) DeleteK8sService(request *DeleteK8sServiceRequest) (_result *DeleteK8sServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteK8sServiceResponse{}
	_body, _err := client.DeleteK8sServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a log directory from an application. This operation is applicable to applications that are deployed in Alibaba Cloud Elastic Compute Service (ECS) clusters and hybrid cloud ECS clusters.
//
// @param request - DeleteLogPathRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogPathResponse
func (client *Client) DeleteLogPathWithOptions(request *DeleteLogPathRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLogPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogPath"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/log/popListLogDirs"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a log directory from an application. This operation is applicable to applications that are deployed in Alibaba Cloud Elastic Compute Service (ECS) clusters and hybrid cloud ECS clusters.
//
// @param request - DeleteLogPathRequest
//
// @return DeleteLogPathResponse
func (client *Client) DeleteLogPath(request *DeleteLogPathRequest) (_result *DeleteLogPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLogPathResponse{}
	_body, _err := client.DeleteLogPathWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Resource Access Management (RAM) role.
//
// @param request - DeleteRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRoleResponse
func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/delete_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
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
// Deletes a Resource Access Management (RAM) role.
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
// Deletes a service group.
//
// @param request - DeleteServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceGroupResponse
func (client *Client) DeleteServiceGroupWithOptions(request *DeleteServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/serviceGroups"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service group.
//
// @param request - DeleteServiceGroupRequest
//
// @return DeleteServiceGroupResponse
func (client *Client) DeleteServiceGroup(request *DeleteServiceGroupRequest) (_result *DeleteServiceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceGroupResponse{}
	_body, _err := client.DeleteServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a lane.
//
// @param request - DeleteSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSwimmingLaneResponse
func (client *Client) DeleteSwimmingLaneWithOptions(request *DeleteSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSwimmingLane"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lanes"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a lane.
//
// @param request - DeleteSwimmingLaneRequest
//
// @return DeleteSwimmingLaneResponse
func (client *Client) DeleteSwimmingLane(request *DeleteSwimmingLaneRequest) (_result *DeleteSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSwimmingLaneResponse{}
	_body, _err := client.DeleteSwimmingLaneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified custom namespace.
//
// @param request - DeleteUserDefineRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDefineRegionResponse
func (client *Client) DeleteUserDefineRegionWithOptions(request *DeleteUserDefineRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteUserDefineRegionResponse, _err error) {
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

	if !dara.IsNil(request.RegionTag) {
		query["RegionTag"] = request.RegionTag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserDefineRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/user_region_def"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserDefineRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified custom namespace.
//
// @param request - DeleteUserDefineRegionRequest
//
// @return DeleteUserDefineRegionResponse
func (client *Client) DeleteUserDefineRegion(request *DeleteUserDefineRegionRequest) (_result *DeleteUserDefineRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserDefineRegionResponse{}
	_body, _err := client.DeleteUserDefineRegionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys an application in an Elastic Compute Service (ECS) cluster.
//
// Description:
//
// > To deploy an application in a Container Service for Kubernetes (ACK) cluster that is imported into Enterprise Distributed Application Service (EDAS), call the DeployK8sApplication operation provided by EDAS. For more information, see [](~~149420~~)DeployK8sApplication.
//
// @param request - DeployApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationResponse
func (client *Client) DeployApplicationWithOptions(request *DeployApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppEnv) {
		query["AppEnv"] = request.AppEnv
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Batch) {
		query["Batch"] = request.Batch
	}

	if !dara.IsNil(request.BatchWaitTime) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !dara.IsNil(request.BuildPackId) {
		query["BuildPackId"] = request.BuildPackId
	}

	if !dara.IsNil(request.ComponentIds) {
		query["ComponentIds"] = request.ComponentIds
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Gray) {
		query["Gray"] = request.Gray
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.ReleaseType) {
		query["ReleaseType"] = request.ReleaseType
	}

	if !dara.IsNil(request.TrafficControlStrategy) {
		query["TrafficControlStrategy"] = request.TrafficControlStrategy
	}

	if !dara.IsNil(request.WarUrl) {
		query["WarUrl"] = request.WarUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an application in an Elastic Compute Service (ECS) cluster.
//
// Description:
//
// > To deploy an application in a Container Service for Kubernetes (ACK) cluster that is imported into Enterprise Distributed Application Service (EDAS), call the DeployK8sApplication operation provided by EDAS. For more information, see [](~~149420~~)DeployK8sApplication.
//
// @param request - DeployApplicationRequest
//
// @return DeployApplicationResponse
func (client *Client) DeployApplication(request *DeployApplicationRequest) (_result *DeployApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployApplicationResponse{}
	_body, _err := client.DeployApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys an application in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - DeployK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployK8sApplicationResponse
func (client *Client) DeployK8sApplicationWithOptions(request *DeployK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Args) {
		query["Args"] = request.Args
	}

	if !dara.IsNil(request.BatchTimeout) {
		query["BatchTimeout"] = request.BatchTimeout
	}

	if !dara.IsNil(request.BatchWaitTime) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !dara.IsNil(request.BuildPackId) {
		query["BuildPackId"] = request.BuildPackId
	}

	if !dara.IsNil(request.CanaryRuleId) {
		query["CanaryRuleId"] = request.CanaryRuleId
	}

	if !dara.IsNil(request.ChangeOrderDesc) {
		query["ChangeOrderDesc"] = request.ChangeOrderDesc
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.ConfigMountDescs) {
		query["ConfigMountDescs"] = request.ConfigMountDescs
	}

	if !dara.IsNil(request.CpuLimit) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.CpuRequest) {
		query["CpuRequest"] = request.CpuRequest
	}

	if !dara.IsNil(request.CustomAffinity) {
		query["CustomAffinity"] = request.CustomAffinity
	}

	if !dara.IsNil(request.CustomAgentVersion) {
		query["CustomAgentVersion"] = request.CustomAgentVersion
	}

	if !dara.IsNil(request.CustomTolerations) {
		query["CustomTolerations"] = request.CustomTolerations
	}

	if !dara.IsNil(request.DeployAcrossNodes) {
		query["DeployAcrossNodes"] = request.DeployAcrossNodes
	}

	if !dara.IsNil(request.DeployAcrossZones) {
		query["DeployAcrossZones"] = request.DeployAcrossZones
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.EmptyDirs) {
		query["EmptyDirs"] = request.EmptyDirs
	}

	if !dara.IsNil(request.EnableAhas) {
		query["EnableAhas"] = request.EnableAhas
	}

	if !dara.IsNil(request.EnableEmptyPushReject) {
		query["EnableEmptyPushReject"] = request.EnableEmptyPushReject
	}

	if !dara.IsNil(request.EnableLosslessRule) {
		query["EnableLosslessRule"] = request.EnableLosslessRule
	}

	if !dara.IsNil(request.EnvFroms) {
		query["EnvFroms"] = request.EnvFroms
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.Image) {
		query["Image"] = request.Image
	}

	if !dara.IsNil(request.ImagePlatforms) {
		query["ImagePlatforms"] = request.ImagePlatforms
	}

	if !dara.IsNil(request.ImageTag) {
		query["ImageTag"] = request.ImageTag
	}

	if !dara.IsNil(request.InitContainers) {
		query["InitContainers"] = request.InitContainers
	}

	if !dara.IsNil(request.JDK) {
		query["JDK"] = request.JDK
	}

	if !dara.IsNil(request.JavaStartUpConfig) {
		query["JavaStartUpConfig"] = request.JavaStartUpConfig
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.LimitEphemeralStorage) {
		query["LimitEphemeralStorage"] = request.LimitEphemeralStorage
	}

	if !dara.IsNil(request.Liveness) {
		query["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.LocalVolume) {
		query["LocalVolume"] = request.LocalVolume
	}

	if !dara.IsNil(request.LosslessRuleAligned) {
		query["LosslessRuleAligned"] = request.LosslessRuleAligned
	}

	if !dara.IsNil(request.LosslessRuleDelayTime) {
		query["LosslessRuleDelayTime"] = request.LosslessRuleDelayTime
	}

	if !dara.IsNil(request.LosslessRuleFuncType) {
		query["LosslessRuleFuncType"] = request.LosslessRuleFuncType
	}

	if !dara.IsNil(request.LosslessRuleRelated) {
		query["LosslessRuleRelated"] = request.LosslessRuleRelated
	}

	if !dara.IsNil(request.LosslessRuleWarmupTime) {
		query["LosslessRuleWarmupTime"] = request.LosslessRuleWarmupTime
	}

	if !dara.IsNil(request.McpuLimit) {
		query["McpuLimit"] = request.McpuLimit
	}

	if !dara.IsNil(request.McpuRequest) {
		query["McpuRequest"] = request.McpuRequest
	}

	if !dara.IsNil(request.MemoryLimit) {
		query["MemoryLimit"] = request.MemoryLimit
	}

	if !dara.IsNil(request.MemoryRequest) {
		query["MemoryRequest"] = request.MemoryRequest
	}

	if !dara.IsNil(request.MountDescs) {
		query["MountDescs"] = request.MountDescs
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PackageVersionId) {
		query["PackageVersionId"] = request.PackageVersionId
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.PvcMountDescs) {
		query["PvcMountDescs"] = request.PvcMountDescs
	}

	if !dara.IsNil(request.Readiness) {
		query["Readiness"] = request.Readiness
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.RequestsEphemeralStorage) {
		query["RequestsEphemeralStorage"] = request.RequestsEphemeralStorage
	}

	if !dara.IsNil(request.RuntimeClassName) {
		query["RuntimeClassName"] = request.RuntimeClassName
	}

	if !dara.IsNil(request.SecurityContext) {
		query["SecurityContext"] = request.SecurityContext
	}

	if !dara.IsNil(request.Sidecars) {
		query["Sidecars"] = request.Sidecars
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.Startup) {
		query["Startup"] = request.Startup
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TerminateGracePeriod) {
		query["TerminateGracePeriod"] = request.TerminateGracePeriod
	}

	if !dara.IsNil(request.TrafficControlStrategy) {
		query["TrafficControlStrategy"] = request.TrafficControlStrategy
	}

	if !dara.IsNil(request.UpdateStrategy) {
		query["UpdateStrategy"] = request.UpdateStrategy
	}

	if !dara.IsNil(request.UriEncoding) {
		query["UriEncoding"] = request.UriEncoding
	}

	if !dara.IsNil(request.UseBodyEncoding) {
		query["UseBodyEncoding"] = request.UseBodyEncoding
	}

	if !dara.IsNil(request.UserBaseImageUrl) {
		query["UserBaseImageUrl"] = request.UserBaseImageUrl
	}

	if !dara.IsNil(request.VolumesStr) {
		query["VolumesStr"] = request.VolumesStr
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	if !dara.IsNil(request.WebContainerConfig) {
		query["WebContainerConfig"] = request.WebContainerConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_apps"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an application in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - DeployK8sApplicationRequest
//
// @return DeployK8sApplicationResponse
func (client *Client) DeployK8sApplication(request *DeployK8sApplicationRequest) (_result *DeployK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployK8sApplicationResponse{}
	_body, _err := client.DeployK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Kubernetes application instances.
//
// @param request - DescribeAppInstanceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppInstanceListResponse
func (client *Client) DescribeAppInstanceListWithOptions(request *DescribeAppInstanceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAppInstanceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.WithNodeInfo) {
		query["WithNodeInfo"] = request.WithNodeInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppInstanceList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/app_instance_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Kubernetes application instances.
//
// @param request - DescribeAppInstanceListRequest
//
// @return DescribeAppInstanceListResponse
func (client *Client) DescribeAppInstanceList(request *DescribeAppInstanceListRequest) (_result *DescribeAppInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppInstanceListResponse{}
	_body, _err := client.DescribeAppInstanceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the auto scaling policies of an application.
//
// @param request - DescribeApplicationScalingRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationScalingRulesResponse
func (client *Client) DescribeApplicationScalingRulesWithOptions(request *DescribeApplicationScalingRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApplicationScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApplicationScalingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/application_scaling_rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApplicationScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the auto scaling policies of an application.
//
// @param request - DescribeApplicationScalingRulesRequest
//
// @return DescribeApplicationScalingRulesResponse
func (client *Client) DescribeApplicationScalingRules(request *DescribeApplicationScalingRulesRequest) (_result *DescribeApplicationScalingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApplicationScalingRulesResponse{}
	_body, _err := client.DescribeApplicationScalingRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeLocalitySettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLocalitySettingResponse
func (client *Client) DescribeLocalitySettingWithOptions(request *DescribeLocalitySettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeLocalitySettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeLocalitySetting"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/applications/locality/setting"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeLocalitySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeLocalitySettingRequest
//
// @return DescribeLocalitySettingResponse
func (client *Client) DescribeLocalitySetting(request *DescribeLocalitySettingRequest) (_result *DescribeLocalitySettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeLocalitySettingResponse{}
	_body, _err := client.DescribeLocalitySettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an auto scaling policy for an application.
//
// @param request - DisableApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableApplicationScalingRuleResponse
func (client *Client) DisableApplicationScalingRuleWithOptions(request *DisableApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableApplicationScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableApplicationScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/disable_application_scaling_rule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an auto scaling policy for an application.
//
// @param request - DisableApplicationScalingRuleRequest
//
// @return DisableApplicationScalingRuleResponse
func (client *Client) DisableApplicationScalingRule(request *DisableApplicationScalingRuleRequest) (_result *DisableApplicationScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableApplicationScalingRuleResponse{}
	_body, _err := client.DisableApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an auto scaling policy for an application.
//
// @param request - EnableApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableApplicationScalingRuleResponse
func (client *Client) EnableApplicationScalingRuleWithOptions(request *EnableApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableApplicationScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableApplicationScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/enable_application_scaling_rule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an auto scaling policy for an application.
//
// @param request - EnableApplicationScalingRuleRequest
//
// @return EnableApplicationScalingRuleResponse
func (client *Client) EnableApplicationScalingRule(request *EnableApplicationScalingRuleRequest) (_result *EnableApplicationScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableApplicationScalingRuleResponse{}
	_body, _err := client.EnableApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the Deployment of a Kubernetes application.
//
// @param request - GetAppDeploymentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppDeploymentResponse
func (client *Client) GetAppDeploymentWithOptions(request *GetAppDeploymentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppDeploymentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppDeployment"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/app_deployment"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the Deployment of a Kubernetes application.
//
// @param request - GetAppDeploymentRequest
//
// @return GetAppDeploymentResponse
func (client *Client) GetAppDeployment(request *GetAppDeploymentRequest) (_result *GetAppDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppDeploymentResponse{}
	_body, _err := client.GetAppDeploymentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a specified application in an Elastic Compute Service (ECS) cluster.
//
// @param request - GetApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/app_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a specified application in an Elastic Compute Service (ECS) cluster.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about a change process.
//
// @param request - GetChangeOrderInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChangeOrderInfoResponse
func (client *Client) GetChangeOrderInfoWithOptions(request *GetChangeOrderInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChangeOrderInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChangeOrderInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/change_order_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChangeOrderInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about a change process.
//
// @param request - GetChangeOrderInfoRequest
//
// @return GetChangeOrderInfoResponse
func (client *Client) GetChangeOrderInfo(request *GetChangeOrderInfoRequest) (_result *GetChangeOrderInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChangeOrderInfoResponse{}
	_body, _err := client.GetChangeOrderInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specific cluster.
//
// @param request - GetClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specific cluster.
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Tomcat configuration of an application or an instance group in which an application is deployed.
//
// @param request - GetContainerConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContainerConfigurationResponse
func (client *Client) GetContainerConfigurationWithOptions(request *GetContainerConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetContainerConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetContainerConfiguration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/container_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContainerConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Tomcat configuration of an application or an instance group in which an application is deployed.
//
// @param request - GetContainerConfigurationRequest
//
// @return GetContainerConfigurationResponse
func (client *Client) GetContainerConfiguration(request *GetContainerConfigurationRequest) (_result *GetContainerConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetContainerConfigurationResponse{}
	_body, _err := client.GetContainerConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of Java startup parameters for an application.
//
// @param request - GetJavaStartUpConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJavaStartUpConfigResponse
func (client *Client) GetJavaStartUpConfigWithOptions(request *GetJavaStartUpConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJavaStartUpConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJavaStartUpConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/java_start_up_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJavaStartUpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of Java startup parameters for an application.
//
// @param request - GetJavaStartUpConfigRequest
//
// @return GetJavaStartUpConfigResponse
func (client *Client) GetJavaStartUpConfig(request *GetJavaStartUpConfigRequest) (_result *GetJavaStartUpConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJavaStartUpConfigResponse{}
	_body, _err := client.GetJavaStartUpConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Java Virtual Machine (JVM) configuration of an application or an instance group in which an application is deployed.
//
// @param request - GetJvmConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJvmConfigurationResponse
func (client *Client) GetJvmConfigurationWithOptions(request *GetJvmConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJvmConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJvmConfiguration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/app_jvm_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJvmConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Java Virtual Machine (JVM) configuration of an application or an instance group in which an application is deployed.
//
// @param request - GetJvmConfigurationRequest
//
// @return GetJvmConfigurationResponse
func (client *Client) GetJvmConfiguration(request *GetJvmConfigurationRequest) (_result *GetJvmConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJvmConfigurationResponse{}
	_body, _err := client.GetJvmConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precheck result of a Kubernetes application.
//
// @param request - GetK8sAppPrecheckResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetK8sAppPrecheckResultResponse
func (client *Client) GetK8sAppPrecheckResultWithOptions(request *GetK8sAppPrecheckResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetK8sAppPrecheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetK8sAppPrecheckResult"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/app_precheck"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetK8sAppPrecheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck result of a Kubernetes application.
//
// @param request - GetK8sAppPrecheckResultRequest
//
// @return GetK8sAppPrecheckResultResponse
func (client *Client) GetK8sAppPrecheckResult(request *GetK8sAppPrecheckResultRequest) (_result *GetK8sAppPrecheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetK8sAppPrecheckResultResponse{}
	_body, _err := client.GetK8sAppPrecheckResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about applications deployed in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - GetK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetK8sApplicationResponse
func (client *Client) GetK8sApplicationWithOptions(request *GetK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_application"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about applications deployed in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - GetK8sApplicationRequest
//
// @return GetK8sApplicationResponse
func (client *Client) GetK8sApplication(request *GetK8sApplicationRequest) (_result *GetK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetK8sApplicationResponse{}
	_body, _err := client.GetK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Container Service for Kubernetes (ACK) clusters or Serverless Kubernetes clusters in a specified region.
//
// @param request - GetK8sClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetK8sClusterResponse
func (client *Client) GetK8sClusterWithOptions(request *GetK8sClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetK8sClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionTag) {
		query["RegionTag"] = request.RegionTag
	}

	if !dara.IsNil(request.SubClusterType) {
		query["SubClusterType"] = request.SubClusterType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetK8sCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s_clusters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetK8sClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Container Service for Kubernetes (ACK) clusters or Serverless Kubernetes clusters in a specified region.
//
// @param request - GetK8sClusterRequest
//
// @return GetK8sClusterResponse
func (client *Client) GetK8sCluster(request *GetK8sClusterRequest) (_result *GetK8sClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetK8sClusterResponse{}
	_body, _err := client.GetK8sClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries application services that are deployed in a Kubernetes cluster.
//
// @param request - GetK8sServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetK8sServicesResponse
func (client *Client) GetK8sServicesWithOptions(request *GetK8sServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetK8sServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetK8sServices"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_service"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetK8sServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries application services that are deployed in a Kubernetes cluster.
//
// @param request - GetK8sServicesRequest
//
// @return GetK8sServicesResponse
func (client *Client) GetK8sServices(request *GetK8sServicesRequest) (_result *GetK8sServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetK8sServicesResponse{}
	_body, _err := client.GetK8sServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Security Token Service (STS) tokens that are required for temporary storage.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPackageStorageCredentialResponse
func (client *Client) GetPackageStorageCredentialWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPackageStorageCredentialResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPackageStorageCredential"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/package_storage_credential"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPackageStorageCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Security Token Service (STS) tokens that are required for temporary storage.
//
// @return GetPackageStorageCredentialResponse
func (client *Client) GetPackageStorageCredential() (_result *GetPackageStorageCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPackageStorageCredentialResponse{}
	_body, _err := client.GetPackageStorageCredentialWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries scaling rules.
//
// @param request - GetScalingRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScalingRulesResponse
func (client *Client) GetScalingRulesWithOptions(request *GetScalingRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScalingRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/scalingRules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scaling rules.
//
// @param request - GetScalingRulesRequest
//
// @return GetScalingRulesResponse
func (client *Client) GetScalingRules(request *GetScalingRulesRequest) (_result *GetScalingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScalingRulesResponse{}
	_body, _err := client.GetScalingRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security token information of a namespace. You can call this operation to query information, such as the AccessKey ID, AccessKey secret, tenant ID, and the domain name of Address Server, for the specified namespace.
//
// @param request - GetSecureTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecureTokenResponse
func (client *Client) GetSecureTokenWithOptions(request *GetSecureTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSecureTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecureToken"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/secure_token"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecureTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security token information of a namespace. You can call this operation to query information, such as the AccessKey ID, AccessKey secret, tenant ID, and the domain name of Address Server, for the specified namespace.
//
// @param request - GetSecureTokenRequest
//
// @return GetSecureTokenResponse
func (client *Client) GetSecureToken(request *GetSecureTokenRequest) (_result *GetSecureTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSecureTokenResponse{}
	_body, _err := client.GetSecureTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service consumers.
//
// @param request - GetServiceConsumersPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceConsumersPageResponse
func (client *Client) GetServiceConsumersPageWithOptions(request *GetServiceConsumersPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceConsumersPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["appId"] = request.AppId
	}

	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.RegistryType) {
		query["registryType"] = request.RegistryType
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["serviceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceConsumersPage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/api/mseForOam/getServiceConsumersPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceConsumersPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service consumers.
//
// @param request - GetServiceConsumersPageRequest
//
// @return GetServiceConsumersPageResponse
func (client *Client) GetServiceConsumersPage(request *GetServiceConsumersPageRequest) (_result *GetServiceConsumersPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceConsumersPageResponse{}
	_body, _err := client.GetServiceConsumersPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service details.
//
// @param request - GetServiceDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceDetailResponse
func (client *Client) GetServiceDetailWithOptions(request *GetServiceDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["appId"] = request.AppId
	}

	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.RegistryType) {
		query["registryType"] = request.RegistryType
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["serviceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceDetail"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/api/mseForOam/getServiceDetail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service details.
//
// @param request - GetServiceDetailRequest
//
// @return GetServiceDetailResponse
func (client *Client) GetServiceDetail(request *GetServiceDetailRequest) (_result *GetServiceDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceDetailResponse{}
	_body, _err := client.GetServiceDetailWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries services.
//
// @param request - GetServiceListPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceListPageResponse
func (client *Client) GetServiceListPageWithOptions(request *GetServiceListPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceListPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.SearchType) {
		query["searchType"] = request.SearchType
	}

	if !dara.IsNil(request.SearchValue) {
		query["searchValue"] = request.SearchValue
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.Side) {
		query["side"] = request.Side
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceListPage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/api/mseForOam/getServiceListPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceListPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries services.
//
// @param request - GetServiceListPageRequest
//
// @return GetServiceListPageResponse
func (client *Client) GetServiceListPage(request *GetServiceListPageRequest) (_result *GetServiceListPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceListPageResponse{}
	_body, _err := client.GetServiceListPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service methods.
//
// @param request - GetServiceMethodPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceMethodPageResponse
func (client *Client) GetServiceMethodPageWithOptions(request *GetServiceMethodPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceMethodPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["appId"] = request.AppId
	}

	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.MethodController) {
		query["methodController"] = request.MethodController
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Path) {
		query["path"] = request.Path
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.RegistryType) {
		query["registryType"] = request.RegistryType
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["serviceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceMethodPage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/api/mseForOam/getServiceMethodPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceMethodPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service methods.
//
// @param request - GetServiceMethodPageRequest
//
// @return GetServiceMethodPageResponse
func (client *Client) GetServiceMethodPage(request *GetServiceMethodPageRequest) (_result *GetServiceMethodPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceMethodPageResponse{}
	_body, _err := client.GetServiceMethodPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service providers.
//
// @param request - GetServiceProvidersPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvidersPageResponse
func (client *Client) GetServiceProvidersPageWithOptions(request *GetServiceProvidersPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceProvidersPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["appId"] = request.AppId
	}

	if !dara.IsNil(request.Group) {
		query["group"] = request.Group
	}

	if !dara.IsNil(request.Ip) {
		query["ip"] = request.Ip
	}

	if !dara.IsNil(request.Namespace) {
		query["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Origin) {
		query["origin"] = request.Origin
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.RegistryType) {
		query["registryType"] = request.RegistryType
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceName) {
		query["serviceName"] = request.ServiceName
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["serviceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceProvidersPage"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/api/mseForOam/getServiceProvidersPage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceProvidersPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service providers.
//
// @param request - GetServiceProvidersPageRequest
//
// @return GetServiceProvidersPageResponse
func (client *Client) GetServiceProvidersPage(request *GetServiceProvidersPageRequest) (_result *GetServiceProvidersPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceProvidersPageResponse{}
	_body, _err := client.GetServiceProvidersPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Tomcat configurations of an application.
//
// Description:
//
// **
//
// @param request - GetWebContainerConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebContainerConfigResponse
func (client *Client) GetWebContainerConfigWithOptions(request *GetWebContainerConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWebContainerConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWebContainerConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/web_container_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWebContainerConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Tomcat configurations of an application.
//
// Description:
//
// **
//
// @param request - GetWebContainerConfigRequest
//
// @return GetWebContainerConfigResponse
func (client *Client) GetWebContainerConfig(request *GetWebContainerConfigRequest) (_result *GetWebContainerConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWebContainerConfigResponse{}
	_body, _err := client.GetWebContainerConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster to Enterprise Distributed Application Service (EDAS).
//
// @param request - ImportK8sClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportK8sClusterResponse
func (client *Client) ImportK8sClusterWithOptions(request *ImportK8sClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportK8sClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EnableAsm) {
		query["EnableAsm"] = request.EnableAsm
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportK8sCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/import_k8s_cluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportK8sClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster to Enterprise Distributed Application Service (EDAS).
//
// @param request - ImportK8sClusterRequest
//
// @return ImportK8sClusterResponse
func (client *Client) ImportK8sCluster(request *ImportK8sClusterRequest) (_result *ImportK8sClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportK8sClusterResponse{}
	_body, _err := client.ImportK8sClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application in an Elastic Compute Service (ECS) cluster.
//
// Description:
//
// > To create an application in a Kubernetes cluster, call the InsertK8sApplication operation provided by Enterprise Distributed Application Service (EDAS).
//
// @param request - InsertApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertApplicationResponse
func (client *Client) InsertApplicationWithOptions(request *InsertApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.BuildPackId) {
		query["BuildPackId"] = request.BuildPackId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentIds) {
		query["ComponentIds"] = request.ComponentIds
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EcuInfo) {
		query["EcuInfo"] = request.EcuInfo
	}

	if !dara.IsNil(request.EnablePortCheck) {
		query["EnablePortCheck"] = request.EnablePortCheck
	}

	if !dara.IsNil(request.EnableUrlCheck) {
		query["EnableUrlCheck"] = request.EnableUrlCheck
	}

	if !dara.IsNil(request.HealthCheckUrl) {
		query["HealthCheckUrl"] = request.HealthCheckUrl
	}

	if !dara.IsNil(request.Hooks) {
		query["Hooks"] = request.Hooks
	}

	if !dara.IsNil(request.Jdk) {
		query["Jdk"] = request.Jdk
	}

	if !dara.IsNil(request.JvmOptions) {
		query["JvmOptions"] = request.JvmOptions
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.MaxHeapSize) {
		query["MaxHeapSize"] = request.MaxHeapSize
	}

	if !dara.IsNil(request.MaxPermSize) {
		query["MaxPermSize"] = request.MaxPermSize
	}

	if !dara.IsNil(request.Mem) {
		query["Mem"] = request.Mem
	}

	if !dara.IsNil(request.MinHeapSize) {
		query["MinHeapSize"] = request.MinHeapSize
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.ReservedPortStr) {
		query["ReservedPortStr"] = request.ReservedPortStr
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_create_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application in an Elastic Compute Service (ECS) cluster.
//
// Description:
//
// > To create an application in a Kubernetes cluster, call the InsertK8sApplication operation provided by Enterprise Distributed Application Service (EDAS).
//
// @param request - InsertApplicationRequest
//
// @return InsertApplicationResponse
func (client *Client) InsertApplication(request *InsertApplicationRequest) (_result *InsertApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertApplicationResponse{}
	_body, _err := client.InsertApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// @param request - InsertClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertClusterResponse
func (client *Client) InsertClusterWithOptions(request *InsertClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.IaasProvider) {
		query["IaasProvider"] = request.IaasProvider
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.NetworkMode) {
		query["NetworkMode"] = request.NetworkMode
	}

	if !dara.IsNil(request.OversoldFactor) {
		query["OversoldFactor"] = request.OversoldFactor
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster.
//
// @param request - InsertClusterRequest
//
// @return InsertClusterResponse
func (client *Client) InsertCluster(request *InsertClusterRequest) (_result *InsertClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertClusterResponse{}
	_body, _err := client.InsertClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports Elastic Compute Service (ECS) instances into an ECS cluster.
//
// Description:
//
// ##
//
// If you call this operation to import an ECS instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all original data of the ECS instance is deleted. In addition, you must set a logon password for the ECS instance. Make sure that no important data exists on the ECS instance that you want to import or data has been backed up for the ECS instance.
//
// > We recommend that you call the InstallAgent operation instead of this operation. For more information, see [InstallAgent](https://help.aliyun.com/document_detail/127023.html).
//
// @param request - InsertClusterMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertClusterMemberResponse
func (client *Client) InsertClusterMemberWithOptions(request *InsertClusterMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertClusterMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["clusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.InstanceIds) {
		query["instanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Password) {
		query["password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertClusterMember"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster_member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertClusterMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports Elastic Compute Service (ECS) instances into an ECS cluster.
//
// Description:
//
// ##
//
// If you call this operation to import an ECS instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all original data of the ECS instance is deleted. In addition, you must set a logon password for the ECS instance. Make sure that no important data exists on the ECS instance that you want to import or data has been backed up for the ECS instance.
//
// > We recommend that you call the InstallAgent operation instead of this operation. For more information, see [InstallAgent](https://help.aliyun.com/document_detail/127023.html).
//
// @param request - InsertClusterMemberRequest
//
// @return InsertClusterMemberResponse
func (client *Client) InsertClusterMember(request *InsertClusterMemberRequest) (_result *InsertClusterMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertClusterMemberResponse{}
	_body, _err := client.InsertClusterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instance group for a specified application.
//
// @param request - InsertDeployGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertDeployGroupResponse
func (client *Client) InsertDeployGroupWithOptions(request *InsertDeployGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertDeployGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InitPackageVersionId) {
		query["InitPackageVersionId"] = request.InitPackageVersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertDeployGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/deploy_group"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertDeployGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instance group for a specified application.
//
// @param request - InsertDeployGroupRequest
//
// @return InsertDeployGroupResponse
func (client *Client) InsertDeployGroup(request *InsertDeployGroupRequest) (_result *InsertDeployGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertDeployGroupResponse{}
	_body, _err := client.InsertDeployGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application in a Container Service for Kubernetes (ACK) cluster or serverless Kubernetes cluster.
//
// @param request - InsertK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertK8sApplicationResponse
func (client *Client) InsertK8sApplicationWithOptions(request *InsertK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.AppConfig) {
		query["AppConfig"] = request.AppConfig
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AppTemplateName) {
		query["AppTemplateName"] = request.AppTemplateName
	}

	if !dara.IsNil(request.ApplicationDescription) {
		query["ApplicationDescription"] = request.ApplicationDescription
	}

	if !dara.IsNil(request.BuildPackId) {
		query["BuildPackId"] = request.BuildPackId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.CommandArgs) {
		query["CommandArgs"] = request.CommandArgs
	}

	if !dara.IsNil(request.ConfigMountDescs) {
		query["ConfigMountDescs"] = request.ConfigMountDescs
	}

	if !dara.IsNil(request.ContainerRegistryId) {
		query["ContainerRegistryId"] = request.ContainerRegistryId
	}

	if !dara.IsNil(request.CsClusterId) {
		query["CsClusterId"] = request.CsClusterId
	}

	if !dara.IsNil(request.CustomAffinity) {
		query["CustomAffinity"] = request.CustomAffinity
	}

	if !dara.IsNil(request.CustomAgentVersion) {
		query["CustomAgentVersion"] = request.CustomAgentVersion
	}

	if !dara.IsNil(request.CustomTolerations) {
		query["CustomTolerations"] = request.CustomTolerations
	}

	if !dara.IsNil(request.DeployAcrossNodes) {
		query["DeployAcrossNodes"] = request.DeployAcrossNodes
	}

	if !dara.IsNil(request.DeployAcrossZones) {
		query["DeployAcrossZones"] = request.DeployAcrossZones
	}

	if !dara.IsNil(request.EdasContainerVersion) {
		query["EdasContainerVersion"] = request.EdasContainerVersion
	}

	if !dara.IsNil(request.EmptyDirs) {
		query["EmptyDirs"] = request.EmptyDirs
	}

	if !dara.IsNil(request.EnableAhas) {
		query["EnableAhas"] = request.EnableAhas
	}

	if !dara.IsNil(request.EnableAsm) {
		query["EnableAsm"] = request.EnableAsm
	}

	if !dara.IsNil(request.EnableEmptyPushReject) {
		query["EnableEmptyPushReject"] = request.EnableEmptyPushReject
	}

	if !dara.IsNil(request.EnableLosslessRule) {
		query["EnableLosslessRule"] = request.EnableLosslessRule
	}

	if !dara.IsNil(request.EnvFroms) {
		query["EnvFroms"] = request.EnvFroms
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.FeatureConfig) {
		query["FeatureConfig"] = request.FeatureConfig
	}

	if !dara.IsNil(request.ImagePlatforms) {
		query["ImagePlatforms"] = request.ImagePlatforms
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.InitContainers) {
		query["InitContainers"] = request.InitContainers
	}

	if !dara.IsNil(request.InternetSlbId) {
		query["InternetSlbId"] = request.InternetSlbId
	}

	if !dara.IsNil(request.InternetSlbPort) {
		query["InternetSlbPort"] = request.InternetSlbPort
	}

	if !dara.IsNil(request.InternetSlbProtocol) {
		query["InternetSlbProtocol"] = request.InternetSlbProtocol
	}

	if !dara.IsNil(request.InternetTargetPort) {
		query["InternetTargetPort"] = request.InternetTargetPort
	}

	if !dara.IsNil(request.IntranetSlbId) {
		query["IntranetSlbId"] = request.IntranetSlbId
	}

	if !dara.IsNil(request.IntranetSlbPort) {
		query["IntranetSlbPort"] = request.IntranetSlbPort
	}

	if !dara.IsNil(request.IntranetSlbProtocol) {
		query["IntranetSlbProtocol"] = request.IntranetSlbProtocol
	}

	if !dara.IsNil(request.IntranetTargetPort) {
		query["IntranetTargetPort"] = request.IntranetTargetPort
	}

	if !dara.IsNil(request.IsMultilingualApp) {
		query["IsMultilingualApp"] = request.IsMultilingualApp
	}

	if !dara.IsNil(request.JDK) {
		query["JDK"] = request.JDK
	}

	if !dara.IsNil(request.JavaStartUpConfig) {
		query["JavaStartUpConfig"] = request.JavaStartUpConfig
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.LimitCpu) {
		query["LimitCpu"] = request.LimitCpu
	}

	if !dara.IsNil(request.LimitEphemeralStorage) {
		query["LimitEphemeralStorage"] = request.LimitEphemeralStorage
	}

	if !dara.IsNil(request.LimitMem) {
		query["LimitMem"] = request.LimitMem
	}

	if !dara.IsNil(request.LimitmCpu) {
		query["LimitmCpu"] = request.LimitmCpu
	}

	if !dara.IsNil(request.Liveness) {
		query["Liveness"] = request.Liveness
	}

	if !dara.IsNil(request.LocalVolume) {
		query["LocalVolume"] = request.LocalVolume
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.LosslessRuleAligned) {
		query["LosslessRuleAligned"] = request.LosslessRuleAligned
	}

	if !dara.IsNil(request.LosslessRuleDelayTime) {
		query["LosslessRuleDelayTime"] = request.LosslessRuleDelayTime
	}

	if !dara.IsNil(request.LosslessRuleFuncType) {
		query["LosslessRuleFuncType"] = request.LosslessRuleFuncType
	}

	if !dara.IsNil(request.LosslessRuleRelated) {
		query["LosslessRuleRelated"] = request.LosslessRuleRelated
	}

	if !dara.IsNil(request.LosslessRuleWarmupTime) {
		query["LosslessRuleWarmupTime"] = request.LosslessRuleWarmupTime
	}

	if !dara.IsNil(request.MountDescs) {
		query["MountDescs"] = request.MountDescs
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.NasId) {
		query["NasId"] = request.NasId
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PackageVersion) {
		query["PackageVersion"] = request.PackageVersion
	}

	if !dara.IsNil(request.PostStart) {
		query["PostStart"] = request.PostStart
	}

	if !dara.IsNil(request.PreStop) {
		query["PreStop"] = request.PreStop
	}

	if !dara.IsNil(request.PvcMountDescs) {
		query["PvcMountDescs"] = request.PvcMountDescs
	}

	if !dara.IsNil(request.Readiness) {
		query["Readiness"] = request.Readiness
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.RepoId) {
		query["RepoId"] = request.RepoId
	}

	if !dara.IsNil(request.RequestsCpu) {
		query["RequestsCpu"] = request.RequestsCpu
	}

	if !dara.IsNil(request.RequestsEphemeralStorage) {
		query["RequestsEphemeralStorage"] = request.RequestsEphemeralStorage
	}

	if !dara.IsNil(request.RequestsMem) {
		query["RequestsMem"] = request.RequestsMem
	}

	if !dara.IsNil(request.RequestsmCpu) {
		query["RequestsmCpu"] = request.RequestsmCpu
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RuntimeClassName) {
		query["RuntimeClassName"] = request.RuntimeClassName
	}

	if !dara.IsNil(request.SecretName) {
		query["SecretName"] = request.SecretName
	}

	if !dara.IsNil(request.SecurityContext) {
		query["SecurityContext"] = request.SecurityContext
	}

	if !dara.IsNil(request.ServiceConfigs) {
		query["ServiceConfigs"] = request.ServiceConfigs
	}

	if !dara.IsNil(request.Sidecars) {
		query["Sidecars"] = request.Sidecars
	}

	if !dara.IsNil(request.SlsConfigs) {
		query["SlsConfigs"] = request.SlsConfigs
	}

	if !dara.IsNil(request.Startup) {
		query["Startup"] = request.Startup
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.TerminateGracePeriod) {
		query["TerminateGracePeriod"] = request.TerminateGracePeriod
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UriEncoding) {
		query["UriEncoding"] = request.UriEncoding
	}

	if !dara.IsNil(request.UseBodyEncoding) {
		query["UseBodyEncoding"] = request.UseBodyEncoding
	}

	if !dara.IsNil(request.UserBaseImageUrl) {
		query["UserBaseImageUrl"] = request.UserBaseImageUrl
	}

	if !dara.IsNil(request.WebContainer) {
		query["WebContainer"] = request.WebContainer
	}

	if !dara.IsNil(request.WebContainerConfig) {
		query["WebContainerConfig"] = request.WebContainerConfig
	}

	if !dara.IsNil(request.WorkloadType) {
		query["WorkloadType"] = request.WorkloadType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/create_k8s_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application in a Container Service for Kubernetes (ACK) cluster or serverless Kubernetes cluster.
//
// @param request - InsertK8sApplicationRequest
//
// @return InsertK8sApplicationResponse
func (client *Client) InsertK8sApplication(request *InsertK8sApplicationRequest) (_result *InsertK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertK8sApplicationResponse{}
	_body, _err := client.InsertK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or edits a custom namespace.
//
// @param request - InsertOrUpdateRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertOrUpdateRegionResponse
func (client *Client) InsertOrUpdateRegionWithOptions(request *InsertOrUpdateRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertOrUpdateRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DebugEnable) {
		query["DebugEnable"] = request.DebugEnable
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MseInstanceId) {
		query["MseInstanceId"] = request.MseInstanceId
	}

	if !dara.IsNil(request.RegionName) {
		query["RegionName"] = request.RegionName
	}

	if !dara.IsNil(request.RegionTag) {
		query["RegionTag"] = request.RegionTag
	}

	if !dara.IsNil(request.RegistryType) {
		query["RegistryType"] = request.RegistryType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertOrUpdateRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/user_region_def"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertOrUpdateRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or edits a custom namespace.
//
// @param request - InsertOrUpdateRegionRequest
//
// @return InsertOrUpdateRegionResponse
func (client *Client) InsertOrUpdateRegion(request *InsertOrUpdateRegionRequest) (_result *InsertOrUpdateRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertOrUpdateRegionResponse{}
	_body, _err := client.InsertOrUpdateRegionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a role.
//
// @param request - InsertRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertRoleResponse
func (client *Client) InsertRoleWithOptions(request *InsertRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionData) {
		query["ActionData"] = request.ActionData
	}

	if !dara.IsNil(request.RoleName) {
		query["RoleName"] = request.RoleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/create_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a role.
//
// @param request - InsertRoleRequest
//
// @return InsertRoleResponse
func (client *Client) InsertRole(request *InsertRoleRequest) (_result *InsertRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertRoleResponse{}
	_body, _err := client.InsertRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service group.
//
// @param request - InsertServiceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertServiceGroupResponse
func (client *Client) InsertServiceGroupWithOptions(request *InsertServiceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertServiceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertServiceGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/serviceGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service group.
//
// @param request - InsertServiceGroupRequest
//
// @return InsertServiceGroupResponse
func (client *Client) InsertServiceGroup(request *InsertServiceGroupRequest) (_result *InsertServiceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertServiceGroupResponse{}
	_body, _err := client.InsertServiceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lane.
//
// @param request - InsertSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertSwimmingLaneResponse
func (client *Client) InsertSwimmingLaneWithOptions(request *InsertSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInfos) {
		query["AppInfos"] = request.AppInfos
	}

	if !dara.IsNil(request.EnableRules) {
		query["EnableRules"] = request.EnableRules
	}

	if !dara.IsNil(request.EntryRules) {
		query["EntryRules"] = request.EntryRules
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertSwimmingLane"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lanes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lane.
//
// @param request - InsertSwimmingLaneRequest
//
// @return InsertSwimmingLaneResponse
func (client *Client) InsertSwimmingLane(request *InsertSwimmingLaneRequest) (_result *InsertSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertSwimmingLaneResponse{}
	_body, _err := client.InsertSwimmingLaneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lane group.
//
// @param request - InsertSwimmingLaneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertSwimmingLaneGroupResponse
func (client *Client) InsertSwimmingLaneGroupWithOptions(request *InsertSwimmingLaneGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InsertSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.EntryApp) {
		query["EntryApp"] = request.EntryApp
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertSwimmingLaneGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lane_groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lane group.
//
// @param request - InsertSwimmingLaneGroupRequest
//
// @return InsertSwimmingLaneGroupResponse
func (client *Client) InsertSwimmingLaneGroup(request *InsertSwimmingLaneGroupRequest) (_result *InsertSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InsertSwimmingLaneGroupResponse{}
	_body, _err := client.InsertSwimmingLaneGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses the Cloud Assistant provided by Elastic Compute Service (ECS) to install Enterprise Distributed Application Service (EDAS) Agent and imports ECS instances to EDAS.
//
// Description:
//
// If you call this operation to import an ECS instance into EDAS, the operating system of the ECS instance is not reinstalled. We recommend that you call this operation to import ECS instances into EDAS.
//
// @param request - InstallAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallAgentResponse
func (client *Client) InstallAgentWithOptions(request *InstallAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DoAsync) {
		query["DoAsync"] = request.DoAsync
	}

	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallAgent"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/ecss/install_agent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses the Cloud Assistant provided by Elastic Compute Service (ECS) to install Enterprise Distributed Application Service (EDAS) Agent and imports ECS instances to EDAS.
//
// Description:
//
// If you call this operation to import an ECS instance into EDAS, the operating system of the ECS instance is not reinstalled. We recommend that you call this operation to import ECS instances into EDAS.
//
// @param request - InstallAgentRequest
//
// @return InstallAgentResponse
func (client *Client) InstallAgent(request *InstallAgentRequest) (_result *InstallAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallAgentResponse{}
	_body, _err := client.InstallAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Alibaba Cloud regions supported by Enterprise Distributed Application Service (EDAS).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAliyunRegionResponse
func (client *Client) ListAliyunRegionWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAliyunRegionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAliyunRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/region_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAliyunRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Alibaba Cloud regions supported by Enterprise Distributed Application Service (EDAS).
//
// @return ListAliyunRegionResponse
func (client *Client) ListAliyunRegion() (_result *ListAliyunRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAliyunRegionResponse{}
	_body, _err := client.ListAliyunRegionWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of applications.
//
// @param request - ListApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationResponse
func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.LogicalRegionIdFilter) {
		query["LogicalRegionIdFilter"] = request.LogicalRegionIdFilter
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/app_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of applications.
//
// @param request - ListApplicationRequest
//
// @return ListApplicationResponse
func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries elastic compute units (ECUs).
//
// @param request - ListApplicationEcuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationEcuResponse
func (client *Client) ListApplicationEcuWithOptions(request *ListApplicationEcuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApplicationEcuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationEcu"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/ecu_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationEcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries elastic compute units (ECUs).
//
// @param request - ListApplicationEcuRequest
//
// @return ListApplicationEcuResponse
func (client *Client) ListApplicationEcu(request *ListApplicationEcuRequest) (_result *ListApplicationEcuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApplicationEcuResponse{}
	_body, _err := client.ListApplicationEcuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all permissions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorityResponse
func (client *Client) ListAuthorityWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAuthorityResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthority"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/authority_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all permissions.
//
// @return ListAuthorityResponse
func (client *Client) ListAuthority() (_result *ListAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAuthorityResponse{}
	_body, _err := client.ListAuthorityWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Enterprise Distributed Application Service (EDAS) Container versions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBuildPackResponse
func (client *Client) ListBuildPackWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBuildPackResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBuildPack"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/build_pack_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBuildPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Enterprise Distributed Application Service (EDAS) Container versions.
//
// @return ListBuildPackResponse
func (client *Client) ListBuildPack() (_result *ListBuildPackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBuildPackResponse{}
	_body, _err := client.ListBuildPackWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries clusters.
//
// @param request - ListClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterResponse
func (client *Client) ListClusterWithOptions(request *ListClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries clusters.
//
// @param request - ListClusterRequest
//
// @return ListClusterResponse
func (client *Client) ListCluster(request *ListClusterRequest) (_result *ListClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterResponse{}
	_body, _err := client.ListClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Elastic Compute Service (ECS) instances in a cluster.
//
// @param request - ListClusterMembersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterMembersResponse
func (client *Client) ListClusterMembersWithOptions(request *ListClusterMembersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListClusterMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EcsList) {
		query["EcsList"] = request.EcsList
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterMembers"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/cluster_member_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Elastic Compute Service (ECS) instances in a cluster.
//
// @param request - ListClusterMembersRequest
//
// @return ListClusterMembersResponse
func (client *Client) ListClusterMembers(request *ListClusterMembersRequest) (_result *ListClusterMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterMembersResponse{}
	_body, _err := client.ListClusterMembersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the components that are related to applications in Elastic Compute Service (ECS) clusters.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/components"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the components that are related to applications in Elastic Compute Service (ECS) clusters.
//
// @return ListComponentsResponse
func (client *Client) ListComponents() (_result *ListComponentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries configuration templates.
//
// @param request - ListConfigTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigTemplatesResponse
func (client *Client) ListConfigTemplatesWithOptions(request *ListConfigTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConfigTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigTemplates"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/config_template"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration templates.
//
// @param request - ListConfigTemplatesRequest
//
// @return ListConfigTemplatesResponse
func (client *Client) ListConfigTemplates(request *ListConfigTemplatesRequest) (_result *ListConfigTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigTemplatesResponse{}
	_body, _err := client.ListConfigTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the services that are consumed by an application.
//
// @param request - ListConsumedServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumedServicesResponse
func (client *Client) ListConsumedServicesWithOptions(request *ListConsumedServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumedServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumedServices"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/listConsumedServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services that are consumed by an application.
//
// @param request - ListConsumedServicesRequest
//
// @return ListConsumedServicesResponse
func (client *Client) ListConsumedServices(request *ListConsumedServicesRequest) (_result *ListConsumedServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumedServicesResponse{}
	_body, _err := client.ListConsumedServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Elastic Compute Service (ECS) instances that can be imported to a specified cluster. This operation is applicable to ECS clusters.
//
// @param request - ListConvertableEcuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConvertableEcuResponse
func (client *Client) ListConvertableEcuWithOptions(request *ListConvertableEcuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConvertableEcuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["clusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConvertableEcu"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/convertable_ecu_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConvertableEcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Elastic Compute Service (ECS) instances that can be imported to a specified cluster. This operation is applicable to ECS clusters.
//
// @param request - ListConvertableEcuRequest
//
// @return ListConvertableEcuResponse
func (client *Client) ListConvertableEcu(request *ListConvertableEcuRequest) (_result *ListConvertableEcuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConvertableEcuResponse{}
	_body, _err := client.ListConvertableEcuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instance groups to which an application is deployed.
//
// @param request - ListDeployGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeployGroupResponse
func (client *Client) ListDeployGroupWithOptions(request *ListDeployGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDeployGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeployGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/deploy_group_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeployGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instance groups to which an application is deployed.
//
// @param request - ListDeployGroupRequest
//
// @return ListDeployGroupResponse
func (client *Client) ListDeployGroup(request *ListDeployGroupRequest) (_result *ListDeployGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeployGroupResponse{}
	_body, _err := client.ListDeployGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all Elastic Compute Service (ECS) instances that have not been imported to clusters.
//
// @param request - ListEcsNotInClusterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsNotInClusterResponse
func (client *Client) ListEcsNotInClusterWithOptions(request *ListEcsNotInClusterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsNotInClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkMode) {
		query["NetworkMode"] = request.NetworkMode
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEcsNotInCluster"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/ecs_not_in_cluster"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEcsNotInClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all Elastic Compute Service (ECS) instances that have not been imported to clusters.
//
// @param request - ListEcsNotInClusterRequest
//
// @return ListEcsNotInClusterResponse
func (client *Client) ListEcsNotInCluster(request *ListEcsNotInClusterRequest) (_result *ListEcsNotInClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsNotInClusterResponse{}
	_body, _err := client.ListEcsNotInClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available elastic compute units (ECUs) in a specified namespace.
//
// Description:
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources such as clusters, ECS instances, and applications, and microservices published in EDAS. This concept involves the default namespace and custom namespaces. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources and microservices.
//
//   - **Elastic compute unit (ECU)**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - ListEcuByRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcuByRegionResponse
func (client *Client) ListEcuByRegionWithOptions(request *ListEcuByRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcuByRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Act) {
		query["Act"] = request.Act
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEcuByRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/ecu_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEcuByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available elastic compute units (ECUs) in a specified namespace.
//
// Description:
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources such as clusters, ECS instances, and applications, and microservices published in EDAS. This concept involves the default namespace and custom namespaces. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources and microservices.
//
//   - **Elastic compute unit (ECU)**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - ListEcuByRegionRequest
//
// @return ListEcuByRegionResponse
func (client *Client) ListEcuByRegion(request *ListEcuByRegionRequest) (_result *ListEcuByRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcuByRegionResponse{}
	_body, _err := client.ListEcuByRegionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries historical deployment packages of an application.
//
// @param request - ListHistoryDeployVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHistoryDeployVersionResponse
func (client *Client) ListHistoryDeployVersionWithOptions(request *ListHistoryDeployVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHistoryDeployVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHistoryDeployVersion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/deploy_history_version_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHistoryDeployVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries historical deployment packages of an application.
//
// @param request - ListHistoryDeployVersionRequest
//
// @return ListHistoryDeployVersionResponse
func (client *Client) ListHistoryDeployVersion(request *ListHistoryDeployVersionRequest) (_result *ListHistoryDeployVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHistoryDeployVersionResponse{}
	_body, _err := client.ListHistoryDeployVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Kubernetes ConfigMaps.
//
// @param request - ListK8sConfigMapsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sConfigMapsResponse
func (client *Client) ListK8sConfigMapsWithOptions(request *ListK8sConfigMapsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListK8sConfigMapsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
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

	if !dara.IsNil(request.ShowRelatedApps) {
		query["ShowRelatedApps"] = request.ShowRelatedApps
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListK8sConfigMaps"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_config_map"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListK8sConfigMapsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Kubernetes ConfigMaps.
//
// @param request - ListK8sConfigMapsRequest
//
// @return ListK8sConfigMapsResponse
func (client *Client) ListK8sConfigMaps(request *ListK8sConfigMapsRequest) (_result *ListK8sConfigMapsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListK8sConfigMapsResponse{}
	_body, _err := client.ListK8sConfigMapsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries ingresses.
//
// @param request - ListK8sIngressRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sIngressRulesResponse
func (client *Client) ListK8sIngressRulesWithOptions(request *ListK8sIngressRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListK8sIngressRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListK8sIngressRules"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_ingress"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListK8sIngressRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries ingresses.
//
// @param request - ListK8sIngressRulesRequest
//
// @return ListK8sIngressRulesResponse
func (client *Client) ListK8sIngressRules(request *ListK8sIngressRulesRequest) (_result *ListK8sIngressRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListK8sIngressRulesResponse{}
	_body, _err := client.ListK8sIngressRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries namespaces for a Kubernetes cluster.
//
// @param request - ListK8sNamespacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sNamespacesResponse
func (client *Client) ListK8sNamespacesWithOptions(request *ListK8sNamespacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListK8sNamespacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListK8sNamespaces"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_namespace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListK8sNamespacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries namespaces for a Kubernetes cluster.
//
// @param request - ListK8sNamespacesRequest
//
// @return ListK8sNamespacesResponse
func (client *Client) ListK8sNamespaces(request *ListK8sNamespacesRequest) (_result *ListK8sNamespacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListK8sNamespacesResponse{}
	_body, _err := client.ListK8sNamespacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Kubernetes Secrets.
//
// @param request - ListK8sSecretsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListK8sSecretsResponse
func (client *Client) ListK8sSecretsWithOptions(request *ListK8sSecretsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListK8sSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Condition) {
		query["Condition"] = request.Condition
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
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

	if !dara.IsNil(request.ShowRelatedApps) {
		query["ShowRelatedApps"] = request.ShowRelatedApps
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListK8sSecrets"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_secret"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListK8sSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Kubernetes Secrets.
//
// @param request - ListK8sSecretsRequest
//
// @return ListK8sSecretsResponse
func (client *Client) ListK8sSecrets(request *ListK8sSecretsRequest) (_result *ListK8sSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListK8sSecretsResponse{}
	_body, _err := client.ListK8sSecretsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries service methods.
//
// @param request - ListMethodsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMethodsResponse
func (client *Client) ListMethodsWithOptions(request *ListMethodsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMethodsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ServiceName) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMethods"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/list_methods"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMethodsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries service methods.
//
// @param request - ListMethodsRequest
//
// @return ListMethodsResponse
func (client *Client) ListMethods(request *ListMethodsRequest) (_result *ListMethodsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMethodsResponse{}
	_body, _err := client.ListMethodsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the services that are published by an application.
//
// @param request - ListPublishedServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedServicesResponse
func (client *Client) ListPublishedServicesWithOptions(request *ListPublishedServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPublishedServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublishedServices"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/listPublishedServices"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the services that are published by an application.
//
// @param request - ListPublishedServicesRequest
//
// @return ListPublishedServicesResponse
func (client *Client) ListPublishedServices(request *ListPublishedServicesRequest) (_result *ListPublishedServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPublishedServicesResponse{}
	_body, _err := client.ListPublishedServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the change processes of an application.
//
// @param request - ListRecentChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecentChangeOrderResponse
func (client *Client) ListRecentChangeOrderWithOptions(request *ListRecentChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRecentChangeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecentChangeOrder"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/change_order_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecentChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the change processes of an application.
//
// @param request - ListRecentChangeOrderRequest
//
// @return ListRecentChangeOrderResponse
func (client *Client) ListRecentChangeOrder(request *ListRecentChangeOrderRequest) (_result *ListRecentChangeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecentChangeOrderResponse{}
	_body, _err := client.ListRecentChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries resource groups.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupResponse
func (client *Client) ListResourceGroupWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListResourceGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/reg_group_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries resource groups.
//
// @return ListResourceGroupResponse
func (client *Client) ListResourceGroup() (_result *ListResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceGroupResponse{}
	_body, _err := client.ListResourceGroupWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries roles.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRoleResponse
func (client *Client) ListRoleWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/role_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries roles.
//
// @return ListRoleResponse
func (client *Client) ListRole() (_result *ListRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRoleResponse{}
	_body, _err := client.ListRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries elastic compute units (ECUs) available for scaling out an application in a specified cluster or the cluster where the application is deployed. This operation is applicable to Elastic Compute Service (ECS) clusters.
//
// Description:
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources such as clusters, ECS instances, and applications, and microservices published in EDAS. This concept involves the default namespace and custom namespaces. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources and microservices.
//
//   - **Elastic compute unit (ECU)**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - ListScaleOutEcuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScaleOutEcuResponse
func (client *Client) ListScaleOutEcuWithOptions(request *ListScaleOutEcuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListScaleOutEcuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Cpu) {
		query["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceNum) {
		query["InstanceNum"] = request.InstanceNum
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	if !dara.IsNil(request.Mem) {
		query["Mem"] = request.Mem
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScaleOutEcu"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/scale_out_ecu_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScaleOutEcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries elastic compute units (ECUs) available for scaling out an application in a specified cluster or the cluster where the application is deployed. This operation is applicable to Elastic Compute Service (ECS) clusters.
//
// Description:
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources such as clusters, ECS instances, and applications, and microservices published in EDAS. This concept involves the default namespace and custom namespaces. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources and microservices.
//
//   - **Elastic compute unit (ECU)**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - ListScaleOutEcuRequest
//
// @return ListScaleOutEcuResponse
func (client *Client) ListScaleOutEcu(request *ListScaleOutEcuRequest) (_result *ListScaleOutEcuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScaleOutEcuResponse{}
	_body, _err := client.ListScaleOutEcuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the service groups of a High-Speed Service Framework (HSF) application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceGroupsResponse
func (client *Client) ListServiceGroupsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServiceGroupsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceGroups"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/service/serviceGroups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the service groups of a High-Speed Service Framework (HSF) application.
//
// @return ListServiceGroupsResponse
func (client *Client) ListServiceGroups() (_result *ListServiceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceGroupsResponse{}
	_body, _err := client.ListServiceGroupsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Server Load Balancer (SLB) instances.
//
// @param request - ListSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSlbResponse
func (client *Client) ListSlbWithOptions(request *ListSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.SlbType) {
		query["SlbType"] = request.SlbType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/slb_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Server Load Balancer (SLB) instances.
//
// @param request - ListSlbRequest
//
// @return ListSlbResponse
func (client *Client) ListSlb(request *ListSlbRequest) (_result *ListSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlbResponse{}
	_body, _err := client.ListSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) users.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubAccountResponse
func (client *Client) ListSubAccountWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSubAccountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubAccount"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/sub_account_list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) users.
//
// @return ListSubAccountResponse
func (client *Client) ListSubAccount() (_result *ListSubAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSubAccountResponse{}
	_body, _err := client.ListSubAccountWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries lanes in a lane group.
//
// @param request - ListSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSwimmingLaneResponse
func (client *Client) ListSwimmingLaneWithOptions(request *ListSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSwimmingLane"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lanes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lanes in a lane group.
//
// @param request - ListSwimmingLaneRequest
//
// @return ListSwimmingLaneResponse
func (client *Client) ListSwimmingLane(request *ListSwimmingLaneRequest) (_result *ListSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSwimmingLaneResponse{}
	_body, _err := client.ListSwimmingLaneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries lane groups.
//
// @param request - ListSwimmingLaneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSwimmingLaneGroupResponse
func (client *Client) ListSwimmingLaneGroupWithOptions(request *ListSwimmingLaneGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSwimmingLaneGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lane_groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries lane groups.
//
// @param request - ListSwimmingLaneGroupRequest
//
// @return ListSwimmingLaneGroupResponse
func (client *Client) ListSwimmingLaneGroup(request *ListSwimmingLaneGroupRequest) (_result *ListSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSwimmingLaneGroupResponse{}
	_body, _err := client.ListSwimmingLaneGroupWithOptions(request, headers, runtime)
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
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/tag/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom namespaces.
//
// @param request - ListUserDefineRegionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDefineRegionResponse
func (client *Client) ListUserDefineRegionWithOptions(request *ListUserDefineRegionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListUserDefineRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DebugEnable) {
		query["DebugEnable"] = request.DebugEnable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDefineRegion"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/user_region_defs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDefineRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom namespaces.
//
// @param request - ListUserDefineRegionRequest
//
// @return ListUserDefineRegionResponse
func (client *Client) ListUserDefineRegion(request *ListUserDefineRegionRequest) (_result *ListUserDefineRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserDefineRegionResponse{}
	_body, _err := client.ListUserDefineRegionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The HTTP status code returned.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVpcResponse
func (client *Client) ListVpcWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVpcResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVpc"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/vpc_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The HTTP status code returned.
//
// @return ListVpcResponse
func (client *Client) ListVpc() (_result *ListVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVpcResponse{}
	_body, _err := client.ListVpcWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Migrates an elastic compute unit (ECU) to the default cluster in a specified namespace.
//
// Description:
//
// ## Limits
//
// We recommend that you do not call this operation. Instead, we recommend that you call the TransformClusterMember operation. For more information, see [TransformClusterMember](https://help.aliyun.com/document_detail/71514.html).
//
// When you call this operation to import an Elastic Compute Service (ECS) instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all data of the ECS instance is deleted. You must set a logon password for the ECS instance. Make sure that no important data exists on or data has been backed up for the ECS instance that you want to import.
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources and microservices in Enterprise Distributed Application Service (EDAS). The resources include clusters, ECS instances, and applications. You can use a default or custom namespace. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources or microservices.
//
//   - **ECU**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - MigrateEcuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateEcuResponse
func (client *Client) MigrateEcuWithOptions(request *MigrateEcuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *MigrateEcuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateEcu"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/migrate_ecu"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateEcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates an elastic compute unit (ECU) to the default cluster in a specified namespace.
//
// Description:
//
// ## Limits
//
// We recommend that you do not call this operation. Instead, we recommend that you call the TransformClusterMember operation. For more information, see [TransformClusterMember](https://help.aliyun.com/document_detail/71514.html).
//
// When you call this operation to import an Elastic Compute Service (ECS) instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all data of the ECS instance is deleted. You must set a logon password for the ECS instance. Make sure that no important data exists on or data has been backed up for the ECS instance that you want to import.
//
// ## Terms
//
//   - **Namespace**: the logical concept that is used to isolate resources and microservices in Enterprise Distributed Application Service (EDAS). The resources include clusters, ECS instances, and applications. You can use a default or custom namespace. Each region has a default namespace and supports multiple custom namespaces. By default, only the default namespace is available. You do not need to create a custom namespace if you do not want to isolate resources or microservices.
//
//   - **ECU**: After an ECS instance is imported to a cluster, the instance becomes an ECU.
//
//   - **Elastic compute container (ECC)**: After you deploy an application to an ECU in a cluster, the ECU becomes an ECC.
//
// @param request - MigrateEcuRequest
//
// @return MigrateEcuResponse
func (client *Client) MigrateEcu(request *MigrateEcuRequest) (_result *MigrateEcuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MigrateEcuResponse{}
	_body, _err := client.MigrateEcuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the scaling rule for an application.
//
// @param request - ModifyScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingRuleResponse
func (client *Client) ModifyScalingRuleWithOptions(request *ModifyScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptEULA) {
		query["AcceptEULA"] = request.AcceptEULA
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InCondition) {
		query["InCondition"] = request.InCondition
	}

	if !dara.IsNil(request.InCpu) {
		query["InCpu"] = request.InCpu
	}

	if !dara.IsNil(request.InDuration) {
		query["InDuration"] = request.InDuration
	}

	if !dara.IsNil(request.InEnable) {
		query["InEnable"] = request.InEnable
	}

	if !dara.IsNil(request.InInstanceNum) {
		query["InInstanceNum"] = request.InInstanceNum
	}

	if !dara.IsNil(request.InLoad) {
		query["InLoad"] = request.InLoad
	}

	if !dara.IsNil(request.InRT) {
		query["InRT"] = request.InRT
	}

	if !dara.IsNil(request.InStep) {
		query["InStep"] = request.InStep
	}

	if !dara.IsNil(request.KeyPairName) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !dara.IsNil(request.MultiAzPolicy) {
		query["MultiAzPolicy"] = request.MultiAzPolicy
	}

	if !dara.IsNil(request.OutCPU) {
		query["OutCPU"] = request.OutCPU
	}

	if !dara.IsNil(request.OutCondition) {
		query["OutCondition"] = request.OutCondition
	}

	if !dara.IsNil(request.OutDuration) {
		query["OutDuration"] = request.OutDuration
	}

	if !dara.IsNil(request.OutEnable) {
		query["OutEnable"] = request.OutEnable
	}

	if !dara.IsNil(request.OutInstanceNum) {
		query["OutInstanceNum"] = request.OutInstanceNum
	}

	if !dara.IsNil(request.OutLoad) {
		query["OutLoad"] = request.OutLoad
	}

	if !dara.IsNil(request.OutRT) {
		query["OutRT"] = request.OutRT
	}

	if !dara.IsNil(request.OutStep) {
		query["OutStep"] = request.OutStep
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ResourceFrom) {
		query["ResourceFrom"] = request.ResourceFrom
	}

	if !dara.IsNil(request.ScalingPolicy) {
		query["ScalingPolicy"] = request.ScalingPolicy
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateInstanceId) {
		query["TemplateInstanceId"] = request.TemplateInstanceId
	}

	if !dara.IsNil(request.TemplateInstanceName) {
		query["TemplateInstanceName"] = request.TemplateInstanceName
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/scaling_rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the scaling rule for an application.
//
// @param request - ModifyScalingRuleRequest
//
// @return ModifyScalingRuleResponse
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (_result *ModifyScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.ModifyScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an application.
//
// @param request - QueryApplicationStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApplicationStatusResponse
func (client *Client) QueryApplicationStatusWithOptions(request *QueryApplicationStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryApplicationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryApplicationStatus"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/app_status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryApplicationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an application.
//
// @param request - QueryApplicationStatusRequest
//
// @return QueryApplicationStatusResponse
func (client *Client) QueryApplicationStatus(request *QueryApplicationStatusRequest) (_result *QueryApplicationStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryApplicationStatusResponse{}
	_body, _err := client.QueryApplicationStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about an elastic compute container (ECC). This operation is applicable to Container Service for Kubernetes (ACK) clusters.
//
// @param request - QueryEccInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEccInfoResponse
func (client *Client) QueryEccInfoWithOptions(request *QueryEccInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryEccInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EccId) {
		query["EccId"] = request.EccId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEccInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/ecc"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEccInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about an elastic compute container (ECC). This operation is applicable to Container Service for Kubernetes (ACK) clusters.
//
// @param request - QueryEccInfoRequest
//
// @return QueryEccInfoResponse
func (client *Client) QueryEccInfo(request *QueryEccInfoRequest) (_result *QueryEccInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryEccInfoResponse{}
	_body, _err := client.QueryEccInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the elastic compute units (ECUs) that can be migrated.
//
// @param request - QueryMigrateEcuListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMigrateEcuListResponse
func (client *Client) QueryMigrateEcuListWithOptions(request *QueryMigrateEcuListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryMigrateEcuListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMigrateEcuList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/migrate_ecu_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMigrateEcuListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the elastic compute units (ECUs) that can be migrated.
//
// @param request - QueryMigrateEcuListRequest
//
// @return QueryMigrateEcuListResponse
func (client *Client) QueryMigrateEcuList(request *QueryMigrateEcuListRequest) (_result *QueryMigrateEcuListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMigrateEcuListResponse{}
	_body, _err := client.QueryMigrateEcuListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the namespaces to which an instance can be migrated.
//
// @param request - QueryMigrateRegionListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMigrateRegionListResponse
func (client *Client) QueryMigrateRegionListWithOptions(request *QueryMigrateRegionListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryMigrateRegionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicalRegionId) {
		query["LogicalRegionId"] = request.LogicalRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMigrateRegionList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/migrate_region_select"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMigrateRegionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the namespaces to which an instance can be migrated.
//
// @param request - QueryMigrateRegionListRequest
//
// @return QueryMigrateRegionListResponse
func (client *Client) QueryMigrateRegionList(request *QueryMigrateRegionListRequest) (_result *QueryMigrateRegionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMigrateRegionListResponse{}
	_body, _err := client.QueryMigrateRegionListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of different regions that are supported by Enterprise Distributed Application Service (EDAS).
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRegionConfigResponse
func (client *Client) QueryRegionConfigWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryRegionConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRegionConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/region_config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRegionConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of different regions that are supported by Enterprise Distributed Application Service (EDAS).
//
// @return QueryRegionConfigResponse
func (client *Client) QueryRegionConfig() (_result *QueryRegionConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRegionConfigResponse{}
	_body, _err := client.QueryRegionConfigWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration details of Log Service for an application.
//
// @param request - QuerySlsLogStoreListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySlsLogStoreListResponse
func (client *Client) QuerySlsLogStoreListWithOptions(request *QuerySlsLogStoreListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySlsLogStoreListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySlsLogStoreList"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/sls/query_sls_log_store_list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySlsLogStoreListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration details of Log Service for an application.
//
// @param request - QuerySlsLogStoreListRequest
//
// @return QuerySlsLogStoreListResponse
func (client *Client) QuerySlsLogStoreList(request *QuerySlsLogStoreListRequest) (_result *QuerySlsLogStoreListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySlsLogStoreListResponse{}
	_body, _err := client.QuerySlsLogStoreListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets an application.
//
// @param request - ResetApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetApplicationResponse
func (client *Client) ResetApplicationWithOptions(request *ResetApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_reset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets an application.
//
// @param request - ResetApplicationRequest
//
// @return ResetApplicationResponse
func (client *Client) ResetApplication(request *ResetApplicationRequest) (_result *ResetApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetApplicationResponse{}
	_body, _err := client.ResetApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an application. This operation is applicable to applications that are deployed in Elastic Compute Service (ECS) clusters.
//
// @param request - RestartApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartApplicationResponse
func (client *Client) RestartApplicationWithOptions(request *RestartApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an application. This operation is applicable to applications that are deployed in Elastic Compute Service (ECS) clusters.
//
// @param request - RestartApplicationRequest
//
// @return RestartApplicationResponse
func (client *Client) RestartApplication(request *RestartApplicationRequest) (_result *RestartApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartApplicationResponse{}
	_body, _err := client.RestartApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an application that is deployed in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - RestartK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartK8sApplicationResponse
func (client *Client) RestartK8sApplicationWithOptions(request *RestartK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/restart_k8s_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an application that is deployed in a Container Service for Kubernetes (ACK) cluster or a serverless Kubernetes cluster.
//
// @param request - RestartK8sApplicationRequest
//
// @return RestartK8sApplicationResponse
func (client *Client) RestartK8sApplication(request *RestartK8sApplicationRequest) (_result *RestartK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartK8sApplicationResponse{}
	_body, _err := client.RestartK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retries a failed process.
//
// @param request - RetryChangeOrderTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryChangeOrderTaskResponse
func (client *Client) RetryChangeOrderTaskWithOptions(request *RetryChangeOrderTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryChangeOrderTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RetryStatus) {
		query["RetryStatus"] = request.RetryStatus
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryChangeOrderTask"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/task_retry"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryChangeOrderTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retries a failed process.
//
// @param request - RetryChangeOrderTaskRequest
//
// @return RetryChangeOrderTaskResponse
func (client *Client) RetryChangeOrderTask(request *RetryChangeOrderTaskRequest) (_result *RetryChangeOrderTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetryChangeOrderTaskResponse{}
	_body, _err := client.RetryChangeOrderTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back an application.
//
// @param request - RollbackApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackApplicationResponse
func (client *Client) RollbackApplicationWithOptions(request *RollbackApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Batch) {
		query["Batch"] = request.Batch
	}

	if !dara.IsNil(request.BatchWaitTime) {
		query["BatchWaitTime"] = request.BatchWaitTime
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HistoryVersion) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_rollback"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back an application.
//
// @param request - RollbackApplicationRequest
//
// @return RollbackApplicationResponse
func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (_result *RollbackApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollbackApplicationResponse{}
	_body, _err := client.RollbackApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates an application change process and rolls back the application. This operation is applicable to applications that are deployed in Elastic Compute Service (ECS) clusters.
//
// @param request - RollbackChangeOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackChangeOrderResponse
func (client *Client) RollbackChangeOrderWithOptions(request *RollbackChangeOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackChangeOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeOrderId) {
		query["ChangeOrderId"] = request.ChangeOrderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackChangeOrder"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/changeorder/rollback"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackChangeOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates an application change process and rolls back the application. This operation is applicable to applications that are deployed in Elastic Compute Service (ECS) clusters.
//
// @param request - RollbackChangeOrderRequest
//
// @return RollbackChangeOrderResponse
func (client *Client) RollbackChangeOrder(request *RollbackChangeOrderRequest) (_result *RollbackChangeOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollbackChangeOrderResponse{}
	_body, _err := client.RollbackChangeOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales in an application.
//
// @param request - ScaleInApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleInApplicationResponse
func (client *Client) ScaleInApplicationWithOptions(request *ScaleInApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleInApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	if !dara.IsNil(request.ForceStatus) {
		query["ForceStatus"] = request.ForceStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleInApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_scale_in"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleInApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales in an application.
//
// @param request - ScaleInApplicationRequest
//
// @return ScaleInApplicationResponse
func (client *Client) ScaleInApplication(request *ScaleInApplicationRequest) (_result *ScaleInApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleInApplicationResponse{}
	_body, _err := client.ScaleInApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out or in an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - ScaleK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleK8sApplicationResponse
func (client *Client) ScaleK8sApplicationWithOptions(request *ScaleK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_apps"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out or in an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - ScaleK8sApplicationRequest
//
// @return ScaleK8sApplicationResponse
func (client *Client) ScaleK8sApplication(request *ScaleK8sApplicationRequest) (_result *ScaleK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleK8sApplicationResponse{}
	_body, _err := client.ScaleK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out an application.
//
// @param request - ScaleOutApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleOutApplicationResponse
func (client *Client) ScaleOutApplicationWithOptions(request *ScaleOutApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleOutApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DeployGroup) {
		query["DeployGroup"] = request.DeployGroup
	}

	if !dara.IsNil(request.EcuInfo) {
		query["EcuInfo"] = request.EcuInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleOutApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_scale_out"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleOutApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out an application.
//
// @param request - ScaleOutApplicationRequest
//
// @return ScaleOutApplicationResponse
func (client *Client) ScaleOutApplication(request *ScaleOutApplicationRequest) (_result *ScaleOutApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleOutApplicationResponse{}
	_body, _err := client.ScaleOutApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases Elastic Compute Service (ECS) instances in the Enterprise Distributed Application Service (EDAS) console and adds the purchased ECS instances to the specified instance group of an application.
//
// Description:
//
// ## Limits
//
// Assume that the auto scaling feature is configured and enabled for an application. When an auto scale-in is triggered for the application, the ECS instances that are purchased by calling this operation are removed first.
//
// @param request - ScaleoutApplicationWithNewInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleoutApplicationWithNewInstancesResponse
func (client *Client) ScaleoutApplicationWithNewInstancesWithOptions(request *ScaleoutApplicationWithNewInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ScaleoutApplicationWithNewInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.InstanceChargePeriod) {
		query["InstanceChargePeriod"] = request.InstanceChargePeriod
	}

	if !dara.IsNil(request.InstanceChargePeriodUnit) {
		query["InstanceChargePeriodUnit"] = request.InstanceChargePeriodUnit
	}

	if !dara.IsNil(request.InstanceChargeType) {
		query["InstanceChargeType"] = request.InstanceChargeType
	}

	if !dara.IsNil(request.ScalingNum) {
		query["ScalingNum"] = request.ScalingNum
	}

	if !dara.IsNil(request.ScalingPolicy) {
		query["ScalingPolicy"] = request.ScalingPolicy
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateInstanceId) {
		query["TemplateInstanceId"] = request.TemplateInstanceId
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScaleoutApplicationWithNewInstances"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/scaling/scale_out"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ScaleoutApplicationWithNewInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases Elastic Compute Service (ECS) instances in the Enterprise Distributed Application Service (EDAS) console and adds the purchased ECS instances to the specified instance group of an application.
//
// Description:
//
// ## Limits
//
// Assume that the auto scaling feature is configured and enabled for an application. When an auto scale-in is triggered for the application, the ECS instances that are purchased by calling this operation are removed first.
//
// @param request - ScaleoutApplicationWithNewInstancesRequest
//
// @return ScaleoutApplicationWithNewInstancesResponse
func (client *Client) ScaleoutApplicationWithNewInstances(request *ScaleoutApplicationWithNewInstancesRequest) (_result *ScaleoutApplicationWithNewInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleoutApplicationWithNewInstancesResponse{}
	_body, _err := client.ScaleoutApplicationWithNewInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an application.
//
// @param request - StartApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartApplicationResponse
func (client *Client) StartApplicationWithOptions(request *StartApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an application.
//
// @param request - StartApplicationRequest
//
// @return StartApplicationResponse
func (client *Client) StartApplication(request *StartApplicationRequest) (_result *StartApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartApplicationResponse{}
	_body, _err := client.StartApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts precheck for Kubernetes application changes.
//
// @param request - StartK8sAppPrecheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartK8sAppPrecheckResponse
func (client *Client) StartK8sAppPrecheckWithOptions(request *StartK8sAppPrecheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartK8sAppPrecheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentIds) {
		query["ComponentIds"] = request.ComponentIds
	}

	if !dara.IsNil(request.ConfigMountDescs) {
		query["ConfigMountDescs"] = request.ConfigMountDescs
	}

	if !dara.IsNil(request.EmptyDirs) {
		query["EmptyDirs"] = request.EmptyDirs
	}

	if !dara.IsNil(request.EnvFroms) {
		query["EnvFroms"] = request.EnvFroms
	}

	if !dara.IsNil(request.Envs) {
		query["Envs"] = request.Envs
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.JavaStartUpConfig) {
		query["JavaStartUpConfig"] = request.JavaStartUpConfig
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.LimitEphemeralStorage) {
		query["LimitEphemeralStorage"] = request.LimitEphemeralStorage
	}

	if !dara.IsNil(request.LimitMem) {
		query["LimitMem"] = request.LimitMem
	}

	if !dara.IsNil(request.LimitmCpu) {
		query["LimitmCpu"] = request.LimitmCpu
	}

	if !dara.IsNil(request.LocalVolume) {
		query["LocalVolume"] = request.LocalVolume
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PackageUrl) {
		query["PackageUrl"] = request.PackageUrl
	}

	if !dara.IsNil(request.PvcMountDescs) {
		query["PvcMountDescs"] = request.PvcMountDescs
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.RequestsEphemeralStorage) {
		query["RequestsEphemeralStorage"] = request.RequestsEphemeralStorage
	}

	if !dara.IsNil(request.RequestsMem) {
		query["RequestsMem"] = request.RequestsMem
	}

	if !dara.IsNil(request.RequestsmCpu) {
		query["RequestsmCpu"] = request.RequestsmCpu
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartK8sAppPrecheck"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/app_precheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartK8sAppPrecheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts precheck for Kubernetes application changes.
//
// @param request - StartK8sAppPrecheckRequest
//
// @return StartK8sAppPrecheckResponse
func (client *Client) StartK8sAppPrecheck(request *StartK8sAppPrecheckRequest) (_result *StartK8sAppPrecheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartK8sAppPrecheckResponse{}
	_body, _err := client.StartK8sAppPrecheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an application in a Container Service for Kubernetes (ACK) cluster or Serverless Kubernetes cluster.
//
// @param request - StartK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartK8sApplicationResponse
func (client *Client) StartK8sApplicationWithOptions(request *StartK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/start_k8s_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an application in a Container Service for Kubernetes (ACK) cluster or Serverless Kubernetes cluster.
//
// @param request - StartK8sApplicationRequest
//
// @return StartK8sApplicationResponse
func (client *Client) StartK8sApplication(request *StartK8sApplicationRequest) (_result *StartK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartK8sApplicationResponse{}
	_body, _err := client.StartK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an application.
//
// @param request - StopApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopApplicationResponse
func (client *Client) StopApplicationWithOptions(request *StopApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EccInfo) {
		query["EccInfo"] = request.EccInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an application.
//
// @param request - StopApplicationRequest
//
// @return StopApplicationResponse
func (client *Client) StopApplication(request *StopApplicationRequest) (_result *StopApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopApplicationResponse{}
	_body, _err := client.StopApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an application in a Container Service for Kubernetes (ACK) cluster or a Serverless Kubernetes cluster.
//
// @param request - StopK8sApplicationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopK8sApplicationResponse
func (client *Client) StopK8sApplicationWithOptions(request *StopK8sApplicationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopK8sApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopK8sApplication"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/stop_k8s_app"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopK8sApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an application in a Container Service for Kubernetes (ACK) cluster or a Serverless Kubernetes cluster.
//
// @param request - StopK8sApplicationRequest
//
// @return StopK8sApplicationResponse
func (client *Client) StopK8sApplication(request *StopK8sApplicationRequest) (_result *StopK8sApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopK8sApplicationResponse{}
	_body, _err := client.StopK8sApplicationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the advanced application monitoring feature or configures this feature for an application that is deployed in an Elastic Compute Service (ECS) or Kubernetes cluster.
//
// Description:
//
// To call the SwitchAdvancedMonitoring operation, you must make sure that the version of Enterprise Distributed Application Service (EDAS) SDK for Java or Python is 3.15.2 or later.
//
// @param request - SwitchAdvancedMonitoringRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchAdvancedMonitoringResponse
func (client *Client) SwitchAdvancedMonitoringWithOptions(request *SwitchAdvancedMonitoringRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SwitchAdvancedMonitoringResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EnableAdvancedMonitoring) {
		query["EnableAdvancedMonitoring"] = request.EnableAdvancedMonitoring
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchAdvancedMonitoring"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/monitor/advancedMonitorInfo"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchAdvancedMonitoringResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the advanced application monitoring feature or configures this feature for an application that is deployed in an Elastic Compute Service (ECS) or Kubernetes cluster.
//
// Description:
//
// To call the SwitchAdvancedMonitoring operation, you must make sure that the version of Enterprise Distributed Application Service (EDAS) SDK for Java or Python is 3.15.2 or later.
//
// @param request - SwitchAdvancedMonitoringRequest
//
// @return SwitchAdvancedMonitoringResponse
func (client *Client) SwitchAdvancedMonitoring(request *SwitchAdvancedMonitoringRequest) (_result *SwitchAdvancedMonitoringResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SwitchAdvancedMonitoringResponse{}
	_body, _err := client.SwitchAdvancedMonitoringWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the basic Alibaba Cloud resources that belong to your account to Enterprise Distributed Application Service (EDAS). This operation is applicable to Elastic Compute Service (ECS) clusters.
//
// Description:
//
// If you call this operation to synchronize ECS resource information, all instance data is synchronized from ECS. If you have more than 100 ECS instances, we recommend that you do not frequently call this operation.
//
// @param request - SynchronizeResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeResourceResponse
func (client *Client) SynchronizeResourceWithOptions(request *SynchronizeResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SynchronizeResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SynchronizeResource"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/pop_sync_resource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SynchronizeResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the basic Alibaba Cloud resources that belong to your account to Enterprise Distributed Application Service (EDAS). This operation is applicable to Elastic Compute Service (ECS) clusters.
//
// Description:
//
// If you call this operation to synchronize ECS resource information, all instance data is synchronized from ECS. If you have more than 100 ECS instances, we recommend that you do not frequently call this operation.
//
// @param request - SynchronizeResourceRequest
//
// @return SynchronizeResourceResponse
func (client *Client) SynchronizeResource(request *SynchronizeResourceRequest) (_result *SynchronizeResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SynchronizeResourceResponse{}
	_body, _err := client.SynchronizeResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates tags and adds the tags to resources at a time.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/tag/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Creates tags and adds the tags to resources at a time.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports or migrates one or more Elastic Compute Service (ECS) instances to a cluster.
//
// Description:
//
// ## Limits
//
// When you call this operation to import an ECS instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all data of the ECS instance is deleted. You must set a logon password for the ECS instance. Make sure that no important data exists on or data has been backed up for the ECS instance that you want to import.
//
// @param request - TransformClusterMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransformClusterMemberResponse
func (client *Client) TransformClusterMemberWithOptions(request *TransformClusterMemberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransformClusterMemberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceIds) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.TargetClusterId) {
		query["TargetClusterId"] = request.TargetClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransformClusterMember"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/resource/transform_cluster_member"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransformClusterMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports or migrates one or more Elastic Compute Service (ECS) instances to a cluster.
//
// Description:
//
// ## Limits
//
// When you call this operation to import an ECS instance, the operating system of the ECS instance is reinstalled. After the operating system is reinstalled, all data of the ECS instance is deleted. You must set a logon password for the ECS instance. Make sure that no important data exists on or data has been backed up for the ECS instance that you want to import.
//
// @param request - TransformClusterMemberRequest
//
// @return TransformClusterMemberResponse
func (client *Client) TransformClusterMember(request *TransformClusterMemberRequest) (_result *TransformClusterMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransformClusterMemberResponse{}
	_body, _err := client.TransformClusterMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a Server Load Balancer (SLB) instance from an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - UnbindK8sSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindK8sSlbResponse
func (client *Client) UnbindK8sSlbWithOptions(request *UnbindK8sSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindK8sSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.SlbName) {
		query["SlbName"] = request.SlbName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindK8sSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_slb_binding"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindK8sSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a Server Load Balancer (SLB) instance from an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - UnbindK8sSlbRequest
//
// @return UnbindK8sSlbResponse
func (client *Client) UnbindK8sSlb(request *UnbindK8sSlbRequest) (_result *UnbindK8sSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindK8sSlbResponse{}
	_body, _err := client.UnbindK8sSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds a Server Load Balancer (SLB) instance from an application.
//
// @param request - UnbindSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSlbResponse
func (client *Client) UnbindSlbWithOptions(request *UnbindSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnbindSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.DeleteListener) {
		query["DeleteListener"] = request.DeleteListener
	}

	if !dara.IsNil(request.SlbId) {
		query["SlbId"] = request.SlbId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/app/unbind_slb_json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds a Server Load Balancer (SLB) instance from an application.
//
// @param request - UnbindSlbRequest
//
// @return UnbindSlbResponse
func (client *Client) UnbindSlb(request *UnbindSlbRequest) (_result *UnbindSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindSlbResponse{}
	_body, _err := client.UnbindSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more tags from one or more resources.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteAll) {
		query["DeleteAll"] = request.DeleteAll
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/tag/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Removes one or more tags from one or more resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about an account.
//
// @param request - UpdateAccountInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountInfoResponse
func (client *Client) UpdateAccountInfoWithOptions(request *UpdateAccountInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAccountInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Telephone) {
		query["Telephone"] = request.Telephone
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccountInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/edit_account_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about an account.
//
// @param request - UpdateAccountInfoRequest
//
// @return UpdateAccountInfoResponse
func (client *Client) UpdateAccountInfo(request *UpdateAccountInfoRequest) (_result *UpdateAccountInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAccountInfoResponse{}
	_body, _err := client.UpdateAccountInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and owner of an application.
//
// @param request - UpdateApplicationBaseInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationBaseInfoResponse
func (client *Client) UpdateApplicationBaseInfoWithOptions(request *UpdateApplicationBaseInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationBaseInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/update_app_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and owner of an application.
//
// @param request - UpdateApplicationBaseInfoRequest
//
// @return UpdateApplicationBaseInfoResponse
func (client *Client) UpdateApplicationBaseInfo(request *UpdateApplicationBaseInfoRequest) (_result *UpdateApplicationBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationBaseInfoResponse{}
	_body, _err := client.UpdateApplicationBaseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an auto scaling policy for an application.
//
// @param request - UpdateApplicationScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationScalingRuleResponse
func (client *Client) UpdateApplicationScalingRuleWithOptions(request *UpdateApplicationScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApplicationScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ScalingBehaviour) {
		query["ScalingBehaviour"] = request.ScalingBehaviour
	}

	if !dara.IsNil(request.ScalingRuleEnable) {
		query["ScalingRuleEnable"] = request.ScalingRuleEnable
	}

	if !dara.IsNil(request.ScalingRuleMetric) {
		query["ScalingRuleMetric"] = request.ScalingRuleMetric
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.ScalingRuleTimer) {
		query["ScalingRuleTimer"] = request.ScalingRuleTimer
	}

	if !dara.IsNil(request.ScalingRuleTrigger) {
		query["ScalingRuleTrigger"] = request.ScalingRuleTrigger
	}

	if !dara.IsNil(request.ScalingRuleType) {
		query["ScalingRuleType"] = request.ScalingRuleType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationScalingRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v1/eam/scale/application_scaling_rule"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an auto scaling policy for an application.
//
// @param request - UpdateApplicationScalingRuleRequest
//
// @return UpdateApplicationScalingRuleResponse
func (client *Client) UpdateApplicationScalingRule(request *UpdateApplicationScalingRuleRequest) (_result *UpdateApplicationScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApplicationScalingRuleResponse{}
	_body, _err := client.UpdateApplicationScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a configuration template.
//
// @param request - UpdateConfigTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigTemplateResponse
func (client *Client) UpdateConfigTemplateWithOptions(request *UpdateConfigTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConfigTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Format) {
		body["Format"] = request.Format
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfigTemplate"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/config_template"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a configuration template.
//
// @param request - UpdateConfigTemplateRequest
//
// @return UpdateConfigTemplateResponse
func (client *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (_result *UpdateConfigTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigTemplateResponse{}
	_body, _err := client.UpdateConfigTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Enterprise Distributed Application Service (EDAS) Container version of a High-Speed Service Framework (HSF) application. EDAS Container includes Ali-Tomcat and Pandora.
//
// @param request - UpdateContainerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContainerResponse
func (client *Client) UpdateContainerWithOptions(request *UpdateContainerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContainerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BuildPackId) {
		query["BuildPackId"] = request.BuildPackId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContainer"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/changeorder/co_update_container"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContainerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Enterprise Distributed Application Service (EDAS) Container version of a High-Speed Service Framework (HSF) application. EDAS Container includes Ali-Tomcat and Pandora.
//
// @param request - UpdateContainerRequest
//
// @return UpdateContainerResponse
func (client *Client) UpdateContainer(request *UpdateContainerRequest) (_result *UpdateContainerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateContainerResponse{}
	_body, _err := client.UpdateContainerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the Tomcat container for an application or application instance group in an Elastic Compute Service (ECS) cluster.
//
// @param request - UpdateContainerConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContainerConfigurationResponse
func (client *Client) UpdateContainerConfigurationWithOptions(request *UpdateContainerConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContainerConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ContextPath) {
		query["ContextPath"] = request.ContextPath
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HttpPort) {
		query["HttpPort"] = request.HttpPort
	}

	if !dara.IsNil(request.MaxThreads) {
		query["MaxThreads"] = request.MaxThreads
	}

	if !dara.IsNil(request.URIEncoding) {
		query["URIEncoding"] = request.URIEncoding
	}

	if !dara.IsNil(request.UseBodyEncoding) {
		query["UseBodyEncoding"] = request.UseBodyEncoding
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContainerConfiguration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/container_config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContainerConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the Tomcat container for an application or application instance group in an Elastic Compute Service (ECS) cluster.
//
// @param request - UpdateContainerConfigurationRequest
//
// @return UpdateContainerConfigurationResponse
func (client *Client) UpdateContainerConfiguration(request *UpdateContainerConfigurationRequest) (_result *UpdateContainerConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateContainerConfigurationResponse{}
	_body, _err := client.UpdateContainerConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the health check URL for an application.
//
// @param request - UpdateHealthCheckUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHealthCheckUrlResponse
func (client *Client) UpdateHealthCheckUrlWithOptions(request *UpdateHealthCheckUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHealthCheckUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.HcURL) {
		query["hcURL"] = request.HcURL
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHealthCheckUrl"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/modify_hc_url"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHealthCheckUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the health check URL for an application.
//
// @param request - UpdateHealthCheckUrlRequest
//
// @return UpdateHealthCheckUrlResponse
func (client *Client) UpdateHealthCheckUrl(request *UpdateHealthCheckUrlRequest) (_result *UpdateHealthCheckUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHealthCheckUrlResponse{}
	_body, _err := client.UpdateHealthCheckUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Mounts a script to an application or application instance group.
//
// @param request - UpdateHookConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHookConfigurationResponse
func (client *Client) UpdateHookConfigurationWithOptions(request *UpdateHookConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHookConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Hooks) {
		query["Hooks"] = request.Hooks
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHookConfiguration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/app/config_app_hook_json"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHookConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Mounts a script to an application or application instance group.
//
// @param request - UpdateHookConfigurationRequest
//
// @return UpdateHookConfigurationResponse
func (client *Client) UpdateHookConfiguration(request *UpdateHookConfigurationRequest) (_result *UpdateHookConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHookConfigurationResponse{}
	_body, _err := client.UpdateHookConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the Java virtual machine (JVM) parameters for an application or an instance group where the application is deployed.
//
// @param request - UpdateJvmConfigurationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJvmConfigurationResponse
func (client *Client) UpdateJvmConfigurationWithOptions(request *UpdateJvmConfigurationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJvmConfigurationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.MaxHeapSize) {
		query["MaxHeapSize"] = request.MaxHeapSize
	}

	if !dara.IsNil(request.MaxPermSize) {
		query["MaxPermSize"] = request.MaxPermSize
	}

	if !dara.IsNil(request.MinHeapSize) {
		query["MinHeapSize"] = request.MinHeapSize
	}

	if !dara.IsNil(request.Options) {
		query["Options"] = request.Options
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJvmConfiguration"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/app/app_jvm_config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJvmConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the Java virtual machine (JVM) parameters for an application or an instance group where the application is deployed.
//
// @param request - UpdateJvmConfigurationRequest
//
// @return UpdateJvmConfigurationResponse
func (client *Client) UpdateJvmConfiguration(request *UpdateJvmConfigurationRequest) (_result *UpdateJvmConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJvmConfigurationResponse{}
	_body, _err := client.UpdateJvmConfigurationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies basic information about an application that is deployed in a Kubernetes cluster.
//
// @param request - UpdateK8sApplicationBaseInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sApplicationBaseInfoResponse
func (client *Client) UpdateK8sApplicationBaseInfoWithOptions(request *UpdateK8sApplicationBaseInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sApplicationBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Owner) {
		query["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sApplicationBaseInfo"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/update_app_basic_info"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sApplicationBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies basic information about an application that is deployed in a Kubernetes cluster.
//
// @param request - UpdateK8sApplicationBaseInfoRequest
//
// @return UpdateK8sApplicationBaseInfoResponse
func (client *Client) UpdateK8sApplicationBaseInfo(request *UpdateK8sApplicationBaseInfoRequest) (_result *UpdateK8sApplicationBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sApplicationBaseInfoResponse{}
	_body, _err := client.UpdateK8sApplicationBaseInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of an application in a Container Service for Kubernetes (ACK) or Serverless Kubernetes cluster.
//
// @param request - UpdateK8sApplicationConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sApplicationConfigResponse
func (client *Client) UpdateK8sApplicationConfigWithOptions(request *UpdateK8sApplicationConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sApplicationConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.CpuLimit) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.CpuRequest) {
		query["CpuRequest"] = request.CpuRequest
	}

	if !dara.IsNil(request.EphemeralStorageLimit) {
		query["EphemeralStorageLimit"] = request.EphemeralStorageLimit
	}

	if !dara.IsNil(request.EphemeralStorageRequest) {
		query["EphemeralStorageRequest"] = request.EphemeralStorageRequest
	}

	if !dara.IsNil(request.McpuLimit) {
		query["McpuLimit"] = request.McpuLimit
	}

	if !dara.IsNil(request.McpuRequest) {
		query["McpuRequest"] = request.McpuRequest
	}

	if !dara.IsNil(request.MemoryLimit) {
		query["MemoryLimit"] = request.MemoryLimit
	}

	if !dara.IsNil(request.MemoryRequest) {
		query["MemoryRequest"] = request.MemoryRequest
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sApplicationConfig"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_app_configuration"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sApplicationConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of an application in a Container Service for Kubernetes (ACK) or Serverless Kubernetes cluster.
//
// @param request - UpdateK8sApplicationConfigRequest
//
// @return UpdateK8sApplicationConfigResponse
func (client *Client) UpdateK8sApplicationConfig(request *UpdateK8sApplicationConfigRequest) (_result *UpdateK8sApplicationConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sApplicationConfigResponse{}
	_body, _err := client.UpdateK8sApplicationConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Kubernetes ConfigMap.
//
// @param request - UpdateK8sConfigMapRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sConfigMapResponse
func (client *Client) UpdateK8sConfigMapWithOptions(request *UpdateK8sConfigMapRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sConfigMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sConfigMap"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_config_map"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sConfigMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Kubernetes ConfigMap.
//
// @param request - UpdateK8sConfigMapRequest
//
// @return UpdateK8sConfigMapResponse
func (client *Client) UpdateK8sConfigMap(request *UpdateK8sConfigMapRequest) (_result *UpdateK8sConfigMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sConfigMapResponse{}
	_body, _err := client.UpdateK8sConfigMapWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an ingress.
//
// @param request - UpdateK8sIngressRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sIngressRuleResponse
func (client *Client) UpdateK8sIngressRuleWithOptions(request *UpdateK8sIngressRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sIngressRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Annotations) {
		query["Annotations"] = request.Annotations
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IngressConf) {
		query["IngressConf"] = request.IngressConf
	}

	if !dara.IsNil(request.Labels) {
		query["Labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sIngressRule"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_ingress"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sIngressRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an ingress.
//
// @param request - UpdateK8sIngressRuleRequest
//
// @return UpdateK8sIngressRuleResponse
func (client *Client) UpdateK8sIngressRule(request *UpdateK8sIngressRuleRequest) (_result *UpdateK8sIngressRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sIngressRuleResponse{}
	_body, _err := client.UpdateK8sIngressRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specified resource in a Kubernetes cluster.
//
// Description:
//
// > You can update only Deployments.
//
// @param request - UpdateK8sResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sResourceResponse
func (client *Client) UpdateK8sResourceWithOptions(request *UpdateK8sResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.ResourceContent) {
		body["ResourceContent"] = request.ResourceContent
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sResource"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/oam/update_k8s_resource_config"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specified resource in a Kubernetes cluster.
//
// Description:
//
// > You can update only Deployments.
//
// @param request - UpdateK8sResourceRequest
//
// @return UpdateK8sResourceResponse
func (client *Client) UpdateK8sResource(request *UpdateK8sResourceRequest) (_result *UpdateK8sResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sResourceResponse{}
	_body, _err := client.UpdateK8sResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Kubernetes Secret.
//
// @param request - UpdateK8sSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sSecretResponse
func (client *Client) UpdateK8sSecretWithOptions(request *UpdateK8sSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Base64Encoded) {
		body["Base64Encoded"] = request.Base64Encoded
	}

	if !dara.IsNil(request.CertId) {
		body["CertId"] = request.CertId
	}

	if !dara.IsNil(request.CertRegionId) {
		body["CertRegionId"] = request.CertRegionId
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Namespace) {
		body["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sSecret"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_secret"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Kubernetes Secret.
//
// @param request - UpdateK8sSecretRequest
//
// @return UpdateK8sSecretResponse
func (client *Client) UpdateK8sSecret(request *UpdateK8sSecretRequest) (_result *UpdateK8sSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sSecretResponse{}
	_body, _err := client.UpdateK8sSecretWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an application service in a Kubernetes cluster.
//
// @param request - UpdateK8sServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sServiceResponse
func (client *Client) UpdateK8sServiceWithOptions(request *UpdateK8sServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ExternalTrafficPolicy) {
		query["ExternalTrafficPolicy"] = request.ExternalTrafficPolicy
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ServicePorts) {
		query["ServicePorts"] = request.ServicePorts
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sService"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_service"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application service in a Kubernetes cluster.
//
// @param request - UpdateK8sServiceRequest
//
// @return UpdateK8sServiceResponse
func (client *Client) UpdateK8sService(request *UpdateK8sServiceRequest) (_result *UpdateK8sServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sServiceResponse{}
	_body, _err := client.UpdateK8sServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Server Load Balancer (SLB) instance bound to an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - UpdateK8sSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateK8sSlbResponse
func (client *Client) UpdateK8sSlbWithOptions(request *UpdateK8sSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateK8sSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DisableForceOverride) {
		query["DisableForceOverride"] = request.DisableForceOverride
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Scheduler) {
		query["Scheduler"] = request.Scheduler
	}

	if !dara.IsNil(request.ServicePortInfos) {
		query["ServicePortInfos"] = request.ServicePortInfos
	}

	if !dara.IsNil(request.SlbName) {
		query["SlbName"] = request.SlbName
	}

	if !dara.IsNil(request.SlbProtocol) {
		query["SlbProtocol"] = request.SlbProtocol
	}

	if !dara.IsNil(request.Specification) {
		query["Specification"] = request.Specification
	}

	if !dara.IsNil(request.TargetPort) {
		query["TargetPort"] = request.TargetPort
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateK8sSlb"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/acs/k8s_slb_binding"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateK8sSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Server Load Balancer (SLB) instance bound to an application that is deployed in a Container Service for Kubernetes (ACK) cluster.
//
// @param request - UpdateK8sSlbRequest
//
// @return UpdateK8sSlbResponse
func (client *Client) UpdateK8sSlb(request *UpdateK8sSlbRequest) (_result *UpdateK8sSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateK8sSlbResponse{}
	_body, _err := client.UpdateK8sSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新本地设置
//
// @param request - UpdateLocalitySettingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocalitySettingResponse
func (client *Client) UpdateLocalitySettingWithOptions(request *UpdateLocalitySettingRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLocalitySettingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.NamespaceId) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.Threshold) {
		query["Threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocalitySetting"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/sp/applications/locality/setting"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocalitySettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新本地设置
//
// @param request - UpdateLocalitySettingRequest
//
// @return UpdateLocalitySettingResponse
func (client *Client) UpdateLocalitySetting(request *UpdateLocalitySettingRequest) (_result *UpdateLocalitySettingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateLocalitySettingResponse{}
	_body, _err := client.UpdateLocalitySettingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a role.
//
// @param request - UpdateRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRoleResponse
func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionData) {
		query["ActionData"] = request.ActionData
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRole"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/account/edit_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
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
// Modifies a role.
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
// Configures a Logstore for an application.
//
// @param request - UpdateSlsLogStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSlsLogStoreResponse
func (client *Client) UpdateSlsLogStoreWithOptions(request *UpdateSlsLogStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSlsLogStoreResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Configs) {
		body["Configs"] = request.Configs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSlsLogStore"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/k8s/sls/update_sls_log_store"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSlsLogStoreResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a Logstore for an application.
//
// @param request - UpdateSlsLogStoreRequest
//
// @return UpdateSlsLogStoreResponse
func (client *Client) UpdateSlsLogStore(request *UpdateSlsLogStoreRequest) (_result *UpdateSlsLogStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSlsLogStoreResponse{}
	_body, _err := client.UpdateSlsLogStoreWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新泳道
//
// @param request - UpdateSwimmingLaneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSwimmingLaneResponse
func (client *Client) UpdateSwimmingLaneWithOptions(request *UpdateSwimmingLaneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSwimmingLaneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInfos) {
		query["AppInfos"] = request.AppInfos
	}

	if !dara.IsNil(request.EnableRules) {
		query["EnableRules"] = request.EnableRules
	}

	if !dara.IsNil(request.EntryRules) {
		query["EntryRules"] = request.EntryRules
	}

	if !dara.IsNil(request.LaneId) {
		query["LaneId"] = request.LaneId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSwimmingLane"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lanes"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSwimmingLaneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新泳道
//
// @param request - UpdateSwimmingLaneRequest
//
// @return UpdateSwimmingLaneResponse
func (client *Client) UpdateSwimmingLane(request *UpdateSwimmingLaneRequest) (_result *UpdateSwimmingLaneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSwimmingLaneResponse{}
	_body, _err := client.UpdateSwimmingLaneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a lane group.
//
// @param request - UpdateSwimmingLaneGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSwimmingLaneGroupResponse
func (client *Client) UpdateSwimmingLaneGroupWithOptions(request *UpdateSwimmingLaneGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSwimmingLaneGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.EntryApp) {
		query["EntryApp"] = request.EntryApp
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSwimmingLaneGroup"),
		Version:     dara.String("2017-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/v5/trafficmgnt/swimming_lane_groups"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSwimmingLaneGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a lane group.
//
// @param request - UpdateSwimmingLaneGroupRequest
//
// @return UpdateSwimmingLaneGroupResponse
func (client *Client) UpdateSwimmingLaneGroup(request *UpdateSwimmingLaneGroupRequest) (_result *UpdateSwimmingLaneGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSwimmingLaneGroupResponse{}
	_body, _err := client.UpdateSwimmingLaneGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
