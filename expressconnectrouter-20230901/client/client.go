// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("ExpressConnectRouter")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hongkong":           dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-zhengzhou-jva":      dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-beijing":            dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-nanjing":            dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-northeast-2":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-northeast-1":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-fuzhou":             dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-chengdu":            dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-wuhan-lr":           dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-guangzhou":          dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-heyuan":             dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-southeast-3":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-huhehaote":          dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-southeast-5":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-southeast-6":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"ap-southeast-7":        dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":           dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"us-west-1":             dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"us-east-1":             dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"eu-central-1":          dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"eu-west-1":             dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"me-central-1":          dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-heyuan-acdr-1":      dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("expressconnectrouter.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("expressconnectrouter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables log delivery for a flow log.
//
// @param request - ActivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLogWithOptions(request *ActivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActivateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogId) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables log delivery for a flow log.
//
// @param request - ActivateFlowLogRequest
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLog(request *ActivateFlowLogRequest) (_result *ActivateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateFlowLogResponse{}
	_body, _err := client.ActivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a child instance (VBR) to an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the **AttachExpressConnectRouterChildInstance*	- operation to load a child instance (VBR) to an Express Connect Router (ECR), make sure that the Express Connect Router (ECR) is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstanceWithOptions(request *AttachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceOwnerId) {
		body["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a child instance (VBR) to an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the **AttachExpressConnectRouterChildInstance*	- operation to load a child instance (VBR) to an Express Connect Router (ECR), make sure that the Express Connect Router (ECR) is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstance(request *AttachExpressConnectRouterChildInstanceRequest) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.AttachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a specified region to an Express Connect Router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouterWithOptions(request *CheckAddRegionToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FreshRegionId) {
		body["FreshRegionId"] = request.FreshRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAddRegionToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a specified region to an Express Connect Router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouter(request *CheckAddRegionToExpressConnectRouterRequest) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CheckAddRegionToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an Express Connect Router (ECR) instance, the instance status is **Active**.
//
// @param request - CreateExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouterWithOptions(request *CreateExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlibabaSideAsn) {
		body["AlibabaSideAsn"] = request.AlibabaSideAsn
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an Express Connect Router (ECR) instance, the instance status is **Active**.
//
// @param request - CreateExpressConnectRouterRequest
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouter(request *CreateExpressConnectRouterRequest) (_result *CreateExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CreateExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an Express Connect Router (ECR) with a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociationWithOptions(request *CreateExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedPrefixes) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !dara.IsNil(request.AllowedPrefixesMode) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !dara.IsNil(request.AssociationRegionId) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateAttachment) {
		body["CreateAttachment"] = request.CreateAttachment
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterOwnerId) {
		body["TransitRouterOwnerId"] = request.TransitRouterOwnerId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		body["VpcOwnerId"] = request.VpcOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an Express Connect Router (ECR) with a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociation(request *CreateExpressConnectRouterAssociationRequest) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CreateExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a flow log and enables log delivery by default.
//
// @param request - CreateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLogWithOptions(request *CreateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowLogResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SamplingRate) {
		query["SamplingRate"] = request.SamplingRate
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetSlsRegionId) {
		query["TargetSlsRegionId"] = request.TargetSlsRegionId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogName) {
		body["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow log and enables log delivery by default.
//
// @param request - CreateFlowLogRequest
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLog(request *CreateFlowLogRequest) (_result *CreateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CreateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an association between an Express Connect Router (ECR) and a route map.
//
// @param request - CreateRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteMapAssociationResponse
func (client *Client) CreateRouteMapAssociationWithOptions(request *CreateRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionIdList) {
		query["RegionIdList"] = request.RegionIdList
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteMapAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an association between an Express Connect Router (ECR) and a route map.
//
// @param request - CreateRouteMapAssociationRequest
//
// @return CreateRouteMapAssociationResponse
func (client *Client) CreateRouteMapAssociation(request *CreateRouteMapAssociationRequest) (_result *CreateRouteMapAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRouteMapAssociationResponse{}
	_body, _err := client.CreateRouteMapAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops log delivery for a flow log.
//
// @param request - DeactivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLogWithOptions(request *DeactivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactivateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogId) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops log delivery for a flow log.
//
// @param request - DeactivateFlowLogRequest
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLog(request *DeactivateFlowLogRequest) (_result *DeactivateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactivateFlowLogResponse{}
	_body, _err := client.DeactivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the DeleteExpressConnectRouter operation to delete an Express Connect Router (ECR) instance, note the following:
//
// - Before you delete an Express Connect Router (ECR) instance, you must delete the resources associated with the Express Connect Router (ECR) instance.
//
// - You can delete only Express Connect Router (ECR) instances that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouterWithOptions(request *DeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the DeleteExpressConnectRouter operation to delete an Express Connect Router (ECR) instance, note the following:
//
// - Before you delete an Express Connect Router (ECR) instance, you must delete the resources associated with the Express Connect Router (ECR) instance.
//
// - You can delete only Express Connect Router (ECR) instances that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouter(request *DeleteExpressConnectRouterRequest) (_result *DeleteExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.DeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect Router (ECR) from a virtual private cloud (VPC) or a transit router.
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociationWithOptions(request *DeleteExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteAttachment) {
		body["DeleteAttachment"] = request.DeleteAttachment
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect Router (ECR) from a virtual private cloud (VPC) or a transit router.
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociation(request *DeleteExpressConnectRouterAssociationRequest) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.DeleteExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a flow log.
//
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithOptions(request *DeleteFlowlogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowlog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a flow log.
//
// @param request - DeleteFlowlogRequest
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlog(request *DeleteFlowlogRequest) (_result *DeleteFlowlogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.DeleteFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the association between an Express Connect Router (ECR) and a route policy.
//
// @param request - DeleteRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteMapAssociationResponse
func (client *Client) DeleteRouteMapAssociationWithOptions(request *DeleteRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteMapAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the association between an Express Connect Router (ECR) and a route policy.
//
// @param request - DeleteRouteMapAssociationRequest
//
// @return DeleteRouteMapAssociationResponse
func (client *Client) DeleteRouteMapAssociation(request *DeleteRouteMapAssociationRequest) (_result *DeleteRouteMapAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteMapAssociationResponse{}
	_body, _err := client.DeleteRouteMapAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the disabled route entry of an Express Connect Router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDisabledExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disabled route entry of an Express Connect Router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntries(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect Router (ECR) instances.
//
// @param request - DescribeExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouterWithOptions(request *DescribeExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect Router (ECR) instances.
//
// @param request - DescribeExpressConnectRouterRequest
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouter(request *DescribeExpressConnectRouterRequest) (_result *DescribeExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.DescribeExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the history of prefix routes for an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterAllowedPrefixHistory"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history of prefix routes for an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistory(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociationWithOptions(request *DescribeExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.AssociationNodeType) {
		body["AssociationNodeType"] = request.AssociationNodeType
	}

	if !dara.IsNil(request.AssociationRegionId) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TransitRouterId) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociation(request *DescribeExpressConnectRouterAssociationRequest) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.DescribeExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstanceWithOptions(request *DescribeExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstance(request *DescribeExpressConnectRouterChildInstanceRequest) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DescribeExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request *DescribeExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterInterRegionTransitMode"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitMode(request *DescribeExpressConnectRouterInterRegionTransitModeRequest) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions where Express Connect Router (ECR) resources are located.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegionWithOptions(request *DescribeExpressConnectRouterRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterRegion"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where Express Connect Router (ECR) resources are located.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegion(request *DescribeExpressConnectRouterRegionRequest) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.DescribeExpressConnectRouterRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntriesWithOptions(request *DescribeExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AsPath) {
		body["AsPath"] = request.AsPath
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Community) {
		body["Community"] = request.Community
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	if !dara.IsNil(request.QueryRegionId) {
		body["QueryRegionId"] = request.QueryRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntries(request *DescribeExpressConnectRouterRouteEntriesRequest) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries flow logs.
//
// @param request - DescribeFlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogsWithOptions(request *DescribeFlowLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowLogs"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries flow logs.
//
// @param request - DescribeFlowLogsRequest
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (_result *DescribeFlowLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowLogsResponse{}
	_body, _err := client.DescribeFlowLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network instances authorized to an Express Connect Router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouterWithOptions(request *DescribeInstanceGrantedToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallerType) {
		body["CallerType"] = request.CallerType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		body["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagModels) {
		body["TagModels"] = request.TagModels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceGrantedToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network instances authorized to an Express Connect Router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouter(request *DescribeInstanceGrantedToExpressConnectRouterRequest) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.DescribeInstanceGrantedToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a virtual border router (VBR) from an Express Connect Router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to detach a VBR instance from an Express Connect Router (ECR), ensure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstanceWithOptions(request *DetachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a virtual border router (VBR) from an Express Connect Router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to detach a VBR instance from an Express Connect Router (ECR), ensure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstance(request *DetachExpressConnectRouterChildInstanceRequest) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DetachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a route entry of an Express Connect Router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntriesWithOptions(request *DisableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a route entry of an Express Connect Router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntries(request *DisableExpressConnectRouterRouteEntriesRequest) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DisableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables route entries for an Express Connect Router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntriesWithOptions(request *EnableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables route entries for an Express Connect Router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntries(request *EnableExpressConnectRouterRouteEntriesRequest) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.EnableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Force deletes an Express Connect Router (ECR) and detaches its associated virtual private clouds (VPCs), Transit Routers (TRs), and virtual border routers (VBRs).
//
// Description:
//
// - Force deleting an Express Connect Router (ECR) instance detaches all associated resources. Before you proceed, confirm that this operation does not disrupt your services.
//
// - You can delete only ECR instances that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouterWithOptions(request *ForceDeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForceDeleteExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Force deletes an Express Connect Router (ECR) and detaches its associated virtual private clouds (VPCs), Transit Routers (TRs), and virtual border routers (VBRs).
//
// Description:
//
// - Force deleting an Express Connect Router (ECR) instance detaches all associated resources. Before you proceed, confirm that this operation does not disrupt your services.
//
// - You can delete only ECR instances that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouter(request *ForceDeleteExpressConnectRouterRequest) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.ForceDeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a virtual private cloud (VPC) or virtual border router (VBR) authorization to a cross-account Express Connect Router (ECR).
//
// Description:
//
// To associate a cross-account network instance with an Express Connect Router (ECR) instance, you must grant the network instance authorization to the Express Connect Router (ECR) instance.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouterWithOptions(request *GrantInstanceToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.EcrOwnerAliUid) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantInstanceToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a virtual private cloud (VPC) or virtual border router (VBR) authorization to a cross-account Express Connect Router (ECR).
//
// Description:
//
// To associate a cross-account network instance with an Express Connect Router (ECR) instance, you must grant the network instance authorization to the Express Connect Router (ECR) instance.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouter(request *GrantInstanceToExpressConnectRouterRequest) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.GrantInstanceToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that support Express Connect Router (ECR).
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegionWithOptions(request *ListExpressConnectRouterSupportedRegionRequest, runtime *dara.RuntimeOptions) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExpressConnectRouterSupportedRegion"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that support Express Connect Router (ECR).
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegion(request *ListExpressConnectRouterSupportedRegionRequest) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.ListExpressConnectRouterSupportedRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the association between an Express Connect Router (ECR) and a route map.
//
// @param request - ListRouteMapAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRouteMapAssociationsResponse
func (client *Client) ListRouteMapAssociationsWithOptions(request *ListRouteMapAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListRouteMapAssociationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRouteMapAssociations"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRouteMapAssociationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an Express Connect Router (ECR) and a route map.
//
// @param request - ListRouteMapAssociationsRequest
//
// @return ListRouteMapAssociationsResponse
func (client *Client) ListRouteMapAssociations(request *ListRouteMapAssociationsRequest) (_result *ListRouteMapAssociationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRouteMapAssociationsResponse{}
	_body, _err := client.ListRouteMapAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of tags bound to an Express Connect Router (ECR).
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2023-09-01"),
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
// Queries the list of tags bound to an Express Connect Router (ECR).
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
// Modifies the properties of an Express Connect Router (ECR), such as its name.
//
// Description:
//
// You can modify an Express Connect Router (ECR) instance only when it is in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouterWithOptions(request *ModifyExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of an Express Connect Router (ECR), such as its name.
//
// Description:
//
// You can modify an Express Connect Router (ECR) instance only when it is in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouter(request *ModifyExpressConnectRouterRequest) (_result *ModifyExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.ModifyExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationResponse
func (client *Client) ModifyExpressConnectRouterAssociationWithOptions(request *ModifyExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationRequest
//
// @return ModifyExpressConnectRouterAssociationResponse
func (client *Client) ModifyExpressConnectRouterAssociation(request *ModifyExpressConnectRouterAssociationRequest) (_result *ModifyExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterAssociationResponse{}
	_body, _err := client.ModifyExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the prefixes for a virtual private cloud (VPC) or a transit router that is attached to an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedPrefixes) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !dara.IsNil(request.AllowedPrefixesMode) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.OwnerAccount) {
		body["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterAssociationAllowedPrefix"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the prefixes for a virtual private cloud (VPC) or a transit router that is attached to an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefix(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance (VBR) associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterChildInstanceResponse
func (client *Client) ModifyExpressConnectRouterChildInstanceWithOptions(request *ModifyExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance (VBR) associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterChildInstanceRequest
//
// @return ModifyExpressConnectRouterChildInstanceResponse
func (client *Client) ModifyExpressConnectRouterChildInstance(request *ModifyExpressConnectRouterChildInstanceRequest) (_result *ModifyExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.ModifyExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the inter-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request *ModifyExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.TransitModeList) {
		body["TransitModeList"] = request.TransitModeList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterInterRegionTransitMode"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the inter-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitMode(request *ModifyExpressConnectRouterInterRegionTransitModeRequest) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, sampling rate, and sampling interval of a flow log.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithOptions(request *ModifyFlowLogAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
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

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SamplingRate) {
		query["SamplingRate"] = request.SamplingRate
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogName) {
		body["FlowLogName"] = request.FlowLogName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFlowLogAttribute"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, sampling rate, and sampling interval of a flow log.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (_result *ModifyFlowLogAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.ModifyFlowLogAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which an Express Connect Router (ECR) belongs.
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2023-09-01"),
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
// Modifies the resource group to which an Express Connect Router (ECR) belongs.
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
// Revokes the authorization of a virtual private cloud (VPC) or virtual border router (VBR) on a cross-account Express Connect Router (ECR).
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouterWithOptions(request *RevokeInstanceFromExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.EcrOwnerAliUid) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeInstanceFromExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the authorization of a virtual private cloud (VPC) or virtual border router (VBR) on a cross-account Express Connect Router (ECR).
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouter(request *RevokeInstanceFromExpressConnectRouterRequest) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.RevokeInstanceFromExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the inter-region bandwidth limits for an Express Connect Router (ECR).
//
// Description:
//
// You can update an ECR instance only when it is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SynchronizeExpressConnectRouterInterRegionBandwidth"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the inter-region bandwidth limits for an Express Connect Router (ECR).
//
// Description:
//
// You can update an ECR instance only when it is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidth(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds tags to an Express Connect Router (ECR). Each call to this operation binds tags to a single ECR instance, but you can bind multiple tags at a time.
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2023-09-01"),
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
// Binds tags to an Express Connect Router (ECR). Each call to this operation binds tags to a single ECR instance, but you can bind multiple tags at a time.
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
// Unbinds tags from an Express Connect Router (ECR).
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		body["All"] = request.All
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		body["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2023-09-01"),
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
// Unbinds tags from an Express Connect Router (ECR).
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
// Updates the association between an Express Connect Router (ECR) and a routing policy.
//
// @param request - UpdateRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteMapAssociationResponse
func (client *Client) UpdateRouteMapAssociationWithOptions(request *UpdateRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.RegionIdList) {
		query["RegionIdList"] = request.RegionIdList
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteMapAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the association between an Express Connect Router (ECR) and a routing policy.
//
// @param request - UpdateRouteMapAssociationRequest
//
// @return UpdateRouteMapAssociationResponse
func (client *Client) UpdateRouteMapAssociation(request *UpdateRouteMapAssociationRequest) (_result *UpdateRouteMapAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRouteMapAssociationResponse{}
	_body, _err := client.UpdateRouteMapAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
