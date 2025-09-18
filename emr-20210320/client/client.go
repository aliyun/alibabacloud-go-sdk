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
		"cn-beijing":                  dara.String("emr.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("emr.aliyuncs.com"),
		"cn-shanghai":                 dara.String("emr.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("emr.aliyuncs.com"),
		"ap-southeast-1":              dara.String("emr.aliyuncs.com"),
		"us-west-1":                   dara.String("emr.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("emr.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("emr.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("emr.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("emr.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("emr.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("emr.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("emr.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("emr.aliyuncs.com"),
		"cn-edge-1":                   dara.String("emr.aliyuncs.com"),
		"cn-fujian":                   dara.String("emr.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("emr.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("emr.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("emr.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("emr.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("emr.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("emr.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("emr.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("emr.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("emr.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("emr.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("emr.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("emr.aliyuncs.com"),
		"cn-wuhan":                    dara.String("emr.aliyuncs.com"),
		"cn-yushanfang":               dara.String("emr.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("emr.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("emr.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("emr.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("emr.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("emr.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("emr.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("emr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a predefined API operation template. The template contains information about an API operation, including the basic structure, request method, URL path, request parameters, and response format.
//
// @param request - CreateApiTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiTemplateResponse
func (client *Client) CreateApiTemplateWithOptions(request *CreateApiTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateApiTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiTemplate"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a predefined API operation template. The template contains information about an API operation, including the basic structure, request method, URL path, request parameters, and response format.
//
// @param request - CreateApiTemplateRequest
//
// @return CreateApiTemplateResponse
func (client *Client) CreateApiTemplate(request *CreateApiTemplateRequest) (_result *CreateApiTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApiTemplateResponse{}
	_body, _err := client.CreateApiTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription cluster.
//
// @param request - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigs) {
		query["ApplicationConfigs"] = request.ApplicationConfigs
	}

	if !dara.IsNil(request.Applications) {
		query["Applications"] = request.Applications
	}

	if !dara.IsNil(request.BootstrapScripts) {
		query["BootstrapScripts"] = request.BootstrapScripts
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DeployMode) {
		query["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NodeAttributes) {
		query["NodeAttributes"] = request.NodeAttributes
	}

	if !dara.IsNil(request.NodeGroups) {
		query["NodeGroups"] = request.NodeGroups
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseVersion) {
		query["ReleaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityMode) {
		query["SecurityMode"] = request.SecurityMode
	}

	if !dara.IsNil(request.SubscriptionConfig) {
		query["SubscriptionConfig"] = request.SubscriptionConfig
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription cluster.
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a node group.
//
// Description:
//
// 创建节点组。
//
// @param request - CreateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroupWithOptions(request *CreateNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroup) {
		query["NodeGroup"] = request.NodeGroup
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeGroup"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a node group.
//
// Description:
//
// 创建节点组。
//
// @param request - CreateNodeGroupRequest
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroup(request *CreateNodeGroupRequest) (_result *CreateNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CreateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param request - CreateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScriptResponse
func (client *Client) CreateScriptWithOptions(request *CreateScriptRequest, runtime *dara.RuntimeOptions) (_result *CreateScriptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScriptType) {
		query["ScriptType"] = request.ScriptType
	}

	if !dara.IsNil(request.Scripts) {
		query["Scripts"] = request.Scripts
	}

	if !dara.IsNil(request.TimeoutSecs) {
		query["TimeoutSecs"] = request.TimeoutSecs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScript"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param request - CreateScriptRequest
//
// @return CreateScriptResponse
func (client *Client) CreateScript(request *CreateScriptRequest) (_result *CreateScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScriptResponse{}
	_body, _err := client.CreateScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple users at a time.
//
// Description:
//
// You can call this operation to create multiple users at a time.
//
// @param request - CreateUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsersResponse
func (client *Client) CreateUsersWithOptions(request *CreateUsersRequest, runtime *dara.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUsers"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple users at a time.
//
// Description:
//
// You can call this operation to create multiple users at a time.
//
// @param request - CreateUsersRequest
//
// @return CreateUsersResponse
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a scale-out operation on the target node group.
//
// @param request - DecreaseNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecreaseNodesResponse
func (client *Client) DecreaseNodesWithOptions(request *DecreaseNodesRequest, runtime *dara.RuntimeOptions) (_result *DecreaseNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchInterval) {
		query["BatchInterval"] = request.BatchInterval
	}

	if !dara.IsNil(request.BatchSize) {
		query["BatchSize"] = request.BatchSize
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DecreaseNodeCount) {
		query["DecreaseNodeCount"] = request.DecreaseNodeCount
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DecreaseNodes"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DecreaseNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a scale-out operation on the target node group.
//
// @param request - DecreaseNodesRequest
//
// @return DecreaseNodesResponse
func (client *Client) DecreaseNodes(request *DecreaseNodesRequest) (_result *DecreaseNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DecreaseNodesResponse{}
	_body, _err := client.DecreaseNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API operation template.
//
// @param request - DeleteApiTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiTemplateResponse
func (client *Client) DeleteApiTemplateWithOptions(request *DeleteApiTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteApiTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiTemplate"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API operation template.
//
// @param request - DeleteApiTemplateRequest
//
// @return DeleteApiTemplateResponse
func (client *Client) DeleteApiTemplate(request *DeleteApiTemplateRequest) (_result *DeleteApiTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApiTemplateResponse{}
	_body, _err := client.DeleteApiTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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

// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param request - DeleteScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScriptResponse
func (client *Client) DeleteScriptWithOptions(request *DeleteScriptRequest, runtime *dara.RuntimeOptions) (_result *DeleteScriptResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptType) {
		query["ScriptType"] = request.ScriptType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScript"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param request - DeleteScriptRequest
//
// @return DeleteScriptResponse
func (client *Client) DeleteScript(request *DeleteScriptRequest) (_result *DeleteScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScriptResponse{}
	_body, _err := client.DeleteScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple users at a time.
//
// Description:
//
// Deletes multiple users at a time.
//
// @param tmpReq - DeleteUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsersWithOptions(tmpReq *DeleteUsersRequest, runtime *dara.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserNames) {
		request.UserNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserNames, dara.String("UserNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.UserNamesShrink) {
		body["UserNames"] = request.UserNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUsers"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple users at a time.
//
// Description:
//
// Deletes multiple users at a time.
//
// @param request - DeleteUsersRequest
//
// @return DeleteUsersResponse
func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ExportApplicationConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportApplicationConfigsResponse
func (client *Client) ExportApplicationConfigsWithOptions(request *ExportApplicationConfigsRequest, runtime *dara.RuntimeOptions) (_result *ExportApplicationConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigFiles) {
		query["ApplicationConfigFiles"] = request.ApplicationConfigFiles
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ExportMode) {
		query["ExportMode"] = request.ExportMode
	}

	if !dara.IsNil(request.FileFormat) {
		query["FileFormat"] = request.FileFormat
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportApplicationConfigs"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportApplicationConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ExportApplicationConfigsRequest
//
// @return ExportApplicationConfigsResponse
func (client *Client) ExportApplicationConfigs(request *ExportApplicationConfigsRequest) (_result *ExportApplicationConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportApplicationConfigsResponse{}
	_body, _err := client.ExportApplicationConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration information about an API operation template.
//
// @param request - GetApiTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiTemplateResponse
func (client *Client) GetApiTemplateWithOptions(request *GetApiTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetApiTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiTemplate"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed configuration information about an API operation template.
//
// @param request - GetApiTemplateRequest
//
// @return GetApiTemplateResponse
func (client *Client) GetApiTemplate(request *GetApiTemplateRequest) (_result *GetApiTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApiTemplateResponse{}
	_body, _err := client.GetApiTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an application.
//
// Description:
//
// 查询应用详情。
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Retrieves the details of an application.
//
// Description:
//
// 查询应用详情。
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an auto scaling activity.
//
// @param request - GetAutoScalingActivityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingActivityResponse
func (client *Client) GetAutoScalingActivityWithOptions(request *GetAutoScalingActivityRequest, runtime *dara.RuntimeOptions) (_result *GetAutoScalingActivityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingActivityId) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoScalingActivity"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoScalingActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an auto scaling activity.
//
// @param request - GetAutoScalingActivityRequest
//
// @return GetAutoScalingActivityResponse
func (client *Client) GetAutoScalingActivity(request *GetAutoScalingActivityRequest) (_result *GetAutoScalingActivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoScalingActivityResponse{}
	_body, _err := client.GetAutoScalingActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries custom auto scaling rules.
//
// @param request - GetAutoScalingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingPolicyResponse
func (client *Client) GetAutoScalingPolicyWithOptions(request *GetAutoScalingPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetAutoScalingPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoScalingPolicy"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoScalingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries custom auto scaling rules.
//
// @param request - GetAutoScalingPolicyRequest
//
// @return GetAutoScalingPolicyResponse
func (client *Client) GetAutoScalingPolicy(request *GetAutoScalingPolicyRequest) (_result *GetAutoScalingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoScalingPolicyResponse{}
	_body, _err := client.GetAutoScalingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a cluster.
//
// @param request - GetClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterResponse
func (client *Client) GetClusterWithOptions(request *GetClusterRequest, runtime *dara.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Obtains the details of a cluster.
//
// @param request - GetClusterRequest
//
// @return GetClusterResponse
func (client *Client) GetCluster(request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains metadata of the E-MapReduce (EMR) cluster that you want to clone. This helps you call the CreateCluster API operation to quickly create an EMR cluster.
//
// @param request - GetClusterCloneMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClusterCloneMetaResponse
func (client *Client) GetClusterCloneMetaWithOptions(request *GetClusterCloneMetaRequest, runtime *dara.RuntimeOptions) (_result *GetClusterCloneMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClusterCloneMeta"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClusterCloneMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains metadata of the E-MapReduce (EMR) cluster that you want to clone. This helps you call the CreateCluster API operation to quickly create an EMR cluster.
//
// @param request - GetClusterCloneMetaRequest
//
// @return GetClusterCloneMetaResponse
func (client *Client) GetClusterCloneMeta(request *GetClusterCloneMetaRequest) (_result *GetClusterCloneMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClusterCloneMetaResponse{}
	_body, _err := client.GetClusterCloneMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains job analysis information on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get one doctor analysis app
//
// @param request - GetDoctorApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorApplicationResponse
func (client *Client) GetDoctorApplicationWithOptions(request *GetDoctorApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorApplication"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains job analysis information on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get one doctor analysis app
//
// @param request - GetDoctorApplicationRequest
//
// @return GetDoctorApplicationResponse
func (client *Client) GetDoctorApplication(request *GetDoctorApplicationRequest) (_result *GetDoctorApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorApplicationResponse{}
	_body, _err := client.GetDoctorApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about resource usage in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get one specific luster engine queue by <type, name>
//
// @param request - GetDoctorComputeSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorComputeSummaryResponse
func (client *Client) GetDoctorComputeSummaryWithOptions(request *GetDoctorComputeSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorComputeSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentInfo) {
		query["ComponentInfo"] = request.ComponentInfo
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorComputeSummary"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorComputeSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about resource usage in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get one specific luster engine queue by <type, name>
//
// @param request - GetDoctorComputeSummaryRequest
//
// @return GetDoctorComputeSummaryResponse
func (client *Client) GetDoctorComputeSummary(request *GetDoctorComputeSummaryRequest) (_result *GetDoctorComputeSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorComputeSummaryResponse{}
	_body, _err := client.GetDoctorComputeSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the metrics of an HBase cluster.
//
// Description:
//
// get Doctor HBaseCluster
//
// @param request - GetDoctorHBaseClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHBaseClusterResponse
func (client *Client) GetDoctorHBaseClusterWithOptions(request *GetDoctorHBaseClusterRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHBaseClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHBaseCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHBaseClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the metrics of an HBase cluster.
//
// Description:
//
// get Doctor HBaseCluster
//
// @param request - GetDoctorHBaseClusterRequest
//
// @return GetDoctorHBaseClusterResponse
func (client *Client) GetDoctorHBaseCluster(request *GetDoctorHBaseClusterRequest) (_result *GetDoctorHBaseClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHBaseClusterResponse{}
	_body, _err := client.GetDoctorHBaseClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get HBase Region information.
//
// Description:
//
// # List Doctor HBase Regions
//
// @param request - GetDoctorHBaseRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHBaseRegionResponse
func (client *Client) GetDoctorHBaseRegionWithOptions(request *GetDoctorHBaseRegionRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHBaseRegionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.HbaseRegionId) {
		query["HbaseRegionId"] = request.HbaseRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHBaseRegion"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHBaseRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get HBase Region information.
//
// Description:
//
// # List Doctor HBase Regions
//
// @param request - GetDoctorHBaseRegionRequest
//
// @return GetDoctorHBaseRegionResponse
func (client *Client) GetDoctorHBaseRegion(request *GetDoctorHBaseRegionRequest) (_result *GetDoctorHBaseRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHBaseRegionResponse{}
	_body, _err := client.GetDoctorHBaseRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about an HBase region server.
//
// Description:
//
// get Doctor HBaseRegionServer
//
// @param request - GetDoctorHBaseRegionServerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHBaseRegionServerResponse
func (client *Client) GetDoctorHBaseRegionServerWithOptions(request *GetDoctorHBaseRegionServerRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHBaseRegionServerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionServerHost) {
		query["RegionServerHost"] = request.RegionServerHost
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHBaseRegionServer"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHBaseRegionServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about an HBase region server.
//
// Description:
//
// get Doctor HBaseRegionServer
//
// @param request - GetDoctorHBaseRegionServerRequest
//
// @return GetDoctorHBaseRegionServerResponse
func (client *Client) GetDoctorHBaseRegionServer(request *GetDoctorHBaseRegionServerRequest) (_result *GetDoctorHBaseRegionServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHBaseRegionServerResponse{}
	_body, _err := client.GetDoctorHBaseRegionServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get HBase Table information.
//
// Description:
//
// get Doctor HBaseTable
//
// @param request - GetDoctorHBaseTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHBaseTableResponse
func (client *Client) GetDoctorHBaseTableWithOptions(request *GetDoctorHBaseTableRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHBaseTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHBaseTable"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHBaseTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get HBase Table information.
//
// Description:
//
// get Doctor HBaseTable
//
// @param request - GetDoctorHBaseTableRequest
//
// @return GetDoctorHBaseTableResponse
func (client *Client) GetDoctorHBaseTable(request *GetDoctorHBaseTableRequest) (_result *GetDoctorHBaseTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHBaseTableResponse{}
	_body, _err := client.GetDoctorHBaseTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of the Hadoop Distributed File System (HDFS) storage resources of a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor HBaseTableRegions
//
// @param request - GetDoctorHDFSClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHDFSClusterResponse
func (client *Client) GetDoctorHDFSClusterWithOptions(request *GetDoctorHDFSClusterRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHDFSClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHDFSCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHDFSClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of the Hadoop Distributed File System (HDFS) storage resources of a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor HBaseTableRegions
//
// @param request - GetDoctorHDFSClusterRequest
//
// @return GetDoctorHDFSClusterResponse
func (client *Client) GetDoctorHDFSCluster(request *GetDoctorHDFSClusterRequest) (_result *GetDoctorHDFSClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHDFSClusterResponse{}
	_body, _err := client.GetDoctorHDFSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a specific Hadoop Distributed File System (HDFS) directory of a cluster. The depth of the directory is not greater than five.
//
// Description:
//
// get Doctor HDFSNode
//
// @param request - GetDoctorHDFSDirectoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHDFSDirectoryResponse
func (client *Client) GetDoctorHDFSDirectoryWithOptions(request *GetDoctorHDFSDirectoryRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHDFSDirectoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.DirPath) {
		query["DirPath"] = request.DirPath
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHDFSDirectory"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHDFSDirectoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a specific Hadoop Distributed File System (HDFS) directory of a cluster. The depth of the directory is not greater than five.
//
// Description:
//
// get Doctor HDFSNode
//
// @param request - GetDoctorHDFSDirectoryRequest
//
// @return GetDoctorHDFSDirectoryResponse
func (client *Client) GetDoctorHDFSDirectory(request *GetDoctorHDFSDirectoryRequest) (_result *GetDoctorHDFSDirectoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHDFSDirectoryResponse{}
	_body, _err := client.GetDoctorHDFSDirectoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a Hive cluster.
//
// Description:
//
// list Doctor Hive Cluster
//
// @param request - GetDoctorHiveClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHiveClusterResponse
func (client *Client) GetDoctorHiveClusterWithOptions(request *GetDoctorHiveClusterRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHiveClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHiveCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHiveClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a Hive cluster.
//
// Description:
//
// list Doctor Hive Cluster
//
// @param request - GetDoctorHiveClusterRequest
//
// @return GetDoctorHiveClusterResponse
func (client *Client) GetDoctorHiveCluster(request *GetDoctorHiveClusterRequest) (_result *GetDoctorHiveClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHiveClusterResponse{}
	_body, _err := client.GetDoctorHiveClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a Hive database.
//
// Description:
//
// get Doctor Hive Database
//
// @param request - GetDoctorHiveDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHiveDatabaseResponse
func (client *Client) GetDoctorHiveDatabaseWithOptions(request *GetDoctorHiveDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHiveDatabaseResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHiveDatabase"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHiveDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a Hive database.
//
// Description:
//
// get Doctor Hive Database
//
// @param request - GetDoctorHiveDatabaseRequest
//
// @return GetDoctorHiveDatabaseResponse
func (client *Client) GetDoctorHiveDatabase(request *GetDoctorHiveDatabaseRequest) (_result *GetDoctorHiveDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHiveDatabaseResponse{}
	_body, _err := client.GetDoctorHiveDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a specific Hive table in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get Doctor Hive Table
//
// @param request - GetDoctorHiveTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorHiveTableResponse
func (client *Client) GetDoctorHiveTableWithOptions(request *GetDoctorHiveTableRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorHiveTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorHiveTable"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorHiveTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of a specific Hive table in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// get Doctor Hive Table
//
// @param request - GetDoctorHiveTableRequest
//
// @return GetDoctorHiveTableResponse
func (client *Client) GetDoctorHiveTable(request *GetDoctorHiveTableRequest) (_result *GetDoctorHiveTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorHiveTableResponse{}
	_body, _err := client.GetDoctorHiveTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the basic running information about a job on E-MapReduce (EMR) Doctor.
//
// Description:
//
// # Get realtime job by yarn
//
// @param request - GetDoctorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorJobResponse
func (client *Client) GetDoctorJobWithOptions(request *GetDoctorJobRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorJob"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the basic running information about a job on E-MapReduce (EMR) Doctor.
//
// Description:
//
// # Get realtime job by yarn
//
// @param request - GetDoctorJobRequest
//
// @return GetDoctorJobResponse
func (client *Client) GetDoctorJob(request *GetDoctorJobRequest) (_result *GetDoctorJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorJobResponse{}
	_body, _err := client.GetDoctorJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the analysis result report of a specified component from EMR Doctor.
//
// Description:
//
// get specify component\\"s report analysis by EMR Doctor
//
// @param request - GetDoctorReportComponentSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDoctorReportComponentSummaryResponse
func (client *Client) GetDoctorReportComponentSummaryWithOptions(request *GetDoctorReportComponentSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetDoctorReportComponentSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentType) {
		query["ComponentType"] = request.ComponentType
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDoctorReportComponentSummary"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDoctorReportComponentSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the analysis result report of a specified component from EMR Doctor.
//
// Description:
//
// get specify component\\"s report analysis by EMR Doctor
//
// @param request - GetDoctorReportComponentSummaryRequest
//
// @return GetDoctorReportComponentSummaryResponse
func (client *Client) GetDoctorReportComponentSummary(request *GetDoctorReportComponentSummaryRequest) (_result *GetDoctorReportComponentSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDoctorReportComponentSummaryResponse{}
	_body, _err := client.GetDoctorReportComponentSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetManagedScalingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagedScalingPolicyResponse
func (client *Client) GetManagedScalingPolicyWithOptions(request *GetManagedScalingPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetManagedScalingPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManagedScalingPolicy"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManagedScalingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetManagedScalingPolicyRequest
//
// @return GetManagedScalingPolicyResponse
func (client *Client) GetManagedScalingPolicy(request *GetManagedScalingPolicyRequest) (_result *GetManagedScalingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManagedScalingPolicyResponse{}
	_body, _err := client.GetManagedScalingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain the details of a node group.
//
// Description:
//
// 获取节点组详情。
//
// @param request - GetNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeGroupResponse
func (client *Client) GetNodeGroupWithOptions(request *GetNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *GetNodeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeGroup"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain the details of a node group.
//
// Description:
//
// 获取节点组详情。
//
// @param request - GetNodeGroupRequest
//
// @return GetNodeGroupResponse
func (client *Client) GetNodeGroup(request *GetNodeGroupRequest) (_result *GetNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeGroupResponse{}
	_body, _err := client.GetNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of an asynchronous operation.
//
// @param request - GetOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOperationResponse
func (client *Client) GetOperationWithOptions(request *GetOperationRequest, runtime *dara.RuntimeOptions) (_result *GetOperationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOperation"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of an asynchronous operation.
//
// @param request - GetOperationRequest
//
// @return GetOperationResponse
func (client *Client) GetOperation(request *GetOperationRequest) (_result *GetOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOperationResponse{}
	_body, _err := client.GetOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales out the node group.
//
// @param request - IncreaseNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IncreaseNodesResponse
func (client *Client) IncreaseNodesWithOptions(request *IncreaseNodesRequest, runtime *dara.RuntimeOptions) (_result *IncreaseNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigs) {
		query["ApplicationConfigs"] = request.ApplicationConfigs
	}

	if !dara.IsNil(request.AutoPayOrder) {
		query["AutoPayOrder"] = request.AutoPayOrder
	}

	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IncreaseNodeCount) {
		query["IncreaseNodeCount"] = request.IncreaseNodeCount
	}

	if !dara.IsNil(request.MinIncreaseNodeCount) {
		query["MinIncreaseNodeCount"] = request.MinIncreaseNodeCount
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PaymentDuration) {
		query["PaymentDuration"] = request.PaymentDuration
	}

	if !dara.IsNil(request.PaymentDurationUnit) {
		query["PaymentDurationUnit"] = request.PaymentDurationUnit
	}

	if !dara.IsNil(request.Promotions) {
		query["Promotions"] = request.Promotions
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IncreaseNodes"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IncreaseNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scales out the node group.
//
// @param request - IncreaseNodesRequest
//
// @return IncreaseNodesResponse
func (client *Client) IncreaseNodes(request *IncreaseNodesRequest) (_result *IncreaseNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IncreaseNodesResponse{}
	_body, _err := client.IncreaseNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an EMR resource to a resource group. A resource can belong to only one resource group.
//
// @param request - JoinResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinResourceGroupResponse
func (client *Client) JoinResourceGroupWithOptions(request *JoinResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *JoinResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("JoinResourceGroup"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an EMR resource to a resource group. A resource can belong to only one resource group.
//
// @param request - JoinResourceGroupRequest
//
// @return JoinResourceGroupResponse
func (client *Client) JoinResourceGroup(request *JoinResourceGroupRequest) (_result *JoinResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &JoinResourceGroupResponse{}
	_body, _err := client.JoinResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询API模板
//
// @param request - ListApiTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiTemplatesResponse
func (client *Client) ListApiTemplatesWithOptions(request *ListApiTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListApiTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateIds) {
		query["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiTemplates"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询API模板
//
// @param request - ListApiTemplatesRequest
//
// @return ListApiTemplatesResponse
func (client *Client) ListApiTemplates(request *ListApiTemplatesRequest) (_result *ListApiTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApiTemplatesResponse{}
	_body, _err := client.ListApiTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of the application.
//
// Description:
//
// 查询应用配置。
//
// @param request - ListApplicationConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationConfigsResponse
func (client *Client) ListApplicationConfigsWithOptions(request *ListApplicationConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigFileName) {
		query["ConfigFileName"] = request.ConfigFileName
	}

	if !dara.IsNil(request.ConfigItemKey) {
		query["ConfigItemKey"] = request.ConfigItemKey
	}

	if !dara.IsNil(request.ConfigItemValue) {
		query["ConfigItemValue"] = request.ConfigItemValue
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationConfigs"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of the application.
//
// Description:
//
// 查询应用配置。
//
// @param request - ListApplicationConfigsRequest
//
// @return ListApplicationConfigsResponse
func (client *Client) ListApplicationConfigs(request *ListApplicationConfigsRequest) (_result *ListApplicationConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationConfigsResponse{}
	_body, _err := client.ListApplicationConfigsWithOptions(request, runtime)
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
// @param request - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationNames) {
		query["ApplicationNames"] = request.ApplicationNames
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplications"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsResponse{}
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
// @param request - ListApplicationsRequest
//
// @return ListApplicationsResponse
func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of auto scaling activities.
//
// @param request - ListAutoScalingActivitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingActivitiesResponse
func (client *Client) ListAutoScalingActivitiesWithOptions(request *ListAutoScalingActivitiesRequest, runtime *dara.RuntimeOptions) (_result *ListAutoScalingActivitiesResponse, _err error) {
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

	if !dara.IsNil(request.InstanceChargeTypes) {
		query["InstanceChargeTypes"] = request.InstanceChargeTypes
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingActivityStates) {
		query["ScalingActivityStates"] = request.ScalingActivityStates
	}

	if !dara.IsNil(request.ScalingActivityType) {
		query["ScalingActivityType"] = request.ScalingActivityType
	}

	if !dara.IsNil(request.ScalingPolicyType) {
		query["ScalingPolicyType"] = request.ScalingPolicyType
	}

	if !dara.IsNil(request.ScalingRuleName) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoScalingActivities"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoScalingActivitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of auto scaling activities.
//
// @param request - ListAutoScalingActivitiesRequest
//
// @return ListAutoScalingActivitiesResponse
func (client *Client) ListAutoScalingActivities(request *ListAutoScalingActivitiesRequest) (_result *ListAutoScalingActivitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoScalingActivitiesResponse{}
	_body, _err := client.ListAutoScalingActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries E-MapReduce (EMR) clusters.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterIds) {
		query["ClusterIds"] = request.ClusterIds
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterStates) {
		query["ClusterStates"] = request.ClusterStates
	}

	if !dara.IsNil(request.ClusterTypes) {
		query["ClusterTypes"] = request.ClusterTypes
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PaymentTypes) {
		query["PaymentTypes"] = request.PaymentTypes
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
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries E-MapReduce (EMR) clusters.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of component instances.
//
// @param request - ListComponentInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentInstancesResponse
func (client *Client) ListComponentInstancesWithOptions(request *ListComponentInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListComponentInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationNames) {
		query["ApplicationNames"] = request.ApplicationNames
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentNames) {
		query["ComponentNames"] = request.ComponentNames
	}

	if !dara.IsNil(request.ComponentStates) {
		query["ComponentStates"] = request.ComponentStates
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.NodeNames) {
		query["NodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponentInstances"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListComponentInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of component instances.
//
// @param request - ListComponentInstancesRequest
//
// @return ListComponentInstancesResponse
func (client *Client) ListComponentInstances(request *ListComponentInstancesRequest) (_result *ListComponentInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComponentInstancesResponse{}
	_body, _err := client.ListComponentInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of components.
//
// @param request - ListComponentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListComponentsResponse
func (client *Client) ListComponentsWithOptions(request *ListComponentsRequest, runtime *dara.RuntimeOptions) (_result *ListComponentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationNames) {
		query["ApplicationNames"] = request.ApplicationNames
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentNames) {
		query["ComponentNames"] = request.ComponentNames
	}

	if !dara.IsNil(request.ComponentStates) {
		query["ComponentStates"] = request.ComponentStates
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListComponents"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Retrieves a list of components.
//
// @param request - ListComponentsRequest
//
// @return ListComponentsResponse
func (client *Client) ListComponents(request *ListComponentsRequest) (_result *ListComponentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListComponentsResponse{}
	_body, _err := client.ListComponentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple jobs on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list all doctor analysis apps
//
// @param request - ListDoctorApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorApplicationsResponse
func (client *Client) ListDoctorApplicationsWithOptions(request *ListDoctorApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Queues) {
		query["Queues"] = request.Queues
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorApplications"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple jobs on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list all doctor analysis apps
//
// @param request - ListDoctorApplicationsRequest
//
// @return ListDoctorApplicationsResponse
func (client *Client) ListDoctorApplications(request *ListDoctorApplicationsRequest) (_result *ListDoctorApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorApplicationsResponse{}
	_body, _err := client.ListDoctorApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about resource usage by resource type in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor analysis result of cluster engine queue view
//
// @param request - ListDoctorComputeSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorComputeSummaryResponse
func (client *Client) ListDoctorComputeSummaryWithOptions(request *ListDoctorComputeSummaryRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorComputeSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentTypes) {
		query["ComponentTypes"] = request.ComponentTypes
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorComputeSummary"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorComputeSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about resource usage by resource type in a cluster on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor analysis result of cluster engine queue view
//
// @param request - ListDoctorComputeSummaryRequest
//
// @return ListDoctorComputeSummaryResponse
func (client *Client) ListDoctorComputeSummary(request *ListDoctorComputeSummaryRequest) (_result *ListDoctorComputeSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorComputeSummaryResponse{}
	_body, _err := client.ListDoctorComputeSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about multiple HBase RegionServers at a time.
//
// Description:
//
// list Doctor HBaseRegionServers
//
// @param request - ListDoctorHBaseRegionServersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHBaseRegionServersResponse
func (client *Client) ListDoctorHBaseRegionServersWithOptions(request *ListDoctorHBaseRegionServersRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHBaseRegionServersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionServerHosts) {
		query["RegionServerHosts"] = request.RegionServerHosts
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorHBaseRegionServers"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHBaseRegionServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about multiple HBase RegionServers at a time.
//
// Description:
//
// list Doctor HBaseRegionServers
//
// @param request - ListDoctorHBaseRegionServersRequest
//
// @return ListDoctorHBaseRegionServersResponse
func (client *Client) ListDoctorHBaseRegionServers(request *ListDoctorHBaseRegionServersRequest) (_result *ListDoctorHBaseRegionServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHBaseRegionServersResponse{}
	_body, _err := client.ListDoctorHBaseRegionServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about multiple HBase tables at a time.
//
// Description:
//
// list Doctor HBaseTables
//
// @param request - ListDoctorHBaseTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHBaseTablesResponse
func (client *Client) ListDoctorHBaseTablesWithOptions(request *ListDoctorHBaseTablesRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHBaseTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableNames) {
		query["TableNames"] = request.TableNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorHBaseTables"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHBaseTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about multiple HBase tables at a time.
//
// Description:
//
// list Doctor HBaseTables
//
// @param request - ListDoctorHBaseTablesRequest
//
// @return ListDoctorHBaseTablesResponse
func (client *Client) ListDoctorHBaseTables(request *ListDoctorHBaseTablesRequest) (_result *ListDoctorHBaseTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHBaseTablesResponse{}
	_body, _err := client.ListDoctorHBaseTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// list Doctor HDFSNodes
//
// @param request - ListDoctorHDFSDirectoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHDFSDirectoriesResponse
func (client *Client) ListDoctorHDFSDirectoriesWithOptions(request *ListDoctorHDFSDirectoriesRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHDFSDirectoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.DirPath) {
		query["DirPath"] = request.DirPath
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorHDFSDirectories"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHDFSDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Description:
//
// list Doctor HDFSNodes
//
// @param request - ListDoctorHDFSDirectoriesRequest
//
// @return ListDoctorHDFSDirectoriesResponse
func (client *Client) ListDoctorHDFSDirectories(request *ListDoctorHDFSDirectoriesRequest) (_result *ListDoctorHDFSDirectoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHDFSDirectoriesResponse{}
	_body, _err := client.ListDoctorHDFSDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of Hadoop Distributed File System (HDFS) storage resources for multiple owners or groups at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor HDFS UGIs
//
// @param request - ListDoctorHDFSUGIRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHDFSUGIResponse
func (client *Client) ListDoctorHDFSUGIWithOptions(request *ListDoctorHDFSUGIRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHDFSUGIResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
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
		Action:      dara.String("ListDoctorHDFSUGI"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHDFSUGIResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of Hadoop Distributed File System (HDFS) storage resources for multiple owners or groups at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor HDFS UGIs
//
// @param request - ListDoctorHDFSUGIRequest
//
// @return ListDoctorHDFSUGIResponse
func (client *Client) ListDoctorHDFSUGI(request *ListDoctorHDFSUGIRequest) (_result *ListDoctorHDFSUGIResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHDFSUGIResponse{}
	_body, _err := client.ListDoctorHDFSUGIWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple Hive databases at a time.
//
// Description:
//
// list Doctor Hive Databases
//
// @param request - ListDoctorHiveDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHiveDatabasesResponse
func (client *Client) ListDoctorHiveDatabasesWithOptions(request *ListDoctorHiveDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHiveDatabasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DatabaseNames) {
		query["DatabaseNames"] = request.DatabaseNames
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorHiveDatabases"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHiveDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple Hive databases at a time.
//
// Description:
//
// list Doctor Hive Databases
//
// @param request - ListDoctorHiveDatabasesRequest
//
// @return ListDoctorHiveDatabasesResponse
func (client *Client) ListDoctorHiveDatabases(request *ListDoctorHiveDatabasesRequest) (_result *ListDoctorHiveDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHiveDatabasesResponse{}
	_body, _err := client.ListDoctorHiveDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple Hive tables at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor Hive Tables
//
// @param request - ListDoctorHiveTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorHiveTablesResponse
func (client *Client) ListDoctorHiveTablesWithOptions(request *ListDoctorHiveTablesRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorHiveTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DateTime) {
		query["DateTime"] = request.DateTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TableNames) {
		query["TableNames"] = request.TableNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorHiveTables"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorHiveTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the analysis results of multiple Hive tables at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list Doctor Hive Tables
//
// @param request - ListDoctorHiveTablesRequest
//
// @return ListDoctorHiveTablesResponse
func (client *Client) ListDoctorHiveTables(request *ListDoctorHiveTablesRequest) (_result *ListDoctorHiveTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorHiveTablesResponse{}
	_body, _err := client.ListDoctorHiveTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the basic running information about multiple jobs at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list realtime jobs by yarn
//
// @param request - ListDoctorJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorJobsResponse
func (client *Client) ListDoctorJobsWithOptions(request *ListDoctorJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppIds) {
		query["AppIds"] = request.AppIds
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndRange) {
		query["EndRange"] = request.EndRange
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Queues) {
		query["Queues"] = request.Queues
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartRange) {
		query["StartRange"] = request.StartRange
	}

	if !dara.IsNil(request.Types) {
		query["Types"] = request.Types
	}

	if !dara.IsNil(request.Users) {
		query["Users"] = request.Users
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorJobs"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the basic running information about multiple jobs at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list realtime jobs by yarn
//
// @param request - ListDoctorJobsRequest
//
// @return ListDoctorJobsResponse
func (client *Client) ListDoctorJobs(request *ListDoctorJobsRequest) (_result *ListDoctorJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorJobsResponse{}
	_body, _err := client.ListDoctorJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the summary of basic running information about multiple jobs at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list stats groupBy jobs by yarn
//
// @param request - ListDoctorJobsStatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorJobsStatsResponse
func (client *Client) ListDoctorJobsStatsWithOptions(request *ListDoctorJobsStatsRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorJobsStatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.EndRange) {
		query["EndRange"] = request.EndRange
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartRange) {
		query["StartRange"] = request.StartRange
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorJobsStats"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorJobsStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the summary of basic running information about multiple jobs at a time on E-MapReduce (EMR) Doctor.
//
// Description:
//
// list stats groupBy jobs by yarn
//
// @param request - ListDoctorJobsStatsRequest
//
// @return ListDoctorJobsStatsResponse
func (client *Client) ListDoctorJobsStats(request *ListDoctorJobsStatsRequest) (_result *ListDoctorJobsStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorJobsStatsResponse{}
	_body, _err := client.ListDoctorJobsStatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the overall analysis result reports of E-MapReduce (EMR) Doctor at a time.
//
// Description:
//
// list all reports analysis by emr doctor
//
// @param request - ListDoctorReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDoctorReportsResponse
func (client *Client) ListDoctorReportsWithOptions(request *ListDoctorReportsRequest, runtime *dara.RuntimeOptions) (_result *ListDoctorReportsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDoctorReports"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDoctorReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the overall analysis result reports of E-MapReduce (EMR) Doctor at a time.
//
// Description:
//
// list all reports analysis by emr doctor
//
// @param request - ListDoctorReportsRequest
//
// @return ListDoctorReportsResponse
func (client *Client) ListDoctorReports(request *ListDoctorReportsRequest) (_result *ListDoctorReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDoctorReportsResponse{}
	_body, _err := client.ListDoctorReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists instance types.
//
// @param request - ListInstanceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceTypesResponse
func (client *Client) ListInstanceTypesWithOptions(request *ListInstanceTypesRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.DeployMode) {
		query["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.IsModification) {
		query["IsModification"] = request.IsModification
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.NodeGroupType) {
		query["NodeGroupType"] = request.NodeGroupType
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReleaseVersion) {
		query["ReleaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceTypes"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists instance types.
//
// @param request - ListInstanceTypesRequest
//
// @return ListInstanceTypesResponse
func (client *Client) ListInstanceTypes(request *ListInstanceTypesRequest) (_result *ListInstanceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.ListInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of node groups in an EMR cluster.
//
// @param request - ListNodeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupIds) {
		query["NodeGroupIds"] = request.NodeGroupIds
	}

	if !dara.IsNil(request.NodeGroupNames) {
		query["NodeGroupNames"] = request.NodeGroupNames
	}

	if !dara.IsNil(request.NodeGroupStates) {
		query["NodeGroupStates"] = request.NodeGroupStates
	}

	if !dara.IsNil(request.NodeGroupTypes) {
		query["NodeGroupTypes"] = request.NodeGroupTypes
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeGroups"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of node groups in an EMR cluster.
//
// @param request - ListNodeGroupsRequest
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the node list of an EMR cluster.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupIds) {
		query["NodeGroupIds"] = request.NodeGroupIds
	}

	if !dara.IsNil(request.NodeIds) {
		query["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.NodeNames) {
		query["NodeNames"] = request.NodeNames
	}

	if !dara.IsNil(request.NodeStates) {
		query["NodeStates"] = request.NodeStates
	}

	if !dara.IsNil(request.PrivateIps) {
		query["PrivateIps"] = request.PrivateIps
	}

	if !dara.IsNil(request.PublicIps) {
		query["PublicIps"] = request.PublicIps
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the node list of an EMR cluster.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the major E-MapReduce (EMR) versions.
//
// Description:
//
// 查询主版本。
//
// @param request - ListReleaseVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReleaseVersionsResponse
func (client *Client) ListReleaseVersionsWithOptions(request *ListReleaseVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListReleaseVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterType) {
		query["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.IaasType) {
		query["IaasType"] = request.IaasType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReleaseVersions"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListReleaseVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the major E-MapReduce (EMR) versions.
//
// Description:
//
// 查询主版本。
//
// @param request - ListReleaseVersionsRequest
//
// @return ListReleaseVersionsResponse
func (client *Client) ListReleaseVersions(request *ListReleaseVersionsRequest) (_result *ListReleaseVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListReleaseVersionsResponse{}
	_body, _err := client.ListReleaseVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query EMR cluster bootstrap scripts or regular scripts.
//
// @param request - ListScriptsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScriptsResponse
func (client *Client) ListScriptsWithOptions(request *ListScriptsRequest, runtime *dara.RuntimeOptions) (_result *ListScriptsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptName) {
		query["ScriptName"] = request.ScriptName
	}

	if !dara.IsNil(request.ScriptType) {
		query["ScriptType"] = request.ScriptType
	}

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScripts"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query EMR cluster bootstrap scripts or regular scripts.
//
// @param request - ListScriptsRequest
//
// @return ListScriptsResponse
func (client *Client) ListScripts(request *ListScriptsRequest) (_result *ListScriptsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScriptsResponse{}
	_body, _err := client.ListScriptsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are bound to an EMR cluster.
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
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2021-03-20"),
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
// Queries the tags that are bound to an EMR cluster.
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
// Queries a user.
//
// Description:
//
// Queries a user.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
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

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.UserNames) {
		query["UserNames"] = request.UserNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a user.
//
// Description:
//
// Queries a user.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a custom auto scaling rule.
//
// Description:
//
// You can call this operation to configure auto scaling policies.
//
// @param request - PutAutoScalingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutAutoScalingPolicyResponse
func (client *Client) PutAutoScalingPolicyWithOptions(request *PutAutoScalingPolicyRequest, runtime *dara.RuntimeOptions) (_result *PutAutoScalingPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Constraints) {
		query["Constraints"] = request.Constraints
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScalingRules) {
		query["ScalingRules"] = request.ScalingRules
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutAutoScalingPolicy"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutAutoScalingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a custom auto scaling rule.
//
// Description:
//
// You can call this operation to configure auto scaling policies.
//
// @param request - PutAutoScalingPolicyRequest
//
// @return PutAutoScalingPolicyResponse
func (client *Client) PutAutoScalingPolicy(request *PutAutoScalingPolicyRequest) (_result *PutAutoScalingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutAutoScalingPolicyResponse{}
	_body, _err := client.PutAutoScalingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutManagedScalingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutManagedScalingPolicyResponse
func (client *Client) PutManagedScalingPolicyWithOptions(request *PutManagedScalingPolicyRequest, runtime *dara.RuntimeOptions) (_result *PutManagedScalingPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Constraints) {
		query["Constraints"] = request.Constraints
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutManagedScalingPolicy"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutManagedScalingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutManagedScalingPolicyRequest
//
// @return PutManagedScalingPolicyResponse
func (client *Client) PutManagedScalingPolicy(request *PutManagedScalingPolicyRequest) (_result *PutManagedScalingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutManagedScalingPolicyResponse{}
	_body, _err := client.PutManagedScalingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes an auto scaling policy.
//
// @param request - RemoveAutoScalingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAutoScalingPolicyResponse
func (client *Client) RemoveAutoScalingPolicyWithOptions(request *RemoveAutoScalingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveAutoScalingPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveAutoScalingPolicy"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveAutoScalingPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an auto scaling policy.
//
// @param request - RemoveAutoScalingPolicyRequest
//
// @return RemoveAutoScalingPolicyResponse
func (client *Client) RemoveAutoScalingPolicy(request *RemoveAutoScalingPolicyRequest) (_result *RemoveAutoScalingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveAutoScalingPolicyResponse{}
	_body, _err := client.RemoveAutoScalingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RunApiTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunApiTemplateResponse
func (client *Client) RunApiTemplateWithOptions(request *RunApiTemplateRequest, runtime *dara.RuntimeOptions) (_result *RunApiTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunApiTemplate"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunApiTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RunApiTemplateRequest
//
// @return RunApiTemplateResponse
func (client *Client) RunApiTemplate(request *RunApiTemplateRequest) (_result *RunApiTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunApiTemplateResponse{}
	_body, _err := client.RunApiTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manages a service deployed in a cluster. For example, you can call this operation to start pr stop a service.
//
// @param request - RunApplicationActionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunApplicationActionResponse
func (client *Client) RunApplicationActionWithOptions(request *RunApplicationActionRequest, runtime *dara.RuntimeOptions) (_result *RunApplicationActionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionName) {
		query["ActionName"] = request.ActionName
	}

	if !dara.IsNil(request.BatchSize) {
		query["BatchSize"] = request.BatchSize
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ComponentInstanceSelector) {
		query["ComponentInstanceSelector"] = request.ComponentInstanceSelector
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecuteStrategy) {
		query["ExecuteStrategy"] = request.ExecuteStrategy
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RollingExecute) {
		query["RollingExecute"] = request.RollingExecute
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunApplicationAction"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunApplicationActionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manages a service deployed in a cluster. For example, you can call this operation to start pr stop a service.
//
// @param request - RunApplicationActionRequest
//
// @return RunApplicationActionResponse
func (client *Client) RunApplicationAction(request *RunApplicationActionRequest) (_result *RunApplicationActionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunApplicationActionResponse{}
	_body, _err := client.RunApplicationActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription E-MapReduce (EMR) cluster.
//
// Description:
//
// RunCluster is an upgraded version of CreateCluster. RunCluster uses HTTPS POST requests and supports more parameters. Complex parameters, such as parameters of the object and array types, are in the JSON format and are more friendly for users who use CLI.
//
// @param tmpReq - RunClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunClusterResponse
func (client *Client) RunClusterWithOptions(tmpReq *RunClusterRequest, runtime *dara.RuntimeOptions) (_result *RunClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RunClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationConfigs) {
		request.ApplicationConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfigs, dara.String("ApplicationConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Applications) {
		request.ApplicationsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Applications, dara.String("Applications"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.BootstrapScripts) {
		request.BootstrapScriptsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BootstrapScripts, dara.String("BootstrapScripts"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeAttributes) {
		request.NodeAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeAttributes, dara.String("NodeAttributes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Promotions) {
		request.PromotionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Promotions, dara.String("Promotions"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SubscriptionConfig) {
		request.SubscriptionConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubscriptionConfig, dara.String("SubscriptionConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PromotionsShrink) {
		query["Promotions"] = request.PromotionsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigsShrink) {
		body["ApplicationConfigs"] = request.ApplicationConfigsShrink
	}

	if !dara.IsNil(request.ApplicationsShrink) {
		body["Applications"] = request.ApplicationsShrink
	}

	if !dara.IsNil(request.BootstrapScriptsShrink) {
		body["BootstrapScripts"] = request.BootstrapScriptsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		body["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.DeletionProtection) {
		body["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DeployMode) {
		body["DeployMode"] = request.DeployMode
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.NodeAttributesShrink) {
		body["NodeAttributes"] = request.NodeAttributesShrink
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !dara.IsNil(request.PaymentType) {
		body["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.ReleaseVersion) {
		body["ReleaseVersion"] = request.ReleaseVersion
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SecurityMode) {
		body["SecurityMode"] = request.SecurityMode
	}

	if !dara.IsNil(request.SubscriptionConfigShrink) {
		body["SubscriptionConfig"] = request.SubscriptionConfigShrink
	}

	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCluster"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go or subscription E-MapReduce (EMR) cluster.
//
// Description:
//
// RunCluster is an upgraded version of CreateCluster. RunCluster uses HTTPS POST requests and supports more parameters. Complex parameters, such as parameters of the object and array types, are in the JSON format and are more friendly for users who use CLI.
//
// @param request - RunClusterRequest
//
// @return RunClusterResponse
func (client *Client) RunCluster(request *RunClusterRequest) (_result *RunClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunClusterResponse{}
	_body, _err := client.RunClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds tags to a specified EMR cluster.
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2021-03-20"),
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
// Binds tags to a specified EMR cluster.
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
// Unbinds tags from a specified column in an EMR cluster. If the tag is not bound to other resources, the tag is automatically deleted.
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

	if !dara.IsNil(request.ResourceIds) {
		query["ResourceIds"] = request.ResourceIds
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeys) {
		query["TagKeys"] = request.TagKeys
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2021-03-20"),
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
// Unbinds tags from a specified column in an EMR cluster. If the tag is not bound to other resources, the tag is automatically deleted.
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
// Updates an API operation template.
//
// Description:
//
// 修改集群模板
//
// @param request - UpdateApiTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiTemplateResponse
func (client *Client) UpdateApiTemplateWithOptions(request *UpdateApiTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateApiTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiName) {
		query["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("UpdateApiTemplate"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an API operation template.
//
// Description:
//
// 修改集群模板
//
// @param request - UpdateApiTemplateRequest
//
// @return UpdateApiTemplateResponse
func (client *Client) UpdateApiTemplate(request *UpdateApiTemplateRequest) (_result *UpdateApiTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApiTemplateResponse{}
	_body, _err := client.UpdateApiTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the application configurations.
//
// @param request - UpdateApplicationConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationConfigsResponse
func (client *Client) UpdateApplicationConfigsWithOptions(request *UpdateApplicationConfigsRequest, runtime *dara.RuntimeOptions) (_result *UpdateApplicationConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationName) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigAction) {
		query["ConfigAction"] = request.ConfigAction
	}

	if !dara.IsNil(request.ConfigScope) {
		query["ConfigScope"] = request.ConfigScope
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.RefreshConfig) {
		query["RefreshConfig"] = request.RefreshConfig
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationConfigs) {
		bodyFlat["ApplicationConfigs"] = request.ApplicationConfigs
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApplicationConfigs"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApplicationConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the application configurations.
//
// @param request - UpdateApplicationConfigsRequest
//
// @return UpdateApplicationConfigsResponse
func (client *Client) UpdateApplicationConfigs(request *UpdateApplicationConfigsRequest) (_result *UpdateApplicationConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApplicationConfigsResponse{}
	_body, _err := client.UpdateApplicationConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates cluster attributes.
//
// @param request - UpdateClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterAttributeResponse
func (client *Client) UpdateClusterAttributeWithOptions(request *UpdateClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterAttribute"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates cluster attributes.
//
// @param request - UpdateClusterAttributeRequest
//
// @return UpdateClusterAttributeResponse
func (client *Client) UpdateClusterAttribute(request *UpdateClusterAttributeRequest) (_result *UpdateClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterAttributeResponse{}
	_body, _err := client.UpdateClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param tmpReq - UpdateScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScriptResponse
func (client *Client) UpdateScriptWithOptions(tmpReq *UpdateScriptRequest, runtime *dara.RuntimeOptions) (_result *UpdateScriptResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateScriptShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Script) {
		request.ScriptShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Script, dara.String("Script"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ScriptShrink) {
		query["Script"] = request.ScriptShrink
	}

	if !dara.IsNil(request.ScriptId) {
		query["ScriptId"] = request.ScriptId
	}

	if !dara.IsNil(request.ScriptType) {
		query["ScriptType"] = request.ScriptType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScript"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a bootstrap action or a common script of an E-MapReduce (EMR) cluster.
//
// @param request - UpdateScriptRequest
//
// @return UpdateScriptResponse
func (client *Client) UpdateScript(request *UpdateScriptRequest) (_result *UpdateScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScriptResponse{}
	_body, _err := client.UpdateScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a user.
//
// Description:
//
// Updates the information about a user.
//
// @param request - UpdateUserAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserAttributeResponse
func (client *Client) UpdateUserAttributeWithOptions(request *UpdateUserAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserAttribute"),
		Version:     dara.String("2021-03-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a user.
//
// Description:
//
// Updates the information about a user.
//
// @param request - UpdateUserAttributeRequest
//
// @return UpdateUserAttributeResponse
func (client *Client) UpdateUserAttribute(request *UpdateUserAttributeRequest) (_result *UpdateUserAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserAttributeResponse{}
	_body, _err := client.UpdateUserAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
