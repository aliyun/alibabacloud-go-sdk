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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dypns"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 配置认证方案APP信息
//
// @param request - ConfigVerifySceneAppInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigVerifySceneAppInfoResponse
func (client *Client) ConfigVerifySceneAppInfoWithOptions(request *ConfigVerifySceneAppInfoRequest, runtime *dara.RuntimeOptions) (_result *ConfigVerifySceneAppInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CM) {
		query["CM"] = request.CM
	}

	if !dara.IsNil(request.CT) {
		query["CT"] = request.CT
	}

	if !dara.IsNil(request.CU) {
		query["CU"] = request.CU
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IpWhitelist) {
		query["IpWhitelist"] = request.IpWhitelist
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

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConfigVerifySceneAppInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConfigVerifySceneAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置认证方案APP信息
//
// @param request - ConfigVerifySceneAppInfoRequest
//
// @return ConfigVerifySceneAppInfoResponse
func (client *Client) ConfigVerifySceneAppInfo(request *ConfigVerifySceneAppInfoRequest) (_result *ConfigVerifySceneAppInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ConfigVerifySceneAppInfoResponse{}
	_body, _err := client.ConfigVerifySceneAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CopySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopySchemeResponse
func (client *Client) CopySchemeWithOptions(request *CopySchemeRequest, runtime *dara.RuntimeOptions) (_result *CopySchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeId) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CopySchemeRequest
//
// @return CopySchemeResponse
func (client *Client) CopyScheme(request *CopySchemeRequest) (_result *CopySchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopySchemeResponse{}
	_body, _err := client.CopySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 复制号码认证方案信息
//
// @param request - CopyVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyVerifySchemeResponse
func (client *Client) CopyVerifySchemeWithOptions(request *CopyVerifySchemeRequest, runtime *dara.RuntimeOptions) (_result *CopyVerifySchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CmApiCode) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !dara.IsNil(request.CmNewFlag) {
		query["CmNewFlag"] = request.CmNewFlag
	}

	if !dara.IsNil(request.CtApiCode) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !dara.IsNil(request.CtNewFlag) {
		query["CtNewFlag"] = request.CtNewFlag
	}

	if !dara.IsNil(request.CuApiCode) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !dara.IsNil(request.CuNewFlag) {
		query["CuNewFlag"] = request.CuNewFlag
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.ModelSceneCode) {
		query["ModelSceneCode"] = request.ModelSceneCode
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CopyVerifyScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CopyVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 复制号码认证方案信息
//
// @param request - CopyVerifySchemeRequest
//
// @return CopyVerifySchemeResponse
func (client *Client) CopyVerifyScheme(request *CopyVerifySchemeRequest) (_result *CopyVerifySchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CopyVerifySchemeResponse{}
	_body, _err := client.CopyVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0http回执测试
//
// @param request - CreateDypnsSmsVerifyCallBackTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDypnsSmsVerifyCallBackTestResponse
func (client *Client) CreateDypnsSmsVerifyCallBackTestWithOptions(request *CreateDypnsSmsVerifyCallBackTestRequest, runtime *dara.RuntimeOptions) (_result *CreateDypnsSmsVerifyCallBackTestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Method) {
		query["Method"] = request.Method
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

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDypnsSmsVerifyCallBackTest"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDypnsSmsVerifyCallBackTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0http回执测试
//
// @param request - CreateDypnsSmsVerifyCallBackTestRequest
//
// @return CreateDypnsSmsVerifyCallBackTestResponse
func (client *Client) CreateDypnsSmsVerifyCallBackTest(request *CreateDypnsSmsVerifyCallBackTestRequest) (_result *CreateDypnsSmsVerifyCallBackTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDypnsSmsVerifyCallBackTestResponse{}
	_body, _err := client.CreateDypnsSmsVerifyCallBackTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0http回执url
//
// @param request - CreateDypnsSmsVerifyMessageCallBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDypnsSmsVerifyMessageCallBackResponse
func (client *Client) CreateDypnsSmsVerifyMessageCallBackWithOptions(request *CreateDypnsSmsVerifyMessageCallBackRequest, runtime *dara.RuntimeOptions) (_result *CreateDypnsSmsVerifyMessageCallBackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
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

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDypnsSmsVerifyMessageCallBack"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDypnsSmsVerifyMessageCallBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0http回执url
//
// @param request - CreateDypnsSmsVerifyMessageCallBackRequest
//
// @return CreateDypnsSmsVerifyMessageCallBackResponse
func (client *Client) CreateDypnsSmsVerifyMessageCallBack(request *CreateDypnsSmsVerifyMessageCallBackRequest) (_result *CreateDypnsSmsVerifyMessageCallBackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDypnsSmsVerifyMessageCallBackResponse{}
	_body, _err := client.CreateDypnsSmsVerifyMessageCallBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0mns回执队列
//
// @param request - CreateDypnsSmsVerifyMessageQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDypnsSmsVerifyMessageQueueResponse
func (client *Client) CreateDypnsSmsVerifyMessageQueueWithOptions(request *CreateDypnsSmsVerifyMessageQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateDypnsSmsVerifyMessageQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryQueueTypes) {
		query["QueryQueueTypes"] = request.QueryQueueTypes
	}

	if !dara.IsNil(request.QueueType) {
		query["QueueType"] = request.QueueType
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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
		Action:      dara.String("CreateDypnsSmsVerifyMessageQueue"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDypnsSmsVerifyMessageQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短信认证2.0mns回执队列
//
// @param request - CreateDypnsSmsVerifyMessageQueueRequest
//
// @return CreateDypnsSmsVerifyMessageQueueResponse
func (client *Client) CreateDypnsSmsVerifyMessageQueue(request *CreateDypnsSmsVerifyMessageQueueRequest) (_result *CreateDypnsSmsVerifyMessageQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDypnsSmsVerifyMessageQueueResponse{}
	_body, _err := client.CreateDypnsSmsVerifyMessageQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFileByBizRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileByBizResponse
func (client *Client) CreateFileByBizWithOptions(request *CreateFileByBizRequest, runtime *dara.RuntimeOptions) (_result *CreateFileByBizResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SaveOssFileName) {
		query["SaveOssFileName"] = request.SaveOssFileName
	}

	if !dara.IsNil(request.UserViewFileName) {
		query["UserViewFileName"] = request.UserViewFileName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileByBiz"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileByBizResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFileByBizRequest
//
// @return CreateFileByBizResponse
func (client *Client) CreateFileByBiz(request *CreateFileByBizRequest) (_result *CreateFileByBizResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileByBizResponse{}
	_body, _err := client.CreateFileByBizWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateGateVerifyExportLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGateVerifyExportLogResponse
func (client *Client) CreateGateVerifyExportLogWithOptions(request *CreateGateVerifyExportLogRequest, runtime *dara.RuntimeOptions) (_result *CreateGateVerifyExportLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryMonth) {
		query["QueryMonth"] = request.QueryMonth
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGateVerifyExportLog"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGateVerifyExportLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateGateVerifyExportLogRequest
//
// @return CreateGateVerifyExportLogResponse
func (client *Client) CreateGateVerifyExportLog(request *CreateGateVerifyExportLogRequest) (_result *CreateGateVerifyExportLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGateVerifyExportLogResponse{}
	_body, _err := client.CreateGateVerifyExportLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建图形认证方案号
//
// @param request - CreateGraphicAuthSchemeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGraphicAuthSchemeConfigResponse
func (client *Client) CreateGraphicAuthSchemeConfigWithOptions(request *CreateGraphicAuthSchemeConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateGraphicAuthSchemeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGraphicAuthSchemeConfig"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGraphicAuthSchemeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图形认证方案号
//
// @param request - CreateGraphicAuthSchemeConfigRequest
//
// @return CreateGraphicAuthSchemeConfigResponse
func (client *Client) CreateGraphicAuthSchemeConfig(request *CreateGraphicAuthSchemeConfigRequest) (_result *CreateGraphicAuthSchemeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateGraphicAuthSchemeConfigResponse{}
	_body, _err := client.CreateGraphicAuthSchemeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductResponse
func (client *Client) CreateProductWithOptions(request *CreateProductRequest, runtime *dara.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ProdId) {
		query["ProdId"] = request.ProdId
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
		Action:      dara.String("CreateProduct"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品
//
// @param request - CreateProductRequest
//
// @return CreateProductResponse
func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 融合认证-新增方案
//
// @param request - CreateSchemeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSchemeConfigResponse
func (client *Client) CreateSchemeConfigWithOptions(request *CreateSchemeConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateSchemeConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPackageName) {
		query["AndroidPackageName"] = request.AndroidPackageName
	}

	if !dara.IsNil(request.AndroidPackageSign) {
		query["AndroidPackageSign"] = request.AndroidPackageSign
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.H5Origin) {
		query["H5Origin"] = request.H5Origin
	}

	if !dara.IsNil(request.H5Url) {
		query["H5Url"] = request.H5Url
	}

	if !dara.IsNil(request.IosBundleId) {
		query["IosBundleId"] = request.IosBundleId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Platform) {
		query["Platform"] = request.Platform
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSchemeConfig"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSchemeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 融合认证-新增方案
//
// @param request - CreateSchemeConfigRequest
//
// @return CreateSchemeConfigResponse
func (client *Client) CreateSchemeConfig(request *CreateSchemeConfigRequest) (_result *CreateSchemeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSchemeConfigResponse{}
	_body, _err := client.CreateSchemeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSignWithOptions(request *CreateSmsSignRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppUrl) {
		query["AppUrl"] = request.AppUrl
	}

	if !dara.IsNil(request.BusinessLicenseCert) {
		query["BusinessLicenseCert"] = request.BusinessLicenseCert
	}

	if !dara.IsNil(request.CertType) {
		query["CertType"] = request.CertType
	}

	if !dara.IsNil(request.ExtendMessage) {
		query["ExtendMessage"] = request.ExtendMessage
	}

	if !dara.IsNil(request.OldSmsSignName) {
		query["OldSmsSignName"] = request.OldSmsSignName
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OrganizationCodeCert) {
		query["OrganizationCodeCert"] = request.OrganizationCodeCert
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.SmsSignRemark) {
		query["SmsSignRemark"] = request.SmsSignRemark
	}

	if !dara.IsNil(request.SmsSignSource) {
		query["SmsSignSource"] = request.SmsSignSource
	}

	if !dara.IsNil(request.TaxRegistrationCert) {
		query["TaxRegistrationCert"] = request.TaxRegistrationCert
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsSign"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSmsSignRequest
//
// @return CreateSmsSignResponse
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (_result *CreateSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsSignResponse{}
	_body, _err := client.CreateSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplateWithOptions(request *CreateSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizUrl) {
		query["BizUrl"] = request.BizUrl
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SmsTemplateContent) {
		query["SmsTemplateContent"] = request.SmsTemplateContent
	}

	if !dara.IsNil(request.SmsTemplateName) {
		query["SmsTemplateName"] = request.SmsTemplateName
	}

	if !dara.IsNil(request.SmsTemplateRule) {
		query["SmsTemplateRule"] = request.SmsTemplateRule
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsTemplate"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateSmsTemplateRequest
//
// @return CreateSmsTemplateResponse
func (client *Client) CreateSmsTemplate(request *CreateSmsTemplateRequest) (_result *CreateSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsTemplateResponse{}
	_body, _err := client.CreateSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短信认证用量统计记录导出任务
//
// @param request - CreateSmsVerifyStatisticRecordsExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmsVerifyStatisticRecordsExportTaskResponse
func (client *Client) CreateSmsVerifyStatisticRecordsExportTaskWithOptions(request *CreateSmsVerifyStatisticRecordsExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSmsVerifyStatisticRecordsExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromDate) {
		query["FromDate"] = request.FromDate
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

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	if !dara.IsNil(request.ToDate) {
		query["ToDate"] = request.ToDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmsVerifyStatisticRecordsExportTask"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmsVerifyStatisticRecordsExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短信认证用量统计记录导出任务
//
// @param request - CreateSmsVerifyStatisticRecordsExportTaskRequest
//
// @return CreateSmsVerifyStatisticRecordsExportTaskResponse
func (client *Client) CreateSmsVerifyStatisticRecordsExportTask(request *CreateSmsVerifyStatisticRecordsExportTaskRequest) (_result *CreateSmsVerifyStatisticRecordsExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSmsVerifyStatisticRecordsExportTaskResponse{}
	_body, _err := client.CreateSmsVerifyStatisticRecordsExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建认证方案
//
// @param request - CreateVerifySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVerifySchemeResponse
func (client *Client) CreateVerifySchemeWithOptions(request *CreateVerifySchemeRequest, runtime *dara.RuntimeOptions) (_result *CreateVerifySchemeResponse, _err error) {
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

	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.BundleId) {
		query["BundleId"] = request.BundleId
	}

	if !dara.IsNil(request.CmApiCode) {
		query["CmApiCode"] = request.CmApiCode
	}

	if !dara.IsNil(request.CtApiCode) {
		query["CtApiCode"] = request.CtApiCode
	}

	if !dara.IsNil(request.CuApiCode) {
		query["CuApiCode"] = request.CuApiCode
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.HmAppIdentifier) {
		query["HmAppIdentifier"] = request.HmAppIdentifier
	}

	if !dara.IsNil(request.HmPackageName) {
		query["HmPackageName"] = request.HmPackageName
	}

	if !dara.IsNil(request.HmSignName) {
		query["HmSignName"] = request.HmSignName
	}

	if !dara.IsNil(request.IpWhiteList) {
		query["IpWhiteList"] = request.IpWhiteList
	}

	if !dara.IsNil(request.Origin) {
		query["Origin"] = request.Origin
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PackName) {
		query["PackName"] = request.PackName
	}

	if !dara.IsNil(request.PackSign) {
		query["PackSign"] = request.PackSign
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneType) {
		query["SceneType"] = request.SceneType
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVerifyScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建认证方案
//
// @param request - CreateVerifySchemeRequest
//
// @return CreateVerifySchemeResponse
func (client *Client) CreateVerifyScheme(request *CreateVerifySchemeRequest) (_result *CreateVerifySchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVerifySchemeResponse{}
	_body, _err := client.CreateVerifySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建短信认证导出任务
//
// @param request - CreatetVerifySmsExportTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatetVerifySmsExportTaskResponse
func (client *Client) CreatetVerifySmsExportTaskWithOptions(request *CreatetVerifySmsExportTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatetVerifySmsExportTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeName) {
		query["SchemeName"] = request.SchemeName
	}

	if !dara.IsNil(request.SendStatus) {
		query["SendStatus"] = request.SendStatus
	}

	if !dara.IsNil(request.SignName) {
		query["SignName"] = request.SignName
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatetVerifySmsExportTask"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatetVerifySmsExportTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建短信认证导出任务
//
// @param request - CreatetVerifySmsExportTaskRequest
//
// @return CreatetVerifySmsExportTaskResponse
func (client *Client) CreatetVerifySmsExportTask(request *CreatetVerifySmsExportTaskRequest) (_result *CreatetVerifySmsExportTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatetVerifySmsExportTaskResponse{}
	_body, _err := client.CreatetVerifySmsExportTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSceneResponse
func (client *Client) DeleteSceneWithOptions(request *DeleteSceneRequest, runtime *dara.RuntimeOptions) (_result *DeleteSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScene"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSceneRequest
//
// @return DeleteSceneResponse
func (client *Client) DeleteScene(request *DeleteSceneRequest) (_result *DeleteSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSceneResponse{}
	_body, _err := client.DeleteSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSchemeResponse
func (client *Client) DeleteSchemeWithOptions(request *DeleteSchemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeId) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteSchemeRequest
//
// @return DeleteSchemeResponse
func (client *Client) DeleteScheme(request *DeleteSchemeRequest) (_result *DeleteSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSchemeResponse{}
	_body, _err := client.DeleteSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDownLoadFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDownLoadFileUrlResponse
func (client *Client) GetDownLoadFileUrlWithOptions(request *GetDownLoadFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetDownLoadFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcctId) {
		query["AcctId"] = request.AcctId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDownLoadFileUrl"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDownLoadFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDownLoadFileUrlRequest
//
// @return GetDownLoadFileUrlResponse
func (client *Client) GetDownLoadFileUrl(request *GetDownLoadFileUrlRequest) (_result *GetDownLoadFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDownLoadFileUrlResponse{}
	_body, _err := client.GetDownLoadFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnterpriseInfoResponse
func (client *Client) GetEnterpriseInfoWithOptions(request *GetEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnterpriseInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetEnterpriseInfoRequest
//
// @return GetEnterpriseInfoResponse
func (client *Client) GetEnterpriseInfo(request *GetEnterpriseInfoRequest) (_result *GetEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEnterpriseInfoResponse{}
	_body, _err := client.GetEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetOpenSubProductStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpenSubProductStatusResponse
func (client *Client) GetOpenSubProductStatusWithOptions(request *GetOpenSubProductStatusRequest, runtime *dara.RuntimeOptions) (_result *GetOpenSubProductStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
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
		Action:      dara.String("GetOpenSubProductStatus"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpenSubProductStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOpenSubProductStatusRequest
//
// @return GetOpenSubProductStatusResponse
func (client *Client) GetOpenSubProductStatus(request *GetOpenSubProductStatusRequest) (_result *GetOpenSubProductStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpenSubProductStatusResponse{}
	_body, _err := client.GetOpenSubProductStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetOssUploadForEnterpriseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssUploadForEnterpriseResponse
func (client *Client) GetOssUploadForEnterpriseWithOptions(request *GetOssUploadForEnterpriseRequest, runtime *dara.RuntimeOptions) (_result *GetOssUploadForEnterpriseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("GetOssUploadForEnterprise"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssUploadForEnterpriseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOssUploadForEnterpriseRequest
//
// @return GetOssUploadForEnterpriseResponse
func (client *Client) GetOssUploadForEnterprise(request *GetOssUploadForEnterpriseRequest) (_result *GetOssUploadForEnterpriseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssUploadForEnterpriseResponse{}
	_body, _err := client.GetOssUploadForEnterpriseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetOssUploadSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssUploadSignResponse
func (client *Client) GetOssUploadSignWithOptions(request *GetOssUploadSignRequest, runtime *dara.RuntimeOptions) (_result *GetOssUploadSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssUploadSign"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssUploadSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetOssUploadSignRequest
//
// @return GetOssUploadSignResponse
func (client *Client) GetOssUploadSign(request *GetOssUploadSignRequest) (_result *GetOssUploadSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssUploadSignResponse{}
	_body, _err := client.GetOssUploadSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSchemeResponse
func (client *Client) GetSchemeWithOptions(request *GetSchemeRequest, runtime *dara.RuntimeOptions) (_result *GetSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SchemeId) {
		query["SchemeId"] = request.SchemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSchemeRequest
//
// @return GetSchemeResponse
func (client *Client) GetScheme(request *GetSchemeRequest) (_result *GetSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSchemeResponse{}
	_body, _err := client.GetSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSmsCodeLimitConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsCodeLimitConfigResponse
func (client *Client) GetSmsCodeLimitConfigWithOptions(request *GetSmsCodeLimitConfigRequest, runtime *dara.RuntimeOptions) (_result *GetSmsCodeLimitConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("GetSmsCodeLimitConfig"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsCodeLimitConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSmsCodeLimitConfigRequest
//
// @return GetSmsCodeLimitConfigResponse
func (client *Client) GetSmsCodeLimitConfig(request *GetSmsCodeLimitConfigRequest) (_result *GetSmsCodeLimitConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmsCodeLimitConfigResponse{}
	_body, _err := client.GetSmsCodeLimitConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSignWithOptions(request *GetSmsSignRequest, runtime *dara.RuntimeOptions) (_result *GetSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
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

	if !dara.IsNil(request.SmsSignName) {
		query["SmsSignName"] = request.SmsSignName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmsSign"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetSmsSignRequest
//
// @return GetSmsSignResponse
func (client *Client) GetSmsSign(request *GetSmsSignRequest) (_result *GetSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmsSignResponse{}
	_body, _err := client.GetSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAuditPassEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuditPassEnterpriseInfoResponse
func (client *Client) ListAuditPassEnterpriseInfoWithOptions(request *ListAuditPassEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListAuditPassEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("ListAuditPassEnterpriseInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuditPassEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAuditPassEnterpriseInfoRequest
//
// @return ListAuditPassEnterpriseInfoResponse
func (client *Client) ListAuditPassEnterpriseInfo(request *ListAuditPassEnterpriseInfoRequest) (_result *ListAuditPassEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuditPassEnterpriseInfoResponse{}
	_body, _err := client.ListAuditPassEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListCertificateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCertificateResponse
func (client *Client) ListCertificateWithOptions(request *ListCertificateRequest, runtime *dara.RuntimeOptions) (_result *ListCertificateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.PhoneNo) {
		query["PhoneNo"] = request.PhoneNo
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCertificate"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListCertificateRequest
//
// @return ListCertificateResponse
func (client *Client) ListCertificate(request *ListCertificateRequest) (_result *ListCertificateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCertificateResponse{}
	_body, _err := client.ListCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseInfoResponse
func (client *Client) ListEnterpriseInfoWithOptions(request *ListEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListEnterpriseInfoRequest
//
// @return ListEnterpriseInfoResponse
func (client *Client) ListEnterpriseInfo(request *ListEnterpriseInfoRequest) (_result *ListEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnterpriseInfoResponse{}
	_body, _err := client.ListEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchemeResponse
func (client *Client) ListSchemeWithOptions(request *ListSchemeRequest, runtime *dara.RuntimeOptions) (_result *ListSchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScheme"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSchemeRequest
//
// @return ListSchemeResponse
func (client *Client) ListScheme(request *ListSchemeRequest) (_result *ListSchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchemeResponse{}
	_body, _err := client.ListSchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSmsSignRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmsSignResponse
func (client *Client) ListSmsSignWithOptions(request *ListSmsSignRequest, runtime *dara.RuntimeOptions) (_result *ListSmsSignResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.QuerySmsSign) {
		query["QuerySmsSign"] = request.QuerySmsSign
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
		Action:      dara.String("ListSmsSign"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmsSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSmsSignRequest
//
// @return ListSmsSignResponse
func (client *Client) ListSmsSign(request *ListSmsSignRequest) (_result *ListSmsSignResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmsSignResponse{}
	_body, _err := client.ListSmsSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListSmsTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmsTemplateResponse
func (client *Client) ListSmsTemplateWithOptions(request *ListSmsTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListSmsTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerId) {
		query["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.QueryTemplate) {
		query["QueryTemplate"] = request.QueryTemplate
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
		Action:      dara.String("ListSmsTemplate"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmsTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListSmsTemplateRequest
//
// @return ListSmsTemplateResponse
func (client *Client) ListSmsTemplate(request *ListSmsTemplateRequest) (_result *ListSmsTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmsTemplateResponse{}
	_body, _err := client.ListSmsTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenCommonProductRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenCommonProductResponse
func (client *Client) OpenCommonProductWithOptions(request *OpenCommonProductRequest, runtime *dara.RuntimeOptions) (_result *OpenCommonProductResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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
		Action:      dara.String("OpenCommonProduct"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenCommonProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - OpenCommonProductRequest
//
// @return OpenCommonProductResponse
func (client *Client) OpenCommonProduct(request *OpenCommonProductRequest) (_result *OpenCommonProductResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenCommonProductResponse{}
	_body, _err := client.OpenCommonProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCommonBillingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommonBillingResponse
func (client *Client) QueryCommonBillingWithOptions(request *QueryCommonBillingRequest, runtime *dara.RuntimeOptions) (_result *QueryCommonBillingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Month) {
		query["Month"] = request.Month
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
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
		Action:      dara.String("QueryCommonBilling"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommonBillingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCommonBillingRequest
//
// @return QueryCommonBillingResponse
func (client *Client) QueryCommonBilling(request *QueryCommonBillingRequest) (_result *QueryCommonBillingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCommonBillingResponse{}
	_body, _err := client.QueryCommonBillingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询客户信息
//
// @param request - QueryCommonCustInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommonCustInfoResponse
func (client *Client) QueryCommonCustInfoWithOptions(request *QueryCommonCustInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryCommonCustInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryCommonCustInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommonCustInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询客户信息
//
// @param request - QueryCommonCustInfoRequest
//
// @return QueryCommonCustInfoResponse
func (client *Client) QueryCommonCustInfo(request *QueryCommonCustInfoRequest) (_result *QueryCommonCustInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCommonCustInfoResponse{}
	_body, _err := client.QueryCommonCustInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCommonStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommonStatisticResponse
func (client *Client) QueryCommonStatisticWithOptions(request *QueryCommonStatisticRequest, runtime *dara.RuntimeOptions) (_result *QueryCommonStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCommonStatistic"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommonStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCommonStatisticRequest
//
// @return QueryCommonStatisticResponse
func (client *Client) QueryCommonStatistic(request *QueryCommonStatisticRequest) (_result *QueryCommonStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCommonStatisticResponse{}
	_body, _err := client.QueryCommonStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryCommonStatisticPreviewRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCommonStatisticPreviewResponse
func (client *Client) QueryCommonStatisticPreviewWithOptions(request *QueryCommonStatisticPreviewRequest, runtime *dara.RuntimeOptions) (_result *QueryCommonStatisticPreviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ProductType) {
		query["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCommonStatisticPreview"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCommonStatisticPreviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryCommonStatisticPreviewRequest
//
// @return QueryCommonStatisticPreviewResponse
func (client *Client) QueryCommonStatisticPreview(request *QueryCommonStatisticPreviewRequest) (_result *QueryCommonStatisticPreviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCommonStatisticPreviewResponse{}
	_body, _err := client.QueryCommonStatisticPreviewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryDayuMigrateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDayuMigrateStatusResponse
func (client *Client) QueryDayuMigrateStatusWithOptions(request *QueryDayuMigrateStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDayuMigrateStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryDayuMigrateStatus"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDayuMigrateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDayuMigrateStatusRequest
//
// @return QueryDayuMigrateStatusResponse
func (client *Client) QueryDayuMigrateStatus(request *QueryDayuMigrateStatusRequest) (_result *QueryDayuMigrateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDayuMigrateStatusResponse{}
	_body, _err := client.QueryDayuMigrateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryDictDataItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDictDataItemResponse
func (client *Client) QueryDictDataItemWithOptions(request *QueryDictDataItemRequest, runtime *dara.RuntimeOptions) (_result *QueryDictDataItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassifyItemCode) {
		query["ClassifyItemCode"] = request.ClassifyItemCode
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDictDataItem"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDictDataItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryDictDataItemRequest
//
// @return QueryDictDataItemResponse
func (client *Client) QueryDictDataItem(request *QueryDictDataItemRequest) (_result *QueryDictDataItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDictDataItemResponse{}
	_body, _err := client.QueryDictDataItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryGateVerifyBillingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyBillingResponse
func (client *Client) QueryGateVerifyBillingWithOptions(request *QueryGateVerifyBillingRequest, runtime *dara.RuntimeOptions) (_result *QueryGateVerifyBillingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.Month) {
		query["Month"] = request.Month
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryGateVerifyBilling"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGateVerifyBillingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryGateVerifyBillingRequest
//
// @return QueryGateVerifyBillingResponse
func (client *Client) QueryGateVerifyBilling(request *QueryGateVerifyBillingRequest) (_result *QueryGateVerifyBillingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGateVerifyBillingResponse{}
	_body, _err := client.QueryGateVerifyBillingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryGateVerifyRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyRecordListResponse
func (client *Client) QueryGateVerifyRecordListWithOptions(request *QueryGateVerifyRecordListRequest, runtime *dara.RuntimeOptions) (_result *QueryGateVerifyRecordListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNum) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.QueryDate) {
		query["QueryDate"] = request.QueryDate
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
		Action:      dara.String("QueryGateVerifyRecordList"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGateVerifyRecordListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryGateVerifyRecordListRequest
//
// @return QueryGateVerifyRecordListResponse
func (client *Client) QueryGateVerifyRecordList(request *QueryGateVerifyRecordListRequest) (_result *QueryGateVerifyRecordListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGateVerifyRecordListResponse{}
	_body, _err := client.QueryGateVerifyRecordListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryGateVerifyStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryGateVerifyStatisticResponse
func (client *Client) QueryGateVerifyStatisticWithOptions(request *QueryGateVerifyStatisticRequest, runtime *dara.RuntimeOptions) (_result *QueryGateVerifyStatisticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryGateVerifyStatistic"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryGateVerifyStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryGateVerifyStatisticRequest
//
// @return QueryGateVerifyStatisticResponse
func (client *Client) QueryGateVerifyStatistic(request *QueryGateVerifyStatisticRequest) (_result *QueryGateVerifyStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryGateVerifyStatisticResponse{}
	_body, _err := client.QueryGateVerifyStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPnsConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPnsConfigResponse
func (client *Client) QueryPnsConfigWithOptions(request *QueryPnsConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryPnsConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryPnsConfig"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPnsConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPnsConfigRequest
//
// @return QueryPnsConfigResponse
func (client *Client) QueryPnsConfig(request *QueryPnsConfigRequest) (_result *QueryPnsConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPnsConfigResponse{}
	_body, _err := client.QueryPnsConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPnsPackageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPnsPackageDetailResponse
func (client *Client) QueryPnsPackageDetailWithOptions(request *QueryPnsPackageDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryPnsPackageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPnsPackageDetail"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPnsPackageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPnsPackageDetailRequest
//
// @return QueryPnsPackageDetailResponse
func (client *Client) QueryPnsPackageDetail(request *QueryPnsPackageDetailRequest) (_result *QueryPnsPackageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPnsPackageDetailResponse{}
	_body, _err := client.QueryPnsPackageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPnsPackageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPnsPackageListResponse
func (client *Client) QueryPnsPackageListWithOptions(request *QueryPnsPackageListRequest, runtime *dara.RuntimeOptions) (_result *QueryPnsPackageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillCycle) {
		query["BillCycle"] = request.BillCycle
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryPnsPackageList"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPnsPackageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPnsPackageListRequest
//
// @return QueryPnsPackageListResponse
func (client *Client) QueryPnsPackageList(request *QueryPnsPackageListRequest) (_result *QueryPnsPackageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPnsPackageListResponse{}
	_body, _err := client.QueryPnsPackageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryPnsPackageStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPnsPackageStatisticsResponse
func (client *Client) QueryPnsPackageStatisticsWithOptions(request *QueryPnsPackageStatisticsRequest, runtime *dara.RuntimeOptions) (_result *QueryPnsPackageStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryPnsPackageStatistics"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPnsPackageStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryPnsPackageStatisticsRequest
//
// @return QueryPnsPackageStatisticsResponse
func (client *Client) QueryPnsPackageStatistics(request *QueryPnsPackageStatisticsRequest) (_result *QueryPnsPackageStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPnsPackageStatisticsResponse{}
	_body, _err := client.QueryPnsPackageStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySceneListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneListWithOptions(request *QuerySceneListRequest, runtime *dara.RuntimeOptions) (_result *QuerySceneListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QuerySceneList"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySceneListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySceneListRequest
//
// @return QuerySceneListResponse
func (client *Client) QuerySceneList(request *QuerySceneListRequest) (_result *QuerySceneListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySceneListResponse{}
	_body, _err := client.QuerySceneListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QuerySdkVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySdkVersionResponse
func (client *Client) QuerySdkVersionWithOptions(request *QuerySdkVersionRequest, runtime *dara.RuntimeOptions) (_result *QuerySdkVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.CustomerId) {
		query["customerId"] = request.CustomerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["prodCode"] = request.ProdCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySdkVersion"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySdkVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QuerySdkVersionRequest
//
// @return QuerySdkVersionResponse
func (client *Client) QuerySdkVersion(request *QuerySdkVersionRequest) (_result *QuerySdkVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySdkVersionResponse{}
	_body, _err := client.QuerySdkVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryTagStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagStatusResponse
func (client *Client) QueryTagStatusWithOptions(request *QueryTagStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryTagStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttrKey) {
		query["AttrKey"] = request.AttrKey
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
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
		Action:      dara.String("QueryTagStatus"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryTagStatusRequest
//
// @return QueryTagStatusResponse
func (client *Client) QueryTagStatus(request *QueryTagStatusRequest) (_result *QueryTagStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagStatusResponse{}
	_body, _err := client.QueryTagStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RemoveGateVerifyExportLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGateVerifyExportLogResponse
func (client *Client) RemoveGateVerifyExportLogWithOptions(request *RemoveGateVerifyExportLogRequest, runtime *dara.RuntimeOptions) (_result *RemoveGateVerifyExportLogResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
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
		Action:      dara.String("RemoveGateVerifyExportLog"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveGateVerifyExportLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveGateVerifyExportLogRequest
//
// @return RemoveGateVerifyExportLogResponse
func (client *Client) RemoveGateVerifyExportLog(request *RemoveGateVerifyExportLogRequest) (_result *RemoveGateVerifyExportLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveGateVerifyExportLogResponse{}
	_body, _err := client.RemoveGateVerifyExportLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitEnterpriseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitEnterpriseInfoResponse
func (client *Client) SubmitEnterpriseInfoWithOptions(request *SubmitEnterpriseInfoRequest, runtime *dara.RuntimeOptions) (_result *SubmitEnterpriseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessLicenseAddress) {
		query["BusinessLicenseAddress"] = request.BusinessLicenseAddress
	}

	if !dara.IsNil(request.BusinessLicensePicture) {
		query["BusinessLicensePicture"] = request.BusinessLicensePicture
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.EnterpriseName) {
		query["EnterpriseName"] = request.EnterpriseName
	}

	if !dara.IsNil(request.LegalPersonCertNumber) {
		query["LegalPersonCertNumber"] = request.LegalPersonCertNumber
	}

	if !dara.IsNil(request.LegalPersonCertPicture) {
		query["LegalPersonCertPicture"] = request.LegalPersonCertPicture
	}

	if !dara.IsNil(request.LegalPersonName) {
		query["LegalPersonName"] = request.LegalPersonName
	}

	if !dara.IsNil(request.ManagerCertNumber) {
		query["ManagerCertNumber"] = request.ManagerCertNumber
	}

	if !dara.IsNil(request.ManagerCertPicture) {
		query["ManagerCertPicture"] = request.ManagerCertPicture
	}

	if !dara.IsNil(request.ManagerContactNumber) {
		query["ManagerContactNumber"] = request.ManagerContactNumber
	}

	if !dara.IsNil(request.ManagerName) {
		query["ManagerName"] = request.ManagerName
	}

	if !dara.IsNil(request.NumberApplicationLetterPicture) {
		query["NumberApplicationLetterPicture"] = request.NumberApplicationLetterPicture
	}

	if !dara.IsNil(request.OrganizationCode) {
		query["OrganizationCode"] = request.OrganizationCode
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UndertakingPicture) {
		query["UndertakingPicture"] = request.UndertakingPicture
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitEnterpriseInfo"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitEnterpriseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitEnterpriseInfoRequest
//
// @return SubmitEnterpriseInfoResponse
func (client *Client) SubmitEnterpriseInfo(request *SubmitEnterpriseInfoRequest) (_result *SubmitEnterpriseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitEnterpriseInfoResponse{}
	_body, _err := client.SubmitEnterpriseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateExportLogStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExportLogStateResponse
func (client *Client) UpdateExportLogStateWithOptions(request *UpdateExportLogStateRequest, runtime *dara.RuntimeOptions) (_result *UpdateExportLogStateResponse, _err error) {
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

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExportLogState"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExportLogStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateExportLogStateRequest
//
// @return UpdateExportLogStateResponse
func (client *Client) UpdateExportLogState(request *UpdateExportLogStateRequest) (_result *UpdateExportLogStateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExportLogStateResponse{}
	_body, _err := client.UpdateExportLogStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateSmsCodeLimitConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmsCodeLimitConfigResponse
func (client *Client) UpdateSmsCodeLimitConfigWithOptions(request *UpdateSmsCodeLimitConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmsCodeLimitConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LimitDay) {
		query["LimitDay"] = request.LimitDay
	}

	if !dara.IsNil(request.LimitHour) {
		query["LimitHour"] = request.LimitHour
	}

	if !dara.IsNil(request.LimitId) {
		query["LimitId"] = request.LimitId
	}

	if !dara.IsNil(request.LimitMinute) {
		query["LimitMinute"] = request.LimitMinute
	}

	if !dara.IsNil(request.LimitOther) {
		query["LimitOther"] = request.LimitOther
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmsCodeLimitConfig"),
		Version:     dara.String("2017-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmsCodeLimitConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateSmsCodeLimitConfigRequest
//
// @return UpdateSmsCodeLimitConfigResponse
func (client *Client) UpdateSmsCodeLimitConfig(request *UpdateSmsCodeLimitConfigRequest) (_result *UpdateSmsCodeLimitConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSmsCodeLimitConfigResponse{}
	_body, _err := client.UpdateSmsCodeLimitConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
