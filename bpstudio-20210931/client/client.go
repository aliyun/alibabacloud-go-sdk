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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("bpstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Switches a disaster recovery application back to the primary zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application back to the primary zone.
//
// @param request - AppFailBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppFailBackResponse
func (client *Client) AppFailBackWithOptions(request *AppFailBackRequest, runtime *dara.RuntimeOptions) (_result *AppFailBackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppFailBack"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppFailBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a disaster recovery application back to the primary zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application back to the primary zone.
//
// @param request - AppFailBackRequest
//
// @return AppFailBackResponse
func (client *Client) AppFailBack(request *AppFailBackRequest) (_result *AppFailBackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AppFailBackResponse{}
	_body, _err := client.AppFailBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches a disaster recovery application to another supported zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application to another supported zone.
//
// @param request - AppFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppFailOverResponse
func (client *Client) AppFailOverWithOptions(request *AppFailOverRequest, runtime *dara.RuntimeOptions) (_result *AppFailOverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.FailZone) {
		body["FailZone"] = request.FailZone
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AppFailOver"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AppFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches a disaster recovery application to another supported zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application to another supported zone.
//
// @param request - AppFailOverRequest
//
// @return AppFailOverResponse
func (client *Client) AppFailOver(request *AppFailOverRequest) (_result *AppFailOverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AppFailOverResponse{}
	_body, _err := client.AppFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which an application or template belongs.
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
	body := map[string]interface{}{}
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
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2021-09-31"),
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
// Changes the resource group to which an application or template belongs.
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
// Creates an application based on an official template or private template in Cloud Architect Design Tool (CADT). Before you call this operation, make sure that you understand the billing methods and prices of Alibaba Cloud services to be used in the application.
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configuration) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, dara.String("Configuration"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProcessVariables) {
		request.ProcessVariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProcessVariables, dara.String("ProcessVariables"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Variables) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, dara.String("Variables"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AreaId) {
		body["AreaId"] = request.AreaId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigurationShrink) {
		body["Configuration"] = request.ConfigurationShrink
	}

	if !dara.IsNil(request.InstancesShrink) {
		body["Instances"] = request.InstancesShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProcessVariablesShrink) {
		body["ProcessVariables"] = request.ProcessVariablesShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.VariablesShrink) {
		body["Variables"] = request.VariablesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application based on an official template or private template in Cloud Architect Design Tool (CADT). Before you call this operation, make sure that you understand the billing methods and prices of Alibaba Cloud services to be used in the application.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
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
// Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.“ You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Description:
//
// Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.“ You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys an application after the payment.
//
// @param request - DeployApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationResponse
func (client *Client) DeployApplicationWithOptions(request *DeployApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Deploys an application after the payment.
//
// @param request - DeployApplicationRequest
//
// @return DeployApplicationResponse
func (client *Client) DeployApplication(request *DeployApplicationRequest) (_result *DeployApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployApplicationResponse{}
	_body, _err := client.DeployApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronous execution of product operation functions.
//
// @param tmpReq - ExecuteOperationASyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteOperationASyncResponse
func (client *Client) ExecuteOperationASyncWithOptions(tmpReq *ExecuteOperationASyncRequest, runtime *dara.RuntimeOptions) (_result *ExecuteOperationASyncResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteOperationASyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AttributesShrink) {
		body["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceType) {
		body["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteOperationASync"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteOperationASyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronous execution of product operation functions.
//
// @param request - ExecuteOperationASyncRequest
//
// @return ExecuteOperationASyncResponse
func (client *Client) ExecuteOperationASync(request *ExecuteOperationASyncRequest) (_result *ExecuteOperationASyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteOperationASyncResponse{}
	_body, _err := client.ExecuteOperationASyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 维护应用下资源API（同步操作）
//
// @param tmpReq - ExecuteOperationSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteOperationSyncResponse
func (client *Client) ExecuteOperationSyncWithOptions(tmpReq *ExecuteOperationSyncRequest, runtime *dara.RuntimeOptions) (_result *ExecuteOperationSyncResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteOperationSyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Attributes) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, dara.String("Attributes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.AttributesShrink) {
		body["Attributes"] = request.AttributesShrink
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceType) {
		body["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteOperationSync"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteOperationSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 维护应用下资源API（同步操作）
//
// @param request - ExecuteOperationSyncRequest
//
// @return ExecuteOperationSyncResponse
func (client *Client) ExecuteOperationSync(request *ExecuteOperationSyncRequest) (_result *ExecuteOperationSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteOperationSyncResponse{}
	_body, _err := client.ExecuteOperationSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The URL of the application topology image.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplication"),
		Version:     dara.String("2021-09-31"),
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
// The URL of the application topology image.
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
// 获取应用输入参数
//
// @param request - GetApplicationVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationVariablesResponse
func (client *Client) GetApplicationVariablesWithOptions(request *GetApplicationVariablesRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationVariablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationVariables"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用输入参数
//
// @param request - GetApplicationVariablesRequest
//
// @return GetApplicationVariablesResponse
func (client *Client) GetApplicationVariables(request *GetApplicationVariablesRequest) (_result *GetApplicationVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationVariablesResponse{}
	_body, _err := client.GetApplicationVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取需要重新配置的变量列表
//
// @param request - GetApplicationVariables4FailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationVariables4FailResponse
func (client *Client) GetApplicationVariables4FailWithOptions(request *GetApplicationVariables4FailRequest, runtime *dara.RuntimeOptions) (_result *GetApplicationVariables4FailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplicationVariables4Fail"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplicationVariables4FailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取需要重新配置的变量列表
//
// @param request - GetApplicationVariables4FailRequest
//
// @return GetApplicationVariables4FailResponse
func (client *Client) GetApplicationVariables4Fail(request *GetApplicationVariables4FailRequest) (_result *GetApplicationVariables4FailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplicationVariables4FailResponse{}
	_body, _err := client.GetApplicationVariables4FailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the result of an operation that is performed on a service instance.
//
// @param request - GetExecuteOperationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteOperationResultResponse
func (client *Client) GetExecuteOperationResultWithOptions(request *GetExecuteOperationResultRequest, runtime *dara.RuntimeOptions) (_result *GetExecuteOperationResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		body["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExecuteOperationResult"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExecuteOperationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Asynchronously queries the result of an operation that is performed on a service instance.
//
// @param request - GetExecuteOperationResultRequest
//
// @return GetExecuteOperationResultResponse
func (client *Client) GetExecuteOperationResult(request *GetExecuteOperationResultRequest) (_result *GetExecuteOperationResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExecuteOperationResultResponse{}
	_body, _err := client.GetExecuteOperationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a disaster recovery switchover task by task ID.
//
// Description:
//
// You can call this operation to query the status of a disaster recovery switchover task by task ID.
//
// @param request - GetFoTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFoTaskStatusResponse
func (client *Client) GetFoTaskStatusWithOptions(request *GetFoTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *GetFoTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFoTaskStatus"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFoTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a disaster recovery switchover task by task ID.
//
// Description:
//
// You can call this operation to query the status of a disaster recovery switchover task by task ID.
//
// @param request - GetFoTaskStatusRequest
//
// @return GetFoTaskStatusResponse
func (client *Client) GetFoTaskStatus(request *GetFoTaskStatusRequest) (_result *GetFoTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFoTaskStatusResponse{}
	_body, _err := client.GetFoTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取模板变参可选值
//
// @param tmpReq - GetLinkageAttributesTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLinkageAttributesTemplateResponse
func (client *Client) GetLinkageAttributesTemplateWithOptions(tmpReq *GetLinkageAttributesTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetLinkageAttributesTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetLinkageAttributesTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Variables) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, dara.String("Variables"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AreaId) {
		body["AreaId"] = request.AreaId
	}

	if !dara.IsNil(request.InstancesShrink) {
		body["Instances"] = request.InstancesShrink
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TargetVariable) {
		body["TargetVariable"] = request.TargetVariable
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.VariablesShrink) {
		body["Variables"] = request.VariablesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLinkageAttributesTemplate"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLinkageAttributesTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取模板变参可选值
//
// @param request - GetLinkageAttributesTemplateRequest
//
// @return GetLinkageAttributesTemplateResponse
func (client *Client) GetLinkageAttributesTemplate(request *GetLinkageAttributesTemplateRequest) (_result *GetLinkageAttributesTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLinkageAttributesTemplateResponse{}
	_body, _err := client.GetLinkageAttributesTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones where the specified disaster recovery service can be switched.
//
// Description:
//
// You can call this operation to query the zones where the specified disaster recovery service can be switched.
//
// @param request - GetPotentialFailZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPotentialFailZonesResponse
func (client *Client) GetPotentialFailZonesWithOptions(request *GetPotentialFailZonesRequest, runtime *dara.RuntimeOptions) (_result *GetPotentialFailZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsPlanId) {
		body["IsPlanId"] = request.IsPlanId
	}

	if !dara.IsNil(request.ObjectId) {
		body["ObjectId"] = request.ObjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPotentialFailZones"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPotentialFailZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones where the specified disaster recovery service can be switched.
//
// Description:
//
// You can call this operation to query the zones where the specified disaster recovery service can be switched.
//
// @param request - GetPotentialFailZonesRequest
//
// @return GetPotentialFailZonesResponse
func (client *Client) GetPotentialFailZones(request *GetPotentialFailZonesRequest) (_result *GetPotentialFailZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPotentialFailZonesResponse{}
	_body, _err := client.GetPotentialFailZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取询价应用变配记录
//
// @param request - GetResource4ModifyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResource4ModifyRecordResponse
func (client *Client) GetResource4ModifyRecordWithOptions(request *GetResource4ModifyRecordRequest, runtime *dara.RuntimeOptions) (_result *GetResource4ModifyRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResource4ModifyRecord"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResource4ModifyRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取询价应用变配记录
//
// @param request - GetResource4ModifyRecordRequest
//
// @return GetResource4ModifyRecordResponse
func (client *Client) GetResource4ModifyRecord(request *GetResource4ModifyRecordRequest) (_result *GetResource4ModifyRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResource4ModifyRecordResponse{}
	_body, _err := client.GetResource4ModifyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取询价结果
//
// @param request - GetResult4QueryInstancePrice4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResult4QueryInstancePrice4ModifyResponse
func (client *Client) GetResult4QueryInstancePrice4ModifyWithOptions(request *GetResult4QueryInstancePrice4ModifyRequest, runtime *dara.RuntimeOptions) (_result *GetResult4QueryInstancePrice4ModifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResult4QueryInstancePrice4Modify"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResult4QueryInstancePrice4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取询价结果
//
// @param request - GetResult4QueryInstancePrice4ModifyRequest
//
// @return GetResult4QueryInstancePrice4ModifyResponse
func (client *Client) GetResult4QueryInstancePrice4Modify(request *GetResult4QueryInstancePrice4ModifyRequest) (_result *GetResult4QueryInstancePrice4ModifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResult4QueryInstancePrice4ModifyResponse{}
	_body, _err := client.GetResult4QueryInstancePrice4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets template images and information about architecture diagrams.
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets template images and information about architecture diagrams.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
//
// Summary:
//
// Obtains a temporary token that is used to read the architecture diagram. The validity period of the token is 30 minutes.
//
// Description:
//
//	Danger:  This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.
//
// @param request - GetTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
//
// Summary:
//
// Obtains a temporary token that is used to read the architecture diagram. The validity period of the token is 30 minutes.
//
// Description:
//
//	Danger:  This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
// Deprecated
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Prepares for application switchover and initiates a switchover task.
//
// Description:
//
// You can call this operation to prepare for application switchover and initiate a switchover task.
//
// @param request - InitAppFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitAppFailOverResponse
func (client *Client) InitAppFailOverWithOptions(request *InitAppFailOverRequest, runtime *dara.RuntimeOptions) (_result *InitAppFailOverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InitAppFailOver"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InitAppFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Prepares for application switchover and initiates a switchover task.
//
// Description:
//
// You can call this operation to prepare for application switchover and initiate a switchover task.
//
// @param request - InitAppFailOverRequest
//
// @return InitAppFailOverResponse
func (client *Client) InitAppFailOver(request *InitAppFailOverRequest) (_result *InitAppFailOverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InitAppFailOverResponse{}
	_body, _err := client.InitAppFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API provides a list of all applications under the current user. The optional keyword parameter defines the keywords contained in the application name.
//
// @param request - ListApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationResponse
func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderType) {
		body["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ShowHide) {
		body["ShowHide"] = request.ShowHide
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// This API provides a list of all applications under the current user. The optional keyword parameter defines the keywords contained in the application name.
//
// @param request - ListApplicationRequest
//
// @return ListApplicationResponse
func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all disaster recovery plans of the current account.
//
// Description:
//
// Queries the information about all disaster recovery plans of the current account.
//
// @param request - ListFoCreatedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoCreatedAppsResponse
func (client *Client) ListFoCreatedAppsWithOptions(runtime *dara.RuntimeOptions) (_result *ListFoCreatedAppsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListFoCreatedApps"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFoCreatedAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all disaster recovery plans of the current account.
//
// Description:
//
// Queries the information about all disaster recovery plans of the current account.
//
// @return ListFoCreatedAppsResponse
func (client *Client) ListFoCreatedApps() (_result *ListFoCreatedAppsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFoCreatedAppsResponse{}
	_body, _err := client.ListFoCreatedAppsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of one or more applications or templates.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		bodyFlat["Tag"] = request.Tag
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2021-09-31"),
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
// Queries the tags of one or more applications or templates.
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
// Queries templates, including information such as the template name, architecture image URL, and URL of the serialized architecture image file.
//
// @param tmpReq - ListTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateResponse
func (client *Client) ListTemplateWithOptions(tmpReq *ListTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderType) {
		body["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		body["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.TagList) {
		body["TagList"] = request.TagList
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplate"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries templates, including information such as the template name, architecture image URL, and URL of the serialized architecture image file.
//
// @param request - ListTemplateRequest
//
// @return ListTemplateResponse
func (client *Client) ListTemplate(request *ListTemplateRequest) (_result *ListTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplateResponse{}
	_body, _err := client.ListTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交应用变配
//
// @param tmpReq - ModifyApplicationSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationSpecResponse
func (client *Client) ModifyApplicationSpecWithOptions(tmpReq *ModifyApplicationSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyApplicationSpecResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ModifyApplicationSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceSpec) {
		request.InstanceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceSpec, dara.String("InstanceSpec"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceSpecShrink) {
		body["InstanceSpec"] = request.InstanceSpecShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyApplicationSpec"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyApplicationSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交应用变配
//
// @param request - ModifyApplicationSpecRequest
//
// @return ModifyApplicationSpecResponse
func (client *Client) ModifyApplicationSpec(request *ModifyApplicationSpecRequest) (_result *ModifyApplicationSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyApplicationSpecResponse{}
	_body, _err := client.ModifyApplicationSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询变配价格
//
// @param tmpReq - QueryInstancePrice4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstancePrice4ModifyResponse
func (client *Client) QueryInstancePrice4ModifyWithOptions(tmpReq *QueryInstancePrice4ModifyRequest, runtime *dara.RuntimeOptions) (_result *QueryInstancePrice4ModifyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryInstancePrice4ModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Configuration) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, dara.String("Configuration"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ConfigurationShrink) {
		body["Configuration"] = request.ConfigurationShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstancePrice4Modify"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstancePrice4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变配价格
//
// @param request - QueryInstancePrice4ModifyRequest
//
// @return QueryInstancePrice4ModifyResponse
func (client *Client) QueryInstancePrice4Modify(request *QueryInstancePrice4ModifyRequest) (_result *QueryInstancePrice4ModifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstancePrice4ModifyResponse{}
	_body, _err := client.QueryInstancePrice4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询变配规格列表
//
// @param tmpReq - QueryInstanceSpec4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceSpec4ModifyResponse
func (client *Client) QueryInstanceSpec4ModifyWithOptions(tmpReq *QueryInstanceSpec4ModifyRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceSpec4ModifyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryInstanceSpec4ModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MethodName) {
		body["MethodName"] = request.MethodName
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstanceSpec4Modify"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceSpec4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变配规格列表
//
// @param request - QueryInstanceSpec4ModifyRequest
//
// @return QueryInstanceSpec4ModifyResponse
func (client *Client) QueryInstanceSpec4Modify(request *QueryInstanceSpec4ModifyRequest) (_result *QueryInstanceSpec4ModifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstanceSpec4ModifyResponse{}
	_body, _err := client.QueryInstanceSpec4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新配置应用
//
// @param request - ReConfigApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReConfigApplicationResponse
func (client *Client) ReConfigApplicationWithOptions(request *ReConfigApplicationRequest, runtime *dara.RuntimeOptions) (_result *ReConfigApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Variables) {
		body["Variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReConfigApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReConfigApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新配置应用
//
// @param request - ReConfigApplicationRequest
//
// @return ReConfigApplicationResponse
func (client *Client) ReConfigApplication(request *ReConfigApplicationRequest) (_result *ReConfigApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReConfigApplicationResponse{}
	_body, _err := client.ReConfigApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the resources of an application.
//
// @param request - ReleaseApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseApplicationResponse
func (client *Client) ReleaseApplicationWithOptions(request *ReleaseApplicationRequest, runtime *dara.RuntimeOptions) (_result *ReleaseApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the resources of an application.
//
// @param request - ReleaseApplicationRequest
//
// @return ReleaseApplicationResponse
func (client *Client) ReleaseApplication(request *ReleaseApplicationRequest) (_result *ReleaseApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseApplicationResponse{}
	_body, _err := client.ReleaseApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the resources of an application. ValidateApplication is an asynchronous operation. You can call the GetApplication operation to query the verification result.
//
// @param request - ValidateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateApplicationResponse
func (client *Client) ValidateApplicationWithOptions(request *ValidateApplicationRequest, runtime *dara.RuntimeOptions) (_result *ValidateApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		query["ApplicationId"] = request.ApplicationId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the resources of an application. ValidateApplication is an asynchronous operation. You can call the GetApplication operation to query the verification result.
//
// @param request - ValidateApplicationRequest
//
// @return ValidateApplicationResponse
func (client *Client) ValidateApplication(request *ValidateApplicationRequest) (_result *ValidateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateApplicationResponse{}
	_body, _err := client.ValidateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the prices of resources of an application. You can call the GetApplication operation to obtain the query results.
//
// @param request - ValuateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValuateApplicationResponse
func (client *Client) ValuateApplicationWithOptions(request *ValuateApplicationRequest, runtime *dara.RuntimeOptions) (_result *ValuateApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValuateApplication"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValuateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prices of resources of an application. You can call the GetApplication operation to obtain the query results.
//
// @param request - ValuateApplicationRequest
//
// @return ValuateApplicationResponse
func (client *Client) ValuateApplication(request *ValuateApplicationRequest) (_result *ValuateApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValuateApplicationResponse{}
	_body, _err := client.ValuateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price of a template.
//
// @param tmpReq - ValuateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValuateTemplateResponse
func (client *Client) ValuateTemplateWithOptions(tmpReq *ValuateTemplateRequest, runtime *dara.RuntimeOptions) (_result *ValuateTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ValuateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Instances) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, dara.String("Instances"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Variables) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, dara.String("Variables"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AreaId) {
		body["AreaId"] = request.AreaId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InstancesShrink) {
		body["Instances"] = request.InstancesShrink
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.VariablesShrink) {
		body["Variables"] = request.VariablesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValuateTemplate"),
		Version:     dara.String("2021-09-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValuateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price of a template.
//
// @param request - ValuateTemplateRequest
//
// @return ValuateTemplateResponse
func (client *Client) ValuateTemplate(request *ValuateTemplateRequest) (_result *ValuateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValuateTemplateResponse{}
	_body, _err := client.ValuateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
