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
	client.Endpoint, _err = client.GetEndpoint(dara.String("websitebuild"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 提交创建Logo任务
//
// @param request - CreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTaskWithOptions(request *CreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateLogoTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogoVersion) {
		query["LogoVersion"] = request.LogoVersion
	}

	if !dara.IsNil(request.NegativePrompt) {
		query["NegativePrompt"] = request.NegativePrompt
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交创建Logo任务
//
// @param request - CreateLogoTaskRequest
//
// @return CreateLogoTaskResponse
func (client *Client) CreateLogoTask(request *CreateLogoTaskRequest) (_result *CreateLogoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogoTaskResponse{}
	_body, _err := client.CreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Logo创建任务
//
// @param request - GetCreateLogoTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTaskWithOptions(request *GetCreateLogoTaskRequest, runtime *dara.RuntimeOptions) (_result *GetCreateLogoTaskResponse, _err error) {
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
		Action:      dara.String("GetCreateLogoTask"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Logo创建任务
//
// @param request - GetCreateLogoTaskRequest
//
// @return GetCreateLogoTaskResponse
func (client *Client) GetCreateLogoTask(request *GetCreateLogoTaskRequest) (_result *GetCreateLogoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreateLogoTaskResponse{}
	_body, _err := client.GetCreateLogoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用
//
// @param request - OperateAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartnerWithOptions(request *OperateAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用
//
// @param request - OperateAppInstanceForPartnerRequest
//
// @return OperateAppInstanceForPartnerResponse
func (client *Client) OperateAppInstanceForPartner(request *OperateAppInstanceForPartnerRequest) (_result *OperateAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppInstanceForPartnerResponse{}
	_body, _err := client.OperateAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用服务
//
// @param request - OperateAppServiceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartnerWithOptions(request *OperateAppServiceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateAppServiceForPartnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.Extend) {
		query["Extend"] = request.Extend
	}

	if !dara.IsNil(request.OperateEvent) {
		query["OperateEvent"] = request.OperateEvent
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAppServiceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴操作应用服务
//
// @param request - OperateAppServiceForPartnerRequest
//
// @return OperateAppServiceForPartnerResponse
func (client *Client) OperateAppServiceForPartner(request *OperateAppServiceForPartnerRequest) (_result *OperateAppServiceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateAppServiceForPartnerResponse{}
	_body, _err := client.OperateAppServiceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴同步应用实例
//
// @param tmpReq - SyncAppInstanceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartnerWithOptions(tmpReq *SyncAppInstanceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SyncAppInstanceForPartnerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AppInstance) {
		request.AppInstanceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInstance, dara.String("AppInstance"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AppInstanceShrink) {
		query["AppInstance"] = request.AppInstanceShrink
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.Operator) {
		query["Operator"] = request.Operator
	}

	if !dara.IsNil(request.SourceBizId) {
		query["SourceBizId"] = request.SourceBizId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncAppInstanceForPartner"),
		Version:     dara.String("2025-04-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴同步应用实例
//
// @param request - SyncAppInstanceForPartnerRequest
//
// @return SyncAppInstanceForPartnerResponse
func (client *Client) SyncAppInstanceForPartner(request *SyncAppInstanceForPartnerRequest) (_result *SyncAppInstanceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncAppInstanceForPartnerResponse{}
	_body, _err := client.SyncAppInstanceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
