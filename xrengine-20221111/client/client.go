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
		"ap-northeast-1":              dara.String("xrengine-daily.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("xrengine-daily.aliyuncs.com"),
		"ap-south-1":                  dara.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-1":              dara.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-2":              dara.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-3":              dara.String("xrengine-daily.aliyuncs.com"),
		"ap-southeast-5":              dara.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing":                  dara.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("xrengine-daily.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("xrengine-daily.aliyuncs.com"),
		"cn-chengdu":                  dara.String("xrengine-daily.aliyuncs.com"),
		"cn-edge-1":                   dara.String("xrengine-daily.aliyuncs.com"),
		"cn-fujian":                   dara.String("xrengine-daily.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hongkong":                 dara.String("xrengine-daily.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("xrengine-daily.aliyuncs.com"),
		"cn-huhehaote":                dara.String("xrengine-daily.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("xrengine-daily.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("xrengine-daily.aliyuncs.com"),
		"cn-qingdao":                  dara.String("xrengine-daily.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai":                 dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("xrengine-daily.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("xrengine-daily.aliyuncs.com"),
		"cn-wuhan":                    dara.String("xrengine-daily.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("xrengine-daily.aliyuncs.com"),
		"cn-yushanfang":               dara.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("xrengine-daily.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("xrengine-daily.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("xrengine-daily.aliyuncs.com"),
		"eu-central-1":                dara.String("xrengine-daily.aliyuncs.com"),
		"eu-west-1":                   dara.String("xrengine-daily.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("xrengine-daily.aliyuncs.com"),
		"me-east-1":                   dara.String("xrengine-daily.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("xrengine-daily.aliyuncs.com"),
		"us-east-1":                   dara.String("xrengine-daily.aliyuncs.com"),
		"us-west-1":                   dara.String("xrengine-daily.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("xrengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 用户签署协议
//
// @param request - AcceptAgreementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AcceptAgreementResponse
func (client *Client) AcceptAgreementWithOptions(request *AcceptAgreementRequest, runtime *dara.RuntimeOptions) (_result *AcceptAgreementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AcceptAgreement"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AcceptAgreementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户签署协议
//
// @param request - AcceptAgreementRequest
//
// @return AcceptAgreementResponse
func (client *Client) AcceptAgreement(request *AcceptAgreementRequest) (_result *AcceptAgreementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AcceptAgreementResponse{}
	_body, _err := client.AcceptAgreementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加白名单（云账号）
//
// @param request - AddWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddWhitelistResponse
func (client *Client) AddWhitelistWithOptions(request *AddWhitelistRequest, runtime *dara.RuntimeOptions) (_result *AddWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunUid) {
		body["AliyunUid"] = request.AliyunUid
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Remark) {
		body["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddWhitelist"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加白名单（云账号）
//
// @param request - AddWhitelistRequest
//
// @return AddWhitelistResponse
func (client *Client) AddWhitelist(request *AddWhitelistRequest) (_result *AddWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddWhitelistResponse{}
	_body, _err := client.AddWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 申请ai试衣服务
//
// @param request - ApplyForTryOnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyForTryOnResponse
func (client *Client) ApplyForTryOnWithOptions(request *ApplyForTryOnRequest, runtime *dara.RuntimeOptions) (_result *ApplyForTryOnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyForTryOn"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyForTryOnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 申请ai试衣服务
//
// @param request - ApplyForTryOnRequest
//
// @return ApplyForTryOnResponse
func (client *Client) ApplyForTryOn(request *ApplyForTryOnRequest) (_result *ApplyForTryOnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApplyForTryOnResponse{}
	_body, _err := client.ApplyForTryOnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AuthUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthUserResponse
func (client *Client) AuthUserWithOptions(request *AuthUserRequest, runtime *dara.RuntimeOptions) (_result *AuthUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AuthUser"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthUserRequest
//
// @return AuthUserResponse
func (client *Client) AuthUser(request *AuthUserRequest) (_result *AuthUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthUserResponse{}
	_body, _err := client.AuthUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - BatchCreateSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateSvcMapBindResponse
func (client *Client) BatchCreateSvcMapBindWithOptions(tmpReq *BatchCreateSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateSvcMapBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MapIds) {
		request.MapIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MapIds, dara.String("MapIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapIdsShrink) {
		body["MapIds"] = request.MapIdsShrink
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateSvcMapBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BatchCreateSvcMapBindRequest
//
// @return BatchCreateSvcMapBindResponse
func (client *Client) BatchCreateSvcMapBind(request *BatchCreateSvcMapBindRequest) (_result *BatchCreateSvcMapBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateSvcMapBindResponse{}
	_body, _err := client.BatchCreateSvcMapBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消关联
//
// @param tmpReq - BatchDeleteSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteSvcMapBindResponse
func (client *Client) BatchDeleteSvcMapBindWithOptions(tmpReq *BatchDeleteSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeleteSvcMapBindShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		body["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteSvcMapBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消关联
//
// @param request - BatchDeleteSvcMapBindRequest
//
// @return BatchDeleteSvcMapBindResponse
func (client *Client) BatchDeleteSvcMapBind(request *BatchDeleteSvcMapBindRequest) (_result *BatchDeleteSvcMapBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteSvcMapBindResponse{}
	_body, _err := client.BatchDeleteSvcMapBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 构建项目
//
// @param request - BuildProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildProjectResponse
func (client *Client) BuildProjectWithOptions(request *BuildProjectRequest, runtime *dara.RuntimeOptions) (_result *BuildProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.VideoName) {
		query["VideoName"] = request.VideoName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 构建项目
//
// @param request - BuildProjectRequest
//
// @return BuildProjectResponse
func (client *Client) BuildProject(request *BuildProjectRequest) (_result *BuildProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BuildProjectResponse{}
	_body, _err := client.BuildProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建，同时创建空白的pai占位
//
// @param request - CreateLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLocationServiceResponse
func (client *Client) CreateLocationServiceWithOptions(request *CreateLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *CreateLocationServiceResponse, _err error) {
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

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Note) {
		body["Note"] = request.Note
	}

	if !dara.IsNil(request.Qps) {
		body["Qps"] = request.Qps
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建，同时创建空白的pai占位
//
// @param request - CreateLocationServiceRequest
//
// @return CreateLocationServiceResponse
func (client *Client) CreateLocationService(request *CreateLocationServiceRequest) (_result *CreateLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLocationServiceResponse{}
	_body, _err := client.CreateLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TryOnParams) {
		request.TryOnParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TryOnParams, dara.String("TryOnParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		body["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.Dependencies) {
		body["Dependencies"] = request.Dependencies
	}

	if !dara.IsNil(request.ExtInfo) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.Gender) {
		body["Gender"] = request.Gender
	}

	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.Intro) {
		body["Intro"] = request.Intro
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapUuid) {
		body["MapUuid"] = request.MapUuid
	}

	if !dara.IsNil(request.Method) {
		body["Method"] = request.Method
	}

	if !dara.IsNil(request.Mode) {
		body["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.TryOnParamsShrink) {
		body["TryOnParams"] = request.TryOnParamsShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.Weight) {
		body["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建项目
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建素材上传policy
//
// @param request - CreateSourcePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSourcePolicyResponse
func (client *Client) CreateSourcePolicyWithOptions(request *CreateSourcePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateSourcePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSourcePolicy"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSourcePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建素材上传policy
//
// @param request - CreateSourcePolicyRequest
//
// @return CreateSourcePolicyResponse
func (client *Client) CreateSourcePolicy(request *CreateSourcePolicyRequest) (_result *CreateSourcePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSourcePolicyResponse{}
	_body, _err := client.CreateSourcePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除项目信息
//
// @param request - DeleteProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectResponse
func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除项目信息
//
// @param request - DeleteProjectRequest
//
// @return DeleteProjectResponse
func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteSourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceFileResponse
func (client *Client) DeleteSourceFileWithOptions(request *DeleteSourceFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSourceFile"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文件
//
// @param request - DeleteSourceFileRequest
//
// @return DeleteSourceFileResponse
func (client *Client) DeleteSourceFile(request *DeleteSourceFileRequest) (_result *DeleteSourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSourceFileResponse{}
	_body, _err := client.DeleteSourceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 1校验部署状态，一期不支持二次部署。相关关联记录里状态智能变更
//
// @param request - DeployLocationTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployLocationTreeResponse
func (client *Client) DeployLocationTreeWithOptions(request *DeployLocationTreeRequest, runtime *dara.RuntimeOptions) (_result *DeployLocationTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ForceUpdate) {
		body["ForceUpdate"] = request.ForceUpdate
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployLocationTree"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployLocationTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 1校验部署状态，一期不支持二次部署。相关关联记录里状态智能变更
//
// @param request - DeployLocationTreeRequest
//
// @return DeployLocationTreeResponse
func (client *Client) DeployLocationTree(request *DeployLocationTreeRequest) (_result *DeployLocationTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployLocationTreeResponse{}
	_body, _err := client.DeployLocationTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查出绑定记录
//
// @param request - FindSvcMapBindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FindSvcMapBindResponse
func (client *Client) FindSvcMapBindWithOptions(request *FindSvcMapBindRequest, runtime *dara.RuntimeOptions) (_result *FindSvcMapBindResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.SvcId) {
		body["SvcId"] = request.SvcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FindSvcMapBind"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FindSvcMapBindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查出绑定记录
//
// @param request - FindSvcMapBindRequest
//
// @return FindSvcMapBindResponse
func (client *Client) FindSvcMapBind(request *FindSvcMapBindRequest) (_result *FindSvcMapBindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FindSvcMapBindResponse{}
	_body, _err := client.FindSvcMapBindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户申请状态
//
// @param request - GetApplyStatusForTryOnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplyStatusForTryOnResponse
func (client *Client) GetApplyStatusForTryOnWithOptions(request *GetApplyStatusForTryOnRequest, runtime *dara.RuntimeOptions) (_result *GetApplyStatusForTryOnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApplyStatusForTryOn"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApplyStatusForTryOnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户申请状态
//
// @param request - GetApplyStatusForTryOnRequest
//
// @return GetApplyStatusForTryOnResponse
func (client *Client) GetApplyStatusForTryOn(request *GetApplyStatusForTryOnRequest) (_result *GetApplyStatusForTryOnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApplyStatusForTryOnResponse{}
	_body, _err := client.GetApplyStatusForTryOnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetArEditCommonMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditCommonMaterialResponse
func (client *Client) GetArEditCommonMaterialWithOptions(request *GetArEditCommonMaterialRequest, runtime *dara.RuntimeOptions) (_result *GetArEditCommonMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditCommonMaterial"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditCommonMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditCommonMaterialRequest
//
// @return GetArEditCommonMaterialResponse
func (client *Client) GetArEditCommonMaterial(request *GetArEditCommonMaterialRequest) (_result *GetArEditCommonMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArEditCommonMaterialResponse{}
	_body, _err := client.GetArEditCommonMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetArEditStsAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditStsAuthResponse
func (client *Client) GetArEditStsAuthWithOptions(request *GetArEditStsAuthRequest, runtime *dara.RuntimeOptions) (_result *GetArEditStsAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditStsAuth"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditStsAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditStsAuthRequest
//
// @return GetArEditStsAuthResponse
func (client *Client) GetArEditStsAuth(request *GetArEditStsAuthRequest) (_result *GetArEditStsAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArEditStsAuthResponse{}
	_body, _err := client.GetArEditStsAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetArEditUgcMaterialRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArEditUgcMaterialResponse
func (client *Client) GetArEditUgcMaterialWithOptions(request *GetArEditUgcMaterialRequest, runtime *dara.RuntimeOptions) (_result *GetArEditUgcMaterialResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetArEditUgcMaterial"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetArEditUgcMaterialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArEditUgcMaterialRequest
//
// @return GetArEditUgcMaterialResponse
func (client *Client) GetArEditUgcMaterial(request *GetArEditUgcMaterialRequest) (_result *GetArEditUgcMaterialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetArEditUgcMaterialResponse{}
	_body, _err := client.GetArEditUgcMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取项目模型详情
//
// @param request - GetProjectDatasetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectDatasetResponse
func (client *Client) GetProjectDatasetWithOptions(request *GetProjectDatasetRequest, runtime *dara.RuntimeOptions) (_result *GetProjectDatasetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectDataset"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取项目模型详情
//
// @param request - GetProjectDatasetRequest
//
// @return GetProjectDatasetResponse
func (client *Client) GetProjectDataset(request *GetProjectDatasetRequest) (_result *GetProjectDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectDatasetResponse{}
	_body, _err := client.GetProjectDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用同步算法
//
// @param tmpReq - InvokeSyncAlgorithmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeSyncAlgorithmResponse
func (client *Client) InvokeSyncAlgorithmWithOptions(tmpReq *InvokeSyncAlgorithmRequest, runtime *dara.RuntimeOptions) (_result *InvokeSyncAlgorithmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &InvokeSyncAlgorithmShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokeSyncAlgorithm"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokeSyncAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用同步算法
//
// @param request - InvokeSyncAlgorithmRequest
//
// @return InvokeSyncAlgorithmResponse
func (client *Client) InvokeSyncAlgorithm(request *InvokeSyncAlgorithmRequest) (_result *InvokeSyncAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InvokeSyncAlgorithmResponse{}
	_body, _err := client.InvokeSyncAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAreaMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAreaMapResponse
func (client *Client) ListAreaMapWithOptions(request *ListAreaMapRequest, runtime *dara.RuntimeOptions) (_result *ListAreaMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAreaMap"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAreaMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAreaMapRequest
//
// @return ListAreaMapResponse
func (client *Client) ListAreaMap(request *ListAreaMapRequest) (_result *ListAreaMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAreaMapResponse{}
	_body, _err := client.ListAreaMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举服饰类型
//
// @param request - ListClothTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClothTypesResponse
func (client *Client) ListClothTypesWithOptions(runtime *dara.RuntimeOptions) (_result *ListClothTypesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListClothTypes"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClothTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举服饰类型
//
// @return ListClothTypesResponse
func (client *Client) ListClothTypes() (_result *ListClothTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClothTypesResponse{}
	_body, _err := client.ListClothTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举服饰
//
// @param tmpReq - ListClothesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClothesResponse
func (client *Client) ListClothesWithOptions(tmpReq *ListClothesRequest, runtime *dara.RuntimeOptions) (_result *ListClothesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListClothesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Categories) {
		request.CategoriesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Categories, dara.String("Categories"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoriesShrink) {
		query["Categories"] = request.CategoriesShrink
	}

	if !dara.IsNil(request.ClothSize) {
		query["ClothSize"] = request.ClothSize
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClothes"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClothesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举服饰
//
// @param request - ListClothesRequest
//
// @return ListClothesResponse
func (client *Client) ListClothes(request *ListClothesRequest) (_result *ListClothesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClothesResponse{}
	_body, _err := client.ListClothesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// hdr文件目录列表
//
// @param request - ListHdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHdrResponse
func (client *Client) ListHdrWithOptions(runtime *dara.RuntimeOptions) (_result *ListHdrResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListHdr"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// hdr文件目录列表
//
// @return ListHdrResponse
func (client *Client) ListHdr() (_result *ListHdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHdrResponse{}
	_body, _err := client.ListHdrWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListLocationPaiImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLocationPaiImageResponse
func (client *Client) ListLocationPaiImageWithOptions(request *ListLocationPaiImageRequest, runtime *dara.RuntimeOptions) (_result *ListLocationPaiImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLocationPaiImage"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLocationPaiImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLocationPaiImageRequest
//
// @return ListLocationPaiImageResponse
func (client *Client) ListLocationPaiImage(request *ListLocationPaiImageRequest) (_result *ListLocationPaiImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLocationPaiImageResponse{}
	_body, _err := client.ListLocationPaiImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLocationServiceResponse
func (client *Client) ListLocationServiceWithOptions(request *ListLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *ListLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.Sort) {
		body["Sort"] = request.Sort
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLocationServiceRequest
//
// @return ListLocationServiceResponse
func (client *Client) ListLocationService(request *ListLocationServiceRequest) (_result *ListLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLocationServiceResponse{}
	_body, _err := client.ListLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出当前用户项目列表
//
// @param request - ListProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectResponse
func (client *Client) ListProjectWithOptions(request *ListProjectRequest, runtime *dara.RuntimeOptions) (_result *ListProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizUsage) {
		body["BizUsage"] = request.BizUsage
	}

	if !dara.IsNil(request.Current) {
		body["Current"] = request.Current
	}

	if !dara.IsNil(request.ExcludedBizUsage) {
		body["ExcludedBizUsage"] = request.ExcludedBizUsage
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Size) {
		body["Size"] = request.Size
	}

	if !dara.IsNil(request.SortField) {
		body["SortField"] = request.SortField
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	if !dara.IsNil(request.WithSource) {
		body["WithSource"] = request.WithSource
	}

	if !dara.IsNil(request.WithUser) {
		body["WithUser"] = request.WithUser
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出当前用户项目列表
//
// @param request - ListProjectRequest
//
// @return ListProjectResponse
func (client *Client) ListProject(request *ListProjectRequest) (_result *ListProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectResponse{}
	_body, _err := client.ListProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据项目Id查出依赖当前项目的项目
//
// @param request - ListProjectsByDependencyIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsByDependencyIdResponse
func (client *Client) ListProjectsByDependencyIdWithOptions(request *ListProjectsByDependencyIdRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsByDependencyIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DependencyProjectId) {
		body["DependencyProjectId"] = request.DependencyProjectId
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapUuid) {
		body["MapUuid"] = request.MapUuid
	}

	if !dara.IsNil(request.WithDataset) {
		body["WithDataset"] = request.WithDataset
	}

	if !dara.IsNil(request.WithSource) {
		body["WithSource"] = request.WithSource
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectsByDependencyId"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsByDependencyIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据项目Id查出依赖当前项目的项目
//
// @param request - ListProjectsByDependencyIdRequest
//
// @return ListProjectsByDependencyIdResponse
func (client *Client) ListProjectsByDependencyId(request *ListProjectsByDependencyIdRequest) (_result *ListProjectsByDependencyIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsByDependencyIdResponse{}
	_body, _err := client.ListProjectsByDependencyIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举项目中上传的文件列表
//
// @param request - ListSourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSourceFileResponse
func (client *Client) ListSourceFileWithOptions(request *ListSourceFileRequest, runtime *dara.RuntimeOptions) (_result *ListSourceFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSourceFile"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举项目中上传的文件列表
//
// @param request - ListSourceFileRequest
//
// @return ListSourceFileResponse
func (client *Client) ListSourceFile(request *ListSourceFileRequest) (_result *ListSourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSourceFileResponse{}
	_body, _err := client.ListSourceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举流水线
//
// @param request - ListWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowResponse
func (client *Client) ListWorkflowWithOptions(runtime *dara.RuntimeOptions) (_result *ListWorkflowResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflow"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举流水线
//
// @return ListWorkflowResponse
func (client *Client) ListWorkflow() (_result *ListWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowResponse{}
	_body, _err := client.ListWorkflowWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - LoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginResponse
func (client *Client) LoginWithOptions(request *LoginRequest, runtime *dara.RuntimeOptions) (_result *LoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EmpId) {
		body["EmpId"] = request.EmpId
	}

	if !dara.IsNil(request.EmpName) {
		body["EmpName"] = request.EmpName
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Login"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - LoginRequest
//
// @return LoginResponse
func (client *Client) Login(request *LoginRequest) (_result *LoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LoginResponse{}
	_body, _err := client.LoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 无线端APP登陆
//
// @param request - LoginAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginAppResponse
func (client *Client) LoginAppWithOptions(request *LoginAppRequest, runtime *dara.RuntimeOptions) (_result *LoginAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EmpId) {
		body["EmpId"] = request.EmpId
	}

	if !dara.IsNil(request.EmpName) {
		body["EmpName"] = request.EmpName
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoginApp"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoginAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 无线端APP登陆
//
// @param request - LoginAppRequest
//
// @return LoginAppResponse
func (client *Client) LoginApp(request *LoginAppRequest) (_result *LoginAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LoginAppResponse{}
	_body, _err := client.LoginAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PublishArEditProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishArEditProjectResponse
func (client *Client) PublishArEditProjectWithOptions(request *PublishArEditProjectRequest, runtime *dara.RuntimeOptions) (_result *PublishArEditProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishArEditProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishArEditProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PublishArEditProjectRequest
//
// @return PublishArEditProjectResponse
func (client *Client) PublishArEditProject(request *PublishArEditProjectRequest) (_result *PublishArEditProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishArEditProjectResponse{}
	_body, _err := client.PublishArEditProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目发布
//
// @param request - PublishProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishProjectResponse
func (client *Client) PublishProjectWithOptions(request *PublishProjectRequest, runtime *dara.RuntimeOptions) (_result *PublishProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目发布
//
// @param request - PublishProjectRequest
//
// @return PublishProjectResponse
func (client *Client) PublishProject(request *PublishProjectRequest) (_result *PublishProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishProjectResponse{}
	_body, _err := client.PublishProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAreaMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAreaMapResponse
func (client *Client) QueryAreaMapWithOptions(request *QueryAreaMapRequest, runtime *dara.RuntimeOptions) (_result *QueryAreaMapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAreaMap"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAreaMapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAreaMapRequest
//
// @return QueryAreaMapResponse
func (client *Client) QueryAreaMap(request *QueryAreaMapRequest) (_result *QueryAreaMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAreaMapResponse{}
	_body, _err := client.QueryAreaMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查找项目构建失败时的信息
//
// @param request - QueryBuildBreakpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBuildBreakpointResponse
func (client *Client) QueryBuildBreakpointWithOptions(request *QueryBuildBreakpointRequest, runtime *dara.RuntimeOptions) (_result *QueryBuildBreakpointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryBuildBreakpoint"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBuildBreakpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找项目构建失败时的信息
//
// @param request - QueryBuildBreakpointRequest
//
// @return QueryBuildBreakpointResponse
func (client *Client) QueryBuildBreakpoint(request *QueryBuildBreakpointRequest) (_result *QueryBuildBreakpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBuildBreakpointResponse{}
	_body, _err := client.QueryBuildBreakpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryLocationServiceResponse
func (client *Client) QueryLocationServiceWithOptions(request *QueryLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *QueryLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryLocationServiceRequest
//
// @return QueryLocationServiceResponse
func (client *Client) QueryLocationService(request *QueryLocationServiceRequest) (_result *QueryLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryLocationServiceResponse{}
	_body, _err := client.QueryLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目构建详情
//
// @param request - QueryProjectBuildDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectBuildDetailResponse
func (client *Client) QueryProjectBuildDetailWithOptions(request *QueryProjectBuildDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryProjectBuildDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProjectBuildDetail"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectBuildDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目构建详情
//
// @param request - QueryProjectBuildDetailRequest
//
// @return QueryProjectBuildDetailResponse
func (client *Client) QueryProjectBuildDetail(request *QueryProjectBuildDetailRequest) (_result *QueryProjectBuildDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryProjectBuildDetailResponse{}
	_body, _err := client.QueryProjectBuildDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询项目详情
//
// @param request - QueryProjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryProjectDetailResponse
func (client *Client) QueryProjectDetailWithOptions(request *QueryProjectDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryProjectDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.OssKey) {
		query["OssKey"] = request.OssKey
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryProjectDetail"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询项目详情
//
// @param request - QueryProjectDetailRequest
//
// @return QueryProjectDetailResponse
func (client *Client) QueryProjectDetail(request *QueryProjectDetailRequest) (_result *QueryProjectDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryProjectDetailResponse{}
	_body, _err := client.QueryProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 该注册接口只用于oneconsole页面的注册
//
// @param request - RegisterUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUserResponse
func (client *Client) RegisterUserWithOptions(request *RegisterUserRequest, runtime *dara.RuntimeOptions) (_result *RegisterUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterUser"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 该注册接口只用于oneconsole页面的注册
//
// @param request - RegisterUserRequest
//
// @return RegisterUserResponse
func (client *Client) RegisterUser(request *RegisterUserRequest) (_result *RegisterUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterUserResponse{}
	_body, _err := client.RegisterUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ResumeLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeLocationServiceResponse
func (client *Client) ResumeLocationServiceWithOptions(request *ResumeLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *ResumeLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ResumeLocationServiceRequest
//
// @return ResumeLocationServiceResponse
func (client *Client) ResumeLocationService(request *ResumeLocationServiceRequest) (_result *ResumeLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeLocationServiceResponse{}
	_body, _err := client.ResumeLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SaveArEditProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveArEditProjectResponse
func (client *Client) SaveArEditProjectWithOptions(request *SaveArEditProjectRequest, runtime *dara.RuntimeOptions) (_result *SaveArEditProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.MapId) {
		body["MapId"] = request.MapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveArEditProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveArEditProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SaveArEditProjectRequest
//
// @return SaveArEditProjectResponse
func (client *Client) SaveArEditProject(request *SaveArEditProjectRequest) (_result *SaveArEditProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveArEditProjectResponse{}
	_body, _err := client.SaveArEditProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param request - SaveSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSourceResponse
func (client *Client) SaveSourceWithOptions(request *SaveSourceRequest, runtime *dara.RuntimeOptions) (_result *SaveSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeStatus) {
		query["ChangeStatus"] = request.ChangeStatus
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.NeedCheck) {
		query["NeedCheck"] = request.NeedCheck
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveSource"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存素材
//
// @param request - SaveSourceRequest
//
// @return SaveSourceResponse
func (client *Client) SaveSource(request *SaveSourceRequest) (_result *SaveSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveSourceResponse{}
	_body, _err := client.SaveSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SuspendLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendLocationServiceResponse
func (client *Client) SuspendLocationServiceWithOptions(request *SuspendLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *SuspendLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SuspendLocationServiceRequest
//
// @return SuspendLocationServiceResponse
func (client *Client) SuspendLocationService(request *SuspendLocationServiceRequest) (_result *SuspendLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendLocationServiceResponse{}
	_body, _err := client.SuspendLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 项目(发布状态改回隐私状态)
//
// @param request - UnPublishProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnPublishProjectResponse
func (client *Client) UnPublishProjectWithOptions(request *UnPublishProjectRequest, runtime *dara.RuntimeOptions) (_result *UnPublishProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnPublishProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnPublishProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 项目(发布状态改回隐私状态)
//
// @param request - UnPublishProjectRequest
//
// @return UnPublishProjectResponse
func (client *Client) UnPublishProject(request *UnPublishProjectRequest) (_result *UnPublishProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnPublishProjectResponse{}
	_body, _err := client.UnPublishProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 如果不包含treeConfig则只是简单更新
//
// @param request - UpdateLocationServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationServiceResponse
func (client *Client) UpdateLocationServiceWithOptions(request *UpdateLocationServiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.Note) {
		body["Note"] = request.Note
	}

	if !dara.IsNil(request.Qps) {
		body["Qps"] = request.Qps
	}

	if !dara.IsNil(request.SvcName) {
		body["SvcName"] = request.SvcName
	}

	if !dara.IsNil(request.SvcState) {
		body["SvcState"] = request.SvcState
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocationService"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocationServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 如果不包含treeConfig则只是简单更新
//
// @param request - UpdateLocationServiceRequest
//
// @return UpdateLocationServiceResponse
func (client *Client) UpdateLocationService(request *UpdateLocationServiceRequest) (_result *UpdateLocationServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLocationServiceResponse{}
	_body, _err := client.UpdateLocationServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂存接口比较复杂，单独拎出来
//
// @param request - UpdateLocationTreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLocationTreeResponse
func (client *Client) UpdateLocationTreeWithOptions(request *UpdateLocationTreeRequest, runtime *dara.RuntimeOptions) (_result *UpdateLocationTreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.JwtToken) {
		body["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.TreeConfig) {
		body["TreeConfig"] = request.TreeConfig
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLocationTree"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLocationTreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂存接口比较复杂，单独拎出来
//
// @param request - UpdateLocationTreeRequest
//
// @return UpdateLocationTreeResponse
func (client *Client) UpdateLocationTree(request *UpdateLocationTreeRequest) (_result *UpdateLocationTreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLocationTreeResponse{}
	_body, _err := client.UpdateLocationTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新项目信息
//
// @param tmpReq - UpdateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProjectResponse
func (client *Client) UpdateProjectWithOptions(tmpReq *UpdateProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ext) {
		request.ExtShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ext, dara.String("Ext"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoBuild) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !dara.IsNil(request.ExtShrink) {
		query["Ext"] = request.ExtShrink
	}

	if !dara.IsNil(request.Intro) {
		query["Intro"] = request.Intro
	}

	if !dara.IsNil(request.JwtToken) {
		query["JwtToken"] = request.JwtToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProject"),
		Version:     dara.String("2022-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新项目信息
//
// @param request - UpdateProjectRequest
//
// @return UpdateProjectResponse
func (client *Client) UpdateProject(request *UpdateProjectRequest) (_result *UpdateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProjectResponse{}
	_body, _err := client.UpdateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
